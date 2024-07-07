import {Link, Outlet, useNavigate} from "react-router-dom";
import {useCallback, useEffect, useState} from "react";
import apiBase from "/src/index.jsx";
import Alert from "/src/components/Alert.jsx";

const App = () => {

	const [jwtToken, setJwtToken] = useState("");
	const [alertInfo, setAlertInfo] = useState({className: "d-none", message: ""});
	const [tickInterval, setTickInterval] = useState();

	const navigate = useNavigate();

	const fetchRefresh = async () => {
		const requestOptions = {
			method: "GET",
			credentials: "include",
		}

		try {
			const res = await fetch(apiBase + "/refresh", requestOptions)
			const data = await res.json();
			if (data.data.access_token) {
				console.log(data.data.access_token)
				setJwtToken(data.data.access_token)
			} else {
				console.log("no access token found", data.message);
			}
		} catch (error) {
			console.log("error logging in", error.message)
		}
	}

	const toggleRefresh = useCallback(status => {
		if (status) {
			let i = setInterval(() => {
				fetchRefresh()
				setTickInterval(i)
			}, 10 * 60 * 1000) // 10 minutes
		} else {
			clearInterval(tickInterval)
		}
	}, [tickInterval])

	const logout = async () => {
		const requestOptions = {
			method: "GET",
			credentials: "include",
		}

		try {
			await fetch(apiBase + "/logout", requestOptions)
		} catch (error) {
			console.log("error logging out", error.message)
		}

		toggleRefresh(false)
		setJwtToken("")
		navigate("/login")
	}

	useEffect(() => {
		if (jwtToken === "") {
			fetchRefresh()
		}
	}, [jwtToken, toggleRefresh]);


	return (
		<div className="container">
			<div className="row">
				<div className="col">
					<h1 className="mt-3">Workout Tracker</h1>
				</div>
				<div className="col text-end mt-3 p-0">
					{jwtToken !== "" &&
						<a href={"#!"} onClick={logout}>
							<span className="badge bg-danger">Logout</span>
						</a>}
				</div>
				<hr className="mb-3"/>
			</div>

			<div className="row">

				<div className="col-md-2">
					<nav>
						<div className="list-group">
							{jwtToken === ""
								? (
									<>
										<Link to={"/login"} className={"list-group-item list-group-item-action"}>Login</Link>
										<Link to={"/register"} className={"list-group-item list-group-item-action"}>Register</Link>
									</>
								) : (
									<>
										<Link to={"/"} className={"list-group-item list-group-item-action"}>Dashboard</Link>
									</>
								)}
						</div>
					</nav>
				</div>

				<div className="col-md-10">
					<Alert message={alertInfo.message} className={alertInfo.className} />
					<Outlet context={{jwtToken, setJwtToken, setAlertInfo, toggleRefresh}}/>
				</div>

			</div>
		</div>
	)
}

export default App





















import {useState} from "react";
import {useNavigate, useOutletContext} from "react-router-dom";
import Input from "/src/components/form/Input";
import apiBase from "/src/index.jsx";

const Login = () => {

	const [credentials, setCredentials] = useState({email: "", password: ""})
	const {setJwtToken, setAlertInfo, toggleRefresh} = useOutletContext()

	const navigate = useNavigate();

	const handleSubmit = async e => {
		e.preventDefault();

		const payload = {
			email: credentials.email,
			password: credentials.password,
		}

		const headers = new Headers({'Content-Type': 'application/json'})

		const requestOptions = {
			method: "POST",
			headers: headers,
			credentials: "include",
			body: JSON.stringify(payload),
		}

		try {
			const res = await fetch(apiBase + "/authenticate", requestOptions)
			const data = await res.json()
			if (data.error) {
				setAlertInfo({message: data.message, className: "alert alert-danger"})
			} else {
				setJwtToken(data.data.access_token)
				setAlertInfo({message: "", className: "d-none"})
				toggleRefresh(true)
				navigate("/")
			}
		} catch (error) {
			setAlertInfo({message: error.message, className: "alert alert-danger"})
		}
	}

	return (
		<div className="container">
			<div className="row">
				<div className="col-md-6 offset-md-2">
					<h2 className="mt-3 mb-3">Login</h2>
					<form onSubmit={handleSubmit}>
						<Input
							title={"Email"}
							type={"email"}
							className={"form-control"}
							name={"email"}
							autoComplete="email-new"
							onChange={(e) => setCredentials({...credentials, email: e.target.value})}
							// TODO: change placeholder CSS
							placeholder={"admin@example.com"}
							required={true}
						/>
						<Input
							title={"Password"}
							type={"password"}
							className={"form-control"}
							name={"password"}
							autoComplete="password-new"
							onChange={(e) => setCredentials({...credentials, password: e.target.value})}
							// TODO: change placeholder CSS
							placeholder={"password"}
							required={true}
						/>
						<input
							type={"submit"}
							className={"btn btn-primary"}
							value={"Login"}
						/>
					</form>
				</div>
			</div>
		</div>
	)
}

export default Login
import {useState} from "react";
import {useNavigate, useOutletContext} from "react-router-dom";
import Input from "/src/components/form/Input.jsx";
import apiBase from "/src/index.jsx";

const Register = () => {

	const [userData, setUserData] = useState({
		email: "",
		password: "",
		name: "",
		patronymic: "",
		surname: "",
		weight: 0.0,
		height: 0.0,
		dob: "",
	})

	const {setAlertInfo, setJwtToken, toggleRefresh} = useOutletContext()

	const navigate = useNavigate()

	const fetchRegister = async () => {
		let requestOptions = {
			method: "POST",
			headers: {'Content-Type': 'application/json'},
			body: JSON.stringify(userData)
		}

		try {
			const res = await fetch(apiBase + "/register", requestOptions)
			const data = await res.json()
			if (data.error) {
				setAlertInfo({message: data.message, className: "alert alert-danger"})
				return false
			}
		} catch (error) {
			setAlertInfo({message: error.message, className: "alert alert-danger"})
			return false
		}

		return true
	}

	const fetchLogin = async () => {
		const payload = {
			email: userData.email,
			password: userData.password,
		}

		const requestOptions = {
			method: "POST",
			headers: {'Content-Type': 'application/json'},
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
				setAlertInfo({message: "New user created", className: "alert alert-success"})
				setTimeout(() => {
					setAlertInfo({message: "", className: "d-none"})
				}, 3 * 1000)
				toggleRefresh(true)
				navigate("/")
			}
		} catch (error) {
			setAlertInfo({message: error.message, className: "alert alert-danger"})
		}
	}

	const handleSubmit = async e => {
		e.preventDefault();

		if (await fetchRegister()) {
			fetchLogin()
		}
	}

	return (
		<div className="container">
			<div className="row">
				<div className="col-md-6 offset-md-2">
					<h2 className="mt-3 mb-4">Register new user</h2>
					<form onSubmit={handleSubmit}>
						<Input
							title={"First Name"}
							type={"text"}
							name={"name"}
							autoComplete={"name-new"}
							placeholder={"John"}
							onChange={e => setUserData({...userData, name: e.target.value})}
							value={userData.name}
							className={"form-control"}
							required={true}
						/>
						<Input
							title={"Patronymic"}
							type={"text"}
							name={"patronymic"}
							autoComplete={"patronymic-new"}
							placeholder={"Ivanovich"}
							onChange={e => setUserData({...userData, patronymic: e.target.value})}
							value={userData.patronymic}
							className={"form-control"}
						/>
						<Input
							title={"Last Name"}
							type={"text"}
							name={"surname"}
							autoComplete={"surname-new"}
							placeholder={"Doe"}
							onChange={e => setUserData({...userData, surname: e.target.value})}
							className={"form-control mb-5"}
							value={userData.surname}
							required={true}
						/>
						<Input
							title={"Email Address"}
							type={"email"}
							name={"email"}
							autoComplete={"email-new"}
							placeholder={"admin@example.com"}
							onChange={e => setUserData({...userData, email: e.target.value})}
							value={userData.email}
							className={"form-control"}
							required={true}
						/>
						<Input
							title={"Password"}
							type={"password"}
							name={"password"}
							autoComplete={"password-new"}
							placeholder={"password"}
							onChange={e => setUserData({...userData, password: e.target.value})}
							value={userData.password}
							className={"form-control mb-5"}
							required={true}
						/>
						<Input
							title={"Weight"}
							type={"number"}
							name={"weight"}
							placeholder={"87.5"}
							onChange={e => setUserData({...userData, weight: e.target.value})}
							value={userData.weight}
							className={"form-control"}
						/>
						<Input
							title={"Height"}
							type={"number"}
							name={"height"}
							placeholder={"183.5"}
							onChange={e => setUserData({...userData, height: e.target.value})}
							value={userData.height}
							className={"form-control"}
						/>
						<Input
							title={"Date of Birth"}
							type={"date"}
							name={"dob"}
							onChange={e => setUserData({...userData, dob: e.target.value})}
							value={userData.dob}
							className={"form-control"}
						/>
						<input
							type={"submit"}
							className={"btn btn-primary mt-3"}
							value={"Submit"}
						/>
					</form>
				</div>
			</div>
		</div>
	)
}

export default Register
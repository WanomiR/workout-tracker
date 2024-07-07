import {useState} from "react";

const Register = () => {

	const [userData, setUserData] = useState({
		email: "",
		password: "",
		name: "",
		patronymic: "",
		surname: "",
		weight: "",
		height: "",
		dob: "",
	})

	const handleSubmit = async e => {
		e.preventDefault();
	}

	return (
		<div className="container">
			<div className="row">
				<div className="col-md-6 offset-md-2">
					<h2 className="mt-3 mb-3">Register</h2>
					<form onSubmit={handleSubmit}>

					</form>
				</div>
			</div>
		</div>
	)
}

export default Register
import {useEffect} from "react";
import {useNavigate, useOutletContext} from "react-router-dom";

const Dashboard = () => {

	const {jwtToken} = useOutletContext()

	const navigate = useNavigate()

	useEffect(() => {
		if (jwtToken === "") {
			navigate("/login")
		}
	}, []);

	return (
		<div className="container">
			<div className="row">
				<div className="col">
					<h2 className="mt-3">Dashboard</h2>
				</div>
			</div>
		</div>
	)
}

export default Dashboard
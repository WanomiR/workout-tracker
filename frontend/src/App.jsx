import {Link} from "react-router-dom";


const App = () => {

	return (
		<div className="container">
			<div className="row">
				<div className="col">
					<h1 className="mt-3">Workout Tracker</h1>
				</div>
				<div className="col text-end mt-3 p-0">
						<Link to={"/login"}>
							<span className="badge bg-success">Login</span>
						</Link>
				</div>
				<hr className="mb-3"/>
			</div>
		</div>
	)
}

export default App
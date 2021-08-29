import React from 'react'
import "./mainBody.css"
import { useState } from 'react'
import { baseUrl } from '../../baseUrl'
import Card from '../Card'





export default function MainBody() {
	const [searchName, setSearchName] = useState("")
	const [isSelected, setIsSelected] = useState(false)
	const [dropDownItems, setDropDownItems] = useState(["hello", "hi"])
	const [selectedDropDownItems, setSelectedDropDownItems] = useState([])

	const [isLoading, setIsLoading] = useState(false)
	const [renderValues, setRenderValues] = useState([])
	const render_cards = (data) => {

		var lis = data.map((elem) => {


			var title = "Html Source"
			var url = elem._source.url
			var lastaccessed = elem._source.lastaccessed
			var dom = elem.highlight.dom[0]

			return {
				title, url, lastaccessed, dom

			}


		})
		if (lis.length === 0) {
			lis = [{ title: "No data found", url: "", lastaccessed: "", dom: "" }]
		}
		return lis


	}

	function search() {



		setIsLoading(true)

		fetch(`${baseUrl}/search/${searchName}`)
			.then(response => response.json())
			.then(data => {
				setIsLoading(false)
				setRenderValues(render_cards(data))

			}

			);

	}

	const handleKeyDown = (e) => {
		if (e.key === 'Enter') {
			search()
		}
	}
	const makeOptions = (lis) => {
		console.log("hello there", lis)
		var obj = lis.map((x, index) => <button className="drop-down-item"
			key={x}

			onClick={() => setSelectedDropDownItems(a => {

				console.log(a, "this is a")
				var lis = a
				lis.push(x)
				setDropDownItems(b => {

					var new_lis = []
					b.forEach((val, ind) => ind != index && new_lis.push(val))
					return new_lis
				})
				return lis

			})}
		>{x}

		</button>
		)
		console.log(obj)
		return <div className="drop-downs"


			onPointerLeave={() => setIsSelected(false)}
		>{obj}</div>

	}

	return (
		<div className="main-container">
			<div className="top-container">
				<input type="text"
					className="search-bar"
					placeholder="Seach any item"
					onKeyPress={handleKeyDown}
					value={searchName}
					onChange={({ target }) => setSearchName(target.value)}


				></input>
				<div className="select-box"
					onClick={() => setIsSelected(x => !x)}
				>

					<button
						className="select"
					>


						{selectedDropDownItems.length < 2 ? `${selectedDropDownItems.length} item selected` : `${selectedDropDownItems.length} items selected`}



					</button>
					{isSelected && makeOptions(dropDownItems)}
				</div>
				<button className="search-button" onClick={() => search()}>Search</button>
			</div>
			<div className="selected-container">
				{selectedDropDownItems.map((x, index) => <div className="bottom-selected-items"
				>{x}
					<button className="cross-icon"

						onClick={() => setDropDownItems(a => {

							console.log(a, "this is a")
							var lis = a
							lis.push(x)
							setSelectedDropDownItems(b => {

								var new_lis = []
								b.forEach((val, ind) => ind != index && new_lis.push(val))
								return new_lis
							})
							return lis

						})}
					>x</button></div>)}
			</div>
			<div className="card-container">

				{isLoading ? <div className="loading">Loading...</div> :
					renderValues.map(x =>

						<Card
							title={x.title}
							innerText={x.dom}
							date={x.lastaccessed}
							link={x.url}

						/>
					)



				}


			</div>
		</div>
	)
}

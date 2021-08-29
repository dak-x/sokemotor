import React from 'react'
import "./mainBody.css"
import { useState } from 'react'

import Card from '../Card'






export default function MainBody() {
	const [searchName, setSearchName] = useState("")

	const [isSelected, setIsSelected] = useState(false)
	const [dropDownItems, setDropDownItems] = useState(["hello", "hi"])
	const [selectedDropDownItems, setSelectedDropDownItems] = useState([])

	const [renderValues,setRenderValues] = useState([])

	function search(){
		var objLis=[
			{
			
				"lastaccessed":"date-time",
	"title":`${searchName}`,
	"url":"google.com",
	"dom":"hkuhlkuhkg",
		
		},
			{
			
				"lastaccessed":"date-time",
	"title":"something",
	"url":"google.com",
	"dom":"hkuhlkuhkg",
		
		},
			{
			
				"lastaccessed":"date-time",
	"title":"something",
	"url":"google.com",
	"dom":"hkuhlkuhkg",
		
		},
		]
		setRenderValues(objLis)
	}
	
	function makeOptions(lis) {
		console.log("hello there", lis)
		var obj = lis.map((x,index) => <button className="drop-down-item"
			key={x}
			
			onClick={()=>setSelectedDropDownItems(a=>{
			
				console.log(a,"this is a")
				var lis = a
				lis.push(x)
				setDropDownItems(b=>{

					var new_lis = []
					b.forEach((val,ind) =>ind!=index&&new_lis.push(val))
					return new_lis
				})
			return 	lis
			
			})}
			>{x}
			
		</button>
		)
		console.log(obj)
		return <div className="drop-downs"


			onPointerLeave={()=> setIsSelected(false)}
		>{obj}</div>

	}

	return (
		<div className="main-container">
			<div className="top-container">
				<input type="text"
					className="search-bar"
					placeholder="Seach any item"
					value={searchName}
					onChange={({ target }) => setSearchName(target.value)}


				></input>
				<div className="select-box"
				onClick={()=>setIsSelected(x=>!x)}
				>
					
					<button
						className="select"
					>
						
						
						{selectedDropDownItems.length<2?`${selectedDropDownItems.length} item selected`:`${selectedDropDownItems.length} items selected`}
						
						
						
						</button>
					{isSelected && makeOptions(dropDownItems)}
				</div>
				<button className="search-button" onClick={()=>search()}>Search</button>
			</div>
			<div style={{display:"flex"}}>
			{selectedDropDownItems.map((x,index) => <div className="bottom-selected-items"
			>{x}
			<button className="cross-icon"

			onClick={()=>setDropDownItems(a=>{
			
				console.log(a,"this is a")
				var lis = a
				lis.push(x)
				setSelectedDropDownItems(b=>{

					var new_lis = []
					b.forEach((val,ind) =>ind!=index&&new_lis.push(val))
					return new_lis
				})
			return 	lis
			
			})}
			>x</button></div>)}
			</div>
<div className="card-container">
			{renderValues.map(x =>
				
<Card
title={x.title}
innerText={x.dom}
date={x.lastaccessed}
link={x.url}

/>
				)}
				</div>
		</div>
	)
}

import React from 'react'
import "./card.css"
export default function Card({title,innerText,date,link}) {
	return (
		<div className="card">
		<div className="title"	>{title}</div>
		<div className="inner-text" div dangerouslySetInnerHTML={{ __html: innerText }}>
			
		</div>
		<div className="flex-container">
		<div className="last-visited"> {date}</div>
		<a className="link" href={link}>visit</a>
		</div>
		</div>
	)
}

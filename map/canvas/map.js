(() => {
	fetch('/shapes.json')
		.then(r => r.json())
		.then(shapes => { draw(shapes) })

	const canvas = document.getElementById('map')
	const ctx = canvas.getContext('2d')
	canvas.width = canvas.offsetWidth;
	canvas.height = canvas.offsetHeight;
	const height = canvas.height
	const width = canvas.width

	ctx.fillStyle = "rgb(255,255,255)"
	ctx.fillRect(0, 0, width, height)

	function draw (shapes) {
		shapes.map(s => {
			ctx.beginPath()

			s.map(p => {
				ctx.lineTo(getLon(p.lon), getLat(p.lat))
			})

			ctx.strokeStyle = `rgba(30,30,30)`
			ctx.lineWidth = 1
			ctx.stroke()
		})

		fetch('/ip.loc.json')
			.then(r => r.json())
			.then(ips => { place(ips) })
	}

	function place (ips) {
		const maxCount = parseInt(ips[0].count)
		ips.map(ip => {
			ctx.beginPath()

			const radius = parseInt(ip.count) * 100 / maxCount
			ctx.arc(getLon(ip.loc.split(',')[1]), getLat(ip.loc.split(',')[0]), radius, 0, 2 * Math.PI)

			ctx.strokeStyle = `rgba(255,${-(radius * 2.55) + 255},0)`
			ctx.lineWidth = 1
			ctx.stroke()
		})

		const dataURL = canvas.toDataURL()
		let img = document.createElement('img')
		img.src = dataURL
		document.body.append(img)
	}

	function getLat(lat) {
		const rawLat = parseFloat(lat) + 90
		let calcLat = 0

		if (rawLat === 0) {
			calcLat = 0
		} else if (rawLat > 180) {
			calcLat = 0
		} else {
			calcLat = (height - (rawLat * (height / 180)))
		}
	
		return calcLat
	}
	
	function getLon(lon) {
		const rowLon = parseFloat(lon) + 180
		let calcLon = 0

		if (rowLon === 0) {
			calcLon = 0
		} else if (rowLon > 360) {
			calcLon = width
		} else {
			calcLon = (rowLon * (width / 360))
		}
	
		return calcLon
	}

})()

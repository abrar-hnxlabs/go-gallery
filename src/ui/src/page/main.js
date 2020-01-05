import React from "react";
import { Container, Header, Segment, Divider, Image } from 'semantic-ui-react'


class MainPage extends React.Component {

    constructor(props) {
        super(props)
        this.state = {}
        this.fetchData()
    }

    fetchData = async () =>{
        let response = await fetch('/api/data')
        let json = await response.json()
        this.reformat(json.data)
    }

    reformat = (data) => {
        let master = {}
        let years = data.map((d) => d.year)
        let uniqueYears = new Set(years)
        uniqueYears.forEach((y) => {
            master[y] = {}
            let months = data.filter((d) => d.year === y).map(f => f.month)
            let uniqueMonths = new Set(months)
            uniqueMonths.forEach((m) => {
                master[y][m] = {}
                let groups = data.filter((d) => d.year === y && d.month === m).map(f => f.dir)
                let uniqueGroups = new Set(groups)
                uniqueGroups.forEach(g => {
                    let images = data.filter(d => d.year === y && d.month === m && d.dir === g);
                    master[y][m][g] = images
                })
            })
        })
    
        this.setState({ data: master})
    }

    toMonth = (monthNumber) => {
        switch(parseInt(monthNumber)){
            case 1:
                return "Jan"
            case 2:
                return "Feb"
            case 3:
                return "Mar"
            case 4:
                return "Apr"
            case 5:
                return "May"
            case 6:
                return "Jun"
            case 7:
                return "Jul"
            case 8:
                return "Aug"
            case 9:
                return "Sep"
            case 10:
                return "Oct"
            case 11:
                return "Nov"
            case 12:
                return "Dec"
            default:
                return "Unknown"
        }

    }
    renderImages = (images) => {
        return images.map(i => <Image key={i.thumbnail} src={i.thumbnail} bordered />)
    }
    renderGroups = (groups) => {
        return Object.keys(groups).map(g => (
            <>
                <Segment key={g}>
                    <Header sub>Count {groups[g].length}</Header>
                    <Image.Group size='small'>
                            {this.renderImages(groups[g])}
                    </Image.Group> 
                </Segment>
                <Divider />
            </>
        ))
    }

    renderMonths = (data) => {
        return Object.keys(data).map(m => <><Header sub>{this.toMonth(m)}</Header>{this.renderGroups(data[m])}</>)
    }

    renderYears = (data) => {
        return Object.keys(data).map((y) => (<>
            <Header as='h2' key={y} >{y}</Header>
            {this.renderMonths(data[y])}
            </>)
        )
    }
    render = () => {
    
        const { data } = this.state
        if(!data) {
            return (
                <Container>
                    <Segment><Header as='h2'>Loading Data...</Header></Segment>
                </Container>
            )
        }
        return(
        <Container>
            { this.renderYears(data)}
        </Container>)
    }
}

export default MainPage;
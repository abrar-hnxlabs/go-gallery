import React from "react";
import { Segment, Container, H2, Overlay,FullScreenImage} from '../components/controls'
import Section from '../components/section'
import styled from 'styled-components'

const ImgContainer = styled.div`
    position: fixed;
    width: 100%;
    top: 0;
    left: 0;
    height: 100%;
    z-index: 1100;
    margin: 0 auto;
`

const CenteredContainer = styled.div`
    display: block;
    width: 1200px;
    margin: 50px auto;
`

class MainPage extends React.Component {

    
    constructor(props) {
        super(props)
        this.state = { showOverlay: false, photo: null}
        this.fetchData()
    }

    fetchData = async () =>{
        let response = await fetch('/api/data')
        let json = await response.json()
        this.reformat(json.data)
    }

    reformat = (data) => {
        let master = []
        let years = data.map((d) => d.year)
        let uniqueYears = new Set(years)
        uniqueYears.forEach((y) => {
            let tempData = { year: y}
            let months = data.filter((d) => d.year === y).map(f => f.month)
            let uniqueMonths = new Set(months)
            uniqueMonths.forEach((m) => {
                tempData.month = m
                let groups = data.filter((d) => d.year === y && d.month === m).map(f => f.dir)
                let uniqueGroups = new Set(groups)
                uniqueGroups.forEach(g => {
                    let images = data.filter(d => d.year === y && d.month === m && d.dir === g);
                    tempData.images = images
                })
            })
            master.push(tempData)
        })
        this.setState({ data: master})
    }

    imageOnClick = (photo) => {
        this.setState({ showOverlay: true, photo })
    }

    closeOverlay = () => {
        this.setState({ showOverlay: false, photo: null })
    }
    render = () => {
    
        const { data, showOverlay, photo } = this.state
        if(!data) {
            return (
                <Container>
                    <Segment><H2>Loading...</H2></Segment>
                </Container>
            )
        }
        return(
        <>
        <Container>
            { data.map(m => <Section key={m.year+'-'+m.month}year={m.year} month={m.month} images={m.images} onClick={this.imageOnClick} />)}
        </Container>
        <Overlay show={showOverlay} />
        {photo && <ImgContainer><CenteredContainer><FullScreenImage src = {photo}  onClick={()=> this.closeOverlay()} /></CenteredContainer></ImgContainer>}
        </>
        )
    }
}

export default MainPage;
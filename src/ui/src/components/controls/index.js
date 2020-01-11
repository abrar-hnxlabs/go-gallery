import styled from 'styled-components'

const Segment = styled.div`
    background-color: #2E1114;
    color: #fff;
    width: 100%;
    height: auto;
    padding: 5px;
    overflow:auto;
`

const FullScreenImage = styled.img`
    width: 100%;
    max-height: 90%;
    margin: 3px;
    cursor: pointer;
    margin: 0 auto;
`

const Image = styled.img`
    height: 150px !important;
    width: 150px;
    margin: 3px;
    cursor: pointer;
`
const H2 = styled.h2`
    color: #83677B;
    margin: 2px;
    font-family: 'Open Sans', sans-serif;
`

const H3 = styled.h3`
    color: #64485C;
    margin: 1px;
    font-family: 'Open Sans', sans-serif;
`

const ImageGroup = styled.div`
    display: block;
    float: left;
    clear: both;
`

const Container = styled.div`
    display: block;
    width: 1280px;
    padding 1px;
    margin: 0 auto;
`

const Overlay = styled.div`
    position: fixed;
    top: 0;
    bottom: 0;
    width: 100%;
    display: ${props => props.show ? "block" : "none"}
    z-index: 1000;
    height: 100%;
    opacity:0.7;
    background-color:#000;
    left: 0;
`
export {Segment, Image, H2, H3, Container, ImageGroup, Overlay, FullScreenImage}
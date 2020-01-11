import React from "react";
import { toMonth } from './utils'
import { Segment, Image, H2, H3, ImageGroup } from './controls'

export default ({ year, month, images, onClick}) => (
    <Segment key={year}>
        <H2>{year}</H2>
        <H3>{toMonth(month)}</H3>
        <ImageGroup>
            { images.map(i => <Image src={i.thumbnail} onClick={() => {onClick(i.file)}} />) }
        </ImageGroup> 
    </Segment>
)

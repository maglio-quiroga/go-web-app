import ImageGallery from "react-image-gallery"
import "react-image-gallery/styles/css/image-gallery.css"

export const GaleriaImagenes = () => {
    const imagenes = [
        {
            original : "https://picsum.photos/id/1018/1920/1000",
            thumbnail : "https://picsum.photos/id/1018/250/150",
        },
        {
            original : "https://picsum.photos/id/1015/1920/1000",
            thumbnail : "https://picsum.photos/id/1015/250/150",
        },
        {
            original : "https://picsum.photos/id/1016/1920/1000",
            thumbnail : "https://picsum.photos/id/1016/250/150",
        },
    ]
    return (
            <div style={{marginTop : "86px"}}>
                <ImageGallery
                    items={imagenes}
                    showPlayButton={false}
                    showThumbnails={false}
                />
            </div>
    )
}
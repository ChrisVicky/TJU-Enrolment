from fastapi import FastAPI, File, UploadFile
import ddddocr

app = FastAPI()


@app.post("/uploadfile/")
async def create_upload_file(image: UploadFile = File(...)):
    ocr = ddddocr.DdddOcr(show_ad=False)
    contents = await image.read()
    result = ocr.classification(contents)
    return {"data": result, "code": 200}

package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "course/image/ai_video_1.png#@#course/image/ai_video_2.png#@#course/image/ai_video_3.png#@#course/image/ai_video_4.png#@#course/image/ai_video_5.png"

	s := strings.Split(s1, "#@#")
	fmt.Println("---->", s)
}

// // 转成slice
// courseP.IntroImageList = strings.Split(courseP.IntroImages, ReadingSplit)
// fmt.Println("---->", courseP.IntroImageList)
// courseP.IntroVideoImageList = strings.Split(courseP.IntroVideoImage, ReadingSplit)
// fmt.Println("---->", courseP.IntroVideoImageList)
// courseP.IntroList = strings.Split(courseP.Intro, ReadingSplit)
// fmt.Println("---->", courseP.IntroList)
// courseP.ExpertList = strings.Split(courseP.Expert, ReadingSplit)
// fmt.Println("---->", courseP.ExpertList)
// courseP.ArchitectureList = strings.Split(courseP.Architecture, ReadingSplit)
// fmt.Println("=====>", courseP)

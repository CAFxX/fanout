package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{("hello" + "foo bar"), "hello", "", ("\n\t" + ("\n\t" + "a"))}
    v3 := []int{128, 4294967295, 275}
    v4 := GenS1{F0: 42, F1: (0 >= 640)}
    v3 = append(v3, 259)
    v5 := GenS2{F0: 255, F1: ("hello" + (("\n\t" + ((("foo bar" + ("hello" + (("\n\t" + "hello") + "\n\t"))) + "") + (("hello" + "foo bar") + "hello"))) + "a"))}
    v2 = append(v2, "")
    var v6 float32 = (1e10 / 1.0)
    v6 += float32(-2.25)
    fmt.Println(v6)
    var v7 uint8 = 255
    if (("foo bar" == "") && false) {
        fmt.Println(v7, v7)
    } else {
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
    }
    var v8 uint32 = 0
    v8 += 42
    v6 = float32(0.0)
    fmt.Println(v6, v6)
    var v9 uint64 = 0
    delete(v1, (("hello" + "foo bar") + "\n\t"))
    v10 := GenS3{F0: (v7 / (((v7 | v7)) | 1))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

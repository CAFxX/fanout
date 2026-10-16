package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int }
func genFn0() uint16 {
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: -834, F1: "a"}
        var v9 bool = false
        _ = fi7
        _ = v8
        _ = v9
    }
    switch 2147483647 {
    case 128:
        var v10 bool = (false || false)
        _ = v10
    }
    if true {
        switch "" {
        case "\n\t":
            v11 := []int{32768, -232}
            _ = v11
        case "a":
            var v12 float32 = 0.0
            _ = v12
        }
    }
    if (true && true) {
        var v13 uint16 = 226
        _ = v13
    }
    return 255
}

func main() {
    var v1 uint = 1
    v2 := make(map[int]uint32)
    v3 := []string{"foo bar", "", ("" + "\n\t")}
    var v4 int64 = 1
    var v5 float32 = 1.5
    var v6 int32 = -1
    fmt.Println(v1, v6, v4)
    v5 += (v5 / -2.25)
    for fi14 := 0; fi14 < 11; fi14++ {
        v5 = 1e10
        _ = fi14
    }
    fmt.Println(v5)
    var v15 uint = (v1 % (((v1 - 1)) | 1))
    v16 := GenS2{F0: false}
    for fi17 := 0; fi17 < 20; fi17++ {
        switch fi17 {
        case -231:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case -9223372036854775808:
            delete(v2, 32768)
        }
        v18 := GenS3{F0: 1, F1: fi17}
        _ = fi17
        _ = v18
    }
    var v19 int64 = (v4 ^ v4)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

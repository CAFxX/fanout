package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint16; F1 int; F2 string }
type GenS3 struct { F0 int16; F1 int32; F2 uint8 }

func main() {
    var v1 int64 = 1
    var v2 int16 = -1
    v3 := []int{}
    v4 := make(map[string]int)
    var v5 uint16 = 2
    if false {
        _ = len(v4)
    } else {
        if false {
            switch "foo bar" {
            case "foo bar":
                v5 = (v5 - 799)
            case "":
                v5 = (v5 % ((v5) | 1))
            case "\n\t":
                var v6 uint = 255
                _ = v6
            }
            v7 := GenS1{F0: 4294967295, F1: (v5 + 42), F2: 42}
            var v8 int32 = 2147483647
            _ = v7
            _ = v8
        }
        fmt.Println(v1)
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        if true {
            fmt.Println(v2)
            switch (fi9 % (((-(fi9))) | 1)) {
            case 2147483647:
                v2 = (v2 | (v2 & v2))
            default:
                if false {
                    v10 := GenS2{F0: (v5 * v5), F1: fi9, F2: (("" + "hello") + "a")}
                    _ = v10
                } else {
                    v2 -= (v2 + v2)
                }
            }
        }
        var v11 string = ""
        _ = fi9
        _ = v11
    }
    var v12 float32 = 0.0
    var v13 int32 = -2147483648
    var v14 bool = false
    _ = len(v3)
    _ = len(v3)
    if (42 == 255) {
        v5 = (v5 % (((v5 + v5)) | 1))
        for fi15 := 0; fi15 < 15; fi15++ {
            v14 = (true && v14)
            v4["a"] = (-(fi15))
            fmt.Println(v12, fi15)
            _ = fi15
        }
    }
    v16 := GenS2{F0: (v5 >> ((130) & 63)), F1: 128, F2: ""}
    for fi17 := 0; fi17 < 20; fi17++ {
        v18 := GenS3{F0: v2, F1: (v13 >> ((493) & 63)), F2: 1}
        _ = fi17
        _ = v18
    }
    fmt.Println(v14, v14, v13)
    v4[(("foo bar" + "foo bar") + "hello")] = 504
    var v19 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS4 struct { F0 int8; F1 int32; F2 uint32 }

func main() {
    var v1 string = ""
    v2 := []string{(v1 + v1), ("\n\t" + "a")}
    var v3 int = 443
    v4 := make(map[uint32]string)
    if (0 != 0) {
        var v5 int = v3
        _ = v5
    }
    v1 = "hello"
    if false {
        if true {
            var v6 float64 = float64(float32(1.5))
            v6 = (0.1 / 1.0)
            _ = v6
        }
    } else {
        var v7 uint64 = 0
        _ = v7
    }
    switch v1 {
    case "":
        v8 := GenS1{F0: 892}
        _ = v8
    case "\n\t":
        for fi9 := 0; fi9 < 2; fi9++ {
            if false {
                if false {
                    if (1e10 < 1.5) {
                        v10 := GenS2{F0: 42}
                        v11 := GenS3{F0: 375, F1: 82, F2: 210}
                        fmt.Println(fi9, v1)
                        _ = v10
                        _ = v11
                    }
                    fmt.Println(v3, v3, v3)
                } else {
                    fmt.Println(v3, fi9)
                    v12 := GenS4{F0: 57, F1: 241, F2: 4294967295}
                    _ = v12
                }
            }
            v13 := v4[1]
            _ = len(v4)
            _ = fi9
            _ = v13
        }
    }
    fmt.Println(v3)
    var v14 bool = true
    delete(v4, 0)
    var v15 float64 = ((0.1 * (0.1 + float64(1.5))) + (((-2.25 * (-2.25 - ((1e10 + 0.0) - 1e10))) / -2.25) + 3.14159))
    var v16 int32 = 2147483647
    delete(v4, 1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}

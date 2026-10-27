package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 bool; F1 int }
func genFn0(p0 uint, p1 uint8) int32 {
    p0 = p0
    var v4 bool = (1 >= 4294967295)
    v4 = (0 <= 0)
    for fi5 := 0; fi5 < 9; fi5++ {
        for fi6 := 0; fi6 < 16; fi6++ {
            var v7 int16 = 833
            _ = fi6
            _ = v7
        }
        if v4 {
            switch ("a" + "\n\t") {
            case "hello":
                v4 = (42 >= -212)
            }
        } else {
            p1 += (p1 + (p1 >> ((p0) & 63)))
        }
        if (p0 >= (p0 - (^(p0)))) {
            if (!v4) {
                v8 := GenS1{F0: fi5}
                v9 := GenS1{F0: (fi5 * (fi5 | 2147483647))}
                switch "foo bar" {
                case "foo bar":
                    var v10 int64 = 1
                    _ = v10
                default:
                    v11 := GenS1{F0: fi5}
                    _ = v11
                }
                _ = v8
                _ = v9
            } else {
                var v12 float32 = 0.0
                _ = v12
            }
            p0 += (p0 % (((p0 ^ p0)) | 1))
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = v4
    return -637
}

func main() {
    var v1 uint = 155
    var v2 int8 = -92
    var v3 bool = (false && true)
    v13 := []int{128, 128}
    var v14 string = ("foo bar" + ("hello" + "a"))
    v3 = (!v3)
    v15 := GenS2{F0: v14, F1: 42}
    v13 = append(v13, -82)
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 int8 = v2
        v17 = v2
        var v18 int16 = 1
        _ = fi16
        _ = v17
        _ = v18
    }
    v2 = v2
    var v19 int16 = -32768
    fmt.Println(v1, v2, v1)
    v1 = 375
    for fi20 := 0; fi20 < 18; fi20++ {
        if false {
            var v21 uint8 = 0
            v2 -= 42
            var v22 int8 = v2
            _ = v21
            _ = v22
        } else {
            v1 -= v1
            v23 := GenS3{F0: v3, F1: fi20}
            _ = len(v13)
            _ = v23
        }
        fmt.Println(v2)
        _ = fi20
    }
    var v24 int = 32768
    fmt.Println(v24, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v24
    _ = v3
}

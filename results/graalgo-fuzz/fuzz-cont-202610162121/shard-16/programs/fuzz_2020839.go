package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 int }
func genFn0() uint {
    var v4 uint = 255
    if (((!false) || true) || false) {
        v5 := GenS1{F0: 2147483647}
        fmt.Println(v4)
        switch "foo bar" {
        case "hello":
            v6 := GenS2{F0: 1000}
            _ = v6
        case "foo bar":
            if ((942 > 1) && true) {
                var v7 float32 = 1e10
                for fi8 := 0; fi8 < 3; fi8++ {
                    v9 := GenS3{F0: (fi8 & fi8), F1: (^(fi8))}
                    _ = fi8
                    _ = v9
                }
                v10 := []int{4294967295, 1}
                _ = v10
                _ = v7
            }
        case "":
            v11 := []int{127, 128, -735, 735}
            _ = v11
        }
        _ = v5
    }
    v12 := GenS1{F0: -840}
    _ = v12
    _ = v4
    return v4
}

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 0
    v3 := []uint32{2, 42, 2, 2147483647}
    v13 := GenS2{F0: 403}
    fmt.Println(v2)
    var v14 string = ("hello" + ("" + "foo bar"))
    for fi15 := 0; fi15 < 3; fi15++ {
        fmt.Println(v14, v2, v14)
        for fi16 := 0; fi16 < 20; fi16++ {
            fmt.Println(v2, fi16)
            if ((true && (65535 >= 516)) && (true || (1 > 42))) {
                var v17 int = fi16
                var v18 uint = 0
                _ = v17
                _ = v18
            }
            v14 = "hello"
            _ = fi16
        }
        _ = fi15
    }
    var v19 uint8 = 24
    v14 = (("hello" + ("foo bar" + v14)) + v14)
    v20 := GenS1{F0: 42}
    var v21 bool = true
    v22 := GenS3{F0: -2147483648, F1: -364}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}

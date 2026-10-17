package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 uint16; F2 int16 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 int32 = 2147483647
    v2 := []int{-222, 127}
    v3 := make(map[int]int)
    fmt.Println(v1)
    switch "" {
    case "foo bar":
        var v4 string = "foo bar"
        _ = v4
    case "a":
        v1 = v1
    default:
        v1 = v1
    }
    v5 := GenS1{F0: 0}
    v6 := GenS2{F0: true, F1: 942, F2: -32768}
    for fi7 := 0; fi7 < 1; fi7++ {
        switch (fi7 & (^(fi7))) {
        case 4294967295:
            if true {
                v1 -= v1
                for fi8 := 0; fi8 < 11; fi8++ {
                    var v9 uint8 = 19
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v10 := v3[(fi7 ^ -95)]
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            } else {
                if ((-2.25 >= 1.5) && true) {
                    var v11 uint = 4294967295
                    var v12 int8 = 0
                    var v13 uint32 = 42
                    _ = v11
                    _ = v12
                    _ = v13
                } else {
                    var v14 float32 = (1e10 + -2.25)
                    _ = v14
                }
                v15 := GenS1{F0: 2147483648}
                _ = v15
            }
        default:
            var v16 string = ("" + "hello")
            _ = v16
        }
        _ = fi7
    }
    var v17 string = "a"
    v1 += v1
    if false {
        v3[1] = 0
        if false {
            var v18 bool = true
            _ = v18
        } else {
            v19 := GenS3{F0: 249}
            var v20 uint8 = 128
            v17 = v17
            _ = v19
            _ = v20
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v21 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 int16; F2 uint8 }
func genFn0(p0 uint, p1 uint32, p2 uint32) int8 {
    if (1 > 1) {
        v4 := GenS1{F0: 496}
        if true {
            if false {
                v5 := []int{128, 128, 0}
                for fi6 := 0; fi6 < 16; fi6++ {
                    if true {
                        if true {
                            if false {
                                p1 = p2
                                switch (-(fi6)) {
                                case -225:
                                    fmt.Println(p1, p1, p1)
                                }
                            }
                            switch (fi6 & (fi6 << (((p0 + p0)) & 63))) {
                            case -534:
                                var v7 int16 = -1
                                _ = v7
                            case -235:
                                v8 := GenS2{F0: 184}
                                _ = v8
                            case 4294967295:
                                var v9 uint16 = 1
                                _ = v9
                            }
                            v10 := GenS2{F0: 128}
                            _ = v10
                        }
                        p0 -= (^(p0))
                    }
                    v11 := GenS2{F0: 98}
                    _ = fi6
                    _ = v11
                }
                _ = v5
            }
        }
        _ = v4
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := GenS3{F0: 4294967295, F1: 1, F2: 42}
        _ = fi12
        _ = v13
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = p2
    return 32
}

func main() {
    v1 := make(map[int]int)
    var v2 int = 42
    v3 := []int{128, v2}
    var v14 string = "foo bar"
    var v15 float64 = float64(0.0)
    fmt.Println(v2, v15)
    var v16 uint32 = 181
    if (false && ((!(4294967295 <= 18446744073709551615)) && true)) {
        switch v14 {
        case "":
            v14 = "\n\t"
        case "hello":
            fmt.Println(v16, v2)
        default:
            fmt.Println(v14, v16)
        }
        v15 = 0.0
        v17 := v1[81]
        _ = v17
    }
    if true {
        var v18 uint16 = 255
        _ = v18
    } else {
        v16 = (v16 - 455)
        var v19 int32 = 1
        v20 := GenS2{F0: 128}
        _ = v19
        _ = v20
    }
    v14 = "foo bar"
    var v21 string = "a"
    var v22 int32 = -351
    var v23 uint = 605
    var v24 bool = (3.14159 <= (((v15 * float64(0.1)) - 0.1) / v15))
    var v25 int16 = 0
    var v26 int64 = 372
    v24 = (1 < uint16(v23))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}

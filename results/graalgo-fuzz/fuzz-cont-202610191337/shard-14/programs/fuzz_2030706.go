package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 int32 }
func genFn0(p0 int, p1 int32, p2 int) uint64 {
    v7 := GenS1{F0: 101, F1: 112, F2: 4294967295}
    var v8 int32 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 629
    var v3 uint32 = 4294967295
    v4 := make(map[string]int)
    v5 := make(map[string]bool)
    var v6 int64 = 111
    fmt.Println(v6, v2, v2)
    v9 := []int{}
    v1["a"] = (v3 >> (((^(v2))) & 63))
    var v10 uint8 = 127
    var v11 bool = true
    v11 = false
    var v12 bool = v11
    var v13 bool = v12
    var v14 uint16 = 42
    for fi15 := 0; fi15 < 9; fi15++ {
        v16 := GenS2{F0: v12, F1: 94}
        fmt.Println(v12, v14, fi15)
        v17 := GenS3{F0: 1, F1: "a"}
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 uint64 = 4294967295
    for fi19 := 0; fi19 < 20; fi19++ {
        for fi20 := 0; fi20 < 10; fi20++ {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            switch fi19 {
            case 32767:
                v11 = v13
            case -694:
                fmt.Println(v18)
            case 32768:
                if (v11 || (!(32767 <= -32768))) {
                    v11 = (1.5 <= ((0.0 / -2.25) * float64(0.0)))
                    v6 = (v6 + v6)
                    if (v11 || v13) {
                        if v13 {
                            var v21 uint8 = 134
                            if true {
                                v22 := GenS3{F0: (^(v18)), F1: ("foo bar" + "a")}
                                _ = v22
                            } else {
                                var v23 int32 = 42
                                var v24 uint = (v2 >> ((v2) & 63))
                                if v12 {
                                    v21 = (v10 ^ uint8(v18))
                                }
                                _ = v23
                                _ = v24
                            }
                            _ = v21
                        }
                        v25 := GenS2{F0: (!("" < "hello")), F1: v2}
                        v26 := GenS4{F0: 42}
                        _ = v25
                        _ = v26
                    }
                } else {
                    var v27 float32 = 0.1
                    v11 = (42 > -32768)
                    _ = v27
                }
            default:
                var v28 int = 32767
                _ = v28
            }
            _ = fi20
        }
        _ = fi19
    }
    v29 := GenS2{F0: v11, F1: uint(v10)}
    for fi30 := 0; fi30 < 17; fi30++ {
        var v31 uint16 = 255
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 10; fi32++ {
        v6 = v6
        _ = fi32
    }
    fmt.Println(v12, v2)
    if true {
        fmt.Println(v6, v18, v6)
        v11 = ((false && false) || ((!v13) && v13))
    } else {
        v3 = (v3 + uint32(v14))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}

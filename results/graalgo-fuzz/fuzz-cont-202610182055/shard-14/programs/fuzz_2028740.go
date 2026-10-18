package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int32; F1 uint16; F2 int8 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 uint8 }
func genFn0() uint8 {
    v6 := []int{4294967295, 2147483647, 1}
    var v7 uint = 0
    v8 := GenS1{F0: 4294967295, F1: 681}
    _ = v6
    _ = v7
    _ = v8
    return 53
}
func genFn1() int64 {
    if true {
        v9 := GenS1{F0: 439, F1: 980}
        _ = v9
    } else {
        var v10 bool = false
        _ = v10
    }
    v11 := []int{4294967295, 128, 949, 127}
    var v12 uint16 = 2
    if len(v11) > 0 {
        v11[0] = 32767
    }
    _ = v11
    _ = v12
    return -9223372036854775808
}

func main() {
    var v1 bool = (42 <= 643)
    var v2 uint32 = 1
    v3 := []int{-1, 32768}
    v4 := make(map[string]string)
    var v5 string = ("" + "a")
    v13 := GenS1{F0: 448, F1: 42}
    v5 = ((v5 + "a") + v5)
    if ((v5 >= "foo bar") || (-1 == 127)) {
        v3 = append(v3, -270)
    }
    fmt.Println(v1)
    switch 761 {
    case 128:
        var v14 int = 4294967295
        _ = v14
    case 612:
        var v15 uint64 = 0
        _ = v15
    case 87:
        if ((3.14159 > -2.25) && v1) {
            var v16 uint8 = 1
            v17 := GenS2{F0: 42, F1: 917, F2: 127}
            v3 = append(v3, 2147483647)
            _ = v16
            _ = v17
        } else {
            v2 = 2147483647
            v2 -= (^(v2))
            v5 = (v5 + v5)
        }
    default:
        switch v5 {
        case "foo bar":
            for fi18 := 0; fi18 < 7; fi18++ {
                if true {
                    v5 = (v5 + v5)
                    fmt.Println(v1, v5)
                    for fi19 := 0; fi19 < 9; fi19++ {
                        var v20 bool = false
                        if v20 {
                            fmt.Println(v1, fi18)
                            if len(v3) > 0 {
                                v3[0] = (fi18 & fi19)
                            }
                            var v21 uint64 = 456
                            _ = v21
                        }
                        if ((-2.25 > ((float32(0.0) - 1e10) - 1e10)) && (4294967295 != 42)) {
                            v22 := GenS3{F0: 136, F1: true}
                            _ = v22
                        }
                        _ = fi19
                        _ = v20
                    }
                }
                _ = fi18
            }
        case "a":
            v23 := GenS1{F0: uint(v2), F1: 938}
            _ = v23
        }
    }
    v24 := GenS3{F0: -751, F1: v1}
    if (v1 || v1) {
        v25 := GenS4{F0: 1}
        var v26 uint64 = 747
        _ = v25
        _ = v26
    } else {
        v27 := v4[""]
        _ = v27
    }
    fmt.Println(v1, v2)
    if (true || (!true)) {
        v1 = v1
    }
    v4[(v5 + v5)] = v5
    fmt.Println(v5, v5, v1)
    for fi28 := 0; fi28 < 17; fi28++ {
        v29 := GenS2{F0: 2147483647, F1: 2, F2: 127}
        _ = fi28
        _ = v29
    }
    var v30 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 int8; F1 uint; F2 int8 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 uint8; F1 int16 }
func genFn0(p0 uint, p1 uint32, p2 int32) string {
    for fi4 := 0; fi4 < 5; fi4++ {
        p0 = (p0 + 42)
        _ = fi4
    }
    if (2 <= 411) {
        p2 += (^(p2))
        var v5 int16 = 344
        for fi6 := 0; fi6 < 10; fi6++ {
            if (false || (((128 >= int(p2)) || ((602 != (p1 >> (((p0 * (^(p0)))) & 63))) && false)) && true)) {
                p2 = (p2 | -253)
                if (true || false) {
                    var v7 string = "hello"
                    var v8 float32 = 0.1
                    _ = v7
                    _ = v8
                } else {
                    var v9 uint8 = 255
                    switch ("" + "hello") {
                    case "hello":
                        v10 := []int{-842, (fi6 - fi6), (fi6 | 175)}
                        _ = v10
                    default:
                        fmt.Println(v9)
                    }
                    _ = v9
                }
                p1 = (^(p1))
            }
            _ = fi6
        }
        _ = v5
    }
    var v11 int = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return "hello"
}
func genFn1(p0 int32, p1 int8, p2 int) string {
    for fi12 := 0; fi12 < 16; fi12++ {
        fmt.Println(p0)
        _ = fi12
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        v14 := GenS1{F0: 2}
        v15 := GenS1{F0: 1}
        fmt.Println(p0, fi13, p2)
        _ = fi13
        _ = v14
        _ = v15
    }
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    return ""
}

func main() {
    var v1 int16 = 1000
    var v2 int8 = 42
    v3 := make(map[string]uint32)
    switch ("foo bar" + "a") {
    case "\n\t":
        v3["hello"] = 1
    }
    if true {
        for fi16 := 0; fi16 < 14; fi16++ {
            for fi17 := 0; fi17 < 6; fi17++ {
                if (!((float32(-2.25) / 0.1) >= ((float32((-2.25 + ((0.0 / -2.25) / 1.0))) - 0.0) - 1.5))) {
                    var v18 uint64 = 983
                    _ = v18
                } else {
                    var v19 uint64 = 2
                    fmt.Println(fi17, v2)
                    _ = v19
                }
                v20 := GenS1{F0: 802}
                fmt.Println(v2)
                _ = fi17
                _ = v20
            }
            v3["hello"] = 2
            switch "\n\t" {
            case "foo bar":
                v1 += v1
            case "":
                fmt.Println(v1, fi16, fi16)
            default:
                _ = len(v3)
            }
            _ = fi16
        }
    } else {
        v2 += v2
        var v21 uint = 4294967295
        if ("a" != "hello") {
            v22 := GenS1{F0: 2147483647}
            _ = v22
        }
        _ = v21
    }
    v23 := GenS2{F0: 876, F1: 42}
    v24 := []int{263}
    var v25 uint32 = 2147483648
    v1 -= (v1 / ((0) | 1))
    v24 = append(v24, 42)
    switch 2147483647 {
    case 127:
        v26 := GenS3{F0: 37, F1: 618, F2: v2}
        _ = v26
    default:
        fmt.Println(v2, v1)
    }
    v2 -= -68
    var v27 uint = 2
    v28 := GenS4{F0: 0, F1: 2}
    var v29 float64 = 1e10
    delete(v3, ("foo bar" + "a"))
    for fi30 := 0; fi30 < 4; fi30++ {
        _ = len(v3)
        var v31 uint32 = (v25 & 2147483647)
        _ = fi30
        _ = v31
    }
    switch -9223372036854775808 {
    case -575:
        switch "" {
        case "foo bar":
            v32 := GenS5{F0: 127, F1: v1}
            _ = v32
        case "a":
            for fi33 := 0; fi33 < 10; fi33++ {
                v2 += (v2 >> (((v27 >> (((v27 | v27)) & 63))) & 63))
                if len(v24) >= 2 {
                    v24 = v24[1:len(v24)]
                }
                v25 = v25
                _ = fi33
            }
        default:
            v34 := GenS2{F0: 127, F1: 1}
            _ = v34
        }
    }
    var v35 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v35
}

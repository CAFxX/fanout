package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int; F1 int8; F2 uint8 }
type GenS6 struct { F0 uint16; F1 uint8 }
func genFn0() bool {
    v5 := GenS1{F0: 0, F1: 9, F2: 216}
    var v6 uint8 = 1
    v7 := []int{}
    _ = v5
    _ = v6
    _ = v7
    return (true && false)
}

func main() {
    var v1 string = "a"
    var v2 float64 = 0.0
    v3 := []int{-351, 332, 128}
    v4 := []uint32{}
    v3 = append(v3, 128)
    v2 += (1e10 * (3.14159 / 0.1))
    var v8 uint8 = 255
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 int64 = 0
        v11 := GenS1{F0: (v10 * v10), F1: (v8 % (((^(v8))) | 1)), F2: v10}
        _ = fi9
        _ = v10
        _ = v11
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if true {
        _ = len(v3)
        v4 = append(v4, 706)
        v1 = v1
    } else {
        if false {
            switch 60 {
            case -384:
                var v12 bool = false
                _ = v12
            default:
                fmt.Println(v2)
            }
            v13 := GenS1{F0: -36, F1: 127, F2: 0}
            _ = v13
        } else {
            v2 -= float64(3.14159)
            if (0 < 0) {
                v2 = float64(1e10)
                var v14 bool = true
                v15 := GenS1{F0: 1, F1: (v8 * v8), F2: -1}
                _ = v14
                _ = v15
            }
            var v16 uint16 = 32767
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        v18 := GenS2{F0: 1, F1: -1}
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v20 := GenS3{F0: 473}
        _ = fi19
        _ = v20
    }
    switch v1 {
    case "a":
        if (1.5 < (v2 / (v2 - 1.5))) {
            var v21 float32 = -2.25
            fmt.Println(v1, v21)
            _ = v21
        } else {
            if (!false) {
                for fi22 := 0; fi22 < 18; fi22++ {
                    var v23 int = (fi22 / (((-(fi22))) | 1))
                    _ = fi22
                    _ = v23
                }
                v24 := GenS4{F0: true}
                _ = v24
            }
            if ("a" != (v1 + ((("foo bar" + "a") + "a") + "\n\t"))) {
                var v25 int64 = -1
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                fmt.Println(v25, v8, v8)
                _ = v25
            }
        }
    case "hello":
        v2 += v2
    case "\n\t":
        switch -9223372036854775808 {
        case -982:
            for fi26 := 0; fi26 < 4; fi26++ {
                v27 := GenS3{F0: 2147483647}
                _ = fi26
                _ = v27
            }
        case 128:
            fmt.Println(v8, v2)
        case -742:
            var v28 uint8 = v8
            _ = v28
        }
    default:
        v8 += v8
    }
    var v29 uint = 411
    v30 := GenS5{F0: 32768, F1: -1, F2: (^(v8))}
    if true {
        v31 := GenS2{F0: 999, F1: 1}
        v32 := GenS6{F0: 181, F1: (v8 / ((v8) | 1))}
        fmt.Println(v2)
        _ = v31
        _ = v32
    } else {
        v33 := GenS4{F0: ((false || (!false)) && (-1 >= 0))}
        var v34 string = v1
        _ = v33
        _ = v34
    }
    v35 := GenS2{F0: 255, F1: 0}
    var v36 float64 = -2.25
    if true {
        v37 := GenS5{F0: 4294967295, F1: 1, F2: (v8 & (v8 & v8))}
        if len(v3) > 0 {
            v3[0] = int(v29)
        }
        switch v1 {
        case "\n\t":
            fmt.Println(v8, v29, v29)
        case "a":
            v8 -= (v8 << ((557) & 63))
        case "foo bar":
            v3 = append(v3, -9223372036854775808)
        }
        _ = v37
    } else {
        var v38 uint16 = 0
        _ = v38
    }
    var v39 uint = v29
    v36 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v8
}

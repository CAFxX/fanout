package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int16 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 string; F1 int64; F2 int8 }
func genFn0() int64 {
    if false {
        v5 := GenS1{F0: -1, F1: 1, F2: -257}
        for fi6 := 0; fi6 < 1; fi6++ {
            var v7 int = -1
            v8 := GenS2{F0: -350, F1: 113}
            fmt.Println(fi6, v7)
            _ = fi6
            _ = v7
            _ = v8
        }
        var v9 uint8 = 2
        _ = v5
        _ = v9
    } else {
        var v10 float64 = (-2.25 + 3.14159)
        var v11 uint8 = 0
        _ = v10
        _ = v11
    }
    if false {
        switch 128 {
        case 438:
            switch 1 {
            case -2147483648:
                var v12 uint16 = 947
                _ = v12
            case 32768:
                var v13 int8 = -128
                _ = v13
            default:
                if (!(true && true)) {
                    v14 := []int{32768, 32767, 0}
                    _ = v14
                } else {
                    var v15 uint = 1
                    v15 = (v15 % ((v15) | 1))
                    _ = v15
                }
            }
        case 128:
            v16 := GenS1{F0: 1, F1: 42, F2: -32768}
            _ = v16
        }
        v17 := []int{127}
        _ = v17
    }
    return -1
}

func main() {
    var v1 int64 = -540
    v2 := []uint32{}
    v3 := []int{153}
    v4 := []string{"hello"}
    switch "a" {
    case "a":
        v1 = (^(v1))
    default:
        _ = len(v3)
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        fmt.Println(fi18, v1)
        if false {
            for fi19 := 0; fi19 < 8; fi19++ {
                v20 := GenS1{F0: -32768, F1: 955, F2: -32768}
                fmt.Println(fi18, fi18, fi18)
                _ = fi19
                _ = v20
            }
        } else {
            var v21 int32 = -507
            if true {
                _ = len(v2)
                var v22 int = 0
                _ = v22
            }
            _ = v21
        }
        var v23 int64 = v1
        _ = fi18
        _ = v23
    }
    v1 -= 957
    v24 := GenS3{F0: 2147483648, F1: 1}
    v25 := GenS4{F0: "", F1: 9223372036854775807, F2: -1}
    v1 -= (v1 & v1)
    var v26 uint64 = 0
    if true {
        v27 := GenS1{F0: 32767, F1: 32767, F2: -472}
        for fi28 := 0; fi28 < 2; fi28++ {
            v26 -= (v26 | v26)
            _ = fi28
        }
        fmt.Println(v26)
        _ = v27
    }
    switch ("foo bar" + "") {
    case "\n\t":
        fmt.Println(v26, v26)
    default:
        switch ("foo bar" + ("\n\t" + "foo bar")) {
        case "\n\t":
            var v29 int64 = genFn0()
            _ = v29
        case "":
            v26 -= (v26 & 963)
        case "a":
            var v30 uint = 65535
            _ = v30
        default:
            if (!(((v1 == (^(v1))) || true) || (1 > 244))) {
                v26 = v26
                v26 = v26
                v1 += v1
            } else {
                v1 += v1
                var v31 int16 = 0
                _ = v31
            }
        }
    }
    v26 += (v26 + v26)
    for fi32 := 0; fi32 < 19; fi32++ {
        var v33 int64 = (^(v1))
        v34 := GenS3{F0: 495, F1: 42}
        var v35 string = "a"
        _ = fi32
        _ = v33
        _ = v34
        _ = v35
    }
    var v36 int = 62
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v36
    _ = v4
}

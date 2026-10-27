package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint }
type GenS3 struct { F0 bool; F1 int16; F2 uint }
type GenS4 struct { F0 int16; F1 string }

func main() {
    var v1 int8 = -1
    var v2 uint64 = 0
    var v3 int8 = (v1 | v1)
    v4 := make(map[uint32]string)
    if (true && true) {
        if ("" <= ("" + ("" + (("" + "\n\t") + "a")))) {
            var v5 uint8 = 1
            if true {
                fmt.Println(v3)
            } else {
                v6 := GenS1{F0: 1}
                _ = v6
            }
            _ = v5
        }
        if false {
            v2 -= (v2 | (v2 + (^(v2))))
        }
        v1 = 0
    } else {
        v1 = -128
    }
    fmt.Println(v1, v2, v1)
    fmt.Println(v2, v1)
    if false {
        for fi7 := 0; fi7 < 13; fi7++ {
            v3 = (v1 / ((1) | 1))
            for fi8 := 0; fi8 < 14; fi8++ {
                var v9 int32 = 42
                v2 -= (v2 % ((18446744073709551615) | 1))
                if (false || ((v9 & v9) == v9)) {
                    v10 := []int{(fi8 >> ((42) & 63))}
                    var v11 int64 = -9223372036854775808
                    _ = v10
                    _ = v11
                }
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
        var v12 uint64 = v2
        v13 := GenS1{F0: -1}
        _ = v12
        _ = v13
    }
    var v14 bool = (true || (!(1.5 < ((1.5 * 1e10) / -2.25))))
    var v15 uint32 = 1
    if ((v3 & (v1 / ((-74) | 1))) == 100) {
        v2 -= (v2 >> ((42) & 63))
    } else {
        var v16 uint64 = v2
        fmt.Println(v1, v14)
        _ = v16
    }
    v3 += (v3 & 42)
    var v17 bool = v14
    v2 = v2
    for fi18 := 0; fi18 < 9; fi18++ {
        if true {
            v3 += v1
            var v19 uint32 = (v15 + v15)
            v20 := GenS2{F0: 212, F1: 173, F2: 604}
            _ = v19
            _ = v20
        } else {
            if (float32(1e10) >= 3.14159) {
                v14 = v17
                switch ((((("foo bar" + ("a" + ("\n\t" + ("hello" + "hello")))) + "foo bar") + "") + "hello") + "a") {
                case "a":
                    fmt.Println(v14, v15)
                case "foo bar":
                    fmt.Println(v1, v15)
                }
            }
            fmt.Println(v14)
            fmt.Println(v15, v17, v15)
        }
        var v21 float32 = 3.14159
        switch fi18 {
        case 128:
            v1 -= 42
        case 0:
            v14 = (!v14)
        case 1:
            v22 := GenS3{F0: v14, F1: 0, F2: 0}
            _ = v22
        default:
            for fi23 := 0; fi23 < 13; fi23++ {
                _ = len(v4)
                _ = fi23
            }
        }
        _ = fi18
        _ = v21
    }
    v24 := []int{}
    v1 -= v1
    v25 := GenS4{F0: 0, F1: "foo bar"}
    var v26 int64 = 42
    var v27 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}

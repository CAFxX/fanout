package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 int64; F1 uint8; F2 uint8 }
func genFn0(p0 uint, p1 uint64) uint8 {
    if (228 < 2) {
        for fi5 := 0; fi5 < 5; fi5++ {
            var v6 uint64 = 0
            v7 := GenS1{F0: 2147483647, F1: "hello"}
            fmt.Println(v6, p1)
            _ = fi5
            _ = v6
            _ = v7
        }
        for fi8 := 0; fi8 < 20; fi8++ {
            fmt.Println(fi8)
            var v9 float32 = 0.1
            _ = fi8
            _ = v9
        }
    } else {
        var v10 int64 = -519
        if (1.5 == -2.25) {
            var v11 bool = false
            _ = v11
        } else {
            v10 = v10
            var v12 int8 = 42
            fmt.Println(p1, v12)
            _ = v12
        }
        _ = v10
    }
    var v13 int16 = 1000
    _ = p0
    _ = p1
    _ = v13
    return uint8(p0)
}
func genFn1(p0 int64, p1 int16) int8 {
    if (128 > 229) {
        v14 := []int{32768}
        _ = v14
    } else {
        if false {
            v15 := []int{127, 128, -464}
            p1 = p1
            for fi16 := 0; fi16 < 13; fi16++ {
                p1 = p1
                _ = fi16
            }
            _ = v15
        } else {
            var v17 int16 = (p1 * (p1 ^ (p1 & p1)))
            var v18 float64 = (3.14159 / 3.14159)
            p1 += (^(p1))
            _ = v17
            _ = v18
        }
        v19 := []int{0, -2147483648, 32767}
        _ = v19
    }
    if true {
        if (47 <= 1) {
            var v20 int = -2147483648
            var v21 uint = 1
            v22 := GenS2{F0: (v21 ^ v21), F1: 128}
            _ = v20
            _ = v21
            _ = v22
        }
    }
    v23 := []int{32768, 32768}
    if len(v23) >= 2 {
        v23 = v23[1:len(v23)]
    }
    var v24 uint = 567
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    return 0
}

func main() {
    var v1 float32 = float32((1e10 * 1e10))
    var v2 int8 = 100
    v3 := make(map[uint32]bool)
    var v4 int = 42
    v1 += 0.0
    if false {
        for fi25 := 0; fi25 < 14; fi25++ {
            delete(v3, 0)
            v1 = 0.0
            fmt.Println(v2, v1)
            _ = fi25
        }
    }
    v4 = (v4 + (v4 << ((42) & 63)))
    v26 := []int{(v4 ^ 428), (v4 & (v4 + -9223372036854775808)), v4}
    fmt.Println(v2, v2)
    v27 := GenS1{F0: -343, F1: "foo bar"}
    var v28 uint = 65535
    var v29 int64 = -9223372036854775808
    delete(v3, 1)
    for fi30 := 0; fi30 < 4; fi30++ {
        v28 = (v28 % ((2) | 1))
        _ = fi30
    }
    v28 += v28
    for fi31 := 0; fi31 < 3; fi31++ {
        switch "hello" {
        case "\n\t":
            for fi32 := 0; fi32 < 20; fi32++ {
                if true {
                    v33 := GenS3{F0: (v2 * (v2 >> ((v28) & 63))), F1: 1}
                    var v34 uint64 = 1
                    var v35 int32 = 440
                    _ = v33
                    _ = v34
                    _ = v35
                }
                switch int(v29) {
                case -2147483648:
                    v1 -= v1
                default:
                    v36 := GenS4{F0: 9223372036854775807, F1: 1, F2: genFn0(v28, 4294967295)}
                    _ = v36
                }
                _ = fi32
            }
        case "hello":
            delete(v3, 0)
        }
        _ = fi31
    }
    fmt.Println(v4, v28)
    fmt.Println(v2, v29, v4)
    var v37 int16 = -1
    if (genFn1(v29, 1) != (v2 + 1)) {
        var v38 uint16 = 0
        _ = v38
    }
    var v39 uint64 = 181
    for fi40 := 0; fi40 < 8; fi40++ {
        switch (v4 ^ (v4 - int(v37))) {
        case -630:
            switch ("a" + "\n\t") {
            case "":
                var v41 bool = true
                _ = v41
            default:
                v29 = v29
            }
        default:
            v28 = v28
        }
        _ = fi40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v37
    _ = v39
    _ = v4
}

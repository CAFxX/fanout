package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint32 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 bool; F1 uint8; F2 bool }
type GenS4 struct { F0 uint8; F1 int32; F2 int }
func genFn0(p0 int) int64 {
    v7 := []int{p0, (^(p0)), -1}
    for fi8 := 0; fi8 < 20; fi8++ {
        fmt.Println(p0, fi8)
        _ = fi8
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 int = (fi9 - 32767)
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v7
    return 945
}

func main() {
    var v1 uint = 1
    v2 := []int{-2147483648}
    var v3 float64 = (1e10 / 1e10)
    v4 := make(map[string]uint32)
    var v5 int32 = -1
    var v6 float32 = (-2.25 * (-2.25 - 0.0))
    v6 += v6
    switch ("hello" + "") {
    case "\n\t":
        var v11 uint8 = 128
        _ = v11
    case "foo bar":
        var v12 uint8 = 42
        _ = v12
    case "hello":
        v6 = float32(v3)
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        v14 := GenS1{F0: v5, F1: "hello", F2: 2}
        if (!false) {
            v3 = (v3 + float64((0.1 - 0.1)))
            if false {
                switch -159 {
                case 42:
                    v5 += (v5 | (v5 & v5))
                case 32767:
                    v3 -= 1.5
                }
                for fi15 := 0; fi15 < 5; fi15++ {
                    _ = len(v4)
                    _ = fi15
                }
                v6 = v6
            }
            v6 = 0.0
        }
        _ = fi13
        _ = v14
    }
    v5 -= 514
    v1 += (v1 << ((v1) & 63))
    fmt.Println(v3)
    v16 := v4[(("" + "\n\t") + "foo bar")]
    for fi17 := 0; fi17 < 12; fi17++ {
        v6 = ((v6 - (0.1 / float32(-2.25))) + (v6 * 0.1))
        for fi18 := 0; fi18 < 5; fi18++ {
            var v19 float32 = 1.5
            if true {
                if (false && true) {
                    delete(v4, "a")
                }
                v20 := GenS1{F0: (v5 + v5), F1: "a", F2: (v16 % ((851) | 1))}
                _ = v20
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    var v21 uint64 = 1
    for fi22 := 0; fi22 < 15; fi22++ {
        v16 = v16
        v16 = 2147483648
        _ = fi22
    }
    var v23 float64 = v3
    var v24 int64 = 970
    if true {
        switch "" {
        case "hello":
            var v25 uint16 = 1
            _ = v25
        case "a":
            v24 = v24
        case "\n\t":
            for fi26 := 0; fi26 < 13; fi26++ {
                if false {
                    v6 -= (v6 / v6)
                    fmt.Println(v6, v24)
                } else {
                    var v27 int8 = 0
                    _ = v27
                }
                for fi28 := 0; fi28 < 20; fi28++ {
                    v2 = append(v2, 1)
                    var v29 int32 = (v5 % ((v5) | 1))
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi28
                    _ = v29
                }
                _ = fi26
            }
        default:
            _ = len(v2)
        }
        if false {
            v24 = 289
            for fi30 := 0; fi30 < 20; fi30++ {
                v31 := GenS2{F0: 42, F1: 42}
                switch (^(fi30)) {
                case 456:
                    v32 := GenS3{F0: false, F1: 167, F2: false}
                    _ = v32
                default:
                    fmt.Println(v21, v5)
                }
                _ = fi30
                _ = v31
            }
        }
        fmt.Println(v24)
    } else {
        v3 = (1e10 - (v23 + (v23 * float64((v6 + v6)))))
    }
    var v33 int8 = 97
    v34 := GenS4{F0: 1, F1: -1, F2: 32767}
    if true {
        v4["a"] = v16
        _ = len(v2)
        fmt.Println(v21, v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}

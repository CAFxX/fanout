package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 string }
func genFn0(p0 int64, p1 uint32, p2 uint8) uint16 {
    for fi7 := 0; fi7 < 14; fi7++ {
        fmt.Println(p0)
        _ = fi7
    }
    switch 4294967295 {
    case 833:
        for fi8 := 0; fi8 < 20; fi8++ {
            if (float64(0.0) == -2.25) {
                v9 := GenS1{F0: p2}
                _ = v9
            } else {
                switch ("hello" + ("a" + ("" + "\n\t"))) {
                case "hello":
                    var v10 int = (fi8 << ((51) & 63))
                    _ = v10
                }
            }
            v11 := []int{-1, (fi8 >> ((1) & 63))}
            _ = fi8
            _ = v11
        }
    case 32767:
        p0 += 1
    }
    _ = p0
    _ = p1
    _ = p2
    return 422
}
func genFn1(p0 uint8, p1 uint32) int32 {
    p1 += (p1 / ((666) | 1))
    switch ("a" + "") {
    case "":
        if (false || false) {
            for fi12 := 0; fi12 < 9; fi12++ {
                v13 := GenS1{F0: p0}
                _ = fi12
                _ = v13
            }
            switch 247 {
            case 4294967295:
                v14 := []int{1}
                _ = v14
            }
            var v15 float64 = 0.1
            _ = v15
        }
    case "a":
        v16 := GenS1{F0: uint8(p1)}
        _ = v16
    default:
        if (((false || (42 != -625)) || true) || true) {
            p1 += p1
            var v17 uint8 = p0
            if ((false && ((!true) && true)) && true) {
                var v18 int64 = -1
                for fi19 := 0; fi19 < 16; fi19++ {
                    if (!true) {
                        v20 := GenS2{F0: (p1 / ((p1) | 1))}
                        p0 = v17
                        _ = v20
                    }
                    v21 := GenS2{F0: (p1 ^ 795)}
                    _ = fi19
                    _ = v21
                }
                var v22 int16 = -624
                _ = v18
                _ = v22
            }
            _ = v17
        } else {
            var v23 int64 = 916
            var v24 uint16 = 1000
            if (-2147483648 > -1) {
                if (32767 <= 1) {
                    v24 = genFn0(42, p1, (p0 + 128))
                }
                for fi25 := 0; fi25 < 15; fi25++ {
                    fmt.Println(p0)
                    var v26 int32 = 42
                    var v27 bool = (1e10 <= -2.25)
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            } else {
                switch ("\n\t" + "foo bar") {
                case "":
                    p1 = 2147483648
                }
                v28 := []int{838, -9223372036854775808, 0}
                _ = v28
            }
            _ = v23
            _ = v24
        }
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[int]bool)
    v3 := make(map[uint32]uint32)
    var v4 int32 = 42
    var v5 int64 = 0
    v6 := []uint32{0, 15, 2147483647}
    _ = len(v6)
    for fi29 := 0; fi29 < 19; fi29++ {
        fmt.Println(v4, fi29, v5)
        var v30 int64 = (v5 ^ -259)
        _ = fi29
        _ = v30
    }
    var v31 int16 = -221
    var v32 float32 = 1.5
    fmt.Println(v32)
    var v33 int8 = -128
    delete(v3, 120)
    v34 := GenS3{F0: 127, F1: ""}
    if (!(!false)) {
        v5 = v5
    }
    v4 -= v4
    var v35 int16 = (v31 + (v31 * 808))
    delete(v3, 1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}

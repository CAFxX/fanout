package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 uint64; F2 int }
type GenS5 struct { F0 uint16; F1 bool; F2 int8 }
func genFn0(p0 int32, p1 uint8, p2 uint16) uint32 {
    switch "\n\t" {
    case "":
        v5 := GenS1{F0: -9223372036854775808}
        _ = v5
    case "\n\t":
        var v6 float64 = float64(0.0)
        _ = v6
    default:
        switch -9223372036854775808 {
        case 127:
            for fi7 := 0; fi7 < 15; fi7++ {
                fmt.Println(p0)
                for fi8 := 0; fi8 < 10; fi8++ {
                    var v9 int64 = -747
                    _ = fi8
                    _ = v9
                }
                _ = fi7
            }
        case 0:
            p0 = -1
        case 4294967295:
            var v10 bool = false
            _ = v10
        }
    }
    v11 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 2147483648
}
func genFn1(p0 int64, p1 uint64, p2 uint) int64 {
    var v12 uint = 42
    if (true && ((float32((0.0 - 3.14159)) / -2.25) > ((1.5 / 0.1) / -2.25))) {
        p0 = p0
        if true {
            p0 = -1
            v13 := []int{-262, -9223372036854775808, int(p2)}
            v14 := GenS2{F0: 0, F1: -1, F2: 127}
            _ = v13
            _ = v14
        }
        for fi15 := 0; fi15 < 3; fi15++ {
            v16 := []int{383, -126}
            _ = fi15
            _ = v16
        }
    } else {
        fmt.Println(v12, p1, p0)
        for fi17 := 0; fi17 < 6; fi17++ {
            var v18 float64 = (0.0 + (0.1 * 0.0))
            _ = fi17
            _ = v18
        }
    }
    p0 = (p0 | int64(v12))
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return -1
}

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[uint32]bool)
    v3 := []uint32{(v1 << ((825) & 63)), (v1 & (v1 % ((v1) | 1)))}
    v4 := []string{((("foo bar" + ("a" + "foo bar")) + "a") + ""), ("foo bar" + ("" + "hello")), "hello", ("" + (((("a" + ("\n\t" + ("" + (("hello" + "") + "foo bar")))) + "a") + "a") + ("hello" + "foo bar")))}
    v19 := GenS3{F0: 0}
    v1 = 4294967295
    fmt.Println(v1)
    v1 += v1
    var v20 uint8 = 255
    v1 = 42
    if false {
        _ = len(v3)
        switch "" {
        case "\n\t":
            v21 := GenS4{F0: -463, F1: uint64(v1), F2: 9223372036854775807}
            _ = v21
        }
    } else {
        for fi22 := 0; fi22 < 4; fi22++ {
            v1 = (v1 * v1)
            if (v1 >= v1) {
                for fi23 := 0; fi23 < 2; fi23++ {
                    v24 := GenS4{F0: -32768, F1: 18446744073709551615, F2: (fi22 & (fi22 % ((fi23) | 1)))}
                    v25 := GenS5{F0: 42, F1: true, F2: -65}
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v26 int64 = 42
                _ = v26
            } else {
                _ = len(v4)
                for fi27 := 0; fi27 < 20; fi27++ {
                    var v28 int = (-(fi22))
                    _ = fi27
                    _ = v28
                }
            }
            _ = fi22
        }
        if (false && false) {
            var v29 int64 = -1
            var v30 uint = 2
            _ = v29
            _ = v30
        } else {
            v4 = append(v4, ("" + "a"))
        }
    }
    var v31 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v4
}

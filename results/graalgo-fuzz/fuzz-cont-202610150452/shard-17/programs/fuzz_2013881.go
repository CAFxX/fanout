package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 int64; F2 uint64 }
func genFn0() uint {
    v7 := GenS1{F0: true}
    var v8 float32 = 1e10
    v9 := GenS2{F0: -608}
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}
func genFn1(p0 int, p1 uint32) string {
    v10 := []int{}
    for fi11 := 0; fi11 < 2; fi11++ {
        fmt.Println(fi11, p1)
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v10
    return ("foo bar" + ((((("hello" + "foo bar") + "a") + (("\n\t" + "a") + "a")) + "\n\t") + "hello"))
}

func main() {
    v1 := []uint32{455, 65535}
    var v2 uint16 = 2
    v3 := []uint32{985, 1, 4294967295}
    v4 := []string{"\n\t", "", (("" + "") + ("hello" + "\n\t"))}
    v5 := make(map[uint32]uint32)
    var v6 int64 = 42
    if false {
        v12 := GenS2{F0: 1}
        _ = v12
    } else {
        v13 := GenS2{F0: 1}
        v5[2147483648] = 0
        v14 := v5[4294967295]
        _ = v13
        _ = v14
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v6 = (v6 ^ (v6 << ((1) & 63)))
    switch "" {
    case "\n\t":
        var v15 int32 = 2147483647
        _ = v15
    case "":
        v6 = (v6 + v6)
    }
    switch 32767 {
    case -9223372036854775808:
        v6 = v6
    case 2147483647:
        v2 = v2
    case -724:
        if (((true && (!(true && false))) && (42 < 2147483647)) || false) {
            var v16 bool = (((float32(0.1) + float32(0.0)) + (-2.25 + float32(float64(1e10)))) != float32((0.1 / 1.5)))
            v16 = (!false)
            _ = v16
        }
    default:
        v6 += v6
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        v2 += v2
        if false {
            v4 = append(v4, "hello")
        }
        _ = fi17
    }
    if ((!(2 < 42)) || (true || false)) {
        v6 -= (v6 ^ v6)
        v18 := GenS3{F0: 0, F1: -478, F2: 1}
        v19 := GenS3{F0: -1, F1: v6, F2: 493}
        _ = v18
        _ = v19
    }
    var v20 int = -530
    for fi21 := 0; fi21 < 12; fi21++ {
        v20 = fi21
        v2 = 32767
        var v22 uint = 65535
        _ = fi21
        _ = v22
    }
    if (("hello" >= (("a" + "hello") + ("hello" + ("\n\t" + (("\n\t" + "") + "a"))))) || (4294967295 >= 18446744073709551615)) {
        var v23 int8 = 1
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v23
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi24 := 0; fi24 < 19; fi24++ {
        if ((float64(3.14159) / 3.14159) < 0.0) {
            _ = len(v5)
            var v25 int64 = (v6 + -9223372036854775808)
            _ = v25
        }
        v20 += -632
        _ = fi24
    }
    v6 = (v6 - 969)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

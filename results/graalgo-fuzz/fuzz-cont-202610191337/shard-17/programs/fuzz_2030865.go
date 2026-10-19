package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8 }
func genFn0() int32 {
    var v7 int8 = 42
    var v8 float32 = 3.14159
    _ = v7
    _ = v8
    return -2147483648
}
func genFn1(p0 int, p1 uint16, p2 uint16) int {
    if (!true) {
        p2 = 0
        p1 = p1
    } else {
        if (!false) {
            for fi9 := 0; fi9 < 7; fi9++ {
                v10 := []int{(p0 * -1), (^(fi9)), (p0 >> ((18446744073709551615) & 63))}
                _ = fi9
                _ = v10
            }
            fmt.Println(p2)
        } else {
            var v11 float32 = ((1e10 / float32(0.1)) + 1.5)
            _ = v11
        }
        var v12 uint = 341
        fmt.Println(p0, v12, p0)
        _ = v12
    }
    v13 := []int{}
    var v14 int64 = 0
    var v15 float32 = -2.25
    var v16 bool = ((^(p2)) == p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    return p0
}

func main() {
    v1 := []int{924, 128, 128, -1}
    var v2 uint64 = 2
    var v3 bool = true
    var v4 int16 = -149
    v5 := []string{"\n\t", "a"}
    v6 := []string{"foo bar", "", "hello"}
    var v17 uint8 = 153
    v3 = (genFn0() > -199)
    v2 = 10
    if len(v5) > 0 {
        v5[0] = "foo bar"
    }
    v17 -= v17
    if len(v1) > 0 {
        v1[0] = 128
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 float32 = (3.14159 - ((-2.25 / float32(0.1)) * 0.1))
        var v20 uint64 = v2
        _ = fi18
        _ = v19
        _ = v20
    }
    v17 = (v17 - v17)
    for fi21 := 0; fi21 < 18; fi21++ {
        for fi22 := 0; fi22 < 5; fi22++ {
            v4 += (v4 / ((v4) | 1))
            _ = fi22
        }
        fmt.Println(v17)
        v2 = v2
        _ = fi21
    }
    if (false && (!(!v3))) {
        var v23 int8 = 81
        v24 := GenS1{F0: (v4 * 32767), F1: genFn0(), F2: (("a" + (("foo bar" + "foo bar") + "")) + "hello")}
        _ = v23
        _ = v24
    }
    fmt.Println(v4)
    if false {
        v25 := GenS1{F0: 42, F1: 304, F2: ""}
        var v26 uint8 = 127
        switch genFn1(32767, 2, 363) {
        case 175:
            if (1 > 2) {
                for fi27 := 0; fi27 < 19; fi27++ {
                    fmt.Println(v26, v2, fi27)
                    var v28 int8 = 100
                    if v3 {
                        v26 -= (v17 ^ v26)
                    }
                    _ = fi27
                    _ = v28
                }
                var v29 float32 = 3.14159
                var v30 uint = 42
                _ = v29
                _ = v30
            } else {
                var v31 uint = 255
                v4 += int16(v17)
                _ = v31
            }
        case 863:
            var v32 float32 = 3.14159
            _ = v32
        case 128:
            switch 240 {
            case 541:
                var v33 int8 = 42
                _ = v33
            case -9223372036854775808:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case 4294967295:
                v34 := GenS2{F0: v2}
                _ = v34
            default:
                fmt.Println(v4)
            }
        }
        _ = v25
        _ = v26
    }
    var v35 int16 = (v4 ^ 1000)
    if (42 < 0) {
        for fi36 := 0; fi36 < 8; fi36++ {
            v35 += -32768
            v37 := GenS1{F0: v4, F1: 2147483647, F2: ("hello" + ("hello" + ("foo bar" + "foo bar")))}
            v4 = v4
            _ = fi36
            _ = v37
        }
    }
    v38 := GenS3{F0: -80}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}

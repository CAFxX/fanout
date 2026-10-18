package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int64; F1 int64; F2 uint }
type GenS3 struct { F0 uint16; F1 int64 }
func genFn0(p0 int16, p1 int64, p2 bool) int8 {
    p1 += 1
    if (0.0 <= (3.14159 - ((1.5 + 3.14159) * ((((1e10 * 1e10) - (1e10 * 1.5)) - 1e10) / -2.25)))) {
        p0 = (p0 % ((p0) | 1))
    } else {
        var v6 uint32 = 2
        switch ((("\n\t" + "\n\t") + "\n\t") + "hello") {
        case "\n\t":
            var v7 int32 = -2147483648
            _ = v7
        case "a":
            var v8 float64 = ((3.14159 + -2.25) + (float64(3.14159) / -2.25))
            _ = v8
        }
        _ = v6
    }
    if ((!p2) || (18446744073709551615 > 4294967295)) {
        p1 += (p1 - (p1 & (p1 >> ((2) & 63))))
        if p2 {
            p2 = (p2 && true)
        } else {
            v9 := GenS1{F0: 222, F1: 2147483647}
            switch ("hello" + "") {
            case "hello":
                for fi10 := 0; fi10 < 13; fi10++ {
                    v11 := GenS1{F0: 65535, F1: 208}
                    v12 := []int{fi10, -1, (fi10 * (fi10 - 1)), 927}
                    fmt.Println(p2)
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            case "foo bar":
                p2 = p2
            default:
                if (p2 && (p2 || (323 != 0))) {
                    p0 = int16(p1)
                }
            }
            _ = v9
        }
        p0 -= -645
    } else {
        var v13 int16 = 820
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}
func genFn1(p0 bool) int16 {
    v14 := []int{}
    var v15 uint8 = 127
    v15 += (^(v15))
    _ = p0
    _ = v14
    _ = v15
    return -472
}

func main() {
    var v1 float32 = -2.25
    v2 := []string{("\n\t" + "\n\t"), ("hello" + ""), "a", ("foo bar" + ("" + ("foo bar" + (("a" + "\n\t") + "\n\t"))))}
    v3 := []string{("\n\t" + "foo bar"), "foo bar", "hello", "a"}
    v4 := []int{128, 0}
    v5 := []uint32{2147483647}
    var v16 int = 1
    var v17 string = ("foo bar" + "hello")
    var v18 int8 = genFn0(32767, -9223372036854775808, false)
    if (!(!(!false))) {
        v16 = (v16 ^ v16)
    } else {
        v19 := GenS2{F0: 9223372036854775807, F1: -691, F2: 1}
        if false {
            fmt.Println(v16, v1, v17)
            switch (v17 + (v17 + v17)) {
            case "a":
                v20 := GenS3{F0: 366, F1: 628}
                _ = v20
            case "foo bar":
                var v21 float32 = 3.14159
                _ = v21
            case "\n\t":
                var v22 int32 = -1
                _ = v22
            default:
                v1 += v1
            }
        } else {
            var v23 string = "\n\t"
            v23 = v17
            if false {
                if len(v4) > 0 {
                    v4[0] = (v16 * 9223372036854775807)
                }
                fmt.Println(v16)
            } else {
                fmt.Println(v18, v18, v1)
                v1 += 0.0
                var v24 uint32 = 739
                _ = v24
            }
            _ = v23
        }
        for fi25 := 0; fi25 < 19; fi25++ {
            var v26 int64 = 42
            _ = fi25
            _ = v26
        }
        _ = v19
    }
    var v27 uint = 2
    v27 = 671
    var v28 uint64 = 235
    v18 = genFn0(32767, 825, true)
    if true {
        if (!(true || true)) {
            _ = len(v2)
            if true {
                v18 += v18
                if true {
                    for fi29 := 0; fi29 < 14; fi29++ {
                        var v30 float32 = 0.1
                        _ = fi29
                        _ = v30
                    }
                    v27 = v27
                } else {
                    var v31 uint32 = 4294967295
                    _ = v31
                }
            }
        }
        var v32 int8 = v18
        _ = len(v5)
        _ = v32
    }
    fmt.Println(v17, v27, v27)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}

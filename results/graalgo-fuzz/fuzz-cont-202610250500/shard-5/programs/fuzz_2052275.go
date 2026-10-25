package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int8; F1 uint8 }
func genFn0(p0 string, p1 uint16) uint8 {
    if (!(42 != 128)) {
        p1 -= p1
        v6 := GenS1{F0: 205, F1: 42}
        if (true && true) {
            v7 := GenS2{F0: -1}
            _ = v7
        } else {
            p0 = ("" + "\n\t")
            var v8 int8 = -2
            p1 += p1
            _ = v8
        }
        _ = v6
    }
    v9 := GenS2{F0: 953}
    if false {
        for fi10 := 0; fi10 < 1; fi10++ {
            v11 := []int{(fi10 * 32768), int(p1), 127}
            p1 += p1
            p0 = (p0 + p0)
            _ = fi10
            _ = v11
        }
        p0 = p0
        v12 := GenS3{F0: 32767, F1: 207}
        _ = v12
    }
    var v13 bool = ((1 > 0) || true)
    v14 := GenS3{F0: 32767, F1: 192}
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v9
    return 128
}
func genFn1() uint64 {
    var v15 int = 1
    v16 := GenS3{F0: 0, F1: 65535}
    fmt.Println(v15)
    var v17 string = "hello"
    if true {
        if true {
            for fi18 := 0; fi18 < 10; fi18++ {
                var v19 float32 = -2.25
                v20 := []int{}
                _ = fi18
                _ = v19
                _ = v20
            }
            var v21 float64 = float64(1e10)
            _ = v21
        } else {
            var v22 bool = (32767 == 836)
            _ = v22
        }
        v23 := GenS3{F0: 447, F1: 0}
        _ = v23
    } else {
        var v24 int64 = -9223372036854775808
        _ = v24
    }
    _ = v15
    _ = v16
    _ = v17
    return 547
}

func main() {
    var v1 int64 = -412
    var v2 uint8 = 2
    var v3 string = "hello"
    v4 := []string{}
    var v5 int32 = 42
    fmt.Println(v2)
    v1 += v1
    switch "\n\t" {
    case "\n\t":
        v2 = genFn0(v3, 0)
    }
    switch (("foo bar" + (((v3 + v3) + v3) + v3)) + v3) {
    case "\n\t":
        var v25 int64 = 918
        _ = v25
    case "hello":
        var v26 uint16 = 889
        _ = v26
    default:
        var v27 uint = 440
        _ = v27
    }
    for fi28 := 0; fi28 < 1; fi28++ {
        v3 = (v3 + v3)
        var v29 int16 = 136
        _ = fi28
        _ = v29
    }
    var v30 float64 = 3.14159
    if ((((v3 + (v3 + "a")) + v3) + "foo bar") >= (v3 + v3)) {
        v3 = ("a" + ("foo bar" + (v3 + v3)))
        if (862 >= -2147483648) {
            v4 = append(v4, "\n\t")
            if (v30 >= float64(0.1)) {
                v31 := GenS4{F0: (-(v5)), F1: 127}
                _ = v31
            } else {
                var v32 int8 = 1
                var v33 uint32 = 696
                if (((((v30 / (v30 + 0.0)) - float64(((0.1 * 0.0) + 0.0))) * 3.14159) + v30) > v30) {
                    if ((v30 - v30) >= 1.5) {
                        v34 := GenS5{F0: -795, F1: v1}
                        v35 := GenS6{F0: (v2 >> ((743) & 63))}
                        v30 += v30
                        _ = v34
                        _ = v35
                    } else {
                        fmt.Println(v33)
                        v32 = (v32 | (-(v32)))
                        fmt.Println(v2, v5, v33)
                    }
                    v2 = (v2 << ((65535) & 63))
                }
                _ = v32
                _ = v33
            }
            v36 := GenS6{F0: (v2 * 143)}
            _ = v36
        } else {
            _ = len(v4)
        }
    }
    v37 := GenS5{F0: 581, F1: 1}
    if ((v1 % ((269) | 1)) != (v1 + v1)) {
        v3 = "hello"
        var v38 int8 = 94
        _ = v38
    }
    v30 = -2.25
    v5 -= (^(v5))
    for fi39 := 0; fi39 < 9; fi39++ {
        var v40 int = fi39
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
        v41 := GenS7{F0: -128, F1: 2}
        _ = fi39
        _ = v40
        _ = v41
    }
    v42 := GenS4{F0: v5, F1: (v2 & 128)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v37
    _ = v4
    _ = v42
    _ = v5
}

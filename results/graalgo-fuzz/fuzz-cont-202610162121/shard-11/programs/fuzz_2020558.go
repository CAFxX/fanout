package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int16; F1 int64; F2 uint }
func genFn0(p0 string, p1 uint32) uint32 {
    v7 := []int{127, 128, -1}
    var v8 int8 = 0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p1
}
func genFn1(p0 int, p1 uint32) uint64 {
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 int8 = 0
        p1 += p1
        _ = fi9
        _ = v10
    }
    switch ("foo bar" + ("a" + "foo bar")) {
    case "\n\t":
        if (2 == 42) {
            fmt.Println(p0, p1)
        }
    case "a":
        var v11 int = -9223372036854775808
        _ = v11
    }
    if true {
        if true {
            for fi12 := 0; fi12 < 5; fi12++ {
                var v13 int32 = 0
                for fi14 := 0; fi14 < 12; fi14++ {
                    switch "hello" {
                    case "a":
                        fmt.Println(fi12, fi12)
                    default:
                        p1 += 65535
                    }
                    v15 := GenS1{F0: 2, F1: 525}
                    p1 = (p1 * genFn0(("" + "hello"), genFn0("", p1)))
                    _ = fi14
                    _ = v15
                }
                _ = fi12
                _ = v13
            }
            if (0.0 > 1.5) {
                v16 := []int{p0, -943, p0, -2147483648}
                p0 = (p0 >> ((255) & 63))
                if len(v16) > 0 {
                    v16[0] = (p0 << ((788) & 63))
                }
                _ = v16
            }
            for fi17 := 0; fi17 < 6; fi17++ {
                p1 -= p1
                for fi18 := 0; fi18 < 13; fi18++ {
                    p1 = p1
                    _ = fi18
                }
                _ = fi17
            }
        } else {
            var v19 int16 = 1
            for fi20 := 0; fi20 < 14; fi20++ {
                v21 := GenS1{F0: 4294967295, F1: 65535}
                if (32767 >= (-(p0))) {
                    for fi22 := 0; fi22 < 13; fi22++ {
                        var v23 float32 = 3.14159
                        _ = fi22
                        _ = v23
                    }
                    var v24 float32 = ((1.5 - ((-2.25 + (1.5 - (3.14159 * -2.25))) - (1e10 - 1e10))) + 0.0)
                    _ = v24
                }
                v25 := GenS1{F0: 518, F1: 365}
                _ = fi20
                _ = v21
                _ = v25
            }
            _ = v19
        }
    } else {
        fmt.Println(p0, p0)
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 uint32 = 915
    var v2 float64 = -2.25
    v3 := make(map[int]string)
    var v4 uint32 = v1
    v5 := make(map[string]string)
    v6 := []string{"\n\t", ("a" + "hello")}
    fmt.Println(v1, v4)
    var v26 uint64 = 4294967295
    v26 = (v26 - v26)
    _ = len(v3)
    if false {
        if (v2 == float64(float32((v2 / v2)))) {
            v27 := v3[-605]
            for fi28 := 0; fi28 < 11; fi28++ {
                var v29 int16 = -1
                _ = fi28
                _ = v29
            }
            v27 = "hello"
            _ = v27
        }
    }
    var v30 int8 = 42
    for fi31 := 0; fi31 < 7; fi31++ {
        if false {
            v32 := GenS2{F0: 0, F1: 492, F2: 255}
            if (2 >= (v26 ^ v26)) {
                var v33 float32 = -2.25
                var v34 int32 = -1
                _ = v33
                _ = v34
            }
            _ = v32
        } else {
            for fi35 := 0; fi35 < 19; fi35++ {
                switch (fi31 << ((894) & 63)) {
                case -857:
                    fmt.Println(fi31, v30)
                case -9223372036854775808:
                    var v36 uint32 = 2147483647
                    _ = v36
                }
                _ = fi35
            }
            v37 := GenS1{F0: 1, F1: 255}
            var v38 bool = (true || false)
            _ = v37
            _ = v38
        }
        _ = fi31
    }
    var v39 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}

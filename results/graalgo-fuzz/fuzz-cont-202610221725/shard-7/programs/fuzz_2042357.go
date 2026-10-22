package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int32; F2 uint }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 string }
func genFn0(p0 int32, p1 uint64) uint64 {
    p0 += -2147483648
    var v6 float32 = (0.0 - float32(0.1))
    v7 := []int{2147483647, 32768, 42}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (^(p1))
}
func genFn1(p0 int32, p1 uint16, p2 uint64) bool {
    switch 42 {
    case 1:
        var v8 bool = true
        _ = v8
    default:
        p0 = 0
    }
    v9 := GenS1{F0: -1}
    if (-241 >= p0) {
        p2 -= (^(p2))
        p0 -= 1
    }
    if false {
        p2 -= (p2 / ((42) | 1))
    } else {
        v10 := GenS1{F0: -1}
        if ((-980 > 328) || false) {
            var v11 int32 = p0
            for fi12 := 0; fi12 < 1; fi12++ {
                for fi13 := 0; fi13 < 19; fi13++ {
                    v14 := []int{}
                    _ = fi13
                    _ = v14
                }
                v15 := GenS1{F0: 520}
                v11 += p0
                _ = fi12
                _ = v15
            }
            _ = v11
        }
        if (true || (false && true)) {
            var v16 int8 = -1
            switch "hello" {
            case "\n\t":
                v17 := GenS2{F0: 255, F1: 42, F2: 334}
                _ = v17
            }
            switch 805 {
            case -672:
                if (!false) {
                    for fi18 := 0; fi18 < 18; fi18++ {
                        if ((("" + "foo bar") + ("foo bar" + "hello")) >= (("" + "a") + "")) {
                            p2 += p2
                            fmt.Println(p1, p2)
                        }
                        _ = fi18
                    }
                    for fi19 := 0; fi19 < 4; fi19++ {
                        var v20 uint32 = 4294967295
                        _ = fi19
                        _ = v20
                    }
                }
            case 1:
                v21 := GenS2{F0: 226, F1: (p0 + p0), F2: 572}
                _ = v21
            default:
                for fi22 := 0; fi22 < 15; fi22++ {
                    v23 := GenS2{F0: 255, F1: p0, F2: 18446744073709551615}
                    _ = fi22
                    _ = v23
                }
            }
            _ = v16
        }
        _ = v10
    }
    var v24 float64 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v9
    return (18446744073709551615 == 0)
}

func main() {
    var v1 uint64 = 42
    v2 := []uint32{42, 2147483648, 2147483647, 2147483648}
    var v3 int8 = 100
    v4 := make(map[string]int)
    v5 := make(map[string]bool)
    switch "foo bar" {
    case "foo bar":
        v1 = 1
    }
    if (!false) {
        v4["hello"] = 442
        var v25 int32 = -185
        if (genFn0(v25, (v1 >> ((0) & 63))) >= v1) {
            if (v25 < -1) {
                switch "hello" {
                case "\n\t":
                    v26 := GenS1{F0: -325}
                    _ = v26
                case "":
                    var v27 uint = 255
                    _ = v27
                }
                v28 := GenS1{F0: 42}
                _ = v28
            }
            v29 := GenS3{F0: 1000, F1: 783}
            if len(v2) > 0 {
                v2[0] = 2
            }
            _ = v29
        }
        _ = v25
    } else {
        fmt.Println(v3)
        var v30 uint16 = 0
        _ = v30
    }
    v31 := GenS1{F0: 1}
    if false {
        var v32 string = "foo bar"
        _ = v32
    }
    for fi33 := 0; fi33 < 10; fi33++ {
        if (32767 < 1) {
            for fi34 := 0; fi34 < 8; fi34++ {
                var v35 int = (fi33 ^ 2147483647)
                _ = fi34
                _ = v35
            }
        } else {
            if ((((fi33 + fi33) != (fi33 >> ((0) & 63))) && true) && (true && true)) {
                _ = len(v4)
            } else {
                v3 += (v3 * v3)
            }
        }
        _ = fi33
    }
    fmt.Println(v1, v1)
    v36 := GenS1{F0: 9223372036854775807}
    v3 = v3
    var v37 uint32 = 0
    fmt.Println(v37, v3)
    if false {
        fmt.Println(v3, v1, v37)
        fmt.Println(v1, v3, v3)
        for fi38 := 0; fi38 < 2; fi38++ {
            v37 -= (v37 ^ 246)
            _ = len(v2)
            _ = fi38
        }
    } else {
        for fi39 := 0; fi39 < 11; fi39++ {
            v40 := GenS4{F0: "hello"}
            _ = fi39
            _ = v40
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}

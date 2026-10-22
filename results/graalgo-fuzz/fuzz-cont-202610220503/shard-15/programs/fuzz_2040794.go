package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32; F1 uint; F2 bool }
type GenS6 struct { F0 uint16; F1 uint64 }
func genFn0(p0 int8, p1 int) int {
    fmt.Println(p1, p0)
    p1 = p1
    for fi6 := 0; fi6 < 7; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := GenS1{F0: "\n\t"}
            var v9 uint16 = 42
            switch "" {
            case "\n\t":
                var v10 uint = 255
                _ = v10
            }
            v9 -= v9
            v11 := GenS1{F0: "a"}
            _ = fi7
            _ = v11
            _ = v8
            _ = v9
        }
        p0 = (p0 | p0)
        v12 := []int{(fi6 ^ 4294967295), (p1 * (p1 | p1))}
        _ = fi6
        _ = v12
    }
    var v13 int16 = 32767
    _ = p0
    _ = p1
    _ = v13
    return (-(p1))
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 2147483647
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]bool)
    var v5 int64 = 0
    v14 := GenS2{F0: ""}
    v15 := GenS3{F0: 18446744073709551615}
    v2 -= (^(v2))
    if false {
        v4[0] = false
        switch genFn0(-1, genFn0(42, genFn0(-1, genFn0(42, -789)))) {
        case -2147483648:
            fmt.Println(v5)
        }
        v16 := []int{}
        _ = v16
    } else {
        v5 -= v5
        v17 := GenS4{F0: 18446744073709551615}
        _ = v17
    }
    if ((((0.1 / 0.1) - (-2.25 + 0.1)) - 3.14159) >= -2.25) {
        for fi18 := 0; fi18 < 16; fi18++ {
            fmt.Println(v2, fi18)
            var v19 bool = (false || true)
            _ = fi18
            _ = v19
        }
        for fi20 := 0; fi20 < 7; fi20++ {
            if false {
                switch (fi20 & 1) {
                case -332:
                    if false {
                        v5 += (-(v5))
                        delete(v3, v2)
                    }
                default:
                    fmt.Println(v2, fi20)
                }
            } else {
                switch (fi20 ^ (fi20 - genFn0(1, -2147483648))) {
                case 930:
                    v5 = (v5 * (v5 | v5))
                case -82:
                    fmt.Println(fi20, v5, v5)
                case 128:
                    v21 := GenS5{F0: (v2 * (v2 & v2)), F1: 620, F2: true}
                    _ = v21
                default:
                    var v22 int32 = 1
                    _ = v22
                }
            }
            v5 -= (v5 ^ 1)
            v23 := []int{9223372036854775807, 1, (fi20 << ((42) & 63))}
            _ = fi20
            _ = v23
        }
        v24 := []int{127}
        _ = v24
    } else {
        v5 = v5
    }
    v2 += (v2 * 441)
    switch 127 {
    case 4294967295:
        for fi25 := 0; fi25 < 19; fi25++ {
            if true {
                var v26 int8 = 127
                _ = v26
            }
            _ = fi25
        }
    case 0:
        switch 32767 {
        case 32768:
            if (65535 == v2) {
                v2 -= 2147483648
            }
        default:
            var v27 float32 = -2.25
            _ = v27
        }
    }
    var v28 int16 = 1
    v28 += v28
    v29 := GenS6{F0: 0, F1: 4294967295}
    v5 = v5
    v30 := []int{42}
    v5 = (v5 * 42)
    v31 := v1[(v2 >> ((1) & 63))]
    v28 -= v28
    switch "hello" {
    case "foo bar":
        var v32 int8 = -1
        _ = v32
    case "a":
        delete(v4, 0)
    case "\n\t":
        fmt.Println(v28)
    default:
        for fi33 := 0; fi33 < 20; fi33++ {
            for fi34 := 0; fi34 < 11; fi34++ {
                var v35 bool = (v31 || false)
                _ = fi34
                _ = v35
            }
            _ = fi33
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 string; F1 int64; F2 uint32 }
type GenS5 struct { F0 uint8; F1 bool }
func genFn0(p0 uint8) uint32 {
    v6 := GenS1{F0: "a", F1: 2}
    if false {
        v7 := GenS1{F0: (("foo bar" + ("hello" + "a")) + ("hello" + (("a" + ("" + "\n\t")) + "foo bar"))), F1: 4294967295}
        p0 -= (p0 | p0)
        v8 := GenS1{F0: "foo bar", F1: 538}
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v6
    return 2147483647
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]bool)
    var v3 bool = ((1e10 / 1.0) >= float64((((0.1 / -2.25) - (0.1 + float32(1.5))) + (1e10 * float32((1e10 - 0.0))))))
    var v4 int64 = 42
    v5 := make(map[int]string)
    v3 = ((42 <= 18446744073709551615) && (!v3))
    if (1 > 911) {
        if (!v3) {
            v3 = (18446744073709551615 >= 18446744073709551615)
            fmt.Println(v3, v4)
        }
        switch "hello" {
        case "a":
            var v9 uint8 = 147
            _ = v9
        }
    }
    fmt.Println(v3)
    var v10 uint32 = 42
    if (false || true) {
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 uint32 = genFn0(215)
            _ = fi11
            _ = v12
        }
    } else {
        for fi13 := 0; fi13 < 19; fi13++ {
            fmt.Println(fi13, fi13, v3)
            v14 := []int{(fi13 % ((-9223372036854775808) | 1)), fi13}
            _ = fi13
            _ = v14
        }
        var v15 float32 = -2.25
        _ = v15
    }
    v16 := GenS2{F0: genFn0(128), F1: (v4 / (((v4 % (((v4 >> ((18446744073709551615) & 63))) | 1))) | 1)), F2: 42}
    var v17 string = "\n\t"
    if (v3 && (32767 == 255)) {
        switch "a" {
        case "hello":
            var v18 float32 = -2.25
            _ = v18
        default:
            _ = len(v5)
        }
        if true {
            if (1 >= 32767) {
                v19 := GenS3{F0: v4, F1: ("" <= v17)}
                var v20 uint16 = 1
                _ = v19
                _ = v20
            } else {
                var v21 float64 = float64(0.0)
                var v22 uint = 4294967295
                for fi23 := 0; fi23 < 18; fi23++ {
                    switch fi23 {
                    case -661:
                        v10 -= 350
                    case 32768:
                        v10 = v10
                    case -2147483648:
                        var v24 int16 = 32767
                        _ = v24
                    }
                    _ = fi23
                }
                _ = v21
                _ = v22
            }
            var v25 float64 = (1e10 * float64(0.0))
            _ = v25
        }
    } else {
        if v3 {
            if (4294967295 == 65535) {
                v4 -= (v4 << ((581) & 63))
                v26 := GenS4{F0: v17, F1: (v4 / ((-9223372036854775808) | 1)), F2: v10}
                v27 := GenS5{F0: 255, F1: true}
                _ = v26
                _ = v27
            }
            v28 := GenS2{F0: genFn0(0), F1: (^(v4)), F2: 42}
            _ = v28
        } else {
            var v29 int8 = 127
            var v30 int = 467
            v31 := GenS5{F0: 0, F1: (false || (-116 < (v29 % ((v29) | 1))))}
            _ = v29
            _ = v30
            _ = v31
        }
    }
    v32 := []int{}
    fmt.Println(v10)
    for fi33 := 0; fi33 < 6; fi33++ {
        fmt.Println(v10)
        _ = fi33
    }
    for fi34 := 0; fi34 < 12; fi34++ {
        _ = len(v5)
        var v35 int32 = -2147483648
        v3 = true
        _ = fi34
        _ = v35
    }
    fmt.Println(v17)
    for fi36 := 0; fi36 < 16; fi36++ {
        v37 := GenS3{F0: v4, F1: v3}
        v38 := GenS3{F0: (-(v4)), F1: ((v4 / (((v4 >> ((528) & 63))) | 1)) < v4)}
        v5[fi36] = v17
        _ = fi36
        _ = v37
        _ = v38
    }
    var v39 uint8 = 183
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v39
    _ = v4
    _ = v5
}

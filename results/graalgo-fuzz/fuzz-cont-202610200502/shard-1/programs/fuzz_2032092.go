package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint64; F1 uint8; F2 uint }
type GenS5 struct { F0 uint64; F1 int64; F2 uint }
func genFn0(p0 string) int {
    p0 = "hello"
    switch 2147483647 {
    case 42:
        v5 := []int{4294967295}
        _ = v5
    default:
        var v6 float32 = float32((1e10 * (0.0 / 1e10)))
        _ = v6
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := []int{fi7, 0}
        if true {
            var v9 int8 = 85
            for fi10 := 0; fi10 < 6; fi10++ {
                v11 := GenS1{F0: 42, F1: 743}
                _ = fi10
                _ = v11
            }
            fmt.Println(fi7, fi7, p0)
            _ = v9
        }
        for fi12 := 0; fi12 < 17; fi12++ {
            if (false || (540 >= -244)) {
                var v13 bool = true
                v13 = v13
                _ = v13
            } else {
                switch 128 {
                case 42:
                    var v14 uint8 = 42
                    _ = v14
                case 2147483647:
                    v15 := GenS2{F0: p0, F1: 2}
                    _ = v15
                case 127:
                    switch 0 {
                    case 128:
                        var v16 uint = 2
                        _ = v16
                    case 9223372036854775807:
                        p0 = "hello"
                    case 127:
                        var v17 float64 = (3.14159 * 0.1)
                        _ = v17
                    }
                default:
                    var v18 int64 = 0
                    _ = v18
                }
            }
            var v19 int8 = -1
            v20 := GenS3{F0: 0, F1: fi12}
            v21 := GenS4{F0: 0, F1: 180, F2: 42}
            _ = fi12
            _ = v19
            _ = v20
            _ = v21
        }
        _ = fi7
        _ = v8
    }
    v22 := []int{32767, 4294967295}
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    _ = p0
    _ = v22
    return -1
}
func genFn1(p0 uint8, p1 int8, p2 uint8) bool {
    if true {
        for fi23 := 0; fi23 < 9; fi23++ {
            v24 := []int{}
            for fi25 := 0; fi25 < 16; fi25++ {
                p2 += (p2 + (^(p2)))
                v24 = append(v24, genFn0("hello"))
                if (false || false) {
                    p1 = (-(p1))
                    v24 = append(v24, 4294967295)
                }
                _ = fi25
            }
            _ = len(v24)
            _ = fi23
            _ = v24
        }
        v26 := []int{127, -680, -8, 1}
        _ = v26
    }
    v27 := GenS5{F0: 810, F1: -79, F2: 26}
    if (248 <= -982) {
        for fi28 := 0; fi28 < 13; fi28++ {
            p0 = (p2 << ((1) & 63))
            p2 -= 239
            _ = fi28
        }
        v29 := []int{}
        _ = v29
    }
    switch ("\n\t" + (("" + "hello") + ("foo bar" + ("a" + "hello")))) {
    case "":
        p1 = 42
    }
    var v30 int64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    _ = v30
    return false
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[string]int)
    v3 := make(map[int]uint32)
    v4 := make(map[string]bool)
    if (0.1 != 3.14159) {
        for fi31 := 0; fi31 < 14; fi31++ {
            v32 := []int{}
            _ = fi31
            _ = v32
        }
        for fi33 := 0; fi33 < 15; fi33++ {
            var v34 int = 0
            _ = fi33
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 2; fi35++ {
        v1 -= v1
        if false {
            v1 += (v1 & 4294967295)
            _ = len(v4)
        }
        var v36 uint = 18446744073709551615
        _ = fi35
        _ = v36
    }
    for fi37 := 0; fi37 < 9; fi37++ {
        var v38 uint = 230
        v39 := []int{genFn0("a"), (fi37 + (fi37 / (((-(fi37))) | 1)))}
        _ = fi37
        _ = v38
        _ = v39
    }
    v1 += (v1 / ((v1) | 1))
    v1 = 295
    v1 = (v1 % ((v1) | 1))
    v1 = v1
    var v40 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
}

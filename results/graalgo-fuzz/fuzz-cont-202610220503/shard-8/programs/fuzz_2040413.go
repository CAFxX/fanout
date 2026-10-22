package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 int16; F2 int16 }
func genFn0(p0 bool, p1 int) uint16 {
    fmt.Println(p1)
    v6 := []int{(p1 >> ((1) & 63)), (p1 ^ 32768)}
    var v7 int8 = -40
    v8 := GenS1{F0: "hello"}
    fmt.Println(v7, v7)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := []uint32{42, 2147483648, 2}
    var v2 int = 0
    var v3 int64 = 42
    v4 := []uint32{2147483647, 0}
    v5 := make(map[int]int)
    for fi9 := 0; fi9 < 12; fi9++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v2 = v2
        _ = fi9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        for fi11 := 0; fi11 < 2; fi11++ {
            if true {
                v12 := v5[(fi10 - fi11)]
                _ = v12
            } else {
                v3 = v3
                if (3.14159 <= ((0.1 + 0.1) + 3.14159)) {
                    if (1.5 >= 1e10) {
                        v2 -= (fi11 << ((2) & 63))
                        fmt.Println(v2)
                    }
                    v13 := GenS1{F0: ""}
                    _ = v13
                } else {
                    var v14 uint64 = 2
                    v15 := GenS1{F0: "foo bar"}
                    v3 += (-(v3))
                    if (993 <= 2147483647) {
                        fmt.Println(v3, v2)
                        v2 = v2
                        v16 := GenS1{F0: "hello"}
                        _ = v16
                    } else {
                        v2 += fi10
                        if (((false && (65535 > 405)) || (true || false)) && (true && ((v2 * (fi10 / ((-325) | 1))) < (fi10 - (v2 ^ (v2 << ((255) & 63))))))) {
                            v17 := GenS1{F0: "a"}
                            _ = v17
                        } else {
                            var v18 int32 = -1
                            v3 = (v3 & 0)
                            v18 = v18
                            _ = v18
                        }
                    }
                    _ = v14
                    _ = v15
                }
                fmt.Println(v2)
            }
            if len(v1) > 0 {
                v1[0] = 1
            }
            var v19 int64 = (v3 & (v3 * v3))
            _ = fi11
            _ = v19
        }
        switch ("a" + "") {
        case "a":
            v3 += v3
        default:
            fmt.Println(v2, v3, v3)
        }
        _ = fi10
    }
    v20 := GenS1{F0: "a"}
    v21 := GenS1{F0: ("hello" + ("\n\t" + "\n\t"))}
    for fi22 := 0; fi22 < 13; fi22++ {
        fmt.Println(v2, fi22)
        _ = fi22
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v2 += -871
    v3 = v3
    v3 += -9223372036854775808
    v23 := GenS2{F0: v3}
    switch ("\n\t" + "\n\t") {
    case "a":
        v3 = 1
    case "\n\t":
        v2 = 127
    case "":
        v3 += v3
    }
    var v24 uint8 = 75
    v25 := GenS3{F0: 65535, F1: 0, F2: -480}
    for fi26 := 0; fi26 < 2; fi26++ {
        var v27 uint64 = uint64(v24)
        _ = fi26
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}

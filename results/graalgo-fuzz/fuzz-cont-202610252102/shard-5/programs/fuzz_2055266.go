package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int16 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint8, p1 uint64, p2 uint16) uint16 {
    var v7 int32 = -2147483648
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 int8 = int8(p2)
        var v10 uint8 = p0
        v11 := []int{(fi8 - 128), -665, (^(fi8)), -2147483648}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    if true {
        var v12 uint32 = 1
        if false {
            v12 += v12
        } else {
            var v13 float64 = 1e10
            p2 += (p2 % ((p2) | 1))
            _ = v13
        }
        switch -1 {
        case 9223372036854775807:
            v14 := []int{}
            _ = v14
        }
        _ = v12
    } else {
        switch "a" {
        case "foo bar":
            fmt.Println(p0, v7, p0)
        }
        if (((float32(((3.14159 + 0.0) + ((float64(1.5) * (0.0 / float64(0.1))) * -2.25))) - ((0.0 / (float32(1e10) / 1.5)) - float32(1.5))) + -2.25) >= (0.1 / 1.0)) {
            if (!((!false) && (false && false))) {
                var v15 int = -9223372036854775808
                _ = v15
            } else {
                fmt.Println(p2, p0)
                p2 = p2
            }
            v16 := GenS1{F0: false, F1: 255, F2: 1}
            if (true && (false && (true || ("hello" >= "foo bar")))) {
                if true {
                    var v17 uint64 = (p1 & (p1 >> ((18446744073709551615) & 63)))
                    v18 := []int{42, 127, 9223372036854775807, 1}
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(p1, p0, p0)
                    case "hello":
                        var v19 uint = 2
                        _ = v19
                    default:
                        if false {
                            var v20 int64 = 851
                            v18 = append(v18, 32768)
                            if len(v18) > 0 {
                                v18[0] = 1
                            }
                            _ = v20
                        } else {
                            p2 = p2
                        }
                    }
                    _ = v17
                    _ = v18
                }
                var v21 uint = 729
                _ = v21
            }
            _ = v16
        } else {
            v22 := []int{-9223372036854775808, -606, -1, 32768}
            v22 = append(v22, 0)
            _ = v22
        }
    }
    var v23 int8 = int8(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    _ = v7
    return p2
}
func genFn1() string {
    for fi24 := 0; fi24 < 16; fi24++ {
        var v25 float32 = -2.25
        v26 := GenS1{F0: true, F1: 949, F2: 42}
        for fi27 := 0; fi27 < 5; fi27++ {
            var v28 uint = 406
            _ = fi27
            _ = v28
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    var v29 int16 = 32767
    v29 -= -32768
    var v30 uint16 = 32767
    var v31 int8 = 127
    _ = v29
    _ = v30
    _ = v31
    return "\n\t"
}

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 1
    v3 := make(map[int]uint32)
    var v4 int32 = 1
    var v5 int = 335
    v6 := make(map[int]bool)
    v4 = -2147483648
    for fi32 := 0; fi32 < 3; fi32++ {
        v4 += (v4 + (^(v4)))
        _ = fi32
    }
    v5 -= (v5 & (^(v5)))
    v33 := GenS2{F0: (v4 - (v4 - v4)), F1: v5}
    v5 -= v5
    v5 = (v5 & -9223372036854775808)
    v5 = (v5 << ((927) & 63))
    var v34 uint16 = v2
    v4 = v4
    v35 := GenS3{F0: false, F1: 255}
    v36 := GenS2{F0: v4, F1: v5}
    switch "hello" {
    case "foo bar":
        v4 = 2147483647
    case "a":
        var v37 int8 = 100
        _ = v37
    default:
        if (13 >= 59) {
            for fi38 := 0; fi38 < 4; fi38++ {
                for fi39 := 0; fi39 < 18; fi39++ {
                    if ((true || true) && ("hello" > "foo bar")) {
                        v34 += genFn0(1, 0, 1)
                        var v40 uint8 = 128
                        _ = v40
                    }
                    var v41 int16 = -403
                    _ = fi39
                    _ = v41
                }
                v2 = 472
                var v42 int16 = -32768
                _ = fi38
                _ = v42
            }
            v34 = (v2 % (((v2 | (v2 ^ (v2 - (v34 >> ((601) & 63)))))) | 1))
        } else {
            var v43 int32 = (v4 >> ((881) & 63))
            _ = v43
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 bool }
type GenS2 struct { F0 string; F1 uint16; F2 string }
type GenS3 struct { F0 uint16; F1 int32; F2 int64 }
type GenS4 struct { F0 int; F1 uint8 }
func genFn0(p0 uint64, p1 uint8) int {
    v7 := GenS1{F0: -1, F1: 920, F2: true}
    var v8 uint32 = 1
    if (!(true || (42 != 0))) {
        var v9 bool = ((!false) && false)
        _ = v9
    } else {
        v10 := []int{}
        if true {
            var v11 int64 = 357
            _ = v11
        } else {
            for fi12 := 0; fi12 < 19; fi12++ {
                p0 += (^(p0))
                _ = fi12
            }
            v8 -= (v8 & (v8 - 65535))
            p1 = p1
        }
        _ = v10
    }
    p1 -= (p1 & (p1 + 255))
    fmt.Println(v8, v8, v8)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{}
    var v3 int32 = 1
    var v4 int = -1
    var v5 int64 = -9223372036854775808
    var v6 int8 = -11
    fmt.Println(v5, v3, v3)
    var v13 int64 = (v5 * (v5 ^ v5))
    for fi14 := 0; fi14 < 4; fi14++ {
        if len(v2) > 0 {
            v2[0] = 739
        }
        _ = fi14
    }
    var v15 uint64 = 18446744073709551615
    v16 := GenS1{F0: 1, F1: 2, F2: ((true && true) || (40 > 478))}
    if (((false && true) && false) && (735 <= 4294967295)) {
        fmt.Println(v3)
        v13 += v13
        for fi17 := 0; fi17 < 11; fi17++ {
            _ = len(v1)
            fmt.Println(v6, v15, v13)
            _ = fi17
        }
    } else {
        v5 -= v13
    }
    v4 = v4
    if false {
        var v18 int8 = v6
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v19 := GenS2{F0: "", F1: 958, F2: ""}
        _ = v18
        _ = v19
    } else {
        fmt.Println(v6, v13, v13)
        if ("hello" < (("foo bar" + "") + (((("foo bar" + ("" + ("\n\t" + "a"))) + "hello") + "hello") + "\n\t"))) {
            for fi20 := 0; fi20 < 5; fi20++ {
                v2 = append(v2, 860)
                _ = fi20
            }
            v13 = int64(v3)
        }
        if len(v2) > 0 {
            v2[0] = 120
        }
    }
    v2 = append(v2, 42)
    for fi21 := 0; fi21 < 12; fi21++ {
        var v22 int16 = int16(v6)
        v23 := GenS2{F0: ("foo bar" + "hello"), F1: 352, F2: ("\n\t" + "\n\t")}
        v24 := GenS3{F0: 65535, F1: v3, F2: (^(v13))}
        _ = fi21
        _ = v22
        _ = v23
        _ = v24
    }
    v4 = 144
    fmt.Println(v13)
    v13 = (v5 << ((265) & 63))
    if (v3 >= -2147483648) {
        v25 := GenS2{F0: ((("" + ("\n\t" + "\n\t")) + ("a" + "\n\t")) + "foo bar"), F1: 372, F2: "hello"}
        if (0 <= 128) {
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if true {
                    v4 -= genFn0((v15 * v15), 1)
                    for fi26 := 0; fi26 < 19; fi26++ {
                        var v27 uint16 = 42
                        _ = fi26
                        _ = v27
                    }
                    var v28 int32 = (^(v3))
                    _ = v28
                }
            }
            var v29 bool = false
            _ = v29
        }
        v30 := GenS1{F0: -32768, F1: 0, F2: false}
        _ = v25
        _ = v30
    }
    var v31 string = ("hello" + "hello")
    v32 := GenS4{F0: v4, F1: 127}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}

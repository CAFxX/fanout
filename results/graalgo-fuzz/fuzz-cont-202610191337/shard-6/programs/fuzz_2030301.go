package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int32, p1 uint64) int64 {
    v4 := GenS1{F0: -9223372036854775808, F1: ("\n\t" + ("a" + ""))}
    v5 := []int{}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 25
}
func genFn1() uint16 {
    var v6 float64 = 3.14159
    var v7 uint64 = 2
    var v8 float64 = (v6 / v6)
    fmt.Println(v7)
    if false {
        for fi9 := 0; fi9 < 17; fi9++ {
            v7 += (v7 ^ v7)
            var v10 bool = (2 < 42)
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 4; fi11++ {
            fmt.Println(v8)
            var v12 uint8 = 128
            for fi13 := 0; fi13 < 11; fi13++ {
                v6 = v6
                _ = fi13
            }
            _ = fi11
            _ = v12
        }
        v6 = 3.14159
    } else {
        var v14 uint64 = (v7 % ((v7) | 1))
        v15 := []int{32768, 0}
        _ = v14
        _ = v15
    }
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{("hello" + ("\n\t" + "hello")), "\n\t", "\n\t"}
    v3 := []uint32{856}
    if true {
        delete(v1, -9223372036854775808)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        var v16 uint64 = 18446744073709551615
        var v17 uint16 = 255
        if (!false) {
            fmt.Println(v16, v16)
        }
        _ = v16
        _ = v17
    }
    var v18 uint32 = 0
    v19 := GenS1{F0: 42, F1: "foo bar"}
    v3 = append(v3, 42)
    fmt.Println(v18)
    v2 = append(v2, "\n\t")
    v20 := GenS1{F0: 32768, F1: "\n\t"}
    v18 += (v18 * 2147483648)
    _ = len(v1)
    for fi21 := 0; fi21 < 19; fi21++ {
        if true {
            v22 := GenS2{F0: 18446744073709551615}
            _ = v22
        }
        _ = fi21
    }
    switch ("\n\t" + ("foo bar" + ((("hello" + "foo bar") + ("foo bar" + ("\n\t" + "hello"))) + "foo bar"))) {
    case "foo bar":
        v23 := GenS2{F0: 242}
        _ = v23
    case "":
        var v24 int32 = 310
        _ = v24
    case "hello":
        v18 = (v18 ^ v18)
    }
    _ = len(v1)
    switch -2147483648 {
    case 2147483647:
        if (4294967295 == 32768) {
            if false {
                if true {
                    for fi25 := 0; fi25 < 16; fi25++ {
                        if (false || (0.0 < 1.5)) {
                            v18 = v18
                        }
                        v26 := GenS2{F0: 536}
                        if true {
                            var v27 uint32 = (^(v18))
                            _ = v27
                        }
                        _ = fi25
                        _ = v26
                    }
                    v18 -= 2
                } else {
                    var v28 int64 = -1
                    _ = v28
                }
                for fi29 := 0; fi29 < 14; fi29++ {
                    v30 := v1[(fi29 >> ((65535) & 63))]
                    fmt.Println(v18, v18, fi29)
                    _ = fi29
                    _ = v30
                }
            } else {
                var v31 uint8 = uint8(v18)
                v1[42] = -317
                var v32 int8 = int8(v18)
                _ = v31
                _ = v32
            }
            var v33 uint8 = 31
            _ = v33
        }
    case 591:
        var v34 uint32 = v18
        _ = v34
    case -114:
        v18 = v18
    default:
        v35 := GenS2{F0: 42}
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint16; F2 string }
func genFn0(p0 bool, p1 bool, p2 uint16) int {
    if p1 {
        if (false && true) {
            var v6 int16 = 382
            p0 = (!(1 <= 32767))
            p0 = p1
            _ = v6
        }
        p1 = p0
    }
    v7 := GenS1{F0: 18446744073709551615, F1: -2147483648}
    var v8 uint8 = 74
    v9 := []int{-889, 127, 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 892
}
func genFn1() uint16 {
    var v10 int64 = 9223372036854775807
    v11 := GenS2{F0: 1, F1: 32767, F2: ("\n\t" + "a")}
    v10 -= v10
    switch 32767 {
    case 2147483647:
        v10 = (v10 / ((v10) | 1))
    }
    _ = v10
    _ = v11
    return 255
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 0
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]uint32)
    var v5 int64 = 9223372036854775807
    v12 := []int{-996, 32768, -603}
    v13 := v3[0]
    v14 := GenS2{F0: v13, F1: genFn1(), F2: "hello"}
    var v15 int8 = v2
    var v16 uint8 = 134
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    v13 = (v13 >> ((838) & 63))
    for fi17 := 0; fi17 < 3; fi17++ {
        v3[650] = 0
        fmt.Println(v5, v15)
        _ = fi17
    }
    var v18 int64 = (v5 * -1)
    var v19 uint32 = (v13 << ((469) & 63))
    for fi20 := 0; fi20 < 4; fi20++ {
        if false {
            var v21 bool = true
            _ = v21
        }
        for fi22 := 0; fi22 < 5; fi22++ {
            if (genFn0(false, false, 496) > fi22) {
                var v23 uint = 0
                v19 -= v19
                var v24 int8 = (v2 >> (((v23 << ((0) & 63))) & 63))
                v5 = (v18 & (v18 << ((697) & 63)))
                switch "\n\t" {
                case "foo bar":
                    var v25 uint16 = 32767
                    _ = v25
                case "hello":
                    v12 = append(v12, 0)
                default:
                    v26 := GenS1{F0: 18446744073709551615, F1: 0}
                    _ = v26
                }
                v27 := GenS2{F0: uint32(v16), F1: 2, F2: ("" + (("hello" + "\n\t") + ""))}
                _ = v23
                _ = v24
                _ = v27
            } else {
                v2 = (v15 * v2)
            }
            v28 := GenS1{F0: 18446744073709551615, F1: 1}
            _ = fi22
            _ = v28
        }
        if true {
            var v29 float64 = -2.25
            _ = v29
        }
        _ = fi20
    }
    v5 = (v18 % ((-850) | 1))
    var v30 int = genFn0(true, (float32(((0.0 - 0.1) * 3.14159)) != 3.14159), 807)
    if true {
        for fi31 := 0; fi31 < 11; fi31++ {
            if (682 > 567) {
                var v32 bool = true
                _ = v32
            }
            v1["a"] = 751
            _ = fi31
        }
        var v33 uint64 = 2
        fmt.Println(v15, v5)
        _ = v33
    } else {
        v34 := GenS1{F0: 759, F1: -452}
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}

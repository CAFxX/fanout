package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int64 }
func genFn0() int64 {
    var v7 uint32 = 2147483648
    for fi8 := 0; fi8 < 15; fi8++ {
        switch ("" + ("" + "a")) {
        case "hello":
            var v9 bool = false
            _ = v9
        case "":
            v10 := GenS1{F0: 42, F1: 0, F2: 42}
            _ = v10
        }
        if false {
            v11 := GenS1{F0: 9223372036854775807, F1: 0, F2: 0}
            v7 += (v7 / ((v7) | 1))
            switch (fi8 * fi8) {
            case 9223372036854775807:
                v12 := GenS1{F0: 42, F1: -1, F2: 9223372036854775807}
                _ = v12
            }
            _ = v11
        } else {
            v13 := []int{}
            switch -354 {
            case 0:
                switch ("foo bar" + "foo bar") {
                case "foo bar":
                    switch ("" + "hello") {
                    case "a":
                        var v14 float32 = -2.25
                        _ = v14
                    case "foo bar":
                        v13 = append(v13, (fi8 << ((42) & 63)))
                    case "hello":
                        for fi15 := 0; fi15 < 1; fi15++ {
                            v7 = 4294967295
                            _ = fi15
                        }
                    }
                default:
                    if false {
                        v16 := GenS1{F0: -9223372036854775808, F1: -128, F2: -592}
                        _ = v16
                    }
                }
            case -811:
                v7 = 65535
            default:
                var v17 uint8 = 216
                _ = v17
            }
            v7 += (v7 | 52)
            _ = v13
        }
        v18 := []int{fi8}
        _ = fi8
        _ = v18
    }
    v19 := []int{489, -869, -9223372036854775808, 32767}
    _ = v19
    _ = v7
    return -576
}
func genFn1() bool {
    var v20 int8 = -1
    var v21 float32 = 0.0
    v21 -= float32(3.14159)
    switch 32767 {
    case 664:
        for fi22 := 0; fi22 < 13; fi22++ {
            if true {
                var v23 int32 = -634
                var v24 int64 = genFn0()
                _ = v23
                _ = v24
            }
            if (((1.5 / (v21 * (-2.25 - v21))) / 1e10) <= 1.5) {
                var v25 int = fi22
                _ = v25
            } else {
                v26 := GenS1{F0: 1, F1: v20, F2: 1}
                v21 = float32(0.0)
                if true {
                    v21 -= (-2.25 * 0.0)
                }
                _ = v26
            }
            _ = fi22
        }
    default:
        for fi27 := 0; fi27 < 7; fi27++ {
            var v28 uint = 4294967295
            v29 := []int{}
            _ = fi27
            _ = v28
            _ = v29
        }
    }
    _ = v20
    _ = v21
    return false
}

func main() {
    var v1 int16 = 1000
    var v2 uint32 = 1
    v3 := make(map[int]bool)
    v4 := make(map[uint32]string)
    var v5 float32 = 1e10
    v6 := []string{("foo bar" + "\n\t"), ("a" + ("a" + "\n\t"))}
    var v30 int16 = v1
    var v31 uint16 = uint16(v2)
    v3[2147483647] = (("hello" + "a") < ("" + "foo bar"))
    if (true || (680 <= 2)) {
        v31 = (v31 - v31)
        var v32 float64 = 3.14159
        var v33 int8 = -128
        _ = v32
        _ = v33
    }
    if (((v1 ^ (^(v1))) > (^(v1))) && ((true || false) && (v5 <= v5))) {
        fmt.Println(v30)
        v31 -= (v31 * v31)
    }
    switch 32767 {
    case 765:
        _ = len(v4)
    case -9223372036854775808:
        switch -16 {
        case 9223372036854775807:
            var v34 int16 = (v1 >> ((889) & 63))
            _ = v34
        }
    default:
        fmt.Println(v31)
    }
    fmt.Println(v2, v30, v31)
    var v35 uint64 = 486
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}

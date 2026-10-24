package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint16; F1 uint16 }
func genFn0() int16 {
    for fi5 := 0; fi5 < 19; fi5++ {
        for fi6 := 0; fi6 < 19; fi6++ {
            if ((209 > 2147483648) || (42 >= 1)) {
                fmt.Println(fi5, fi5)
                if false {
                    v7 := []int{(fi5 ^ fi5)}
                    fmt.Println(fi5, fi5)
                    _ = v7
                }
            } else {
                var v8 int64 = -1
                if false {
                    v9 := []int{fi5}
                    if len(v9) > 0 {
                        v9[0] = -704
                    }
                    _ = v9
                } else {
                    var v10 int32 = 1
                    v8 = v8
                    var v11 uint32 = 42
                    _ = v10
                    _ = v11
                }
                _ = v8
            }
            var v12 uint32 = 2147483647
            v13 := []int{}
            _ = fi6
            _ = v12
            _ = v13
        }
        _ = fi5
    }
    switch ("hello" + (("hello" + "hello") + "\n\t")) {
    case "":
        v14 := []int{9223372036854775807, 1, 2147483647, 462}
        _ = v14
    default:
        switch 32767 {
        case 42:
            if false {
                var v15 uint16 = 0
                fmt.Println(v15)
                _ = v15
            }
        case 127:
            var v16 uint8 = 233
            _ = v16
        case 330:
            if true {
                v17 := GenS1{F0: -2147483648, F1: "a"}
                var v18 int = -1
                for fi19 := 0; fi19 < 16; fi19++ {
                    fmt.Println(v18, fi19)
                    _ = fi19
                }
                _ = v17
                _ = v18
            } else {
                var v20 uint8 = 2
                _ = v20
            }
        }
    }
    var v21 string = "foo bar"
    switch 32767 {
    case 2147483647:
        v22 := GenS2{F0: 42}
        _ = v22
    case 127:
        v23 := GenS1{F0: 0, F1: (((v21 + "hello") + v21) + v21)}
        _ = v23
    default:
        v24 := GenS3{F0: v21, F1: 42, F2: 102}
        _ = v24
    }
    fmt.Println(v21)
    _ = v21
    return 1
}
func genFn1(p0 uint, p1 string) uint64 {
    var v25 float64 = (0.1 * 3.14159)
    p0 = (^(p0))
    _ = p0
    _ = p1
    _ = v25
    return 42
}

func main() {
    var v1 int = 127
    v2 := []uint32{}
    var v3 int16 = -32768
    v4 := []int{127, v1, 32768, v1}
    if true {
        fmt.Println(v1, v1)
        v1 -= (v1 << ((18446744073709551615) & 63))
    }
    if len(v4) > 0 {
        v4[0] = v1
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v26 float32 = 0.0
    for fi27 := 0; fi27 < 7; fi27++ {
        v26 = v26
        _ = fi27
    }
    v28 := GenS4{F0: 622, F1: 140}
    v1 = v1
    for fi29 := 0; fi29 < 13; fi29++ {
        v1 += (v1 & v1)
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v4
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint16 }
type GenS2 struct { F0 uint; F1 int8; F2 string }
func genFn0(p0 bool, p1 int32) string {
    p1 = p1
    var v7 string = "foo bar"
    v8 := []int{32767, 896, -2147483648, 1}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return v7
}
func genFn1() int {
    var v9 uint16 = 1
    var v10 int8 = 63
    fmt.Println(v10, v10)
    var v11 int8 = -87
    v12 := []int{-730, 1}
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 32767
}

func main() {
    v1 := []uint32{}
    var v2 uint32 = 2147483647
    var v3 uint8 = uint8(v2)
    v4 := make(map[uint32]int)
    v5 := []int{0}
    v6 := make(map[int]string)
    var v13 int32 = 22
    var v14 int = 32768
    v14 += v14
    v13 = (v13 << ((1) & 63))
    var v15 uint8 = (^(v3))
    v16 := GenS1{F0: (("hello" + "foo bar") + (("" + ((("foo bar" + "hello") + "foo bar") + ("a" + ((("hello" + ("a" + "a")) + "foo bar") + ("hello" + "\n\t"))))) + "\n\t")), F1: (v14 & (v14 >> ((65535) & 63))), F2: 1}
    v17 := GenS1{F0: "\n\t", F1: v14, F2: 2}
    for fi18 := 0; fi18 < 6; fi18++ {
        v15 = v3
        var v19 int = genFn1()
        _ = fi18
        _ = v19
    }
    v20 := GenS2{F0: 904, F1: -80, F2: ("a" + "a")}
    var v21 int64 = -234
    var v22 uint16 = 2
    v13 = v13
    if (("hello" + ("foo bar" + ("\n\t" + ""))) > "") {
        var v23 uint32 = (v2 << ((18446744073709551615) & 63))
        for fi24 := 0; fi24 < 5; fi24++ {
            v25 := GenS1{F0: "\n\t", F1: fi24, F2: uint16(v3)}
            v4[(^(v2))] = 899
            _ = fi24
            _ = v25
        }
        _ = v23
    }
    if ("foo bar" < "\n\t") {
        v22 = 2
    } else {
        v22 += (v22 / (((v22 / (((^(v22))) | 1))) | 1))
    }
    v22 = (v22 - v22)
    v3 = v3
    for fi26 := 0; fi26 < 6; fi26++ {
        for fi27 := 0; fi27 < 11; fi27++ {
            var v28 uint32 = v2
            _ = fi27
            _ = v28
        }
        var v29 uint = 65535
        if (v2 == v2) {
            var v30 float32 = 3.14159
            _ = v30
        } else {
            for fi31 := 0; fi31 < 15; fi31++ {
                fmt.Println(v14, v2)
                _ = fi31
            }
        }
        _ = fi26
        _ = v29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

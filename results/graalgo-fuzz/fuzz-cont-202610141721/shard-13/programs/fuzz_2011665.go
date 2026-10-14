package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 int64; F2 bool }
func genFn0() uint8 {
    switch -9223372036854775808 {
    case 32767:
        switch "" {
        case "hello":
            var v4 int32 = -850
            _ = v4
        case "\n\t":
            var v5 uint64 = 18446744073709551615
            _ = v5
        default:
            if false {
                var v6 float32 = 1.5
                _ = v6
            }
        }
    }
    var v7 int = 365
    for fi8 := 0; fi8 < 18; fi8++ {
        if (-102 != 127) {
            v7 = 128
        }
        v9 := []int{}
        _ = fi8
        _ = v9
    }
    _ = v7
    return 185
}
func genFn1(p0 uint8, p1 int32, p2 int8) int8 {
    var v10 uint = 2
    v11 := GenS1{F0: (p2 | (p2 ^ p2))}
    p0 += (p0 << ((1) & 63))
    p2 = p2
    for fi12 := 0; fi12 < 1; fi12++ {
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 uint = (v10 ^ 42)
            v15 := []int{(fi13 * fi12), fi13, (fi12 * fi13)}
            _ = fi13
            _ = v14
            _ = v15
        }
        var v16 float32 = (1e10 / -2.25)
        _ = fi12
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return p2
}

func main() {
    var v1 int32 = 221
    v2 := make(map[string]string)
    var v3 uint64 = 604
    v3 = (v3 / (((v3 % ((v3) | 1))) | 1))
    fmt.Println(v3, v3)
    for fi17 := 0; fi17 < 13; fi17++ {
        v2[("hello" + "")] = "a"
        v1 = v1
        _ = fi17
    }
    var v18 string = ((("\n\t" + ("a" + "foo bar")) + "foo bar") + "hello")
    fmt.Println(v3)
    if true {
        for fi19 := 0; fi19 < 7; fi19++ {
            for fi20 := 0; fi20 < 12; fi20++ {
                v21 := GenS1{F0: genFn1(127, (v1 ^ (v1 >> ((0) & 63))), -9)}
                fmt.Println(v1, fi19, v3)
                v22 := GenS1{F0: 42}
                if false {
                    v1 = v1
                }
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = fi19
        }
        for fi23 := 0; fi23 < 14; fi23++ {
            for fi24 := 0; fi24 < 9; fi24++ {
                v1 = 171
                switch (v18 + "a") {
                case "hello":
                    v25 := GenS2{F0: v18, F1: 9223372036854775807, F2: true}
                    _ = v25
                }
                _ = fi24
            }
            _ = fi23
        }
        var v26 int16 = 32767
        _ = v26
    }
    v27 := v2[v18]
    fmt.Println(v1)
    for fi28 := 0; fi28 < 18; fi28++ {
        v1 = v1
        v29 := GenS1{F0: 126}
        for fi30 := 0; fi30 < 2; fi30++ {
            v31 := []int{fi30, (-(fi30)), (fi30 << ((255) & 63))}
            var v32 uint32 = 0
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = fi28
        _ = v29
    }
    var v33 int16 = 32767
    var v34 int32 = (v1 & int32(v3))
    v18 = ""
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v34
}

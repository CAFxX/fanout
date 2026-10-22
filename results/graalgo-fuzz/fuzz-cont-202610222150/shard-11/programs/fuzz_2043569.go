package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int64; F1 string; F2 int64 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 int32 }
func genFn0(p0 uint32) uint16 {
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := GenS1{F0: (p0 - 2147483648), F1: (p0 >> ((554) & 63))}
        _ = fi5
        _ = v6
    }
    fmt.Println(p0)
    var v7 uint8 = 128
    _ = p0
    _ = v7
    return 32767
}
func genFn1() uint8 {
    var v8 bool = true
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := GenS1{F0: 1, F1: 2147483648}
        v8 = (("foo bar" + (("a" + "foo bar") + "hello")) >= "a")
        _ = fi9
        _ = v10
    }
    v11 := []int{}
    fmt.Println(v8)
    v8 = (v8 && ((v8 || (v8 || (260 > 2))) && v8))
    _ = v11
    _ = v8
    return 228
}

func main() {
    var v1 int16 = -32768
    var v2 uint8 = 2
    var v3 int64 = 1
    var v4 float64 = -2.25
    fmt.Println(v3, v2)
    v2 = (v2 & v2)
    v12 := []int{}
    v13 := GenS2{F0: (v3 ^ (v3 - v3)), F1: "a", F2: (v3 % ((42) | 1))}
    v14 := GenS2{F0: (v3 * (v3 ^ v3)), F1: (("hello" + "hello") + "\n\t"), F2: (v3 << ((2) & 63))}
    fmt.Println(v3, v3, v2)
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    v12 = append(v12, 4294967295)
    var v15 bool = (548 < 2)
    for fi16 := 0; fi16 < 8; fi16++ {
        for fi17 := 0; fi17 < 19; fi17++ {
            v18 := GenS3{F0: -273, F1: 2}
            _ = fi17
            _ = v18
        }
        v4 = v4
        _ = fi16
    }
    if len(v12) > 0 {
        v12[0] = -9223372036854775808
    }
    var v19 int8 = -37
    if ((uint64(v2) <= 2) || v15) {
        if v15 {
            v20 := GenS4{F0: false, F1: 2}
            fmt.Println(v1)
            v21 := GenS2{F0: 1, F1: "a", F2: v3}
            _ = v20
            _ = v21
        } else {
            if true {
                switch "" {
                case "a":
                    var v22 int32 = -1
                    _ = v22
                case "\n\t":
                    for fi23 := 0; fi23 < 6; fi23++ {
                        v4 = v4
                        _ = fi23
                    }
                }
            } else {
                v12 = append(v12, 2147483647)
            }
            v1 = (v1 ^ (v1 - v1))
            v24 := GenS3{F0: -88, F1: 501}
            _ = v24
        }
    }
    if (v2 >= (v2 ^ (^(v2)))) {
        v25 := GenS5{F0: -800}
        _ = v25
    }
    v1 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}

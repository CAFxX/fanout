package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint64 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int; F1 uint16; F2 uint }
func genFn0() int {
    var v7 int16 = 32767
    v8 := []int{-810, 42, 2147483647}
    for fi9 := 0; fi9 < 9; fi9++ {
        if (!true) {
            for fi10 := 0; fi10 < 18; fi10++ {
                fmt.Println(v7, fi9)
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = fi10
            }
        }
        for fi11 := 0; fi11 < 2; fi11++ {
            v12 := GenS1{F0: "foo bar", F1: (v7 - int16(fi11)), F2: 723}
            var v13 uint32 = 0
            _ = fi11
            _ = v12
            _ = v13
        }
        var v14 uint32 = 0
        _ = fi9
        _ = v14
    }
    switch ((((("a" + "") + (("hello" + "\n\t") + "hello")) + "") + "") + "\n\t") {
    case "hello":
        _ = len(v8)
    case "a":
        if true {
            for fi15 := 0; fi15 < 9; fi15++ {
                v7 += -32768
                for fi16 := 0; fi16 < 14; fi16++ {
                    var v17 uint = 480
                    fmt.Println(v7)
                    v17 = v17
                    _ = fi16
                    _ = v17
                }
                _ = fi15
            }
            v7 = v7
        }
    default:
        var v18 string = "hello"
        _ = v18
    }
    v7 -= (v7 * 32767)
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 int8 = 127
    v2 := make(map[int]int)
    var v3 int16 = 42
    var v4 int8 = (v1 << ((42) & 63))
    v5 := []string{"\n\t", ("\n\t" + "hello"), (((("a" + ("hello" + ("foo bar" + ("\n\t" + ("\n\t" + "hello"))))) + "hello") + "hello") + "\n\t")}
    var v6 int32 = 1
    var v19 float32 = 1e10
    for fi20 := 0; fi20 < 11; fi20++ {
        v21 := GenS2{F0: v3, F1: true}
        _ = fi20
        _ = v21
    }
    fmt.Println(v3, v3)
    v22 := GenS3{F0: 128, F1: 714, F2: 65535}
    if (!false) {
        v23 := GenS3{F0: 9223372036854775807, F1: 884, F2: 42}
        fmt.Println(v1, v3, v1)
        for fi24 := 0; fi24 < 6; fi24++ {
            v2[genFn0()] = fi24
            _ = fi24
        }
        _ = v23
    }
    fmt.Println(v19)
    v19 += 0.1
    var v25 int8 = (-(v1))
    for fi26 := 0; fi26 < 6; fi26++ {
        if true {
            v27 := GenS2{F0: (v3 % ((-32768) | 1)), F1: true}
            _ = v27
        }
        v3 += v3
        fmt.Println(v25, v3, v3)
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

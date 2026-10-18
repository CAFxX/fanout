package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint; F1 uint32 }
func genFn0(p0 uint32) int16 {
    v4 := GenS1{F0: (!true), F1: 128}
    for fi5 := 0; fi5 < 15; fi5++ {
        p0 = (p0 & (^(p0)))
        v6 := GenS1{F0: (true && true), F1: 76}
        _ = fi5
        _ = v6
    }
    if false {
        v7 := GenS2{F0: 42, F1: (^(p0))}
        v8 := []int{}
        if len(v8) > 0 {
            v8[0] = 860
        }
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 20; fi9++ {
            v10 := []int{int(p0), 128, fi9, 11}
            _ = fi9
            _ = v10
        }
        if true {
            if ((0 >= 42) || false) {
                var v11 uint = 0
                _ = v11
            } else {
                fmt.Println(p0)
            }
        } else {
            v12 := []int{}
            _ = v12
        }
    }
    p0 += 2147483647
    v13 := GenS1{F0: false, F1: 2}
    _ = p0
    _ = v13
    _ = v4
    return 32767
}
func genFn1(p0 uint8, p1 int64) int32 {
    fmt.Println(p0, p0)
    p0 += p0
    var v14 uint8 = (p0 + 0)
    for fi15 := 0; fi15 < 2; fi15++ {
        p0 = v14
        _ = fi15
    }
    if (p1 < (p1 % ((p1) | 1))) {
        fmt.Println(p0, v14)
        p0 = (p0 << ((42) & 63))
        if true {
            for fi16 := 0; fi16 < 15; fi16++ {
                p0 = 255
                _ = fi16
            }
        }
    } else {
        for fi17 := 0; fi17 < 20; fi17++ {
            v18 := GenS2{F0: 2, F1: 65535}
            v19 := GenS1{F0: false, F1: (^(p0))}
            v20 := []int{(fi17 >> ((181) & 63)), fi17, fi17}
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        var v21 uint = 427
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v14
    return 2147483647
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = -9223372036854775808
    var v3 uint16 = 32767
    v2 = (v2 - -30)
    v2 -= (^(v2))
    var v22 int16 = genFn0(42)
    fmt.Println(v22, v3)
    var v23 float32 = 0.0
    fmt.Println(v22, v3, v2)
    fmt.Println(v2)
    v23 += (0.0 * v23)
    fmt.Println(v22, v3, v3)
    v23 = 0.1
    fmt.Println(v22)
    for fi24 := 0; fi24 < 15; fi24++ {
        _ = len(v1)
        v25 := GenS2{F0: 191, F1: 2}
        _ = fi24
        _ = v25
    }
    fmt.Println(v23)
    switch ("a" + "\n\t") {
    case "hello":
        switch ("hello" + "\n\t") {
        case "foo bar":
            if (v22 != v22) {
                v1[65535] = ""
                fmt.Println(v2)
            }
        }
    }
    var v26 int = -373
    v27 := GenS2{F0: 42, F1: uint32(v3)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v3
}

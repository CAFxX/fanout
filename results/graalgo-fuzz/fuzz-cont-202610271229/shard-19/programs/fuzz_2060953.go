package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint64; F1 int32; F2 string }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 uint32; F1 uint16; F2 int16 }

func main() {
    var v1 uint16 = 1
    var v2 uint32 = 0
    var v3 int64 = 42
    v4 := []uint32{}
    var v5 int32 = 435
    var v6 uint = 65535
    v2 += v2
    if true {
        switch 42 {
        case 32767:
            v7 := GenS1{F0: false, F1: ((((("" + ("hello" + "hello")) >= (("" + (("foo bar" + "a") + (("a" + (("a" + "") + "hello")) + "hello"))) + "a")) && ("hello" < "\n\t")) || true) || (156 > 158))}
            _ = v7
        case 127:
            v2 += v2
        case 4294967295:
            v4 = append(v4, 0)
        default:
            v8 := GenS2{F0: 18446744073709551615, F1: (^(v5)), F2: "hello"}
            _ = v8
        }
    } else {
        for fi9 := 0; fi9 < 2; fi9++ {
            for fi10 := 0; fi10 < 16; fi10++ {
                v11 := GenS2{F0: 0, F1: v5, F2: "foo bar"}
                _ = len(v4)
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
        var v12 uint8 = 255
        var v13 int64 = -724
        _ = v12
        _ = v13
    }
    v4 = append(v4, v2)
    v14 := GenS1{F0: ((v5 & (^(v5))) >= v5), F1: ((96 >= 100) || false)}
    _ = len(v4)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 += v3
    if (v6 < v6) {
        v5 += (v5 | v5)
        if (!(0.1 <= (1e10 / 0.1))) {
            var v15 int16 = 0
            _ = v15
        }
    } else {
        v16 := GenS3{F0: 4294967295, F1: (v2 * v2)}
        v3 = 190
        var v17 int8 = -128
        _ = v16
        _ = v17
    }
    v3 -= v3
    v18 := GenS4{F0: (v2 - v2), F1: 1, F2: 1000}
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

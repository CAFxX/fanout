package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint32; F1 int }
func genFn0(p0 uint8) uint64 {
    var v7 uint64 = 2
    v7 += (v7 * v7)
    _ = p0
    _ = v7
    return (v7 | 696)
}

func main() {
    v1 := []int{-890}
    var v2 int8 = -23
    var v3 int64 = -616
    v4 := []uint32{130, 4294967295}
    v5 := []string{"", ("a" + "")}
    v6 := make(map[uint32]bool)
    fmt.Println(v3)
    v3 += (v3 ^ v3)
    if false {
        for fi8 := 0; fi8 < 3; fi8++ {
            v2 -= v2
            switch ((("hello" + "hello") + ("foo bar" + ("foo bar" + "hello"))) + "") {
            case "a":
                var v9 float32 = 0.0
                _ = v9
            default:
                var v10 uint32 = 4294967295
                _ = v10
            }
            _ = fi8
        }
        for fi11 := 0; fi11 < 13; fi11++ {
            if false {
                v12 := GenS1{F0: -32768, F1: (-(v2))}
                var v13 int = (fi11 & fi11)
                v14 := GenS2{F0: 255, F1: ((fi11 >> ((42) & 63)) < fi11), F2: 32767}
                _ = v12
                _ = v13
                _ = v14
            } else {
                _ = len(v6)
            }
            _ = fi11
        }
    } else {
        v15 := GenS1{F0: 32767, F1: 0}
        if (0 > 1) {
            v6[2147483647] = (!false)
        }
        var v16 float32 = (3.14159 / float32(((float64(3.14159) + -2.25) / (1.5 - 0.1))))
        _ = v15
        _ = v16
    }
    var v17 string = ("foo bar" + ("a" + (((("hello" + "a") + "\n\t") + "a") + "foo bar")))
    v18 := GenS1{F0: int16(v3), F1: v2}
    v17 = v17
    v3 -= (v3 * -9223372036854775808)
    fmt.Println(v3, v2)
    var v19 int = 1
    var v20 uint32 = 65535
    var v21 uint = 255
    var v22 uint32 = (^(v20))
    v19 += 32767
    v23 := GenS3{F0: 128, F1: int32(v22)}
    for fi24 := 0; fi24 < 19; fi24++ {
        if false {
            v25 := GenS4{F0: (^(v22)), F1: fi24}
            delete(v6, (v20 * (v20 >> ((v21) & 63))))
            _ = v25
        }
        if true {
            for fi26 := 0; fi26 < 16; fi26++ {
                fmt.Println(v17)
                _ = fi26
            }
            if (!(!false)) {
                v3 = v3
                var v27 uint16 = 733
                v28 := GenS1{F0: -700, F1: (v2 - v2)}
                _ = v27
                _ = v28
            }
            v2 = v2
        }
        v29 := GenS3{F0: 113, F1: 42}
        _ = fi24
        _ = v29
    }
    _ = len(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

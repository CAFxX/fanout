package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 bool }
type GenS2 struct { F0 int }

func main() {
    var v1 int = 32768
    v2 := []uint32{4294967295}
    var v3 uint16 = 42
    var v4 uint16 = v3
    var v5 uint16 = v4
    v6 := make(map[uint32]uint32)
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint64 = uint64(v3)
        switch 0 {
        case -2147483648:
            var v9 string = "a"
            _ = v9
        case 9223372036854775807:
            var v10 float32 = 0.0
            _ = v10
        case 0:
            if (!false) {
                switch (("hello" + "") + ("hello" + "")) {
                case "a":
                    for fi11 := 0; fi11 < 12; fi11++ {
                        v8 -= (v8 & v8)
                        _ = fi11
                    }
                case "":
                    var v12 uint32 = 4294967295
                    _ = v12
                default:
                    v13 := GenS1{F0: 81, F1: "", F2: (!(!false))}
                    _ = v13
                }
                var v14 uint16 = 0
                var v15 string = ""
                _ = v14
                _ = v15
            }
        default:
            if true {
                var v16 uint32 = 42
                var v17 int64 = 42
                _ = v16
                _ = v17
            }
        }
        _ = fi7
        _ = v8
    }
    var v18 uint16 = v5
    v5 = (v5 - (v4 - 2))
    v4 += (v3 ^ 154)
    fmt.Println(v1)
    var v19 int64 = 1
    v18 = v3
    fmt.Println(v18, v3, v5)
    v20 := GenS2{F0: (v1 / ((v1) | 1))}
    for fi21 := 0; fi21 < 1; fi21++ {
        for fi22 := 0; fi22 < 9; fi22++ {
            v23 := GenS1{F0: 118, F1: "hello", F2: (-838 <= 0)}
            var v24 bool = false
            _ = len(v2)
            _ = fi22
            _ = v23
            _ = v24
        }
        var v25 uint16 = (v4 ^ (^(v18)))
        if ((float32(3.14159) * 0.0) >= (((-2.25 * 0.1) / 0.1) - -2.25)) {
            var v26 uint32 = 0
            v18 -= 389
            _ = v26
        }
        _ = fi21
        _ = v25
    }
    fmt.Println(v4)
    v6[2147483647] = 359
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

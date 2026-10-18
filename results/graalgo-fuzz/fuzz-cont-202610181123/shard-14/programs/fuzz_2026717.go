package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint32 }
type GenS2 struct { F0 int8; F1 bool; F2 int }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 int16; F1 uint32; F2 int64 }
type GenS5 struct { F0 bool; F1 uint32; F2 int16 }
type GenS6 struct { F0 string; F1 int8 }
type GenS7 struct { F0 int8; F1 bool; F2 int }
func genFn0(p0 uint64) uint32 {
    v5 := GenS1{F0: ("a" + ("\n\t" + "\n\t")), F1: 65535, F2: 42}
    v6 := []int{-290, 1, 9223372036854775807, -9223372036854775808}
    switch "hello" {
    case "hello":
        p0 -= p0
    case "\n\t":
        for fi7 := 0; fi7 < 8; fi7++ {
            for fi8 := 0; fi8 < 17; fi8++ {
                if true {
                    v9 := GenS1{F0: "hello", F1: 2147483648, F2: 2}
                    if true {
                        var v10 int16 = -829
                        p0 = 460
                        v11 := GenS2{F0: -128, F1: true, F2: 0}
                        _ = v10
                        _ = v11
                    }
                    _ = len(v6)
                    _ = v9
                }
                v12 := GenS2{F0: 42, F1: true, F2: fi8}
                v13 := GenS3{F0: 0, F1: 0}
                _ = fi8
                _ = v12
                _ = v13
            }
            _ = len(v6)
            _ = fi7
        }
    case "a":
        v14 := GenS4{F0: 42, F1: 0, F2: -1}
        _ = v14
    default:
        v15 := GenS5{F0: (255 > 1), F1: 2147483647, F2: 42}
        _ = v15
    }
    _ = p0
    _ = v5
    _ = v6
    return 325
}

func main() {
    var v1 int = -1
    var v2 uint16 = 0
    var v3 uint8 = 42
    v4 := []string{("hello" + "foo bar")}
    var v16 uint64 = 4294967295
    var v17 int = -893
    var v18 uint8 = (^(v3))
    var v19 int16 = -1
    v20 := GenS1{F0: "", F1: 337, F2: 75}
    v21 := GenS6{F0: "", F1: 100}
    var v22 uint8 = 19
    v2 -= 32767
    var v23 int64 = -1
    v19 += v19
    fmt.Println(v23, v18)
    v24 := GenS7{F0: -45, F1: true, F2: (v17 ^ v17)}
    var v25 int8 = -1
    fmt.Println(v23)
    v3 = uint8(v16)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}

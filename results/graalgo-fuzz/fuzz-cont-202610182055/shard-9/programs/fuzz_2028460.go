package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 uint8; F1 int16; F2 bool }
func genFn0(p0 bool, p1 uint) uint64 {
    var v4 uint = 250
    var v5 int8 = 1
    var v6 int16 = -849
    v4 -= (p1 >> (((v4 - (^(v4)))) & 63))
    v7 := []int{0, 2147483647}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[string]string)
    var v2 int32 = -2147483648
    v3 := make(map[string]uint32)
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        switch "a" {
        case "\n\t":
            var v8 int8 = -88
            _ = v8
        default:
            v2 -= v2
        }
    case "hello":
        for fi9 := 0; fi9 < 1; fi9++ {
            v2 += int32(fi9)
            _ = fi9
        }
    default:
        var v10 int = -2147483648
        _ = v10
    }
    var v11 float32 = 0.0
    v11 = v11
    var v12 int16 = 942
    v13 := GenS1{F0: 551, F1: 255}
    switch -2147483648 {
    case 32767:
        v14 := GenS2{F0: 400}
        _ = v14
    }
    var v15 int64 = 42
    for fi16 := 0; fi16 < 16; fi16++ {
        if true {
            for fi17 := 0; fi17 < 14; fi17++ {
                v12 += v12
                v18 := GenS3{F0: v2, F1: v15}
                if (!(((v11 + (v11 * (v11 + v11))) / 1e10) >= v11)) {
                    switch ("foo bar" + "foo bar") {
                    case "hello":
                        var v19 uint64 = 42
                        _ = v19
                    case "a":
                        v20 := GenS4{F0: 92, F1: 186, F2: (42 < 0)}
                        _ = v20
                    case "":
                        v21 := GenS2{F0: 65535}
                        _ = v21
                    default:
                        v22 := []int{}
                        _ = v22
                    }
                    fmt.Println(v11, v12)
                    var v23 float32 = v11
                    _ = v23
                }
                v24 := []int{(fi17 >> ((951) & 63)), -2147483648}
                if (false && (293 == genFn0(true, 266))) {
                    v25 := GenS1{F0: 737, F1: 65535}
                    _ = v25
                } else {
                    var v26 string = "hello"
                    if len(v24) > 0 {
                        v24[0] = fi17
                    }
                    _ = v26
                }
                fmt.Println(v15, v2, v2)
                v27 := GenS1{F0: 589, F1: 326}
                _ = fi17
                _ = v18
                _ = v24
                _ = v27
            }
        }
        _ = fi16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
}

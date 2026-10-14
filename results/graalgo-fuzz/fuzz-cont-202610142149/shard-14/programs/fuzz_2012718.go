package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 int; F2 uint }
type GenS4 struct { F0 int32; F1 string; F2 string }
func genFn0(p0 uint8, p1 uint8) int16 {
    switch -244 {
    case -9223372036854775808:
        switch "" {
        case "a":
            for fi5 := 0; fi5 < 19; fi5++ {
                var v6 float32 = ((0.1 * 1.5) / 3.14159)
                _ = fi5
                _ = v6
            }
        }
    case 127:
        p0 = 2
    case 0:
        p0 += (^(p1))
    default:
        v7 := []int{2147483647, 2147483647}
        _ = v7
    }
    switch ("" + "") {
    case "":
        var v8 int16 = 0
        _ = v8
    default:
        p0 -= 127
    }
    _ = p0
    _ = p1
    return 32767
}

func main() {
    v1 := []string{}
    var v2 int8 = 100
    var v3 uint64 = 4294967295
    v4 := []uint32{}
    if (9223372036854775807 <= 423) {
        var v9 int32 = -2147483648
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        if len(v4) > 0 {
            v4[0] = 2
        }
        v11 := GenS1{F0: 0, F1: ("hello" + ""), F2: 130}
        _ = fi10
        _ = v11
    }
    v12 := GenS2{F0: 1}
    v3 = 42
    switch "\n\t" {
    case "a":
        v13 := GenS3{F0: -1, F1: 884, F2: 255}
        _ = v13
    case "foo bar":
        var v14 uint32 = uint32(v3)
        _ = v14
    default:
        v2 = v2
    }
    v15 := GenS4{F0: 1, F1: "\n\t", F2: ("foo bar" + "\n\t")}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v16 int16 = genFn0(4, 2)
    for fi17 := 0; fi17 < 3; fi17++ {
        if len(v1) > 0 {
            v1[0] = "hello"
        }
        _ = fi17
    }
    fmt.Println(v2, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}

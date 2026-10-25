package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 uint; F2 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 int64 }
func genFn0(p0 uint8) int64 {
    p0 -= 125
    fmt.Println(p0)
    switch "hello" {
    case "foo bar":
        p0 = 128
    case "\n\t":
        v5 := GenS1{F0: 2147483648}
        _ = v5
    case "":
        fmt.Println(p0)
    default:
        var v6 int = 42
        _ = v6
    }
    _ = p0
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = 0
    var v3 int32 = 0
    v4 := make(map[string]int)
    v3 = (v3 >> ((429) & 63))
    switch (("" + "") + "hello") {
    case "":
        fmt.Println(v2, v2)
    case "foo bar":
        v7 := GenS1{F0: 65535}
        _ = v7
    default:
        switch 716 {
        case 1:
            if ((float64(0.1) / 1.5) <= (1e10 + ((1e10 - (((float64((0.1 - 0.0)) * float64(1e10)) + 1e10) + 0.1)) - 1e10))) {
                v8 := GenS2{F0: (v2 * (v2 - v2)), F1: 65535, F2: 65535}
                fmt.Println(v3, v2)
                v2 = int16(v3)
                _ = v8
            } else {
                v9 := []int{119, 128, 4294967295, 9223372036854775807}
                _ = v9
            }
        case 32768:
            var v10 int64 = genFn0(1)
            _ = v10
        case 0:
            v11 := GenS3{F0: -128}
            _ = v11
        }
    }
    switch 127 {
    case 4294967295:
        for fi12 := 0; fi12 < 17; fi12++ {
            v2 -= -32768
            for fi13 := 0; fi13 < 9; fi13++ {
                fmt.Println(v2)
                _ = fi13
            }
            _ = fi12
        }
    }
    v14 := GenS4{F0: 9223372036854775807, F1: genFn0(1)}
    var v15 float32 = float32(3.14159)
    if false {
        var v16 string = ("a" + "\n\t")
        for fi17 := 0; fi17 < 13; fi17++ {
            var v18 uint16 = 744
            _ = fi17
            _ = v18
        }
        v2 = int16(v3)
        _ = v16
    } else {
        var v19 uint32 = 720
        _ = v19
    }
    delete(v1, 2)
    if true {
        var v20 float32 = v15
        _ = v20
    }
    v3 = 0
    var v21 int32 = v3
    v21 -= v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint16; F1 string }
func genFn0() uint64 {
    v4 := GenS1{F0: -616, F1: true}
    var v5 uint8 = 60
    var v6 uint64 = 1
    v7 := GenS1{F0: 32767, F1: true}
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return (v6 << ((0) & 63))
}

func main() {
    var v1 bool = (true || (!(42 <= -2147483648)))
    var v2 float64 = 1e10
    var v3 uint = 682
    switch "a" {
    case "hello":
        var v8 int64 = 913
        _ = v8
    case "a":
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 int = 32768
            if v1 {
                switch v10 {
                case -798:
                    if (128 > 128) {
                        var v11 uint16 = 42
                        v2 = (v2 / (v2 + (((float64(0.1) - v2) - ((float64(-2.25) * (1e10 + v2)) - (0.0 * (0.1 + (v2 * v2))))) + v2)))
                        _ = v11
                    }
                }
                v12 := GenS1{F0: fi9, F1: true}
                _ = v12
            }
            _ = fi9
            _ = v10
        }
    case "":
        v13 := []int{2147483647, 128}
        _ = v13
    }
    fmt.Println(v3, v2, v1)
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 float32 = 1.5
        _ = fi14
        _ = v15
    }
    v16 := []int{-1, 32768}
    for fi17 := 0; fi17 < 6; fi17++ {
        if false {
            var v18 uint = v3
            var v19 int16 = -298
            v2 += (-2.25 - ((v2 * v2) - v2))
            _ = v18
            _ = v19
        } else {
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            var v20 string = "\n\t"
            _ = v20
        }
        var v21 uint64 = 0
        if (65 >= 0) {
            v22 := GenS2{F0: 255, F1: ("hello" + "foo bar")}
            fmt.Println(v1, v3, v21)
            if true {
                var v23 uint32 = 42
                _ = v23
            }
            _ = v22
        }
        _ = fi17
        _ = v21
    }
    switch -9223372036854775808 {
    case 9223372036854775807:
        var v24 int32 = 0
        _ = v24
    default:
        var v25 int16 = 32767
        _ = v25
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        v3 = 4294967295
        _ = fi26
    }
    if v1 {
        if v1 {
            v2 = 0.0
        }
        if len(v16) > 0 {
            v16[0] = 32768
        }
        var v27 uint16 = 2
        _ = v27
    }
    var v28 int = 127
    v1 = true
    v3 = (v3 - v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v28
    _ = v3
}

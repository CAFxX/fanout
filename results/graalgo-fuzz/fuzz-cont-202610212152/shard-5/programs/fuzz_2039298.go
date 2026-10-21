package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 uint32 }
func genFn0() uint8 {
    v4 := []int{}
    var v5 int8 = 100
    var v6 float64 = (float64(1.5) / 1e10)
    _ = v4
    _ = v5
    _ = v6
    return 255
}

func main() {
    var v1 int16 = 32767
    v2 := []uint32{0, 879, 2147483648, 4294967295}
    v3 := make(map[uint32]int)
    var v7 int32 = -550
    v8 := GenS1{F0: 2}
    switch 326 {
    case 127:
        switch 0 {
        case -9223372036854775808:
            v9 := GenS1{F0: 1}
            _ = v9
        case 128:
            for fi10 := 0; fi10 < 15; fi10++ {
                var v11 int32 = v7
                _ = fi10
                _ = v11
            }
        case 127:
            if (65535 < 98) {
                for fi12 := 0; fi12 < 6; fi12++ {
                    v13 := GenS2{F0: (v1 | (v1 >> ((394) & 63)))}
                    var v14 int32 = v7
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                var v15 int64 = 0
                v16 := GenS2{F0: (v1 % ((v1) | 1))}
                _ = v15
                _ = v16
            } else {
                for fi17 := 0; fi17 < 9; fi17++ {
                    var v18 int32 = (v7 % ((42) | 1))
                    _ = fi17
                    _ = v18
                }
            }
        }
    case -9223372036854775808:
        switch ("\n\t" + (("foo bar" + ("hello" + "a")) + ("\n\t" + ("hello" + (("hello" + "a") + "foo bar"))))) {
        case "hello":
            if len(v2) > 0 {
                v2[0] = 2
            }
        case "":
            v19 := GenS1{F0: 438}
            _ = v19
        case "\n\t":
            v7 += v7
        default:
            v1 = (v1 % (((v1 >> ((931) & 63))) | 1))
        }
    default:
        v1 += (^(v1))
    }
    switch "hello" {
    case "foo bar":
        var v20 float32 = float32((float64(-2.25) - 0.1))
        _ = v20
    case "":
        v21 := GenS1{F0: 0}
        _ = v21
    }
    v2 = append(v2, 290)
    switch ("\n\t" + (("" + ("hello" + "")) + "foo bar")) {
    case "foo bar":
        var v22 string = "hello"
        _ = v22
    case "":
        _ = len(v2)
    case "\n\t":
        v23 := GenS3{F0: 0, F1: 889}
        _ = v23
    default:
        v7 -= v7
    }
    v24 := GenS1{F0: 4294967295}
    var v25 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v7
    _ = v8
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint32; F2 uint32 }
func genFn0() uint16 {
    v5 := []int{32768, 189}
    if false {
        for fi6 := 0; fi6 < 19; fi6++ {
            switch ("\n\t" + "") {
            case "foo bar":
                v7 := GenS1{F0: 0}
                _ = v7
            }
            _ = fi6
        }
        for fi8 := 0; fi8 < 17; fi8++ {
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 uint32 = 897
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
        var v11 int64 = 42
        _ = v11
    } else {
        var v12 uint16 = 65535
        var v13 uint64 = 18446744073709551615
        _ = v12
        _ = v13
    }
    _ = v5
    return 32767
}

func main() {
    var v1 uint32 = 0
    var v2 bool = true
    var v3 int32 = 0
    var v4 float64 = ((0.0 * (0.0 + (((float64(3.14159) * -2.25) - ((0.1 / 1.0) * 1e10)) - 0.0))) - -2.25)
    if v2 {
        switch ("" + "\n\t") {
        case "hello":
            v14 := []int{-1, -93, 868}
            _ = v14
        default:
            if (2 > 127) {
                v2 = (!v2)
                v1 = v1
                var v15 int32 = v3
                _ = v15
            }
        }
        if v2 {
            v3 -= v3
            var v16 uint8 = 23
            v2 = ((int32(v16) != v3) && (("" + "") < (("foo bar" + "") + "hello")))
            _ = v16
        }
        if true {
            var v17 int64 = 9223372036854775807
            v17 = 1
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        v4 = v4
        _ = fi18
    }
    v19 := []int{-1, 396, -118, 2147483647}
    v1 = (v1 ^ v1)
    v3 += v3
    var v20 uint16 = 255
    if (uint32(v20) == (v1 | v1)) {
        _ = len(v19)
    } else {
        if true {
            switch "" {
            case "a":
                v21 := GenS2{F0: "hello", F1: 0, F2: 706}
                _ = v21
            case "hello":
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
            default:
                var v22 uint8 = 128
                _ = v22
            }
        }
        v4 = ((v4 * float64((1e10 - -2.25))) * 0.1)
    }
    v20 -= 255
    fmt.Println(v4, v20, v20)
    if len(v19) > 0 {
        v19[0] = 128
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}

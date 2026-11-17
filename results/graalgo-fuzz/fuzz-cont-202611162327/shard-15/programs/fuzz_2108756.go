package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 uint; F2 int8 }
func genFn0(p0 uint8, p1 int32, p2 int32) int16 {
    fmt.Println(p1, p1, p2)
    v7 := GenS1{F0: 18446744073709551615, F1: 1, F2: 42}
    var v8 int8 = -38
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 string = "a"
    v2 := make(map[string]bool)
    var v3 uint16 = 42
    var v4 string = (v1 + "hello")
    v5 := make(map[uint32]string)
    v6 := []int{-2147483648, 32767}
    if true {
        v9 := GenS2{F0: false, F1: 42}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 17; fi10++ {
            if ("a" <= "foo bar") {
                _ = len(v6)
                var v11 string = v4
                var v12 uint16 = (v3 ^ 65535)
                _ = v11
                _ = v12
            } else {
                delete(v5, 65535)
            }
            var v13 uint8 = 203
            var v14 int = 0
            _ = fi10
            _ = v13
            _ = v14
        }
        var v15 int8 = 127
        var v16 int8 = v15
        _ = v15
        _ = v16
    }
    v6 = append(v6, 127)
    fmt.Println(v4)
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 float32 = ((-2.25 + (float32(0.0) - 0.0)) / 1.0)
        _ = fi17
        _ = v18
    }
    v19 := GenS3{F0: 4294967295}
    v20 := GenS2{F0: (false || false), F1: -521}
    v4 = "a"
    v3 -= (v3 - (v3 * (v3 >> ((682) & 63))))
    v4 = "foo bar"
    for fi21 := 0; fi21 < 9; fi21++ {
        if (!(!true)) {
            fmt.Println(v4, v1)
            switch fi21 {
            case 32767:
                var v22 float64 = 0.0
                _ = v22
            case 32768:
                if false {
                    var v23 int64 = -1
                    var v24 int8 = 42
                    for fi25 := 0; fi25 < 20; fi25++ {
                        var v26 bool = (255 <= 127)
                        _ = len(v6)
                        v1 = ((v4 + v1) + ((v1 + v4) + "hello"))
                        _ = fi25
                        _ = v26
                    }
                    _ = v23
                    _ = v24
                }
            default:
                var v27 float64 = float64((float32((float64(0.0) * 0.0)) * (0.1 - 1e10)))
                _ = v27
            }
            var v28 string = ("foo bar" + v1)
            _ = v28
        }
        _ = fi21
    }
    v29 := GenS4{F0: 475, F1: 42, F2: -125}
    v30 := GenS1{F0: 42, F1: 2147483648, F2: -1}
    fmt.Println(v1, v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}

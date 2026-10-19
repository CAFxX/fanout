package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int64; F2 uint }
func genFn0() uint64 {
    if (("foo bar" + "") < (("foo bar" + "\n\t") + "foo bar")) {
        var v7 int64 = -1
        _ = v7
    } else {
        var v8 bool = true
        _ = v8
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        if (((fi9 >> ((65535) & 63)) > fi9) || (!(2 <= 42))) {
            for fi10 := 0; fi10 < 10; fi10++ {
                v11 := GenS1{F0: 42}
                v12 := []int{32767, fi9, fi10}
                _ = fi10
                _ = v11
                _ = v12
            }
            if (float64(-2.25) < (0.1 - float64(float32(0.1)))) {
                var v13 uint32 = 4294967295
                _ = v13
            } else {
                var v14 uint16 = 42
                v14 += v14
                for fi15 := 0; fi15 < 2; fi15++ {
                    v16 := GenS1{F0: v14}
                    _ = fi15
                    _ = v16
                }
                _ = v14
            }
            for fi17 := 0; fi17 < 10; fi17++ {
                if false {
                    v18 := []int{}
                    var v19 string = ("foo bar" + "a")
                    _ = v18
                    _ = v19
                }
                var v20 int32 = 0
                var v21 float32 = -2.25
                var v22 int16 = 0
                v21 = 1e10
                _ = fi17
                _ = v20
                _ = v21
                _ = v22
            }
        } else {
            var v23 float64 = 1e10
            var v24 int64 = -64
            _ = v23
            _ = v24
        }
        _ = fi9
    }
    return 4294967295
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]string)
    v4 := make(map[uint32]int)
    v5 := make(map[int]int)
    v6 := make(map[int]string)
    v25 := []int{32767, -9223372036854775808, -551, 2147483647}
    v26 := GenS2{F0: 108, F1: 432, F2: 65535}
    var v27 int64 = -9223372036854775808
    for fi28 := 0; fi28 < 6; fi28++ {
        if (fi28 != (^(fi28))) {
            v27 -= 464
        }
        v29 := v3["\n\t"]
        switch "foo bar" {
        case "foo bar":
            v27 = 0
        default:
            fmt.Println(v29, fi28)
        }
        _ = fi28
        _ = v29
    }
    var v30 float32 = 0.1
    var v31 bool = false
    fmt.Println(v27, v30, v1)
    v30 -= v30
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}

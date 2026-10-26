package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 uint }
func genFn0(p0 int) int8 {
    var v7 bool = false
    v8 := GenS1{F0: true, F1: 226}
    var v9 int = (p0 % (((p0 / ((0) | 1))) | 1))
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 111
}

func main() {
    var v1 float32 = float32(0.1)
    v2 := make(map[string]int)
    v3 := make(map[string]int)
    var v4 int = -54
    v5 := []int{}
    v6 := []int{v4}
    for fi10 := 0; fi10 < 8; fi10++ {
        if (890 <= 4294967295) {
            var v11 float32 = (-2.25 + v1)
            fmt.Println(v4)
            _ = v11
        } else {
            if ((!((0.1 / -2.25) == (-2.25 + 3.14159))) || (18446744073709551615 > 4294967295)) {
                v12 := GenS2{F0: 388, F1: 42}
                var v13 uint16 = 32767
                _ = v12
                _ = v13
            }
            switch 785 {
            case -18:
                var v14 uint64 = 1
                _ = v14
            }
            if (679 == 288) {
                v15 := GenS3{F0: -1}
                _ = v15
            }
        }
        _ = fi10
    }
    v16 := GenS2{F0: 255, F1: 0}
    for fi17 := 0; fi17 < 19; fi17++ {
        v1 += v1
        _ = fi17
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        v1 += v1
        switch "\n\t" {
        case "\n\t":
            v19 := GenS4{F0: 65535, F1: 60}
            _ = v19
        }
        _ = fi18
    }
    fmt.Println(v4)
    var v20 float32 = v1
    v4 = (v4 * v4)
    var v21 uint8 = 0
    var v22 int16 = -32768
    var v23 uint16 = 65535
    var v24 float64 = (0.1 / 1e10)
    v2[("\n\t" + "a")] = 32768
    v25 := GenS1{F0: false, F1: uint32(v21)}
    v21 = (v21 | v21)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

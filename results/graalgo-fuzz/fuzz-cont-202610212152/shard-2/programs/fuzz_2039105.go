package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 int64 }
func genFn0(p0 int, p1 int8, p2 uint8) uint8 {
    switch ("hello" + "\n\t") {
    case "foo bar":
        if (float32(0.1) > 0.1) {
            var v5 uint8 = (p2 - (^(p2)))
            _ = v5
        } else {
            p0 = p0
        }
    case "a":
        if (!true) {
            var v6 uint8 = 224
            _ = v6
        }
    }
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    return (p2 | 127)
}

func main() {
    var v1 int64 = 829
    v2 := make(map[uint32]int)
    v3 := []int{0, 42}
    var v4 string = "hello"
    fmt.Println(v4, v4)
    var v7 bool = true
    for fi8 := 0; fi8 < 20; fi8++ {
        fmt.Println(v1)
        _ = fi8
    }
    var v9 uint64 = 392
    var v10 uint8 = genFn0(-829, 1, 2)
    var v11 int = 9223372036854775807
    fmt.Println(v7, v10, v10)
    if (!(v1 <= v1)) {
        for fi12 := 0; fi12 < 14; fi12++ {
            for fi13 := 0; fi13 < 18; fi13++ {
                var v14 int = (^(v11))
                if len(v3) > 0 {
                    v3[0] = v11
                }
                _ = fi13
                _ = v14
            }
            var v15 int = (fi12 / ((42) | 1))
            fmt.Println(v11)
            _ = fi12
            _ = v15
        }
        v3 = append(v3, (v11 / ((-528) | 1)))
        if (false || (true && (-32768 != int16(v10)))) {
            var v16 int32 = -949
            var v17 bool = (32767 != 65535)
            for fi18 := 0; fi18 < 5; fi18++ {
                v9 -= v9
                _ = fi18
            }
            _ = v16
            _ = v17
        } else {
            if v7 {
                var v19 string = v4
                _ = v19
            } else {
                v20 := GenS1{F0: v9, F1: "\n\t"}
                v21 := GenS1{F0: (v9 & 0), F1: "a"}
                _ = v20
                _ = v21
            }
            v4 = v4
            _ = len(v3)
        }
    }
    v11 -= (v11 - 2147483647)
    v22 := GenS2{F0: 42, F1: 107}
    delete(v2, 0)
    var v23 int8 = 100
    v24 := GenS3{F0: v1}
    fmt.Println(v23)
    var v25 uint64 = 42
    var v26 int = v11
    fmt.Println(v10)
    v11 = (-(v26))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}

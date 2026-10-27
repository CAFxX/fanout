package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int32 }
func genFn0(p0 int16, p1 int32) int16 {
    v5 := GenS1{F0: 42, F1: 756, F2: (-(p1))}
    v6 := GenS1{F0: -1, F1: 65535, F2: p1}
    var v7 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return p0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "hello"
    var v3 int16 = 32767
    v4 := make(map[string]uint32)
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 uint8 = 42
        v2 = ""
        _ = fi8
        _ = v9
    }
    delete(v4, v2)
    switch 127 {
    case -596:
        v2 = (((v2 + v2) + "\n\t") + (("hello" + v2) + v2))
    default:
        if (!(false || true)) {
            var v10 int16 = 1
            if (true && true) {
                var v11 uint16 = 278
                for fi12 := 0; fi12 < 15; fi12++ {
                    for fi13 := 0; fi13 < 11; fi13++ {
                        fmt.Println(fi13, v3, v3)
                        _ = fi13
                    }
                    switch -2147483648 {
                    case 42:
                        v11 = v11
                    case 541:
                        if (0.0 < -2.25) {
                            fmt.Println(fi12, v10, fi12)
                        }
                    case -9223372036854775808:
                        var v14 bool = true
                        _ = v14
                    }
                    v15 := GenS1{F0: 1, F1: 0, F2: -2147483648}
                    _ = fi12
                    _ = v15
                }
                _ = v11
            } else {
                v2 = "a"
            }
            _ = v10
        }
    }
    v16 := GenS1{F0: -9223372036854775808, F1: 4294967295, F2: 0}
    v3 -= v3
    for fi17 := 0; fi17 < 8; fi17++ {
        var v18 int16 = 441
        v19 := GenS1{F0: -878, F1: 47, F2: 2147483647}
        fmt.Println(v2, v2)
        _ = fi17
        _ = v18
        _ = v19
    }
    switch v2 {
    case "":
        var v20 uint32 = 246
        _ = v20
    case "a":
        fmt.Println(v3, v3)
    case "foo bar":
        v21 := GenS1{F0: 0, F1: 2, F2: -1}
        _ = v21
    }
    v3 -= v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}

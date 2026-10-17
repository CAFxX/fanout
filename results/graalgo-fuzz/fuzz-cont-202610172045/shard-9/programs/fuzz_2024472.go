package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 string; F1 int; F2 uint16 }
type GenS3 struct { F0 bool; F1 uint16; F2 int16 }

func main() {
    v1 := make(map[int]int)
    var v2 int64 = 42
    var v3 int32 = -2147483648
    v4 := []string{("a" + "hello")}
    v5 := make(map[uint32]uint32)
    var v6 int16 = -1
    var v7 float64 = -2.25
    var v8 uint32 = 4294967295
    if (-393 <= v6) {
        for fi9 := 0; fi9 < 6; fi9++ {
            if (127 != 1) {
                _ = len(v5)
            }
            _ = fi9
        }
    }
    switch "foo bar" {
    case "a":
        var v10 uint64 = 0
        _ = v10
    case "hello":
        v11 := GenS1{F0: 1, F1: v8}
        _ = v11
    default:
        for fi12 := 0; fi12 < 12; fi12++ {
            v13 := GenS1{F0: 0, F1: v8}
            _ = fi12
            _ = v13
        }
    }
    var v14 uint64 = 18446744073709551615
    v6 += (v6 * v6)
    v15 := GenS2{F0: "", F1: -110, F2: 255}
    var v16 int8 = 127
    v17 := GenS2{F0: "a", F1: 32767, F2: 879}
    v16 += v16
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 int = 42
        var v20 float64 = ((v7 / 3.14159) - 3.14159)
        for fi21 := 0; fi21 < 14; fi21++ {
            var v22 float64 = (v7 - (v7 + v20))
            var v23 uint64 = (v14 | v14)
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    switch -343 {
    case -419:
        if false {
            if (!((3.14159 - 1.5) <= 1e10)) {
                v24 := GenS3{F0: (!(false || (true && (1 != v8)))), F1: 537, F2: int16(v14)}
                for fi25 := 0; fi25 < 3; fi25++ {
                    for fi26 := 0; fi26 < 3; fi26++ {
                        var v27 uint = 42
                        _ = fi26
                        _ = v27
                    }
                    v2 = -1
                    _ = fi25
                }
                _ = v24
            }
        }
    case 128:
        v1[int(v2)] = -465
    }
    var v28 int8 = v16
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

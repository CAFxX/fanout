package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int8 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint32 }

func main() {
    var v1 int8 = 0
    var v2 uint32 = 1
    v3 := make(map[int]int)
    var v4 float32 = (1e10 * 3.14159)
    v5 := make(map[int]string)
    if ((-1 >= 100) && true) {
        v2 += v2
        var v6 int16 = 32767
        switch 2147483647 {
        case -1:
            if ((v6 & v6) < v6) {
                v7 := GenS1{F0: 1, F1: 0, F2: -1}
                _ = v7
            }
        case 128:
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 string = ("a" + ("hello" + ("\n\t" + ("hello" + ""))))
                fmt.Println(v6)
                v6 += v6
                _ = fi8
                _ = v9
            }
        }
        _ = v6
    } else {
        v10 := GenS1{F0: 0, F1: 0, F2: 42}
        v2 -= 1
        var v11 int = 42
        _ = v10
        _ = v11
    }
    fmt.Println(v4)
    fmt.Println(v4, v4, v1)
    delete(v5, -932)
    for fi12 := 0; fi12 < 4; fi12++ {
        switch "\n\t" {
        case "\n\t":
            if false {
                var v13 int8 = 100
                fmt.Println(v4)
                delete(v5, fi12)
                _ = v13
            }
        case "a":
            var v14 int = (fi12 >> ((4294967295) & 63))
            _ = v14
        default:
            v15 := GenS1{F0: 2, F1: (^(fi12)), F2: int8(v2)}
            _ = v15
        }
        v16 := GenS2{F0: 2147483647, F1: (v1 | 111), F2: v2}
        _ = fi12
        _ = v16
    }
    delete(v5, 1)
    for fi17 := 0; fi17 < 15; fi17++ {
        for fi18 := 0; fi18 < 6; fi18++ {
            var v19 int16 = 42
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    fmt.Println(v1, v1, v2)
    if ((0.0 - (0.1 + 1.5)) <= 1.5) {
        fmt.Println(v4)
        var v20 uint = 1
        for fi21 := 0; fi21 < 9; fi21++ {
            var v22 int = 42
            var v23 int64 = int64(fi21)
            v24 := []int{}
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        _ = v20
    } else {
        for fi25 := 0; fi25 < 16; fi25++ {
            fmt.Println(v4)
            for fi26 := 0; fi26 < 3; fi26++ {
                var v27 uint64 = 433
                var v28 int16 = 42
                var v29 string = "hello"
                _ = fi26
                _ = v27
                _ = v28
                _ = v29
            }
            v2 = (^(v2))
            _ = fi25
        }
        v30 := GenS2{F0: (v2 | v2), F1: -117, F2: (v2 - v2)}
        _ = v30
    }
    v31 := []int{128, 0, 42}
    var v32 bool = (true || false)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int16 }
type GenS2 struct { F0 uint64; F1 int; F2 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint, p1 string) uint64 {
    var v4 float64 = (0.1 / -2.25)
    var v5 int64 = -1
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 692
}
func genFn1() string {
    var v6 uint16 = 42
    v7 := GenS1{F0: 128, F1: -242, F2: -32768}
    _ = v6
    _ = v7
    return ("a" + "")
}

func main() {
    var v1 int8 = 0
    v2 := make(map[string]int)
    var v3 int = int(v1)
    v3 = v3
    v8 := GenS1{F0: 128, F1: (v3 >> ((2) & 63)), F2: 1}
    var v9 int8 = 100
    if true {
        v1 = v1
        v10 := GenS1{F0: (v3 / (((v3 * v3)) | 1)), F1: (v3 % (((v3 / ((-443) | 1))) | 1)), F2: 42}
        _ = v10
    }
    v11 := GenS2{F0: 687, F1: (-(v3)), F2: 40}
    if (133 > 18446744073709551615) {
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 string = "hello"
            _ = fi12
            _ = v13
        }
    } else {
        v14 := GenS3{F0: 1000}
        _ = v14
    }
    if false {
        var v15 bool = false
        v2["foo bar"] = (^(v3))
        _ = len(v2)
        _ = v15
    }
    var v16 uint = 4294967295
    if (!true) {
        var v17 float32 = (((3.14159 / (((1.5 * 0.1) + 0.0) / 1.0)) - ((((-2.25 - (3.14159 - 1e10)) - 3.14159) * (0.0 + 0.1)) + 1e10)) - 0.0)
        _ = v17
    }
    v18 := GenS2{F0: 18446744073709551615, F1: (v3 % ((v3) | 1)), F2: 32767}
    if (96 == v1) {
        for fi19 := 0; fi19 < 17; fi19++ {
            fmt.Println(v1, v9, v16)
            v20 := GenS1{F0: fi19, F1: 128, F2: 1}
            v3 -= fi19
            _ = fi19
            _ = v20
        }
    }
    v21 := GenS4{F0: true}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v16
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v8
    _ = v9
}

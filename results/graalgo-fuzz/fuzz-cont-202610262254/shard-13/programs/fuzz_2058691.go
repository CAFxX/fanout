package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint64 }
type GenS2 struct { F0 int32; F1 uint; F2 uint32 }
type GenS3 struct { F0 int8; F1 bool; F2 string }
type GenS4 struct { F0 uint64; F1 uint16; F2 int8 }
func genFn0(p0 int16) uint64 {
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 float32 = (0.1 * float32((3.14159 * float64((((-2.25 - 1e10) + 3.14159) * (3.14159 / 1e10))))))
        _ = fi7
        _ = v8
    }
    v9 := []int{9223372036854775807, 404, 2147483647}
    var v10 uint8 = 2
    _ = p0
    _ = v10
    _ = v9
    return 42
}

func main() {
    v1 := make(map[int]string)
    var v2 string = ""
    var v3 uint64 = 42
    v4 := make(map[uint32]string)
    var v5 bool = true
    v6 := make(map[string]bool)
    fmt.Println(v3, v5)
    v2 = "\n\t"
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := GenS1{F0: 42, F1: 1, F2: v3}
        v13 := []int{(fi11 | fi11), 1, -855, fi11}
        _ = fi11
        _ = v12
        _ = v13
    }
    switch 853 {
    case 0:
        if true {
            fmt.Println(v2, v2, v5)
            v14 := GenS2{F0: 664, F1: 42, F2: 65535}
            var v15 string = v2
            _ = v14
            _ = v15
        } else {
            v3 -= (v3 + v3)
        }
    case 32768:
        for fi16 := 0; fi16 < 14; fi16++ {
            v2 = (v2 + ((("hello" + v2) + (v2 + v2)) + v2))
            _ = fi16
        }
    case -2147483648:
        var v17 uint8 = 42
        _ = v17
    }
    fmt.Println(v3, v2)
    var v18 uint64 = (v3 | v3)
    v19 := GenS3{F0: 127, F1: v5, F2: v2}
    v18 = (v3 >> ((936) & 63))
    v5 = (127 <= 255)
    fmt.Println(v3)
    for fi20 := 0; fi20 < 14; fi20++ {
        for fi21 := 0; fi21 < 9; fi21++ {
            delete(v6, (("foo bar" + "a") + v2))
            fmt.Println(fi21)
            _ = fi21
        }
        _ = fi20
    }
    if v5 {
        fmt.Println(v5)
    } else {
        var v22 int64 = -981
        _ = v22
    }
    v18 = (v3 & v18)
    if (2147483648 == 750) {
        v23 := GenS4{F0: (v18 >> ((980) & 63)), F1: 164, F2: -128}
        v24 := GenS1{F0: 255, F1: 127, F2: genFn0(458)}
        var v25 bool = (0 == -41)
        _ = v23
        _ = v24
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

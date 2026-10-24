package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 uint32; F2 int64 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]int)
    var v3 uint32 = 857
    v4 := make(map[uint32]string)
    var v5 int = 127
    v6 := GenS1{F0: 0, F1: ("hello" + "a")}
    var v7 bool = true
    delete(v2, v5)
    var v8 int8 = 0
    v7 = (413 <= 733)
    for fi9 := 0; fi9 < 19; fi9++ {
        if (!true) {
            var v10 uint8 = 1
            _ = v10
        } else {
            var v11 uint = 274
            _ = v11
        }
        v12 := GenS2{F0: (("\n\t" + "hello") + ("a" + "\n\t")), F1: v3}
        v13 := GenS3{F0: -268}
        _ = fi9
        _ = v12
        _ = v13
    }
    var v14 uint16 = 239
    if v7 {
        fmt.Println(v5, v8)
        if (42 != 18446744073709551615) {
            v3 -= v3
        } else {
            if v7 {
                v4[v3] = (("" + ("foo bar" + (("" + ("" + "")) + "hello"))) + "a")
                v15 := GenS1{F0: v8, F1: "hello"}
                _ = v15
            } else {
                var v16 float64 = (3.14159 - 0.1)
                v5 -= 519
                _ = v16
            }
            v17 := GenS4{F0: 0, F1: v3, F2: -467}
            _ = v17
        }
    }
    v8 = (v8 & v8)
    var v18 bool = true
    switch v5 {
    case 2147483647:
        switch ("hello" + "") {
        case "foo bar":
            var v19 float64 = float64((0.0 + (float32(((-2.25 * 3.14159) * float64(3.14159))) - (-2.25 * 3.14159))))
            _ = v19
        case "hello":
            _ = len(v1)
        case "a":
            v14 = (^(v14))
        }
    default:
        v20 := []int{(v5 >> ((4294967295) & 63)), (^(v5)), v5}
        _ = v20
    }
    switch -9223372036854775808 {
    case 703:
        var v21 float64 = 1e10
        _ = v21
    default:
        if v18 {
            fmt.Println(v5, v8, v3)
        } else {
            v18 = v18
            v18 = (32767 >= 794)
            var v22 uint64 = 12
            _ = v22
        }
    }
    if v18 {
        v7 = (v18 && (952 > -928))
        delete(v2, v5)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

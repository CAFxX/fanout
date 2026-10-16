package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint16 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 bool }
func genFn0() int64 {
    var v5 uint32 = 661
    v6 := GenS1{F0: 873, F1: ("hello" + "\n\t"), F2: 988}
    _ = v5
    _ = v6
    return 9223372036854775807
}
func genFn1() uint64 {
    switch 32767 {
    case -1:
        var v7 bool = (2147483648 < 2147483648)
        _ = v7
    case 4294967295:
        switch 127 {
        case 128:
            var v8 string = "\n\t"
            _ = v8
        case 32768:
            var v9 int32 = 1
            _ = v9
        }
    default:
        var v10 float64 = 0.0
        _ = v10
    }
    var v11 uint = 18446744073709551615
    fmt.Println(v11)
    for fi12 := 0; fi12 < 12; fi12++ {
        v11 = (v11 | 18446744073709551615)
        _ = fi12
    }
    _ = v11
    return 2
}

func main() {
    var v1 string = ""
    v2 := make(map[uint32]string)
    v3 := []string{v1, v1, "hello"}
    var v4 int64 = 1
    if true {
        if len(v3) > 0 {
            v3[0] = v1
        }
    } else {
        if false {
            v13 := GenS2{F0: -1, F1: -223}
            v1 = (v1 + v1)
            _ = v13
        }
        fmt.Println(v1, v4)
    }
    var v14 uint = 391
    v15 := GenS1{F0: 1, F1: (("foo bar" + "") + (v1 + "hello")), F2: 258}
    v16 := GenS3{F0: false}
    if (!true) {
        var v17 uint32 = 2147483647
        _ = v17
    }
    v4 = (^(v4))
    if (true && (v14 <= (v14 + v14))) {
        var v18 float64 = ((0.0 * -2.25) * ((-2.25 * -2.25) / -2.25))
        _ = v18
    }
    v1 = v1
    v14 -= (v14 | (v14 | (v14 + (v14 ^ (v14 & (v14 << (((^(v14))) & 63)))))))
    v19 := GenS2{F0: 2147483647, F1: 756}
    v14 += 4294967295
    switch "\n\t" {
    case "foo bar":
        v14 = v14
    default:
        fmt.Println(v4, v1, v14)
    }
    fmt.Println(v4, v14, v14)
    v14 = (v14 | 510)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}

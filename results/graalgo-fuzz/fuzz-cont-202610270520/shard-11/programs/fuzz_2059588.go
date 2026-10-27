package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 int64; F2 int64 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 uint8; F1 bool }
func genFn0(p0 uint8) string {
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := GenS1{F0: 369, F1: 413}
        var v8 bool = (-41 < 42)
        fmt.Println(fi6)
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(p0)
    v9 := GenS1{F0: 1, F1: 32767}
    _ = p0
    _ = v9
    return "\n\t"
}

func main() {
    var v1 uint64 = 1
    var v2 bool = false
    var v3 int64 = -276
    v4 := make(map[int]string)
    var v5 int = -878
    switch v5 {
    case 2147483647:
        var v10 uint16 = 255
        _ = v10
    case 843:
        v11 := []int{v5, (v5 - v5), v5}
        _ = v11
    default:
        for fi12 := 0; fi12 < 9; fi12++ {
            v13 := []int{v5}
            v5 = (v5 / (((fi12 * fi12)) | 1))
            _ = fi12
            _ = v13
        }
    }
    var v14 uint = 138
    for fi15 := 0; fi15 < 14; fi15++ {
        v16 := GenS2{F0: (v3 << ((v14) & 63))}
        v17 := GenS3{F0: 4294967295, F1: 0, F2: v3}
        _ = fi15
        _ = v16
        _ = v17
    }
    switch ("" + ("" + "a")) {
    case "foo bar":
        v2 = v2
    case "hello":
        var v18 uint32 = 1
        _ = v18
    default:
        switch (v5 | (v5 >> ((255) & 63))) {
        case 2147483647:
            for fi19 := 0; fi19 < 15; fi19++ {
                fmt.Println(v1, fi19)
                fmt.Println(v3, v5, v5)
                for fi20 := 0; fi20 < 2; fi20++ {
                    v3 = v3
                    var v21 float32 = (3.14159 * 3.14159)
                    var v22 uint8 = 127
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                _ = fi19
            }
        case 42:
            var v23 string = (("a" + ("\n\t" + ((("hello" + ("foo bar" + ("hello" + "hello"))) + "hello") + "\n\t"))) + "a")
            _ = v23
        case 4294967295:
            for fi24 := 0; fi24 < 19; fi24++ {
                var v25 uint64 = v1
                var v26 float64 = 0.1
                v27 := GenS3{F0: 2147483648, F1: int64(v1), F2: (^(v3))}
                _ = fi24
                _ = v25
                _ = v26
                _ = v27
            }
        default:
            if v2 {
                v28 := GenS4{F0: (723 >= uint16(v14)), F1: -9223372036854775808}
                v2 = (!v2)
                _ = v28
            }
        }
    }
    v3 = (v3 & (v3 * v3))
    if v2 {
        fmt.Println(v14, v3, v3)
        switch (v5 + v5) {
        case -9223372036854775808:
            v3 += -1
        default:
            v5 += v5
        }
    } else {
        v29 := []int{v5, (v5 + v5), 127, v5}
        v30 := GenS1{F0: 0, F1: 32767}
        _ = v29
        _ = v30
    }
    v31 := GenS5{F0: 161, F1: false}
    v5 += -924
    var v32 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}

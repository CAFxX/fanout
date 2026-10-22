package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 int16; F2 int64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]string)
    v3 := make(map[string]bool)
    var v4 bool = true
    var v5 int = 9223372036854775807
    switch "\n\t" {
    case "\n\t":
        if false {
            v6 := []int{(v5 % (((v5 & v5)) | 1))}
            _ = v6
        } else {
            fmt.Println(v4)
            fmt.Println(v5)
        }
    case "a":
        v7 := v2[2147483648]
        _ = v7
    case "foo bar":
        v8 := GenS1{F0: 9223372036854775807, F1: -1, F2: 0}
        _ = v8
    default:
        v9 := GenS2{F0: 42}
        _ = v9
    }
    v5 -= v5
    if (v5 >= v5) {
        var v10 uint16 = 551
        _ = v10
    }
    if false {
        var v11 float32 = (0.1 - (3.14159 - 0.0))
        v5 += v5
        v12 := []int{(v5 ^ v5), v5}
        _ = v11
        _ = v12
    } else {
        if (1 >= 222) {
            delete(v1, v5)
            v4 = (!v4)
            for fi13 := 0; fi13 < 10; fi13++ {
                v5 = fi13
                var v14 bool = v4
                v5 -= v5
                _ = fi13
                _ = v14
            }
        } else {
            v15 := GenS2{F0: 0}
            _ = v15
        }
        v5 = -764
    }
    fmt.Println(v4)
    if (2147483648 == 2147483647) {
        var v16 uint64 = 18446744073709551615
        _ = v16
    }
    switch v5 {
    case 42:
        v17 := GenS2{F0: 0}
        _ = v17
    case -2147483648:
        var v18 int32 = -2147483648
        _ = v18
    default:
        v5 -= (v5 << ((4294967295) & 63))
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        fmt.Println(v5, fi19)
        v20 := GenS3{F0: 62, F1: 0, F2: int64(v5)}
        _ = fi19
        _ = v20
    }
    v21 := v2[1]
    for fi22 := 0; fi22 < 6; fi22++ {
        if ((float32((3.14159 + 0.1)) / (float32(1.5) + 1e10)) <= -2.25) {
            v23 := GenS3{F0: 4294967295, F1: 512, F2: 42}
            fmt.Println(v5, v21)
            _ = v23
        }
        for fi24 := 0; fi24 < 7; fi24++ {
            v21 = "\n\t"
            if v4 {
                fmt.Println(fi22, fi22, v4)
                var v25 uint8 = 1
                v5 = (v5 | 2147483647)
                _ = v25
            } else {
                fmt.Println(v4, fi24)
                var v26 float64 = float64(-2.25)
                _ = v26
            }
            v27 := v3[((v21 + "\n\t") + (v21 + ((v21 + (("hello" + ("" + (v21 + ("hello" + "\n\t")))) + (v21 + ("a" + v21)))) + v21)))]
            _ = fi24
            _ = v27
        }
        _ = fi22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}

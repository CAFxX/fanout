package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 int64; F2 string }
type GenS4 struct { F0 int32; F1 uint; F2 int16 }
func genFn0(p0 int16) uint8 {
    v5 := GenS1{F0: 9223372036854775807, F1: 127}
    v6 := GenS1{F0: 741, F1: 1}
    v7 := []int{128, -333, 127, 32768}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 255
}
func genFn1(p0 uint) int64 {
    v8 := GenS1{F0: 0, F1: 42}
    var v9 uint8 = 1
    switch ("a" + ("\n\t" + ("hello" + ""))) {
    case "":
        v9 -= 233
    case "\n\t":
        p0 = p0
    }
    if false {
        if (false || true) {
            switch -388 {
            case -764:
                v9 = genFn0(1)
            case 2147483647:
                p0 += 2
            }
        } else {
            v10 := []int{2147483647}
            p0 = (p0 + p0)
            fmt.Println(p0, p0)
            _ = v10
        }
        p0 -= (p0 << (((^(p0))) & 63))
    }
    fmt.Println(p0, v9)
    _ = p0
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[int]string)
    var v3 float32 = (float32(1e10) - float32(0.0))
    v4 := []uint32{2, 2147483648}
    var v11 int64 = 1
    v12 := GenS2{F0: 788}
    var v13 uint32 = 233
    for fi14 := 0; fi14 < 15; fi14++ {
        v11 = 9223372036854775807
        v11 += v11
        if true {
            v15 := GenS3{F0: 42, F1: (v11 & v11), F2: ""}
            var v16 int = (fi14 * fi14)
            _ = v15
            _ = v16
        } else {
            v17 := GenS4{F0: 504, F1: 42, F2: 32767}
            var v18 float32 = (0.0 - 3.14159)
            _ = v17
            _ = v18
        }
        _ = fi14
    }
    if (576 != 0) {
        for fi19 := 0; fi19 < 11; fi19++ {
            fmt.Println(v13)
            var v20 int16 = 42
            _ = fi19
            _ = v20
        }
    } else {
        var v21 int32 = 0
        _ = v21
    }
    var v22 uint64 = 1
    fmt.Println(v11, v11, v1)
    v3 -= (float32((3.14159 * 0.1)) / ((-2.25 * float32(((float64(-2.25) / (1.5 - 1e10)) - 1.5))) - v3))
    fmt.Println(v13, v1)
    fmt.Println(v13, v11)
    v3 = ((0.0 / v3) / (v3 * (v3 / (0.1 - v3))))
    if (1 == 65535) {
        var v23 uint = uint(v1)
        var v24 int = 127
        var v25 string = "a"
        _ = v23
        _ = v24
        _ = v25
    } else {
        var v26 float32 = v3
        _ = v26
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}

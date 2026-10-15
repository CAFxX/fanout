package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 bool }
func genFn0() int32 {
    switch "\n\t" {
    case "":
        for fi4 := 0; fi4 < 4; fi4++ {
            var v5 int64 = -836
            var v6 float32 = 1.5
            v6 += float32(0.1)
            _ = fi4
            _ = v5
            _ = v6
        }
    case "hello":
        var v7 uint = 4294967295
        _ = v7
    }
    switch -535 {
    case -9223372036854775808:
        var v8 float32 = (1e10 - (((0.0 - 3.14159) * ((1e10 + float32(0.0)) * (-2.25 * (((((0.1 - 3.14159) * 3.14159) + 0.0) - 3.14159) - 1.5)))) * (3.14159 - (1.5 - (1e10 * 1e10)))))
        _ = v8
    default:
        v9 := GenS1{F0: 391, F1: 81, F2: 654}
        _ = v9
    }
    var v10 string = ""
    var v11 int = 32768
    _ = v10
    _ = v11
    return -2147483648
}
func genFn1(p0 int64, p1 uint16, p2 int) int16 {
    for fi12 := 0; fi12 < 6; fi12++ {
        p2 = p2
        for fi13 := 0; fi13 < 16; fi13++ {
            fmt.Println(fi13, fi13)
            p2 -= (p2 ^ 188)
            _ = fi13
        }
        _ = fi12
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        p1 -= (p1 | (^(p1)))
        var v15 int64 = -9223372036854775808
        _ = fi14
        _ = v15
    }
    v16 := []int{128, p2, (p2 + 4294967295), (p2 | p2)}
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return -242
}

func main() {
    var v1 int8 = 0
    var v2 int64 = 964
    v3 := make(map[string]bool)
    var v17 uint8 = 127
    var v18 int = 32767
    fmt.Println(v18, v2)
    switch 9223372036854775807 {
    case -14:
        var v19 uint64 = 18446744073709551615
        _ = v19
    case 127:
        for fi20 := 0; fi20 < 3; fi20++ {
            if true {
                v18 += fi20
            } else {
                for fi21 := 0; fi21 < 20; fi21++ {
                    v2 = v2
                    _ = fi21
                }
                var v22 int64 = v2
                _ = v22
            }
            v23 := GenS2{F0: v17}
            _ = fi20
            _ = v23
        }
    default:
        if (false && true) {
            switch (v18 - (v18 - 519)) {
            case -9223372036854775808:
                _ = len(v3)
            default:
                var v24 bool = false
                _ = v24
            }
            switch (((("hello" + "hello") + "hello") + "a") + ("hello" + "hello")) {
            case "":
                var v25 int16 = 1
                _ = v25
            case "hello":
                var v26 uint8 = (v17 | 42)
                _ = v26
            case "\n\t":
                var v27 bool = (!(852 >= 42))
                _ = v27
            default:
                fmt.Println(v2, v17)
            }
            fmt.Println(v2, v1, v17)
        } else {
            var v28 uint32 = 2147483648
            var v29 uint32 = (v28 << ((0) & 63))
            v17 = (^(v17))
            _ = v28
            _ = v29
        }
    }
    fmt.Println(v17)
    if true {
        var v30 string = "\n\t"
        _ = v30
    }
    switch (v18 + (-(v18))) {
    case 911:
        var v31 string = (("" + ("\n\t" + "\n\t")) + ("hello" + "\n\t"))
        _ = v31
    default:
        _ = len(v3)
    }
    v32 := []int{2147483647, -748, 2147483647}
    delete(v3, ("\n\t" + "hello"))
    v33 := GenS3{F0: 0, F1: (float64((-2.25 * ((1e10 + float32((0.0 * 0.0))) - 1.5))) >= 3.14159)}
    v17 = 128
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v32:", len(v32), v32)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v33
}

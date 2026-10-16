package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint8, p1 string) uint8 {
    p1 = (p1 + "\n\t")
    v7 := GenS1{F0: 0, F1: false}
    _ = p0
    _ = p1
    _ = v7
    return (^(p0))
}
func genFn1(p0 bool, p1 uint64) int64 {
    fmt.Println(p1, p0)
    if (65535 < 2147483647) {
        for fi8 := 0; fi8 < 18; fi8++ {
            v9 := GenS2{F0: 1, F1: 32767}
            var v10 bool = p0
            v11 := GenS2{F0: 127, F1: 32767}
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
    } else {
        v12 := GenS1{F0: 42, F1: true}
        for fi13 := 0; fi13 < 20; fi13++ {
            v14 := []int{(fi13 + (fi13 + 4294967295))}
            _ = fi13
            _ = v14
        }
        _ = v12
    }
    v15 := GenS3{F0: 415}
    _ = p0
    _ = p1
    _ = v15
    return -9223372036854775808
}

func main() {
    var v1 int8 = 100
    var v2 int = -585
    v3 := make(map[int]string)
    v4 := make(map[string]uint32)
    v5 := make(map[string]bool)
    var v6 string = "\n\t"
    if true {
        fmt.Println(v2, v2, v2)
    } else {
        var v16 string = "a"
        for fi17 := 0; fi17 < 15; fi17++ {
            v16 = "\n\t"
            v18 := []int{0, -9223372036854775808, (^(fi17)), fi17}
            v16 = v16
            _ = fi17
            _ = v18
        }
        if false {
            v19 := GenS1{F0: 267, F1: true}
            switch v2 {
            case 4294967295:
                fmt.Println(v2, v16, v2)
            default:
                v20 := []int{(v2 >> ((4294967295) & 63))}
                _ = v20
            }
            v21 := GenS2{F0: (v1 ^ v1), F1: 740}
            _ = v19
            _ = v21
        }
        _ = v16
    }
    fmt.Println(v6, v2, v1)
    v22 := GenS2{F0: (v1 / (((v1 + -128)) | 1)), F1: 583}
    _ = len(v3)
    fmt.Println(v1, v6)
    v23 := []int{}
    switch (("hello" + "") + v6) {
    case "hello":
        fmt.Println(v6, v1)
    case "a":
        v1 = (v1 >> ((896) & 63))
    case "\n\t":
        delete(v5, (v6 + v6))
    }
    if len(v23) > 0 {
        v23[0] = (v2 & (v2 >> ((42) & 63)))
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        v25 := GenS2{F0: (v1 | v1), F1: 687}
        v26 := GenS3{F0: 748}
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 int8 = v1
    v28 := GenS2{F0: (v1 | v27), F1: 255}
    var v29 uint32 = 2147483647
    v27 = 28
    v1 = -128
    fmt.Println(v27, v1, v2)
    for fi30 := 0; fi30 < 20; fi30++ {
        for fi31 := 0; fi31 < 8; fi31++ {
            switch 2147483647 {
            case 127:
                v27 = -1
                var v32 int = int(v29)
                _ = v32
            case -313:
                var v33 int64 = 0
                _ = v33
            }
            if len(v23) > 0 {
                v23[0] = fi31
            }
            var v34 float32 = ((3.14159 - (0.1 - float32(((3.14159 - (((float64((3.14159 - 0.0)) + 3.14159) - (float64(1e10) / ((3.14159 / -2.25) * 1.5))) * 1.5)) + -2.25)))) * ((1e10 + ((-2.25 * (3.14159 - 1e10)) - 0.1)) - ((0.0 / 0.1) / ((0.1 * 1e10) + (1e10 + (0.0 - 0.1))))))
            _ = fi31
            _ = v34
        }
        _ = fi30
    }
    v2 -= (v2 + (v2 | (v2 % ((v2) | 1))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

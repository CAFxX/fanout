package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int8; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int16, p1 bool, p2 uint) uint16 {
    switch 32768 {
    case 68:
        p1 = false
    case 1:
        v7 := GenS1{F0: 625, F1: 723}
        _ = v7
    case -2147483648:
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := GenS1{F0: 2147483647, F1: 2147483647}
            for fi10 := 0; fi10 < 9; fi10++ {
                v11 := GenS2{F0: 0, F1: 2147483647, F2: 2}
                v12 := []int{fi8}
                _ = fi10
                _ = v11
                _ = v12
            }
            v13 := []int{}
            _ = fi8
            _ = v13
            _ = v9
        }
    default:
        var v14 bool = p1
        _ = v14
    }
    v15 := []int{2147483647, 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return uint16(p2)
}

func main() {
    v1 := []uint32{2147483647, 0, 498}
    v2 := make(map[string]uint32)
    v3 := []uint32{2, 65535}
    var v4 int = 42
    var v5 int8 = -128
    v6 := make(map[uint32]uint32)
    if true {
        var v16 uint = 0
        _ = v16
    }
    if (0 >= 127) {
        v4 = (v4 / ((v4) | 1))
    } else {
        switch "\n\t" {
        case "":
            var v17 string = (("\n\t" + ("a" + "hello")) + "\n\t")
            _ = v17
        }
    }
    if ((float64(3.14159) >= 3.14159) || (genFn0(1, ((1e10 + 0.1) != (3.14159 / 0.1)), 150) <= genFn0(32767, true, 255))) {
        switch (v4 * v4) {
        case -9223372036854775808:
            fmt.Println(v5)
        case 1:
            v5 = v5
        case 128:
            v18 := v6[0]
            _ = v18
        }
        for fi19 := 0; fi19 < 7; fi19++ {
            if len(v3) > 0 {
                v3[0] = 0
            }
            _ = fi19
        }
    }
    var v20 float64 = 1e10
    v21 := GenS1{F0: 2147483648, F1: 15}
    if true {
        var v22 uint16 = 42
        v5 -= v5
        _ = v22
    } else {
        v5 = int8(v4)
        var v23 string = "a"
        _ = v23
    }
    if true {
        var v24 uint32 = 65535
        for fi25 := 0; fi25 < 6; fi25++ {
            for fi26 := 0; fi26 < 6; fi26++ {
                v27 := GenS2{F0: (v5 ^ (v5 >> ((65535) & 63))), F1: -2147483648, F2: 255}
                v28 := GenS2{F0: v5, F1: -2147483648, F2: 127}
                _ = fi26
                _ = v27
                _ = v28
            }
            _ = fi25
        }
        var v29 string = (("" + "") + ("" + "hello"))
        _ = v24
        _ = v29
    } else {
        v20 -= float64(float32(v20))
        v30 := GenS3{F0: 42}
        _ = v30
    }
    if true {
        var v31 int16 = 1000
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v31
    } else {
        v4 = v4
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

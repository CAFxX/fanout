package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 uint; F2 uint }
func genFn0() uint16 {
    var v6 string = "\n\t"
    v7 := GenS1{F0: 65535, F1: 500, F2: 0}
    switch ((v6 + (("foo bar" + "foo bar") + v6)) + "a") {
    case "\n\t":
        for fi8 := 0; fi8 < 13; fi8++ {
            fmt.Println(fi8, fi8)
            for fi9 := 0; fi9 < 5; fi9++ {
                var v10 string = "hello"
                var v11 uint16 = 714
                v11 = v11
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi8
        }
    case "hello":
        var v12 uint8 = 176
        _ = v12
    case "a":
        fmt.Println(v6)
    }
    _ = v6
    _ = v7
    return 32767
}
func genFn1(p0 int8) uint16 {
    for fi13 := 0; fi13 < 13; fi13++ {
        fmt.Println(p0, fi13)
        p0 = (p0 << ((4294967295) & 63))
        switch (fi13 >> ((65535) & 63)) {
        case 391:
            if (2 > 672) {
                fmt.Println(p0, fi13)
                v14 := GenS1{F0: 2147483648, F1: 65535, F2: 42}
                for fi15 := 0; fi15 < 8; fi15++ {
                    fmt.Println(fi13)
                    var v16 float64 = (-2.25 + (3.14159 * 0.1))
                    _ = fi15
                    _ = v16
                }
                _ = v14
            }
        default:
            v17 := GenS1{F0: 965, F1: 2147483648, F2: 633}
            _ = v17
        }
        _ = fi13
    }
    fmt.Println(p0)
    var v18 int = 9223372036854775807
    _ = p0
    _ = v18
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = 45
    v3 := make(map[string]int)
    var v4 uint8 = 127
    var v5 int8 = (-(v2))
    v4 = v4
    v19 := GenS2{F0: v4}
    fmt.Println(v5, v2, v5)
    v4 += v4
    switch 32768 {
    case -9223372036854775808:
        for fi20 := 0; fi20 < 7; fi20++ {
            v21 := GenS2{F0: v4}
            var v22 int = fi20
            _ = fi20
            _ = v21
            _ = v22
        }
    case 32767:
        switch "hello" {
        case "\n\t":
            v23 := GenS1{F0: 65535, F1: 612, F2: 42}
            _ = v23
        case "":
            v24 := GenS3{F0: 1000, F1: 255, F2: 2}
            _ = v24
        case "foo bar":
            var v25 bool = false
            _ = v25
        }
    case 15:
        v26 := []int{128}
        _ = v26
    default:
        var v27 int8 = (v2 & v5)
        _ = v27
    }
    if true {
        v4 = (v4 - v4)
    }
    v28 := GenS2{F0: v4}
    var v29 float64 = (float64(float32(0.0)) / 1.5)
    v30 := v3["\n\t"]
    if ((((v5 * v2) >= (v2 << ((833) & 63))) && (!((0 < -347) || false))) || true) {
        v31 := GenS3{F0: 1000, F1: 65535, F2: 65535}
        _ = v31
    } else {
        if true {
            v1[("\n\t" + ("hello" + ("\n\t" + "a")))] = 564
        }
        var v32 uint = 0
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}

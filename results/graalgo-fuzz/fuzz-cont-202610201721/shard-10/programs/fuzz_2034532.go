package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS3 struct { F0 int16; F1 int32; F2 uint8 }
type GenS4 struct { F0 int8; F1 int16 }
func genFn0(p0 int32, p1 bool) int32 {
    if (!true) {
        v7 := []int{128, 32768, 19}
        v8 := GenS1{F0: false}
        for fi9 := 0; fi9 < 3; fi9++ {
            if (1 <= 4294967295) {
                var v10 string = "hello"
                v11 := GenS2{F0: 2, F1: 4294967295, F2: 2147483647}
                _ = v10
                _ = v11
            }
            _ = fi9
        }
        _ = v7
        _ = v8
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        v13 := GenS3{F0: -339, F1: 1, F2: 37}
        switch (fi12 * (fi12 * (fi12 % (((fi12 | 32767)) | 1)))) {
        case 4294967295:
            var v14 uint8 = 2
            _ = v14
        default:
            v15 := []int{4294967295, (^(fi12))}
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    v16 := GenS4{F0: -128, F1: -189}
    for fi17 := 0; fi17 < 12; fi17++ {
        fmt.Println(p0, p0, p0)
        _ = fi17
    }
    _ = p0
    _ = p1
    _ = v16
    return (p0 % (((p0 * p0)) | 1))
}
func genFn1(p0 bool, p1 uint, p2 string) uint32 {
    var v18 int16 = 32767
    var v19 int = 0
    fmt.Println(p1, p2, p2)
    for fi20 := 0; fi20 < 17; fi20++ {
        p2 = ((p2 + p2) + (("" + p2) + p2))
        v18 = 1000
        v21 := []int{}
        _ = fi20
        _ = v21
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return 2147483647
}

func main() {
    var v1 string = ("hello" + "foo bar")
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]string)
    var v5 uint8 = 42
    var v6 int64 = -20
    var v22 int8 = 100
    for fi23 := 0; fi23 < 6; fi23++ {
        var v24 int = (fi23 ^ -9223372036854775808)
        v1 = (v1 + "foo bar")
        _ = fi23
        _ = v24
    }
    if (true || true) {
        var v25 uint8 = v5
        _ = v25
    } else {
        fmt.Println(v6)
        switch (v1 + (((v1 + v1) + (v1 + v1)) + v1)) {
        case "\n\t":
            delete(v2, 2147483648)
        case "a":
            fmt.Println(v22, v6, v22)
        default:
            var v26 uint8 = (v5 - (v5 - v5))
            _ = v26
        }
    }
    v27 := v4[(v1 + ((v1 + (v1 + ((v1 + "a") + ("\n\t" + "foo bar")))) + v1))]
    v2[0] = 32768
    if (!(!false)) {
        var v28 uint = 4294967295
        _ = v28
    }
    fmt.Println(v27, v6, v27)
    fmt.Println(v22)
    var v29 uint = 468
    var v30 uint16 = 65535
    v30 += (v30 << ((857) & 63))
    if (3.14159 != 1.5) {
        v31 := GenS2{F0: 2, F1: 910, F2: 65535}
        delete(v4, (("a" + v1) + v1))
        v22 = 100
        _ = v31
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}

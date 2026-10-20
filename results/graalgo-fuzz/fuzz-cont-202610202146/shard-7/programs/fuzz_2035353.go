package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 int8 }
func genFn0(p0 uint8, p1 int8) int32 {
    p0 = p0
    if (true && (false && true)) {
        for fi6 := 0; fi6 < 20; fi6++ {
            var v7 uint64 = 1
            p0 -= p0
            if true {
                v8 := GenS1{F0: (^(p0))}
                for fi9 := 0; fi9 < 15; fi9++ {
                    p1 = -128
                    fmt.Println(fi9)
                    _ = fi9
                }
                _ = v8
            } else {
                var v10 uint32 = 65535
                _ = v10
            }
            _ = fi6
            _ = v7
        }
    }
    p1 += (-(p1))
    v11 := []int{}
    if (18446744073709551615 != 157) {
        var v12 int16 = int16(p1)
        switch 128 {
        case 0:
            var v13 uint32 = 4294967295
            _ = v13
        }
        _ = v12
    } else {
        v14 := GenS1{F0: 255}
        fmt.Println(p1, p1)
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v11
    return -2147483648
}
func genFn1(p0 uint16) uint8 {
    var v15 float64 = (float64(1.5) / 3.14159)
    v15 -= -2.25
    fmt.Println(v15)
    _ = p0
    _ = v15
    return 0
}

func main() {
    var v1 uint8 = 127
    var v2 uint16 = 42
    var v3 int8 = -125
    v4 := []uint32{756, 2}
    v5 := make(map[uint32]bool)
    fmt.Println(v1, v1, v2)
    v2 += (v2 - v2)
    v2 = (v2 ^ 278)
    var v16 uint32 = 65535
    v17 := GenS2{F0: 0, F1: 1}
    fmt.Println(v1, v1, v3)
    v1 += (v1 * v1)
    v18 := GenS1{F0: (v1 - 2)}
    v2 += 2
    fmt.Println(v16, v16, v1)
    switch 2147483647 {
    case 266:
        fmt.Println(v16, v2)
    case 706:
        for fi19 := 0; fi19 < 19; fi19++ {
            delete(v5, v16)
            var v20 uint = 4294967295
            _ = fi19
            _ = v20
        }
    case -85:
        for fi21 := 0; fi21 < 4; fi21++ {
            v22 := GenS3{F0: 0}
            _ = fi21
            _ = v22
        }
    default:
        v23 := GenS2{F0: 980, F1: -1}
        _ = v23
    }
    switch "foo bar" {
    case "\n\t":
        v3 = v3
    case "a":
        v24 := GenS3{F0: 42}
        _ = v24
    }
    for fi25 := 0; fi25 < 10; fi25++ {
        var v26 int64 = -942
        _ = fi25
        _ = v26
    }
    v27 := GenS1{F0: 128}
    if false {
        for fi28 := 0; fi28 < 19; fi28++ {
            var v29 float64 = 0.0
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v30 bool = (((v3 << ((884) & 63)) != -68) || true)
            _ = fi28
            _ = v29
            _ = v30
        }
        if (65535 >= 0) {
            v2 = 65535
            switch -1 {
            case -2147483648:
                v1 += 218
            case 718:
                var v31 uint16 = (^(v2))
                _ = v31
            default:
                v4 = append(v4, (v16 * v16))
            }
        } else {
            switch (("foo bar" + ("a" + "foo bar")) + "") {
            case "a":
                var v32 uint16 = v2
                _ = v32
            case "":
                fmt.Println(v1)
            default:
                if (!true) {
                    var v33 int16 = -32768
                    _ = v33
                } else {
                    var v34 uint16 = v2
                    _ = v34
                }
            }
        }
    } else {
        var v35 int32 = 42
        _ = len(v4)
        _ = v35
    }
    v36 := GenS4{F0: v2, F1: 0}
    fmt.Println(v3, v2)
    fmt.Println(v1, v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}

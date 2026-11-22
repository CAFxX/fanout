package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint8) uint32 {
    var v7 float64 = (1.5 * float64(0.0))
    if ((("a" + "hello") + "\n\t") > "foo bar") {
        fmt.Println(p0, v7)
    } else {
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 string = "\n\t"
            v9 = ((v9 + (v9 + v9)) + "a")
            v9 = v9
            _ = fi8
            _ = v9
        }
        switch -564 {
        case 9223372036854775807:
            var v10 uint = 2
            _ = v10
        case -2147483648:
            var v11 uint64 = 18446744073709551615
            _ = v11
        default:
            for fi12 := 0; fi12 < 2; fi12++ {
                v13 := GenS1{F0: 2, F1: 2147483647}
                v7 += float64(-2.25)
                _ = fi12
                _ = v13
            }
        }
        var v14 uint16 = 1
        _ = v14
    }
    if (!(!false)) {
        var v15 float32 = (1e10 - ((0.1 / 1.5) / 1.5))
        v16 := GenS1{F0: 4294967295, F1: 766}
        _ = v15
        _ = v16
    } else {
        if ("a" >= "a") {
            var v17 uint16 = 858
            _ = v17
        }
    }
    switch 784 {
    case 32768:
        v7 -= v7
    default:
        fmt.Println(p0, p0)
    }
    _ = p0
    _ = v7
    return 0
}
func genFn1(p0 uint16, p1 uint64, p2 int8) uint32 {
    switch ("a" + "hello") {
    case "foo bar":
        var v18 uint8 = 14
        _ = v18
    case "":
        v19 := GenS2{F0: 4294967295, F1: 255}
        _ = v19
    default:
        for fi20 := 0; fi20 < 2; fi20++ {
            for fi21 := 0; fi21 < 20; fi21++ {
                switch (-(fi20)) {
                case -9223372036854775808:
                    v22 := []int{fi21, (fi21 ^ -1)}
                    _ = v22
                default:
                    var v23 bool = true
                    _ = v23
                }
                _ = fi21
            }
            v24 := []int{fi20, (-(fi20)), 1, (^(fi20))}
            _ = fi20
            _ = v24
        }
    }
    var v25 uint64 = (p1 | (p1 % ((2) | 1)))
    var v26 int16 = 448
    v27 := GenS3{F0: -1}
    v28 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    return 42
}

func main() {
    var v1 int16 = -285
    v2 := make(map[uint32]uint32)
    var v3 float64 = 0.0
    var v4 uint8 = 1
    var v5 string = "\n\t"
    var v6 int32 = -2147483648
    v29 := []int{928, 1, 0}
    for fi30 := 0; fi30 < 8; fi30++ {
        if (true && ((!true) && true)) {
            v4 = v4
            v31 := GenS1{F0: 42, F1: 2147483647}
            _ = v31
        } else {
            var v32 int8 = 127
            v3 -= float64((float32((v3 / v3)) - (1.5 / 0.1)))
            _ = v32
        }
        _ = len(v2)
        fmt.Println(v4, v1, fi30)
        _ = fi30
    }
    fmt.Println(v5, v4, v5)
    var v33 int16 = v1
    fmt.Println(v1)
    switch ("" + "hello") {
    case "hello":
        var v34 int64 = int64(v6)
        _ = v34
    case "\n\t":
        v33 -= v1
    case "a":
        v35 := GenS3{F0: -1}
        _ = v35
    }
    var v36 uint32 = 4294967295
    var v37 int64 = -1
    var v38 bool = false
    v5 = v5
    if v38 {
        v39 := v2[(v36 - v36)]
        v2[2147483648] = v36
        v38 = (((!(!(v4 > v4))) && (!(387 == 65535))) && false)
        _ = v39
    } else {
        if ((v38 || v38) || (-1 > (v6 & (v6 + 446)))) {
            var v40 float32 = 0.1
            _ = v40
        }
        for fi41 := 0; fi41 < 7; fi41++ {
            v42 := GenS3{F0: -1}
            delete(v2, genFn1(42, 18446744073709551615, -1))
            _ = fi41
            _ = v42
        }
        delete(v2, (v36 / ((genFn1(237, 2, 0)) | 1)))
    }
    v5 = "foo bar"
    v1 = v33
    v43 := GenS1{F0: (v36 >> ((319) & 63)), F1: v36}
    for fi44 := 0; fi44 < 8; fi44++ {
        v1 = (v1 - v33)
        var v45 uint8 = 1
        _ = fi44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}

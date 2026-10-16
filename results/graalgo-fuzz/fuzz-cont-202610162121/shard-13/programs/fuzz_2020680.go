package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS5 struct { F0 uint; F1 int32 }
func genFn0() bool {
    var v5 uint32 = 2
    v6 := GenS1{F0: 580, F1: 0}
    _ = v5
    _ = v6
    return (509 != 65535)
}
func genFn1(p0 int64, p1 int16) string {
    var v7 int16 = 32767
    var v8 int8 = -96
    for fi9 := 0; fi9 < 12; fi9++ {
        switch "" {
        case "a":
            v10 := []int{(fi9 * (fi9 % ((fi9) | 1))), 2147483647, fi9, fi9}
            _ = v10
        case "":
            switch "hello" {
            case "a":
                fmt.Println(p0)
            case "hello":
                v11 := GenS2{F0: 1}
                _ = v11
            case "foo bar":
                v12 := GenS1{F0: 11, F1: 53}
                _ = v12
            }
        case "hello":
            v8 -= v8
        }
        var v13 float32 = float32(1.5)
        var v14 float32 = v13
        _ = fi9
        _ = v13
        _ = v14
    }
    fmt.Println(p1, v7, p1)
    p0 = (p0 | -1)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return "\n\t"
}

func main() {
    var v1 float32 = 0.1
    v2 := make(map[uint32]uint32)
    var v3 int64 = 0
    var v4 int32 = -899
    var v15 float32 = 0.0
    fmt.Println(v1, v3)
    v16 := GenS1{F0: 18446744073709551615, F1: 915}
    v17 := []int{9223372036854775807}
    for fi18 := 0; fi18 < 13; fi18++ {
        v1 -= float32(1e10)
        _ = len(v17)
        for fi19 := 0; fi19 < 14; fi19++ {
            v20 := GenS3{F0: 159}
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        for fi22 := 0; fi22 < 14; fi22++ {
            switch fi22 {
            case -9223372036854775808:
                fmt.Println(fi21)
            }
            _ = fi22
        }
        _ = fi21
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        fmt.Println(v4)
        fmt.Println(v15, v4)
        _ = fi23
    }
    for fi24 := 0; fi24 < 12; fi24++ {
        var v25 bool = false
        v3 = (v3 % ((-477) | 1))
        _ = fi24
        _ = v25
    }
    switch 128 {
    case -2147483648:
        var v26 int16 = 297
        _ = v26
    case 127:
        v27 := GenS1{F0: 736, F1: 1}
        _ = v27
    case 1:
        switch ("hello" + "") {
        case "hello":
            v4 -= v4
        case "a":
            if false {
                v15 = ((v1 * 0.1) * v15)
            } else {
                v2[2] = 42
            }
        case "\n\t":
            var v28 int16 = int16(v4)
            _ = v28
        }
    }
    v29 := GenS2{F0: 1}
    for fi30 := 0; fi30 < 3; fi30++ {
        for fi31 := 0; fi31 < 10; fi31++ {
            v32 := GenS4{F0: 900, F1: 65, F2: 2}
            var v33 uint = 605
            v34 := GenS1{F0: 18446744073709551615, F1: 777}
            _ = fi31
            _ = v32
            _ = v33
            _ = v34
        }
        _ = fi30
    }
    v35 := GenS1{F0: 648, F1: 32767}
    if (255 >= 241) {
        if (!false) {
            v4 = v4
            for fi36 := 0; fi36 < 14; fi36++ {
                v15 = (v15 + v1)
                var v37 float64 = (((0.1 - 0.0) * ((((0.1 - 0.1) - 0.0) - 3.14159) + 0.1)) - float64(v1))
                v4 -= (-(v4))
                _ = fi36
                _ = v37
            }
        }
    } else {
        if len(v17) > 0 {
            v17[0] = 9223372036854775807
        }
        if (255 > 127) {
            v38 := GenS1{F0: 4294967295, F1: 217}
            var v39 int32 = (v4 / (((v4 | v4)) | 1))
            v1 += float32((0.1 - 1.5))
            _ = v38
            _ = v39
        } else {
            v40 := GenS5{F0: 42, F1: v4}
            var v41 uint16 = 934
            _ = v40
            _ = v41
        }
        v1 -= (1e10 + float32((1.5 - float64(v15))))
    }
    v4 += v4
    var v42 uint32 = 2147483648
    for fi43 := 0; fi43 < 19; fi43++ {
        v4 += v4
        v15 += (-2.25 * v15)
        v44 := v2[v42]
        _ = fi43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v35
    _ = v4
    _ = v42
}

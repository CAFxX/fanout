package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint64; F1 int; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS4 struct { F0 uint64; F1 string; F2 uint8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int16; F1 int8; F2 uint32 }
type GenS7 struct { F0 int16; F1 int64 }
type GenS8 struct { F0 uint64 }
func genFn0(p0 int, p1 string) int64 {
    switch p0 {
    case 426:
        var v7 uint64 = 18446744073709551615
        _ = v7
    case 42:
        switch 32767 {
        case 128:
            switch (p1 + (p1 + "\n\t")) {
            case "":
                var v8 uint = 414
                _ = v8
            }
        case 4294967295:
            p0 += (p0 | (p0 >> ((255) & 63)))
        }
    default:
        for fi9 := 0; fi9 < 17; fi9++ {
            v10 := GenS1{F0: 128, F1: 1, F2: 4294967295}
            _ = fi9
            _ = v10
        }
    }
    p0 = 32767
    v11 := GenS1{F0: -1, F1: 255, F2: 521}
    p1 = p1
    _ = p0
    _ = p1
    _ = v11
    return 9223372036854775807
}
func genFn1() uint64 {
    v12 := GenS2{F0: 1, F1: -885, F2: 1}
    if ((-2.25 - 3.14159) >= (1e10 - (float32(float64((float32((3.14159 + float64(float32((-2.25 / 1.0))))) + ((3.14159 + 0.0) / 1.0)))) / 1.0))) {
        var v13 int = -71
        var v14 int64 = -1
        v15 := GenS3{F0: 42, F1: 65535, F2: 1}
        _ = v13
        _ = v14
        _ = v15
    }
    v16 := GenS4{F0: 18446744073709551615, F1: "hello", F2: 127}
    v17 := GenS5{F0: 1}
    var v18 int8 = 100
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    return 18446744073709551615
}

func main() {
    var v1 float64 = float64((0.0 / 1.0))
    v2 := make(map[string]bool)
    var v3 float64 = v1
    var v4 uint8 = 1
    var v5 string = ""
    var v6 int16 = 0
    if false {
        _ = len(v2)
        v19 := []int{-494, 2147483647, 630}
        _ = v19
    } else {
        v20 := []int{852, 766}
        delete(v2, (v5 + v5))
        _ = v20
    }
    var v21 uint8 = 1
    v22 := []int{32767, 1, 32767}
    if ((!(!true)) && ((v21 & v4) != 127)) {
        v3 = -2.25
    }
    var v23 uint32 = 2147483648
    switch -9223372036854775808 {
    case -396:
        v4 = v21
    default:
        for fi24 := 0; fi24 < 4; fi24++ {
            v25 := GenS6{F0: (-(v6)), F1: 100, F2: (v23 >> ((387) & 63))}
            var v26 bool = ((v4 % ((0) | 1)) != (v21 % (((v4 | 127)) | 1)))
            _ = fi24
            _ = v25
            _ = v26
        }
    }
    switch 125 {
    case 4294967295:
        v27 := GenS7{F0: v6, F1: 1}
        _ = v27
    case 127:
        var v28 uint8 = (v21 + 1)
        _ = v28
    default:
        fmt.Println(v5, v6, v21)
    }
    if false {
        for fi29 := 0; fi29 < 9; fi29++ {
            switch -816 {
            case -1:
                var v30 int8 = -118
                _ = v30
            case 12:
                var v31 float64 = v1
                _ = v31
            case -705:
                fmt.Println(v6)
            }
            var v32 int16 = 420
            _ = fi29
            _ = v32
        }
        if len(v22) > 0 {
            v22[0] = -914
        }
    } else {
        v33 := GenS8{F0: 2}
        switch -9223372036854775808 {
        case -2147483648:
            v2[v5] = ((true && false) && true)
        case 32768:
            switch v5 {
            case "":
                _ = len(v2)
            case "hello":
                v23 -= v23
            default:
                switch int(v6) {
                case 1:
                    var v34 int8 = int8(v4)
                    _ = v34
                }
            }
        default:
            v23 = 0
        }
        _ = v33
    }
    if len(v22) > 0 {
        v22[0] = 127
    }
    if false {
        fmt.Println(v23, v3, v4)
    } else {
        v5 = "hello"
        var v35 float32 = (1e10 + ((3.14159 - 1e10) / (0.1 / -2.25)))
        v21 = 255
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

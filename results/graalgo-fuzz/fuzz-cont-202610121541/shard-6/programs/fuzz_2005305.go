package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 int16; F1 uint32 }
func genFn0(p0 uint8, p1 int16, p2 uint32) string {
    v5 := GenS1{F0: 989, F1: true, F2: 0}
    fmt.Println(p2, p0, p2)
    v6 := GenS2{F0: int32(p1)}
    v7 := []int{-613, 229, 32768, 2147483647}
    var v8 uint64 = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return (("\n\t" + "\n\t") + ("\n\t" + (("hello" + ("foo bar" + "\n\t")) + "\n\t")))
}
func genFn1(p0 int, p1 uint32, p2 int) int32 {
    if (!true) {
        var v9 int64 = 9223372036854775807
        var v10 bool = true
        _ = v10
        _ = v9
    }
    fmt.Println(p2, p1)
    var v11 int = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 2147483647
}

func main() {
    var v1 int = 9223372036854775807
    var v2 uint64 = 42
    v3 := []uint32{42, 104}
    var v4 int16 = 372
    for fi12 := 0; fi12 < 2; fi12++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v2 = v2
        switch (v1 / (((v1 + 0)) | 1)) {
        case 1:
            for fi13 := 0; fi13 < 9; fi13++ {
                v14 := GenS3{F0: true, F1: 127}
                _ = fi13
                _ = v14
            }
        case 9223372036854775807:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            fmt.Println(v4, v4, v1)
        }
        _ = fi12
    }
    var v15 uint64 = v2
    var v16 uint64 = (v15 & 4294967295)
    switch (v1 | 42) {
    case -354:
        switch 930 {
        case 32768:
            fmt.Println(v1)
        case 2147483647:
            var v17 int16 = (v4 / ((-455) | 1))
            _ = v17
        case -749:
            var v18 int8 = 100
            _ = v18
        }
    case 127:
        for fi19 := 0; fi19 < 6; fi19++ {
            var v20 float64 = (float64(3.14159) + float64(3.14159))
            _ = len(v3)
            _ = fi19
            _ = v20
        }
    case 1:
        if ((v2 ^ (v15 ^ (v16 / (((v15 >> ((4294967295) & 63))) | 1)))) <= v2) {
            if len(v3) > 0 {
                v3[0] = 2
            }
            v3 = append(v3, 719)
        } else {
            v16 -= v15
        }
    default:
        for fi21 := 0; fi21 < 7; fi21++ {
            var v22 int64 = -332
            _ = fi21
            _ = v22
        }
    }
    if true {
        v15 -= (v15 - v16)
    } else {
        if false {
            _ = len(v3)
            var v23 bool = ((0.1 / 3.14159) <= 0.0)
            if len(v3) > 0 {
                v3[0] = 42
            }
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        for fi25 := 0; fi25 < 20; fi25++ {
            switch fi25 {
            case -1:
                if false {
                    var v26 uint8 = 255
                    var v27 int64 = -832
                    v28 := GenS4{F0: (v27 / ((v27) | 1)), F1: ""}
                    var v29 int64 = -1
                    var v30 int16 = (v4 * -351)
                    _ = v26
                    _ = v27
                    _ = v28
                    _ = v29
                    _ = v30
                }
            }
            _ = fi25
        }
        _ = fi24
    }
    if true {
        for fi31 := 0; fi31 < 8; fi31++ {
            for fi32 := 0; fi32 < 6; fi32++ {
                v15 = 42
                _ = len(v3)
                _ = fi32
            }
            if len(v3) > 0 {
                v3[0] = 204
            }
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
            _ = fi31
        }
        v16 = (v15 / ((v16) | 1))
    } else {
        if true {
            v15 = 496
            var v33 int8 = 127
            _ = v33
        } else {
            var v34 int = 128
            switch "" {
            case "foo bar":
                v35 := GenS2{F0: genFn1(32768, 206, v34)}
                _ = v35
            default:
                v36 := GenS3{F0: (2 > 148), F1: 255}
                _ = v36
            }
            for fi37 := 0; fi37 < 15; fi37++ {
                for fi38 := 0; fi38 < 7; fi38++ {
                    fmt.Println(v2, fi38, fi37)
                    v15 = v2
                    _ = fi38
                }
                _ = fi37
            }
            _ = v34
        }
    }
    v4 += v4
    var v39 uint = 0
    v40 := GenS5{F0: (v4 / (((v4 % (((^(v4))) | 1))) | 1)), F1: 783}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v3 = append(v3, 2147483647)
    if len(v3) > 0 {
        v3[0] = 464
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint64; F1 int64; F2 int }
type GenS3 struct { F0 string; F1 uint64; F2 string }

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 65535
    var v3 uint8 = 1
    v4 := make(map[int]bool)
    for fi5 := 0; fi5 < 4; fi5++ {
        v2 = 162
        _ = fi5
    }
    for fi6 := 0; fi6 < 17; fi6++ {
        for fi7 := 0; fi7 < 15; fi7++ {
            fmt.Println(v2, fi6)
            switch "hello" {
            case "hello":
                v3 += (v3 - v3)
            case "":
                v8 := GenS1{F0: (-(fi6)), F1: 474}
                v9 := v4[fi7]
                switch 2147483647 {
                case -721:
                    if true {
                        switch (fi6 % (((^(fi6))) | 1)) {
                        case 846:
                            var v10 float32 = 1e10
                            _ = v10
                        case 129:
                            v11 := []int{fi6, (fi6 / (((fi6 | (fi7 + -61))) | 1)), (fi6 % ((fi6) | 1))}
                            fmt.Println(v9, fi7, v3)
                            var v12 int8 = -1
                            _ = v11
                            _ = v12
                        case 127:
                            fmt.Println(v3)
                        }
                        fmt.Println(fi6, fi6)
                        if (-18 != -1) {
                            v13 := []int{}
                            v3 -= (v3 % ((v3) | 1))
                            v3 = v3
                            _ = v13
                        }
                        v9 = (!true)
                        if (((true || false) || ((1.5 / float64(0.1)) >= (1.5 - (((((0.1 / 0.1) * (float64(float32(float64(((0.0 * 0.1) + 0.1)))) - (-2.25 + (1e10 * (((1e10 + 1.5) + float64((0.1 * 0.0))) + 1e10))))) / (float64(3.14159) + float64((3.14159 - 3.14159)))) - 0.0) - 0.0)))) && (780 > 2)) {
                            v14 := GenS1{F0: (fi7 / ((fi7) | 1)), F1: 0}
                            v15 := GenS1{F0: fi6, F1: 1}
                            _ = v14
                            _ = v15
                        }
                        v16 := GenS1{F0: -9223372036854775808, F1: 9}
                        _ = v16
                    }
                case -1:
                    v17 := GenS2{F0: 2, F1: 9223372036854775807, F2: (-(fi7))}
                    _ = v17
                default:
                    v2 = 42
                }
                _ = v8
                _ = v9
            }
            var v18 uint8 = 46
            _ = fi7
            _ = v18
        }
        var v19 uint16 = 42
        if false {
            if ("\n\t" > (("" + ("hello" + "foo bar")) + "")) {
                v3 += (^(v3))
            }
        }
        _ = fi6
        _ = v19
    }
    var v20 int = -2147483648
    if ((v20 % (((v20 << ((849) & 63))) | 1)) >= (^(v20))) {
        v20 = v20
    } else {
        fmt.Println(v3, v2, v20)
        switch (^(v20)) {
        case 9223372036854775807:
            var v21 int8 = 1
            _ = v21
        case 222:
            v2 -= 65535
        case 574:
            v20 -= v20
        default:
            for fi22 := 0; fi22 < 16; fi22++ {
                var v23 int32 = 543
                var v24 int64 = -9223372036854775808
                _ = fi22
                _ = v23
                _ = v24
            }
        }
    }
    for fi25 := 0; fi25 < 1; fi25++ {
        fmt.Println(v2, v3, v3)
        _ = fi25
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        _ = len(v4)
        _ = fi26
    }
    switch ("hello" + "hello") {
    case "":
        if false {
            var v27 int64 = 42
            if false {
                var v28 uint32 = (v2 / (((v2 + v2)) | 1))
                var v29 int64 = (v27 + 9223372036854775807)
                _ = v28
                _ = v29
            } else {
                var v30 bool = true
                _ = v30
            }
            _ = v27
        }
    case "hello":
        fmt.Println(v2, v2)
    case "\n\t":
        var v31 int = v20
        _ = v31
    }
    fmt.Println(v2, v3)
    for fi32 := 0; fi32 < 15; fi32++ {
        v33 := GenS3{F0: (("foo bar" + "a") + "\n\t"), F1: 1, F2: ("" + "\n\t")}
        switch ("foo bar" + "\n\t") {
        case "":
            var v34 uint = 44
            _ = v34
        default:
            var v35 uint64 = 217
            _ = v35
        }
        _ = fi32
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}

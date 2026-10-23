package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int; F1 uint32; F2 uint8 }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int32) bool {
    v7 := GenS1{F0: 380, F1: 0}
    v8 := GenS1{F0: 4294967295, F1: 403}
    _ = p0
    _ = v7
    _ = v8
    return true
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = -9223372036854775808
    var v3 int16 = -344
    v4 := []int{v2, v2}
    var v5 int16 = v3
    v6 := make(map[uint32]int)
    fmt.Println(v5, v2, v3)
    v9 := GenS1{F0: 18446744073709551615, F1: 42}
    if len(v4) > 0 {
        v4[0] = (v2 + (v2 << ((18446744073709551615) & 63)))
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        v2 = 0
        var v11 int = (-(fi10))
        _ = fi10
        _ = v11
    }
    v2 = -996
    switch "" {
    case "":
        if false {
            if (v2 != (v2 >> ((2) & 63))) {
                v12 := GenS1{F0: 4294967295, F1: 639}
                _ = v12
            }
            for fi13 := 0; fi13 < 10; fi13++ {
                fmt.Println(fi13)
                v14 := GenS2{F0: (fi13 << ((18446744073709551615) & 63)), F1: 75, F2: 128}
                _ = fi13
                _ = v14
            }
            var v15 int = 127
            _ = v15
        } else {
            v5 = (v3 & (v5 + 32767))
            var v16 int64 = 746
            _ = v16
        }
    case "\n\t":
        for fi17 := 0; fi17 < 5; fi17++ {
            if false {
                fmt.Println(fi17, v5, v5)
            }
            v18 := GenS1{F0: 1, F1: 0}
            for fi19 := 0; fi19 < 10; fi19++ {
                v20 := GenS2{F0: (fi17 + fi17), F1: 2147483648, F2: 1}
                switch "" {
                case "foo bar":
                    v21 := GenS3{F0: (true || true), F1: 810}
                    _ = v21
                }
                var v22 float32 = (1e10 / (1e10 - 1.5))
                _ = fi19
                _ = v20
                _ = v22
            }
            _ = fi17
            _ = v18
        }
    case "a":
        if (861 > 2) {
            v23 := GenS2{F0: (v2 % (((v2 - v2)) | 1)), F1: 347, F2: 0}
            var v24 uint32 = 910
            _ = v23
            _ = v24
        } else {
            fmt.Println(v5, v2, v5)
            var v25 int64 = 9223372036854775807
            _ = v25
        }
    default:
        if (true && true) {
            for fi26 := 0; fi26 < 15; fi26++ {
                v27 := GenS4{F0: 926}
                v28 := GenS2{F0: -2147483648, F1: 1, F2: 1}
                if len(v4) > 0 {
                    v4[0] = (fi26 << ((0) & 63))
                }
                _ = fi26
                _ = v27
                _ = v28
            }
            switch ("a" + "\n\t") {
            case "foo bar":
                v2 = v2
            case "a":
                v4 = append(v4, v2)
            default:
                var v29 int32 = 992
                _ = v29
            }
        }
    }
    var v30 float64 = 1e10
    var v31 int8 = 0
    if true {
        var v32 float64 = (1.5 + float64((3.14159 * 0.0)))
        _ = v32
    } else {
        v2 += 762
        v31 += v31
    }
    var v33 uint64 = 1
    var v34 int64 = 42
    for fi35 := 0; fi35 < 18; fi35++ {
        v33 = 311
        _ = fi35
    }
    var v36 int8 = v31
    v30 += float64(1e10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v34
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}

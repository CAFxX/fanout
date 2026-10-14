package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 string }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 int64; F1 int64; F2 uint16 }
type GenS4 struct { F0 int64; F1 bool; F2 uint8 }
func genFn0(p0 int, p1 bool, p2 int16) int32 {
    p2 += p2
    var v4 int8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -532
}
func genFn1(p0 int16) string {
    switch 32768 {
    case -1:
        var v5 uint32 = 938
        _ = v5
    case 9223372036854775807:
        switch (("a" + ("" + "hello")) + (((("" + "foo bar") + (("hello" + ("" + "a")) + ("foo bar" + "a"))) + "") + ("hello" + "a"))) {
        case "a":
            v6 := GenS1{F0: 2, F1: 0, F2: (("" + "a") + "\n\t")}
            _ = v6
        case "hello":
            p0 = (-(p0))
        case "foo bar":
            for fi7 := 0; fi7 < 17; fi7++ {
                p0 = (p0 * p0)
                var v8 int64 = 0
                v9 := []int{(fi7 >> ((741) & 63)), -9223372036854775808, fi7, 0}
                _ = fi7
                _ = v8
                _ = v9
            }
        }
    case 127:
        if false {
            fmt.Println(p0)
            var v10 uint64 = 442
            _ = v10
        }
    default:
        var v11 int8 = 42
        _ = v11
    }
    v12 := GenS2{F0: (true || (false || true)), F1: -81}
    _ = p0
    _ = v12
    return "hello"
}

func main() {
    var v1 int8 = -1
    var v2 uint = 821
    var v3 bool = false
    for fi13 := 0; fi13 < 16; fi13++ {
        if (!(!(1 <= 255))) {
            v14 := GenS1{F0: 0, F1: 2, F2: "\n\t"}
            var v15 uint64 = 18446744073709551615
            v3 = v3
            _ = v14
            _ = v15
        }
        _ = fi13
    }
    v16 := GenS2{F0: v3, F1: -109}
    for fi17 := 0; fi17 < 16; fi17++ {
        var v18 bool = true
        _ = fi17
        _ = v18
    }
    v2 = v2
    fmt.Println(v2, v2)
    for fi19 := 0; fi19 < 18; fi19++ {
        var v20 int64 = 0
        var v21 bool = true
        if (true && (255 >= 127)) {
            fmt.Println(fi19)
            v20 = v20
        } else {
            for fi22 := 0; fi22 < 18; fi22++ {
                switch fi19 {
                case -885:
                    var v23 uint8 = 56
                    v2 = 18446744073709551615
                    _ = v23
                case -464:
                    var v24 int16 = -839
                    _ = v24
                default:
                    v25 := []int{(^(fi22))}
                    _ = v25
                }
                v26 := []int{32767, fi22}
                _ = fi22
                _ = v26
            }
            if (0.1 < 3.14159) {
                var v27 float64 = (float64(0.1) - 3.14159)
                if (v1 <= (v1 * (v1 << ((783) & 63)))) {
                    switch (fi19 & (fi19 * -2147483648)) {
                    case -9223372036854775808:
                        v3 = (1 == (fi19 << (((v2 + 42)) & 63)))
                    case 32768:
                        switch "" {
                        case "":
                            if true {
                                v1 += -128
                            } else {
                                var v28 float32 = float32(float64(3.14159))
                                _ = v28
                            }
                        default:
                            var v29 int = int(v1)
                            _ = v29
                        }
                    default:
                        for fi30 := 0; fi30 < 6; fi30++ {
                            v31 := GenS3{F0: (v20 | (-(v20))), F1: v20, F2: 255}
                            var v32 bool = v21
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                    }
                }
                _ = v27
            }
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println(v1, v1)
    v33 := GenS4{F0: 0, F1: ((v1 % (((v1 * v1)) | 1)) == v1), F2: 127}
    switch 4294967295 {
    case 0:
        for fi34 := 0; fi34 < 11; fi34++ {
            v2 = v2
            _ = fi34
        }
    default:
        v2 += v2
    }
    if true {
        var v35 uint16 = 32767
        _ = v35
    }
    var v36 uint64 = 42
    v37 := []int{int(v36), 127}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v37
}

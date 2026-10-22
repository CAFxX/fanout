package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint8; F2 int }
type GenS3 struct { F0 uint; F1 uint16; F2 uint8 }
func genFn0(p0 int32) uint16 {
    v6 := GenS1{F0: -66, F1: 4294967295}
    v7 := GenS1{F0: int64(p0), F1: 65535}
    for fi8 := 0; fi8 < 12; fi8++ {
        p0 += (p0 & 0)
        _ = fi8
    }
    p0 -= (p0 % (((p0 << ((18446744073709551615) & 63))) | 1))
    _ = p0
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = 3.14159
    v3 := make(map[uint32]int)
    var v4 float32 = (float32(3.14159) / float32(v2))
    v5 := []string{"\n\t", "hello", "\n\t", "a"}
    v3[42] = 127
    switch 506 {
    case 32768:
        v2 += -2.25
    case 127:
        if (true || true) {
            var v9 int16 = 691
            fmt.Println(v9, v2, v2)
            if (!((128 == 255) || (!(2 != 0)))) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v4 -= v4
            } else {
                fmt.Println(v9, v9, v4)
            }
            _ = v9
        } else {
            var v10 uint = 42
            fmt.Println(v4, v2, v4)
            _ = v10
        }
    default:
        var v11 uint32 = 2147483648
        _ = v11
    }
    var v12 string = "foo bar"
    v12 = ("a" + "")
    var v13 bool = false
    v14 := GenS1{F0: 534, F1: 2147483647}
    switch (((v12 + "a") + "hello") + "") {
    case "foo bar":
        if len(v5) > 0 {
            v5[0] = v12
        }
    case "a":
        v12 = "\n\t"
    case "":
        var v15 uint8 = 128
        _ = v15
    default:
        var v16 uint64 = 224
        _ = v16
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        if (v12 == "") {
            for fi18 := 0; fi18 < 15; fi18++ {
                fmt.Println(v12)
                _ = fi18
            }
            v2 = v2
        } else {
            if ((v12 + "") > v12) {
                if ((936 != 255) || false) {
                    v19 := GenS2{F0: -1, F1: 12, F2: (fi17 >> ((4294967295) & 63))}
                    fmt.Println(v2, v2)
                    _ = v19
                } else {
                    fmt.Println(v2)
                    v12 = v12
                    var v20 uint16 = 32767
                    _ = v20
                }
                v12 = "hello"
            }
            if v13 {
                v13 = true
                var v21 int32 = 2147483647
                switch fi17 {
                case 174:
                    for fi22 := 0; fi22 < 20; fi22++ {
                        v23 := GenS1{F0: 0, F1: 1}
                        _ = fi22
                        _ = v23
                    }
                case 129:
                    v5 = append(v5, v12)
                case 820:
                    _ = len(v5)
                }
                _ = v21
            } else {
                switch v12 {
                case "\n\t":
                    var v24 uint16 = 42
                    _ = v24
                default:
                    if v13 {
                        var v25 int16 = 0
                        fmt.Println(v4)
                        for fi26 := 0; fi26 < 13; fi26++ {
                            _ = len(v5)
                            v1[0] = 2147483648
                            var v27 bool = ((v4 / float32(v2)) > v4)
                            v28 := GenS3{F0: 763, F1: 65535, F2: 127}
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        _ = v25
                    } else {
                        var v29 uint8 = 52
                        var v30 uint8 = v29
                        _ = v29
                        _ = v30
                    }
                }
                v12 = ((v12 + v12) + ((("" + v12) + v12) + v12))
                fmt.Println(fi17, v2)
            }
        }
        _ = fi17
    }
    var v31 float32 = v4
    fmt.Println(v31)
    switch v12 {
    case "":
        var v32 uint = 18446744073709551615
        _ = v32
    case "\n\t":
        v2 = 1.5
    case "foo bar":
        for fi33 := 0; fi33 < 13; fi33++ {
            if v13 {
                var v34 float32 = (v31 / v31)
                _ = v34
            } else {
                delete(v1, 42)
                fmt.Println(v2)
                _ = len(v1)
            }
            v31 = (v4 / v4)
            v3[0] = (fi33 + fi33)
            _ = fi33
        }
    }
    var v35 float64 = v2
    var v36 uint8 = 42
    v37 := GenS3{F0: 0, F1: genFn0(0), F2: (^(v36))}
    v31 = 1.5
    if true {
        v38 := GenS1{F0: 0, F1: 2}
        _ = v38
    } else {
        var v39 float32 = v4
        fmt.Println(v13)
        v40 := GenS2{F0: 1, F1: v36, F2: 128}
        _ = v39
        _ = v40
    }
    v5 = append(v5, v12)
    var v41 float64 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v41
    _ = v5
}

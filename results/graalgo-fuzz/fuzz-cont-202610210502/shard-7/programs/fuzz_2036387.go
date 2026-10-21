package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint8; F2 int8 }
type GenS3 struct { F0 bool; F1 bool; F2 int8 }
type GenS4 struct { F0 int }
func genFn0(p0 uint8) string {
    var v7 uint64 = 42
    v8 := []int{2147483647, 32768, 0, 127}
    v7 += (v7 * (v7 % (((v7 & 0)) | 1)))
    fmt.Println(v7, v7)
    _ = p0
    _ = v7
    _ = v8
    return ("a" + "")
}

func main() {
    var v1 int = 602
    v2 := make(map[int]int)
    v3 := []string{"foo bar"}
    var v4 int32 = -1
    v5 := []uint32{}
    v6 := make(map[string]bool)
    fmt.Println(v1, v1)
    v9 := GenS1{F0: 0, F1: 65535}
    fmt.Println(v1)
    if (0 == 18446744073709551615) {
        v4 += v4
    } else {
        var v10 uint = 927
        var v11 int32 = (v4 % ((v4) | 1))
        _ = v10
        _ = v11
    }
    switch (-(v1)) {
    case -519:
        v12 := GenS2{F0: 4294967295, F1: 1, F2: -51}
        _ = v12
    case -1:
        var v13 bool = ((v4 >> ((4294967295) & 63)) < (v4 % ((v4) | 1)))
        _ = v13
    case 128:
        if false {
            v3 = append(v3, "a")
        } else {
            fmt.Println(v4)
        }
    }
    if true {
        if false {
            for fi14 := 0; fi14 < 3; fi14++ {
                v1 += (fi14 & fi14)
                fmt.Println(v1, v1)
                var v15 uint = 2
                _ = fi14
                _ = v15
            }
            if (false && true) {
                for fi16 := 0; fi16 < 10; fi16++ {
                    v1 = 127
                    _ = fi16
                }
                v17 := v6["foo bar"]
                _ = v17
            } else {
                v18 := GenS1{F0: -1, F1: 255}
                _ = v18
            }
        } else {
            v4 -= (v4 % ((544) | 1))
        }
        v19 := GenS1{F0: 1000, F1: 931}
        _ = v19
    }
    if true {
        var v20 string = ""
        switch 4294967295 {
        case -9223372036854775808:
            var v21 float64 = (3.14159 - 1e10)
            _ = v21
        case -871:
            var v22 uint16 = 65535
            _ = v22
        case 2147483647:
            for fi23 := 0; fi23 < 14; fi23++ {
                v1 = v1
                _ = fi23
            }
        default:
            for fi24 := 0; fi24 < 8; fi24++ {
                var v25 int32 = v4
                _ = fi24
                _ = v25
            }
        }
        _ = len(v6)
        _ = v20
    } else {
        v1 -= v1
        if (!true) {
            if true {
                switch (v1 << ((42) & 63)) {
                case 694:
                    for fi26 := 0; fi26 < 2; fi26++ {
                        var v27 float32 = (1.5 / 0.1)
                        switch fi26 {
                        case 32768:
                            var v28 uint16 = 65535
                            _ = v28
                        default:
                            v29 := v2[-990]
                            _ = v29
                        }
                        v4 -= -2147483648
                        _ = fi26
                        _ = v27
                    }
                case -324:
                    if false {
                        if ((-2147483648 <= (v1 - v1)) || false) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        } else {
                            var v30 float64 = (1e10 + float64(((1e10 + (1.5 + -2.25)) + (3.14159 / 1.0))))
                            var v31 uint = 0
                            _ = v30
                            _ = v31
                        }
                        _ = len(v3)
                    }
                default:
                    v2[v1] = 127
                }
                var v32 int = -1
                var v33 uint16 = 255
                _ = v32
                _ = v33
            } else {
                v1 = int(v4)
                var v34 int8 = -21
                _ = v34
            }
        }
    }
    switch v1 {
    case -9223372036854775808:
        v4 = (v4 / (((v4 & v4)) | 1))
    case 9223372036854775807:
        v35 := GenS3{F0: ((255 > 128) && (false || false)), F1: (0 != v4), F2: 127}
        _ = v35
    case -2147483648:
        v1 -= 523
    }
    v36 := GenS4{F0: v1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}

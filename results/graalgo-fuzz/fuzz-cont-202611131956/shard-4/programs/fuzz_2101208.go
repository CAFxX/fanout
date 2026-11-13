package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint32; F1 uint8 }
func genFn0() int8 {
    if false {
        var v6 float64 = (1e10 / 1.0)
        switch 1 {
        case -9223372036854775808:
            v6 -= (v6 - (v6 - float64((float32(0.1) * 1e10))))
        case -2147483648:
            v7 := []int{0, 1, 781, -951}
            _ = v7
        default:
            switch -72 {
            case -9223372036854775808:
                switch 2147483647 {
                case 32768:
                    for fi8 := 0; fi8 < 4; fi8++ {
                        v9 := []int{4294967295, -1}
                        _ = fi8
                        _ = v9
                    }
                }
            default:
                fmt.Println(v6)
            }
        }
        var v10 uint64 = 759
        _ = v10
        _ = v6
    } else {
        v11 := []int{42, 2147483647, 4294967295, 2147483647}
        _ = v11
    }
    var v12 uint32 = 2147483647
    v13 := []int{32767, -2147483648, 2147483647, 42}
    for fi14 := 0; fi14 < 1; fi14++ {
        if false {
            v15 := GenS1{F0: 42, F1: 42, F2: -9223372036854775808}
            for fi16 := 0; fi16 < 6; fi16++ {
                v17 := GenS1{F0: -892, F1: 555, F2: 9223372036854775807}
                var v18 int32 = 1
                _ = fi16
                _ = v17
                _ = v18
            }
            var v19 uint8 = 199
            _ = v15
            _ = v19
        }
        _ = fi14
    }
    _ = v12
    _ = v13
    return 0
}

func main() {
    var v1 uint8 = 176
    var v2 bool = false
    v3 := []uint32{}
    var v4 uint32 = 432
    v5 := []string{""}
    var v20 int32 = 2147483647
    v2 = v2
    var v21 uint = 419
    v22 := GenS1{F0: -1, F1: -1, F2: -393}
    for fi23 := 0; fi23 < 19; fi23++ {
        if ((v2 && ((v4 & v4) != (v4 / ((v4) | 1)))) || v2) {
            for fi24 := 0; fi24 < 1; fi24++ {
                _ = len(v3)
                if (42 <= 42) {
                    v25 := GenS1{F0: 104, F1: int32(fi23), F2: 9223372036854775807}
                    fmt.Println(v2, v4, v21)
                    _ = v25
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        v4 = v4
                    case "a":
                        v26 := GenS2{F0: 32767, F1: 4294967295, F2: v20}
                        _ = v26
                    case "":
                        var v27 float32 = (((0.1 - 1e10) * 3.14159) * (1.5 - -2.25))
                        _ = v27
                    }
                    _ = len(v5)
                    v20 += -97
                    if (92 != 4294967295) {
                        switch (fi24 & -1) {
                        case 127:
                            var v28 int16 = -32768
                            _ = v28
                        case 32768:
                            var v29 uint32 = v4
                            if v2 {
                                v1 = v1
                                fmt.Println(v21)
                                v30 := GenS1{F0: -1, F1: (v20 % ((v20) | 1)), F2: int64(v20)}
                                _ = v30
                            }
                            var v31 uint16 = 65535
                            _ = v29
                            _ = v31
                        case 4294967295:
                            var v32 bool = (!v2)
                            _ = v32
                        default:
                            v20 = 2147483647
                        }
                    }
                    v3 = append(v3, v4)
                }
                switch fi24 {
                case 32768:
                    switch fi23 {
                    case 4294967295:
                        _ = len(v3)
                    case 740:
                        fmt.Println(fi24, v20, fi23)
                    }
                }
                _ = fi24
            }
        }
        if len(v5) > 0 {
            v5[0] = "hello"
        }
        v21 = 255
        _ = fi23
    }
    switch (("\n\t" + "\n\t") + (("foo bar" + ("hello" + "")) + "a")) {
    case "foo bar":
        switch "a" {
        case "hello":
            switch 389 {
            case -574:
                v2 = v2
            case 1:
                v2 = ((v4 / ((837) | 1)) > (v4 * (v4 << ((389) & 63))))
            default:
                fmt.Println(v21)
            }
        case "a":
            v4 = (v4 + 65535)
        default:
            v33 := GenS3{F0: uint32(v21), F1: uint8(v21)}
            _ = v33
        }
    case "a":
        var v34 uint16 = 841
        _ = v34
    }
    switch "\n\t" {
    case "a":
        var v35 uint64 = 0
        _ = v35
    case "foo bar":
        v4 = 810
    }
    v1 -= v1
    if ((0.0 - 0.1) >= 0.0) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    var v36 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}

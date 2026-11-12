package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 string; F1 uint32 }

func main() {
    var v1 float32 = (((0.0 * -2.25) * 0.0) * (float32(((-2.25 / (3.14159 / 1e10)) * 3.14159)) / -2.25))
    var v2 int32 = 42
    var v3 uint = 18446744073709551615
    var v4 string = ""
    var v5 int8 = 105
    fmt.Println(v1)
    for fi6 := 0; fi6 < 7; fi6++ {
        switch v4 {
        case "foo bar":
            var v7 uint16 = 2
            _ = v7
        case "":
            switch "\n\t" {
            case "foo bar":
                v8 := GenS1{F0: (fi6 << ((v3) & 63)), F1: "\n\t"}
                _ = v8
            case "a":
                v1 += ((0.1 + v1) - 0.1)
            case "\n\t":
                for fi9 := 0; fi9 < 20; fi9++ {
                    v10 := GenS1{F0: (^(fi6)), F1: "foo bar"}
                    if true {
                        switch v4 {
                        case "\n\t":
                            if (("a" + "\n\t") == v4) {
                                v1 = (((float32(((3.14159 * 1.5) * float64(0.0))) + -2.25) / v1) * -2.25)
                                v11 := GenS2{F0: v4, F1: 940}
                                _ = v11
                            }
                        }
                        v12 := GenS2{F0: (v4 + v4), F1: 0}
                        _ = v12
                    } else {
                        v13 := []int{fi6, 128, fi6, fi9}
                        var v14 bool = ((v3 / (((v3 * v3)) | 1)) >= (v3 * v3))
                        switch v4 {
                        case "\n\t":
                            var v15 int32 = 647
                            _ = v15
                        case "foo bar":
                            if true {
                                v16 := GenS2{F0: "", F1: 0}
                                v1 = (v1 + v1)
                                _ = v16
                            }
                        case "a":
                            v1 = v1
                            v14 = (((v1 - (1.5 + v1)) / 1.0) < (v1 - (0.1 + (v1 / 3.14159))))
                        }
                        _ = v13
                        _ = v14
                    }
                    v17 := []int{(fi9 | fi9), fi9}
                    _ = fi9
                    _ = v10
                    _ = v17
                }
            }
        }
        fmt.Println(v5)
        fmt.Println(v3, v4, v1)
        _ = fi6
    }
    fmt.Println(v4)
    v2 -= v2
    if true {
        fmt.Println(v2, v1)
        for fi18 := 0; fi18 < 19; fi18++ {
            var v19 int16 = -344
            v19 += 42
            v20 := GenS2{F0: ("\n\t" + v4), F1: 1}
            _ = fi18
            _ = v19
            _ = v20
        }
    } else {
        switch -2147483648 {
        case -2147483648:
            fmt.Println(v4, v3, v5)
        case 379:
            switch "foo bar" {
            case "hello":
                if true {
                    var v21 int64 = 42
                    _ = v21
                }
            case "foo bar":
                v22 := []int{-608, 2147483647, 127}
                _ = v22
            default:
                for fi23 := 0; fi23 < 19; fi23++ {
                    for fi24 := 0; fi24 < 2; fi24++ {
                        var v25 int16 = 1
                        v1 = v1
                        fmt.Println(fi23, v3, v2)
                        _ = fi24
                        _ = v25
                    }
                    v26 := []int{(-(fi23)), -513}
                    v27 := GenS1{F0: fi23, F1: ""}
                    _ = fi23
                    _ = v26
                    _ = v27
                }
            }
        default:
            if ((!(true && false)) && (32767 <= 422)) {
                for fi28 := 0; fi28 < 15; fi28++ {
                    var v29 int16 = -817
                    v4 = (((v4 + v4) + v4) + "\n\t")
                    v4 = ("a" + "\n\t")
                    _ = fi28
                    _ = v29
                }
                v3 = 991
                for fi30 := 0; fi30 < 9; fi30++ {
                    if (0 != 42) {
                        var v31 int8 = int8(v2)
                        _ = v31
                    }
                    v2 += (-(v2))
                    var v32 bool = ((v2 - v2) <= (v2 - (v2 - v2)))
                    _ = fi30
                    _ = v32
                }
            }
        }
    }
    v33 := []int{919, 127, 128}
    var v34 string = "a"
    for fi35 := 0; fi35 < 16; fi35++ {
        _ = len(v33)
        _ = fi35
    }
    var v36 uint64 = 42
    var v37 uint64 = 42
    if (!false) {
        var v38 uint32 = 4294967295
        _ = v38
    } else {
        v37 = (^(v36))
    }
    if len(v33) > 0 {
        v33[0] = 4294967295
    }
    v39 := GenS2{F0: (v4 + (v34 + v34)), F1: 2147483648}
    fmt.Println(v4, v36)
    var v40 uint64 = (v36 ^ v36)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}

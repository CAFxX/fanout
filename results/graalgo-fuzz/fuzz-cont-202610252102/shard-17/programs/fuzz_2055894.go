package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
func genFn0(p0 uint16) int {
    p0 = (p0 >> ((42) & 63))
    for fi6 := 0; fi6 < 10; fi6++ {
        switch "" {
        case "a":
            if (1e10 == 1.5) {
                switch ("" + "\n\t") {
                case "":
                    v7 := []int{int(p0), fi6, -9223372036854775808}
                    _ = v7
                case "\n\t":
                    v8 := GenS1{F0: 42, F1: 4294967295}
                    _ = v8
                case "hello":
                    v9 := []int{fi6, fi6, fi6, 32767}
                    _ = v9
                }
                var v10 uint16 = 255
                _ = v10
            }
        }
        if true {
            v11 := []int{(-(fi6)), 128, (fi6 - 32767)}
            for fi12 := 0; fi12 < 9; fi12++ {
                if (988 >= 18446744073709551615) {
                    p0 -= p0
                    if ("" >= "\n\t") {
                        v13 := GenS1{F0: 1, F1: 1}
                        p0 = 0
                        var v14 uint32 = 745
                        _ = v13
                        _ = v14
                    } else {
                        if false {
                            if len(v11) > 0 {
                                v11[0] = 32768
                            }
                            var v15 int64 = 9223372036854775807
                            _ = v15
                        }
                    }
                } else {
                    var v16 int16 = 0
                    _ = v16
                }
                var v17 uint8 = 0
                _ = fi12
                _ = v17
            }
            switch (^(fi6)) {
            case 32767:
                p0 += (p0 % ((341) | 1))
            case 0:
                if (-2147483648 != -1) {
                    var v18 int = (fi6 - fi6)
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 14; fi19++ {
                        var v20 uint = 398
                        if (0.0 > ((0.1 * 1.5) / 1.5)) {
                            if false {
                                var v21 string = "foo bar"
                                switch "foo bar" {
                                case "foo bar":
                                    p0 = (^(p0))
                                case "a":
                                    var v22 string = ""
                                    _ = v22
                                }
                                _ = v21
                            }
                        }
                        _ = fi19
                        _ = v20
                    }
                }
            }
            _ = v11
        } else {
            for fi23 := 0; fi23 < 14; fi23++ {
                fmt.Println(p0, fi6)
                _ = fi23
            }
        }
        if true {
            p0 = p0
            fmt.Println(p0)
            switch "a" {
            case "":
                p0 = 375
            case "foo bar":
                fmt.Println(fi6)
            }
        } else {
            fmt.Println(p0, fi6)
            var v24 uint16 = 819
            fmt.Println(v24, p0)
            _ = v24
        }
        _ = fi6
    }
    _ = p0
    return -9223372036854775808
}
func genFn1(p0 bool) uint32 {
    p0 = ((1000 == -369) || p0)
    v25 := GenS1{F0: -1, F1: 4294967295}
    _ = p0
    _ = v25
    return 1
}

func main() {
    var v1 int8 = -128
    var v2 bool = true
    v3 := make(map[int]string)
    v4 := make(map[string]bool)
    var v5 uint8 = 202
    v1 = 1
    if (v2 || (-1 <= 9223372036854775807)) {
        v26 := []int{4294967295}
        _ = v26
    }
    v1 -= v1
    var v27 uint = 18446744073709551615
    switch 32768 {
    case -2147483648:
        fmt.Println(v5)
    case -9223372036854775808:
        if v2 {
            var v28 uint32 = 993
            if (!(-2.25 >= (0.1 * (((((-2.25 + 3.14159) * float32(3.14159)) / 0.1) * 3.14159) - ((-2.25 + (1e10 * (1e10 / 1e10))) / (0.0 + -2.25)))))) {
                var v29 float64 = 1.5
                for fi30 := 0; fi30 < 8; fi30++ {
                    v31 := GenS1{F0: 42, F1: uint32(v27)}
                    v2 = v2
                    _ = fi30
                    _ = v31
                }
                _ = v29
            }
            v28 = v28
            _ = v28
        }
    case -1:
        fmt.Println(v1)
    }
    v32 := []int{-9223372036854775808, genFn0(747)}
    if ((false || (true || false)) && (!v2)) {
        var v33 bool = (true || (665 >= 0))
        _ = v33
    } else {
        if (!true) {
            var v34 int = 4294967295
            _ = v34
        }
        delete(v3, 128)
    }
    var v35 float64 = float64((1.5 - -2.25))
    var v36 uint16 = 621
    fmt.Println(v36, v36, v35)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}

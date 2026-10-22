package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32 }

func main() {
    var v1 string = "\n\t"
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    v4 := GenS1{F0: 42, F1: 4294967295, F2: 4294967295}
    fmt.Println(v1)
    switch ("" + ("\n\t" + v1)) {
    case "":
        switch v1 {
        case "":
            v5 := GenS1{F0: 0, F1: 65535, F2: 170}
            _ = v5
        case "\n\t":
            if (!true) {
                var v6 int = 457
                _ = v6
            }
        case "foo bar":
            for fi7 := 0; fi7 < 10; fi7++ {
                for fi8 := 0; fi8 < 7; fi8++ {
                    delete(v3, fi8)
                    _ = fi8
                }
                _ = fi7
            }
        }
    case "foo bar":
        fmt.Println(v1)
    default:
        v9 := GenS1{F0: 42, F1: 853, F2: 2147483647}
        _ = v9
    }
    fmt.Println(v1)
    var v10 uint64 = 18446744073709551615
    v1 = "hello"
    v11 := GenS2{F0: -2147483648, F1: 127}
    switch ("hello" + v1) {
    case "hello":
        v12 := GenS2{F0: -2147483648, F1: -846}
        _ = v12
    case "a":
        switch "\n\t" {
        case "a":
            var v13 float32 = 0.0
            _ = v13
        case "":
            for fi14 := 0; fi14 < 5; fi14++ {
                fmt.Println(fi14, v10, v1)
                var v15 uint8 = 42
                if ((^(v15)) != 128) {
                    for fi16 := 0; fi16 < 10; fi16++ {
                        fmt.Println(fi14, v15)
                        var v17 string = v1
                        _ = fi16
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 8; fi18++ {
                        if ((!((fi14 >> ((2) & 63)) <= (fi18 + fi18))) && (!(147 < 2))) {
                            switch ((v1 + "hello") + (v1 + v1)) {
                            case "foo bar":
                                v19 := GenS3{F0: 62}
                                _ = v19
                            case "\n\t":
                                var v20 uint32 = 42
                                _ = v20
                            case "hello":
                                var v21 bool = (1 >= -1)
                                _ = v21
                            }
                            switch "hello" {
                            case "\n\t":
                                var v22 uint = 18446744073709551615
                                _ = v22
                            case "":
                                var v23 uint = 2
                                _ = v23
                            }
                            v10 = v10
                        }
                        _ = fi18
                    }
                    for fi24 := 0; fi24 < 15; fi24++ {
                        v25 := v3[-1]
                        _ = fi24
                        _ = v25
                    }
                } else {
                    v10 -= 895
                }
                _ = fi14
                _ = v15
            }
        case "\n\t":
            if true {
                var v26 uint64 = (^(v10))
                _ = v26
            } else {
                if (false && true) {
                    v3[-9223372036854775808] = 42
                    var v27 uint = 0
                    _ = v27
                } else {
                    switch 757 {
                    case 128:
                        for fi28 := 0; fi28 < 19; fi28++ {
                            var v29 int64 = 9223372036854775807
                            for fi30 := 0; fi30 < 18; fi30++ {
                                if (817 < 0) {
                                    switch v1 {
                                    case "":
                                        v1 = v1
                                    case "foo bar":
                                        var v31 int8 = 97
                                        _ = v31
                                    case "a":
                                        v2[2] = true
                                    }
                                }
                                var v32 float32 = ((float32(1.5) + 1e10) - 1e10)
                                _ = fi30
                                _ = v32
                            }
                            v29 -= v29
                            _ = fi28
                            _ = v29
                        }
                    case 32767:
                        switch v1 {
                        case "\n\t":
                            if false {
                                v33 := []int{32768, 32767}
                                _ = v33
                            }
                        }
                    case 9223372036854775807:
                        if (true || (2 <= 255)) {
                            v1 = v1
                        }
                    default:
                        v10 -= 0
                    }
                    v10 = v10
                    v1 = v1
                }
            }
        }
    }
    if false {
        var v34 int = -691
        v35 := GenS4{F0: 42}
        if (((v1 + "foo bar") + "a") >= ((((v1 + ("" + v1)) + v1) + ("foo bar" + "\n\t")) + "foo bar")) {
            var v36 float64 = (1e10 * 3.14159)
            if (!(!false)) {
                v37 := []int{(v34 & (v34 | (v34 - v34))), -1, -34}
                v38 := GenS3{F0: 100}
                _ = v37
                _ = v38
            } else {
                if ((3.14159 >= 1e10) && ((v34 & (v34 << ((18446744073709551615) & 63))) >= v34)) {
                    v34 -= v34
                }
                var v39 uint32 = 0
                v40 := v2[v39]
                _ = v39
                _ = v40
            }
            _ = v36
        } else {
            v41 := GenS2{F0: -976, F1: -157}
            v10 = v10
            _ = v41
        }
        _ = v34
        _ = v35
    }
    _ = len(v2)
    if true {
        fmt.Println(v10, v1)
        if (969 <= 432) {
            var v42 uint32 = 0
            if (0 > 386) {
                var v43 int16 = 1
                v1 = v1
                v10 = v10
                _ = v43
            }
            fmt.Println(v42, v1)
            _ = v42
        } else {
            v44 := GenS1{F0: 333, F1: 187, F2: 2147483647}
            v45 := GenS4{F0: -1}
            _ = v44
            _ = v45
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}

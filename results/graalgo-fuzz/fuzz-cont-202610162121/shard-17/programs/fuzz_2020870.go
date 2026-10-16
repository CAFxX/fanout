package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 int32; F2 uint64 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 uint; F1 uint8; F2 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int; F1 int16; F2 uint8 }
type GenS8 struct { F0 uint16; F1 string }
func genFn0(p0 uint8) int16 {
    v5 := []int{-9223372036854775808}
    p0 = 255
    var v6 uint = 188
    _ = p0
    _ = v5
    _ = v6
    return 32767
}

func main() {
    var v1 int = -1
    v2 := make(map[int]string)
    var v3 string = "a"
    v4 := make(map[string]bool)
    switch "hello" {
    case "a":
        v1 = v1
    case "hello":
        v7 := GenS1{F0: 1}
        _ = v7
    case "":
        v8 := []int{-1, v1}
        _ = v8
    default:
        fmt.Println(v3, v1)
    }
    delete(v4, ("a" + ("\n\t" + v3)))
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 float32 = 0.1
        switch v3 {
        case "a":
            v1 += v1
        case "\n\t":
            var v11 uint32 = 4294967295
            _ = v11
        }
        if false {
            switch v1 {
            case -12:
                fmt.Println(v3, v1, v10)
            case -2147483648:
                var v12 int64 = -9223372036854775808
                _ = v12
            case 50:
                v10 = v10
            }
            switch v3 {
            case "hello":
                v13 := GenS2{F0: 306, F1: 668, F2: 739}
                _ = v13
            }
        } else {
            switch "foo bar" {
            case "\n\t":
                fmt.Println(fi9, v1, fi9)
            case "hello":
                v14 := GenS3{F0: 42, F1: 128}
                _ = v14
            }
            switch -615 {
            case 30:
                if (1 != 0) {
                    switch (v3 + v3) {
                    case "\n\t":
                        v15 := GenS4{F0: 4294967295, F1: 127, F2: 4294967295}
                        _ = v15
                    }
                } else {
                    switch (v1 - v1) {
                    case 0:
                        var v16 uint = 65535
                        _ = v16
                    case 2147483647:
                        var v17 int16 = -152
                        _ = v17
                    }
                    v1 -= v1
                    var v18 bool = (0 != 1)
                    _ = v18
                }
            case 290:
                switch fi9 {
                case 9223372036854775807:
                    v2[-9223372036854775808] = "a"
                case -38:
                    if false {
                        var v19 float64 = (3.14159 / ((((0.0 / ((-2.25 * 1.5) + 1e10)) * float64(0.0)) - ((1e10 + (float64(v10) / (0.0 + ((1e10 * (-2.25 + (3.14159 * 1.5))) / (0.1 * 0.1))))) * -2.25)) + 3.14159))
                        v20 := GenS4{F0: 4294967295, F1: 78, F2: 607}
                        fmt.Println(v10, v3)
                        _ = v19
                        _ = v20
                    }
                case 0:
                    switch (fi9 >> ((1) & 63)) {
                    case 9223372036854775807:
                        if true {
                            v21 := GenS5{F0: -128}
                            switch v3 {
                            case "\n\t":
                                v22 := []int{fi9, (fi9 * v1)}
                                _ = v22
                            case "a":
                                _ = len(v4)
                            case "foo bar":
                                switch v1 {
                                case 487:
                                    fmt.Println(v3, v3, v3)
                                case 42:
                                    v4[((v3 + v3) + "foo bar")] = true
                                }
                            }
                            _ = v21
                        }
                    }
                }
            default:
                if false {
                    var v23 uint16 = 42
                    _ = v23
                }
            }
            for fi24 := 0; fi24 < 14; fi24++ {
                v25 := GenS3{F0: -339, F1: 21}
                if (-991 > 42) {
                    v1 -= v1
                    delete(v2, 4294967295)
                    v3 = v3
                    switch (v3 + v3) {
                    case "a":
                        v3 = v3
                    case "\n\t":
                        var v26 string = v3
                        _ = v26
                    }
                }
                v27 := GenS4{F0: 18446744073709551615, F1: 2, F2: 742}
                _ = fi24
                _ = v25
                _ = v27
            }
        }
        _ = fi9
        _ = v10
    }
    v3 = v3
    switch (v1 << ((607) & 63)) {
    case 32767:
        delete(v4, (v3 + v3))
    case 332:
        var v28 int32 = -2147483648
        _ = v28
    default:
        v29 := GenS1{F0: 32767}
        _ = v29
    }
    v30 := GenS4{F0: 18446744073709551615, F1: 84, F2: 255}
    v31 := GenS6{F0: 934}
    _ = len(v4)
    v32 := v4["hello"]
    var v33 int = (v1 / ((v1) | 1))
    for fi34 := 0; fi34 < 9; fi34++ {
        for fi35 := 0; fi35 < 20; fi35++ {
            v36 := GenS5{F0: -33}
            if (42 < 1) {
                var v37 int32 = -277
                _ = v37
            }
            if v32 {
                switch v33 {
                case -24:
                    delete(v2, (v33 + -1))
                }
                fmt.Println(v33)
                if v32 {
                    v3 = "a"
                    v38 := GenS7{F0: (v1 & (v33 >> ((4294967295) & 63))), F1: -694, F2: 122}
                    if false {
                        var v39 uint16 = 42
                        _ = v39
                    } else {
                        if v32 {
                            if v32 {
                                v40 := GenS2{F0: -428, F1: 1, F2: 1}
                                _ = v40
                            }
                        }
                        v41 := GenS8{F0: 42, F1: (v3 + "a")}
                        _ = v41
                    }
                    _ = v38
                }
            }
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    if ((v33 >> ((42) & 63)) > v1) {
        v1 += (v1 | (v1 * (v1 - 127)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}

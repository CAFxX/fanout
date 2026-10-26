package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 int32; F2 uint8 }
func genFn0() uint16 {
    v5 := GenS1{F0: 42, F1: 4294967295}
    var v6 float64 = 1e10
    v7 := []int{32768, 1}
    fmt.Println(v6)
    _ = v5
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 bool = (702 == 32767)
    var v2 uint8 = 104
    var v3 uint = uint(v2)
    var v4 string = ("hello" + "foo bar")
    fmt.Println(v1)
    for fi8 := 0; fi8 < 2; fi8++ {
        for fi9 := 0; fi9 < 10; fi9++ {
            switch "\n\t" {
            case "a":
                switch v4 {
                case "":
                    fmt.Println(v2, v2, v2)
                case "foo bar":
                    var v10 int8 = 100
                    _ = v10
                case "\n\t":
                    if (v2 >= (v2 % ((v2) | 1))) {
                        switch fi8 {
                        case -9223372036854775808:
                            v4 = "\n\t"
                        default:
                            fmt.Println(v2)
                        }
                        var v11 uint8 = (v2 - v2)
                        if v1 {
                            fmt.Println(fi8, v3, v2)
                            switch "" {
                            case "a":
                                if true {
                                    var v12 uint = 0
                                    v11 = 177
                                    _ = v12
                                } else {
                                    v13 := GenS1{F0: (v11 ^ (^(v2))), F1: (fi9 | -9223372036854775808)}
                                    var v14 uint8 = v2
                                    _ = v13
                                    _ = v14
                                }
                                var v15 uint32 = 4294967295
                                _ = v15
                            case "foo bar":
                                if (v1 || v1) {
                                    fmt.Println(v3)
                                    var v16 uint32 = 42
                                    v17 := GenS1{F0: 244, F1: fi9}
                                    _ = v16
                                    _ = v17
                                }
                            }
                            v18 := []int{}
                            fmt.Println(v11, v4, fi9)
                            if v1 {
                                v2 = (^(v2))
                                if v1 {
                                    v2 = v11
                                    v19 := GenS1{F0: v11, F1: fi9}
                                    fmt.Println(v3, fi8)
                                    _ = v19
                                }
                            } else {
                                var v20 int8 = 1
                                _ = v20
                            }
                            v2 = v2
                            v18 = append(v18, 212)
                            _ = v18
                        } else {
                            var v21 float32 = 1e10
                            _ = v21
                        }
                        v22 := []int{(^(fi9)), 42, fi9, fi8}
                        var v23 int = fi9
                        _ = v11
                        _ = v22
                        _ = v23
                    } else {
                        v4 = v4
                        switch (v4 + v4) {
                        case "foo bar":
                            v24 := GenS1{F0: (^(v2)), F1: fi8}
                            _ = v24
                        case "\n\t":
                            switch fi9 {
                            case 128:
                                var v25 int32 = 1
                                _ = v25
                            case 0:
                                fmt.Println(v1)
                            case 2147483647:
                                var v26 float64 = 1.5
                                _ = v26
                            default:
                                var v27 uint64 = 0
                                _ = v27
                            }
                        default:
                            var v28 int32 = 42
                            _ = v28
                        }
                        v3 = (v3 / (((^(v3))) | 1))
                    }
                }
            default:
                v29 := []int{(^(fi8))}
                _ = v29
            }
            v1 = v1
            if (!v1) {
                fmt.Println(v1)
                v1 = false
                var v30 uint16 = 0
                _ = v30
            }
            _ = fi9
        }
        v31 := []int{fi8}
        _ = fi8
        _ = v31
    }
    var v32 string = "a"
    fmt.Println(v3, v32, v1)
    v2 = v2
    fmt.Println(v4, v1, v32)
    for fi33 := 0; fi33 < 14; fi33++ {
        for fi34 := 0; fi34 < 12; fi34++ {
            switch v32 {
            case "\n\t":
                v35 := []int{(fi33 / ((fi33) | 1)), (fi33 ^ (fi34 / (((fi33 / ((fi33) | 1))) | 1))), 464, -9223372036854775808}
                _ = v35
            case "hello":
                if v1 {
                    var v36 int = 2147483647
                    _ = v36
                }
            case "foo bar":
                v37 := []int{(fi34 >> (((v3 - v3)) & 63)), (fi34 * (fi33 ^ (fi34 + fi33)))}
                _ = v37
            }
            v38 := GenS2{F0: (v4 + "")}
            _ = fi34
            _ = v38
        }
        var v39 float32 = ((1.5 / -2.25) + 0.1)
        _ = fi33
        _ = v39
    }
    if v1 {
        v4 = v32
        switch 980 {
        case 1:
            switch -588 {
            case 32768:
                switch 32767 {
                case -1:
                    v40 := []int{-40, 128, 411}
                    _ = v40
                case 1:
                    var v41 float64 = ((0.1 / 1e10) / 1.0)
                    _ = v41
                }
            case 1:
                v32 = v32
            default:
                v42 := GenS1{F0: (v2 | (v2 << ((0) & 63))), F1: 127}
                _ = v42
            }
        }
    } else {
        fmt.Println(v32)
        if true {
            var v43 uint16 = 65535
            var v44 int16 = 0
            _ = v43
            _ = v44
        } else {
            v4 = "a"
            v45 := GenS3{F0: 32767, F1: 42, F2: (^(v2))}
            _ = v45
        }
        fmt.Println(v2, v4, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
}

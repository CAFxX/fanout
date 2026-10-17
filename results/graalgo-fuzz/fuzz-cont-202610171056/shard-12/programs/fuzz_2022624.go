package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint8; F1 uint8; F2 int16 }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 float64 = 1.5
    v2 := make(map[string]int)
    var v3 int8 = -128
    v1 = v1
    switch "a" {
    case "\n\t":
        v4 := GenS1{F0: 1}
        _ = v4
    case "foo bar":
        fmt.Println(v1)
    }
    v3 += (v3 >> ((42) & 63))
    for fi5 := 0; fi5 < 14; fi5++ {
        v6 := GenS2{F0: (false || ((2 >= 2) || (18446744073709551615 > 324))), F1: 263, F2: 4294967295}
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        v3 -= -30
        _ = fi7
    }
    var v8 uint32 = 1
    var v9 float32 = 0.1
    var v10 string = ""
    v10 = ((v10 + "foo bar") + v10)
    fmt.Println(v8, v1)
    if false {
        switch 128 {
        case 9223372036854775807:
            switch -328 {
            case 128:
                for fi11 := 0; fi11 < 11; fi11++ {
                    var v12 int = (^(fi11))
                    for fi13 := 0; fi13 < 6; fi13++ {
                        var v14 uint = 2
                        if ((v14 - 4294967295) < v14) {
                            v15 := []int{fi13, (^(v12))}
                            fmt.Println(v1, fi13)
                            v16 := GenS1{F0: v14}
                            _ = v15
                            _ = v16
                        }
                        _ = fi13
                        _ = v14
                    }
                    _ = fi11
                    _ = v12
                }
            case 4294967295:
                var v17 string = "\n\t"
                _ = v17
            case 2147483647:
                var v18 int32 = -2147483648
                _ = v18
            default:
                v9 -= (float32(-2.25) + v9)
            }
        case 1:
            v8 -= 2147483648
        case 32768:
            v8 = (v8 << ((0) & 63))
        default:
            for fi19 := 0; fi19 < 5; fi19++ {
                switch (fi19 << ((65535) & 63)) {
                case 1:
                    if true {
                        switch fi19 {
                        case 1:
                            v3 -= v3
                        case -760:
                            fmt.Println(v1)
                        case 128:
                            v20 := GenS1{F0: 84}
                            _ = v20
                        }
                        fmt.Println(v1, v3, v1)
                    }
                case -1:
                    for fi21 := 0; fi21 < 8; fi21++ {
                        v8 += (v8 << ((18446744073709551615) & 63))
                        v22 := v2[v10]
                        var v23 int = fi21
                        switch "\n\t" {
                        case "\n\t":
                            v24 := []int{}
                            _ = v24
                        default:
                            if (false || false) {
                                var v25 uint32 = (v8 - v8)
                                switch ((v10 + v10) + v10) {
                                case "foo bar":
                                    fmt.Println(v9)
                                case "\n\t":
                                    delete(v2, v10)
                                }
                                v10 = v10
                                _ = v25
                            } else {
                                var v26 uint16 = 255
                                _ = v26
                            }
                        }
                        if true {
                            switch v23 {
                            case 32768:
                                switch v23 {
                                case 9223372036854775807:
                                    v27 := []int{9223372036854775807, (v23 >> ((65535) & 63)), (fi19 ^ v22)}
                                    _ = v27
                                case -239:
                                    var v28 uint32 = v8
                                    _ = v28
                                default:
                                    var v29 int32 = 2147483647
                                    _ = v29
                                }
                            case -217:
                                switch "\n\t" {
                                case "a":
                                    v30 := GenS3{F0: 255, F1: 2, F2: 247}
                                    _ = v30
                                case "foo bar":
                                    switch v10 {
                                    case "":
                                        var v31 int = -798
                                        _ = v31
                                    case "\n\t":
                                        v32 := v2[v10]
                                        _ = v32
                                    case "a":
                                        v10 = v10
                                    default:
                                        if (!true) {
                                            var v33 int64 = 42
                                            switch v10 {
                                            case "\n\t":
                                                var v34 uint = 65535
                                                _ = v34
                                            case "foo bar":
                                                var v35 int32 = 2147483647
                                                _ = v35
                                            case "":
                                                v36 := []int{}
                                                delete(v2, "a")
                                                switch v10 {
                                                case "\n\t":
                                                    v23 += (v23 - 4294967295)
                                                    v22 -= v22
                                                    v22 += (-(fi21))
                                                case "a":
                                                    v3 = (v3 % ((25) | 1))
                                                }
                                                _ = v36
                                            default:
                                                v37 := []int{-1, fi19}
                                                _ = v37
                                            }
                                            var v38 float32 = 0.0
                                            _ = v33
                                            _ = v38
                                        }
                                    }
                                default:
                                    var v39 int64 = 0
                                    _ = v39
                                }
                            case 2147483647:
                                var v40 int64 = 1
                                _ = v40
                            }
                        }
                        var v41 uint = 2
                        _ = fi21
                        _ = v22
                        _ = v23
                        _ = v41
                    }
                }
                fmt.Println(v3, v9)
                _ = fi19
            }
        }
        v10 = (v10 + v10)
    } else {
        v42 := GenS2{F0: (!(!true)), F1: 1, F2: 2}
        switch 4294967295 {
        case 202:
            v43 := GenS4{F0: 1}
            _ = v43
        case -1:
            _ = len(v2)
        }
        switch v10 {
        case "foo bar":
            var v44 uint16 = 2
            _ = v44
        case "":
            if true {
                var v45 int8 = v3
                v46 := []int{2147483647, 4294967295}
                v1 = (v1 / float64(v9))
                _ = v45
                _ = v46
            }
        }
        _ = v42
    }
    v2[v10] = -307
    v3 = (v3 ^ (v3 / ((42) | 1)))
    v1 += (v1 * 0.0)
    v8 += (^(v8))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}

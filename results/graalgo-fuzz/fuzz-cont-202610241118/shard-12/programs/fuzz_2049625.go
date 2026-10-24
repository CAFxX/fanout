package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 bool; F1 int }
func genFn0(p0 int16) int16 {
    fmt.Println(p0)
    p0 -= p0
    fmt.Println(p0)
    _ = p0
    return p0
}

func main() {
    var v1 int = 42
    var v2 uint = 65535
    var v3 string = ""
    var v4 uint16 = 33
    for fi5 := 0; fi5 < 14; fi5++ {
        if ((-1 != 9223372036854775807) || true) {
            v1 = v1
            v6 := []int{(fi5 | v1)}
            var v7 float64 = (0.0 * ((-2.25 / 1.0) / (-2.25 + 3.14159)))
            _ = v6
            _ = v7
        } else {
            var v8 float64 = (((float64(0.0) + (float64(1e10) * (((1e10 / ((-2.25 * 0.1) + 1.5)) * 1.5) * 1e10))) * -2.25) * (3.14159 / 1e10))
            _ = v8
        }
        var v9 int64 = 814
        for fi10 := 0; fi10 < 20; fi10++ {
            v1 = (fi5 - 270)
            var v11 bool = (("\n\t" <= ((v3 + (v3 + ("foo bar" + v3))) + ("a" + ("a" + v3)))) || true)
            if v11 {
                fmt.Println(v2)
                switch "" {
                case "foo bar":
                    fmt.Println(fi5, v3, v1)
                case "":
                    switch ("hello" + ((v3 + "foo bar") + "a")) {
                    case "hello":
                        fmt.Println(fi5)
                    case "\n\t":
                        v12 := []int{}
                        _ = v12
                    case "a":
                        var v13 int = 0
                        var v14 uint64 = 18446744073709551615
                        _ = v13
                        _ = v14
                    }
                }
                v15 := []int{}
                fmt.Println(v2, v1)
                _ = v15
            }
            _ = fi10
            _ = v11
        }
        _ = fi5
        _ = v9
    }
    var v16 float64 = 3.14159
    v4 += v4
    var v17 uint8 = 2
    v18 := []int{(v1 - v1), v1}
    v16 += v16
    v19 := GenS1{F0: v1, F1: v4, F2: -1}
    fmt.Println(v16)
    switch v1 {
    case 393:
        v3 = v3
    case 9223372036854775807:
        if (32767 >= 32767) {
            v20 := GenS1{F0: 9223372036854775807, F1: 65535, F2: 0}
            var v21 uint32 = 42
            fmt.Println(v3, v17, v3)
            _ = v20
            _ = v21
        } else {
            if (("a" <= v3) && (2 != 2)) {
                for fi22 := 0; fi22 < 3; fi22++ {
                    for fi23 := 0; fi23 < 12; fi23++ {
                        v4 += (v4 | v4)
                        switch "" {
                        case "\n\t":
                            switch v3 {
                            case "a":
                                v24 := GenS2{F0: 65535, F1: -74}
                                v3 = v3
                                fmt.Println(v16, v17, v1)
                                _ = v24
                            case "foo bar":
                                if true {
                                    v4 = uint16(v2)
                                    v3 = (((("hello" + v3) + v3) + "a") + "a")
                                    v2 -= (^(v2))
                                    if len(v18) >= 2 {
                                        v18 = v18[1:len(v18)]
                                    }
                                } else {
                                    fmt.Println(fi22)
                                }
                            case "\n\t":
                                v25 := GenS1{F0: fi23, F1: 357, F2: 1}
                                _ = v25
                            default:
                                if (4294967295 <= 1) {
                                    var v26 uint8 = v17
                                    v27 := GenS3{F0: ("" + "\n\t")}
                                    if len(v18) >= 2 {
                                        v18 = v18[1:len(v18)]
                                    }
                                    _ = v26
                                    _ = v27
                                } else {
                                    v28 := GenS4{F0: v1, F1: 18446744073709551615}
                                    v1 += v1
                                    v29 := GenS3{F0: (v3 + "hello")}
                                    fmt.Println(v2)
                                    _ = v28
                                    _ = v29
                                }
                            }
                        case "hello":
                            var v30 uint8 = v17
                            _ = v30
                        case "foo bar":
                            v31 := GenS2{F0: uint(v4), F1: 42}
                            _ = v31
                        default:
                            var v32 uint32 = 2147483647
                            _ = v32
                        }
                        v33 := GenS1{F0: (-(fi22)), F1: 848, F2: 9223372036854775807}
                        _ = fi23
                        _ = v33
                    }
                    if true {
                        v18 = append(v18, (fi22 << ((4294967295) & 63)))
                    } else {
                        for fi34 := 0; fi34 < 6; fi34++ {
                            v35 := GenS5{F0: true, F1: (v1 | 777)}
                            switch ("" + v3) {
                            case "hello":
                                fmt.Println(v17, fi22)
                            case "\n\t":
                                switch ("foo bar" + v3) {
                                case "":
                                    fmt.Println(v1, fi34, fi22)
                                case "\n\t":
                                    v36 := GenS1{F0: fi22, F1: 499, F2: 0}
                                    _ = v36
                                case "foo bar":
                                    var v37 int16 = 684
                                    _ = v37
                                default:
                                    var v38 float64 = v16
                                    _ = v38
                                }
                            case "foo bar":
                                var v39 uint32 = 2147483647
                                _ = v39
                            }
                            v1 += -1
                            _ = fi34
                            _ = v35
                        }
                    }
                    var v40 int32 = -1
                    _ = fi22
                    _ = v40
                }
                var v41 float64 = 0.0
                _ = v41
            } else {
                var v42 int16 = 292
                _ = v42
            }
            v3 = ""
        }
    case 1:
        switch (v3 + v3) {
        case "hello":
            v1 -= (v1 & v1)
        }
    default:
        v17 += (v17 / (((v17 & 2)) | 1))
    }
    for fi43 := 0; fi43 < 10; fi43++ {
        switch fi43 {
        case 32768:
            v44 := GenS1{F0: 9223372036854775807, F1: 0, F2: -1}
            _ = v44
        case -2147483648:
            switch -9223372036854775808 {
            case -9223372036854775808:
                var v45 string = v3
                _ = v45
            case -1:
                var v46 uint16 = v4
                _ = v46
            }
        case 287:
            fmt.Println(v1)
        }
        var v47 float64 = ((((v16 + 3.14159) - ((3.14159 * (v16 * (1e10 + (1.5 / float64(1.5))))) + (0.0 + float64((3.14159 + ((0.0 / 1.5) + -2.25)))))) / v16) + 1.5)
        var v48 string = ((((("\n\t" + v3) + "foo bar") + (v3 + v3)) + "a") + v3)
        _ = fi43
        _ = v47
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}

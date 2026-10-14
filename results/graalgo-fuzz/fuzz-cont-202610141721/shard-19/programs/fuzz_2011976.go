package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int32; F1 uint8; F2 bool }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 int32; F1 int8; F2 int8 }

func main() {
    var v1 bool = true
    var v2 float32 = (0.0 + 0.1)
    v3 := []int{2147483647, 128, 32768, 127}
    var v4 uint = 65535
    v3 = append(v3, 9223372036854775807)
    var v5 int64 = 591
    for fi6 := 0; fi6 < 6; fi6++ {
        _ = len(v3)
        var v7 int8 = 127
        v8 := GenS1{F0: ("\n\t" + "hello")}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint16 = 2
    switch 1 {
    case 42:
        if (!(42 == 967)) {
            for fi10 := 0; fi10 < 7; fi10++ {
                if v1 {
                    fmt.Println(v1)
                    v11 := GenS2{F0: ("hello" + ""), F1: -128}
                    if ((v5 | -1) == (v5 ^ 42)) {
                        v12 := GenS3{F0: 42, F1: 127, F2: false}
                        _ = len(v3)
                        _ = v12
                    }
                    _ = v11
                }
                switch fi10 {
                case 127:
                    if v1 {
                        fmt.Println(v4, fi10)
                        v2 = v2
                    } else {
                        var v13 float64 = 1e10
                        _ = v13
                    }
                case -299:
                    switch "hello" {
                    case "hello":
                        fmt.Println(v4, fi10)
                    case "":
                        var v14 uint64 = 51
                        _ = v14
                    default:
                        var v15 uint = 4294967295
                        _ = v15
                    }
                case 32768:
                    fmt.Println(v9, fi10, fi10)
                default:
                    switch (fi10 * (fi10 >> ((v4) & 63))) {
                    case -9223372036854775808:
                        if v1 {
                            var v16 uint = (^(v4))
                            for fi17 := 0; fi17 < 16; fi17++ {
                                var v18 bool = (v1 || true)
                                v19 := GenS1{F0: "\n\t"}
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                            _ = v16
                        } else {
                            var v20 string = ("\n\t" + ("hello" + ("hello" + "foo bar")))
                            for fi21 := 0; fi21 < 10; fi21++ {
                                if v1 {
                                    if ((v2 * v2) == v2) {
                                        var v22 string = ((v20 + (v20 + "hello")) + v20)
                                        _ = v22
                                    }
                                    if v1 {
                                        switch v20 {
                                        case "foo bar":
                                            if (3.14159 < 1e10) {
                                                v2 += ((-2.25 * v2) * v2)
                                            }
                                        default:
                                            if (!(!true)) {
                                                var v23 uint16 = (v9 / (((v9 * v9)) | 1))
                                                _ = v23
                                            } else {
                                                fmt.Println(fi10, v2)
                                                switch "" {
                                                case "a":
                                                    var v24 uint = 18446744073709551615
                                                    _ = v24
                                                case "\n\t":
                                                    fmt.Println(v2)
                                                case "foo bar":
                                                    if len(v3) > 0 {
                                                        v3[0] = (fi10 | fi10)
                                                    }
                                                default:
                                                    v25 := GenS4{F0: -32768, F1: (v9 ^ (v9 | v9))}
                                                    _ = v25
                                                }
                                                var v26 uint8 = 197
                                                _ = v26
                                            }
                                        }
                                        v20 = v20
                                        if v1 {
                                            fmt.Println(fi21, v2)
                                        }
                                    } else {
                                        if v1 {
                                            fmt.Println(v1)
                                            if (!v1) {
                                                var v27 uint64 = 1
                                                v28 := GenS2{F0: "foo bar", F1: 0}
                                                if (31 == -1) {
                                                    v5 = (-(v5))
                                                }
                                                switch (fi10 * fi10) {
                                                case 32768:
                                                    v2 = (v2 / v2)
                                                case -2147483648:
                                                    v20 = v20
                                                }
                                                _ = len(v3)
                                                _ = v27
                                                _ = v28
                                            }
                                            switch "hello" {
                                            case "hello":
                                                if ((v4 << (((v4 << ((v4) & 63))) & 63)) > v4) {
                                                    var v29 uint64 = 917
                                                    v29 += (v29 >> (((^(v4))) & 63))
                                                    _ = v29
                                                } else {
                                                    var v30 bool = true
                                                    var v31 uint8 = uint8(v4)
                                                    v30 = v1
                                                    fmt.Println(fi10, v5)
                                                    _ = v30
                                                    _ = v31
                                                }
                                            case "foo bar":
                                                var v32 float64 = (0.0 * (float64(1e10) + 0.0))
                                                _ = v32
                                            default:
                                                _ = len(v3)
                                            }
                                        } else {
                                            var v33 int = (fi21 % ((fi21) | 1))
                                            _ = v33
                                        }
                                        switch (v20 + (v20 + ((v20 + "") + ((v20 + v20) + ("a" + v20))))) {
                                        case "hello":
                                            var v34 float64 = 3.14159
                                            _ = v34
                                        case "":
                                            switch ("foo bar" + v20) {
                                            case "foo bar":
                                                fmt.Println(v20, fi10, v20)
                                            default:
                                                v20 = "a"
                                            }
                                        case "foo bar":
                                            _ = len(v3)
                                        default:
                                            if len(v3) > 0 {
                                                v3[0] = -2147483648
                                            }
                                        }
                                    }
                                    v20 = v20
                                } else {
                                    var v35 uint32 = uint32(v9)
                                    fmt.Println(v35)
                                    switch v20 {
                                    case "a":
                                        v36 := GenS5{F0: 161, F1: ((v1 || v1) && ((v9 & (^(v9))) > (^(v9))))}
                                        _ = v36
                                    case "foo bar":
                                        if ((("a" + v20) <= (v20 + v20)) && ((v35 | (v35 / ((4294967295) | 1))) != 1)) {
                                            v20 = "\n\t"
                                            if (42 != -1) {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                var v37 float64 = 0.0
                                                _ = v37
                                            }
                                            var v38 int32 = 42
                                            _ = v38
                                        } else {
                                            var v39 int = (fi21 ^ 127)
                                            var v40 uint32 = v35
                                            if len(v3) > 0 {
                                                v3[0] = (-(fi21))
                                            }
                                            _ = v39
                                            _ = v40
                                        }
                                    case "hello":
                                        v1 = false
                                    }
                                    _ = v35
                                }
                                _ = len(v3)
                                _ = fi21
                            }
                            _ = v20
                        }
                    case 1:
                        for fi41 := 0; fi41 < 4; fi41++ {
                            v42 := GenS3{F0: -2147483648, F1: 127, F2: (true || (1e10 != v2))}
                            _ = fi41
                            _ = v42
                        }
                    case 9223372036854775807:
                        var v43 float32 = (-2.25 + ((v2 - v2) + (float32(3.14159) * 0.0)))
                        _ = v43
                    }
                }
                _ = len(v3)
                _ = fi10
            }
            var v44 float32 = (v2 - (1e10 + (v2 - float32((float64(3.14159) - 1e10)))))
            _ = v44
        }
    case -9223372036854775808:
        _ = len(v3)
    default:
        if v1 {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if v1 {
                var v45 int32 = -1
                fmt.Println(v2, v2, v5)
                v4 -= v4
                _ = v45
            } else {
                v46 := GenS3{F0: 2147483647, F1: 236, F2: (int(v5) > 506)}
                for fi47 := 0; fi47 < 9; fi47++ {
                    if true {
                        var v48 uint8 = 29
                        v5 += v5
                        _ = v48
                    }
                    _ = fi47
                }
                _ = v46
            }
            fmt.Println(v1, v4, v4)
        }
    }
    if v1 {
        v1 = (77 == 100)
    } else {
        var v49 int16 = 1
        v4 = uint(v9)
        _ = v49
    }
    var v50 int8 = 95
    var v51 int8 = 0
    v52 := GenS1{F0: "foo bar"}
    if v1 {
        var v53 uint64 = 0
        _ = v53
    } else {
        v54 := GenS6{F0: 42, F1: -125, F2: (v51 & (^(v51)))}
        if len(v3) > 0 {
            v3[0] = int(v4)
        }
        _ = v54
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v9
}

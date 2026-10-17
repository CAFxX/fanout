package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int8 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint32; F1 int32; F2 string }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int8; F1 uint32; F2 uint32 }

func main() {
    var v1 int64 = -659
    var v2 int8 = 91
    var v3 uint64 = 700
    v4 := []string{}
    var v5 uint16 = 32767
    switch "a" {
    case "a":
        switch "hello" {
        case "\n\t":
            if (((((-2.25 + -2.25) * (0.1 - ((1e10 - 1.5) / -2.25))) + 0.0) < float32(float64(1.5))) && (42 != (v1 | (v1 | (v1 ^ -9223372036854775808))))) {
                v2 = 0
            } else {
                v2 = (v2 % (((v2 | v2)) | 1))
                var v6 uint = 65535
                _ = v6
            }
        default:
            v7 := GenS1{F0: (v2 & (v2 ^ v2)), F1: "foo bar", F2: v2}
            _ = v7
        }
    default:
        var v8 int32 = -2147483648
        _ = v8
    }
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    var v9 uint16 = (v5 | (v5 * v5))
    var v10 int16 = 0
    if len(v4) > 0 {
        v4[0] = ""
    }
    fmt.Println(v1, v9, v5)
    var v11 string = "\n\t"
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        var v13 bool = (1 != 2147483648)
        var v14 string = ((("hello" + v11) + "a") + v11)
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 int = 9223372036854775807
    v16 := GenS1{F0: (v2 / ((v2) | 1)), F1: ((v11 + "a") + "hello"), F2: v2}
    for fi17 := 0; fi17 < 4; fi17++ {
        switch ("foo bar" + v11) {
        case "":
            v18 := GenS1{F0: (v2 % ((v2) | 1)), F1: (v11 + "a"), F2: 127}
            _ = v18
        case "a":
            var v19 int16 = -1
            _ = v19
        case "foo bar":
            if (v10 >= (v10 << ((65535) & 63))) {
                v20 := GenS2{F0: (v15 - 9223372036854775807), F1: v3}
                switch v11 {
                case "\n\t":
                    v21 := GenS3{F0: 42, F1: -276, F2: v11}
                    _ = v21
                case "hello":
                    var v22 float32 = float32((3.14159 + (float64(-2.25) - float64((((float32(1.5) / ((((0.1 / 0.1) + -2.25) * 1e10) / 1e10)) - 3.14159) * -2.25)))))
                    _ = v22
                case "":
                    switch v11 {
                    case "\n\t":
                        for fi23 := 0; fi23 < 3; fi23++ {
                            fmt.Println(v15, fi17)
                            if (v5 >= (v9 - (^(v9)))) {
                                _ = len(v4)
                            } else {
                                if (32767 < (-(v10))) {
                                    if false {
                                        v4 = append(v4, v11)
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        var v24 uint8 = 142
                                        if (-2147483648 != 42) {
                                            v9 = (v5 << ((250) & 63))
                                            if len(v4) > 0 {
                                                v4[0] = "a"
                                            }
                                        } else {
                                            switch fi23 {
                                            case -501:
                                                var v25 int64 = (-(v1))
                                                _ = v25
                                            case -2147483648:
                                                var v26 bool = false
                                                _ = v26
                                            case 128:
                                                switch v15 {
                                                case 42:
                                                    var v27 float64 = 1e10
                                                    _ = v27
                                                }
                                            default:
                                                var v28 uint8 = v24
                                                _ = v28
                                            }
                                            if ((v1 % ((int64(v3)) | 1)) <= (v1 >> ((40) & 63))) {
                                                v10 += v10
                                                v24 = (v24 ^ v24)
                                                var v29 float64 = 0.1
                                                _ = v29
                                            } else {
                                                var v30 uint = 69
                                                var v31 bool = true
                                                var v32 uint = (v30 ^ 18446744073709551615)
                                                _ = v30
                                                _ = v31
                                                _ = v32
                                            }
                                        }
                                        if (318 == 2147483647) {
                                            if (0 != 42) {
                                                var v33 uint64 = v3
                                                v9 = 2
                                                switch (^(v15)) {
                                                case -692:
                                                    fmt.Println(v5, v33)
                                                }
                                                _ = v33
                                            } else {
                                                fmt.Println(v9)
                                                v34 := GenS1{F0: -1, F1: "hello", F2: (-(v2))}
                                                _ = v34
                                            }
                                            if (0 < v10) {
                                                var v35 int16 = (v10 | 42)
                                                _ = v35
                                            } else {
                                                fmt.Println(v2)
                                                if true {
                                                    var v36 int64 = v1
                                                    var v37 int = v15
                                                    v37 = 32768
                                                    v38 := GenS2{F0: 46, F1: (v3 % ((v3) | 1))}
                                                    _ = v36
                                                    _ = v37
                                                    _ = v38
                                                } else {
                                                    v10 = -777
                                                    var v39 int32 = -780
                                                    switch (fi17 >> ((2) & 63)) {
                                                    case -61:
                                                        fmt.Println(v2, v11)
                                                    }
                                                    _ = v39
                                                }
                                            }
                                            fmt.Println(v11)
                                            if false {
                                                if len(v4) > 0 {
                                                    v4[0] = (v11 + ("\n\t" + ((v11 + ((v11 + "\n\t") + (("\n\t" + ("hello" + v11)) + "\n\t"))) + v11)))
                                                }
                                                v15 = fi23
                                            }
                                            var v40 int64 = (v1 + 9223372036854775807)
                                            var v41 uint16 = v9
                                            _ = v40
                                            _ = v41
                                        } else {
                                            v11 = (v11 + (v11 + v11))
                                            var v42 uint64 = v3
                                            var v43 string = (v11 + (v11 + ((v11 + "") + v11)))
                                            var v44 uint16 = 647
                                            v4 = append(v4, v43)
                                            if false {
                                                v45 := GenS3{F0: 927, F1: 42, F2: (("foo bar" + v43) + "")}
                                                v44 = 206
                                                switch "a" {
                                                case "a":
                                                    var v46 uint16 = 2
                                                    v5 += uint16(v24)
                                                    _ = len(v4)
                                                    _ = v46
                                                case "foo bar":
                                                    if (0 == 42) {
                                                        v43 = v11
                                                    } else {
                                                        switch 9223372036854775807 {
                                                        case 0:
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        default:
                                                            if len(v4) > 0 {
                                                                v4[0] = v43
                                                            }
                                                        }
                                                        var v47 int = -9223372036854775808
                                                        _ = v47
                                                    }
                                                case "\n\t":
                                                    var v48 int = (fi17 + v15)
                                                    _ = v48
                                                }
                                                _ = v45
                                            }
                                            _ = v42
                                            _ = v43
                                            _ = v44
                                        }
                                        if true {
                                            if (3.14159 == 1e10) {
                                                if true {
                                                    fmt.Println(v1, v3, v3)
                                                    var v49 float64 = float64((((3.14159 / ((1e10 - (-2.25 - 3.14159)) / 1.5)) + 0.1) / 1.5))
                                                    _ = v49
                                                }
                                            } else {
                                                v24 = (v24 >> ((uint(v3)) & 63))
                                                v50 := GenS4{F0: (v1 ^ int64(v5)), F1: 2147483647}
                                                var v51 int8 = v2
                                                _ = v50
                                                _ = v51
                                            }
                                            switch 957 {
                                            case 32768:
                                                fmt.Println(v3, v1, v5)
                                            case 190:
                                                v52 := GenS3{F0: 993, F1: -1, F2: v11}
                                                _ = v52
                                            case 9223372036854775807:
                                                if (v3 >= v3) {
                                                    v53 := GenS1{F0: v2, F1: v11, F2: (v2 % ((0) | 1))}
                                                    v11 = v11
                                                    fmt.Println(v9, v24)
                                                    _ = v53
                                                } else {
                                                    var v54 int16 = v10
                                                    fmt.Println(v2, v3, v54)
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    _ = v54
                                                }
                                            default:
                                                switch v11 {
                                                case "hello":
                                                    var v55 uint64 = uint64(v5)
                                                    _ = v55
                                                }
                                            }
                                            _ = len(v4)
                                            if true {
                                                v2 -= v2
                                            }
                                        } else {
                                            if true {
                                                v56 := GenS5{F0: v15}
                                                v5 -= (v5 + v5)
                                                _ = v56
                                            }
                                            v15 += (fi17 & (^(fi23)))
                                        }
                                        _ = v24
                                    }
                                    fmt.Println(fi17)
                                    var v57 int8 = (^(v2))
                                    _ = v57
                                } else {
                                    var v58 int = 2147483647
                                    if true {
                                        fmt.Println(v2, v3, fi23)
                                    }
                                    v59 := GenS5{F0: (fi17 >> ((956) & 63))}
                                    if (0.1 >= (1.5 / ((1.5 / (((1.5 * (0.0 + 0.1)) * ((1e10 * (1.5 * 1.5)) / 1.0)) * (1e10 + 1e10))) / (1e10 + 3.14159)))) {
                                        v15 = fi23
                                    } else {
                                        v3 += (v3 % ((2) | 1))
                                        fmt.Println(v2, fi17)
                                    }
                                    _ = v58
                                    _ = v59
                                }
                                v15 = (v15 << ((1) & 63))
                            }
                            _ = fi23
                        }
                    default:
                        v60 := GenS3{F0: uint32(v3), F1: -98, F2: ""}
                        _ = v60
                    }
                }
                v2 = v2
                _ = v20
            }
        default:
            var v61 string = v11
            _ = v61
        }
        v62 := GenS6{F0: 987}
        _ = fi17
        _ = v62
    }
    switch -9223372036854775808 {
    case -611:
        var v63 uint16 = 0
        _ = v63
    case -578:
        switch v15 {
        case 0:
            v64 := GenS7{F0: v2, F1: 2147483647, F2: 1}
            _ = v64
        }
    }
    var v65 int16 = v10
    v66 := GenS4{F0: v1, F1: 677}
    var v67 uint64 = v3
    var v68 uint = 42
    v69 := GenS2{F0: (^(v15)), F1: 500}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v65:", v65)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v66
    _ = v67
    _ = v68
    _ = v69
    _ = v9
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint32; F1 string }
type GenS8 struct { F0 uint16; F1 int64 }
type GenS9 struct { F0 int; F1 uint; F2 uint64 }

func main() {
    var v1 string = (("\n\t" + "a") + "a")
    v2 := []uint32{756}
    var v3 bool = true
    for fi4 := 0; fi4 < 9; fi4++ {
        fmt.Println(fi4)
        _ = fi4
    }
    v3 = (!(!true))
    var v5 int16 = 42
    switch v1 {
    case "":
        fmt.Println(v5)
    case "hello":
        var v6 int = 32767
        _ = v6
    case "foo bar":
        for fi7 := 0; fi7 < 7; fi7++ {
            switch (fi7 * 573) {
            case 218:
                v8 := GenS1{F0: 1, F1: -128}
                _ = v8
            case 286:
                v3 = (0 > 645)
            case 836:
                v9 := GenS1{F0: 42, F1: 42}
                _ = v9
            default:
                for fi10 := 0; fi10 < 17; fi10++ {
                    v3 = (203 <= 2)
                    _ = len(v2)
                    _ = fi10
                }
            }
            if v3 {
                if true {
                    var v11 int64 = -1
                    switch "\n\t" {
                    case "hello":
                        var v12 string = ""
                        _ = v12
                    case "":
                        v3 = (60 < 100)
                    case "foo bar":
                        var v13 bool = false
                        _ = v13
                    default:
                        var v14 uint = 65535
                        _ = v14
                    }
                    for fi15 := 0; fi15 < 16; fi15++ {
                        if ((!v3) && true) {
                            var v16 int64 = int64(fi15)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v5 = (v5 << ((4294967295) & 63))
                            _ = v16
                        } else {
                            v1 = (v1 + ((v1 + "hello") + v1))
                        }
                        switch fi7 {
                        case 140:
                            var v17 int32 = 42
                            _ = v17
                        }
                        switch fi7 {
                        case 9223372036854775807:
                            if (42 > 100) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            } else {
                                var v18 uint64 = 18446744073709551615
                                v19 := GenS1{F0: -2147483648, F1: 0}
                                if v3 {
                                    v2 = append(v2, 234)
                                }
                                _ = v18
                                _ = v19
                            }
                        case 2147483647:
                            if (v3 || (1 == 902)) {
                                v3 = (!false)
                                if (v3 || (-71 <= 530)) {
                                    if false {
                                        var v20 int8 = 34
                                        var v21 uint = 294
                                        v5 += v5
                                        v3 = ((v20 ^ v20) > v20)
                                        _ = v20
                                        _ = v21
                                    } else {
                                        fmt.Println(fi15, v5)
                                        if len(v2) > 0 {
                                            v2[0] = 815
                                        }
                                    }
                                    var v22 float32 = float32(0.0)
                                    v2 = append(v2, 271)
                                    _ = v22
                                } else {
                                    v11 += v11
                                    v23 := GenS2{F0: ("hello" + "a"), F1: 55}
                                    switch (^(fi15)) {
                                    case -939:
                                        v24 := GenS3{F0: 1}
                                        _ = v24
                                    }
                                    _ = v23
                                }
                            } else {
                                var v25 int64 = (v11 | v11)
                                v26 := GenS4{F0: fi7, F1: int8(v11)}
                                _ = v25
                                _ = v26
                            }
                        default:
                            fmt.Println(v1)
                            fmt.Println(v5)
                            if (!v3) {
                                if (false || (!false)) {
                                    v5 = v5
                                    if (v3 && true) {
                                        var v27 uint64 = 42
                                        var v28 float32 = 0.0
                                        switch fi7 {
                                        case -1:
                                            v28 += 0.0
                                        case 2147483647:
                                            v11 = (v11 & (^(v11)))
                                        case 330:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            var v29 int64 = (v11 % (((v11 & (v11 % (((v11 ^ 462)) | 1)))) | 1))
                                            v3 = v3
                                            switch (fi15 | fi7) {
                                            case 4294967295:
                                                if true {
                                                    v1 = (((v1 + v1) + v1) + v1)
                                                } else {
                                                    v30 := GenS5{F0: (^(v27))}
                                                    fmt.Println(v1, v3)
                                                    v2 = append(v2, 2147483648)
                                                    _ = v30
                                                }
                                            default:
                                                v31 := GenS6{F0: 100}
                                                _ = v31
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            switch (^(fi7)) {
                                            case 0:
                                                if v3 {
                                                    var v32 uint64 = v27
                                                    fmt.Println(v27, v1, v11)
                                                    _ = v32
                                                } else {
                                                    v5 -= v5
                                                    v33 := GenS5{F0: 888}
                                                    _ = v33
                                                }
                                                _ = len(v2)
                                                if len(v2) > 0 {
                                                    v2[0] = 2147483648
                                                }
                                            default:
                                                if len(v2) > 0 {
                                                    v2[0] = 65535
                                                }
                                            }
                                            if v3 {
                                                var v34 string = "a"
                                                if ((v27 | (^(v27))) != 4294967295) {
                                                    v35 := GenS4{F0: -782, F1: -99}
                                                    fmt.Println(v1, v28)
                                                    _ = v35
                                                }
                                                _ = v34
                                            }
                                            switch fi7 {
                                            case -9223372036854775808:
                                                fmt.Println(v11)
                                            case 194:
                                                if v3 {
                                                    v29 += (v29 + (v11 % (((v11 << ((971) & 63))) | 1)))
                                                } else {
                                                    v29 = (v29 & (v11 / ((v29) | 1)))
                                                    fmt.Println(v3, v11)
                                                    if false {
                                                        if len(v2) > 0 {
                                                            v2[0] = 65535
                                                        }
                                                    } else {
                                                        var v36 int16 = (v5 + v5)
                                                        _ = v36
                                                    }
                                                }
                                            case 759:
                                                if (-1 < -7) {
                                                    v11 = v29
                                                    fmt.Println(v28, v1, v3)
                                                    if v3 {
                                                        v29 = (v29 >> ((0) & 63))
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        if v3 {
                                                            if false {
                                                                v3 = false
                                                                var v37 int64 = (-(v11))
                                                                fmt.Println(fi15)
                                                                if (0 == 2147483647) {
                                                                    _ = len(v2)
                                                                    v28 = v28
                                                                    _ = len(v2)
                                                                }
                                                                _ = v37
                                                            }
                                                            if v3 {
                                                                var v38 int32 = 888
                                                                v5 += v5
                                                                var v39 float32 = 3.14159
                                                                _ = v38
                                                                _ = v39
                                                            }
                                                            fmt.Println(v27, v3)
                                                            _ = len(v2)
                                                            var v40 uint32 = 978
                                                            _ = v40
                                                        } else {
                                                            if len(v2) > 0 {
                                                                v2[0] = 83
                                                            }
                                                            v41 := GenS1{F0: 2147483647, F1: 127}
                                                            _ = v41
                                                        }
                                                    }
                                                }
                                            }
                                            if (false || v3) {
                                                v1 = v1
                                                v1 = v1
                                            } else {
                                                switch (fi15 & 1) {
                                                case -2147483648:
                                                    if (v3 && ((^(v27)) > 248)) {
                                                        v42 := GenS7{F0: 771, F1: "a"}
                                                        var v43 int32 = 2147483647
                                                        if v3 {
                                                            var v44 bool = v3
                                                            v5 = (-(v5))
                                                            _ = v44
                                                        }
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                default:
                                                    fmt.Println(v29)
                                                }
                                            }
                                            _ = v29
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = 42
                                        }
                                        var v45 float64 = 1e10
                                        _ = v27
                                        _ = v28
                                        _ = v45
                                    }
                                    fmt.Println(v1, fi7, fi15)
                                } else {
                                    var v46 uint64 = 4294967295
                                    _ = v46
                                }
                            }
                        }
                        _ = fi15
                    }
                    _ = v11
                } else {
                    var v47 float32 = (3.14159 - 1e10)
                    fmt.Println(v5)
                    _ = v47
                }
                var v48 int32 = -976
                v49 := GenS4{F0: 0, F1: -50}
                _ = v48
                _ = v49
            } else {
                for fi50 := 0; fi50 < 17; fi50++ {
                    if v3 {
                        if false {
                            var v51 uint64 = 18446744073709551615
                            if true {
                                var v52 uint32 = 548
                                if v3 {
                                    var v53 int16 = (v5 >> ((2) & 63))
                                    _ = v53
                                }
                                _ = v52
                            } else {
                                v54 := GenS2{F0: v1, F1: 158}
                                _ = v54
                            }
                            _ = v51
                        }
                    } else {
                        v55 := GenS2{F0: (v1 + "a"), F1: 128}
                        if true {
                            _ = len(v2)
                            var v56 uint = 957
                            _ = v56
                        }
                        v57 := GenS8{F0: 65535, F1: 1}
                        _ = v55
                        _ = v57
                    }
                    if (-2.25 == 1.5) {
                        switch (^(fi7)) {
                        case 4294967295:
                            var v58 uint16 = 863
                            _ = v58
                        }
                    }
                    switch fi50 {
                    case 42:
                        var v59 int64 = 0
                        _ = v59
                    case -780:
                        var v60 int64 = 0
                        _ = v60
                    default:
                        if (v1 == v1) {
                            v3 = ((!false) || false)
                            v61 := GenS9{F0: 446, F1: 4294967295, F2: 42}
                            var v62 int16 = (v5 * v5)
                            _ = v61
                            _ = v62
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v63 bool = false
                            v2 = append(v2, 2)
                            v1 = v1
                            _ = v63
                        }
                    }
                    var v64 uint32 = 2
                    v65 := GenS8{F0: 32767, F1: 799}
                    _ = fi50
                    _ = v64
                    _ = v65
                }
            }
            _ = fi7
        }
    }
    v1 = "hello"
    fmt.Println(v3)
    var v66 float32 = 1.5
    if ((3.14159 / ((v66 + ((v66 / v66) * ((v66 * float32((-2.25 / -2.25))) + v66))) - (float32(float64(v66)) - (3.14159 - (3.14159 / v66))))) > v66) {
        _ = len(v2)
    }
    if (!v3) {
        v5 += (^(v5))
        if (!(1 < 2147483647)) {
            var v67 int = -9223372036854775808
            if (0 <= -1) {
                v1 = v1
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            if (v5 <= (v5 | v5)) {
                v68 := GenS2{F0: v1, F1: 42}
                _ = v68
            }
            _ = v67
        } else {
            v1 = "hello"
            v1 = v1
            v5 += v5
        }
        if v3 {
            if (42 >= 2) {
                if v3 {
                    fmt.Println(v5, v3, v66)
                    v69 := GenS7{F0: 2147483648, F1: "\n\t"}
                    var v70 int32 = -1
                    _ = v69
                    _ = v70
                }
                var v71 uint32 = 1
                _ = v71
            } else {
                v1 = ("a" + (v1 + "foo bar"))
                v5 += (v5 | (v5 << ((65535) & 63)))
                v72 := GenS6{F0: 0}
                _ = v72
            }
        }
    } else {
        var v73 int8 = 100
        _ = v73
    }
    v74 := GenS2{F0: v1, F1: 236}
    v66 = v66
    v2 = append(v2, 4294967295)
    var v75 float64 = (-2.25 - 3.14159)
    if false {
        var v76 int8 = 127
        _ = v76
    } else {
        v75 += 0.0
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v66:", v66)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v66
    _ = v74
    _ = v75
}

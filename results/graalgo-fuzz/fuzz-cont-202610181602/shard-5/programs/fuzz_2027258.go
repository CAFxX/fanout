package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 int32; F1 uint; F2 uint32 }
type GenS7 struct { F0 int64; F1 int; F2 int }
type GenS8 struct { F0 uint; F1 uint8; F2 uint }

func main() {
    v1 := []string{"\n\t", "hello"}
    var v2 string = "a"
    v3 := make(map[string]uint32)
    var v4 int = 4294967295
    v5 := make(map[int]uint32)
    var v6 int = 32768
    var v7 int = (^(v4))
    var v8 int64 = -9223372036854775808
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        switch (v6 >> ((4294967295) & 63)) {
        case 32767:
            delete(v5, v6)
        }
        v9 := GenS1{F0: 437}
        _ = v9
    }
    v7 += (v4 & 0)
    if false {
        fmt.Println(v2, v7, v7)
        v4 = -1
    } else {
        v7 = (v4 >> ((42) & 63))
        fmt.Println(v2)
    }
    if (-98 < -84) {
        v8 += -831
    }
    var v10 int = v7
    if true {
        switch (v2 + (((v2 + ((v2 + v2) + "a")) + "") + "")) {
        case "foo bar":
            if false {
                v1 = append(v1, "hello")
            } else {
                v11 := GenS1{F0: 4294967295}
                v12 := v3[((v2 + (("" + v2) + ((v2 + v2) + (("\n\t" + (("a" + v2) + v2)) + v2)))) + v2)]
                var v13 int = v7
                _ = v11
                _ = v12
                _ = v13
            }
        case "":
            if ((-128 >= -1) && (-1 != 42)) {
                for fi14 := 0; fi14 < 13; fi14++ {
                    if false {
                        if false {
                            v4 += (^(fi14))
                            v8 = v8
                        }
                        var v15 int64 = 1
                        for fi16 := 0; fi16 < 11; fi16++ {
                            if len(v1) > 0 {
                                v1[0] = v2
                            }
                            v8 = v15
                            var v17 bool = false
                            _ = fi16
                            _ = v17
                        }
                        _ = v15
                    } else {
                        if len(v1) > 0 {
                            v1[0] = v2
                        }
                        v18 := GenS2{F0: (v2 + "")}
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 20; fi19++ {
                        switch fi14 {
                        case 1:
                            v20 := GenS2{F0: (v2 + ((v2 + (((((v2 + (("\n\t" + ((v2 + "") + v2)) + (("hello" + v2) + (v2 + ("\n\t" + v2))))) + v2) + v2) + (("" + (("\n\t" + v2) + v2)) + v2)) + "")) + v2))}
                            _ = v20
                        case 4294967295:
                            if true {
                                if (4294967295 >= 2147483647) {
                                    v21 := GenS2{F0: (v2 + v2)}
                                    switch v4 {
                                    case -9223372036854775808:
                                        if false {
                                            v10 -= fi14
                                            v6 -= v4
                                            v8 = v8
                                            var v22 uint16 = 684
                                            _ = v22
                                        }
                                    case 460:
                                        v1 = append(v1, "\n\t")
                                    case 201:
                                        var v23 uint64 = 2
                                        _ = v23
                                    default:
                                        v1 = append(v1, v2)
                                    }
                                    _ = v21
                                }
                                var v24 int8 = -1
                                v6 = int(v8)
                                _ = v24
                            } else {
                                fmt.Println(v10, v2)
                                v25 := GenS2{F0: v2}
                                _ = v25
                            }
                        }
                        switch (v6 * v7) {
                        case 499:
                            var v26 uint = 2
                            _ = v26
                        case 4294967295:
                            v7 = v10
                        case 2147483647:
                            v10 = 9223372036854775807
                        }
                        _ = fi19
                    }
                    _ = len(v1)
                    _ = fi14
                }
                fmt.Println(v2, v4, v7)
                if false {
                    fmt.Println(v2, v8)
                    var v27 bool = (0.1 > 1e10)
                    var v28 uint32 = 4294967295
                    _ = v27
                    _ = v28
                }
            }
        }
        for fi29 := 0; fi29 < 17; fi29++ {
            v6 += (v10 / ((v7) | 1))
            if false {
                v30 := GenS3{F0: 42}
                _ = v30
            }
            _ = fi29
        }
        v7 = (v10 | (v7 >> ((699) & 63)))
    } else {
        if false {
            v31 := GenS4{F0: v6}
            if true {
                for fi32 := 0; fi32 < 5; fi32++ {
                    fmt.Println(v10)
                    switch v2 {
                    case "foo bar":
                        switch v2 {
                        case "":
                            v33 := GenS5{F0: -1, F1: 1000}
                            _ = v33
                        case "a":
                            fmt.Println(v7)
                        default:
                            _ = len(v1)
                        }
                    case "\n\t":
                        if false {
                            switch "" {
                            case "":
                                if false {
                                    var v34 int = 32768
                                    v34 = (v4 * v4)
                                    if false {
                                        v35 := GenS6{F0: 42, F1: 255, F2: 2147483648}
                                        var v36 string = ((v2 + v2) + v2)
                                        v1 = append(v1, v36)
                                        _ = v35
                                        _ = v36
                                    } else {
                                        var v37 int64 = v8
                                        v38 := GenS7{F0: (v37 >> ((2) & 63)), F1: (v34 >> ((376) & 63)), F2: (v6 >> ((0) & 63))}
                                        _ = v37
                                        _ = v38
                                    }
                                    _ = v34
                                }
                            }
                            v39 := GenS8{F0: 18446744073709551615, F1: 128, F2: 42}
                            _ = v39
                        } else {
                            switch v2 {
                            case "a":
                                if true {
                                    v40 := GenS8{F0: 721, F1: 255, F2: 285}
                                    var v41 uint32 = 547
                                    _ = v40
                                    _ = v41
                                }
                            default:
                                fmt.Println(fi32)
                            }
                        }
                    case "":
                        if false {
                            if (3.14159 <= float64(float32(0.1))) {
                                v42 := GenS8{F0: 139, F1: 255, F2: 18446744073709551615}
                                v43 := GenS6{F0: 0, F1: 822, F2: 910}
                                _ = v42
                                _ = v43
                            }
                            var v44 uint8 = 2
                            v45 := GenS3{F0: -1}
                            _ = v44
                            _ = v45
                        } else {
                            var v46 int16 = 1
                            v47 := GenS2{F0: (v2 + v2)}
                            _ = v46
                            _ = v47
                        }
                    default:
                        var v48 float32 = 3.14159
                        _ = v48
                    }
                    switch (("foo bar" + v2) + v2) {
                    case "":
                        v10 = 127
                    case "a":
                        var v49 float32 = (1.5 * -2.25)
                        _ = v49
                    case "\n\t":
                        if true {
                            fmt.Println(v7, v2, v10)
                            var v50 uint32 = 315
                            _ = v50
                        }
                    }
                    _ = fi32
                }
            } else {
                v10 = (v7 - v10)
                for fi51 := 0; fi51 < 13; fi51++ {
                    var v52 string = (v2 + v2)
                    _ = fi51
                    _ = v52
                }
            }
            _ = v31
        }
        var v53 int64 = (v8 * v8)
        _ = v53
    }
    v7 = v7
    var v54 uint64 = 2
    if false {
        v55 := GenS1{F0: v54}
        v7 += 414
        _ = v55
    }
    v56 := GenS3{F0: 42}
    var v57 uint8 = 127
    for fi58 := 0; fi58 < 16; fi58++ {
        var v59 bool = false
        _ = fi58
        _ = v59
    }
    if true {
        v60 := GenS7{F0: v8, F1: (^(v4)), F2: int(v8)}
        _ = v60
    } else {
        v5[316] = 2147483647
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v56
    _ = v57
    _ = v6
    _ = v7
    _ = v8
}

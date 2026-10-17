package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16; F1 bool }
type GenS7 struct { F0 uint64; F1 string }
type GenS8 struct { F0 uint32; F1 string; F2 int16 }
type GenS9 struct { F0 int16; F1 int }

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 656
    var v3 string = "foo bar"
    v4 := make(map[uint32]uint32)
    v5 := []int{1, -782, 249}
    switch ((("\n\t" + ((("" + v3) + v3) + (v3 + (v3 + "")))) + v3) + "a") {
    case "":
        for fi6 := 0; fi6 < 14; fi6++ {
            fmt.Println(fi6, fi6, v2)
            for fi7 := 0; fi7 < 12; fi7++ {
                var v8 int = (fi7 % ((0) | 1))
                if (((fi7 <= (fi6 ^ 706)) && (!false)) || false) {
                    if true {
                        if false {
                            switch (fi6 * fi7) {
                            case -2147483648:
                                v9 := GenS1{F0: true}
                                _ = v9
                            }
                            v10 := GenS1{F0: false}
                            if true {
                                switch 128 {
                                case -734:
                                    switch -9223372036854775808 {
                                    case 4294967295:
                                        fmt.Println(fi6, v8, v8)
                                    default:
                                        var v11 int32 = -413
                                        if (!((18446744073709551615 <= 4294967295) || (!true))) {
                                            var v12 int64 = 16
                                            _ = v12
                                        }
                                        v13 := GenS2{F0: 192}
                                        _ = v11
                                        _ = v13
                                    }
                                case -10:
                                    v8 = -920
                                default:
                                    if true {
                                        if false {
                                            delete(v1, fi6)
                                            v2 += v2
                                        } else {
                                            v14 := GenS3{F0: 4294967295}
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            fmt.Println(v3)
                                            _ = v14
                                        }
                                        v2 = 753
                                    } else {
                                        var v15 uint = 590
                                        _ = v15
                                    }
                                }
                                v2 -= v2
                                var v16 uint64 = 1
                                _ = v16
                            }
                            _ = v10
                        }
                    } else {
                        fmt.Println(fi7)
                        v17 := GenS1{F0: true}
                        _ = v17
                    }
                    fmt.Println(fi7)
                } else {
                    if (0 != 919) {
                        v2 = 1
                        switch (v8 % ((fi7) | 1)) {
                        case 1:
                            var v18 int8 = -1
                            _ = v18
                        }
                    } else {
                        v3 = ""
                        v4[847] = 508
                    }
                    var v19 int64 = 926
                    v20 := GenS3{F0: 1}
                    _ = v19
                    _ = v20
                }
                if (true || (false || (!(true && true)))) {
                    if false {
                        _ = len(v5)
                        switch v3 {
                        case "foo bar":
                            v21 := GenS1{F0: (uint64(v2) == 151)}
                            _ = v21
                        case "a":
                            v1[-1] = v3
                        default:
                            var v22 int16 = -1
                            _ = v22
                        }
                        v23 := GenS3{F0: 42}
                        _ = v23
                    } else {
                        v3 = "hello"
                        if true {
                            v3 = v3
                            var v24 int32 = 2147483647
                            v3 = ""
                            v25 := GenS1{F0: false}
                            _ = v24
                            _ = v25
                        }
                        v1[(^(fi6))] = "hello"
                    }
                    v8 = v8
                    v2 = v2
                } else {
                    v26 := GenS4{F0: 255, F1: int64(v2)}
                    _ = v26
                }
                var v27 int32 = -2147483648
                var v28 uint16 = 348
                var v29 uint32 = (v2 + v2)
                _ = fi7
                _ = v27
                _ = v28
                _ = v29
                _ = v8
            }
            fmt.Println(v2, v2)
            _ = fi6
        }
    case "\n\t":
        var v30 uint64 = 236
        _ = v30
    case "foo bar":
        for fi31 := 0; fi31 < 4; fi31++ {
            var v32 uint32 = v2
            _ = fi31
            _ = v32
        }
    }
    v33 := GenS5{F0: 255}
    v2 = v2
    v3 = v3
    if true {
        var v34 int16 = 0
        _ = v34
    } else {
        for fi35 := 0; fi35 < 6; fi35++ {
            var v36 float64 = ((-2.25 - 1.5) - 3.14159)
            _ = fi35
            _ = v36
        }
        for fi37 := 0; fi37 < 9; fi37++ {
            fmt.Println(fi37)
            switch -642 {
            case -9223372036854775808:
                var v38 uint8 = 0
                _ = v38
            }
            for fi39 := 0; fi39 < 2; fi39++ {
                if false {
                    v40 := GenS6{F0: 2, F1: (191 >= 127)}
                    _ = v40
                }
                var v41 int64 = 42
                _ = fi39
                _ = v41
            }
            _ = fi37
        }
    }
    v42 := GenS7{F0: 559, F1: "a"}
    var v43 int16 = 0
    for fi44 := 0; fi44 < 18; fi44++ {
        for fi45 := 0; fi45 < 17; fi45++ {
            var v46 int16 = v43
            _ = fi45
            _ = v46
        }
        if false {
            v47 := GenS8{F0: v2, F1: v3, F2: (v43 / (((v43 << ((255) & 63))) | 1))}
            fmt.Println(v2, v43)
            v48 := GenS9{F0: 1000, F1: (fi44 / ((fi44) | 1))}
            _ = v47
            _ = v48
        } else {
            for fi49 := 0; fi49 < 20; fi49++ {
                _ = len(v1)
                v2 = 630
                _ = fi49
            }
            var v50 float32 = 3.14159
            switch fi44 {
            case 825:
                switch fi44 {
                case 1:
                    v43 = 42
                default:
                    _ = len(v4)
                }
            case 42:
                v51 := GenS7{F0: 0, F1: v3}
                _ = v51
            case 9223372036854775807:
                var v52 int64 = -222
                _ = v52
            }
            _ = v50
        }
        v1[fi44] = (v3 + "\n\t")
        _ = fi44
    }
    v2 = (v2 & v2)
    for fi53 := 0; fi53 < 17; fi53++ {
        if false {
            v43 = v43
        } else {
            fmt.Println(fi53, v3, v2)
        }
        fmt.Println(v3, v43, fi53)
        _ = fi53
    }
    for fi54 := 0; fi54 < 19; fi54++ {
        v43 = (^(v43))
        if (!(fi54 <= fi54)) {
            v2 += 435
            for fi55 := 0; fi55 < 11; fi55++ {
                var v56 int64 = -1
                v5 = append(v5, (fi55 >> ((1) & 63)))
                _ = fi55
                _ = v56
            }
        } else {
            if true {
                var v57 int = -9223372036854775808
                _ = v57
            }
        }
        _ = fi54
    }
    var v58 uint64 = 1
    v59 := GenS2{F0: 255}
    fmt.Println(v58, v2)
    delete(v1, -9223372036854775808)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v58
    _ = v59
}

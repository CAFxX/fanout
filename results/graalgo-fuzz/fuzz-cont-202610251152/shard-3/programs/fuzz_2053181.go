package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 string; F1 int }
type GenS7 struct { F0 int64; F1 int; F2 uint8 }
type GenS8 struct { F0 uint32; F1 uint64; F2 int }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int32 }

func main() {
    var v1 uint = 42
    v2 := []uint32{}
    var v3 uint64 = 929
    v4 := []int{9223372036854775807}
    if false {
        for fi5 := 0; fi5 < 2; fi5++ {
            var v6 int16 = -1
            for fi7 := 0; fi7 < 10; fi7++ {
                switch "a" {
                case "a":
                    v8 := GenS1{F0: false, F1: (42 <= 1), F2: 585}
                    _ = v8
                case "":
                    switch (fi7 / (((fi5 % ((fi5) | 1))) | 1)) {
                    case 128:
                        v1 = (v1 << (((^(v1))) & 63))
                    case 0:
                        v1 -= (^(v1))
                    default:
                        v9 := GenS1{F0: true, F1: true, F2: 901}
                        _ = v9
                    }
                case "hello":
                    v6 = 42
                }
                var v10 float64 = 1.5
                _ = fi7
                _ = v10
            }
            _ = fi5
            _ = v6
        }
    } else {
        v4 = append(v4, 0)
        if true {
            var v11 int8 = 1
            fmt.Println(v3)
            _ = v11
        }
        for fi12 := 0; fi12 < 12; fi12++ {
            v13 := GenS2{F0: -279, F1: 348, F2: 1}
            _ = fi12
            _ = v13
        }
    }
    var v14 int32 = 858
    var v15 string = (("\n\t" + "foo bar") + "a")
    v1 += 65535
    var v16 uint64 = v3
    v1 = v1
    var v17 bool = false
    for fi18 := 0; fi18 < 20; fi18++ {
        var v19 int16 = 42
        var v20 uint8 = 127
        _ = fi18
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        v3 += 1
        v22 := GenS2{F0: 264, F1: 65535, F2: 2}
        _ = fi21
        _ = v22
    }
    v23 := GenS3{F0: 1}
    v15 = ("\n\t" + ((v15 + "a") + (((v15 + "") + (v15 + "\n\t")) + v15)))
    var v24 float64 = -2.25
    switch -319 {
    case -9223372036854775808:
        if (v15 == "foo bar") {
            if v17 {
                switch -647 {
                case 840:
                    v14 = v14
                case 9223372036854775807:
                    var v25 uint64 = v16
                    _ = v25
                case 32767:
                    v26 := GenS4{F0: 1, F1: 0}
                    _ = v26
                }
                for fi27 := 0; fi27 < 1; fi27++ {
                    v15 = v15
                    v14 = v14
                    v17 = (2147483647 > 4294967295)
                    _ = fi27
                }
            }
            _ = len(v2)
            if v17 {
                var v28 uint8 = 128
                v16 -= 0
                _ = v28
            }
        } else {
            v17 = (v17 || v17)
            v29 := GenS5{F0: 0}
            _ = v29
        }
    case -135:
        if false {
            var v30 float64 = 1.5
            fmt.Println(v15, v3)
            switch 1 {
            case -9223372036854775808:
                fmt.Println(v15, v3, v15)
            }
            _ = v30
        }
    case -283:
        if v17 {
            v31 := GenS6{F0: (v15 + (v15 + v15)), F1: 9223372036854775807}
            fmt.Println(v3)
            _ = v31
        } else {
            if v17 {
                switch 1 {
                case 0:
                    var v32 int64 = 723
                    _ = v32
                case 32768:
                    v33 := GenS4{F0: 2, F1: -9223372036854775808}
                    _ = v33
                default:
                    fmt.Println(v15, v15)
                }
                var v34 float64 = 0.1
                _ = v34
            } else {
                v35 := GenS5{F0: v16}
                _ = v35
            }
            v36 := GenS7{F0: 42, F1: 128, F2: 255}
            _ = v36
        }
    default:
        if (v17 && false) {
            for fi37 := 0; fi37 < 19; fi37++ {
                v2 = append(v2, 2147483648)
                var v38 uint8 = 0
                if false {
                    fmt.Println(v15, v24)
                    switch 959 {
                    case -9223372036854775808:
                        fmt.Println(v15)
                    case 2147483647:
                        for fi39 := 0; fi39 < 14; fi39++ {
                            if (!((v14 ^ v14) > 42)) {
                                switch fi37 {
                                case -2147483648:
                                    v40 := GenS7{F0: -9223372036854775808, F1: (fi37 & fi39), F2: v38}
                                    _ = v40
                                case 128:
                                    v14 += (v14 * (^(v14)))
                                case -836:
                                    fmt.Println(v14)
                                    switch fi37 {
                                    case 4294967295:
                                        var v41 bool = (int(v3) != 9223372036854775807)
                                        _ = v41
                                    }
                                    v42 := GenS6{F0: (v15 + ""), F1: fi37}
                                    _ = v42
                                default:
                                    if (-128 < 40) {
                                        v43 := GenS1{F0: (1 >= -7), F1: (((v1 * v1) >= v1) && (1 != 765)), F2: 876}
                                        v24 = v24
                                        _ = v43
                                    }
                                }
                                if v17 {
                                    v44 := GenS7{F0: 9223372036854775807, F1: fi37, F2: (v38 - 2)}
                                    if v17 {
                                        fmt.Println(v14, fi37, v15)
                                        if len(v2) > 0 {
                                            v2[0] = 2147483647
                                        }
                                        fmt.Println(v1, v16)
                                    }
                                    v17 = (-1 <= int64(v16))
                                    _ = v44
                                }
                                if (v17 && true) {
                                    if false {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v45 := GenS1{F0: ((v14 << (((v1 % ((v1) | 1))) & 63)) < v14), F1: true, F2: 255}
                                        v1 -= 1
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v45
                                    } else {
                                        v24 = 1e10
                                        v38 -= 42
                                    }
                                    var v46 int8 = -128
                                    var v47 float32 = (1e10 * 0.1)
                                    var v48 bool = (824 <= v1)
                                    v14 = (v14 >> ((v1) & 63))
                                    _ = v46
                                    _ = v47
                                    _ = v48
                                }
                            }
                            _ = fi39
                        }
                    default:
                        fmt.Println(v15, v3)
                    }
                } else {
                    if (42 > (v3 | v16)) {
                        for fi49 := 0; fi49 < 4; fi49++ {
                            if (!v17) {
                                if v17 {
                                    if v17 {
                                        if (v1 < (^(v1))) {
                                            fmt.Println(v16)
                                        } else {
                                            switch (fi37 + fi49) {
                                            case 9223372036854775807:
                                                switch ((v15 + v15) + v15) {
                                                case "":
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    var v50 int16 = -510
                                                    _ = v50
                                                case "foo bar":
                                                    var v51 string = (v15 + (((v15 + v15) + v15) + "a"))
                                                    _ = v51
                                                default:
                                                    var v52 uint64 = (v16 / (((v16 / ((v3) | 1))) | 1))
                                                    _ = v52
                                                }
                                            case 32768:
                                                if v17 {
                                                    v3 = 741
                                                    var v53 int16 = 1
                                                    _ = v53
                                                } else {
                                                    var v54 float32 = 3.14159
                                                    v14 -= v14
                                                    _ = v54
                                                }
                                            case 4294967295:
                                                var v55 uint16 = 149
                                                _ = v55
                                            }
                                        }
                                        if (!v17) {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            if v17 {
                                                v56 := GenS8{F0: 0, F1: v16, F2: 32768}
                                                var v57 string = "hello"
                                                if v17 {
                                                    var v58 int = fi37
                                                    v15 = v15
                                                    var v59 float64 = (0.0 / float64(3.14159))
                                                    _ = v58
                                                    _ = v59
                                                } else {
                                                    v57 = "foo bar"
                                                }
                                                _ = v56
                                                _ = v57
                                            } else {
                                                var v60 uint64 = 382
                                                v4 = append(v4, (fi49 + fi49))
                                                _ = v60
                                            }
                                        } else {
                                            v2 = append(v2, uint32(v38))
                                            var v61 float32 = float32(float64(0.1))
                                            _ = v61
                                        }
                                    }
                                    var v62 int64 = -1
                                    var v63 int8 = 42
                                    _ = v62
                                    _ = v63
                                }
                                switch fi37 {
                                case 4294967295:
                                    if len(v2) > 0 {
                                        v2[0] = 633
                                    }
                                case 9223372036854775807:
                                    v64 := GenS7{F0: 9223372036854775807, F1: (fi37 / ((int(v16)) | 1)), F2: v38}
                                    _ = v64
                                case 127:
                                    v65 := GenS9{F0: 128}
                                    if len(v4) > 0 {
                                        v4[0] = fi37
                                    }
                                    _ = v65
                                }
                            } else {
                                v66 := GenS10{F0: (^(v14))}
                                v24 = v24
                                _ = v66
                            }
                            v67 := GenS9{F0: fi49}
                            _ = fi49
                            _ = v67
                        }
                        v38 = v38
                    }
                }
                _ = fi37
                _ = v38
            }
            if (255 >= 255) {
                var v68 uint = 42
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v68
            } else {
                for fi69 := 0; fi69 < 18; fi69++ {
                    _ = len(v2)
                    _ = fi69
                }
            }
            var v70 int32 = v14
            _ = v70
        } else {
            v71 := GenS8{F0: 706, F1: (v3 / (((v3 % ((v3) | 1))) | 1)), F2: 128}
            var v72 int16 = 42
            v73 := GenS2{F0: v14, F1: 65535, F2: 65535}
            _ = v71
            _ = v72
            _ = v73
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}

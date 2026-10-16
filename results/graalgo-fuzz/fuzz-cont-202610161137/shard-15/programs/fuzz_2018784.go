package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS3 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint8; F1 uint32; F2 string }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS7 struct { F0 uint16; F1 bool }
type GenS8 struct { F0 int8 }

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 2
    var v3 uint8 = (v2 - v2)
    v4 := []int{4294967295, -9223372036854775808, 42, int(v2)}
    var v5 int16 = -32768
    if true {
        v6 := GenS1{F0: 735}
        fmt.Println(v5, v3, v3)
        _ = v6
    }
    if true {
        var v7 int16 = v5
        switch 662 {
        case 53:
            if (1 != 65535) {
                for fi8 := 0; fi8 < 6; fi8++ {
                    v3 -= (v3 >> ((42) & 63))
                    var v9 int8 = 0
                    var v10 uint32 = 888
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                var v11 int8 = 83
                _ = v11
            } else {
                var v12 float32 = -2.25
                _ = v12
            }
        case -734:
            v3 -= v3
        default:
            if true {
                for fi13 := 0; fi13 < 4; fi13++ {
                    var v14 float32 = 3.14159
                    switch 128 {
                    case 42:
                        var v15 int16 = v5
                        _ = v15
                    case 1:
                        for fi16 := 0; fi16 < 15; fi16++ {
                            v7 += (^(v7))
                            _ = len(v1)
                            _ = fi16
                        }
                    }
                    switch "" {
                    case "\n\t":
                        v17 := GenS1{F0: 32767}
                        _ = v17
                    case "foo bar":
                        v18 := GenS2{F0: 0, F1: 351, F2: 414}
                        _ = v18
                    default:
                        v2 = 55
                    }
                    _ = fi13
                    _ = v14
                }
            }
        }
        _ = v7
    } else {
        if false {
            if ((46 < 2147483648) || false) {
                switch -9223372036854775808 {
                case -1:
                    if true {
                        v3 = v2
                    } else {
                        var v19 uint8 = 1
                        fmt.Println(v5, v2, v2)
                        var v20 bool = ((262 > 65535) || true)
                        _ = v19
                        _ = v20
                    }
                case 128:
                    if (false && true) {
                        fmt.Println(v5, v3)
                    } else {
                        if (!false) {
                            for fi21 := 0; fi21 < 6; fi21++ {
                                if true {
                                    v2 = (v2 % (((v2 & (v3 + 2))) | 1))
                                    var v22 int64 = 42
                                    var v23 uint32 = 1
                                    _ = v22
                                    _ = v23
                                } else {
                                    for fi24 := 0; fi24 < 18; fi24++ {
                                        switch (("" + "a") + "hello") {
                                        case "a":
                                            v4 = append(v4, -752)
                                        case "foo bar":
                                            if (1 <= 2147483647) {
                                                if (v2 <= 2) {
                                                    var v25 int8 = 100
                                                    _ = v25
                                                }
                                                v2 -= 1
                                                switch fi24 {
                                                case -746:
                                                    _ = len(v4)
                                                }
                                            } else {
                                                v26 := GenS1{F0: 421}
                                                switch -9223372036854775808 {
                                                case 865:
                                                    v2 += v3
                                                }
                                                _ = v26
                                            }
                                        case "":
                                            v27 := GenS3{F0: 0, F1: 42, F2: 2}
                                            _ = v27
                                        }
                                        _ = fi24
                                    }
                                    v3 -= v3
                                    fmt.Println(fi21, fi21, v5)
                                }
                                _ = fi21
                            }
                            for fi28 := 0; fi28 < 3; fi28++ {
                                v2 += (^(v3))
                                _ = fi28
                            }
                            delete(v1, ("hello" + "hello"))
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v3 -= (v2 + v2)
                    }
                case 4294967295:
                    switch -514 {
                    case -234:
                        v29 := GenS4{F0: v2, F1: 2147483647, F2: "foo bar"}
                        _ = v29
                    }
                }
                v4 = append(v4, 596)
            } else {
                switch "hello" {
                case "foo bar":
                    var v30 int64 = 0
                    _ = v30
                case "":
                    if false {
                        v3 += (v3 >> ((987) & 63))
                        var v31 int64 = 833
                        switch ("a" + "a") {
                        case "hello":
                            var v32 string = ("\n\t" + "foo bar")
                            _ = v32
                        case "a":
                            v5 = (v5 << ((2) & 63))
                        }
                        _ = v31
                    }
                case "\n\t":
                    var v33 string = ""
                    _ = v33
                default:
                    var v34 float64 = (1e10 + 3.14159)
                    _ = v34
                }
            }
        } else {
            _ = len(v4)
            if true {
                var v35 int32 = -2147483648
                v36 := GenS4{F0: 1, F1: 42, F2: "\n\t"}
                if ((v35 + v35) == int32(v3)) {
                    for fi37 := 0; fi37 < 6; fi37++ {
                        var v38 bool = true
                        v39 := GenS1{F0: 32767}
                        var v40 uint64 = 18446744073709551615
                        _ = fi37
                        _ = v38
                        _ = v39
                        _ = v40
                    }
                    v1["\n\t"] = 275
                    var v41 int = 2147483647
                    _ = v41
                } else {
                    switch "hello" {
                    case "hello":
                        var v42 bool = (2147483647 < -485)
                        _ = v42
                    case "foo bar":
                        v43 := GenS2{F0: 100, F1: 199, F2: 255}
                        _ = v43
                    default:
                        var v44 int64 = -1
                        _ = v44
                    }
                }
                _ = v35
                _ = v36
            } else {
                if (!(0 < 88)) {
                    v3 += v3
                    for fi45 := 0; fi45 < 9; fi45++ {
                        var v46 int64 = 9223372036854775807
                        v46 = v46
                        _ = fi45
                        _ = v46
                    }
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                for fi47 := 0; fi47 < 7; fi47++ {
                    fmt.Println(v5, v3, v3)
                    _ = fi47
                }
            }
            for fi48 := 0; fi48 < 7; fi48++ {
                v49 := GenS5{F0: 0}
                _ = fi48
                _ = v49
            }
        }
    }
    v3 = v3
    v4 = append(v4, -9223372036854775808)
    var v50 uint64 = 0
    if (105 > 133) {
        for fi51 := 0; fi51 < 19; fi51++ {
            var v52 float32 = 3.14159
            v53 := GenS5{F0: 65535}
            if false {
                fmt.Println(v52)
            }
            _ = fi51
            _ = v52
            _ = v53
        }
        var v54 float64 = (-2.25 - ((float64((-2.25 - 1.5)) * 0.1) + -2.25))
        v2 = (^(v3))
        _ = v54
    } else {
        if (v50 < 0) {
            var v55 bool = (!(true || false))
            _ = v55
        } else {
            v56 := GenS5{F0: 42}
            switch "\n\t" {
            case "hello":
                var v57 uint16 = 991
                _ = v57
            }
            _ = v56
        }
    }
    switch -2147483648 {
    case 1:
        for fi58 := 0; fi58 < 9; fi58++ {
            v50 -= v50
            fmt.Println(fi58, v50, v50)
            _ = fi58
        }
    case -2147483648:
        for fi59 := 0; fi59 < 4; fi59++ {
            var v60 int32 = -2147483648
            if false {
                var v61 uint16 = 42
                v62 := GenS6{F0: 110, F1: v61, F2: uint32(v50)}
                _ = v61
                _ = v62
            }
            _ = fi59
            _ = v60
        }
    default:
        if (!(-1 >= -1)) {
            fmt.Println(v50, v50, v2)
            switch 4294967295 {
            case -1:
                v63 := GenS7{F0: 255, F1: true}
                _ = v63
            }
        } else {
            if (877 <= 128) {
                if false {
                    var v64 uint = 18446744073709551615
                    v65 := GenS8{F0: -128}
                    v66 := GenS2{F0: 1, F1: 65535, F2: 32767}
                    _ = v64
                    _ = v65
                    _ = v66
                }
                if (!true) {
                    v67 := GenS5{F0: 877}
                    v68 := GenS3{F0: 93, F1: 42, F2: uint16(v50)}
                    v50 -= v50
                    _ = v67
                    _ = v68
                } else {
                    switch 127 {
                    case -43:
                        for fi69 := 0; fi69 < 16; fi69++ {
                            fmt.Println(v3, v5)
                            if false {
                                var v70 uint16 = 0
                                _ = v70
                            }
                            _ = fi69
                        }
                    }
                }
                fmt.Println(v3, v2, v5)
            } else {
                var v71 string = "hello"
                var v72 uint8 = (v2 * (^(v2)))
                _ = v71
                _ = v72
            }
        }
    }
    var v73 uint32 = 42
    fmt.Println(v5, v73, v50)
    v74 := GenS8{F0: -32}
    v2 += v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v73
    _ = v74
}

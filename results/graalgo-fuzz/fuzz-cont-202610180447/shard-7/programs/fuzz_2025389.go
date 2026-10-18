package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 int64; F2 uint64 }
type GenS4 struct { F0 string; F1 uint; F2 int64 }
type GenS5 struct { F0 int32; F1 int; F2 uint64 }
type GenS6 struct { F0 uint64 }
func genFn0(p0 int8) int8 {
    switch -263 {
    case 127:
        fmt.Println(p0)
    default:
        if true {
            v7 := GenS1{F0: -113, F1: (200 == 65535)}
            var v8 int16 = 1000
            _ = v7
            _ = v8
        }
    }
    v9 := GenS1{F0: 2147483647, F1: false}
    var v10 int8 = p0
    for fi11 := 0; fi11 < 6; fi11++ {
        v10 -= p0
        v12 := []int{(fi11 & 203), fi11, 4294967295}
        v13 := GenS2{F0: 4294967295}
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 uint8 = 255
    _ = p0
    _ = v10
    _ = v14
    _ = v9
    return (p0 % ((-128) | 1))
}
func genFn1(p0 bool, p1 bool, p2 bool) int64 {
    var v15 uint = 4294967295
    var v16 bool = (uint64(v15) <= 42)
    v17 := GenS3{F0: 2, F1: 1, F2: 42}
    for fi18 := 0; fi18 < 17; fi18++ {
        v16 = (1e10 >= float64((1.5 * 0.0)))
        v16 = (!p2)
        _ = fi18
    }
    v19 := []int{-9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = -261
    v3 := make(map[uint32]int)
    var v4 bool = (false && true)
    v5 := []string{"hello", "\n\t", ""}
    var v6 uint = 311
    for fi20 := 0; fi20 < 15; fi20++ {
        if (!(!v4)) {
            v21 := GenS3{F0: v6, F1: (v2 - v2), F2: 2}
            _ = v21
        }
        _ = fi20
    }
    switch 127 {
    case 128:
        var v22 uint = v6
        _ = v22
    case 1:
        var v23 float32 = ((float32((3.14159 / 1.0)) / -2.25) * 1.5)
        _ = v23
    case -1:
        for fi24 := 0; fi24 < 1; fi24++ {
            v6 = (v6 % (((^(v6))) | 1))
            for fi25 := 0; fi25 < 7; fi25++ {
                v2 = -1
                _ = fi25
            }
            var v26 uint64 = 42
            _ = fi24
            _ = v26
        }
    default:
        var v27 int = 42
        _ = v27
    }
    delete(v1, "hello")
    var v28 uint = (^(v6))
    v2 = v2
    var v29 int64 = v2
    if v4 {
        v30 := GenS4{F0: "hello", F1: v28, F2: genFn1(v4, (1 > 2), v4)}
        v29 -= v2
        _ = v30
    }
    for fi31 := 0; fi31 < 4; fi31++ {
        v28 = v28
        if v4 {
            var v32 uint = v6
            v33 := GenS5{F0: -823, F1: (fi31 - 4294967295), F2: 2}
            var v34 int8 = -94
            _ = v32
            _ = v33
            _ = v34
        } else {
            fmt.Println(v2, v2)
        }
        var v35 string = ""
        _ = fi31
        _ = v35
    }
    _ = len(v5)
    var v36 int8 = 111
    for fi37 := 0; fi37 < 7; fi37++ {
        v38 := GenS1{F0: -2147483648, F1: v4}
        _ = fi37
        _ = v38
    }
    v39 := GenS5{F0: -2147483648, F1: 1, F2: 42}
    if (0.1 <= 0.1) {
        switch -9223372036854775808 {
        case 135:
            switch (("\n\t" + "foo bar") + "hello") {
            case "\n\t":
                fmt.Println(v4, v28)
            case "":
                v28 += 1
            }
        case -979:
            var v40 uint32 = 731
            _ = v40
        case -1:
            var v41 int32 = 0
            _ = v41
        default:
            v28 += (v28 << ((65535) & 63))
        }
    } else {
        var v42 int32 = 2147483647
        v4 = (v4 || v4)
        _ = v42
    }
    v43 := GenS3{F0: (v28 + (v6 ^ v28)), F1: v29, F2: 1}
    var v44 uint = v6
    for fi45 := 0; fi45 < 1; fi45++ {
        var v46 string = "foo bar"
        switch (fi45 >> ((v44) & 63)) {
        case 1:
            v46 = "a"
        case 225:
            if false {
                if true {
                    v44 = (^(v6))
                    for fi47 := 0; fi47 < 1; fi47++ {
                        _ = len(v1)
                        _ = fi47
                    }
                }
            } else {
                v48 := GenS2{F0: 2147483647}
                if (18446744073709551615 != 440) {
                    v28 = 892
                    v3[2147483647] = 2147483647
                    v49 := GenS6{F0: 18446744073709551615}
                    _ = v49
                } else {
                    switch v46 {
                    case "":
                        _ = len(v1)
                    default:
                        if ((v4 && true) && false) {
                            var v50 uint16 = 340
                            v51 := GenS3{F0: v44, F1: genFn1((135 >= 255), (v50 <= 1), v4), F2: 42}
                            _ = v50
                            _ = v51
                        } else {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                    }
                    var v52 uint8 = 162
                    var v53 int16 = -1
                    _ = v52
                    _ = v53
                }
                _ = v48
            }
        case 127:
            v5 = append(v5, v46)
        default:
            if (!true) {
                for fi54 := 0; fi54 < 19; fi54++ {
                    var v55 int8 = (v36 >> ((v28) & 63))
                    _ = fi54
                    _ = v55
                }
                if true {
                    var v56 uint32 = 42
                    _ = v56
                } else {
                    var v57 bool = true
                    _ = v57
                }
            } else {
                switch fi45 {
                case 1:
                    v44 += (v28 + v44)
                case 0:
                    if false {
                        if v4 {
                            if false {
                                var v58 uint8 = 128
                                delete(v1, (v46 + v46))
                                _ = v58
                            } else {
                                v2 -= v2
                                if v4 {
                                    var v59 int16 = -32768
                                    var v60 uint8 = 128
                                    v6 += 0
                                    _ = v59
                                    _ = v60
                                } else {
                                    if (2147483647 <= 2147483647) {
                                        if v4 {
                                            var v61 uint32 = 2147483648
                                            if ((2147483647 < -1) && v4) {
                                                var v62 uint64 = 549
                                                delete(v1, (v46 + "hello"))
                                                _ = v62
                                            }
                                            var v63 uint8 = 127
                                            _ = v61
                                            _ = v63
                                        }
                                        v29 = -9223372036854775808
                                    }
                                    v28 += 18446744073709551615
                                    fmt.Println(v46, v44)
                                }
                                delete(v1, v46)
                            }
                        } else {
                            v64 := v1["\n\t"]
                            v2 = genFn1((int32(v64) >= -220), v4, true)
                            _ = v64
                        }
                        switch (v46 + v46) {
                        case "a":
                            v44 += (v28 - 42)
                        case "":
                            v36 += v36
                        case "\n\t":
                            var v65 uint8 = 42
                            _ = v65
                        }
                    } else {
                        v6 -= v44
                        for fi66 := 0; fi66 < 11; fi66++ {
                            var v67 float64 = -2.25
                            _ = fi66
                            _ = v67
                        }
                        fmt.Println(v2, v28, fi45)
                    }
                }
                if false {
                    var v68 int16 = 42
                    _ = v68
                }
                switch 4294967295 {
                case -141:
                    fmt.Println(v28, fi45)
                case -2147483648:
                    if len(v5) > 0 {
                        v5[0] = (v46 + "a")
                    }
                case 32768:
                    v3[4294967295] = 4294967295
                default:
                    for fi69 := 0; fi69 < 14; fi69++ {
                        var v70 uint32 = 2
                        v70 -= 0
                        _ = fi69
                        _ = v70
                    }
                }
            }
        }
        _ = fi45
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v36
    _ = v39
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}

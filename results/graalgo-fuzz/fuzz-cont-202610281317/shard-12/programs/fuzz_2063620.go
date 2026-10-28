package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint32; F1 int64; F2 int }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 int16; F1 bool; F2 uint64 }
type GenS5 struct { F0 uint64; F1 uint16 }
type GenS6 struct { F0 string; F1 uint16; F2 uint32 }
func genFn0(p0 int, p1 uint16, p2 string) uint16 {
    switch (p0 << ((18446744073709551615) & 63)) {
    case 42:
        if (9223372036854775807 == (p0 - 44)) {
            for fi7 := 0; fi7 < 11; fi7++ {
                v8 := GenS1{F0: (true || (!true)), F1: 32767, F2: 410}
                var v9 uint32 = 563
                _ = fi7
                _ = v8
                _ = v9
            }
        } else {
            p2 = (p2 + "a")
        }
    case 9223372036854775807:
        v10 := GenS1{F0: false, F1: 42, F2: 65535}
        _ = v10
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        v12 := []int{(p0 >> ((255) & 63))}
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return 456
}

func main() {
    var v1 uint32 = 65535
    var v2 int16 = 42
    var v3 uint8 = 42
    var v4 int8 = -84
    var v5 uint32 = 937
    var v6 int = 127
    v6 += (v6 & v6)
    v3 += 2
    if (false || false) {
        var v13 int16 = (-(v2))
        _ = v13
    } else {
        if (("a" < "\n\t") && (42 == 42)) {
            fmt.Println(v5, v1)
            if true {
                v14 := GenS2{F0: v5, F1: -453, F2: v6}
                v15 := []int{int(v4), v6, (v6 & v6), -9223372036854775808}
                _ = v14
                _ = v15
            }
        }
        if ("a" != "\n\t") {
            fmt.Println(v3, v3)
            for fi16 := 0; fi16 < 19; fi16++ {
                var v17 int64 = 42
                switch (fi16 / (((^(v6))) | 1)) {
                case -140:
                    switch int(v3) {
                    case -9223372036854775808:
                        var v18 uint = 2
                        _ = v18
                    case -862:
                        v19 := []int{(v6 - fi16)}
                        _ = v19
                    case -2147483648:
                        var v20 uint16 = uint16(v3)
                        _ = v20
                    default:
                        if true {
                            var v21 int32 = 0
                            fmt.Println(v4, v1, v3)
                            v17 = v17
                            _ = v21
                        }
                    }
                case -756:
                    v2 = (v2 % ((-490) | 1))
                default:
                    if true {
                        v4 += 42
                        var v22 int32 = 42
                        if (-2.25 <= -2.25) {
                            fmt.Println(v6, v4)
                            switch 32768 {
                            case -1:
                                if ("" == "") {
                                    switch 32767 {
                                    case 2147483647:
                                        v23 := []int{v6, fi16, (fi16 * (v6 + (v6 ^ fi16))), (fi16 & v6)}
                                        _ = v23
                                    default:
                                        for fi24 := 0; fi24 < 9; fi24++ {
                                            v25 := []int{32768, (v6 / (((fi24 / ((32768) | 1))) | 1)), 32767, (-(fi16))}
                                            var v26 float64 = float64(1.5)
                                            _ = fi24
                                            _ = v25
                                            _ = v26
                                        }
                                    }
                                }
                            default:
                                v2 += v2
                            }
                        }
                        _ = v22
                    } else {
                        if true {
                            v27 := []int{0, 127, (v6 % ((int(v17)) | 1))}
                            v28 := GenS1{F0: (620 <= (v2 - 0)), F1: -1, F2: v5}
                            _ = v27
                            _ = v28
                        } else {
                            var v29 float32 = 3.14159
                            for fi30 := 0; fi30 < 16; fi30++ {
                                v31 := GenS2{F0: (v5 | uint32(v3)), F1: -547, F2: 128}
                                _ = fi30
                                _ = v31
                            }
                            fmt.Println(v3, v6, fi16)
                            _ = v29
                        }
                        v5 += v5
                    }
                }
                _ = fi16
                _ = v17
            }
        } else {
            if true {
                if false {
                    fmt.Println(v6, v4, v5)
                }
                for fi32 := 0; fi32 < 20; fi32++ {
                    var v33 uint8 = (v3 + v3)
                    _ = fi32
                    _ = v33
                }
            }
        }
        v34 := []int{2147483647, int(v2)}
        _ = v34
    }
    v35 := GenS3{F0: v6, F1: 42}
    v36 := []int{v6, (v6 % ((0) | 1)), (v6 - v6), 42}
    for fi37 := 0; fi37 < 4; fi37++ {
        switch -944 {
        case -1:
            var v38 uint32 = v1
            _ = v38
        case -972:
            var v39 float64 = 0.1
            _ = v39
        default:
            v36 = append(v36, (fi37 | (fi37 + 886)))
        }
        var v40 uint64 = 0
        var v41 int32 = -316
        _ = fi37
        _ = v40
        _ = v41
    }
    var v42 int16 = 113
    var v43 bool = ("\n\t" != ("foo bar" + ""))
    fmt.Println(v43, v42)
    v4 = v4
    for fi44 := 0; fi44 < 13; fi44++ {
        for fi45 := 0; fi45 < 1; fi45++ {
            _ = len(v36)
            v36 = append(v36, fi45)
            var v46 string = ""
            _ = fi45
            _ = v46
        }
        v47 := GenS4{F0: v42, F1: v43, F2: 899}
        _ = fi44
        _ = v47
    }
    switch "a" {
    case "foo bar":
        for fi48 := 0; fi48 < 6; fi48++ {
            var v49 int16 = int16(v4)
            switch "a" {
            case "a":
                v2 -= v2
            case "":
                v36 = append(v36, int(v3))
            case "\n\t":
                v50 := GenS5{F0: 169, F1: 2}
                _ = v50
            default:
                for fi51 := 0; fi51 < 8; fi51++ {
                    fmt.Println(v1)
                    _ = fi51
                }
            }
            v52 := GenS6{F0: "a", F1: 942, F2: 2147483648}
            _ = fi48
            _ = v49
            _ = v52
        }
    case "":
        var v53 string = "\n\t"
        _ = v53
    case "hello":
        var v54 uint32 = v1
        _ = v54
    default:
        switch v6 {
        case -1:
            var v55 uint32 = (v1 * (v5 & v5))
            _ = v55
        case 0:
            fmt.Println(v1, v4, v42)
        default:
            v56 := GenS5{F0: 360, F1: 65535}
            _ = v56
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint64 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint; F1 uint8; F2 bool }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 uint64; F1 int; F2 bool }
type GenS6 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 uint32 = 42
    var v2 float64 = -2.25
    v3 := make(map[string]int)
    var v4 string = "hello"
    var v5 uint16 = 1
    for fi6 := 0; fi6 < 16; fi6++ {
        for fi7 := 0; fi7 < 6; fi7++ {
            v8 := GenS1{F0: 100, F1: (false || true), F2: 0}
            switch "hello" {
            case "foo bar":
                delete(v3, ("a" + v4))
                v1 = 226
                v9 := GenS2{F0: 42, F1: v1}
                var v10 float64 = v2
                _ = v10
                _ = v9
            case "a":
                fmt.Println(v1, fi7, v2)
            case "\n\t":
                v11 := GenS3{F0: 1, F1: 42, F2: false}
                _ = v11
            default:
                if false {
                    v12 := GenS1{F0: 0, F1: true, F2: 1}
                    _ = v12
                } else {
                    v5 = 2
                    if true {
                        v5 = 65535
                        v13 := GenS1{F0: 8, F1: false, F2: 42}
                        fmt.Println(fi6, v4)
                        switch fi7 {
                        case 32767:
                            var v14 int8 = 0
                            _ = v14
                        case 32768:
                            v5 -= 248
                        case 2147483647:
                            v5 = 221
                        default:
                            var v15 int8 = -128
                            fmt.Println(v4, fi6, v1)
                            _ = v15
                        }
                        _ = v13
                    }
                }
            }
            _ = fi7
            _ = v8
        }
        v16 := []int{}
        _ = fi6
        _ = v16
    }
    var v17 int16 = 1
    for fi18 := 0; fi18 < 7; fi18++ {
        v2 = (v2 / float64(-2.25))
        _ = fi18
    }
    var v19 int32 = -636
    fmt.Println(v4, v4, v1)
    v3[(v4 + ("foo bar" + "hello"))] = 128
    var v20 uint64 = 167
    for fi21 := 0; fi21 < 11; fi21++ {
        v22 := []int{4294967295, 968, fi21, 127}
        _ = fi21
        _ = v22
    }
    if false {
        var v23 uint32 = v1
        for fi24 := 0; fi24 < 12; fi24++ {
            var v25 bool = false
            v26 := v3[(v4 + "a")]
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = v23
    } else {
        var v27 float64 = 0.0
        if (2147483648 <= (v1 & 42)) {
            if true {
                fmt.Println(v2)
                var v28 uint16 = v5
                _ = v28
            }
            if true {
                _ = len(v3)
                var v29 float32 = 1.5
                _ = v29
            } else {
                v30 := []int{127}
                v31 := GenS3{F0: 233, F1: 127, F2: false}
                _ = v30
                _ = v31
            }
        } else {
            var v32 uint32 = (v1 % ((uint32(v5)) | 1))
            var v33 uint16 = (v5 * 807)
            var v34 uint8 = 47
            _ = v32
            _ = v33
            _ = v34
        }
        var v35 uint8 = 31
        _ = v27
        _ = v35
    }
    v36 := []int{32768, 9223372036854775807, -2147483648}
    var v37 int32 = v19
    fmt.Println(v19, v2)
    fmt.Println(v19, v5, v4)
    var v38 uint = 316
    for fi39 := 0; fi39 < 5; fi39++ {
        var v40 bool = true
        v4 = (v4 + v4)
        _ = fi39
        _ = v40
    }
    var v41 int64 = int64(v37)
    if false {
        for fi42 := 0; fi42 < 17; fi42++ {
            var v43 uint = v38
            switch (fi42 % ((-573) | 1)) {
            case 851:
                v44 := GenS2{F0: (v17 | 1000), F1: 65535}
                _ = v44
            default:
                v17 = v17
            }
            _ = fi42
            _ = v43
        }
        var v45 int8 = 42
        var v46 int = 0
        _ = v45
        _ = v46
    } else {
        v17 -= v17
        var v47 uint8 = 203
        for fi48 := 0; fi48 < 5; fi48++ {
            for fi49 := 0; fi49 < 1; fi49++ {
                if false {
                    fmt.Println(v2)
                    var v50 uint8 = 128
                    _ = v50
                } else {
                    v3["a"] = 9223372036854775807
                    switch v4 {
                    case "foo bar":
                        fmt.Println(v1)
                    }
                }
                if false {
                    var v51 uint16 = v5
                    if len(v36) >= 2 {
                        v36 = v36[1:len(v36)]
                    }
                    _ = v51
                }
                if (true && (true || (v5 == v5))) {
                    if (!(uint64(v38) < 4294967295)) {
                        v47 = v47
                        if (v5 < 2) {
                            if false {
                                if true {
                                    var v52 uint64 = (v20 + (v20 / ((v20) | 1)))
                                    v20 -= (v52 >> (((v38 ^ 786)) & 63))
                                    switch 127 {
                                    case -918:
                                        if (!false) {
                                            var v53 float64 = (v2 / ((v2 + v2) * (float64(0.1) - 3.14159)))
                                            var v54 uint = 918
                                            var v55 int8 = -128
                                            _ = v53
                                            _ = v54
                                            _ = v55
                                        }
                                    case 4294967295:
                                        v56 := GenS4{F0: (v47 + (v47 ^ uint8(v5))), F1: v5}
                                        v47 = 77
                                        if true {
                                            v57 := GenS5{F0: v52, F1: (-(fi49)), F2: (int8(v17) < 76)}
                                            _ = v57
                                        }
                                        _ = v56
                                    }
                                    _ = v52
                                } else {
                                    v37 -= 2147483647
                                    v4 = "\n\t"
                                }
                                v58 := GenS2{F0: (v17 / (((v17 << ((255) & 63))) | 1)), F1: 976}
                                _ = v58
                            }
                            v38 -= v38
                            var v59 bool = false
                            _ = v59
                        }
                    } else {
                        var v60 uint32 = uint32(v47)
                        v38 -= v38
                        _ = v60
                    }
                } else {
                    v61 := v3[(v4 + (v4 + (v4 + v4)))]
                    var v62 uint8 = v47
                    _ = v61
                    _ = v62
                }
                _ = fi49
            }
            var v63 float64 = v2
            if true {
                if false {
                    var v64 int16 = v17
                    v65 := GenS6{F0: v47, F1: (v1 >> ((uint(v5)) & 63))}
                    _ = v64
                    _ = v65
                }
            }
            _ = fi48
            _ = v63
        }
        _ = v47
    }
    if len(v36) >= 2 {
        v36 = v36[1:len(v36)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v41
    _ = v5
}

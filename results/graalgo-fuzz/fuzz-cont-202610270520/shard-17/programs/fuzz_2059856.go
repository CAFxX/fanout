package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 uint32; F1 int8; F2 int8 }
type GenS5 struct { F0 int8; F1 uint32; F2 uint8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint16; F1 bool; F2 uint8 }
type GenS8 struct { F0 int32 }

func main() {
    var v1 int8 = 73
    var v2 int8 = (v1 ^ (v1 & v1))
    var v3 bool = true
    var v4 int = 32768
    var v5 uint32 = 42
    var v6 int32 = 849
    switch -1 {
    case 32767:
        v1 -= (v1 | v2)
    case 158:
        v7 := GenS1{F0: 4294967295, F1: (v3 && (v2 <= (v1 - v1))), F2: 2}
        _ = v7
    }
    v4 = v4
    if v3 {
        if false {
            var v8 int32 = (v6 * 626)
            _ = v8
        }
        v3 = v3
        v9 := []int{32768, (v4 << ((2) & 63)), 2147483647, v4}
        _ = v9
    } else {
        v1 = v1
    }
    if v3 {
        var v10 int16 = -46
        var v11 int = (v4 % (((-(v4))) | 1))
        _ = v10
        _ = v11
    }
    v12 := []int{-927, v4}
    if v3 {
        for fi13 := 0; fi13 < 12; fi13++ {
            if ((("hello" + ("\n\t" + ("a" + ""))) + "\n\t") <= "hello") {
                switch "" {
                case "\n\t":
                    v14 := GenS2{F0: -844}
                    _ = v14
                case "":
                    v5 += 42
                case "a":
                    switch "" {
                    case "":
                        fmt.Println(v4, v4)
                    default:
                        var v15 int64 = 9223372036854775807
                        _ = v15
                    }
                }
                var v16 int16 = 32767
                switch 4294967295 {
                case -517:
                    for fi17 := 0; fi17 < 12; fi17++ {
                        v12 = append(v12, (fi13 * fi13))
                        if v3 {
                            var v18 uint16 = 65535
                            _ = v18
                        }
                        if (729 <= 18446744073709551615) {
                            var v19 uint8 = 167
                            v4 -= 2147483647
                            _ = v19
                        } else {
                            v1 += 127
                        }
                        _ = fi17
                    }
                case 32767:
                    v20 := GenS1{F0: 42, F1: v3, F2: 2}
                    _ = v20
                case -1:
                    v2 = v1
                }
                _ = v16
            } else {
                fmt.Println(v2, v2, v4)
                v5 = v5
            }
            if (549 <= 0) {
                _ = len(v12)
                fmt.Println(v5, v6)
            } else {
                for fi21 := 0; fi21 < 17; fi21++ {
                    v6 = (v6 - (v6 % (((v6 | v6)) | 1)))
                    var v22 uint64 = 0
                    _ = fi21
                    _ = v22
                }
                if v3 {
                    var v23 uint16 = 716
                    if len(v12) > 0 {
                        v12[0] = (^(v4))
                    }
                    _ = v23
                } else {
                    switch fi13 {
                    case -572:
                        for fi24 := 0; fi24 < 18; fi24++ {
                            v2 = (v1 % ((v2) | 1))
                            if false {
                                if false {
                                    v25 := GenS3{F0: (-(fi24)), F1: true}
                                    v26 := GenS3{F0: 42, F1: v3}
                                    _ = v25
                                    _ = v26
                                }
                            } else {
                                if ((float32(3.14159) + (0.1 - ((3.14159 + 1.5) - 0.1))) == 1e10) {
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                } else {
                                    v4 = (fi24 + fi13)
                                    v27 := GenS2{F0: -1}
                                    _ = v27
                                }
                                v28 := GenS4{F0: 4294967295, F1: v1, F2: 127}
                                var v29 string = ("a" + "\n\t")
                                _ = v28
                                _ = v29
                            }
                            var v30 uint = 746
                            _ = fi24
                            _ = v30
                        }
                    case -2147483648:
                        if v3 {
                            v31 := GenS4{F0: (v5 ^ 596), F1: -62, F2: (v2 ^ v2)}
                            v4 += v4
                            if len(v12) > 0 {
                                v12[0] = v4
                            }
                            _ = v31
                        } else {
                            var v32 float64 = -2.25
                            var v33 int = 32767
                            switch -9223372036854775808 {
                            case 32768:
                                v34 := GenS5{F0: v1, F1: (^(v5)), F2: 205}
                                _ = v34
                            }
                            _ = v32
                            _ = v33
                        }
                    case -407:
                        v4 = 42
                    }
                    if v3 {
                        var v35 float32 = 1e10
                        if (!(772 == 1)) {
                            _ = len(v12)
                            v1 -= v2
                            fmt.Println(v2)
                        }
                        if true {
                            var v36 uint = 4294967295
                            _ = v36
                        }
                        _ = v35
                    }
                }
                var v37 uint8 = 255
                _ = v37
            }
            _ = fi13
        }
        if false {
            if false {
                v6 = -1
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
            }
            fmt.Println(v5, v3, v6)
        } else {
            v2 += 45
            v2 -= v2
            var v38 uint8 = 166
            _ = v38
        }
        for fi39 := 0; fi39 < 13; fi39++ {
            v40 := GenS6{F0: v5}
            _ = fi39
            _ = v40
        }
    }
    if (42 == 32767) {
        var v41 float64 = 1e10
        _ = v41
    }
    if v3 {
        for fi42 := 0; fi42 < 18; fi42++ {
            v5 += v5
            _ = fi42
        }
        v6 = (^(v6))
    }
    v5 += 2
    if (false || v3) {
        v43 := GenS1{F0: 0, F1: (v5 < v5), F2: 42}
        v44 := GenS7{F0: 790, F1: v3, F2: 0}
        var v45 int16 = -560
        _ = v43
        _ = v44
        _ = v45
    }
    v46 := GenS4{F0: v5, F1: v1, F2: v1}
    v47 := GenS7{F0: 32767, F1: (!(32767 <= -9223372036854775808)), F2: 127}
    v48 := GenS8{F0: v6}
    fmt.Println(v4)
    var v49 string = "foo bar"
    var v50 int8 = (v2 & int8(v4))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}

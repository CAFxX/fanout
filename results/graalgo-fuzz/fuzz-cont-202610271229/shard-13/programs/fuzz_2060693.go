package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int64 }
type GenS2 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS3 struct { F0 int; F1 uint16; F2 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS6 struct { F0 int8 }
func genFn0(p0 int32, p1 uint32) uint8 {
    fmt.Println(p0, p1)
    v5 := GenS1{F0: 42, F1: (!true), F2: 1}
    v6 := GenS1{F0: 42, F1: (!true), F2: 42}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 string, p1 uint16, p2 int32) int32 {
    var v7 uint8 = 127
    var v8 int16 = -32768
    var v9 uint = 538
    fmt.Println(v7)
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 uint = (v9 / (((v9 + (v9 << (((^(v9))) & 63)))) | 1))
        v12 := GenS2{F0: 66, F1: v11, F2: 0}
        v13 := GenS3{F0: int(v8), F1: 32767, F2: 42}
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return p2
}

func main() {
    var v1 int8 = 42
    var v2 uint32 = 486
    var v3 uint64 = 4294967295
    var v4 uint32 = v2
    v14 := []int{4294967295, 32768, 128, 32768}
    for fi15 := 0; fi15 < 5; fi15++ {
        v16 := GenS3{F0: fi15, F1: 1, F2: 42}
        _ = fi15
        _ = v16
    }
    if len(v14) > 0 {
        v14[0] = 570
    }
    switch -354 {
    case 2147483647:
        fmt.Println(v4)
    case -9223372036854775808:
        switch "a" {
        case "":
            if (v2 < (v2 - (v4 ^ (^(v2))))) {
                if true {
                    v17 := GenS3{F0: 0, F1: 65535, F2: 18446744073709551615}
                    _ = v17
                }
                switch -2147483648 {
                case -9223372036854775808:
                    if false {
                        v1 = (^(v1))
                        v4 = v4
                    }
                case 32767:
                    var v18 int64 = -603
                    _ = v18
                case 32768:
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                default:
                    fmt.Println(v2, v2, v4)
                }
                var v19 bool = true
                _ = v19
            } else {
                for fi20 := 0; fi20 < 15; fi20++ {
                    v1 -= (-(v1))
                    v1 = v1
                    v21 := GenS4{F0: ("foo bar" + ((((("foo bar" + "a") + "\n\t") + "a") + "a") + "\n\t"))}
                    _ = fi20
                    _ = v21
                }
            }
        case "foo bar":
            switch 32767 {
            case -2147483648:
                fmt.Println(v1, v4, v4)
            case 128:
                v4 -= v4
            case 9223372036854775807:
                switch ("a" + "a") {
                case "hello":
                    for fi22 := 0; fi22 < 8; fi22++ {
                        if (true && false) {
                            var v23 int32 = 2147483647
                            _ = v23
                        } else {
                            fmt.Println(v1, v1, fi22)
                        }
                        for fi24 := 0; fi24 < 11; fi24++ {
                            v2 = v4
                            var v25 int32 = -168
                            v26 := GenS5{F0: v3, F1: 539, F2: v25}
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                        switch fi22 {
                        case 32767:
                            var v27 uint = 1
                            _ = v27
                        case -295:
                            v4 += 379
                        case 127:
                            v2 += (v2 * v2)
                        default:
                            for fi28 := 0; fi28 < 20; fi28++ {
                                v3 -= (v3 & (v3 * v3))
                                fmt.Println(fi28)
                                v29 := GenS3{F0: fi22, F1: 32767, F2: 771}
                                _ = fi28
                                _ = v29
                            }
                        }
                        _ = fi22
                    }
                case "\n\t":
                    v1 = v1
                case "foo bar":
                    if true {
                        var v30 uint8 = 42
                        _ = v30
                    } else {
                        v31 := GenS4{F0: "hello"}
                        var v32 uint32 = (v2 & (^(v4)))
                        v33 := GenS6{F0: v1}
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                default:
                    for fi34 := 0; fi34 < 16; fi34++ {
                        var v35 float64 = (float64(1e10) - 1.5)
                        var v36 uint64 = v3
                        var v37 bool = true
                        _ = fi34
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                }
            default:
                v38 := GenS1{F0: v1, F1: true, F2: 42}
                _ = v38
            }
        case "a":
            var v39 uint32 = 4294967295
            _ = v39
        default:
            if true {
                var v40 int = -2147483648
                _ = v40
            } else {
                if (true && false) {
                    v2 = (v4 / ((693) | 1))
                    v41 := GenS3{F0: -156, F1: 27, F2: 42}
                    fmt.Println(v4, v2)
                    _ = v41
                }
                for fi42 := 0; fi42 < 1; fi42++ {
                    fmt.Println(v3)
                    switch (fi42 % ((fi42) | 1)) {
                    case 919:
                        v3 = 4294967295
                    case 128:
                        var v43 int32 = -2147483648
                        _ = v43
                    }
                    if true {
                        if len(v14) > 0 {
                            v14[0] = 32767
                        }
                        v3 = v3
                        switch fi42 {
                        case -563:
                            v14 = append(v14, fi42)
                        case -1:
                            if true {
                                var v44 string = ("\n\t" + "")
                                if len(v14) > 0 {
                                    v14[0] = fi42
                                }
                                _ = v44
                            } else {
                                for fi45 := 0; fi45 < 7; fi45++ {
                                    if len(v14) > 0 {
                                        v14[0] = fi45
                                    }
                                    _ = fi45
                                }
                            }
                        }
                    }
                    _ = fi42
                }
            }
        }
    case 32768:
        if (-9223372036854775808 != -9223372036854775808) {
            var v46 int = 32768
            _ = v46
        }
    }
    v4 = v2
    fmt.Println(v3)
    if false {
        v47 := GenS6{F0: -1}
        _ = v47
    }
    var v48 int16 = int16(v2)
    v49 := GenS3{F0: 530, F1: 20, F2: 265}
    switch 9223372036854775807 {
    case -830:
        fmt.Println(v1, v4)
    case -2147483648:
        switch 1 {
        case 42:
            var v50 float32 = 0.0
            _ = v50
        default:
            v2 = (v2 / (((v4 - v2)) | 1))
        }
    case 4294967295:
        for fi51 := 0; fi51 < 14; fi51++ {
            fmt.Println(fi51)
            _ = len(v14)
            _ = fi51
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 int8; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint8; F1 uint32; F2 uint8 }
type GenS5 struct { F0 int16; F1 int8; F2 uint32 }
type GenS6 struct { F0 uint32; F1 int64 }
type GenS7 struct { F0 int16; F1 int }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 string }
func genFn0(p0 int32, p1 uint64) uint16 {
    var v7 uint64 = 1
    var v8 uint32 = 2147483647
    v9 := []int{42, -861, 127, 127}
    if (-106 == 0) {
        var v10 uint8 = 2
        var v11 int8 = -1
        var v12 bool = true
        _ = v10
        _ = v11
        _ = v12
    } else {
        fmt.Println(v8, v7)
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 string = (("a" + ("a" + "")) + (("a" + "hello") + "foo bar"))
    var v2 float32 = -2.25
    var v3 uint32 = 42
    var v4 int64 = -110
    var v5 string = "hello"
    var v6 uint32 = (v3 / ((v3) | 1))
    if true {
        v6 -= v3
        var v13 int64 = v4
        v14 := []int{-304, -332, -9223372036854775808, 168}
        _ = v13
        _ = v14
    }
    var v15 int8 = -128
    for fi16 := 0; fi16 < 3; fi16++ {
        if ("\n\t" <= v5) {
            if false {
                v17 := GenS1{F0: -55, F1: (v3 / ((588) | 1))}
                _ = v17
            } else {
                v4 = v4
                for fi18 := 0; fi18 < 7; fi18++ {
                    v19 := GenS2{F0: fi16, F1: 127}
                    var v20 uint32 = (v3 << ((0) & 63))
                    switch "a" {
                    case "a":
                        switch "foo bar" {
                        case "a":
                            if false {
                                v21 := GenS3{F0: (-(v15)), F1: 1, F2: 18446744073709551615}
                                _ = v21
                            }
                        }
                    }
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                var v22 uint32 = (^(v3))
                _ = v22
            }
            var v23 uint64 = 4294967295
            for fi24 := 0; fi24 < 3; fi24++ {
                var v25 uint16 = 0
                var v26 uint64 = v23
                _ = fi24
                _ = v25
                _ = v26
            }
            _ = v23
        } else {
            if (2147483647 < -1) {
                switch v5 {
                case "":
                    v15 = (v15 / (((v15 ^ v15)) | 1))
                case "\n\t":
                    v27 := GenS4{F0: 128, F1: (^(v6)), F2: 128}
                    _ = v27
                default:
                    for fi28 := 0; fi28 < 4; fi28++ {
                        if (!false) {
                            var v29 uint8 = 1
                            v30 := []int{fi28, (fi16 & int(v4)), fi28, (-(fi28))}
                            v6 -= v6
                            _ = v29
                            _ = v30
                        } else {
                            v4 = -446
                            v31 := GenS5{F0: 1000, F1: 127, F2: 948}
                            switch ((((v1 + "\n\t") + ("\n\t" + "hello")) + "\n\t") + v5) {
                            case "a":
                                v2 = v2
                            case "\n\t":
                                v1 = v1
                            case "foo bar":
                                var v32 int16 = 0
                                _ = v32
                            }
                            _ = v31
                        }
                        fmt.Println(v15)
                        _ = fi28
                    }
                }
                v33 := GenS2{F0: fi16, F1: 2}
                _ = v33
            }
            switch (v1 + v1) {
            case "foo bar":
                for fi34 := 0; fi34 < 17; fi34++ {
                    v35 := GenS5{F0: 42, F1: (^(v15)), F2: v6}
                    if false {
                        switch v5 {
                        case "a":
                            var v36 float32 = ((1e10 / (v2 * v2)) - (0.0 + (v2 + v2)))
                            _ = v36
                        case "hello":
                            v37 := GenS6{F0: (v6 ^ v6), F1: 14}
                            _ = v37
                        case "":
                            fmt.Println(v1, v5, v2)
                        }
                    }
                    _ = fi34
                    _ = v35
                }
            case "":
                v2 = v2
            case "a":
                if false {
                    for fi38 := 0; fi38 < 2; fi38++ {
                        v39 := GenS3{F0: (v15 + (-(v15))), F1: v4, F2: 1}
                        v40 := GenS7{F0: -240, F1: fi38}
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                }
            default:
                v41 := []int{fi16, (fi16 << ((255) & 63))}
                _ = v41
            }
        }
        v6 -= v3
        if (!false) {
            v2 = (v2 + 1.5)
            if (!(1000 < 1)) {
                for fi42 := 0; fi42 < 5; fi42++ {
                    fmt.Println(fi42, v2)
                    v43 := []int{fi16, (fi42 ^ fi16)}
                    _ = fi42
                    _ = v43
                }
                switch 32767 {
                case 0:
                    v44 := []int{}
                    _ = v44
                }
            } else {
                for fi45 := 0; fi45 < 8; fi45++ {
                    v46 := []int{}
                    if len(v46) > 0 {
                        v46[0] = (fi16 - (fi16 >> ((4294967295) & 63)))
                    }
                    _ = fi45
                    _ = v46
                }
                switch -9223372036854775808 {
                case 32768:
                    for fi47 := 0; fi47 < 5; fi47++ {
                        if true {
                            v48 := GenS3{F0: v15, F1: 761, F2: 1}
                            var v49 uint64 = 955
                            v3 = v3
                            switch (fi16 >> ((4294967295) & 63)) {
                            case -702:
                                v50 := []int{(fi47 << ((65535) & 63)), fi16, 514}
                                _ = v50
                            case 2147483647:
                                v51 := []int{fi47, (fi16 >> ((65535) & 63))}
                                _ = v51
                            }
                            fmt.Println(v5)
                            _ = v48
                            _ = v49
                        } else {
                            if true {
                                v1 = ((v5 + ("\n\t" + v5)) + "hello")
                                v52 := []int{int(v15), fi16, 2147483647, (^(fi16))}
                                if len(v52) >= 2 {
                                    v52 = v52[1:len(v52)]
                                }
                                _ = v52
                            }
                        }
                        if false {
                            v6 = v3
                        }
                        switch "hello" {
                        case "\n\t":
                            v1 = "a"
                        }
                        _ = fi47
                    }
                case 0:
                    v1 = ("" + v1)
                case 586:
                    v5 = v5
                default:
                    v53 := GenS8{F0: 32767}
                    _ = v53
                }
                switch fi16 {
                case 0:
                    var v54 uint16 = genFn0(365, 990)
                    _ = v54
                case -70:
                    v55 := []int{fi16, fi16}
                    _ = v55
                }
            }
            switch "a" {
            case "\n\t":
                var v56 int64 = (v4 - (-(v4)))
                _ = v56
            default:
                var v57 bool = true
                _ = v57
            }
        }
        _ = fi16
    }
    v5 = ((v1 + v1) + "a")
    var v58 uint32 = (v3 + (v6 - 0))
    for fi59 := 0; fi59 < 3; fi59++ {
        v60 := GenS9{F0: (((v1 + ("\n\t" + "foo bar")) + ("" + v1)) + (v1 + (("\n\t" + v5) + v1)))}
        v4 -= 1
        fmt.Println(v5)
        _ = fi59
        _ = v60
    }
    for fi61 := 0; fi61 < 2; fi61++ {
        var v62 float64 = float64(3.14159)
        _ = fi61
        _ = v62
    }
    switch -1 {
    case -630:
        var v63 string = "\n\t"
        _ = v63
    }
    if true {
        v58 = v6
    }
    var v64 float32 = (((v2 * v2) * 1.5) / (v2 * (0.1 - (-2.25 + ((1.5 + (v2 - v2)) + (v2 - ((v2 - v2) / (v2 + v2))))))))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
    _ = v64
}

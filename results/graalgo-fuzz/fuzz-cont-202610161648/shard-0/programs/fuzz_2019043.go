package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int16; F1 bool; F2 uint16 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32; F1 uint; F2 uint64 }
type GenS7 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS8 struct { F0 int64; F1 int64 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int }

func main() {
    var v1 int16 = 0
    v2 := []uint32{929, 2, 428}
    v3 := make(map[string]bool)
    var v4 int8 = -128
    fmt.Println(v4)
    v5 := GenS1{F0: v1, F1: 1}
    v6 := GenS2{F0: 0, F1: (v1 & (v1 % ((506) | 1)))}
    fmt.Println(v1)
    for fi7 := 0; fi7 < 11; fi7++ {
        if (127 < 97) {
            if ((!(!(!true))) && (42 == 1)) {
                for fi8 := 0; fi8 < 7; fi8++ {
                    v9 := GenS2{F0: (-(v4)), F1: (-(v1))}
                    _ = fi8
                    _ = v9
                }
                if true {
                    fmt.Println(v1, v4)
                }
            }
            for fi10 := 0; fi10 < 6; fi10++ {
                if (!(419 <= 4294967295)) {
                    var v11 uint16 = 65535
                    _ = v11
                }
                v4 -= 0
                _ = fi10
            }
            v12 := GenS3{F0: (v1 ^ (v1 & v1)), F1: (42 > 42), F2: 2}
            _ = v12
        }
        v2 = append(v2, 600)
        _ = fi7
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 uint16 = 32767
            fmt.Println(fi13, v15)
            switch (^(fi14)) {
            case 32768:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 583:
                switch (fi14 ^ (fi13 + (^(fi13)))) {
                case -177:
                    v1 = v1
                default:
                    var v16 int = 1
                    _ = v16
                }
            }
            _ = fi14
            _ = v15
        }
        fmt.Println(v1, v1)
        _ = fi13
    }
    fmt.Println(v1, v4)
    switch 127 {
    case -895:
        for fi17 := 0; fi17 < 20; fi17++ {
            v18 := GenS4{F0: 2147483647, F1: 4294967295}
            _ = fi17
            _ = v18
        }
    case 127:
        var v19 float64 = 1.5
        _ = v19
    case 180:
        v20 := GenS3{F0: 32767, F1: (!(false && (1e10 != 1e10))), F2: 2}
        _ = v20
    default:
        for fi21 := 0; fi21 < 18; fi21++ {
            switch "\n\t" {
            case "\n\t":
                if true {
                    if true {
                        v1 -= (v1 * v1)
                        switch "hello" {
                        case "foo bar":
                            if true {
                                v4 = (-(v4))
                            }
                        default:
                            v22 := GenS5{F0: 0}
                            _ = v22
                        }
                        fmt.Println(v4)
                    }
                    _ = len(v3)
                }
            case "foo bar":
                var v23 int8 = 100
                _ = v23
            case "":
                v24 := GenS6{F0: 1, F1: 0, F2: 891}
                _ = v24
            }
            v1 = v1
            if (!false) {
                if (!false) {
                    if false {
                        for fi25 := 0; fi25 < 15; fi25++ {
                            var v26 uint16 = 152
                            _ = fi25
                            _ = v26
                        }
                        v1 = (v1 - 0)
                        var v27 uint = 533
                        _ = v27
                    } else {
                        switch 42 {
                        case -9223372036854775808:
                            fmt.Println(v1)
                        case -143:
                            var v28 bool = false
                            _ = v28
                        default:
                            v4 = v4
                        }
                        var v29 int16 = 568
                        fmt.Println(v4, v29, v4)
                        _ = v29
                    }
                    switch "a" {
                    case "hello":
                        for fi30 := 0; fi30 < 10; fi30++ {
                            v4 = v4
                            fmt.Println(v1, fi21)
                            _ = len(v2)
                            _ = fi30
                        }
                    case "":
                        fmt.Println(v1, fi21)
                    }
                } else {
                    v31 := GenS7{F0: 2147483648, F1: 184, F2: -1}
                    v1 += v1
                    _ = v31
                }
            } else {
                if (2 > 745) {
                    if true {
                        fmt.Println(v4)
                    }
                    v1 -= (^(v1))
                }
                if false {
                    for fi32 := 0; fi32 < 19; fi32++ {
                        v33 := GenS1{F0: v1, F1: -1}
                        if true {
                            v34 := GenS8{F0: 42, F1: 42}
                            var v35 int = (fi32 % ((-362) | 1))
                            _ = v34
                            _ = v35
                        }
                        var v36 int64 = 9223372036854775807
                        _ = fi32
                        _ = v33
                        _ = v36
                    }
                }
                switch fi21 {
                case 34:
                    if false {
                        for fi37 := 0; fi37 < 11; fi37++ {
                            var v38 int32 = -450
                            _ = fi37
                            _ = v38
                        }
                        v39 := GenS9{F0: -862}
                        _ = v39
                    }
                case 845:
                    v1 = (^(v1))
                case 481:
                    var v40 int8 = v4
                    _ = v40
                default:
                    v1 = v1
                }
            }
            _ = fi21
        }
    }
    for fi41 := 0; fi41 < 7; fi41++ {
        var v42 uint16 = 1
        _ = fi41
        _ = v42
    }
    v4 = v4
    v43 := GenS10{F0: int(v4)}
    var v44 int8 = v4
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v45 := GenS1{F0: (v1 ^ (v1 * v1)), F1: -2147483648}
        _ = v45
    }
    if true {
        if true {
            var v46 int32 = -221
            _ = v46
        } else {
            v1 += (^(v1))
            v44 = v4
            for fi47 := 0; fi47 < 5; fi47++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi47
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}

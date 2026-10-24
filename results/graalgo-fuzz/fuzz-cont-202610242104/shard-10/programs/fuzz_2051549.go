package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint8; F1 int8; F2 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 bool; F2 uint8 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 int) int8 {
    p0 -= (^(p0))
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 string = "foo bar"
        v8 := []int{(p0 & fi6)}
        var v9 int = 1
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := GenS1{F0: 65535, F1: 419}
        _ = fi10
        _ = v11
    }
    var v12 uint64 = 1
    _ = p0
    _ = v12
    return -1
}
func genFn1(p0 string, p1 uint) uint16 {
    v13 := []int{32767}
    var v14 int64 = -198
    v13 = append(v13, 625)
    p1 = 0
    switch -767 {
    case -14:
        var v15 uint = 2
        _ = v15
    case 1:
        var v16 float32 = 0.0
        _ = v16
    default:
        fmt.Println(p1, v14, p0)
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 255
}

func main() {
    var v1 string = ("\n\t" + "hello")
    v2 := []int{-2147483648, -2147483648}
    var v3 int16 = -1
    var v4 int8 = -128
    var v5 int64 = -1
    if true {
        fmt.Println(v4, v5)
    } else {
        for fi17 := 0; fi17 < 15; fi17++ {
            for fi18 := 0; fi18 < 10; fi18++ {
                if true {
                    v4 = (v4 + 79)
                } else {
                    fmt.Println(v5, fi17, v4)
                    v5 = v5
                    if true {
                        var v19 int16 = 0
                        _ = v19
                    }
                }
                v5 = v5
                v3 += v3
                fmt.Println(v1)
                _ = fi18
            }
            _ = fi17
        }
        var v20 int32 = -214
        _ = v20
    }
    var v21 uint16 = 98
    v1 = (v1 + "hello")
    fmt.Println(v3)
    if (!true) {
        switch v1 {
        case "\n\t":
            v4 -= v4
        default:
            if (488 < 0) {
                var v22 int = 696
                var v23 uint8 = 0
                _ = v22
                _ = v23
            } else {
                if true {
                    if len(v2) > 0 {
                        v2[0] = -915
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    v5 = (v5 & v5)
                    v4 = (v4 + 42)
                }
            }
        }
        switch -278 {
        case -2147483648:
            if len(v2) > 0 {
                v2[0] = 0
            }
        case 9223372036854775807:
            var v24 int = int(v4)
            _ = v24
        case 1:
            var v25 uint64 = 918
            _ = v25
        default:
            switch 9223372036854775807 {
            case -676:
                v21 += 2
            default:
                for fi26 := 0; fi26 < 9; fi26++ {
                    var v27 uint64 = 220
                    fmt.Println(v3)
                    _ = fi26
                    _ = v27
                }
            }
        }
    } else {
        for fi28 := 0; fi28 < 4; fi28++ {
            fmt.Println(v5, v3, v1)
            _ = fi28
        }
        for fi29 := 0; fi29 < 8; fi29++ {
            if true {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                for fi30 := 0; fi30 < 10; fi30++ {
                    if (!(!true)) {
                        v4 -= (v4 | v4)
                    }
                    _ = fi30
                }
            }
            v31 := GenS2{F0: 0, F1: (v4 % ((v4) | 1)), F2: 0}
            _ = fi29
            _ = v31
        }
    }
    v2 = append(v2, 0)
    v32 := GenS1{F0: (v21 << ((1) & 63)), F1: 42}
    switch v1 {
    case "":
        if true {
            for fi33 := 0; fi33 < 3; fi33++ {
                fmt.Println(fi33, v21)
                v3 = v3
                if false {
                    for fi34 := 0; fi34 < 19; fi34++ {
                        if (v1 >= v1) {
                            switch v1 {
                            case "hello":
                                v4 += -128
                            case "foo bar":
                                var v35 int64 = v5
                                _ = v35
                            }
                            fmt.Println(v21)
                        }
                        v36 := GenS1{F0: (v21 + genFn1(v1, 255)), F1: -1}
                        v37 := GenS3{F0: int32(v3)}
                        v4 -= v4
                        _ = fi34
                        _ = v36
                        _ = v37
                    }
                    var v38 bool = true
                    _ = v38
                } else {
                    if false {
                        v4 = 32
                    }
                }
                _ = fi33
            }
        }
    default:
        if ((2 <= 65535) || (false || (!(!(false || (false && false)))))) {
            v39 := GenS4{F0: 0, F1: false, F2: 46}
            _ = v39
        }
    }
    v40 := GenS5{F0: 245}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v32
    _ = v4
    _ = v40
    _ = v5
}

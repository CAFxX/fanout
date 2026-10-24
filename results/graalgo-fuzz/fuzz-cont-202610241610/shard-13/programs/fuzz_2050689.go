package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int8; F1 int; F2 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 int32; F2 uint32 }
func genFn0() uint8 {
    if true {
        v4 := GenS1{F0: ("foo bar" + ("foo bar" + "\n\t")), F1: -56}
        _ = v4
    } else {
        var v5 int64 = -1
        _ = v5
    }
    switch 42 {
    case -759:
        var v6 int8 = 29
        _ = v6
    case 32767:
        var v7 uint16 = 1
        _ = v7
    default:
        switch "a" {
        case "":
            var v8 float32 = (-2.25 + 0.1)
            _ = v8
        case "\n\t":
            var v9 uint32 = 4294967295
            _ = v9
        default:
            switch -1 {
            case -9223372036854775808:
                var v10 int = 9223372036854775807
                _ = v10
            case 747:
                var v11 int32 = -552
                _ = v11
            case -653:
                if (4294967295 != 750) {
                    var v12 bool = (9223372036854775807 == 330)
                    v13 := []int{4294967295, -9223372036854775808, -214}
                    _ = v12
                    _ = v13
                } else {
                    var v14 int8 = -1
                    v14 -= 1
                    _ = v14
                }
            default:
                if false {
                    for fi15 := 0; fi15 < 5; fi15++ {
                        var v16 uint32 = 4294967295
                        if false {
                            v16 -= (v16 << ((18446744073709551615) & 63))
                        } else {
                            if (!false) {
                                if (!true) {
                                    var v17 bool = (("foo bar" == "\n\t") || true)
                                    var v18 uint16 = 65535
                                    _ = v17
                                    _ = v18
                                } else {
                                    for fi19 := 0; fi19 < 12; fi19++ {
                                        var v20 int64 = 26
                                        _ = fi19
                                        _ = v20
                                    }
                                }
                            }
                            v21 := GenS2{F0: 40, F1: 32768, F2: 72}
                            v16 = (^(v16))
                            _ = v21
                        }
                        v22 := GenS1{F0: ("\n\t" + ("hello" + "\n\t")), F1: 0}
                        _ = fi15
                        _ = v16
                        _ = v22
                    }
                    switch "" {
                    case "":
                        var v23 int = 127
                        _ = v23
                    }
                } else {
                    var v24 int = 32767
                    var v25 int16 = 1
                    v24 -= -198
                    _ = v24
                    _ = v25
                }
            }
        }
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        var v27 float64 = 1.5
        for fi28 := 0; fi28 < 13; fi28++ {
            if ((false && (523 <= 174)) || (!(!false))) {
                var v29 bool = false
                _ = v29
            }
            v30 := []int{-9223372036854775808, fi28, fi26, -1}
            v31 := GenS3{F0: false}
            _ = fi28
            _ = v30
            _ = v31
        }
        _ = fi26
        _ = v27
    }
    v32 := GenS1{F0: "", F1: -642}
    _ = v32
    return 28
}
func genFn1(p0 uint16) uint16 {
    for fi33 := 0; fi33 < 4; fi33++ {
        fmt.Println(p0, p0)
        _ = fi33
    }
    p0 = 32767
    _ = p0
    return p0
}

func main() {
    v1 := []string{"\n\t"}
    v2 := []int{2147483647, 32768}
    var v3 uint64 = 442
    if len(v1) > 0 {
        v1[0] = (((("a" + "hello") + "a") + "") + (("\n\t" + (("a" + "\n\t") + "foo bar")) + "hello"))
    }
    for fi34 := 0; fi34 < 11; fi34++ {
        v3 += (^(v3))
        if (65535 != 2) {
            for fi35 := 0; fi35 < 20; fi35++ {
                v3 += v3
                v3 = (v3 % ((v3) | 1))
                _ = fi35
            }
            for fi36 := 0; fi36 < 14; fi36++ {
                v37 := GenS3{F0: true}
                v38 := GenS3{F0: true}
                _ = fi36
                _ = v37
                _ = v38
            }
            v39 := GenS1{F0: "foo bar", F1: 0}
            _ = v39
        }
        var v40 uint = 42
        _ = fi34
        _ = v40
    }
    v3 -= v3
    switch ("a" + "") {
    case "foo bar":
        var v41 int8 = -128
        _ = v41
    case "hello":
        var v42 bool = (!(32767 > 9223372036854775807))
        _ = v42
    default:
        var v43 bool = (!false)
        _ = v43
    }
    v3 += 2
    if true {
        for fi44 := 0; fi44 < 20; fi44++ {
            v3 += (v3 % (((v3 + (v3 - v3))) | 1))
            _ = fi44
        }
    } else {
        var v45 uint8 = 250
        _ = v45
    }
    var v46 int = 32768
    v46 = (^(v46))
    var v47 bool = (!true)
    v46 -= v46
    v48 := GenS4{F0: 110, F1: -34, F2: 0}
    v49 := GenS3{F0: v47}
    switch "" {
    case "hello":
        if len(v2) > 0 {
            v2[0] = -328
        }
    case "\n\t":
        v47 = v47
    case "a":
        for fi50 := 0; fi50 < 3; fi50++ {
            v3 = v3
            _ = fi50
        }
    }
    v3 -= v3
    v47 = true
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v46
    _ = v47
    _ = v48
    _ = v49
}

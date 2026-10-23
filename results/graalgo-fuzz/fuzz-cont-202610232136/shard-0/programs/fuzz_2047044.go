package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int8 }
type GenS2 struct { F0 int; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 int; F1 int32; F2 uint }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 int16; F1 uint; F2 string }
func genFn0(p0 int64, p1 int) int16 {
    v7 := []int{(p1 << ((910) & 63)), -893, 127, p1}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v7
    return -1
}
func genFn1(p0 uint64, p1 uint32) int16 {
    if (!true) {
        v8 := []int{306}
        var v9 bool = false
        _ = v8
        _ = v9
    } else {
        fmt.Println(p0, p1)
        p1 += (^(p1))
        v10 := GenS1{F0: 18446744073709551615, F1: -1, F2: 1}
        _ = v10
    }
    p0 -= p0
    p1 = p1
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 uint16 = uint16(p1)
        var v13 uint64 = (p0 & (^(p0)))
        fmt.Println(p1, v13)
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := []int{2147483647, 1, 9223372036854775807, 32767}
    _ = p0
    _ = p1
    _ = v14
    return 1
}

func main() {
    var v1 uint32 = 0
    var v2 uint16 = 65535
    var v3 int16 = -1
    var v4 uint16 = 32767
    v5 := []uint32{(^(v1)), v1, 2147483647, v1}
    var v6 int8 = -20
    var v15 int = 32768
    v15 = (v15 / ((int(v2)) | 1))
    if false {
        var v16 uint8 = 42
        _ = v16
    } else {
        fmt.Println(v1, v3, v6)
    }
    v2 = 32767
    for fi17 := 0; fi17 < 15; fi17++ {
        for fi18 := 0; fi18 < 14; fi18++ {
            v19 := GenS2{F0: v15, F1: -1, F2: (^(v1))}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v20 float64 = ((1.5 + 0.1) / 0.1)
            _ = fi18
            _ = v19
            _ = v20
        }
        switch v15 {
        case 4294967295:
            v21 := GenS2{F0: v15, F1: (^(v6)), F2: (v1 * v1)}
            _ = v21
        case 32767:
            if false {
                _ = len(v5)
                if (65535 < 438) {
                    fmt.Println(v15, v3)
                    v4 = v2
                } else {
                    if false {
                        if (v3 != (v3 % (((v3 << ((65535) & 63))) | 1))) {
                            var v22 float64 = 0.1
                            v23 := GenS2{F0: (fi17 | v15), F1: 42, F2: v1}
                            for fi24 := 0; fi24 < 12; fi24++ {
                                fmt.Println(fi17)
                                _ = fi24
                            }
                            _ = v22
                            _ = v23
                        }
                    }
                    if true {
                        fmt.Println(v6)
                        for fi25 := 0; fi25 < 17; fi25++ {
                            var v26 bool = true
                            v5 = append(v5, 65535)
                            _ = fi25
                            _ = v26
                        }
                    } else {
                        for fi27 := 0; fi27 < 1; fi27++ {
                            v28 := GenS3{F0: 18446744073709551615, F1: (v1 | v1)}
                            if (false || false) {
                                v29 := GenS1{F0: 18446744073709551615, F1: (v3 & -115), F2: (v6 >> ((1) & 63))}
                                _ = v29
                            }
                            var v30 uint = 65535
                            _ = fi27
                            _ = v28
                            _ = v30
                        }
                        var v31 string = ("foo bar" + "\n\t")
                        v2 = (v2 >> ((2) & 63))
                        _ = v31
                    }
                }
            }
        default:
            v2 = v4
        }
        switch "a" {
        case "":
            if true {
                fmt.Println(v1, v6, fi17)
                v32 := GenS4{F0: int(v3), F1: 0, F2: 18446744073709551615}
                for fi33 := 0; fi33 < 20; fi33++ {
                    fmt.Println(v4, v1)
                    var v34 uint8 = 170
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi33
                    _ = v34
                }
                _ = v32
            }
        case "a":
            fmt.Println(v15)
        default:
            v3 += 910
        }
        _ = fi17
    }
    fmt.Println(v1, v3)
    var v35 int8 = v6
    v6 = v35
    switch 127 {
    case 904:
        var v36 string = "a"
        _ = v36
    default:
        v6 -= 100
    }
    if ((uint64(v4) >= 1) || (!true)) {
        v37 := GenS1{F0: 914, F1: 749, F2: int8(v15)}
        fmt.Println(v1, v6, v1)
        _ = v37
    }
    var v38 uint16 = (v2 & 1)
    for fi39 := 0; fi39 < 9; fi39++ {
        var v40 int8 = (-(v6))
        for fi41 := 0; fi41 < 1; fi41++ {
            v42 := GenS5{F0: 2, F1: true}
            _ = fi41
            _ = v42
        }
        v2 += (^(v38))
        _ = fi39
        _ = v40
    }
    v35 = 127
    switch (v15 / ((v15) | 1)) {
    case -853:
        fmt.Println(v15)
    case 127:
        for fi43 := 0; fi43 < 13; fi43++ {
            switch (fi43 - v15) {
            case -103:
                for fi44 := 0; fi44 < 19; fi44++ {
                    v45 := GenS6{F0: genFn0(1, fi43), F1: 0, F2: ("\n\t" + (("hello" + "foo bar") + "\n\t"))}
                    _ = fi44
                    _ = v45
                }
            case 4294967295:
                var v46 int8 = int8(v2)
                _ = v46
            case 423:
                v38 -= (v4 & v4)
            default:
                v2 -= v2
            }
            switch (fi43 / (((fi43 / ((fi43) | 1))) | 1)) {
            case -117:
                v15 += (fi43 | v15)
            case 9223372036854775807:
                var v47 uint32 = 2147483647
                _ = v47
            default:
                if (true || false) {
                    fmt.Println(v15, v4, v15)
                } else {
                    var v48 uint = 65535
                    _ = v48
                }
            }
            _ = fi43
        }
    }
    if len(v5) > 0 {
        v5[0] = (v1 - (v1 + (v1 >> ((65535) & 63))))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}

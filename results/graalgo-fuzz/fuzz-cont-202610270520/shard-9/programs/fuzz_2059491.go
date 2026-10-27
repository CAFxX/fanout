package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 int32; F2 uint16 }
type GenS4 struct { F0 int16; F1 int64; F2 string }
type GenS5 struct { F0 int; F1 uint8; F2 uint64 }
type GenS6 struct { F0 uint16; F1 uint64 }
type GenS7 struct { F0 int32; F1 uint8; F2 uint }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 string }
func genFn0(p0 bool) uint16 {
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := []int{}
        for fi8 := 0; fi8 < 13; fi8++ {
            if ("\n\t" == "") {
                var v9 int8 = 42
                v10 := GenS1{F0: 0}
                fmt.Println(p0, fi6, v9)
                _ = v10
                _ = v9
            } else {
                if (542 <= 4294967295) {
                    if false {
                        p0 = ((fi6 / (((^(fi6))) | 1)) == (fi8 - (fi6 - -157)))
                        p0 = p0
                    }
                    if p0 {
                        p0 = (-926 <= 409)
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        if p0 {
                            v7 = append(v7, (fi8 | (fi8 + 372)))
                        }
                    }
                    var v11 int32 = 1
                    _ = v11
                } else {
                    fmt.Println(fi6, fi8, fi6)
                    if (!(p0 && (!(!((((float32(3.14159) / -2.25) / -2.25) <= -2.25) && true))))) {
                        p0 = (fi8 > 32768)
                        v12 := GenS2{F0: 860}
                        _ = v12
                    } else {
                        fmt.Println(fi8, fi6)
                    }
                    v13 := GenS3{F0: 42, F1: -2147483648, F2: 65535}
                    _ = v13
                }
                var v14 int = (fi6 >> ((357) & 63))
                fmt.Println(fi6, v14)
                _ = v14
            }
            v15 := GenS4{F0: 32767, F1: -493, F2: ("a" + "hello")}
            _ = fi8
            _ = v15
        }
        _ = fi6
        _ = v7
    }
    v16 := GenS5{F0: 32768, F1: 42, F2: 260}
    v17 := GenS6{F0: 2, F1: 639}
    var v18 float64 = ((0.1 * 0.1) - 0.0)
    var v19 int = 127
    _ = p0
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    return 65535
}

func main() {
    var v1 bool = true
    var v2 uint32 = 2147483647
    var v3 uint32 = (v2 << ((1) & 63))
    var v4 int16 = 42
    var v5 uint16 = 2
    fmt.Println(v1, v3, v4)
    v20 := GenS7{F0: -1, F1: 131, F2: 988}
    v21 := []int{}
    fmt.Println(v1)
    v22 := GenS3{F0: 438, F1: 797, F2: v5}
    for fi23 := 0; fi23 < 5; fi23++ {
        var v24 int = 0
        _ = fi23
        _ = v24
    }
    if v1 {
        v25 := GenS8{F0: 42}
        var v26 int = 524
        _ = v25
        _ = v26
    }
    switch ("\n\t" + ("foo bar" + ("foo bar" + ("foo bar" + ("" + "\n\t"))))) {
    case "":
        var v27 uint8 = 128
        _ = v27
    case "\n\t":
        if v1 {
            v28 := GenS2{F0: 205}
            var v29 uint32 = v2
            for fi30 := 0; fi30 < 3; fi30++ {
                v3 = 454
                _ = fi30
            }
            _ = v28
            _ = v29
        } else {
            var v31 float32 = 3.14159
            if (140 > 2) {
                fmt.Println(v2, v4, v4)
                switch -417 {
                case 42:
                    var v32 bool = (255 > 65535)
                    _ = v32
                default:
                    fmt.Println(v4, v2)
                }
            } else {
                if len(v21) > 0 {
                    v21[0] = 1
                }
            }
            _ = v31
        }
    }
    var v33 uint16 = (v5 ^ v5)
    var v34 uint8 = 56
    if (42 >= uint(v34)) {
        switch "foo bar" {
        case "":
            if v1 {
                v5 -= (v5 + (^(v33)))
                if v1 {
                    var v35 uint32 = v3
                    var v36 int64 = 42
                    _ = v35
                    _ = v36
                }
                var v37 int64 = -429
                _ = v37
            } else {
                fmt.Println(v3, v33, v3)
                switch 0 {
                case 32767:
                    _ = len(v21)
                case 0:
                    if len(v21) > 0 {
                        v21[0] = -9223372036854775808
                    }
                case 1:
                    v3 -= (v3 % ((2) | 1))
                default:
                    v34 += v34
                }
                fmt.Println(v3)
            }
        case "\n\t":
            var v38 uint64 = 4294967295
            _ = v38
        case "hello":
            v39 := GenS2{F0: 0}
            _ = v39
        }
        switch 0 {
        case 9223372036854775807:
            v40 := GenS9{F0: (("\n\t" + "\n\t") + "")}
            _ = v40
        case 892:
            v41 := GenS4{F0: v4, F1: -9223372036854775808, F2: "hello"}
            _ = v41
        case 2147483647:
            var v42 uint32 = (v3 + v3)
            _ = v42
        }
        fmt.Println(v4, v3)
    }
    fmt.Println(v34, v3, v2)
    fmt.Println(v33, v5)
    switch 16 {
    case 32768:
        var v43 uint32 = v3
        _ = v43
    default:
        var v44 bool = (!(220 == (^(v34))))
        _ = v44
    }
    fmt.Println(v34, v1)
    v3 = 0
    var v45 int32 = -2147483648
    if v1 {
        for fi46 := 0; fi46 < 19; fi46++ {
            v5 = (^(v5))
            _ = fi46
        }
        v34 = v34
    } else {
        v1 = (!(1.5 == (0.0 / 3.14159)))
        switch ((("\n\t" + "\n\t") + (("foo bar" + "hello") + "a")) + ("foo bar" + "\n\t")) {
        case "a":
            fmt.Println(v2, v45)
        case "\n\t":
            var v47 float64 = 0.0
            _ = v47
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v45
    _ = v5
}

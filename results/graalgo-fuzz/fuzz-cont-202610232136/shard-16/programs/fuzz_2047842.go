package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int64; F1 int32; F2 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 uint16; F1 uint64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint8; F1 uint32 }
func genFn0() int64 {
    var v6 uint32 = 0
    v7 := GenS1{F0: 127, F1: (!false)}
    _ = v6
    _ = v7
    return -453
}
func genFn1(p0 int16, p1 int16, p2 uint16) int8 {
    p0 += p0
    p2 = (p2 + p2)
    var v8 uint64 = 527
    switch "\n\t" {
    case "foo bar":
        v9 := []int{4294967295, -9223372036854775808, 32767, 32767}
        _ = v9
    case "a":
        v10 := GenS1{F0: -2147483648, F1: true}
        _ = v10
    default:
        p0 -= p1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 1
}

func main() {
    var v1 int16 = -32768
    var v2 int16 = (v1 * v1)
    var v3 float32 = (1e10 - 3.14159)
    v4 := []string{"\n\t", ("a" + ("a" + "\n\t")), ("hello" + (("hello" + "") + ("\n\t" + (("\n\t" + (("hello" + ("a" + "")) + (("foo bar" + "a") + ""))) + "\n\t")))), "\n\t"}
    var v5 int32 = 336
    var v11 uint8 = 42
    if false {
        v11 += (v11 + 0)
        v12 := GenS1{F0: 128, F1: false}
        var v13 string = "a"
        _ = v12
        _ = v13
    }
    if ((42 > 699) || true) {
        v14 := GenS1{F0: 98, F1: (!(925 <= 18446744073709551615))}
        switch 1 {
        case 128:
            if len(v4) > 0 {
                v4[0] = ("\n\t" + ("foo bar" + "a"))
            }
        }
        v2 -= v2
        _ = v14
    } else {
        v15 := GenS2{F0: -9223372036854775808, F1: (v5 - 1), F2: 837}
        _ = v15
    }
    if true {
        for fi16 := 0; fi16 < 9; fi16++ {
            if (false || (0.1 >= 0.1)) {
                if (!(((false || false) && (false && false)) && true)) {
                    if ((42 <= -175) || true) {
                        v17 := GenS1{F0: (fi16 * (fi16 / (((fi16 & fi16)) | 1))), F1: (!true)}
                        if (!(128 <= (^(v11)))) {
                            v11 -= (v11 << ((4294967295) & 63))
                            var v18 uint64 = 18446744073709551615
                            _ = v18
                        } else {
                            v19 := GenS3{F0: 42}
                            switch "\n\t" {
                            case "foo bar":
                                fmt.Println(v2, v2, v3)
                            case "\n\t":
                                fmt.Println(fi16)
                            case "":
                                fmt.Println(v1, v2)
                            }
                            _ = v19
                        }
                        _ = v17
                    }
                    if len(v4) > 0 {
                        v4[0] = ("\n\t" + "hello")
                    }
                }
                for fi20 := 0; fi20 < 18; fi20++ {
                    switch fi20 {
                    case 32768:
                        var v21 bool = false
                        _ = v21
                    }
                    v22 := GenS1{F0: (fi16 << ((0) & 63)), F1: true}
                    _ = fi20
                    _ = v22
                }
                for fi23 := 0; fi23 < 4; fi23++ {
                    v24 := GenS4{F0: 618, F1: -116}
                    var v25 int16 = v1
                    v3 = v3
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
            switch fi16 {
            case -193:
                for fi26 := 0; fi26 < 16; fi26++ {
                    if (32767 > 42) {
                        fmt.Println(fi16, fi26)
                        switch (("foo bar" + "foo bar") + "") {
                        case "foo bar":
                            fmt.Println(fi16, v3)
                        case "hello":
                            v11 = (v11 + (v11 & (v11 ^ v11)))
                        case "\n\t":
                            _ = len(v4)
                        }
                        fmt.Println(fi16)
                    } else {
                        var v27 uint64 = 0
                        _ = v27
                    }
                    if false {
                        v11 += 0
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    v1 = v1
                    _ = fi26
                }
            case 275:
                v2 = v1
            case 0:
                for fi28 := 0; fi28 < 8; fi28++ {
                    fmt.Println(v1)
                    var v29 uint64 = 18446744073709551615
                    _ = fi28
                    _ = v29
                }
            default:
                fmt.Println(fi16, fi16, v5)
            }
            v5 += (v5 >> ((uint(v11)) & 63))
            _ = fi16
        }
    }
    v30 := GenS5{F0: 594, F1: 0}
    v2 = int16(v11)
    var v31 int8 = -62
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1 = (v1 - -892)
    v2 += v1
    v3 = 1e10
    v11 = 255
    fmt.Println(v3, v2, v1)
    if true {
        var v32 uint32 = 52
        if true {
            v3 = v3
            v33 := GenS6{F0: 9223372036854775807}
            v32 = v32
            _ = v33
        }
        _ = v32
    } else {
        var v34 int64 = -1
        for fi35 := 0; fi35 < 12; fi35++ {
            if (2 >= 1) {
                v5 += (v5 / (((-(v5))) | 1))
            }
            var v36 uint8 = (v11 >> ((2) & 63))
            fmt.Println(v1, v34)
            _ = fi35
            _ = v36
        }
        _ = len(v4)
        _ = v34
    }
    switch 219 {
    case 127:
        for fi37 := 0; fi37 < 9; fi37++ {
            if (v3 != 0.0) {
                v38 := GenS7{F0: 2}
                v5 = (v5 ^ v5)
                _ = v38
            } else {
                var v39 uint32 = 1
                switch (^(fi37)) {
                case 127:
                    switch "a" {
                    case "":
                        if true {
                            if false {
                                v3 += v3
                                v2 = v2
                            }
                            for fi40 := 0; fi40 < 16; fi40++ {
                                v41 := GenS6{F0: -9223372036854775808}
                                _ = fi40
                                _ = v41
                            }
                        } else {
                            var v42 int64 = genFn0()
                            v39 = v39
                            _ = v42
                        }
                    case "hello":
                        var v43 int16 = v2
                        _ = v43
                    default:
                        v4 = append(v4, "hello")
                    }
                case 185:
                    v44 := GenS3{F0: (v5 / ((v5) | 1))}
                    _ = v44
                case 42:
                    v3 += ((v3 - (v3 / ((v3 + 1e10) + (v3 * (1e10 * (0.0 * 1e10)))))) * 0.0)
                default:
                    v45 := GenS1{F0: fi37, F1: (1 > 607)}
                    _ = v45
                }
                fmt.Println(v39, v5)
                _ = v39
            }
            _ = len(v4)
            v2 = (v1 + v1)
            _ = fi37
        }
    default:
        for fi46 := 0; fi46 < 19; fi46++ {
            v1 = 45
            v11 = (v11 % ((v11) | 1))
            v47 := GenS8{F0: 2, F1: 2}
            _ = fi46
            _ = v47
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}

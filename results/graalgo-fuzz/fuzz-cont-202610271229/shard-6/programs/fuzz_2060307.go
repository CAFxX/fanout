package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 int; F1 uint; F2 uint16 }
func genFn0(p0 uint) int {
    var v7 int16 = -32768
    var v8 uint32 = 802
    var v9 int8 = 39
    var v10 uint32 = 65535
    v8 += (v10 - (v8 >> (((p0 % ((p0) | 1))) & 63)))
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return -735
}

func main() {
    v1 := []string{(("a" + "\n\t") + (("foo bar" + "a") + "hello")), ("hello" + "foo bar"), ""}
    var v2 bool = false
    var v3 float32 = 1e10
    v4 := make(map[uint32]string)
    var v5 int32 = 1
    var v6 int32 = v5
    switch "a" {
    case "a":
        for fi11 := 0; fi11 < 13; fi11++ {
            v12 := GenS1{F0: 32767, F1: 42}
            v13 := GenS2{F0: 2, F1: (v6 * v5)}
            v14 := GenS2{F0: 1, F1: v6}
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
    default:
        var v15 int8 = 42
        _ = v15
    }
    switch "hello" {
    case "hello":
        var v16 uint64 = 4294967295
        _ = v16
    case "\n\t":
        v5 += 579
    default:
        if v2 {
            for fi17 := 0; fi17 < 7; fi17++ {
                for fi18 := 0; fi18 < 11; fi18++ {
                    if (2 == 122) {
                        switch genFn0(2) {
                        case 1:
                            fmt.Println(v6, v3)
                        case -703:
                            switch (fi18 + (fi17 * fi17)) {
                            case -9223372036854775808:
                                fmt.Println(v5, fi17)
                            case 1:
                                switch "" {
                                case "foo bar":
                                    v3 -= -2.25
                                    if (((-2.25 + 0.1) / 1e10) < float64(float32(1.5))) {
                                        var v19 uint32 = 708
                                        _ = v19
                                    } else {
                                        v20 := GenS2{F0: 255, F1: (v6 << ((638) & 63))}
                                        delete(v4, 4294967295)
                                        v6 += 1
                                        switch (fi17 * genFn0(225)) {
                                        case 2147483647:
                                            fmt.Println(v6, v5)
                                        case -777:
                                            v3 -= float32((-2.25 - (1e10 / 1.0)))
                                        case 9223372036854775807:
                                            v2 = ((v6 % ((v6) | 1)) == v5)
                                        }
                                        if v2 {
                                            var v21 string = "\n\t"
                                            _ = v21
                                        } else {
                                            if (fi17 <= fi17) {
                                                var v22 string = "a"
                                                v23 := v4[2147483648]
                                                if (true && (1 != 21)) {
                                                    fmt.Println(v23, v2)
                                                } else {
                                                    fmt.Println(v22, v6, v3)
                                                    var v24 uint32 = 2147483647
                                                    _ = v24
                                                }
                                                if (1 == 2147483647) {
                                                    v23 = "foo bar"
                                                    var v25 string = ""
                                                    var v26 uint8 = 128
                                                    _ = v25
                                                    _ = v26
                                                }
                                                _ = v22
                                                _ = v23
                                            } else {
                                                switch genFn0(255) {
                                                case 965:
                                                    var v27 uint16 = 65535
                                                    _ = v27
                                                default:
                                                    v28 := GenS1{F0: (fi17 ^ (^(fi17))), F1: 1}
                                                    _ = v28
                                                }
                                            }
                                        }
                                        fmt.Println(v2, v2)
                                        v5 = -1
                                        _ = v20
                                    }
                                    _ = len(v1)
                                }
                            case 974:
                                var v29 int32 = (v6 * v6)
                                _ = v29
                            }
                        default:
                            if (("hello" > "\n\t") && v2) {
                                fmt.Println(v2, v5)
                                var v30 uint16 = 847
                                v6 += v6
                                _ = v30
                            }
                        }
                        var v31 int16 = 477
                        _ = v31
                    }
                    _ = fi18
                }
                _ = fi17
            }
            for fi32 := 0; fi32 < 9; fi32++ {
                if (!(!v2)) {
                    if len(v1) > 0 {
                        v1[0] = "\n\t"
                    }
                    var v33 uint = 1
                    v34 := GenS3{F0: 1}
                    _ = v33
                    _ = v34
                } else {
                    fmt.Println(v5, v3)
                    v35 := v4[457]
                    fmt.Println(v2, v2)
                    _ = v35
                }
                _ = fi32
            }
        }
    }
    fmt.Println(v5, v5, v5)
    fmt.Println(v2, v3, v6)
    var v36 uint = 1
    for fi37 := 0; fi37 < 5; fi37++ {
        var v38 int = (fi37 | (fi37 * (fi37 & -529)))
        var v39 uint32 = 1
        _ = fi37
        _ = v38
        _ = v39
    }
    v2 = v2
    v6 = 496
    delete(v4, 4294967295)
    var v40 uint32 = 2147483648
    for fi41 := 0; fi41 < 18; fi41++ {
        for fi42 := 0; fi42 < 4; fi42++ {
            if (128 >= 199) {
                var v43 string = "hello"
                if v2 {
                    if (!(1 >= 42)) {
                        var v44 uint8 = 128
                        fmt.Println(v44, v43, fi42)
                        _ = v44
                    }
                } else {
                    v2 = ((!(v6 <= v5)) && (!v2))
                    v45 := GenS3{F0: 32767}
                    fmt.Println(v40, v6)
                    _ = v45
                }
                _ = v43
            }
            var v46 int8 = 0
            v47 := GenS4{F0: 255, F1: 0}
            _ = fi42
            _ = v46
            _ = v47
        }
        _ = fi41
    }
    switch "" {
    case "":
        v40 += v40
    case "foo bar":
        v48 := GenS5{F0: 32768, F1: v36, F2: 255}
        _ = v48
    case "\n\t":
        fmt.Println(v5)
    default:
        for fi49 := 0; fi49 < 12; fi49++ {
            for fi50 := 0; fi50 < 9; fi50++ {
                if ((^(v40)) > (v40 ^ (v40 ^ v40))) {
                    v3 += v3
                }
                _ = fi50
            }
            var v51 uint16 = 1
            _ = fi49
            _ = v51
        }
    }
    v3 = (v3 / (3.14159 * v3))
    if (48 > -671) {
        var v52 uint8 = 7
        _ = len(v1)
        fmt.Println(v40, v52)
        _ = v52
    } else {
        var v53 uint16 = 65535
        _ = v53
    }
    v40 -= v40
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}

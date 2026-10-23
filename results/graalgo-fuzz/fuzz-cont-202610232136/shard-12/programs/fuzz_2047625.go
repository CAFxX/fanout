package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 int16; F2 uint8 }
type GenS4 struct { F0 int16; F1 uint32; F2 string }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 int32; F1 int16; F2 int8 }
type GenS7 struct { F0 uint64; F1 int16; F2 string }
func genFn0(p0 int64, p1 int32) int16 {
    v6 := GenS1{F0: 1, F1: 178}
    v7 := GenS1{F0: -105, F1: 42}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 456
}
func genFn1() int64 {
    v8 := []int{127, -921}
    _ = len(v8)
    var v9 int16 = genFn0(42, 290)
    _ = v8
    _ = v9
    return 610
}

func main() {
    v1 := []uint32{346, 2147483648, 2}
    var v2 uint = 294
    var v3 uint64 = 1
    var v4 int32 = -1
    var v5 int16 = 42
    for fi10 := 0; fi10 < 19; fi10++ {
        switch "hello" {
        case "foo bar":
            if false {
                for fi11 := 0; fi11 < 9; fi11++ {
                    v12 := GenS1{F0: -1, F1: 1}
                    v5 = v5
                    fmt.Println(v2, fi11)
                    _ = fi11
                    _ = v12
                }
                fmt.Println(fi10, fi10)
            } else {
                v4 -= (v4 ^ (v4 % ((v4) | 1)))
                var v13 uint16 = 975
                _ = v13
            }
        case "":
            for fi14 := 0; fi14 < 1; fi14++ {
                if (true || (32768 != -930)) {
                    fmt.Println(fi14)
                    fmt.Println(v3)
                }
                _ = fi14
            }
        default:
            v15 := GenS2{F0: 706}
            _ = v15
        }
        _ = fi10
    }
    fmt.Println(v3)
    var v16 int8 = int8(v4)
    var v17 uint16 = 0
    v17 -= v17
    for fi18 := 0; fi18 < 6; fi18++ {
        v19 := GenS3{F0: "hello", F1: v5, F2: 1}
        v20 := GenS4{F0: genFn0(42, (^(v4))), F1: 0, F2: "hello"}
        _ = fi18
        _ = v19
        _ = v20
    }
    v17 -= 1
    var v21 int64 = int64(v3)
    switch 0 {
    case 9223372036854775807:
        v4 = v4
    case 1:
        var v22 uint8 = 103
        _ = v22
    default:
        fmt.Println(v16)
    }
    var v23 int64 = v21
    for fi24 := 0; fi24 < 11; fi24++ {
        var v25 float32 = 0.1
        _ = fi24
        _ = v25
    }
    v26 := GenS4{F0: v5, F1: 1, F2: "\n\t"}
    switch ("foo bar" + "\n\t") {
    case "hello":
        v27 := GenS1{F0: (-(v5)), F1: 59}
        _ = v27
    default:
        fmt.Println(v23, v3)
    }
    _ = len(v1)
    switch ("a" + "\n\t") {
    case "a":
        switch 32767 {
        case 128:
            v4 += v4
        default:
            v28 := GenS2{F0: 980}
            _ = v28
        }
    case "":
        switch 0 {
        case 117:
            v2 -= 4294967295
        case 287:
            _ = len(v1)
        }
    case "foo bar":
        var v29 uint64 = 2
        _ = v29
    default:
        switch -72 {
        case 2147483647:
            v30 := GenS2{F0: v2}
            _ = v30
        default:
            for fi31 := 0; fi31 < 13; fi31++ {
                for fi32 := 0; fi32 < 19; fi32++ {
                    switch "foo bar" {
                    case "foo bar":
                        var v33 float64 = float64(1e10)
                        fmt.Println(v5, v2)
                        var v34 int64 = (v23 & (-(v21)))
                        _ = v33
                        _ = v34
                    case "\n\t":
                        v35 := GenS5{F0: (v3 * v3), F1: "hello"}
                        _ = v35
                    case "hello":
                        if false {
                            if false {
                                var v36 float32 = (0.1 + (1e10 + (float32((0.1 * 0.1)) + 1e10)))
                                v37 := GenS5{F0: 0, F1: "a"}
                                _ = v36
                                _ = v37
                            } else {
                                v38 := GenS6{F0: 42, F1: v5, F2: 100}
                                v1 = append(v1, 4294967295)
                                _ = v38
                            }
                            if (127 > 128) {
                                v5 = genFn0(genFn1(), int32(v17))
                            }
                        }
                    }
                    v17 = v17
                    if false {
                        switch (fi31 / ((fi31) | 1)) {
                        case 32767:
                            v39 := GenS7{F0: (v3 % (((v3 >> ((255) & 63))) | 1)), F1: (v5 | (v5 & int16(v2))), F2: ""}
                            _ = v39
                        case 32768:
                            v4 = 2147483647
                        default:
                            v5 -= (v5 << ((v2) & 63))
                        }
                        v2 = 4294967295
                        var v40 int32 = v4
                        _ = v40
                    } else {
                        fmt.Println(v5, v17)
                        var v41 string = ("" + "a")
                        _ = v41
                    }
                    _ = fi32
                }
                v1 = append(v1, 4294967295)
                for fi42 := 0; fi42 < 6; fi42++ {
                    if true {
                        v1 = append(v1, 4294967295)
                    } else {
                        if true {
                            if false {
                                v5 = (v5 + 196)
                                var v43 uint8 = 42
                                v44 := GenS5{F0: v3, F1: ("hello" + ((("" + "\n\t") + "a") + ("" + "foo bar")))}
                                _ = v43
                                _ = v44
                            }
                            var v45 uint32 = 166
                            if (!true) {
                                switch (fi31 & -717) {
                                case 1:
                                    if len(v1) > 0 {
                                        v1[0] = (^(v45))
                                    }
                                case 127:
                                    v4 += (v4 << ((v2) & 63))
                                case 32767:
                                    var v46 uint = (v2 | (v2 % ((231) | 1)))
                                    _ = v46
                                }
                                _ = len(v1)
                            }
                            _ = v45
                        }
                    }
                    if true {
                        var v47 int32 = (v4 / ((v4) | 1))
                        _ = v47
                    }
                    _ = fi42
                }
                _ = fi31
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16; F1 bool; F2 uint8 }
type GenS5 struct { F0 bool; F1 bool; F2 uint16 }

func main() {
    var v1 bool = ((-1 <= -32768) && false)
    var v2 uint64 = 1
    var v3 uint32 = 120
    var v4 uint16 = 0
    v5 := []int{-335, 42, int(v3)}
    var v6 int8 = 24
    fmt.Println(v6, v1, v1)
    var v7 uint32 = 972
    if false {
        v3 = v3
        var v8 string = "hello"
        _ = v8
    } else {
        if len(v5) > 0 {
            v5[0] = 127
        }
    }
    _ = len(v5)
    var v9 uint32 = (v7 >> ((255) & 63))
    if (v9 <= v9) {
        for fi10 := 0; fi10 < 20; fi10++ {
            v9 -= (v9 | 0)
            _ = fi10
        }
        if (18446744073709551615 < (v2 << ((912) & 63))) {
            v11 := GenS1{F0: 0, F1: (v4 - v4), F2: (v4 << ((1) & 63))}
            v9 -= 2
            v12 := GenS1{F0: -1, F1: v4, F2: (v4 / ((v4) | 1))}
            _ = v11
            _ = v12
        }
        switch 1 {
        case 42:
            var v13 int8 = v6
            _ = v13
        case 32767:
            if true {
                if (true || (1.5 > 1e10)) {
                    var v14 uint16 = 0
                    v4 -= v4
                    if len(v5) > 0 {
                        v5[0] = 754
                    }
                    _ = v14
                }
                if v1 {
                    _ = len(v5)
                    for fi15 := 0; fi15 < 9; fi15++ {
                        v16 := GenS1{F0: -2147483648, F1: (^(v4)), F2: v4}
                        _ = fi15
                        _ = v16
                    }
                }
            } else {
                if (!true) {
                    var v17 int64 = 406
                    for fi18 := 0; fi18 < 19; fi18++ {
                        for fi19 := 0; fi19 < 8; fi19++ {
                            fmt.Println(v3, v4, v7)
                            switch "a" {
                            case "\n\t":
                                if true {
                                    v20 := GenS2{F0: -2147483648}
                                    if true {
                                        v21 := GenS1{F0: 2147483647, F1: v4, F2: v4}
                                        if v1 {
                                            var v22 bool = false
                                            _ = v22
                                        } else {
                                            if (!true) {
                                                switch (fi19 / (((fi18 & (fi18 << ((87) & 63)))) | 1)) {
                                                case 128:
                                                    v23 := GenS3{F0: v6}
                                                    _ = len(v5)
                                                    v1 = (false && ((uint16(v7) >= (^(v4))) || ((fi18 % (((fi19 + 597)) | 1)) == fi19)))
                                                    _ = v23
                                                case -9223372036854775808:
                                                    v5 = append(v5, fi18)
                                                case 964:
                                                    if v1 {
                                                        v17 = (v17 * -100)
                                                        var v24 int32 = -1
                                                        v17 = (^(v17))
                                                        _ = v24
                                                    }
                                                    var v25 float32 = 0.0
                                                    v26 := GenS4{F0: (^(v4)), F1: v1, F2: 1}
                                                    _ = v25
                                                    _ = v26
                                                }
                                                switch "foo bar" {
                                                case "\n\t":
                                                    if (-9223372036854775808 <= v17) {
                                                        var v27 uint8 = 1
                                                        fmt.Println(v2, v7, fi18)
                                                        _ = v27
                                                    } else {
                                                        v28 := GenS1{F0: -2147483648, F1: v4, F2: 2}
                                                        var v29 uint16 = (v4 - (v4 + (v4 ^ 255)))
                                                        v30 := GenS3{F0: (v6 % ((v6) | 1))}
                                                        _ = v28
                                                        _ = v29
                                                        _ = v30
                                                    }
                                                default:
                                                    v3 += v3
                                                }
                                                v6 += (^(v6))
                                            }
                                            if true {
                                                v1 = v1
                                                fmt.Println(fi18, v9, v7)
                                                switch "\n\t" {
                                                case "hello":
                                                    var v31 int32 = -2147483648
                                                    _ = v31
                                                case "foo bar":
                                                    switch (^(fi18)) {
                                                    case -152:
                                                        fmt.Println(fi19, v7, v17)
                                                    case -1:
                                                        v1 = v1
                                                    case 2147483647:
                                                        v4 -= (v4 << ((636) & 63))
                                                    }
                                                }
                                                v32 := GenS3{F0: (-(v6))}
                                                _ = v32
                                            }
                                            fmt.Println(fi18)
                                        }
                                        if v1 {
                                            var v33 uint16 = (v4 + v4)
                                            fmt.Println(v3, fi19, v2)
                                            _ = v33
                                        } else {
                                            var v34 float32 = (0.1 * 0.1)
                                            _ = v34
                                        }
                                        var v35 int64 = 659
                                        if (v1 || v1) {
                                            var v36 float64 = 0.1
                                            switch ((("foo bar" + ("a" + ("foo bar" + "hello"))) + "") + "hello") {
                                            case "a":
                                                v5 = append(v5, 0)
                                            default:
                                                v37 := GenS5{F0: v1, F1: v1, F2: (v4 ^ uint16(v2))}
                                                _ = v37
                                            }
                                            _ = v36
                                        } else {
                                            if (v17 != v17) {
                                                if (uint(v7) > 66) {
                                                    var v38 float32 = float32(((0.1 - (float64(1.5) / 1.0)) * 3.14159))
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        v39 := GenS3{F0: int8(v7)}
                                                        _ = v39
                                                    case "":
                                                        v4 += 255
                                                        v9 -= 42
                                                    case "a":
                                                        v40 := GenS2{F0: -312}
                                                        var v41 int64 = (v17 >> ((255) & 63))
                                                        _ = v40
                                                        _ = v41
                                                    default:
                                                        if (!v1) {
                                                            if (3.14159 >= (1e10 - 0.0)) {
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                                v5 = append(v5, fi19)
                                                            } else {
                                                                var v42 uint16 = v4
                                                                _ = v42
                                                            }
                                                            v43 := GenS4{F0: 32767, F1: (v35 <= (v17 + v35)), F2: 1}
                                                            _ = v43
                                                        } else {
                                                            _ = len(v5)
                                                        }
                                                    }
                                                    var v44 int8 = v6
                                                    _ = v38
                                                    _ = v44
                                                }
                                                var v45 uint64 = (v2 % ((v2) | 1))
                                                _ = len(v5)
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                _ = v45
                                            }
                                        }
                                        _ = v21
                                        _ = v35
                                    } else {
                                        var v46 bool = (!false)
                                        _ = v46
                                    }
                                    if v1 {
                                        if len(v5) > 0 {
                                            v5[0] = (fi19 >> ((577) & 63))
                                        }
                                        var v47 uint8 = 1
                                        v2 = (^(v2))
                                        _ = v47
                                    }
                                    var v48 float32 = 1e10
                                    _ = v20
                                    _ = v48
                                } else {
                                    var v49 bool = true
                                    _ = v49
                                }
                            case "a":
                                fmt.Println(v3, v2, v9)
                            }
                            var v50 string = (("hello" + "\n\t") + "hello")
                            _ = fi19
                            _ = v50
                        }
                        v1 = v1
                        v1 = v1
                        _ = fi18
                    }
                    _ = v17
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}

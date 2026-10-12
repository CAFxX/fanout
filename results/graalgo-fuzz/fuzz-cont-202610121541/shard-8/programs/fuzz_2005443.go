package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint32; F1 int8; F2 uint }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint, p1 int, p2 uint64) int16 {
    p1 -= (p1 ^ -9223372036854775808)
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := []uint32{4294967295, 2147483648, 245}
    v2 := make(map[uint32]uint32)
    v3 := []string{("a" + "foo bar"), ""}
    var v4 uint8 = 128
    var v5 uint64 = 0
    var v6 uint16 = 2
    var v7 int16 = 1000
    switch "hello" {
    case "\n\t":
        v8 := GenS1{F0: v5, F1: "hello"}
        _ = v8
    default:
        _ = len(v2)
    }
    var v9 string = "foo bar"
    if true {
        fmt.Println(v7, v7)
        switch 129 {
        case -265:
            var v10 uint = 65535
            _ = v10
        case 0:
            for fi11 := 0; fi11 < 13; fi11++ {
                for fi12 := 0; fi12 < 4; fi12++ {
                    v3 = append(v3, v9)
                    _ = fi12
                }
                _ = fi11
            }
        case -430:
            switch v9 {
            case "hello":
                for fi13 := 0; fi13 < 5; fi13++ {
                    var v14 int = 32768
                    if (1 > 65535) {
                        _ = len(v2)
                        v7 -= v7
                        for fi15 := 0; fi15 < 15; fi15++ {
                            if (v4 < (^(v4))) {
                                if true {
                                    if ((!false) && (!(v4 > (v4 * (^(v4)))))) {
                                        var v16 string = v9
                                        _ = v16
                                    } else {
                                        delete(v2, 610)
                                    }
                                    _ = len(v2)
                                } else {
                                    var v17 int = (^(fi15))
                                    switch ("foo bar" + v9) {
                                    case "foo bar":
                                        if len(v3) > 0 {
                                            v3[0] = "\n\t"
                                        }
                                    case "":
                                        var v18 int32 = -2147483648
                                        _ = v18
                                    default:
                                        v19 := v2[1]
                                        _ = v19
                                    }
                                    _ = v17
                                }
                            } else {
                                var v20 int8 = 1
                                if (false || true) {
                                    if true {
                                        if ((v5 | (v5 << ((4294967295) & 63))) != v5) {
                                            fmt.Println(v5)
                                        }
                                    } else {
                                        switch ("\n\t" + "") {
                                        case "\n\t":
                                            v21 := GenS1{F0: (v5 >> ((42) & 63)), F1: (v9 + v9)}
                                            _ = v21
                                        }
                                        v22 := GenS1{F0: 298, F1: v9}
                                        var v23 float64 = 3.14159
                                        _ = v22
                                        _ = v23
                                    }
                                }
                                var v24 uint32 = 0
                                _ = v20
                                _ = v24
                            }
                            _ = fi15
                        }
                    } else {
                        for fi25 := 0; fi25 < 18; fi25++ {
                            if true {
                                v4 = (v4 + v4)
                                var v26 float32 = 3.14159
                                if ((!true) && false) {
                                    var v27 uint16 = v6
                                    v6 = (v6 ^ v6)
                                    _ = v27
                                }
                                _ = v26
                            } else {
                                v9 = v9
                            }
                            fmt.Println(v4, fi25)
                            _ = fi25
                        }
                        for fi28 := 0; fi28 < 10; fi28++ {
                            v29 := GenS2{F0: 65535, F1: 0, F2: 0}
                            v4 -= v4
                            var v30 bool = true
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                    }
                    switch fi13 {
                    case 127:
                        _ = len(v3)
                    case 0:
                        for fi31 := 0; fi31 < 8; fi31++ {
                            fmt.Println(v4, v6)
                            switch v14 {
                            case -9223372036854775808:
                                v9 = (v9 + "a")
                            case -1:
                                v32 := GenS3{F0: 65535}
                                _ = v32
                            }
                            fmt.Println(v7, v4, v9)
                            _ = fi31
                        }
                    default:
                        for fi33 := 0; fi33 < 10; fi33++ {
                            if false {
                                v34 := GenS1{F0: 2, F1: "\n\t"}
                                _ = v34
                            } else {
                                v5 = 2
                            }
                            if false {
                                v6 = (v6 - (^(v6)))
                                var v35 uint8 = (v4 % (((v4 & v4)) | 1))
                                v36 := GenS1{F0: v5, F1: v9}
                                _ = v35
                                _ = v36
                            } else {
                                v2[uint32(v4)] = 978
                                v37 := GenS3{F0: 0}
                                var v38 int64 = 9223372036854775807
                                _ = v37
                                _ = v38
                            }
                            v4 += (^(v4))
                            _ = fi33
                        }
                    }
                    _ = fi13
                    _ = v14
                }
            case "\n\t":
                var v39 int64 = 241
                _ = v39
            }
        default:
            _ = len(v1)
        }
        v6 = v6
    }
    for fi40 := 0; fi40 < 7; fi40++ {
        if true {
            v5 -= v5
        }
        _ = fi40
    }
    var v41 int16 = v7
    v5 -= v5
    fmt.Println(v6)
    for fi42 := 0; fi42 < 12; fi42++ {
        var v43 string = "foo bar"
        if true {
            v7 += (v7 ^ genFn0(18446744073709551615, -9223372036854775808, v5))
            switch (-(fi42)) {
            case 251:
                v4 = v4
            case 32768:
                v44 := GenS1{F0: v5, F1: v9}
                _ = v44
            default:
                if ((-2.25 * (1.5 + 3.14159)) >= float64((0.0 + 1.5))) {
                    var v45 uint64 = (^(v5))
                    for fi46 := 0; fi46 < 14; fi46++ {
                        delete(v2, 0)
                        _ = fi46
                    }
                    _ = v45
                }
            }
            fmt.Println(v4)
        } else {
            var v47 uint32 = 42
            v41 -= genFn0(270, fi42, v5)
            var v48 uint32 = v47
            _ = v47
            _ = v48
        }
        _ = fi42
        _ = v43
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v49 := GenS3{F0: 2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v49
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}

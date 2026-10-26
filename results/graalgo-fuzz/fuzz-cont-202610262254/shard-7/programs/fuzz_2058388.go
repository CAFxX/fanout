package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 int32; F1 uint16; F2 int8 }
func genFn0(p0 uint64, p1 int, p2 int8) bool {
    var v7 int64 = -1
    var v8 int32 = -1
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := []int{}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (2 < 65535)
}

func main() {
    var v1 int32 = 1
    var v2 uint64 = 412
    var v3 uint = 4294967295
    v4 := make(map[int]bool)
    var v5 int64 = -269
    v6 := make(map[int]bool)
    var v11 uint32 = 0
    fmt.Println(v1, v2)
    if false {
        v12 := []int{9223372036854775807, 909, 128, 32767}
        switch "" {
        case "\n\t":
            var v13 int32 = 42
            _ = v13
        }
        if ((v2 << ((751) & 63)) > 0) {
            v14 := GenS1{F0: (^(v11))}
            if false {
                var v15 uint32 = v11
                fmt.Println(v3, v15, v3)
                _ = v15
            }
            _ = v14
        }
        _ = v12
    }
    if (18446744073709551615 >= (v3 | v3)) {
        v16 := GenS1{F0: v11}
        v17 := GenS1{F0: v11}
        fmt.Println(v2)
        _ = v16
        _ = v17
    }
    v6[511] = false
    for fi18 := 0; fi18 < 14; fi18++ {
        v3 = (^(v3))
        fmt.Println(v2, v5)
        _ = fi18
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        v4[fi19] = ((v5 - -252) > v5)
        switch (fi19 >> ((v3) & 63)) {
        case 9223372036854775807:
            if (0 != -1) {
                switch "hello" {
                case "foo bar":
                    for fi20 := 0; fi20 < 5; fi20++ {
                        var v21 float64 = 0.1
                        if false {
                            if ((127 >= 47) && (!(!false))) {
                                if (0.1 > 3.14159) {
                                    var v22 uint32 = 2147483648
                                    v5 -= (v5 >> ((v3) & 63))
                                    if (0 != 1000) {
                                        v23 := v6[fi19]
                                        _ = v23
                                    }
                                    _ = v22
                                }
                            } else {
                                var v24 float64 = (v21 - v21)
                                var v25 uint32 = (v11 % ((v11) | 1))
                                if false {
                                    var v26 int8 = -18
                                    _ = v26
                                }
                                if false {
                                    var v27 float64 = (((v21 * float64((-2.25 / -2.25))) * ((v24 * -2.25) - (1.5 + 0.0))) / v21)
                                    _ = v27
                                }
                                var v28 string = ""
                                _ = v24
                                _ = v25
                                _ = v28
                            }
                        } else {
                            switch fi19 {
                            case -2147483648:
                                v2 -= uint64(v3)
                            case 4294967295:
                                v1 = 620
                                v2 -= 1
                            default:
                                fmt.Println(v1, v2, v5)
                            }
                            v29 := GenS1{F0: 42}
                            _ = v29
                        }
                        switch (fi20 >> (((v3 | 652)) & 63)) {
                        case 127:
                            if false {
                                v30 := v6[0]
                                v5 = 9223372036854775807
                                _ = v30
                            } else {
                                var v31 int = 127
                                var v32 uint32 = v11
                                switch 128 {
                                case 632:
                                    switch (fi20 + (fi20 >> ((255) & 63))) {
                                    case -1:
                                        _ = len(v6)
                                    default:
                                        fmt.Println(v32, fi20)
                                    }
                                case -604:
                                    fmt.Println(fi19, v2)
                                case 717:
                                    v3 = 389
                                }
                                _ = v31
                                _ = v32
                            }
                            v33 := GenS1{F0: (^(v11))}
                            _ = v33
                        case -362:
                            if false {
                                var v34 uint64 = (v2 & v2)
                                _ = v34
                            }
                        default:
                            fmt.Println(fi19)
                        }
                        _ = fi20
                        _ = v21
                    }
                case "\n\t":
                    fmt.Println(fi19)
                }
                if (((1.5 * (1e10 - float64((((1.5 / 1e10) - 3.14159) + (-2.25 / 0.1))))) / (0.1 - 1.5)) <= 0.1) {
                    v35 := []int{int(v3), 644}
                    var v36 uint16 = 864
                    for fi37 := 0; fi37 < 6; fi37++ {
                        var v38 uint8 = 1
                        v39 := GenS2{F0: (v36 >> (((v3 & (^(v3)))) & 63))}
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                    _ = v35
                    _ = v36
                } else {
                    v40 := GenS3{F0: -667, F1: (^(fi19))}
                    var v41 float64 = float64(0.0)
                    _ = v40
                    _ = v41
                }
                delete(v6, fi19)
            }
        case 4294967295:
            v1 = v1
        case 42:
            for fi42 := 0; fi42 < 16; fi42++ {
                if false {
                    var v43 int32 = (v1 - v1)
                    if true {
                        switch ("hello" + "foo bar") {
                        case "hello":
                            v3 = (v3 + (v3 & v3))
                        case "foo bar":
                            v44 := GenS3{F0: -1, F1: 42}
                            _ = v44
                        }
                    }
                    _ = v43
                } else {
                    v11 -= uint32(v2)
                    v2 = (v2 | v2)
                }
                _ = fi42
            }
        }
        _ = fi19
    }
    if (!true) {
        if (2147483647 == int(v3)) {
            var v45 int32 = v1
            _ = v45
        }
    } else {
        var v46 uint8 = 183
        v47 := GenS4{F0: (v1 / ((int32(v11)) | 1)), F1: 928, F2: 100}
        fmt.Println(v1, v5)
        _ = v46
        _ = v47
    }
    var v48 string = ("" + ("a" + "\n\t"))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
}

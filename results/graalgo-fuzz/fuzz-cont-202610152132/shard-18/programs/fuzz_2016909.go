package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 int8; F2 string }
type GenS5 struct { F0 int8; F1 uint32; F2 int }

func main() {
    var v1 uint64 = 1
    var v2 int16 = -32768
    v3 := []uint32{}
    v4 := []string{"a", ("hello" + ("\n\t" + "a"))}
    var v5 uint64 = 433
    for fi6 := 0; fi6 < 20; fi6++ {
        for fi7 := 0; fi7 < 12; fi7++ {
            fmt.Println(v1, v2, v1)
            v8 := GenS1{F0: (!true)}
            var v9 uint16 = 2
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    switch -376 {
    case 127:
        var v10 int32 = 2147483647
        _ = v10
    default:
        if true {
            v1 = (v1 << ((1) & 63))
        }
    }
    if false {
        if true {
            for fi11 := 0; fi11 < 18; fi11++ {
                if len(v4) > 0 {
                    v4[0] = ""
                }
                _ = fi11
            }
            switch "" {
            case "a":
                v12 := GenS1{F0: ((v2 / ((v2) | 1)) == v2)}
                _ = v12
            }
            v2 = (v2 + (v2 | (v2 - v2)))
        } else {
            var v13 int8 = 42
            fmt.Println(v13, v1)
            v14 := GenS1{F0: true}
            _ = v13
            _ = v14
        }
    }
    fmt.Println(v2, v5)
    if true {
        v1 -= v5
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 uint8 = 2
            var v17 bool = false
            _ = fi15
            _ = v16
            _ = v17
        }
    } else {
        fmt.Println(v5, v5)
    }
    v2 = v2
    fmt.Println(v1, v2, v1)
    v18 := GenS1{F0: (678 != 2147483648)}
    var v19 string = "a"
    if true {
        v20 := GenS2{F0: (v2 << ((568) & 63)), F1: -9223372036854775808}
        if true {
            switch -256 {
            case 32767:
                var v21 float64 = ((float64((0.0 / 1.0)) * 1.5) * 1e10)
                _ = v21
            case 2147483647:
                for fi22 := 0; fi22 < 17; fi22++ {
                    var v23 int8 = -128
                    v24 := GenS2{F0: (v2 + v2), F1: 1}
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            case -71:
                v25 := GenS3{F0: 4294967295}
                _ = v25
            default:
                v2 -= (v2 >> ((4294967295) & 63))
            }
        } else {
            if (false && (!true)) {
                v3 = append(v3, 42)
            }
        }
        if ((v2 % ((-1) | 1)) <= (v2 + v2)) {
            fmt.Println(v1)
            if (("hello" + ("\n\t" + (((v19 + v19) + v19) + v19))) >= ((("a" + v19) + v19) + v19)) {
                switch 42 {
                case -1:
                    for fi26 := 0; fi26 < 6; fi26++ {
                        if (2 <= 1) {
                            for fi27 := 0; fi27 < 7; fi27++ {
                                var v28 uint8 = 127
                                var v29 int32 = -1
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                            if (true && true) {
                                switch v19 {
                                case "foo bar":
                                    if (840 == 42) {
                                        var v30 uint = 18446744073709551615
                                        _ = v30
                                    } else {
                                        if len(v3) > 0 {
                                            v3[0] = 2
                                        }
                                        v1 += (v1 ^ v5)
                                    }
                                }
                            }
                        }
                        _ = fi26
                    }
                }
                v5 = (^(v1))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                var v31 int16 = v2
                _ = v31
            }
        } else {
            switch 1 {
            case -9223372036854775808:
                if true {
                    v32 := GenS4{F0: v19, F1: -128, F2: v19}
                    _ = v32
                } else {
                    _ = len(v3)
                    v33 := GenS5{F0: 127, F1: 846, F2: 242}
                    if true {
                        for fi34 := 0; fi34 < 8; fi34++ {
                            v35 := GenS5{F0: 127, F1: 2, F2: fi34}
                            _ = fi34
                            _ = v35
                        }
                    } else {
                        var v36 bool = true
                        _ = v36
                    }
                    _ = v33
                }
            case 127:
                for fi37 := 0; fi37 < 8; fi37++ {
                    v2 = (v2 % (((v2 >> ((285) & 63))) | 1))
                    for fi38 := 0; fi38 < 9; fi38++ {
                        if true {
                            fmt.Println(v19)
                        } else {
                            switch "a" {
                            case "foo bar":
                                var v39 uint16 = 65535
                                _ = v39
                            case "a":
                                v5 -= v1
                            case "\n\t":
                                v5 += v1
                            }
                            v2 += v2
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        var v40 float64 = (1e10 - 1.5)
                        var v41 int32 = 0
                        var v42 int8 = 26
                        var v43 uint16 = 768
                        var v44 uint64 = 18446744073709551615
                        _ = fi38
                        _ = v40
                        _ = v41
                        _ = v42
                        _ = v43
                        _ = v44
                    }
                    _ = fi37
                }
            case 243:
                v45 := GenS5{F0: 127, F1: 268, F2: 873}
                _ = v45
            }
        }
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

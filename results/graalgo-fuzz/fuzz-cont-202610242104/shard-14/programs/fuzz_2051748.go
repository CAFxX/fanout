package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 uint32; F1 uint }
type GenS7 struct { F0 int8; F1 int; F2 bool }

func main() {
    var v1 int16 = -545
    v2 := []string{"\n\t", "", "foo bar", ("foo bar" + ("hello" + "hello"))}
    var v3 int = -1
    v4 := make(map[uint32]int)
    var v5 int8 = 1
    if false {
        v3 = v3
        v5 += -1
        v3 = 1
    } else {
        var v6 int = (v3 | 0)
        fmt.Println(v6, v3, v5)
        _ = v6
    }
    if (0 < 42) {
        switch (v3 | 128) {
        case 1:
            if ((v5 << ((831) & 63)) == v5) {
                v7 := GenS1{F0: 0, F1: 0, F2: 127}
                var v8 float32 = 3.14159
                _ = v7
                _ = v8
            }
        case 32768:
            if (!(!false)) {
                v5 = -128
                fmt.Println(v3, v5)
                var v9 uint64 = 42
                _ = v9
            } else {
                for fi10 := 0; fi10 < 2; fi10++ {
                    for fi11 := 0; fi11 < 14; fi11++ {
                        v3 = fi11
                        _ = fi11
                    }
                    if false {
                        v12 := GenS1{F0: 18446744073709551615, F1: 215, F2: 0}
                        v13 := GenS2{F0: 698}
                        _ = v12
                        _ = v13
                    }
                    _ = fi10
                }
            }
        }
        switch ("foo bar" + (("hello" + ("" + "hello")) + ("\n\t" + ((("a" + "foo bar") + "") + "a")))) {
        case "\n\t":
            switch 886 {
            case -993:
                for fi14 := 0; fi14 < 20; fi14++ {
                    fmt.Println(v5, v5, v1)
                    v1 = -323
                    v1 = (-(v1))
                    _ = fi14
                }
            case 204:
                v3 += v3
            case 2147483647:
                var v15 float64 = ((3.14159 - (1.5 + 0.0)) * ((0.0 / (0.1 * 3.14159)) + -2.25))
                _ = v15
            }
        case "a":
            var v16 int16 = v1
            _ = v16
        }
    } else {
        var v17 uint32 = 4294967295
        var v18 bool = false
        _ = v17
        _ = v18
    }
    if false {
        if ((-30 < (^(v5))) || ((((1e10 * 0.1) == ((-2.25 / (0.0 + -2.25)) * (1e10 / -2.25))) && true) && (false || true))) {
            v1 -= (v1 ^ (v1 | v1))
            if false {
                switch (v3 ^ v3) {
                case -1:
                    var v19 int32 = -507
                    _ = v19
                default:
                    v20 := GenS3{F0: -32768, F1: 128}
                    _ = v20
                }
                var v21 int64 = 1
                delete(v4, 42)
                _ = v21
            } else {
                var v22 bool = ("a" != "a")
                var v23 int = v3
                _ = v22
                _ = v23
            }
        } else {
            if false {
                var v24 int64 = 150
                v25 := GenS3{F0: 42, F1: v3}
                v1 = (v1 << ((506) & 63))
                _ = v24
                _ = v25
            } else {
                var v26 int32 = 0
                _ = v26
            }
        }
    } else {
        for fi27 := 0; fi27 < 17; fi27++ {
            fmt.Println(v5, fi27, v1)
            _ = fi27
        }
        fmt.Println(v3)
        v28 := v4[0]
        _ = v28
    }
    var v29 uint = 42
    v30 := GenS3{F0: v1, F1: -862}
    var v31 string = "foo bar"
    switch (v3 * v3) {
    case -730:
        if false {
            if ((!(v1 != v1)) || (4294967295 == 0)) {
                switch v31 {
                case "foo bar":
                    for fi32 := 0; fi32 < 14; fi32++ {
                        v31 = "\n\t"
                        v33 := GenS4{F0: v31, F1: (^(v1))}
                        _ = fi32
                        _ = v33
                    }
                case "a":
                    var v34 string = (("hello" + ((v31 + "foo bar") + (v31 + v31))) + v31)
                    _ = v34
                case "":
                    if (true || (1 <= 2147483647)) {
                        v35 := GenS5{F0: v1, F1: 127}
                        v1 += -934
                        _ = v35
                    }
                }
                switch v3 {
                case 127:
                    for fi36 := 0; fi36 < 20; fi36++ {
                        switch fi36 {
                        case 1:
                            fmt.Println(v1, v3)
                        case 458:
                            if true {
                                fmt.Println(fi36)
                                if len(v2) > 0 {
                                    v2[0] = v31
                                }
                                v37 := GenS3{F0: (v1 - v1), F1: -602}
                                _ = v37
                            }
                        }
                        if len(v2) > 0 {
                            v2[0] = v31
                        }
                        _ = fi36
                    }
                case -835:
                    switch (v3 / ((v3) | 1)) {
                    case -2147483648:
                        fmt.Println(v31)
                    case -788:
                        if true {
                            if len(v2) > 0 {
                                v2[0] = (v31 + "foo bar")
                            }
                        } else {
                            if true {
                                v3 += (v3 & v3)
                                if len(v2) > 0 {
                                    v2[0] = v31
                                }
                            } else {
                                switch v31 {
                                case "":
                                    var v38 int8 = v5
                                    _ = v38
                                case "\n\t":
                                    v39 := GenS6{F0: 0, F1: (v29 - v29)}
                                    _ = v39
                                case "hello":
                                    v40 := GenS7{F0: (v5 & 42), F1: v3, F2: (0 > -1)}
                                    _ = v40
                                }
                                var v41 uint64 = 42
                                for fi42 := 0; fi42 < 4; fi42++ {
                                    var v43 uint = (v29 * v29)
                                    var v44 float32 = 1.5
                                    _ = fi42
                                    _ = v43
                                    _ = v44
                                }
                                _ = v41
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v45 := GenS2{F0: 296}
                            _ = v45
                        }
                    default:
                        v1 += (v1 + v1)
                    }
                }
                v2 = append(v2, "hello")
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if len(v2) > 0 {
                v2[0] = v31
            }
        }
    }
    var v46 int8 = v5
    for fi47 := 0; fi47 < 17; fi47++ {
        switch "hello" {
        case "":
            var v48 uint64 = 88
            _ = v48
        }
        if len(v2) > 0 {
            v2[0] = ("\n\t" + v31)
        }
        _ = fi47
    }
    if true {
        var v49 int64 = 9223372036854775807
        _ = v49
    } else {
        fmt.Println(v5, v29)
        v29 = (v29 & v29)
    }
    var v50 float64 = (1.5 * (1.5 * -2.25))
    v31 = "hello"
    for fi51 := 0; fi51 < 10; fi51++ {
        v52 := GenS4{F0: v31, F1: int16(v29)}
        fmt.Println(v3)
        _ = fi51
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v46
    _ = v5
    _ = v50
}

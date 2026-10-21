package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 uint64; F2 uint }
type GenS4 struct { F0 uint16; F1 int; F2 bool }
type GenS5 struct { F0 int; F1 uint64 }
type GenS6 struct { F0 int8; F1 uint8; F2 uint8 }

func main() {
    var v1 uint8 = 134
    var v2 uint32 = 42
    v3 := []string{""}
    var v4 int16 = -1
    v4 = v4
    var v5 string = (("a" + "\n\t") + "a")
    if false {
        var v6 bool = true
        if (0.1 <= 3.14159) {
            var v7 bool = (0.1 == ((1.5 * (0.1 / 1.0)) + 3.14159))
            for fi8 := 0; fi8 < 4; fi8++ {
                var v9 int64 = 697
                fmt.Println(v5, v2, v9)
                var v10 uint32 = (v2 << ((4294967295) & 63))
                _ = fi8
                _ = v10
                _ = v9
            }
            if (4294967295 > 2) {
                var v11 uint = 42
                if v6 {
                    if (65535 >= 711) {
                        var v12 bool = (!(!v7))
                        if v7 {
                            switch "\n\t" {
                            case "foo bar":
                                v13 := GenS1{F0: (v5 + ("\n\t" + ("" + ((v5 + v5) + (("\n\t" + (v5 + v5)) + v5)))))}
                                _ = v13
                            default:
                                var v14 uint = v11
                                _ = v14
                            }
                            var v15 uint32 = (v2 * (v2 + (^(v2))))
                            v16 := GenS2{F0: 42}
                            _ = v15
                            _ = v16
                        } else {
                            v3 = append(v3, v5)
                            if true {
                                var v17 float32 = (-2.25 * 3.14159)
                                if (1 > int(v11)) {
                                    switch 1 {
                                    case 0:
                                        for fi18 := 0; fi18 < 13; fi18++ {
                                            var v19 string = ("a" + "hello")
                                            _ = fi18
                                            _ = v19
                                        }
                                    case 2147483647:
                                        if (v6 || (42 > 0)) {
                                            v20 := GenS2{F0: -2147483648}
                                            if (v5 == "hello") {
                                                v17 = (((((((v17 + v17) * -2.25) / 1e10) / ((3.14159 - (float32(1.5) / v17)) * (((3.14159 * v17) - 1e10) - 1.5))) + float32(1.5)) + (1.5 + -2.25)) * v17)
                                            } else {
                                                fmt.Println(v1, v11)
                                                v6 = false
                                            }
                                            var v21 int64 = -1
                                            _ = v20
                                            _ = v21
                                        }
                                    case 4294967295:
                                        v1 += v1
                                    }
                                    switch -9223372036854775808 {
                                    case 4294967295:
                                        var v22 uint64 = 18446744073709551615
                                        _ = v22
                                    case 32767:
                                        for fi23 := 0; fi23 < 11; fi23++ {
                                            v24 := GenS2{F0: (fi23 % ((fi23) | 1))}
                                            _ = fi23
                                            _ = v24
                                        }
                                    }
                                } else {
                                    for fi25 := 0; fi25 < 10; fi25++ {
                                        v7 = false
                                        _ = fi25
                                    }
                                    switch -458 {
                                    case -802:
                                        v5 = v5
                                    case 9223372036854775807:
                                        var v26 bool = v12
                                        _ = v26
                                    case -515:
                                        fmt.Println(v12, v6, v17)
                                    }
                                }
                                var v27 bool = (382 <= 1)
                                _ = v17
                                _ = v27
                            }
                        }
                        _ = v12
                    } else {
                        var v28 uint8 = v1
                        _ = v28
                    }
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v1 = (^(v1))
                    v4 = v4
                }
                _ = v11
            } else {
                fmt.Println(v7, v5)
                fmt.Println(v1, v4)
                var v29 uint16 = 1
                _ = v29
            }
            _ = v7
        }
        _ = v6
    } else {
        fmt.Println(v2, v4, v5)
        v4 += 32767
    }
    v30 := GenS3{F0: 128, F1: 1, F2: 18446744073709551615}
    fmt.Println(v2, v4)
    if true {
        for fi31 := 0; fi31 < 2; fi31++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v32 uint8 = 128
            _ = fi31
            _ = v32
        }
        for fi33 := 0; fi33 < 2; fi33++ {
            fmt.Println(v5, v2)
            _ = fi33
        }
        v1 += v1
    } else {
        var v34 uint = 297
        for fi35 := 0; fi35 < 3; fi35++ {
            var v36 int = fi35
            var v37 uint32 = v2
            for fi38 := 0; fi38 < 19; fi38++ {
                switch v5 {
                case "\n\t":
                    switch (v36 % (((fi35 % ((fi38) | 1))) | 1)) {
                    case 128:
                        var v39 int64 = -9223372036854775808
                        _ = v39
                    case -119:
                        v5 = "hello"
                    case 0:
                        v4 = (v4 * (v4 & v4))
                    }
                case "foo bar":
                    _ = len(v3)
                }
                _ = fi38
            }
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = v34
    }
    for fi40 := 0; fi40 < 17; fi40++ {
        v41 := GenS4{F0: 2, F1: (fi40 | 2147483647), F2: (v5 <= ("foo bar" + v5))}
        if (((1.5 - 0.1) * float32(float64(1.5))) < 1.5) {
            switch v5 {
            case "foo bar":
                v4 = (v4 / (((v4 ^ v4)) | 1))
            case "a":
                v3 = append(v3, v5)
            default:
                v42 := GenS5{F0: fi40, F1: 42}
                _ = v42
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if ((^(v1)) > (^(v1))) {
                if len(v3) > 0 {
                    v3[0] = (((v5 + v5) + v5) + ((v5 + v5) + "foo bar"))
                }
                var v43 uint64 = 18446744073709551615
                var v44 uint64 = (^(v43))
                _ = v43
                _ = v44
            } else {
                v45 := GenS3{F0: fi40, F1: 2, F2: 2}
                for fi46 := 0; fi46 < 1; fi46++ {
                    fmt.Println(v4, v5, fi46)
                    fmt.Println(fi40)
                    _ = fi46
                }
                _ = v45
            }
        } else {
            if true {
                v5 = (v5 + v5)
                var v47 uint8 = 20
                if len(v3) > 0 {
                    v3[0] = v5
                }
                _ = v47
            }
        }
        var v48 float32 = ((-2.25 * 0.0) - 0.0)
        _ = fi40
        _ = v41
        _ = v48
    }
    if (((v2 / (((v2 - v2)) | 1)) <= v2) || (128 != 1)) {
        fmt.Println(v2)
    }
    for fi49 := 0; fi49 < 3; fi49++ {
        var v50 uint8 = (v1 - (v1 & v1))
        for fi51 := 0; fi51 < 15; fi51++ {
            fmt.Println(fi51, v50)
            if false {
                var v52 float32 = 1.5
                v53 := GenS6{F0: -1, F1: (v1 ^ v1), F2: (v1 * 127)}
                _ = v52
                _ = v53
            }
            v4 -= v4
            _ = fi51
        }
        _ = fi49
        _ = v50
    }
    for fi54 := 0; fi54 < 2; fi54++ {
        var v55 int = 401
        if false {
            var v56 uint16 = 586
            for fi57 := 0; fi57 < 20; fi57++ {
                var v58 int8 = 100
                v58 = v58
                _ = fi57
                _ = v58
            }
            _ = v56
        }
        v2 = 65535
        _ = fi54
        _ = v55
    }
    for fi59 := 0; fi59 < 13; fi59++ {
        fmt.Println(v5, v1, fi59)
        _ = fi59
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}

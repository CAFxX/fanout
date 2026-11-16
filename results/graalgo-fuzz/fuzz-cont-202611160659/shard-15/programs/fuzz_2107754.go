package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int8; F1 uint8; F2 uint16 }

func main() {
    var v1 int32 = -1
    var v2 string = "a"
    var v3 int16 = -32768
    v4 := make(map[uint32]bool)
    v5 := []uint32{1, 335}
    var v6 uint64 = 2
    var v7 int32 = (v1 >> ((2) & 63))
    v5 = append(v5, 2147483648)
    var v8 uint8 = 1
    var v9 int32 = (v7 % ((-1) | 1))
    switch "a" {
    case "foo bar":
        v10 := GenS1{F0: -1}
        _ = v10
    case "hello":
        v11 := GenS2{F0: 32768, F1: v3}
        _ = v11
    case "\n\t":
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 float32 = 1e10
            _ = fi12
            _ = v13
        }
    default:
        for fi14 := 0; fi14 < 3; fi14++ {
            for fi15 := 0; fi15 < 6; fi15++ {
                v9 = 42
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if (!((v7 != (v7 & (v9 >> ((4294967295) & 63)))) || (282 < 0))) {
                    var v16 string = "foo bar"
                    v1 -= v9
                    if true {
                        var v17 uint8 = uint8(v6)
                        v5 = append(v5, 4294967295)
                        _ = v17
                    }
                    _ = v16
                } else {
                    v2 = v2
                }
                _ = fi15
            }
            var v18 int64 = 9223372036854775807
            if false {
                var v19 float32 = (((3.14159 + 3.14159) / (-2.25 * (0.1 * float32(1e10)))) - 0.0)
                v1 = (v9 & (v7 - 2147483647))
                fmt.Println(v3, v3, v19)
                _ = v19
            } else {
                fmt.Println(fi14, v7)
            }
            _ = fi14
            _ = v18
        }
    }
    if (v2 != "\n\t") {
        v4[0] = true
        for fi20 := 0; fi20 < 3; fi20++ {
            if (true && (!true)) {
                fmt.Println(v1, v9, v6)
                v6 -= (v6 >> ((1) & 63))
            } else {
                if (true && (true && false)) {
                    if ((36 == -128) && false) {
                        var v21 string = (v2 + v2)
                        if true {
                            v6 -= v6
                        }
                        _ = v21
                    } else {
                        if false {
                            fmt.Println(v6, v9)
                            _ = len(v4)
                        } else {
                            for fi22 := 0; fi22 < 13; fi22++ {
                                var v23 float32 = 1e10
                                var v24 float32 = v23
                                fmt.Println(v23, v3)
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                        }
                    }
                    var v25 uint64 = v6
                    _ = v25
                } else {
                    fmt.Println(fi20)
                    if true {
                        switch int(v8) {
                        case 891:
                            if (true || true) {
                                var v26 float64 = -2.25
                                _ = v26
                            } else {
                                switch fi20 {
                                case 32767:
                                    var v27 uint16 = 534
                                    _ = v27
                                case -887:
                                    var v28 uint64 = 18446744073709551615
                                    _ = v28
                                default:
                                    _ = len(v4)
                                }
                            }
                        default:
                            switch (fi20 | fi20) {
                            case -181:
                                var v29 uint16 = 1
                                _ = v29
                            case 128:
                                var v30 int64 = -9223372036854775808
                                _ = v30
                            case 0:
                                var v31 int16 = 1000
                                _ = v31
                            }
                        }
                        var v32 int64 = -9223372036854775808
                        _ = v32
                    }
                    v7 = (v7 | 44)
                }
            }
            var v33 float32 = -2.25
            switch ("a" + (((("\n\t" + ("a" + (v2 + (v2 + "")))) + "foo bar") + v2) + v2)) {
            case "":
                var v34 uint64 = 2
                _ = v34
            case "a":
                if false {
                    if (!(127 < 100)) {
                        for fi35 := 0; fi35 < 17; fi35++ {
                            v6 += v6
                            _ = fi35
                        }
                    }
                } else {
                    var v36 uint = 18446744073709551615
                    if len(v5) > 0 {
                        v5[0] = 1
                    }
                    v37 := GenS1{F0: 100}
                    _ = v36
                    _ = v37
                }
            default:
                var v38 bool = ((true && false) || false)
                _ = v38
            }
            _ = fi20
            _ = v33
        }
        switch v2 {
        case "\n\t":
            var v39 int8 = 127
            _ = v39
        case "a":
            v40 := v4[2147483647]
            _ = v40
        case "":
            v41 := GenS1{F0: 0}
            _ = v41
        }
    }
    if (!(false && true)) {
        v42 := GenS3{F0: 0, F1: v8, F2: 911}
        _ = v42
    } else {
        if (!false) {
            switch "\n\t" {
            case "a":
                v6 -= uint64(v8)
            case "\n\t":
                for fi43 := 0; fi43 < 5; fi43++ {
                    fmt.Println(v7, v1)
                    var v44 int = -9223372036854775808
                    _ = fi43
                    _ = v44
                }
            case "":
                if len(v5) > 0 {
                    v5[0] = 994
                }
            default:
                fmt.Println(v9)
            }
        } else {
            var v45 int8 = 127
            var v46 uint = 0
            _ = v45
            _ = v46
        }
        v1 = (v7 ^ (v1 | v9))
    }
    if ((!false) || true) {
        v2 = v2
    } else {
        _ = len(v4)
        v47 := GenS2{F0: 127, F1: 1000}
        var v48 uint8 = v8
        _ = v47
        _ = v48
    }
    var v49 uint8 = 1
    for fi50 := 0; fi50 < 20; fi50++ {
        for fi51 := 0; fi51 < 8; fi51++ {
            if ((fi51 - fi50) == (fi51 ^ (fi51 ^ 32768))) {
                var v52 int16 = (v3 ^ v3)
                v7 = (v9 >> ((180) & 63))
                _ = v52
            }
            v4[118] = true
            fmt.Println(v8)
            v53 := v4[uint32(v6)]
            _ = fi51
            _ = v53
        }
        var v54 uint8 = (v49 - (v49 | 255))
        switch fi50 {
        case 127:
            v2 = (((v2 + v2) + v2) + v2)
        case 4294967295:
            switch "a" {
            case "\n\t":
                var v55 int = fi50
                _ = v55
            }
        }
        _ = fi50
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

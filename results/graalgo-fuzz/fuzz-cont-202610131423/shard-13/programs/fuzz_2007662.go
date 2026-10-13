package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 bool; F2 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 int8; F2 uint8 }

func main() {
    var v1 uint = 0
    var v2 uint16 = 360
    v3 := make(map[int]uint32)
    v4 := []uint32{1}
    var v5 int32 = 705
    switch ("a" + "a") {
    case "\n\t":
        var v6 float32 = 1.5
        _ = v6
    default:
        if (false && true) {
            _ = len(v3)
        }
    }
    var v7 uint32 = 333
    if false {
        var v8 bool = true
        switch "foo bar" {
        case "\n\t":
            var v9 uint32 = v7
            _ = v9
        case "foo bar":
            _ = len(v3)
        default:
            var v10 uint = uint(v7)
            _ = v10
        }
        _ = v8
    } else {
        v7 += (v7 >> ((uint(v2)) & 63))
    }
    var v11 uint = (v1 << ((65535) & 63))
    for fi12 := 0; fi12 < 5; fi12++ {
        v5 -= (v5 & (v5 ^ (-(v5))))
        var v13 uint32 = 2147483647
        fmt.Println(v13, v7)
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: -9223372036854775808}
    fmt.Println(v5, v7, v2)
    v15 := GenS1{F0: 753}
    for fi16 := 0; fi16 < 9; fi16++ {
        switch (fi16 & (^(fi16))) {
        case 128:
            switch (fi16 >> ((0) & 63)) {
            case -9223372036854775808:
                v17 := GenS1{F0: -256}
                _ = v17
            case 933:
                fmt.Println(fi16, v7, v1)
            case 9223372036854775807:
                var v18 int = fi16
                _ = v18
            }
        case 4294967295:
            var v19 uint64 = 110
            _ = v19
        }
        for fi20 := 0; fi20 < 17; fi20++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v21 int32 = v5
            _ = fi20
            _ = v21
        }
        _ = fi16
    }
    switch ("" + "hello") {
    case "hello":
        v7 -= (v7 ^ v7)
    case "a":
        var v22 int32 = (v5 % ((v5) | 1))
        _ = v22
    default:
        v23 := GenS1{F0: int64(v11)}
        _ = v23
    }
    if len(v4) > 0 {
        v4[0] = v7
    }
    if (!false) {
        var v24 bool = false
        _ = v24
    }
    fmt.Println(v1, v1, v1)
    switch -2147483648 {
    case -477:
        v1 = uint(v2)
    case 32768:
        var v25 int16 = 42
        _ = v25
    case 32767:
        for fi26 := 0; fi26 < 15; fi26++ {
            if len(v4) > 0 {
                v4[0] = v7
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi26
        }
    default:
        switch 1 {
        case -9223372036854775808:
            for fi27 := 0; fi27 < 5; fi27++ {
                switch ("hello" + "a") {
                case "\n\t":
                    if (!(622 > 4294967295)) {
                        fmt.Println(v11, v1)
                        for fi28 := 0; fi28 < 10; fi28++ {
                            fmt.Println(v7)
                            v2 = (^(v2))
                            _ = fi28
                        }
                        var v29 float32 = float32(0.0)
                        _ = v29
                    } else {
                        v2 = v2
                        var v30 string = "foo bar"
                        _ = v30
                    }
                case "":
                    for fi31 := 0; fi31 < 15; fi31++ {
                        switch "" {
                        case "foo bar":
                            if true {
                                switch 42 {
                                case -240:
                                    var v32 uint32 = (v7 / ((v7) | 1))
                                    var v33 string = "foo bar"
                                    var v34 int32 = v5
                                    v5 -= (v5 >> (((v1 >> ((v1) & 63))) & 63))
                                    fmt.Println(v34, v11, fi31)
                                    v11 = v1
                                    fmt.Println(fi31)
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                case 32767:
                                    var v35 int = (fi27 | fi31)
                                    _ = v35
                                default:
                                    _ = len(v3)
                                }
                                v2 = v2
                                if false {
                                    var v36 float64 = float64(float32(0.0))
                                    var v37 string = "a"
                                    _ = v36
                                    _ = v37
                                }
                            } else {
                                v1 = v11
                                v7 = 2147483648
                            }
                        case "\n\t":
                            v38 := GenS2{F0: -1, F1: (!true), F2: 42}
                            _ = v38
                        case "hello":
                            if false {
                                v39 := v3[fi27]
                                v3[fi31] = 2147483648
                                v2 += 255
                                _ = v39
                            } else {
                                var v40 int16 = int16(v2)
                                _ = len(v3)
                                _ = v40
                            }
                        }
                        if false {
                            if (fi31 >= (fi31 & (fi27 ^ (fi31 ^ -2147483648)))) {
                                var v41 int = fi31
                                v41 = (fi27 / ((-603) | 1))
                                _ = v41
                            }
                            v7 += (v7 << ((v1) & 63))
                        } else {
                            switch fi31 {
                            case 0:
                                var v42 string = "foo bar"
                                _ = v42
                            default:
                                if true {
                                    v4 = append(v4, 0)
                                    fmt.Println(fi31, v5)
                                } else {
                                    var v43 int = int(v2)
                                    _ = len(v4)
                                    _ = v43
                                }
                            }
                            var v44 float32 = 0.1
                            v45 := GenS3{F0: 902}
                            _ = v44
                            _ = v45
                        }
                        _ = fi31
                    }
                case "a":
                    fmt.Println(v1, v11, v7)
                default:
                    if true {
                        if ("\n\t" != "foo bar") {
                            var v46 uint16 = v2
                            v7 = 1
                            _ = v46
                        } else {
                            v47 := GenS3{F0: (v1 << (((^(v1))) & 63))}
                            _ = v47
                        }
                        if true {
                            var v48 uint8 = uint8(v7)
                            _ = len(v4)
                            _ = v48
                        }
                    }
                }
                if len(v4) > 0 {
                    v4[0] = (^(v7))
                }
                v11 = v1
                _ = fi27
            }
        case 32768:
            var v49 float32 = 0.0
            _ = v49
        }
    }
    switch 128 {
    case 2147483647:
        v1 = uint(v7)
    default:
        if false {
            for fi50 := 0; fi50 < 19; fi50++ {
                _ = len(v3)
                for fi51 := 0; fi51 < 18; fi51++ {
                    if (595 < 590) {
                        v52 := GenS4{F0: (1000 == 42), F1: -1, F2: 1}
                        var v53 uint64 = 42
                        _ = v52
                        _ = v53
                    } else {
                        switch -9223372036854775808 {
                        case 42:
                            var v54 bool = false
                            _ = v54
                        default:
                            var v55 uint32 = (v7 * 65535)
                            _ = v55
                        }
                    }
                    _ = fi51
                }
                _ = fi50
            }
            var v56 string = ("a" + ("\n\t" + "foo bar"))
            _ = v56
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}

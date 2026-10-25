package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 uint8; F1 int8; F2 uint8 }

func main() {
    var v1 uint8 = 128
    var v2 uint8 = (v1 - v1)
    var v3 uint64 = 2
    var v4 uint32 = 436
    var v5 uint = 65535
    v6 := []uint32{(v4 << (((^(v5))) & 63)), (v4 - (v4 << ((v5) & 63))), (v4 + 0), 4294967295}
    switch "a" {
    case "\n\t":
        switch 2147483647 {
        case 1:
            fmt.Println(v4)
        default:
            switch 4294967295 {
            case 32768:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            default:
                switch -1 {
                case 4294967295:
                    if false {
                        fmt.Println(v3)
                    } else {
                        v7 := GenS1{F0: -115, F1: -1, F2: int(v4)}
                        for fi8 := 0; fi8 < 1; fi8++ {
                            if ((^(v2)) < (^(v2))) {
                                fmt.Println(v2)
                            } else {
                                fmt.Println(v2, v4)
                                v9 := GenS1{F0: 42, F1: 1, F2: fi8}
                                if (255 < 786) {
                                    v5 += 839
                                } else {
                                    v10 := GenS1{F0: 42, F1: -1, F2: 9223372036854775807}
                                    _ = v10
                                }
                                _ = v9
                            }
                            _ = fi8
                        }
                        v11 := GenS2{F0: v3, F1: 42, F2: 0}
                        _ = v11
                        _ = v7
                    }
                default:
                    v3 = 554
                }
            }
        }
    case "foo bar":
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 uint32 = v4
            _ = fi12
            _ = v13
        }
    case "a":
        v4 = v4
    default:
        var v14 int16 = 32767
        _ = v14
    }
    var v15 uint16 = 65535
    var v16 int32 = 42
    fmt.Println(v4, v1, v5)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if true {
        if len(v6) > 0 {
            v6[0] = (^(v4))
        }
        v15 = v15
    }
    var v17 uint32 = uint32(v15)
    switch -2147483648 {
    case 128:
        _ = len(v6)
    case 9223372036854775807:
        v18 := GenS2{F0: v3, F1: 434, F2: v15}
        _ = v18
    case 127:
        var v19 uint8 = 25
        _ = v19
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi20
    }
    v4 = 42
    var v21 bool = (!true)
    switch 4294967295 {
    case 4294967295:
        if (!v21) {
            var v22 float64 = float64((0.1 * 0.1))
            v23 := GenS3{F0: 291}
            v16 += v16
            _ = v22
            _ = v23
        }
    case 0:
        switch "a" {
        case "a":
            if false {
                if false {
                    var v24 string = ((("a" + "hello") + "") + "foo bar")
                    var v25 uint64 = (^(v3))
                    _ = v24
                    _ = v25
                }
                if true {
                    _ = len(v6)
                    if len(v6) > 0 {
                        v6[0] = 2147483648
                    }
                }
                if (0.0 > (((3.14159 * float32(1e10)) * (0.0 - 3.14159)) - 0.1)) {
                    var v26 int8 = 100
                    v21 = false
                    _ = v26
                } else {
                    v27 := GenS2{F0: (v3 + v3), F1: 2, F2: (v15 << ((743) & 63))}
                    v6 = append(v6, uint32(v3))
                    for fi28 := 0; fi28 < 16; fi28++ {
                        var v29 string = "hello"
                        v30 := GenS4{F0: "\n\t", F1: 1}
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    _ = v27
                }
            } else {
                switch "a" {
                case "foo bar":
                    _ = len(v6)
                }
                v2 += 127
            }
        case "hello":
            v31 := GenS5{F0: uint8(v3), F1: 38, F2: v2}
            _ = v31
        default:
            if v21 {
                _ = len(v6)
            } else {
                for fi32 := 0; fi32 < 17; fi32++ {
                    v21 = v21
                    for fi33 := 0; fi33 < 10; fi33++ {
                        v2 -= 176
                        _ = fi33
                    }
                    v6 = append(v6, v17)
                    _ = fi32
                }
            }
        }
    default:
        _ = len(v6)
    }
    _ = len(v6)
    v34 := GenS1{F0: -12, F1: 42, F2: 128}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}

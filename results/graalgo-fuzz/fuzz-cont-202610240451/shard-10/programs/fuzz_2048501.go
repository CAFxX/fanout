package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 bool; F2 int }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 int8 }

func main() {
    var v1 float64 = (1.5 - 1e10)
    v2 := []int{127}
    var v3 int64 = -612
    var v4 int64 = v3
    var v5 string = ""
    var v6 int64 = v3
    v3 += (v3 >> ((1) & 63))
    var v7 float64 = v1
    var v8 float64 = 3.14159
    v2 = append(v2, -385)
    v4 = (v3 / (((v4 >> ((2) & 63))) | 1))
    if ((v5 != v5) || false) {
        fmt.Println(v3, v8, v4)
        v9 := GenS1{F0: 56}
        var v10 uint16 = 2
        _ = v10
        _ = v9
    }
    var v11 uint = 255
    v4 = (v6 ^ v4)
    for fi12 := 0; fi12 < 14; fi12++ {
        v8 = v8
        v13 := GenS1{F0: 42}
        v14 := GenS1{F0: 1}
        _ = fi12
        _ = v13
        _ = v14
    }
    v15 := GenS2{F0: 0, F1: (32767 != 0), F2: 1}
    var v16 uint = v11
    var v17 string = v5
    if (!false) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v11 -= (v16 / ((v16) | 1))
    }
    if true {
        v8 -= v1
        v11 += 4294967295
        v6 -= (v4 << (((v11 & (v11 - v11))) & 63))
    }
    var v18 string = ((v17 + "\n\t") + "foo bar")
    switch -445 {
    case 42:
        if true {
            for fi19 := 0; fi19 < 3; fi19++ {
                v8 -= v8
                for fi20 := 0; fi20 < 15; fi20++ {
                    var v21 uint = v16
                    v4 = (v3 + (v4 << ((v21) & 63)))
                    if false {
                        if true {
                            v22 := GenS3{F0: 0, F1: 0}
                            if true {
                                var v23 int = -603
                                if false {
                                    var v24 uint32 = 42
                                    v16 = v16
                                    _ = v24
                                } else {
                                    var v25 int32 = 383
                                    if len(v2) > 0 {
                                        v2[0] = -871
                                    }
                                    var v26 uint32 = 873
                                    fmt.Println(v26, v6)
                                    v2 = append(v2, (v23 >> ((v11) & 63)))
                                    v27 := GenS2{F0: 113, F1: (true && true), F2: v23}
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                }
                                switch v18 {
                                case "\n\t":
                                    v28 := GenS2{F0: 42, F1: (true && true), F2: fi20}
                                    _ = v28
                                case "a":
                                    v8 += ((v8 / v8) * v1)
                                case "foo bar":
                                    var v29 int8 = 100
                                    _ = v29
                                }
                                _ = v23
                            } else {
                                fmt.Println(v6, fi19, v16)
                            }
                            _ = v22
                        }
                        if true {
                            switch 0 {
                            case -250:
                                _ = len(v2)
                            case 545:
                                v30 := GenS4{F0: -21}
                                _ = v30
                            case -9223372036854775808:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                var v31 int16 = 183
                                _ = v31
                            }
                        } else {
                            var v32 bool = false
                            var v33 uint16 = 290
                            _ = v32
                            _ = v33
                        }
                        v34 := GenS1{F0: 127}
                        v4 += v4
                        _ = v34
                    }
                    _ = fi20
                    _ = v21
                }
                _ = fi19
            }
        } else {
            v11 += v11
        }
    case 0:
        var v35 bool = true
        _ = v35
    case 127:
        var v36 uint8 = 0
        _ = v36
    default:
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v3)
        default:
            var v37 uint16 = 282
            _ = v37
        }
    }
    if true {
        v1 += float64(((0.1 + float32(v1)) / 1.0))
    } else {
        v2 = append(v2, 32768)
        var v38 uint = (v16 ^ (^(v16)))
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

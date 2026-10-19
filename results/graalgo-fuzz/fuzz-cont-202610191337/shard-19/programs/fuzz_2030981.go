package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 int8; F2 uint64 }
type GenS6 struct { F0 int16 }
func genFn0(p0 int64) uint {
    v7 := GenS1{F0: true}
    v8 := []int{42, 4294967295, 242}
    var v9 float64 = -2.25
    switch "" {
    case "":
        v8 = append(v8, 750)
    case "hello":
        v10 := GenS1{F0: false}
        _ = v10
    case "foo bar":
        var v11 float64 = ((3.14159 * -2.25) * (v9 - (3.14159 + 0.0)))
        _ = v11
    }
    v9 -= 1e10
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 uint8 = 128
    var v2 bool = false
    v3 := []uint32{1}
    var v4 int8 = -10
    var v5 uint = 18446744073709551615
    var v6 uint8 = 2
    switch -9223372036854775808 {
    case 0:
        for fi12 := 0; fi12 < 14; fi12++ {
            if len(v3) > 0 {
                v3[0] = 49
            }
            var v13 uint8 = v6
            _ = fi12
            _ = v13
        }
    case 42:
        v6 = 252
    default:
        var v14 uint8 = (v1 + v6)
        _ = v14
    }
    fmt.Println(v1, v6)
    if v2 {
        var v15 float64 = float64((float32((1e10 * 0.1)) + -2.25))
        fmt.Println(v1)
        switch -9223372036854775808 {
        case 0:
            v5 = 24
        case 2147483647:
            for fi16 := 0; fi16 < 18; fi16++ {
                if v2 {
                    v15 += 1.5
                    fmt.Println(v1, v4)
                    v15 -= -2.25
                }
                for fi17 := 0; fi17 < 4; fi17++ {
                    var v18 uint16 = 0
                    v5 -= (v5 - (^(v5)))
                    _ = fi17
                    _ = v18
                }
                _ = fi16
            }
        case -9223372036854775808:
            switch -409 {
            case -736:
                fmt.Println(v1)
            case 788:
                var v19 int = 128
                _ = v19
            case -2147483648:
                v15 += v15
            }
        default:
            var v20 uint = 65535
            _ = v20
        }
        _ = v15
    }
    if true {
        if v2 {
            for fi21 := 0; fi21 < 4; fi21++ {
                if false {
                    switch "" {
                    case "\n\t":
                        if false {
                            if true {
                                v22 := GenS1{F0: v2}
                                _ = v22
                            }
                            if (!(156 <= v1)) {
                                v2 = false
                                fmt.Println(v6, v5)
                            } else {
                                switch "" {
                                case "\n\t":
                                    var v23 string = "a"
                                    _ = v23
                                }
                            }
                            if (42 > 0) {
                                v24 := GenS1{F0: v2}
                                v5 -= (v5 ^ genFn0(9223372036854775807))
                                v25 := GenS2{F0: v5}
                                _ = v24
                                _ = v25
                            } else {
                                v1 = v6
                                var v26 uint = 65535
                                _ = v26
                            }
                        }
                    }
                    var v27 int8 = v4
                    _ = v27
                }
                v28 := GenS3{F0: -9223372036854775808}
                _ = fi21
                _ = v28
            }
        }
        v29 := GenS4{F0: 917}
        _ = len(v3)
        _ = v29
    }
    v1 += 86
    v6 -= 128
    v4 = 42
    v30 := GenS2{F0: (^(v5))}
    for fi31 := 0; fi31 < 15; fi31++ {
        for fi32 := 0; fi32 < 13; fi32++ {
            var v33 uint8 = v6
            if len(v3) > 0 {
                v3[0] = 0
            }
            switch "foo bar" {
            case "hello":
                v34 := GenS5{F0: (fi31 ^ int(v6)), F1: int8(v6), F2: 0}
                _ = v34
            case "\n\t":
                var v35 int = fi31
                _ = v35
            }
            v2 = true
            v5 = genFn0(-1)
            if ((1 <= -1) && v2) {
                var v36 int64 = 0
                _ = v36
            } else {
                switch 0 {
                case -924:
                    v3 = append(v3, 631)
                case 9223372036854775807:
                    var v37 uint8 = (v6 >> ((v5) & 63))
                    _ = v37
                default:
                    v5 = 42
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v6 = 11
            }
            var v38 int = 127
            fmt.Println(v5)
            _ = fi32
            _ = v33
            _ = v38
        }
        if true {
            var v39 string = "foo bar"
            _ = v39
        }
        var v40 float64 = (1.5 - 1.5)
        _ = fi31
        _ = v40
    }
    var v41 int64 = 1
    if v2 {
        v41 = (v41 * 1)
        if (v41 > -905) {
            for fi42 := 0; fi42 < 3; fi42++ {
                v43 := GenS6{F0: 0}
                var v44 uint8 = 127
                var v45 uint32 = 42
                _ = fi42
                _ = v43
                _ = v44
                _ = v45
            }
            v6 += v1
        } else {
            v6 -= v1
        }
        fmt.Println(v2)
    }
    for fi46 := 0; fi46 < 15; fi46++ {
        v3 = append(v3, 498)
        _ = fi46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int8 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 0
    var v3 int64 = -790
    var v4 string = ""
    switch int(v2) {
    case -977:
        var v5 uint16 = 42
        _ = v5
    case -711:
        v6 := []int{-1, -488, -1, -1}
        _ = v6
    }
    v3 = v3
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 uint32 = v2
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: (v4 + v4)}
    switch "a" {
    case "\n\t":
        var v10 int32 = -402
        _ = v10
    case "a":
        v1[v4] = (134 != 4294967295)
    case "":
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 uint16 = 1
            var v13 int32 = -1
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    switch 9223372036854775807 {
    case 0:
        v14 := []int{}
        _ = v14
    case 9223372036854775807:
        switch 728 {
        case 42:
            switch 851 {
            case 9223372036854775807:
                var v15 uint32 = (v2 - (v2 & (^(v2))))
                _ = v15
            case -952:
                v2 = (v2 * (v2 * v2))
            case -566:
                switch 127 {
                case 9223372036854775807:
                    v16 := []int{127, 134}
                    _ = v16
                }
            default:
                var v17 bool = ((("\n\t" + ("hello" + "foo bar")) >= (("" + "") + (v4 + v4))) && (false || true))
                _ = v17
            }
        case 131:
            v18 := GenS1{F0: "\n\t"}
            _ = v18
        case -2147483648:
            v19 := GenS2{F0: -57, F1: 107}
            _ = v19
        default:
            delete(v1, (v4 + (v4 + v4)))
        }
    default:
        fmt.Println(v3, v4)
    }
    var v20 int16 = -305
    v21 := []int{52, -2147483648, 32767, 128}
    for fi22 := 0; fi22 < 10; fi22++ {
        if false {
            var v23 uint = 42
            switch v4 {
            case "a":
                switch fi22 {
                case -847:
                    fmt.Println(v20)
                default:
                    if false {
                        if true {
                            v23 = v23
                            if len(v21) >= 2 {
                                v21 = v21[1:len(v21)]
                            }
                            v20 -= (v20 * (^(v20)))
                        }
                        var v24 int16 = (^(v20))
                        var v25 uint8 = 222
                        _ = v24
                        _ = v25
                    } else {
                        _ = len(v21)
                    }
                }
            case "":
                var v26 bool = (true || (((-2.25 / 3.14159) - 1e10) > float64(-2.25)))
                _ = v26
            case "foo bar":
                v23 = (v23 - 0)
            }
            var v27 bool = (v4 < "foo bar")
            _ = v23
            _ = v27
        } else {
            v21 = append(v21, (-(fi22)))
            fmt.Println(v3, v2, v2)
            if (false || (fi22 >= 9223372036854775807)) {
                v4 = "foo bar"
            }
        }
        delete(v1, v4)
        _ = fi22
    }
    if len(v21) > 0 {
        v21[0] = 128
    }
    switch 9223372036854775807 {
    case -2147483648:
        switch "\n\t" {
        case "\n\t":
            var v28 string = (v4 + "")
            _ = v28
        default:
            if true {
                var v29 string = ("" + "foo bar")
                var v30 string = v29
                fmt.Println(v20, v4, v3)
                _ = v29
                _ = v30
            }
        }
    case -9223372036854775808:
        switch v4 {
        case "hello":
            var v31 uint8 = 1
            _ = v31
        default:
            for fi32 := 0; fi32 < 17; fi32++ {
                v21 = append(v21, 42)
                _ = fi32
            }
        }
    }
    for fi33 := 0; fi33 < 6; fi33++ {
        switch fi33 {
        case -1:
            if false {
                fmt.Println(fi33)
                v20 = (v20 % ((v20) | 1))
            }
        case 127:
            switch "" {
            case "\n\t":
                fmt.Println(v2, v20)
            case "hello":
                fmt.Println(v4)
            case "a":
                var v34 int = fi33
                _ = v34
            }
        case -661:
            var v35 string = "foo bar"
            _ = v35
        default:
            fmt.Println(v3, v2)
        }
        var v36 int64 = v3
        _ = fi33
        _ = v36
    }
    var v37 uint32 = 88
    v1[v4] = ((18446744073709551615 < 4294967295) && false)
    if (65535 >= 587) {
        var v38 int64 = (v3 + v3)
        _ = v38
    }
    delete(v1, "")
    var v39 uint8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v37
    _ = v39
    _ = v4
    _ = v9
}

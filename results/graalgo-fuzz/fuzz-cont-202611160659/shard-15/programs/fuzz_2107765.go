package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 uint32; F1 string; F2 int64 }

func main() {
    v1 := []uint32{949, 2147483647}
    v2 := make(map[int]bool)
    var v3 uint16 = 929
    v4 := []uint32{84, 453}
    v5 := []uint32{2, 149, 855, 92}
    var v6 uint16 = (v3 ^ 463)
    var v7 string = "foo bar"
    switch v7 {
    case "hello":
        for fi8 := 0; fi8 < 7; fi8++ {
            switch v7 {
            case "foo bar":
                var v9 int32 = 0
                _ = v9
            case "\n\t":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                fmt.Println(fi8)
            }
            var v10 uint64 = 0
            switch 2147483647 {
            case 1:
                fmt.Println(v10, fi8)
            }
            _ = fi8
            _ = v10
        }
    }
    switch -1 {
    case 101:
        v6 -= v6
    case 32768:
        var v11 int64 = 0
        _ = v11
    }
    fmt.Println(v7, v6, v7)
    if false {
        if true {
            _ = len(v1)
            for fi12 := 0; fi12 < 8; fi12++ {
                v3 = (v3 - (v3 * v3))
                fmt.Println(v7, v6)
                v13 := GenS1{F0: -9223372036854775808}
                _ = fi12
                _ = v13
            }
            v14 := GenS2{F0: (548 > 1), F1: 0}
            _ = v14
        } else {
            v15 := GenS2{F0: (!(!(-581 >= -398))), F1: 0}
            _ = v15
        }
    }
    switch -9223372036854775808 {
    case 71:
        v3 -= (v3 >> ((101) & 63))
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        if len(v1) > 0 {
            v1[0] = 52
        }
        if (true && (!((v7 > (v7 + "foo bar")) && false))) {
            for fi17 := 0; fi17 < 1; fi17++ {
                v7 = ("" + "hello")
                _ = fi17
            }
            for fi18 := 0; fi18 < 3; fi18++ {
                var v19 float64 = (0.0 - (1e10 * (float64(0.1) / float64(0.1))))
                v7 = v7
                _ = fi18
                _ = v19
            }
            for fi20 := 0; fi20 < 6; fi20++ {
                var v21 int64 = 42
                v3 = 949
                fmt.Println(fi16)
                _ = fi20
                _ = v21
            }
        } else {
            v6 = v6
            v6 = v3
            var v22 int8 = 127
            _ = v22
        }
        if len(v1) > 0 {
            v1[0] = 2
        }
        _ = fi16
    }
    var v23 uint16 = (v3 - (^(v6)))
    if true {
        fmt.Println(v3)
        for fi24 := 0; fi24 < 12; fi24++ {
            if false {
                var v25 int16 = 1000
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v25
            }
            _ = fi24
        }
    } else {
        if (4294967295 >= 1) {
            var v26 uint64 = 42
            _ = v26
        }
    }
    v1 = append(v1, 312)
    v3 += v3
    if ((!false) && (!(!(-429 <= 476)))) {
        var v27 int64 = 0
        v28 := GenS2{F0: false, F1: v27}
        v6 = v3
        _ = v27
        _ = v28
    } else {
        var v29 uint8 = 2
        _ = v29
    }
    v3 -= v3
    v30 := GenS3{F0: 42, F1: "foo bar", F2: -9223372036854775808}
    var v31 uint16 = v23
    if (-266 < 1000) {
        if true {
            v31 = (v3 + (v3 + v6))
            v31 -= v31
            v23 += v6
        } else {
            v6 = 42
        }
    } else {
        fmt.Println(v6)
    }
    if (-2147483648 != 42) {
        v31 = (v6 ^ 2)
        for fi32 := 0; fi32 < 17; fi32++ {
            switch fi32 {
            case 2147483647:
                if true {
                    v33 := GenS2{F0: ((v3 >= 1) && false), F1: 1}
                    v34 := GenS1{F0: 9223372036854775807}
                    v1 = append(v1, 2147483647)
                    _ = v33
                    _ = v34
                }
            default:
                if true {
                    fmt.Println(v3, v31, v6)
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v7 = ""
                }
            }
            var v35 int8 = 81
            _ = fi32
            _ = v35
        }
        v7 = "hello"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

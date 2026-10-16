package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int32; F2 int64 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 int }

func main() {
    var v1 float64 = 1e10
    var v2 uint8 = 2
    var v3 float32 = -2.25
    v4 := []int{32767, 284, 1, -9223372036854775808}
    var v5 uint = 209
    v6 := []string{"\n\t", ""}
    if false {
        var v7 int = -1
        _ = len(v4)
        _ = v7
    } else {
        var v8 uint = v5
        v3 -= 1.5
        _ = v8
    }
    if (uint16(v5) != 42) {
        var v9 uint64 = 1
        fmt.Println(v2, v1)
        switch 4294967295 {
        case -989:
            switch 128 {
            case 2147483647:
                var v10 int16 = -32768
                _ = v10
            case 32767:
                var v11 uint16 = 773
                _ = v11
            case 926:
                for fi12 := 0; fi12 < 13; fi12++ {
                    v13 := GenS1{F0: 0}
                    v5 = uint(v2)
                    v5 = v5
                    _ = fi12
                    _ = v13
                }
            default:
                for fi14 := 0; fi14 < 13; fi14++ {
                    var v15 uint = (v5 + (v5 | 912))
                    _ = fi14
                    _ = v15
                }
            }
        }
        _ = v9
    }
    switch -9223372036854775808 {
    case 823:
        var v16 float64 = float64(((v3 - 0.0) / v3))
        _ = v16
    default:
        switch "hello" {
        case "a":
            var v17 uint16 = 949
            _ = v17
        }
    }
    v18 := GenS2{F0: 42, F1: 0, F2: -285}
    switch ("" + ("\n\t" + "hello")) {
    case "\n\t":
        v19 := GenS1{F0: 1}
        _ = v19
    case "":
        if (!true) {
            v2 -= v2
            var v20 int32 = 42
            _ = v20
        } else {
            for fi21 := 0; fi21 < 9; fi21++ {
                fmt.Println(v3)
                _ = fi21
            }
            var v22 int16 = 1000
            if ((9223372036854775807 != -1) && (0 <= -1)) {
                v2 = (v2 >> ((uint(v2)) & 63))
            }
            _ = v22
        }
    default:
        var v23 uint8 = v2
        _ = v23
    }
    if ((true || true) || (!false)) {
        var v24 uint16 = 985
        var v25 string = "a"
        v26 := GenS1{F0: 42}
        _ = v24
        _ = v25
        _ = v26
    } else {
        for fi27 := 0; fi27 < 14; fi27++ {
            for fi28 := 0; fi28 < 2; fi28++ {
                var v29 int32 = -2147483648
                v3 = v3
                _ = fi28
                _ = v29
            }
            for fi30 := 0; fi30 < 11; fi30++ {
                v5 += v5
                v2 -= (v2 | v2)
                _ = fi30
            }
            fmt.Println(fi27, fi27)
            _ = fi27
        }
        v2 = (^(v2))
        v31 := GenS1{F0: int64(v2)}
        _ = v31
    }
    for fi32 := 0; fi32 < 5; fi32++ {
        if len(v4) > 0 {
            v4[0] = (fi32 + (-(fi32)))
        }
        for fi33 := 0; fi33 < 1; fi33++ {
            switch "" {
            case "hello":
                v5 = (v5 & v5)
            }
            var v34 int64 = -1
            var v35 int = (fi32 % ((-204) | 1))
            _ = fi33
            _ = v34
            _ = v35
        }
        v36 := GenS3{F0: -552, F1: false}
        _ = fi32
        _ = v36
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v37 int = -1
    var v38 float64 = (v1 - v1)
    var v39 string = ""
    switch -2147483648 {
    case -910:
        v37 -= (v37 + v37)
    case 128:
        if ((v3 - v3) == v3) {
            var v40 string = "hello"
            v41 := GenS4{F0: (v37 / ((32768) | 1))}
            for fi42 := 0; fi42 < 11; fi42++ {
                if ((!false) || false) {
                    v4 = append(v4, -863)
                } else {
                    for fi43 := 0; fi43 < 5; fi43++ {
                        var v44 string = v39
                        v38 = v1
                        _ = fi43
                        _ = v44
                    }
                }
                _ = fi42
            }
            _ = v40
            _ = v41
        } else {
            switch 1 {
            case 127:
                var v45 uint16 = 533
                _ = v45
            case 0:
                v37 = (^(v37))
            case -1:
                for fi46 := 0; fi46 < 18; fi46++ {
                    v3 = ((v3 / v3) * (v3 + (((v3 - -2.25) - 0.1) - 1.5)))
                    var v47 float64 = v1
                    _ = fi46
                    _ = v47
                }
            default:
                for fi48 := 0; fi48 < 1; fi48++ {
                    v38 = float64((v3 / 1e10))
                    var v49 int64 = int64(fi48)
                    _ = fi48
                    _ = v49
                }
            }
            var v50 string = v39
            _ = v50
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}

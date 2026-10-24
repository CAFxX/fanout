package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 uint16, p1 int64, p2 uint) uint64 {
    var v6 uint8 = 127
    for fi7 := 0; fi7 < 7; fi7++ {
        v6 -= uint8(p0)
        _ = fi7
    }
    v8 := GenS1{F0: 32768}
    if ((p1 | 127) > p1) {
        fmt.Println(p2)
    } else {
        if true {
            switch 377 {
            case -9223372036854775808:
                p1 = (p1 & 42)
            default:
                var v9 string = "a"
                _ = v9
            }
            fmt.Println(v6)
            v10 := []int{128, -999}
            _ = v10
        }
        var v11 int32 = 42
        _ = v11
    }
    p2 += 0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v8
    return 42
}

func main() {
    v1 := []string{"hello"}
    var v2 uint = 649
    var v3 float64 = 3.14159
    var v4 float64 = -2.25
    var v5 int16 = 0
    if (0 != 0) {
        for fi12 := 0; fi12 < 13; fi12++ {
            fmt.Println(v5, v3)
            _ = fi12
        }
    }
    var v13 int8 = 1
    var v14 bool = (32767 > (-(v5)))
    if (1 < 481) {
        var v15 float32 = 1.5
        var v16 uint64 = 2
        v5 += 42
        _ = v15
        _ = v16
    } else {
        var v17 uint = 2
        if ("a" != "foo bar") {
            fmt.Println(v3)
            for fi18 := 0; fi18 < 12; fi18++ {
                for fi19 := 0; fi19 < 20; fi19++ {
                    fmt.Println(v4)
                    _ = fi19
                }
                _ = fi18
            }
            v20 := GenS1{F0: 243}
            _ = v20
        } else {
            if v14 {
                var v21 bool = (3.14159 <= (0.0 - float64((1.5 + 0.0))))
                v21 = v14
                _ = v21
            }
            v17 = (v17 >> ((2) & 63))
            var v22 int = -762
            _ = v22
        }
        v23 := GenS2{F0: 42, F1: v13, F2: 42}
        _ = v17
        _ = v23
    }
    v2 = v2
    switch "" {
    case "hello":
        v13 -= (v13 * (v13 + (v13 + v13)))
    case "foo bar":
        if (((float64(-2.25) / 1e10) + float64(0.1)) != float64(1.5)) {
            switch "foo bar" {
            case "\n\t":
                var v24 float32 = 0.1
                _ = v24
            case "hello":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            if ((1e10 / 1e10) > 0.0) {
                v25 := GenS3{F0: 42}
                var v26 int16 = -32768
                _ = v25
                _ = v26
            } else {
                for fi27 := 0; fi27 < 12; fi27++ {
                    fmt.Println(v4, v4)
                    _ = fi27
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
    case "a":
        switch 1 {
        case 528:
            for fi28 := 0; fi28 < 2; fi28++ {
                var v29 uint = (v2 >> ((858) & 63))
                _ = fi28
                _ = v29
            }
        case 9223372036854775807:
            for fi30 := 0; fi30 < 5; fi30++ {
                v31 := GenS2{F0: 42, F1: v13, F2: 255}
                _ = fi30
                _ = v31
            }
        case 664:
            var v32 int = 2147483647
            _ = v32
        }
    default:
        if (1 <= 2147483648) {
            for fi33 := 0; fi33 < 3; fi33++ {
                if (v14 && ((v13 & (v13 - (v13 | v13))) == (-(v13)))) {
                    for fi34 := 0; fi34 < 15; fi34++ {
                        v14 = ((0.1 < ((((1.5 / 1.0) * -2.25) * 1e10) / (float32(v3) + 1e10))) && v14)
                        _ = fi34
                    }
                    v35 := GenS3{F0: 0}
                    _ = v35
                }
                _ = fi33
            }
            v13 += (v13 >> ((v2) & 63))
            v36 := GenS3{F0: 32767}
            _ = v36
        } else {
            _ = len(v1)
            fmt.Println(v2)
            var v37 int32 = 1
            _ = v37
        }
    }
    v38 := GenS4{F0: 738}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch "hello" {
    case "foo bar":
        for fi39 := 0; fi39 < 9; fi39++ {
            v40 := GenS4{F0: 1}
            switch ("hello" + ("\n\t" + (("\n\t" + "foo bar") + ("foo bar" + "")))) {
            case "\n\t":
                if true {
                    var v41 uint = v2
                    v13 += (^(v13))
                    _ = v41
                } else {
                    if v14 {
                        for fi42 := 0; fi42 < 9; fi42++ {
                            fmt.Println(v3, v14, v2)
                            _ = fi42
                        }
                    }
                    var v43 uint8 = 54
                    v4 = v3
                    _ = v43
                }
            }
            _ = fi39
            _ = v40
        }
    default:
        v13 = 127
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}

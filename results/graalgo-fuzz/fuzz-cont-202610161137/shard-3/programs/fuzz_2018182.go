package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 int8; F2 int8 }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 uint; F1 int8; F2 string }
type GenS6 struct { F0 uint32; F1 int32 }
type GenS7 struct { F0 uint16 }
func genFn0() bool {
    for fi7 := 0; fi7 < 18; fi7++ {
        fmt.Println(fi7)
        _ = fi7
    }
    v8 := GenS1{F0: -2147483648, F1: 799, F2: "a"}
    var v9 int8 = 1
    v9 = -1
    _ = v8
    _ = v9
    return (((-2.25 / 1e10) * 3.14159) != (1.5 - float32(3.14159)))
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = ((!false) || true)
    var v3 bool = ("hello" > "")
    var v4 uint64 = 18446744073709551615
    var v5 uint16 = 0
    var v6 int8 = 127
    var v10 int16 = 0
    var v11 bool = v2
    v12 := GenS2{F0: 2}
    var v13 bool = v2
    for fi14 := 0; fi14 < 11; fi14++ {
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 uint8 = 127
            _ = fi15
            _ = v16
        }
        if (0 != 0) {
            v1["\n\t"] = 2147483648
            switch (fi14 & (fi14 | (fi14 % ((fi14) | 1)))) {
            case -53:
                var v17 uint64 = (v4 % ((v4) | 1))
                _ = v17
            }
        } else {
            for fi18 := 0; fi18 < 2; fi18++ {
                if (204 >= 2) {
                    v19 := GenS3{F0: "foo bar", F1: (v6 + (v6 + (v6 - v6))), F2: (v6 * v6)}
                    _ = v19
                }
                switch fi18 {
                case 0:
                    v20 := GenS4{F0: 0, F1: v4}
                    _ = v20
                case 9223372036854775807:
                    v21 := GenS2{F0: 4294967295}
                    _ = v21
                default:
                    v22 := []int{9223372036854775807, (^(fi18)), (^(fi14))}
                    v1["hello"] = 0
                    _ = v22
                }
                _ = fi18
            }
            v3 = v13
        }
        _ = fi14
    }
    v23 := []int{}
    if v2 {
        fmt.Println(v10, v11, v6)
        var v24 int8 = v6
        var v25 int64 = 9223372036854775807
        _ = v24
        _ = v25
    }
    v11 = true
    v4 = 42
    v4 = v4
    switch "foo bar" {
    case "a":
        if false {
            for fi26 := 0; fi26 < 15; fi26++ {
                var v27 int = (fi26 + 127)
                _ = fi26
                _ = v27
            }
            v28 := GenS5{F0: 18446744073709551615, F1: v6, F2: "hello"}
            _ = v28
        }
    case "":
        _ = len(v1)
    case "foo bar":
        var v29 bool = v2
        _ = v29
    }
    v30 := v1["foo bar"]
    var v31 int32 = -340
    v4 += v4
    if ((v6 * v6) < 0) {
        switch "foo bar" {
        case "foo bar":
            v32 := GenS6{F0: (^(v30)), F1: 2147483647}
            _ = v32
        case "\n\t":
            var v33 int8 = 1
            _ = v33
        case "":
            switch 9223372036854775807 {
            case 4294967295:
                for fi34 := 0; fi34 < 20; fi34++ {
                    if v11 {
                        var v35 string = "foo bar"
                        switch (fi34 - 9223372036854775807) {
                        case 316:
                            var v36 int = (fi34 & 127)
                            _ = v36
                        default:
                            for fi37 := 0; fi37 < 10; fi37++ {
                                v38 := GenS5{F0: 18446744073709551615, F1: v6, F2: ("\n\t" + v35)}
                                _ = fi37
                                _ = v38
                            }
                        }
                        _ = v35
                    }
                    fmt.Println(v5, v6)
                    var v39 uint = 2
                    _ = fi34
                    _ = v39
                }
            case 0:
                var v40 float64 = 0.1
                _ = v40
            default:
                var v41 uint64 = 349
                _ = v41
            }
        }
        if (!((0.0 + 1e10) <= (0.0 + (3.14159 / (-2.25 * 0.1))))) {
            var v42 string = "foo bar"
            var v43 int32 = (v31 - v31)
            _ = v42
            _ = v43
        }
        if ((uint8(v30) <= 1) && (-2.25 > 0.0)) {
            v4 = (v4 - v4)
            var v44 uint64 = 133
            _ = v44
        }
    } else {
        if (!(0.0 < 3.14159)) {
            v45 := v1[""]
            for fi46 := 0; fi46 < 4; fi46++ {
                v11 = true
                _ = fi46
            }
            _ = v45
        } else {
            var v47 uint8 = 127
            if v3 {
                v13 = false
                v48 := GenS2{F0: uint32(v4)}
                var v49 int = -2147483648
                _ = v48
                _ = v49
            }
            v50 := GenS6{F0: v30, F1: -2147483648}
            _ = v47
            _ = v50
        }
        fmt.Println(v4)
        var v51 float64 = 1e10
        _ = v51
    }
    v10 += v10
    switch -698 {
    case -1:
        if (!(float64(1e10) <= (1.5 / (1e10 - (1e10 + -2.25))))) {
            fmt.Println(v2, v2, v11)
        }
    case -210:
        for fi52 := 0; fi52 < 15; fi52++ {
            if false {
                v5 = v5
                v6 = 81
            }
            switch -2147483648 {
            case 1:
                for fi53 := 0; fi53 < 16; fi53++ {
                    fmt.Println(v31, fi53, v30)
                    _ = fi53
                }
            case -2147483648:
                fmt.Println(v11, v11)
            case 4294967295:
                fmt.Println(v4, v4)
            default:
                for fi54 := 0; fi54 < 12; fi54++ {
                    fmt.Println(v31, v10)
                    _ = fi54
                }
            }
            v55 := GenS7{F0: v5}
            _ = fi52
            _ = v55
        }
    default:
        fmt.Println(v4, v31)
    }
    v23 = append(v23, 32767)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}

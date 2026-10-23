package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 int64; F1 string; F2 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint; F1 uint64 }
func genFn0(p0 int64, p1 int64, p2 bool) uint16 {
    if p2 {
        var v6 string = "\n\t"
        _ = v6
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: 826, F1: (!(660 > 42))}
        v9 := GenS1{F0: 0, F1: p2}
        fmt.Println(p1, p0)
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1(p0 string) string {
    for fi10 := 0; fi10 < 15; fi10++ {
        if false {
            v11 := []int{(fi10 - fi10), fi10}
            _ = v11
        } else {
            var v12 float64 = 3.14159
            _ = v12
        }
        p0 = ((p0 + p0) + p0)
        if false {
            switch p0 {
            case "\n\t":
                v13 := []int{}
                _ = v13
            }
            var v14 int32 = 0
            _ = v14
        }
        _ = fi10
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        if (179 < 944) {
            if true {
                for fi16 := 0; fi16 < 13; fi16++ {
                    p0 = (((p0 + p0) + "hello") + (p0 + p0))
                    v17 := []int{(fi16 - 9223372036854775807), (fi15 - fi16), fi16}
                    fmt.Println(p0, fi15)
                    _ = fi16
                    _ = v17
                }
                for fi18 := 0; fi18 < 9; fi18++ {
                    v19 := []int{fi15, fi18, fi15}
                    var v20 int16 = 1000
                    fmt.Println(v20, p0)
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                switch (p0 + p0) {
                case "\n\t":
                    if (-128 < -66) {
                        fmt.Println(fi15, p0)
                    }
                default:
                    var v21 uint64 = 4294967295
                    _ = v21
                }
            }
        }
        _ = fi15
    }
    p0 = "foo bar"
    if false {
        v22 := GenS1{F0: 2, F1: (float64(((0.0 + 1e10) / 0.1)) < 0.0)}
        var v23 uint32 = 65535
        switch p0 {
        case "hello":
            for fi24 := 0; fi24 < 1; fi24++ {
                fmt.Println(v23, p0)
                fmt.Println(v23, fi24, p0)
                _ = fi24
            }
        case "a":
            var v25 float64 = (1.5 - float64(1e10))
            _ = v25
        case "foo bar":
            var v26 int8 = 42
            _ = v26
        default:
            switch ("a" + (p0 + p0)) {
            case "\n\t":
                switch p0 {
                case "foo bar":
                    var v27 int16 = 0
                    _ = v27
                case "\n\t":
                    if (-1 < 386) {
                        v28 := GenS2{F0: (("" + "foo bar") + (p0 + p0)), F1: 127}
                        _ = v28
                    } else {
                        var v29 int = 127
                        _ = v29
                    }
                default:
                    if true {
                        v30 := []int{1, 1, 0}
                        _ = v30
                    }
                }
            case "a":
                for fi31 := 0; fi31 < 10; fi31++ {
                    var v32 int32 = int32(fi31)
                    _ = fi31
                    _ = v32
                }
            case "hello":
                fmt.Println(v23)
            default:
                v33 := GenS3{F0: 42, F1: false}
                _ = v33
            }
        }
        _ = v22
        _ = v23
    } else {
        v34 := []int{}
        _ = v34
    }
    v35 := GenS3{F0: -1, F1: (!true)}
    _ = p0
    _ = v35
    return "hello"
}

func main() {
    var v1 int16 = 1
    var v2 int16 = (v1 * v1)
    var v3 int32 = 0
    var v4 int = 2147483647
    v5 := []int{}
    for fi36 := 0; fi36 < 3; fi36++ {
        var v37 string = "a"
        switch (fi36 & v4) {
        case 1:
            switch v4 {
            case -767:
                for fi38 := 0; fi38 < 11; fi38++ {
                    v37 = (v37 + (((v37 + v37) + "a") + (v37 + "a")))
                    _ = fi38
                }
            case -530:
                for fi39 := 0; fi39 < 1; fi39++ {
                    var v40 uint32 = 129
                    v1 -= v2
                    v1 = (^(v1))
                    if (0 > (v4 | (fi36 * (fi36 | 484)))) {
                        v41 := GenS4{F0: -9223372036854775808, F1: v37, F2: 26}
                        _ = v41
                    }
                    _ = fi39
                    _ = v40
                }
            case -2147483648:
                var v42 int8 = int8(v4)
                _ = v42
            }
        case 32768:
            if len(v5) > 0 {
                v5[0] = fi36
            }
        case 548:
            v43 := GenS2{F0: v37, F1: v4}
            _ = v43
        }
        _ = fi36
        _ = v37
    }
    for fi44 := 0; fi44 < 12; fi44++ {
        var v45 uint64 = 4294967295
        if ("" == "a") {
            switch (v4 - (fi44 << ((1) & 63))) {
            case 1:
                v5 = append(v5, fi44)
            case 0:
                var v46 uint64 = v45
                _ = v46
            default:
                if true {
                    v47 := GenS3{F0: 118, F1: (((true && (0 == 0)) && (!(false || (42 > 42)))) || ((false || ((42 != 96) || (0 >= 100))) || false))}
                    v48 := GenS5{F0: 127}
                    v5 = append(v5, (fi44 % ((v4) | 1)))
                    _ = v47
                    _ = v48
                }
            }
            v49 := GenS6{F0: 983, F1: v45}
            _ = v49
        } else {
            fmt.Println(v45, v3, v2)
        }
        _ = fi44
        _ = v45
    }
    if len(v5) > 0 {
        v5[0] = v4
    }
    if true {
        fmt.Println(v2, v4)
    } else {
        v50 := GenS4{F0: -9223372036854775808, F1: ("foo bar" + "hello"), F2: 425}
        v3 = (v3 * 1)
        for fi51 := 0; fi51 < 18; fi51++ {
            v4 = (fi51 + fi51)
            v52 := GenS5{F0: 255}
            _ = fi51
            _ = v52
        }
        _ = v50
    }
    v1 = v2
    var v53 uint64 = 420
    if len(v5) > 0 {
        v5[0] = 443
    }
    v2 += (v2 << ((2) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
}

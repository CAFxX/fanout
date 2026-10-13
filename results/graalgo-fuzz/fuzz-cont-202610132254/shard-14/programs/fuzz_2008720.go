package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 string; F1 int; F2 uint8 }
type GenS3 struct { F0 int; F1 int32; F2 uint32 }
type GenS4 struct { F0 uint16; F1 uint8; F2 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64 }
func genFn0(p0 int16, p1 bool, p2 int8) int {
    if p1 {
        var v6 int64 = 322
        for fi7 := 0; fi7 < 10; fi7++ {
            for fi8 := 0; fi8 < 8; fi8++ {
                var v9 int8 = 100
                v6 += 178
                _ = fi8
                _ = v9
            }
            if (p1 || (p1 || p1)) {
                for fi10 := 0; fi10 < 12; fi10++ {
                    fmt.Println(p1)
                    _ = fi10
                }
            }
            _ = fi7
        }
        _ = v6
    }
    p1 = ((((-2.25 / 1.0) * (1.5 - 0.0)) + -2.25) >= float64(0.1))
    switch "\n\t" {
    case "foo bar":
        fmt.Println(p2)
    case "a":
        for fi11 := 0; fi11 < 9; fi11++ {
            v12 := GenS1{F0: 1, F1: -121}
            fmt.Println(fi11)
            p0 = p0
            _ = fi11
            _ = v12
        }
    case "":
        var v13 string = "foo bar"
        _ = v13
    default:
        switch -2147483648 {
        case 773:
            var v14 int = 1
            _ = v14
        case 703:
            p2 -= (-(p2))
        default:
            var v15 int64 = -9223372036854775808
            _ = v15
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 32768
}
func genFn1(p0 uint64, p1 int, p2 int32) int16 {
    var v16 bool = (true && (!(!true)))
    var v17 int8 = int8(p0)
    switch (p1 & (p1 << ((1) & 63))) {
    case 9223372036854775807:
        v16 = v16
    case -9223372036854775808:
        fmt.Println(p0)
    case 104:
        p1 = genFn0(-195, ((p0 % ((p0) | 1)) <= p0), v17)
    default:
        v18 := GenS2{F0: "\n\t", F1: (-(p1)), F2: 255}
        _ = v18
    }
    if true {
        if v16 {
            if (0 != 2) {
                if (-93 <= -2147483648) {
                    switch "a" {
                    case "\n\t":
                        if v16 {
                            var v19 float64 = float64((1.5 * 0.1))
                            p0 -= 42
                            var v20 bool = v16
                            _ = v19
                            _ = v20
                        }
                    case "a":
                        var v21 float64 = 3.14159
                        _ = v21
                    case "hello":
                        var v22 uint16 = 674
                        _ = v22
                    default:
                        var v23 uint16 = 65535
                        _ = v23
                    }
                    var v24 float32 = 0.1
                    _ = v24
                } else {
                    v25 := []int{int(p2)}
                    v17 -= v17
                    var v26 uint8 = 255
                    _ = v25
                    _ = v26
                }
                p0 = 18446744073709551615
                switch p1 {
                case 9223372036854775807:
                    v16 = (v16 && v16)
                case 128:
                    p0 = (^(p0))
                case 975:
                    if true {
                        v27 := GenS1{F0: p2, F1: 9223372036854775807}
                        for fi28 := 0; fi28 < 14; fi28++ {
                            v29 := GenS1{F0: (p2 % ((p2) | 1)), F1: -9223372036854775808}
                            _ = fi28
                            _ = v29
                        }
                        _ = v27
                    }
                default:
                    v30 := []int{(p1 ^ (p1 % ((-1) | 1))), genFn0(0, v16, 69)}
                    _ = v30
                }
            }
            var v31 bool = (-9223372036854775808 < 1)
            _ = v31
        } else {
            v32 := []int{p1, p1, int(p0), (p1 & p1)}
            v33 := GenS2{F0: "foo bar", F1: (-(p1)), F2: 7}
            for fi34 := 0; fi34 < 9; fi34++ {
                fmt.Println(v17, p2, v16)
                v32 = append(v32, (fi34 | (fi34 % ((p1) | 1))))
                fmt.Println(p2)
                _ = fi34
            }
            _ = v32
            _ = v33
        }
        var v35 int = (p1 / ((p1) | 1))
        fmt.Println(p2, v16, v35)
        _ = v35
    } else {
        for fi36 := 0; fi36 < 18; fi36++ {
            switch "a" {
            case "foo bar":
                v37 := GenS3{F0: (fi36 / ((genFn0(32767, (((float64(3.14159) + 1.5) * -2.25) != 3.14159), (v17 ^ v17))) | 1)), F1: (-(p2)), F2: 490}
                _ = v37
            case "\n\t":
                p0 = (p0 % ((p0) | 1))
            case "":
                p2 -= (^(p2))
            default:
                v16 = v16
            }
            var v38 float32 = 0.0
            switch "hello" {
            case "\n\t":
                v39 := GenS4{F0: 255, F1: 44, F2: ""}
                _ = v39
            case "":
                var v40 float32 = v38
                _ = v40
            }
            _ = fi36
            _ = v38
        }
    }
    v41 := GenS5{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    _ = v41
    return -32768
}

func main() {
    var v1 int = 1
    var v2 uint64 = 0
    v3 := make(map[string]uint32)
    var v4 int8 = 0
    var v5 uint64 = 0
    v42 := GenS3{F0: -722, F1: 42, F2: 1}
    v43 := GenS5{F0: 1}
    for fi44 := 0; fi44 < 8; fi44++ {
        delete(v3, ("\n\t" + "hello"))
        _ = fi44
    }
    v1 = 42
    v45 := []int{}
    var v46 int8 = -118
    v1 += (v1 - v1)
    v2 = v5
    v46 = v46
    v4 = (-(v4))
    delete(v3, "hello")
    switch (v1 | (v1 << ((2) & 63))) {
    case 510:
        if len(v45) >= 2 {
            v45 = v45[1:len(v45)]
        }
    case -2147483648:
        v46 = (^(v4))
    case -287:
        switch ("a" + "foo bar") {
        case "foo bar":
            switch (v1 << ((0) & 63)) {
            case 761:
                fmt.Println(v1, v1)
            case 127:
                v47 := GenS5{F0: 380}
                _ = v47
            case -600:
                v3[("\n\t" + ("" + "foo bar"))] = 716
            default:
                v2 = (v2 + 0)
            }
        case "":
            for fi48 := 0; fi48 < 12; fi48++ {
                fmt.Println(v2, v46, fi48)
                v5 -= (v5 & 42)
                if len(v45) > 0 {
                    v45[0] = (fi48 % (((v1 & genFn0(42, true, v4))) | 1))
                }
                _ = fi48
            }
        default:
            switch "a" {
            case "\n\t":
                v4 -= (-(v4))
            case "a":
                fmt.Println(v1)
            case "":
                v49 := GenS6{F0: v5}
                _ = v49
            default:
                var v50 float64 = 1e10
                _ = v50
            }
        }
    default:
        if (!(111 != 42)) {
            v5 += v5
        } else {
            var v51 uint = 533
            _ = v51
        }
    }
    v52 := GenS2{F0: "\n\t", F1: (v1 & (-(v1))), F2: 127}
    v4 += -81
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v45
    _ = v46
    _ = v5
    _ = v52
}

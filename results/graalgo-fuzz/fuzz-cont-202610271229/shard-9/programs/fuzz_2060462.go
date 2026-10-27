package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 uint }
type GenS5 struct { F0 uint16; F1 int }
func genFn0() string {
    var v7 uint8 = 128
    fmt.Println(v7)
    var v8 int = 32767
    if (v8 > v8) {
        v9 := GenS1{F0: 42, F1: 2, F2: 42}
        for fi10 := 0; fi10 < 18; fi10++ {
            fmt.Println(v8)
            v8 = (-(fi10))
            for fi11 := 0; fi11 < 12; fi11++ {
                if (true && true) {
                    fmt.Println(fi10)
                    var v12 int64 = 0
                    v8 += 1
                    var v13 string = ("" + ("\n\t" + (("" + (("\n\t" + "a") + "a")) + "hello")))
                    _ = v12
                    _ = v13
                }
                if (42 == 18446744073709551615) {
                    var v14 float64 = (1e10 + 1.5)
                    v15 := GenS2{F0: 100, F1: 12}
                    _ = v14
                    _ = v15
                } else {
                    var v16 int = v8
                    v17 := []int{(v8 - (fi10 - 908)), v16, (fi10 & int(v7)), (v8 * (fi11 / ((fi11) | 1)))}
                    var v18 float64 = 0.0
                    _ = v16
                    _ = v17
                    _ = v18
                }
                _ = fi11
            }
            _ = fi10
        }
        _ = v9
    }
    _ = v7
    _ = v8
    return "a"
}

func main() {
    v1 := []int{}
    v2 := []string{"", "hello", "foo bar"}
    var v3 uint16 = 42
    v4 := []string{"foo bar", ("foo bar" + "\n\t"), "hello", "hello"}
    var v5 float32 = (3.14159 * (0.0 + 3.14159))
    v6 := make(map[int]int)
    v19 := GenS2{F0: 6, F1: -179}
    switch 526 {
    case 990:
        v3 = (v3 * (v3 << ((120) & 63)))
    case 42:
        v3 -= (v3 * (v3 ^ v3))
    case 127:
        v3 = (v3 + v3)
    default:
        switch 9223372036854775807 {
        case 32768:
            _ = len(v6)
        case 0:
            var v20 bool = (!(((2147483648 <= 670) || true) || false))
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        if true {
            switch (fi21 ^ fi21) {
            case 2147483647:
                fmt.Println(v3, v3)
            default:
                if false {
                    fmt.Println(fi21)
                    for fi22 := 0; fi22 < 16; fi22++ {
                        if (-1 >= 2147483647) {
                            fmt.Println(v3, v3, fi21)
                            v23 := GenS1{F0: v3, F1: 927, F2: 4294967295}
                            v24 := GenS1{F0: v3, F1: 2, F2: 42}
                            _ = v23
                            _ = v24
                        }
                        switch fi22 {
                        case -683:
                            v3 -= 2
                        case 9223372036854775807:
                            if (false && (!true)) {
                                switch (fi22 - fi21) {
                                case 32767:
                                    fmt.Println(v5, fi21)
                                case -2147483648:
                                    switch (("a" + "") + "a") {
                                    case "\n\t":
                                        v3 += 480
                                    case "a":
                                        if true {
                                            if false {
                                                delete(v6, (fi22 >> ((4294967295) & 63)))
                                            }
                                            v5 = (0.0 - v5)
                                        }
                                    }
                                default:
                                    v25 := v6[(fi22 / ((398) | 1))]
                                    _ = v25
                                }
                                v5 = ((v5 * -2.25) * v5)
                            } else {
                                v26 := v6[(fi21 % ((fi22) | 1))]
                                var v27 int = (fi21 ^ 276)
                                _ = v26
                                _ = v27
                            }
                        }
                        var v28 int8 = -128
                        _ = fi22
                        _ = v28
                    }
                }
            }
        }
        v3 += (v3 % (((^(v3))) | 1))
        _ = fi21
    }
    v3 += 448
    if true {
        v6[2147483647] = -715
        var v29 int8 = 127
        _ = v29
    }
    fmt.Println(v3, v3)
    v4 = append(v4, "a")
    v30 := GenS3{F0: -32768}
    for fi31 := 0; fi31 < 5; fi31++ {
        v5 = v5
        for fi32 := 0; fi32 < 16; fi32++ {
            _ = len(v1)
            _ = fi32
        }
        fmt.Println(v5, v3, v5)
        _ = fi31
    }
    switch ((("a" + "") + "foo bar") + "") {
    case "foo bar":
        var v33 uint8 = 1
        _ = v33
    case "":
        for fi34 := 0; fi34 < 15; fi34++ {
            v35 := GenS4{F0: 2147483647, F1: 255}
            _ = fi34
            _ = v35
        }
    }
    for fi36 := 0; fi36 < 17; fi36++ {
        var v37 int64 = 9223372036854775807
        fmt.Println(v37, v5)
        var v38 int8 = 100
        _ = fi36
        _ = v37
        _ = v38
    }
    v5 += 3.14159
    switch "a" {
    case "a":
        v39 := GenS5{F0: v3, F1: 42}
        _ = v39
    case "foo bar":
        switch 378 {
        case 42:
            var v40 uint64 = 401
            _ = v40
        }
    }
    if true {
        var v41 float64 = 1.5
        if (v5 > v5) {
            fmt.Println(v41, v41)
            var v42 int16 = 1
            fmt.Println(v5)
            _ = v42
        }
        if false {
            if true {
                if ("\n\t" < "foo bar") {
                    if false {
                        fmt.Println(v5, v5, v41)
                        var v43 float32 = v5
                        _ = v43
                    } else {
                        if true {
                            fmt.Println(v3)
                        } else {
                            switch "a" {
                            case "a":
                                for fi44 := 0; fi44 < 5; fi44++ {
                                    _ = len(v4)
                                    _ = len(v6)
                                    v45 := GenS2{F0: 42, F1: 1}
                                    _ = fi44
                                    _ = v45
                                }
                            case "hello":
                                var v46 uint32 = 42
                                _ = v46
                            }
                            for fi47 := 0; fi47 < 12; fi47++ {
                                var v48 uint = 18446744073709551615
                                v5 += v5
                                _ = fi47
                                _ = v48
                            }
                            if false {
                                v49 := GenS2{F0: 127, F1: 9223372036854775807}
                                v50 := GenS5{F0: v3, F1: 578}
                                _ = v49
                                _ = v50
                            }
                        }
                        for fi51 := 0; fi51 < 7; fi51++ {
                            v52 := v6[fi51]
                            _ = fi51
                            _ = v52
                        }
                    }
                    if true {
                        v41 = 3.14159
                        if (-365 >= 495) {
                            v41 = v41
                        }
                        v5 = float32(3.14159)
                    }
                }
            } else {
                _ = len(v6)
            }
            if false {
                v53 := GenS4{F0: 642, F1: 42}
                _ = v53
            }
            v41 = 3.14159
        }
        _ = v41
    } else {
        var v54 int = 0
        var v55 string = "foo bar"
        _ = v54
        _ = v55
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}

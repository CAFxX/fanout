package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 int; F1 uint16; F2 int16 }
func genFn0(p0 bool) int {
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 uint16 = 42
        v7 -= (v7 & (v7 - v7))
        _ = fi6
        _ = v7
    }
    fmt.Println(p0)
    var v8 string = "a"
    _ = p0
    _ = v8
    return -9223372036854775808
}
func genFn1() string {
    var v9 uint = 2
    if true {
        if (!(!(42 == 127))) {
            v9 = (^(v9))
            switch 127 {
            case 0:
                v9 += (v9 - v9)
            case 21:
                for fi10 := 0; fi10 < 11; fi10++ {
                    v11 := []int{}
                    var v12 int8 = 100
                    var v13 uint32 = 2147483648
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
            case 200:
                switch (("foo bar" + "hello") + "") {
                case "hello":
                    for fi14 := 0; fi14 < 20; fi14++ {
                        v15 := GenS1{F0: -9223372036854775808}
                        _ = fi14
                        _ = v15
                    }
                case "a":
                    v9 = (v9 + v9)
                case "\n\t":
                    v9 = 665
                }
            default:
                fmt.Println(v9)
            }
        }
    }
    switch ("hello" + "hello") {
    case "\n\t":
        v16 := GenS1{F0: -295}
        _ = v16
    case "foo bar":
        v17 := []int{788, -1, 907, -396}
        _ = v17
    }
    var v18 uint16 = 32767
    switch 735 {
    case 182:
        v18 = 32767
    case -9223372036854775808:
        v19 := []int{9223372036854775807, 32768, 0}
        _ = v19
    default:
        if (("a" == (("hello" + "\n\t") + "hello")) && false) {
            v20 := []int{9223372036854775807, -2147483648, 32768}
            fmt.Println(v9, v18)
            for fi21 := 0; fi21 < 10; fi21++ {
                switch "hello" {
                case "foo bar":
                    var v22 uint64 = 4294967295
                    _ = v22
                }
                if false {
                    for fi23 := 0; fi23 < 5; fi23++ {
                        var v24 int8 = 17
                        _ = fi23
                        _ = v24
                    }
                } else {
                    var v25 int32 = 1
                    v9 += v9
                    v9 = (v9 ^ (v9 >> ((v9) & 63)))
                    _ = v25
                }
                _ = fi21
            }
            _ = v20
        } else {
            v18 = v18
            v26 := []int{-576, -613, 144}
            _ = v26
        }
    }
    _ = v18
    _ = v9
    return "foo bar"
}

func main() {
    var v1 string = (("foo bar" + "hello") + "")
    v2 := []string{}
    v3 := make(map[string]uint32)
    var v4 bool = (!(-715 > -9223372036854775808))
    var v5 string = ("\n\t" + "foo bar")
    v1 = v5
    if (-2147483648 > 42) {
        fmt.Println(v5, v4, v4)
    }
    switch 0 {
    case 686:
        for fi27 := 0; fi27 < 16; fi27++ {
            v5 = v1
            fmt.Println(v5, v5, v4)
            if (v4 && false) {
                fmt.Println(v5, fi27)
            } else {
                switch v5 {
                case "a":
                    for fi28 := 0; fi28 < 8; fi28++ {
                        var v29 string = v5
                        var v30 uint32 = 0
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                case "hello":
                    var v31 int64 = 9223372036854775807
                    _ = v31
                }
                v1 = v5
                var v32 int32 = 0
                _ = v32
            }
            _ = fi27
        }
    case -936:
        for fi33 := 0; fi33 < 19; fi33++ {
            var v34 int = (-(fi33))
            v2 = append(v2, v5)
            if (!v4) {
                if false {
                    if len(v2) > 0 {
                        v2[0] = "\n\t"
                    }
                } else {
                    v1 = ""
                }
                for fi35 := 0; fi35 < 8; fi35++ {
                    v4 = (!false)
                    _ = fi35
                }
            }
            _ = fi33
            _ = v34
        }
    case 42:
        _ = len(v2)
    default:
        v2 = append(v2, ("\n\t" + (v1 + v1)))
    }
    fmt.Println(v1, v1, v4)
    v36 := v3[""]
    for fi37 := 0; fi37 < 20; fi37++ {
        if (708 <= 255) {
            var v38 int16 = 42
            if (1 != 42) {
                for fi39 := 0; fi39 < 11; fi39++ {
                    switch v1 {
                    case "":
                        var v40 int64 = 42
                        _ = v40
                    case "hello":
                        switch v5 {
                        case "foo bar":
                            fmt.Println(v1)
                            var v41 uint32 = v36
                            v42 := GenS2{F0: 42}
                            _ = v41
                            _ = v42
                        case "\n\t":
                            if len(v2) > 0 {
                                v2[0] = (("\n\t" + v1) + "")
                            }
                        case "a":
                            var v43 int8 = 1
                            _ = v43
                        default:
                            v44 := GenS2{F0: 2}
                            _ = v44
                        }
                    }
                    _ = fi39
                }
                var v45 uint32 = v36
                v4 = (false && (!(-1 != 422)))
                _ = v45
            }
            _ = v38
        }
        _ = fi37
    }
    fmt.Println(v36)
    var v46 uint16 = 0
    var v47 uint16 = v46
    for fi48 := 0; fi48 < 10; fi48++ {
        v49 := GenS3{F0: 4294967295, F1: (v4 || v4)}
        fmt.Println(fi48, fi48)
        _ = fi48
        _ = v49
    }
    var v50 int64 = 42
    v46 = 2
    switch 127 {
    case -905:
        _ = len(v2)
    case 32768:
        if (1e10 > (float32((1e10 * ((0.1 + ((1e10 + 0.1) - (3.14159 / (3.14159 - (1e10 + 0.0))))) / ((float64(1.5) * ((1.5 * 1.5) * 0.1)) + (-2.25 * float64(1.5)))))) + (0.0 / 1e10))) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            v51 := GenS4{F0: -1, F1: 479, F2: -1}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
            _ = v51
        }
    case -2147483648:
        v4 = ((3.14159 >= 1e10) && (v4 && (0 <= (v50 & (v50 - (v50 ^ -1))))))
    default:
        var v52 int8 = -87
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v50
}

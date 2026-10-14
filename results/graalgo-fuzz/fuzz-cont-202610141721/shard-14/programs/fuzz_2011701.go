package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint64; F1 uint }
type GenS7 struct { F0 string; F1 uint64; F2 int }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint8; F1 string }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint8 }
func genFn0(p0 int16, p1 string, p2 uint64) bool {
    fmt.Println(p1, p0, p1)
    if false {
        p0 = (-(p0))
        var v6 int = -178
        v7 := GenS1{F0: (true && true), F1: -1}
        _ = v6
        _ = v7
    } else {
        fmt.Println(p2, p2, p0)
        if (193 != 2147483647) {
            if true {
                var v8 int64 = -9223372036854775808
                switch (p1 + ((("hello" + "\n\t") + ("a" + p1)) + "\n\t")) {
                case "a":
                    p0 = 32767
                case "":
                    if false {
                        fmt.Println(p0, v8)
                    }
                case "\n\t":
                    v9 := GenS2{F0: 2, F1: 0}
                    _ = v9
                }
                fmt.Println(p2)
                _ = v8
            }
        } else {
            if true {
                p1 = ("a" + p1)
            }
            var v10 bool = (!false)
            _ = v10
        }
        p1 = ((("a" + p1) + ((p1 + p1) + p1)) + p1)
    }
    _ = p0
    _ = p1
    _ = p2
    return false
}

func main() {
    v1 := []string{"\n\t"}
    v2 := make(map[uint32]string)
    var v3 uint64 = 192
    var v4 float64 = float64(3.14159)
    v5 := make(map[uint32]bool)
    fmt.Println(v4, v3)
    switch "foo bar" {
    case "a":
        v1 = append(v1, "")
    case "hello":
        for fi11 := 0; fi11 < 18; fi11++ {
            v3 = v3
            v12 := GenS3{F0: -839}
            _ = fi11
            _ = v12
        }
    case "\n\t":
        v13 := GenS4{F0: ("" + "hello")}
        _ = v13
    }
    v14 := GenS1{F0: ((-3 == 9223372036854775807) || (800 < 9223372036854775807)), F1: 100}
    if (true || false) {
        var v15 int32 = -6
        _ = v15
    }
    var v16 bool = true
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    if v16 {
        if false {
            v17 := GenS4{F0: "\n\t"}
            v16 = (v16 || (!(32767 != 578)))
            v18 := GenS3{F0: 255}
            _ = v17
            _ = v18
        } else {
            for fi19 := 0; fi19 < 5; fi19++ {
                v20 := GenS5{F0: 470}
                var v21 uint = 42
                var v22 int = 32768
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
            switch "" {
            case "a":
                if false {
                    switch ("a" + "foo bar") {
                    case "hello":
                        v23 := GenS6{F0: v3, F1: 0}
                        _ = v23
                    case "\n\t":
                        for fi24 := 0; fi24 < 17; fi24++ {
                            var v25 uint8 = 255
                            _ = fi24
                            _ = v25
                        }
                    }
                    switch (("" + ("" + (("" + (("a" + (("hello" + ("\n\t" + "\n\t")) + (("" + (("hello" + "") + "hello")) + "a"))) + "a")) + "a"))) + ("foo bar" + "hello")) {
                    case "foo bar":
                        var v26 int32 = -2147483648
                        _ = v26
                    case "\n\t":
                        for fi27 := 0; fi27 < 13; fi27++ {
                            switch (("foo bar" + "a") + "\n\t") {
                            case "":
                                var v28 uint64 = (^(v3))
                                _ = v28
                            default:
                                for fi29 := 0; fi29 < 18; fi29++ {
                                    switch "a" {
                                    case "foo bar":
                                        v3 -= (^(v3))
                                    case "a":
                                        var v30 int32 = -189
                                        _ = v30
                                    }
                                    v4 -= (v4 * float64((1.5 / (float32((v4 * v4)) + 1e10))))
                                    _ = fi29
                                }
                            }
                            v31 := GenS7{F0: "\n\t", F1: v3, F2: (fi27 + (-(fi27)))}
                            _ = fi27
                            _ = v31
                        }
                    default:
                        fmt.Println(v4, v16)
                    }
                }
            case "hello":
                v3 += v3
            default:
                var v32 float32 = (((3.14159 - float32(v4)) / (((0.1 + 1.5) + 0.0) * ((0.0 / (float32(float64(0.1)) / float32((-2.25 * v4)))) * ((0.1 + float32(v4)) * (((1.5 + (((float32(v4) / 1.0) * 0.1) + (((-2.25 / 1.0) - 3.14159) + 3.14159))) / float32(v4)) / 3.14159))))) - 0.1)
                _ = v32
            }
        }
        v5[65535] = v16
        var v33 float64 = v4
        _ = v33
    } else {
        fmt.Println(v4, v16, v3)
        v34 := GenS8{F0: 127}
        _ = v34
    }
    var v35 int64 = -630
    if false {
        if false {
            fmt.Println(v4)
            if (1 == 128) {
                switch "foo bar" {
                case "":
                    fmt.Println(v3, v4)
                case "foo bar":
                    fmt.Println(v4)
                }
                v36 := GenS9{F0: 14, F1: "foo bar"}
                fmt.Println(v16, v4, v3)
                _ = v36
            } else {
                fmt.Println(v16)
                if len(v1) > 0 {
                    v1[0] = ((("hello" + "\n\t") + (("hello" + ("" + "hello")) + "hello")) + (("a" + "hello") + ((((((("" + ("hello" + "a")) + "\n\t") + ((("" + "hello") + "\n\t") + "foo bar")) + "a") + "a") + "\n\t") + "foo bar")))
                }
                var v37 float64 = float64((0.0 / (1e10 - 3.14159)))
                _ = v37
            }
            v2[985] = "\n\t"
        } else {
            fmt.Println(v3, v16)
        }
        v38 := GenS1{F0: ((v35 & -9223372036854775808) < v35), F1: 0}
        _ = v38
    } else {
        for fi39 := 0; fi39 < 20; fi39++ {
            var v40 uint32 = 2
            v41 := GenS10{F0: 4294967295}
            _ = fi39
            _ = v40
            _ = v41
        }
        v42 := v2[2147483648]
        _ = v42
    }
    v43 := GenS11{F0: 0}
    v3 = 42
    var v44 uint32 = 2147483647
    for fi45 := 0; fi45 < 2; fi45++ {
        var v46 int32 = 2147483647
        var v47 int16 = int16(v3)
        _ = fi45
        _ = v46
        _ = v47
    }
    if v16 {
        if v16 {
            v5[4294967295] = (((!(!((v44 == 0) || (100 >= -104)))) && (!v16)) && (v4 == float64(1e10)))
            var v48 uint32 = (^(v44))
            _ = v48
        }
    } else {
        fmt.Println(v16, v3)
        v49 := GenS12{F0: 255}
        v50 := GenS5{F0: 0}
        _ = v49
        _ = v50
    }
    switch "\n\t" {
    case "a":
        v3 = 291
    case "foo bar":
        for fi51 := 0; fi51 < 10; fi51++ {
            delete(v2, 4294967295)
            v4 -= float64(1.5)
            _ = fi51
        }
    case "\n\t":
        if v16 {
            var v52 uint = 255
            if v16 {
                for fi53 := 0; fi53 < 17; fi53++ {
                    if (!(v52 <= (v52 ^ (v52 * v52)))) {
                        if (!v16) {
                            fmt.Println(v16)
                        }
                        if len(v1) > 0 {
                            v1[0] = "a"
                        }
                        var v54 bool = v16
                        _ = v54
                    } else {
                        v52 = (v52 >> ((v52) & 63))
                    }
                    _ = fi53
                }
            } else {
                if v16 {
                    fmt.Println(v44, v4)
                    fmt.Println(v3, v52)
                }
                v35 = (v35 + v35)
            }
            _ = v52
        }
    default:
        v55 := GenS12{F0: 255}
        _ = v55
    }
    if v16 {
        var v56 uint16 = 65535
        fmt.Println(v44, v16, v16)
        v56 -= v56
        _ = v56
    }
    var v57 int16 = 1000
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v57:", v57)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v57
}

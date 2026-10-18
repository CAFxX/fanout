package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int32; F1 uint16; F2 bool }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool; F1 string; F2 string }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint16 }
func genFn0(p0 string, p1 int32) uint64 {
    p1 -= (^(p1))
    var v5 int32 = p1
    _ = p0
    _ = p1
    _ = v5
    return 42
}
func genFn1(p0 uint8) uint8 {
    p0 = 183
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 16; fi7++ {
            v8 := []int{fi6, 127, (fi6 >> ((367) & 63))}
            var v9 uint = 340
            v10 := GenS1{F0: 42, F1: -2147483648}
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    v11 := []int{42, 32768, -9223372036854775808}
    if true {
        fmt.Println(p0)
        if true {
            switch ("foo bar" + "a") {
            case "hello":
                if true {
                    _ = len(v11)
                } else {
                    v12 := GenS2{F0: -32768, F1: 255}
                    _ = v12
                }
            case "\n\t":
                var v13 float32 = 0.0
                _ = v13
            }
            fmt.Println(p0)
            if len(v11) > 0 {
                v11[0] = -1000
            }
        }
        p0 += 127
    }
    if false {
        if (false || true) {
            v14 := GenS1{F0: 32767, F1: 0}
            _ = v14
        } else {
            p0 = p0
            if false {
                switch "foo bar" {
                case "foo bar":
                    var v15 bool = false
                    _ = v15
                case "a":
                    v16 := GenS3{F0: 805, F1: 576, F2: false}
                    _ = v16
                case "\n\t":
                    switch ((((("foo bar" + (("\n\t" + ("" + ("hello" + "\n\t"))) + ("" + "\n\t"))) + "a") + "foo bar") + ("foo bar" + "a")) + (("" + ("a" + "\n\t")) + ("a" + "hello"))) {
                    case "\n\t":
                        var v17 int = 32768
                        _ = v17
                    case "":
                        switch 0 {
                        case 42:
                            _ = len(v11)
                        }
                    default:
                        if (!true) {
                            for fi18 := 0; fi18 < 12; fi18++ {
                                var v19 string = ("\n\t" + "foo bar")
                                _ = fi18
                                _ = v19
                            }
                            _ = len(v11)
                        } else {
                            if true {
                                for fi20 := 0; fi20 < 9; fi20++ {
                                    var v21 uint = 0
                                    switch (fi20 - 32768) {
                                    case -9223372036854775808:
                                        v22 := GenS2{F0: 1, F1: 905}
                                        _ = v22
                                    }
                                    _ = fi20
                                    _ = v21
                                }
                                switch 127 {
                                case 32767:
                                    p0 += p0
                                case 1:
                                    var v23 int64 = -382
                                    _ = v23
                                case -9223372036854775808:
                                    fmt.Println(p0)
                                }
                            }
                        }
                    }
                default:
                    v24 := GenS4{F0: 342, F1: -9223372036854775808}
                    _ = v24
                }
                var v25 uint64 = 787
                for fi26 := 0; fi26 < 6; fi26++ {
                    var v27 uint32 = 2147483648
                    for fi28 := 0; fi28 < 5; fi28++ {
                        _ = len(v11)
                        v29 := GenS4{F0: 255, F1: (fi26 % ((fi26) | 1))}
                        var v30 uint8 = 46
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    var v31 bool = false
                    _ = fi26
                    _ = v27
                    _ = v31
                }
                _ = v25
            } else {
                var v32 string = ("\n\t" + "hello")
                _ = v32
            }
            v33 := GenS2{F0: -32768, F1: 0}
            _ = v33
        }
        for fi34 := 0; fi34 < 19; fi34++ {
            v35 := GenS5{F0: (!false)}
            var v36 uint8 = 127
            _ = fi34
            _ = v35
            _ = v36
        }
    } else {
        if (2 != 4294967295) {
            p0 = (^(p0))
        } else {
            v37 := GenS4{F0: 32767, F1: 0}
            v38 := GenS5{F0: false}
            _ = v37
            _ = v38
        }
        for fi39 := 0; fi39 < 17; fi39++ {
            for fi40 := 0; fi40 < 8; fi40++ {
                var v41 uint8 = 42
                _ = fi40
                _ = v41
            }
            for fi42 := 0; fi42 < 14; fi42++ {
                fmt.Println(fi39, fi42, fi42)
                _ = fi42
            }
            _ = fi39
        }
        p0 = (p0 + (^(p0)))
    }
    _ = p0
    _ = v11
    return (p0 + (^(p0)))
}

func main() {
    var v1 bool = true
    var v2 int8 = -128
    var v3 string = "\n\t"
    v4 := []string{(("a" + (v3 + "hello")) + ""), "foo bar", v3}
    switch (v3 + v3) {
    case "hello":
        v43 := GenS6{F0: true, F1: v3, F2: "a"}
        _ = v43
    }
    switch (((v3 + (v3 + ("" + v3))) + (("hello" + v3) + (v3 + v3))) + ((v3 + "\n\t") + (v3 + ("hello" + v3)))) {
    case "foo bar":
        v4 = append(v4, "a")
    case "a":
        for fi44 := 0; fi44 < 6; fi44++ {
            _ = len(v4)
            _ = fi44
        }
    }
    var v45 string = "a"
    v46 := GenS7{F0: 2}
    v47 := GenS8{F0: 18446744073709551615}
    for fi48 := 0; fi48 < 15; fi48++ {
        var v49 float64 = 1.5
        v1 = (v45 >= v3)
        v50 := GenS7{F0: 4294967295}
        _ = fi48
        _ = v49
        _ = v50
    }
    v51 := GenS2{F0: 0, F1: 65535}
    switch 0 {
    case -1:
        fmt.Println(v2, v2, v2)
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v52 := GenS2{F0: 32767, F1: 2}
    fmt.Println(v3, v1)
    switch -248 {
    case 128:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case -9223372036854775808:
        if true {
            var v53 uint16 = 1
            var v54 float32 = 0.0
            _ = v53
            _ = v54
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v3 = ""
        }
    default:
        for fi55 := 0; fi55 < 5; fi55++ {
            v56 := GenS9{F0: 42}
            v2 -= -128
            switch ("" + ("\n\t" + ((((((((v45 + v3) + v45) + "a") + "") + v3) + v45) + (v3 + (v45 + v3))) + v45))) {
            case "foo bar":
                _ = len(v4)
            case "\n\t":
                v4 = append(v4, v3)
            }
            _ = fi55
            _ = v56
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v51
    _ = v52
}

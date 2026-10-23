package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 string; F1 uint64; F2 int8 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int8; F1 uint }
type GenS8 struct { F0 int32 }
func genFn0(p0 bool, p1 int, p2 uint64) bool {
    p1 = p1
    var v4 float32 = 0.1
    switch p1 {
    case -9223372036854775808:
        fmt.Println(p0, p1, p1)
    }
    for fi5 := 0; fi5 < 15; fi5++ {
        v6 := GenS1{F0: fi5, F1: -677}
        _ = fi5
        _ = v6
    }
    v7 := []int{p1, p1, 9223372036854775807, -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    return (795 == 1)
}
func genFn1(p0 int32, p1 int32) string {
    if false {
        v8 := GenS1{F0: 32768, F1: p0}
        var v9 int8 = -1
        _ = v8
        _ = v9
    } else {
        switch 9223372036854775807 {
        case 32768:
            var v10 string = ("a" + "foo bar")
            _ = v10
        case 915:
            fmt.Println(p0)
        case -321:
            for fi11 := 0; fi11 < 17; fi11++ {
                switch "a" {
                case "\n\t":
                    var v12 int = fi11
                    _ = v12
                case "foo bar":
                    fmt.Println(p0, p1, p0)
                case "":
                    p0 -= p0
                default:
                    v13 := GenS2{F0: -9223372036854775808, F1: 0}
                    _ = v13
                }
                var v14 int = 32767
                _ = fi11
                _ = v14
            }
        }
        fmt.Println(p0, p0)
    }
    p0 += p0
    _ = p0
    _ = p1
    return (("" + "foo bar") + "foo bar")
}

func main() {
    var v1 int64 = -357
    var v2 int8 = int8(v1)
    v3 := []uint32{42, 1, 285}
    if true {
        v1 = v1
        v1 = 0
    } else {
        v2 += 0
        var v15 float64 = 0.1
        _ = v15
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        fmt.Println(v1)
        _ = fi16
    }
    var v17 uint16 = 255
    v17 -= v17
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    switch "a" {
    case "\n\t":
        if (0 < 0) {
            var v18 int = 127
            _ = v18
        } else {
            var v19 float64 = ((((1.5 - 3.14159) * 1e10) - float64(-2.25)) + (1.5 + -2.25))
            _ = v19
        }
    default:
        switch "foo bar" {
        case "hello":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "foo bar":
            fmt.Println(v1, v1, v1)
        default:
            v2 = (v2 >> ((980) & 63))
        }
    }
    switch ("foo bar" + "a") {
    case "\n\t":
        if (788 <= 47) {
            switch ("hello" + ((("" + "hello") + ("hello" + "hello")) + (("a" + "a") + ""))) {
            case "\n\t":
                if (!(v1 > (v1 << ((0) & 63)))) {
                    v2 = (v2 * -1)
                    for fi20 := 0; fi20 < 20; fi20++ {
                        for fi21 := 0; fi21 < 1; fi21++ {
                            v22 := GenS3{F0: "a", F1: 408, F2: v2}
                            var v23 uint64 = 18446744073709551615
                            v24 := GenS3{F0: ("" + "hello"), F1: 18446744073709551615, F2: (v2 / ((v2) | 1))}
                            var v25 uint16 = v17
                            fmt.Println(v1, v17, v17)
                            _ = fi21
                            _ = v22
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                        _ = fi20
                    }
                } else {
                    var v26 bool = false
                    v1 = (v1 | v1)
                    _ = v26
                }
            case "":
                fmt.Println(v17)
            }
        } else {
            var v27 uint64 = 730
            v28 := GenS4{F0: (("foo bar" + "foo bar") + "a"), F1: 0}
            _ = v27
            _ = v28
        }
    case "":
        if true {
            fmt.Println(v17, v2, v17)
            fmt.Println(v1, v2)
            v29 := GenS5{F0: 4294967295, F1: 766, F2: 93}
            _ = v29
        } else {
            var v30 float64 = (((0.1 - 0.1) - 0.0) / ((-2.25 + float64((1.5 - (1e10 / 3.14159)))) * 1e10))
            v17 = v17
            _ = v30
        }
    case "hello":
        if (-676 != -1) {
            fmt.Println(v1, v2, v17)
            for fi31 := 0; fi31 < 10; fi31++ {
                var v32 int8 = (v2 * v2)
                if len(v3) > 0 {
                    v3[0] = 1
                }
                _ = fi31
                _ = v32
            }
        } else {
            if len(v3) > 0 {
                v3[0] = 42
            }
        }
    default:
        v33 := GenS6{F0: ""}
        _ = v33
    }
    v34 := GenS7{F0: (v2 & v2), F1: 2}
    for fi35 := 0; fi35 < 12; fi35++ {
        switch "hello" {
        case "hello":
            switch ((("" + (("hello" + ((("foo bar" + "hello") + ("hello" + "hello")) + ("\n\t" + ((((("" + (("a" + "\n\t") + "a")) + "") + "hello") + (("\n\t" + ("a" + "a")) + "foo bar")) + "")))) + ("\n\t" + ""))) + ("a" + ("a" + "a"))) + "") {
            case "hello":
                var v36 uint64 = 671
                _ = v36
            }
        case "\n\t":
            v2 -= (v2 + v2)
        }
        _ = fi35
    }
    for fi37 := 0; fi37 < 19; fi37++ {
        fmt.Println(v17)
        _ = fi37
    }
    switch (("" + ("\n\t" + ("" + "a"))) + "") {
    case "":
        v38 := GenS4{F0: "a", F1: 127}
        _ = v38
    case "foo bar":
        for fi39 := 0; fi39 < 12; fi39++ {
            v40 := GenS6{F0: (("foo bar" + (("hello" + "foo bar") + "foo bar")) + "foo bar")}
            var v41 float32 = 3.14159
            fmt.Println(v2)
            _ = fi39
            _ = v40
            _ = v41
        }
    default:
        if true {
            if len(v3) > 0 {
                v3[0] = 778
            }
            if ((!true) || (42 == 1)) {
                if (!(1e10 <= 0.0)) {
                    var v42 float64 = ((0.1 / -2.25) * (-2.25 - -2.25))
                    var v43 uint = 4294967295
                    fmt.Println(v2, v42, v17)
                    _ = v42
                    _ = v43
                }
                switch "\n\t" {
                case "":
                    var v44 int32 = 0
                    _ = v44
                case "foo bar":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "a":
                    if true {
                        var v45 float32 = ((0.0 / 1.5) / 3.14159)
                        _ = v45
                    }
                default:
                    v2 = v2
                }
                v17 += (^(v17))
            }
        }
    }
    for fi46 := 0; fi46 < 5; fi46++ {
        if len(v3) > 0 {
            v3[0] = 2147483648
        }
        _ = fi46
    }
    var v47 int = -385
    if (!false) {
        for fi48 := 0; fi48 < 20; fi48++ {
            v49 := GenS8{F0: 1}
            v50 := GenS8{F0: 837}
            _ = fi48
            _ = v49
            _ = v50
        }
        var v51 int32 = 872
        var v52 int8 = (v2 | (v2 & v2))
        _ = v51
        _ = v52
    }
    v1 += (v1 % (((v1 + (^(v1)))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v47
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 uint16; F2 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 int; F2 uint32 }
type GenS6 struct { F0 int16; F1 int64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 string }
func genFn0() string {
    var v5 int = 32768
    v5 += 128
    v5 = (v5 | v5)
    switch v5 {
    case 4294967295:
        switch "a" {
        case "\n\t":
            v5 = v5
        case "":
            var v6 float64 = 3.14159
            _ = v6
        case "a":
            fmt.Println(v5)
        }
    case 32768:
        switch "hello" {
        case "\n\t":
            switch ("foo bar" + "\n\t") {
            case "hello":
                if true {
                    v5 = v5
                    for fi7 := 0; fi7 < 4; fi7++ {
                        var v8 float64 = 0.1
                        _ = fi7
                        _ = v8
                    }
                    v9 := GenS1{F0: true, F1: 0}
                    _ = v9
                }
            }
        case "a":
            v10 := []int{}
            _ = v10
        default:
            for fi11 := 0; fi11 < 16; fi11++ {
                v12 := GenS2{F0: 110}
                _ = fi11
                _ = v12
            }
        }
    case -9223372036854775808:
        var v13 int16 = 1000
        _ = v13
    default:
        v14 := GenS3{F0: 1, F1: 1, F2: (v5 >> ((4294967295) & 63))}
        _ = v14
    }
    _ = v5
    return ("" + "")
}
func genFn1(p0 int16) int {
    fmt.Println(p0)
    switch "a" {
    case "foo bar":
        v15 := GenS4{F0: (1.5 == 1.5)}
        _ = v15
    case "\n\t":
        if true {
            switch "a" {
            case "\n\t":
                fmt.Println(p0)
            }
        }
    case "":
        switch 0 {
        case 32767:
            var v16 int64 = 42
            _ = v16
        case -376:
            v17 := GenS3{F0: 2, F1: 521, F2: 42}
            _ = v17
        case 390:
            v18 := GenS5{F0: 2, F1: 1, F2: 2}
            _ = v18
        }
    default:
        switch -138 {
        case 127:
            if (!(-1 >= 2147483647)) {
                var v19 int = -29
                for fi20 := 0; fi20 < 14; fi20++ {
                    var v21 string = "hello"
                    v19 -= (v19 ^ (v19 + -2147483648))
                    _ = fi20
                    _ = v21
                }
                v22 := GenS5{F0: 4294967295, F1: 9223372036854775807, F2: 2147483648}
                _ = v19
                _ = v22
            } else {
                var v23 uint64 = 18446744073709551615
                var v24 int8 = int8(v23)
                v25 := GenS6{F0: int16(v24), F1: 9223372036854775807}
                _ = v23
                _ = v24
                _ = v25
            }
        case 32768:
            p0 -= (p0 * (p0 * (^(p0))))
        }
    }
    v26 := GenS7{F0: 127}
    p0 = (p0 + (^(p0)))
    _ = p0
    _ = v26
    return -2147483648
}

func main() {
    var v1 int16 = -581
    v2 := make(map[uint32]string)
    v3 := []string{"", "a"}
    v4 := make(map[string]bool)
    var v27 float32 = -2.25
    if (!true) {
        var v28 int32 = 1
        _ = v28
    }
    switch 65 {
    case 127:
        fmt.Println(v1)
    case 4294967295:
        var v29 int32 = -876
        _ = v29
    }
    var v30 uint64 = 1
    for fi31 := 0; fi31 < 1; fi31++ {
        var v32 int = fi31
        v33 := GenS7{F0: 0}
        _ = fi31
        _ = v32
        _ = v33
    }
    v34 := GenS5{F0: 2147483647, F1: 1, F2: 494}
    v35 := GenS8{F0: 42}
    var v36 uint32 = 2147483647
    switch "\n\t" {
    case "hello":
        v3 = append(v3, ("hello" + "a"))
    case "":
        fmt.Println(v1)
    case "a":
        fmt.Println(v36)
    default:
        v37 := GenS6{F0: (v1 << ((1) & 63)), F1: 883}
        _ = v37
    }
    v30 += (v30 / ((2) | 1))
    var v38 int16 = 32767
    if false {
        if (255 < 42) {
            v38 -= (v38 ^ v1)
            switch 127 {
            case -854:
                v1 -= (v1 ^ (v1 / ((v38) | 1)))
            case -69:
                for fi39 := 0; fi39 < 2; fi39++ {
                    var v40 bool = (!(true && true))
                    v41 := GenS5{F0: v36, F1: fi39, F2: v36}
                    _ = fi39
                    _ = v40
                    _ = v41
                }
            }
        } else {
            for fi42 := 0; fi42 < 2; fi42++ {
                v43 := GenS9{F0: -128}
                _ = fi42
                _ = v43
            }
        }
        v3 = append(v3, "a")
        fmt.Println(v30, v1)
    }
    for fi44 := 0; fi44 < 5; fi44++ {
        if (false && (true || false)) {
            v45 := GenS3{F0: v30, F1: 2, F2: fi44}
            if len(v3) > 0 {
                v3[0] = ("a" + "foo bar")
            }
            _ = v45
        }
        _ = fi44
    }
    fmt.Println(v1, v1, v30)
    switch -2147483648 {
    case -551:
        for fi46 := 0; fi46 < 19; fi46++ {
            v30 = (^(v30))
            fmt.Println(v36, v27, v30)
            _ = fi46
        }
    case 2147483647:
        v30 -= v30
    }
    switch 4294967295 {
    case 9223372036854775807:
        if len(v3) > 0 {
            v3[0] = ("\n\t" + (("foo bar" + "") + ("foo bar" + "a")))
        }
    case 32768:
        for fi47 := 0; fi47 < 1; fi47++ {
            v48 := GenS1{F0: ((((3.14159 / float64((0.0 / v27))) / 1e10) * float64(float32(1.5))) == ((-2.25 * 0.1) * (0.1 * 1e10))), F1: (v36 - (v36 << ((uint(v36)) & 63)))}
            _ = fi47
            _ = v48
        }
    case 97:
        fmt.Println(v27)
    }
    switch "foo bar" {
    case "a":
        var v49 bool = true
        _ = v49
    case "\n\t":
        switch -365 {
        case 127:
            if (255 >= 1) {
                for fi50 := 0; fi50 < 15; fi50++ {
                    var v51 int16 = v1
                    _ = fi50
                    _ = v51
                }
            }
        default:
            var v52 string = (("a" + "hello") + "\n\t")
            _ = v52
        }
    }
    v53 := GenS10{F0: "foo bar"}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v36
    _ = v38
    _ = v4
    _ = v53
}

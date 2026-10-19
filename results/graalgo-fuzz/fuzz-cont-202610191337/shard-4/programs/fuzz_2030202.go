package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 bool }
type GenS2 struct { F0 bool; F1 int16; F2 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 int16 }

func main() {
    var v1 uint32 = 42
    var v2 bool = false
    var v3 uint32 = 2147483648
    v4 := []string{}
    v5 := make(map[string]bool)
    var v6 int = -300
    v6 = 4294967295
    v7 := GenS1{F0: v6, F1: -1, F2: v2}
    for fi8 := 0; fi8 < 16; fi8++ {
        fmt.Println(fi8, v2, fi8)
        _ = fi8
    }
    delete(v5, ("" + "\n\t"))
    v9 := GenS1{F0: v6, F1: 1, F2: (-807 != int16(v6))}
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 int16 = -32768
        fmt.Println(v2, v2, v6)
        _ = fi10
        _ = v11
    }
    v1 = v3
    for fi12 := 0; fi12 < 12; fi12++ {
        v13 := GenS2{F0: (!(-587 == 42)), F1: -1, F2: -913}
        var v14 int8 = 119
        switch (fi12 ^ 42) {
        case 1:
            var v15 int = -1
            _ = v15
        case 845:
            v16 := GenS3{F0: 2}
            _ = v16
        default:
            var v17 float64 = -2.25
            _ = v17
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    var v18 int32 = 42
    if (v2 || (!v2)) {
        for fi19 := 0; fi19 < 11; fi19++ {
            for fi20 := 0; fi20 < 12; fi20++ {
                v1 = (v1 - 2147483648)
                v21 := GenS3{F0: 0}
                if ((v18 + (v18 % ((v18) | 1))) > (-(v18))) {
                    v22 := GenS1{F0: fi19, F1: v18, F2: (42 <= -9223372036854775808)}
                    _ = v22
                }
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
        v6 += v6
        if ((float32(float64(float32(1e10))) / -2.25) > (float32(0.1) / 3.14159)) {
            var v23 int64 = -9223372036854775808
            fmt.Println(v23, v6)
            if (!("foo bar" == "")) {
                v18 = v18
                fmt.Println(v18)
            }
            _ = v23
        } else {
            switch ("\n\t" + ("hello" + "hello")) {
            case "foo bar":
                var v24 uint64 = 307
                _ = v24
            case "hello":
                if v2 {
                    v1 -= v1
                    var v25 int64 = 42
                    if (v25 == v25) {
                        var v26 int8 = -1
                        _ = len(v5)
                        switch int(v3) {
                        case 9223372036854775807:
                            fmt.Println(v18, v1, v25)
                        }
                        _ = v26
                    } else {
                        var v27 float64 = ((-2.25 * (-2.25 * 0.1)) * 1e10)
                        _ = v27
                    }
                    _ = v25
                }
            default:
                var v28 int8 = 1
                _ = v28
            }
            v3 -= 112
        }
    } else {
        v6 += (v6 % ((v6) | 1))
    }
    if true {
        if (0.0 < (-2.25 + -2.25)) {
            var v29 int16 = 585
            v30 := GenS4{F0: -653, F1: 652}
            _ = v29
            _ = v30
        }
    } else {
        switch 1 {
        case 9223372036854775807:
            v4 = append(v4, "\n\t")
        case 4294967295:
            v31 := GenS3{F0: 781}
            _ = v31
        case -820:
            switch 32767 {
            case 32768:
                if (v2 || v2) {
                    if len(v4) > 0 {
                        v4[0] = ""
                    }
                }
            case -824:
                fmt.Println(v6)
            case 9223372036854775807:
                delete(v5, ("a" + ("\n\t" + "a")))
            }
        }
        switch ("hello" + "a") {
        case "hello":
            fmt.Println(v6, v1, v2)
        default:
            for fi32 := 0; fi32 < 9; fi32++ {
                switch "\n\t" {
                case "\n\t":
                    var v33 uint16 = 255
                    _ = v33
                case "foo bar":
                    fmt.Println(v18, v2)
                case "hello":
                    if true {
                        _ = len(v4)
                        for fi34 := 0; fi34 < 3; fi34++ {
                            var v35 float32 = 3.14159
                            _ = fi34
                            _ = v35
                        }
                        v2 = v2
                    }
                default:
                    fmt.Println(v3, v2, v3)
                }
                _ = fi32
            }
        }
        for fi36 := 0; fi36 < 15; fi36++ {
            var v37 int16 = -1
            _ = fi36
            _ = v37
        }
    }
    v2 = v2
    v3 = (v1 | (v1 + v1))
    var v38 float64 = -2.25
    v39 := GenS2{F0: true, F1: -1, F2: 1}
    v3 = 2147483648
    switch (v6 / ((1) | 1)) {
    case 391:
        var v40 float32 = (1.5 + (((3.14159 * 1e10) * 0.0) * (3.14159 / (0.1 / -2.25))))
        _ = v40
    default:
        var v41 bool = v2
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}

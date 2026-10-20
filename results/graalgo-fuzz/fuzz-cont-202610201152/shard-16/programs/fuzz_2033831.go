package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"foo bar", "foo bar", "hello"}
    v3 := []uint32{65535}
    var v4 int = 0
    v5 := make(map[int]bool)
    var v6 int32 = 2147483647
    switch (v4 >> ((731) & 63)) {
    case -226:
        v6 = 42
    }
    delete(v1, "hello")
    v4 = -488
    var v7 uint = 65535
    var v8 uint32 = 795
    v8 += v8
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 bool = true
        for fi11 := 0; fi11 < 16; fi11++ {
            switch "a" {
            case "foo bar":
                _ = len(v2)
            case "\n\t":
                if (v6 > v6) {
                    var v12 uint = v7
                    v8 = (v8 ^ v8)
                    fmt.Println(v7, v7, v10)
                    _ = v12
                }
            case "":
                v7 += (v7 * v7)
            }
            _ = fi11
        }
        v10 = v10
        _ = fi9
        _ = v10
    }
    var v13 int16 = 0
    fmt.Println(v7)
    switch (("\n\t" + ("hello" + ("hello" + ((("a" + ("" + "")) + "\n\t") + "a")))) + "hello") {
    case "":
        _ = len(v5)
    case "hello":
        v4 += (v4 % ((-319) | 1))
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        v15 := GenS1{F0: (v13 & v13), F1: int8(v7), F2: 863}
        if true {
            v16 := GenS1{F0: v13, F1: 100, F2: (v6 % ((-1) | 1))}
            v4 = (v4 ^ -9223372036854775808)
            var v17 float64 = ((((-2.25 + (-2.25 / 3.14159)) + 0.1) - -2.25) - 0.0)
            _ = v16
            _ = v17
        } else {
            fmt.Println(v7, v4, v7)
            switch ("" + ("hello" + ("a" + ("hello" + ((((("foo bar" + "a") + "\n\t") + "\n\t") + "") + "foo bar"))))) {
            case "\n\t":
                if (true && true) {
                    var v18 int8 = -120
                    v8 = (v8 + v8)
                    _ = v18
                }
            case "a":
                switch ("\n\t" + "a") {
                case "a":
                    switch "a" {
                    case "hello":
                        var v19 uint16 = 42
                        _ = v19
                    case "\n\t":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    default:
                        var v20 float64 = (0.0 * (-2.25 + 0.0))
                        _ = v20
                    }
                case "":
                    v7 -= 947
                case "hello":
                    if true {
                        v3 = append(v3, v8)
                    } else {
                        v4 += (v4 ^ v4)
                        v21 := GenS1{F0: v13, F1: 90, F2: v6}
                        v22 := GenS1{F0: v13, F1: -61, F2: v6}
                        _ = v21
                        _ = v22
                    }
                }
            }
        }
        _ = fi14
        _ = v15
    }
    v2 = append(v2, "a")
    switch ("\n\t" + "") {
    case "foo bar":
        switch (v4 << ((18446744073709551615) & 63)) {
        case 128:
            if (42 < 322) {
                fmt.Println(v13, v4)
                fmt.Println(v7, v13)
            }
        case 42:
            switch ("hello" + (("" + "") + "\n\t")) {
            case "foo bar":
                v7 -= 65535
            case "\n\t":
                if (65535 >= 1) {
                    var v23 float64 = (1e10 * 0.1)
                    _ = v23
                } else {
                    var v24 int64 = -363
                    var v25 int16 = v13
                    var v26 int = (v4 >> ((v7) & 63))
                    _ = v24
                    _ = v25
                    _ = v26
                }
            default:
                var v27 float64 = 3.14159
                _ = v27
            }
        case 9223372036854775807:
            var v28 int64 = 0
            _ = v28
        default:
            var v29 int8 = 42
            _ = v29
        }
    case "hello":
        var v30 int = v4
        _ = v30
    case "\n\t":
        switch 365 {
        case 399:
            var v31 uint32 = 158
            _ = v31
        case 315:
            switch "hello" {
            case "\n\t":
                _ = len(v3)
            case "a":
                v13 = v13
            case "foo bar":
                var v32 bool = false
                _ = v32
            default:
                if (true || (false && ((true && true) && (uint16(v7) != 0)))) {
                    var v33 float64 = 0.0
                    v8 += v8
                    _ = v33
                }
            }
        case -534:
            var v34 float64 = 0.0
            _ = v34
        }
    default:
        var v35 uint64 = 2
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

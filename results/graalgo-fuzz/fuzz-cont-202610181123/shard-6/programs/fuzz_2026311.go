package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 int32; F1 int64 }

func main() {
    v1 := []uint32{57, 65535, 65535}
    var v2 uint = 4294967295
    var v3 float32 = 0.0
    var v4 string = "\n\t"
    fmt.Println(v4, v4, v2)
    switch ("" + ("foo bar" + (v4 + v4))) {
    case "a":
        v4 = v4
    case "":
        var v5 string = "\n\t"
        _ = v5
    case "\n\t":
        _ = len(v1)
    default:
        for fi6 := 0; fi6 < 15; fi6++ {
            for fi7 := 0; fi7 < 5; fi7++ {
                fmt.Println(fi6, v3, v3)
                _ = fi7
            }
            for fi8 := 0; fi8 < 8; fi8++ {
                v9 := GenS1{F0: (0 > 1), F1: "foo bar", F2: (^(v2))}
                switch (v4 + "hello") {
                case "hello":
                    var v10 uint8 = 42
                    _ = v10
                case "":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "a":
                    v3 = ((((-2.25 / (v3 + (0.0 * ((1.5 * (float32(((-2.25 / 3.14159) - (((3.14159 / 1.5) + (1.5 - (1e10 + 1.5))) + -2.25))) - v3)) * (1e10 - 0.1))))) * ((v3 + v3) + v3)) / ((v3 / (v3 * v3)) + (v3 * (0.1 + ((v3 * (v3 / (-2.25 - v3))) - (v3 - (float32((0.1 / (0.0 - -2.25))) + v3))))))) - 1e10)
                }
                var v11 uint64 = 4294967295
                _ = fi8
                _ = v11
                _ = v9
            }
            for fi12 := 0; fi12 < 13; fi12++ {
                var v13 float32 = ((v3 + v3) + -2.25)
                _ = fi12
                _ = v13
            }
            _ = fi6
        }
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 int32 = -2147483648
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: (109 < 90), F1: ("hello" + v4), F2: v2}
    for fi17 := 0; fi17 < 1; fi17++ {
        v1 = append(v1, 42)
        if (18446744073709551615 < 381) {
            for fi18 := 0; fi18 < 19; fi18++ {
                if false {
                    var v19 int16 = -32768
                    var v20 uint = (v2 + (^(v2)))
                    fmt.Println(fi18, fi17)
                    _ = v19
                    _ = v20
                } else {
                    fmt.Println(v4, v4, v3)
                    if len(v1) > 0 {
                        v1[0] = 205
                    }
                }
                v1 = append(v1, 2)
                var v21 uint32 = 4294967295
                _ = fi18
                _ = v21
            }
            fmt.Println(v3, fi17)
            v1 = append(v1, 2147483647)
        }
        var v22 int64 = -9223372036854775808
        _ = fi17
        _ = v22
    }
    fmt.Println(v2, v2)
    v4 = v4
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v23 float64 = (1e10 / (float64(v3) / -2.25))
    for fi24 := 0; fi24 < 11; fi24++ {
        switch fi24 {
        case 884:
            v1 = append(v1, 89)
        case -101:
            if (!false) {
                if (false && false) {
                    v25 := GenS2{F0: 2147483647, F1: 2}
                    v23 = (((v23 + 0.1) + v23) / v23)
                    v4 = v4
                    _ = v25
                } else {
                    if len(v1) > 0 {
                        v1[0] = 2147483648
                    }
                    if (fi24 != (^(fi24))) {
                        v3 -= v3
                        v4 = v4
                        _ = len(v1)
                    }
                    if (true || (false || true)) {
                        fmt.Println(v23, fi24, v23)
                    }
                }
                var v26 uint64 = 311
                for fi27 := 0; fi27 < 9; fi27++ {
                    switch (v4 + v4) {
                    case "\n\t":
                        if len(v1) > 0 {
                            v1[0] = 0
                        }
                    }
                    _ = fi27
                }
                _ = v26
            } else {
                var v28 uint8 = 151
                var v29 float32 = 0.0
                _ = v28
                _ = v29
            }
        case -484:
            if true {
                v30 := GenS3{F0: 0}
                var v31 float32 = 3.14159
                _ = v30
                _ = v31
            } else {
                v23 = (((3.14159 / v23) / v23) * (v23 - ((float64(1.5) * v23) + 1.5)))
                var v32 bool = ((v23 / (-2.25 + v23)) > v23)
                _ = v32
            }
        }
        v1 = append(v1, 1)
        _ = fi24
    }
    if true {
        v2 = 255
        v4 = ((v4 + "\n\t") + "foo bar")
        var v33 int = 42
        _ = v33
    } else {
        v1 = append(v1, 4294967295)
    }
    for fi34 := 0; fi34 < 3; fi34++ {
        fmt.Println(v3, v4)
        if true {
            if (int8(v2) != -128) {
                var v35 string = "a"
                v35 = ("" + (((((("" + (((("" + ((v35 + v4) + (((v35 + (v4 + "a")) + v4) + (((v4 + "foo bar") + v35) + v35)))) + v4) + (("a" + v4) + "foo bar")) + v35)) + v35) + (("\n\t" + "") + v4)) + v35) + "\n\t") + v35))
                for fi36 := 0; fi36 < 9; fi36++ {
                    v37 := GenS1{F0: (!(!(70 <= 128))), F1: v35, F2: 4294967295}
                    var v38 uint8 = 0
                    _ = fi36
                    _ = v37
                    _ = v38
                }
                _ = v35
            } else {
                if (0 < 42) {
                    v2 += (^(v2))
                } else {
                    switch ("foo bar" + v4) {
                    case "":
                        v39 := GenS4{F0: 65535, F1: 2}
                        _ = v39
                    case "a":
                        _ = len(v1)
                    case "hello":
                        switch (fi34 / ((-9223372036854775808) | 1)) {
                        case -87:
                            var v40 int64 = 42
                            _ = v40
                        case -9223372036854775808:
                            var v41 uint8 = 2
                            _ = v41
                        case 42:
                            v42 := GenS5{F0: 2147483647, F1: 0}
                            _ = v42
                        }
                    }
                    fmt.Println(v23, v4)
                    v2 -= 65535
                }
                v23 = 1e10
                v23 = 1.5
            }
            for fi43 := 0; fi43 < 7; fi43++ {
                var v44 int32 = 2147483647
                v44 -= v44
                v44 = v44
                _ = fi43
                _ = v44
            }
            var v45 bool = false
            _ = v45
        }
        _ = fi34
    }
    for fi46 := 0; fi46 < 12; fi46++ {
        fmt.Println(fi46, fi46, v4)
        v23 -= (0.0 - v23)
        if true {
            v1 = append(v1, 2147483647)
            var v47 int8 = 1
            _ = v47
        }
        _ = fi46
    }
    fmt.Println(v4, v4)
    var v48 float64 = 3.14159
    for fi49 := 0; fi49 < 17; fi49++ {
        if ((v2 >= v2) || false) {
            v23 = ((v48 / v48) + 1.5)
        } else {
            fmt.Println(v2, v2)
        }
        _ = fi49
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v48
}

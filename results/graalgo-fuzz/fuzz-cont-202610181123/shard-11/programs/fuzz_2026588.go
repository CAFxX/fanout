package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 int64; F1 bool; F2 bool }

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = 0.0
    var v3 int64 = -1
    var v4 bool = ((-1 > 100) && (2 > 255))
    v5 := make(map[int]string)
    fmt.Println(v4, v2)
    fmt.Println(v4)
    v2 = v2
    v6 := GenS1{F0: 571}
    fmt.Println(v3)
    fmt.Println(v3, v3, v3)
    for fi7 := 0; fi7 < 2; fi7++ {
        switch "a" {
        case "foo bar":
            switch "\n\t" {
            case "a":
                fmt.Println(v2)
            }
        case "hello":
            if (42 > 4294967295) {
                v2 += (-2.25 * (v2 * v2))
                switch "a" {
                case "\n\t":
                    var v8 uint16 = 500
                    _ = v8
                case "":
                    v3 -= v3
                case "foo bar":
                    if v4 {
                        switch -665 {
                        case 1:
                            v4 = (429 == 0)
                        case 4294967295:
                            switch fi7 {
                            case 0:
                                v9 := []int{-20}
                                _ = v9
                            case 128:
                                fmt.Println(v4)
                            }
                        case -518:
                            for fi10 := 0; fi10 < 14; fi10++ {
                                v11 := []int{}
                                v12 := v1["\n\t"]
                                v3 = v3
                                _ = fi10
                                _ = v11
                                _ = v12
                            }
                        default:
                            switch fi7 {
                            case 9223372036854775807:
                                var v13 int32 = -30
                                _ = v13
                            case 42:
                                var v14 int16 = 0
                                _ = v14
                            default:
                                var v15 float64 = 3.14159
                                _ = v15
                            }
                        }
                        v2 += 3.14159
                        v4 = ((3.14159 / ((0.0 - -2.25) + (1e10 + (-2.25 + 0.0)))) == 0.1)
                    } else {
                        v3 -= 0
                        if (!v4) {
                            var v16 int = 32768
                            _ = v16
                        } else {
                            for fi17 := 0; fi17 < 5; fi17++ {
                                v18 := []int{(fi7 >> ((1) & 63))}
                                v4 = false
                                fmt.Println(v2)
                                _ = fi17
                                _ = v18
                            }
                            var v19 int64 = (v3 << ((600) & 63))
                            _ = v19
                        }
                        for fi20 := 0; fi20 < 14; fi20++ {
                            switch ("foo bar" + "\n\t") {
                            case "":
                                v21 := []int{}
                                _ = v21
                            case "a":
                                v3 -= -61
                                v2 = (v2 - v2)
                            }
                            v3 = v3
                            _ = fi20
                        }
                    }
                }
            } else {
                var v22 string = ((("a" + "a") + ("hello" + "")) + "a")
                _ = v22
            }
        default:
            v3 = (v3 ^ v3)
        }
        delete(v5, -2147483648)
        var v23 int32 = 0
        _ = fi7
        _ = v23
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        if (255 >= 1) {
            var v25 bool = false
            var v26 float32 = v2
            _ = v25
            _ = v26
        }
        v3 -= -471
        if (!v4) {
            v2 = v2
            var v27 uint64 = 18446744073709551615
            if false {
                if (2 < 2) {
                    var v28 int8 = -128
                    v29 := GenS1{F0: (fi24 & (fi24 % ((fi24) | 1)))}
                    _ = v28
                    _ = v29
                } else {
                    v3 -= -1
                    switch (fi24 * 121) {
                    case 542:
                        var v30 bool = (v4 && (42 > 1))
                        _ = v30
                    }
                    v4 = true
                }
            }
            _ = v27
        } else {
            if (v4 && false) {
                v31 := v1["a"]
                switch "foo bar" {
                case "\n\t":
                    switch ((("foo bar" + "\n\t") + "hello") + ("\n\t" + "\n\t")) {
                    case "":
                        fmt.Println(fi24)
                    case "\n\t":
                        if true {
                            v4 = ("" == "foo bar")
                            v2 = (v2 / v2)
                        } else {
                            var v32 uint16 = 42
                            _ = v32
                        }
                    case "a":
                        for fi33 := 0; fi33 < 12; fi33++ {
                            var v34 uint32 = 1
                            v3 -= (v3 >> ((255) & 63))
                            fmt.Println(fi33, fi33, fi24)
                            _ = fi33
                            _ = v34
                        }
                    default:
                        v35 := GenS2{F0: 42}
                        _ = v35
                    }
                }
                _ = v31
            }
        }
        _ = fi24
    }
    for fi36 := 0; fi36 < 4; fi36++ {
        fmt.Println(v3)
        fmt.Println(v3, fi36)
        if false {
            var v37 int32 = 42
            var v38 uint64 = 458
            _ = v37
            _ = v38
        }
        _ = fi36
    }
    fmt.Println(v3, v4)
    for fi39 := 0; fi39 < 16; fi39++ {
        v40 := GenS1{F0: fi39}
        switch (fi39 * (fi39 + (-(fi39)))) {
        case 1:
            for fi41 := 0; fi41 < 1; fi41++ {
                switch fi41 {
                case 228:
                    v2 = (0.0 / 0.1)
                case -761:
                    if ("\n\t" <= ("\n\t" + "foo bar")) {
                        var v42 string = "hello"
                        _ = len(v5)
                        v43 := GenS1{F0: (fi41 % ((fi39) | 1))}
                        var v44 uint8 = 215
                        v45 := GenS1{F0: -2147483648}
                        _ = v42
                        _ = v43
                        _ = v44
                        _ = v45
                    }
                    v3 += v3
                case 1:
                    v3 = (v3 & 9223372036854775807)
                }
                var v46 uint32 = 42
                _ = fi41
                _ = v46
            }
        case 0:
            var v47 uint16 = 563
            _ = v47
        case 4294967295:
            v48 := GenS2{F0: -104}
            _ = v48
        default:
            v2 += 1.5
        }
        _ = fi39
        _ = v40
    }
    v2 += v2
    v49 := GenS3{F0: (!(!(2 > 255))), F1: v3}
    fmt.Println(v3)
    for fi50 := 0; fi50 < 3; fi50++ {
        if true {
            var v51 int16 = 42
            _ = v51
        } else {
            for fi52 := 0; fi52 < 8; fi52++ {
                if (fi52 == (^(fi52))) {
                    if (v4 || (true && v4)) {
                        v1["\n\t"] = true
                        v4 = (!(("foo bar" + "a") < ""))
                        v2 -= v2
                    }
                    if true {
                        switch ("hello" + "hello") {
                        case "":
                            var v53 int8 = -11
                            _ = v53
                        }
                        v54 := GenS2{F0: 10}
                        _ = v54
                    }
                    var v55 string = "\n\t"
                    _ = v55
                } else {
                    fmt.Println(fi52)
                    var v56 uint16 = 1
                    _ = v56
                }
                v3 = (v3 | (v3 % ((0) | 1)))
                _ = fi52
            }
        }
        v57 := GenS4{F0: (v3 | (v3 / (((-(v3))) | 1))), F1: (9223372036854775807 != (v3 - v3)), F2: v4}
        fmt.Println(fi50, v2, v3)
        _ = fi50
        _ = v57
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}

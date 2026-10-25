package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16 }

func main() {
    var v1 float32 = (1e10 + 3.14159)
    var v2 uint64 = 1
    var v3 uint8 = 127
    v4 := make(map[int]int)
    v5 := []string{("hello" + ("" + "")), ("foo bar" + "foo bar"), ((("" + ((("a" + ("foo bar" + "")) + "") + "hello")) + "foo bar") + ("hello" + (("" + "") + "hello"))), ("\n\t" + "a")}
    v6 := make(map[int]string)
    v1 += 1e10
    fmt.Println(v3)
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 int32 = 42
        fmt.Println(v1)
        _ = fi7
        _ = v8
    }
    v1 = 0.0
    var v9 uint8 = (^(v3))
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 uint8 = v3
        var v12 int = fi10
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 uint32 = 42
    switch "hello" {
    case "foo bar":
        if ((!(((false && (v1 == 0.1)) || true) && (31 < -92))) || false) {
            v13 = 57
            delete(v6, 90)
        }
    default:
        v2 -= v2
    }
    fmt.Println(v1, v1, v3)
    var v14 int = 632
    var v15 uint64 = (v2 / (((v2 ^ (v2 - (v2 | 4294967295)))) | 1))
    if (!true) {
        for fi16 := 0; fi16 < 7; fi16++ {
            v13 -= (v13 / ((187) | 1))
            _ = fi16
        }
        var v17 uint16 = uint16(v13)
        switch (((("\n\t" + "a") + ("a" + "hello")) + (("" + "hello") + "")) + "foo bar") {
        case "":
            var v18 float32 = v1
            _ = v18
        case "foo bar":
            switch (v14 + -103) {
            case 283:
                v1 -= 1.5
            default:
                for fi19 := 0; fi19 < 15; fi19++ {
                    if ((v1 * v1) < v1) {
                        var v20 uint16 = (v17 >> ((65535) & 63))
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 17; fi21++ {
                        var v22 uint = 924
                        v1 += (v1 + v1)
                        _ = fi21
                        _ = v22
                    }
                    var v23 float64 = ((float64((v1 + (0.0 + 1e10))) - -2.25) + float64((float32(((float64(-2.25) - 1.5) * (((0.1 - (0.0 + float64(v1))) * -2.25) - float64(v1)))) + 1e10)))
                    _ = fi19
                    _ = v23
                }
            }
        case "\n\t":
            v1 += v1
        }
        _ = v17
    } else {
        for fi24 := 0; fi24 < 7; fi24++ {
            switch ("foo bar" + "hello") {
            case "\n\t":
                var v25 uint = 18446744073709551615
                _ = v25
            }
            var v26 int8 = -40
            _ = fi24
            _ = v26
        }
        for fi27 := 0; fi27 < 1; fi27++ {
            switch fi27 {
            case 32768:
                v28 := GenS1{F0: -141, F1: 2147483647}
                _ = v28
            case 2147483647:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            default:
                var v29 bool = (32767 < 42)
                _ = v29
            }
            if (42 >= 1) {
                v30 := GenS2{F0: 912}
                for fi31 := 0; fi31 < 18; fi31++ {
                    if (!(-1 == 1)) {
                        fmt.Println(v1)
                        v13 = v13
                        _ = len(v4)
                        fmt.Println(v13, v14)
                    }
                    _ = fi31
                }
                switch "\n\t" {
                case "\n\t":
                    v3 = v9
                case "hello":
                    for fi32 := 0; fi32 < 17; fi32++ {
                        switch v14 {
                        case 4294967295:
                            switch ("a" + ((("" + "\n\t") + "\n\t") + "\n\t")) {
                            case "a":
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                        default:
                            v33 := GenS3{F0: (1 != 32767)}
                            _ = v33
                        }
                        _ = fi32
                    }
                case "foo bar":
                    v2 = v2
                default:
                    if ((((float64(-2.25) + (0.0 - 1.5)) * 1.5) / 1e10) != 1e10) {
                        var v34 uint16 = 42
                        v35 := GenS4{F0: 744}
                        if (true && (v13 > v13)) {
                            v14 = (fi27 << ((941) & 63))
                        }
                        _ = v34
                        _ = v35
                    }
                }
                _ = v30
            }
            _ = fi27
        }
    }
    switch ("hello" + ("" + ("" + ("a" + "")))) {
    case "":
        var v36 string = ""
        _ = v36
    case "\n\t":
        for fi37 := 0; fi37 < 11; fi37++ {
            v6[0] = ("a" + "hello")
            var v38 uint64 = (v2 / ((v2) | 1))
            _ = fi37
            _ = v38
        }
    case "hello":
        for fi39 := 0; fi39 < 17; fi39++ {
            switch (v14 | v14) {
            case 9223372036854775807:
                v9 = (v3 * v9)
            }
            for fi40 := 0; fi40 < 3; fi40++ {
                v9 -= 1
                var v41 uint16 = 42
                _ = fi40
                _ = v41
            }
            _ = fi39
        }
    default:
        v42 := GenS4{F0: -1}
        _ = v42
    }
    v4[v14] = (v14 | v14)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}

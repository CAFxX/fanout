package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 string }

func main() {
    var v1 uint16 = 65535
    var v2 float64 = float64(0.0)
    var v3 int8 = 100
    v4 := []string{"", "hello", ("hello" + "foo bar")}
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := GenS1{F0: v3}
        switch 465 {
        case 32767:
            v2 -= v2
        case -250:
            v2 += (v2 - (v2 + v2))
        }
        _ = fi5
        _ = v6
    }
    v1 += (v1 - v1)
    fmt.Println(v3, v2, v3)
    fmt.Println(v3, v3, v2)
    fmt.Println(v1)
    v2 = v2
    v7 := GenS2{F0: 127, F1: 2, F2: 32767}
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 int = int(v3)
        var v10 float32 = 1e10
        v2 = (((1e10 - v2) - ((1e10 / v2) / v2)) - -2.25)
        _ = fi8
        _ = v10
        _ = v9
    }
    if (42 != 1) {
        if (false || (false || true)) {
            var v11 int64 = -1
            var v12 float64 = v2
            var v13 string = (("foo bar" + "a") + ("a" + "a"))
            _ = v11
            _ = v12
            _ = v13
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        if (1 >= 0) {
            for fi14 := 0; fi14 < 8; fi14++ {
                for fi15 := 0; fi15 < 9; fi15++ {
                    v2 += float64(-2.25)
                    switch "" {
                    case "a":
                        var v16 int32 = -2147483648
                        _ = v16
                    case "foo bar":
                        v2 += (-2.25 * 3.14159)
                    case "\n\t":
                        if len(v4) > 0 {
                            v4[0] = "\n\t"
                        }
                        if len(v4) > 0 {
                            v4[0] = (("\n\t" + "hello") + ("foo bar" + "hello"))
                        }
                    default:
                        if len(v4) > 0 {
                            v4[0] = ""
                        }
                    }
                    if len(v4) > 0 {
                        v4[0] = ("foo bar" + (("\n\t" + "\n\t") + ""))
                    }
                    _ = fi15
                }
                switch (("hello" + "") + ((("hello" + "\n\t") + ("a" + ("hello" + "a"))) + "a")) {
                case "":
                    if true {
                        var v17 uint = 2
                        v4 = append(v4, "foo bar")
                        _ = v17
                    } else {
                        var v18 int8 = (-(v3))
                        v19 := GenS1{F0: v3}
                        _ = v18
                        _ = v19
                    }
                case "foo bar":
                    v20 := GenS1{F0: (v3 << ((0) & 63))}
                    _ = v20
                case "\n\t":
                    var v21 uint8 = 128
                    _ = v21
                default:
                    v3 = -37
                }
                _ = fi14
            }
        }
    } else {
        v2 = float64((1.5 / 1.5))
        if len(v4) > 0 {
            v4[0] = "a"
        }
        if (299 == 679) {
            v2 = (v2 - 0.0)
        } else {
            v22 := GenS1{F0: (v3 % ((v3) | 1))}
            _ = v22
        }
    }
    if len(v4) > 0 {
        v4[0] = (("" + ("\n\t" + "a")) + ("hello" + "a"))
    }
    v3 += v3
    if (!false) {
        for fi23 := 0; fi23 < 7; fi23++ {
            switch "hello" {
            case "a":
                fmt.Println(v3)
            case "hello":
                fmt.Println(fi23, v2, v1)
            case "\n\t":
                for fi24 := 0; fi24 < 13; fi24++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v25 int = fi23
                    _ = fi24
                    _ = v25
                }
            }
            var v26 int16 = -32768
            v27 := GenS3{F0: -9223372036854775808, F1: -480}
            _ = fi23
            _ = v26
            _ = v27
        }
    } else {
        var v28 float32 = float32(float64(0.1))
        var v29 int64 = 0
        v2 -= v2
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 6; fi30++ {
        switch ((((("\n\t" + "foo bar") + "a") + (("a" + "hello") + ("foo bar" + "\n\t"))) + "\n\t") + "\n\t") {
        case "foo bar":
            for fi31 := 0; fi31 < 7; fi31++ {
                var v32 float64 = 1.5
                v2 += (3.14159 * v32)
                _ = fi31
                _ = v32
            }
        case "a":
            if (1.5 >= float64(3.14159)) {
                v2 = ((v2 * v2) / v2)
                _ = len(v4)
                v33 := GenS4{F0: ""}
                _ = v33
            }
        case "\n\t":
            v3 = (v3 ^ (-(v3)))
        }
        _ = fi30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}

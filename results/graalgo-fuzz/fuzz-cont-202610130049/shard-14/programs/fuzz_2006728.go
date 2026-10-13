package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int64; F1 int8 }

func main() {
    v1 := []string{(("hello" + "foo bar") + ("a" + "hello")), "hello", "hello"}
    var v2 string = "a"
    v3 := []int{9223372036854775807}
    var v4 float32 = (float32(3.14159) / -2.25)
    v5 := GenS1{F0: -128}
    var v6 uint8 = 255
    switch ((((((v2 + "a") + v2) + "a") + "a") + v2) + (v2 + ("foo bar" + "hello"))) {
    case "a":
        var v7 uint16 = 498
        _ = v7
    }
    switch 259 {
    case 127:
        v8 := GenS2{F0: (v6 >> ((255) & 63))}
        _ = v8
    }
    switch (v2 + (v2 + "hello")) {
    case "a":
        fmt.Println(v6, v6, v6)
    case "hello":
        if true {
            if (0 >= 0) {
                if (!(!(100 == 127))) {
                    v2 = "a"
                    if ((1 == 133) || true) {
                        switch -2147483648 {
                        case 901:
                            v9 := GenS1{F0: 100}
                            _ = v9
                        case -1:
                            v10 := GenS3{F0: 65535}
                            _ = v10
                        default:
                            if true {
                                var v11 int = 2147483647
                                _ = v11
                            }
                        }
                        var v12 uint16 = 0
                        var v13 uint8 = 73
                        _ = v12
                        _ = v13
                    } else {
                        if len(v3) > 0 {
                            v3[0] = 2147483647
                        }
                    }
                }
                fmt.Println(v4, v4)
            } else {
                _ = len(v3)
                fmt.Println(v4)
                if ((false && (42 >= 42)) && (1e10 <= (v4 + (((float32(3.14159) - v4) + v4) - 0.1)))) {
                    var v14 int64 = 42
                    _ = v14
                } else {
                    switch v2 {
                    case "hello":
                        fmt.Println(v4)
                    case "foo bar":
                        fmt.Println(v4, v4, v4)
                    default:
                        v4 -= float32(0.1)
                    }
                    if true {
                        var v15 uint8 = (v6 & 1)
                        v2 = "hello"
                        _ = v15
                    } else {
                        if (!false) {
                            v6 = (v6 ^ v6)
                            var v16 uint = 2
                            _ = v16
                        } else {
                            switch -783 {
                            case 32768:
                                if (false || false) {
                                    for fi17 := 0; fi17 < 13; fi17++ {
                                        fmt.Println(v2, v2, fi17)
                                        for fi18 := 0; fi18 < 17; fi18++ {
                                            var v19 uint64 = 858
                                            var v20 int16 = 0
                                            if ((v6 << ((652) & 63)) == v6) {
                                                var v21 bool = ((v20 >> ((96) & 63)) >= (v20 | v20))
                                                v3 = append(v3, fi17)
                                                v22 := GenS3{F0: 42}
                                                _ = v21
                                                _ = v22
                                            } else {
                                                var v23 uint32 = 2147483647
                                                v2 = v2
                                                _ = v23
                                            }
                                            _ = fi18
                                            _ = v19
                                            _ = v20
                                        }
                                        _ = fi17
                                    }
                                    v24 := GenS2{F0: v6}
                                    _ = v24
                                } else {
                                    switch -622 {
                                    case 51:
                                        var v25 uint64 = 4294967295
                                        _ = v25
                                    default:
                                        v4 -= (((v4 * v4) + 0.1) * v4)
                                    }
                                    for fi26 := 0; fi26 < 7; fi26++ {
                                        for fi27 := 0; fi27 < 3; fi27++ {
                                            if len(v3) > 0 {
                                                v3[0] = (fi26 + (-(fi26)))
                                            }
                                            v28 := GenS3{F0: 198}
                                            v3 = append(v3, 9223372036854775807)
                                            _ = fi27
                                            _ = v28
                                        }
                                        v29 := GenS4{F0: 0}
                                        _ = fi26
                                        _ = v29
                                    }
                                }
                            default:
                                switch v2 {
                                case "":
                                    v30 := GenS4{F0: 526}
                                    _ = v30
                                case "foo bar":
                                    switch -1 {
                                    case -2147483648:
                                        switch (v2 + (v2 + "hello")) {
                                        case "hello":
                                            v4 -= float32(3.14159)
                                        default:
                                            var v31 float32 = v4
                                            _ = v31
                                        }
                                    case 4294967295:
                                        for fi32 := 0; fi32 < 16; fi32++ {
                                            v6 -= 255
                                            var v33 string = v2
                                            _ = fi32
                                            _ = v33
                                        }
                                    case 32767:
                                        if false {
                                            fmt.Println(v2, v2)
                                        }
                                    default:
                                        fmt.Println(v6, v2)
                                    }
                                case "hello":
                                    v34 := GenS3{F0: 585}
                                    _ = v34
                                }
                            }
                            var v35 uint32 = 1
                            _ = v35
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        switch -2147483648 {
                        case -902:
                            v6 += 248
                        case 128:
                            v36 := GenS5{F0: ("\n\t" + v2)}
                            _ = v36
                        default:
                            v4 -= v4
                        }
                    }
                }
            }
        } else {
            var v37 int8 = 100
            v37 -= v37
            v2 = v2
            _ = v37
        }
    case "\n\t":
        v2 = v2
    }
    var v38 int16 = -32768
    if (true || false) {
        v39 := GenS4{F0: -32768}
        _ = v39
    }
    for fi40 := 0; fi40 < 8; fi40++ {
        var v41 string = (v2 + v2)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if (2147483647 >= -2147483648) {
            for fi42 := 0; fi42 < 7; fi42++ {
                v3 = append(v3, 9223372036854775807)
                if false {
                    if true {
                        v41 = v41
                        switch "foo bar" {
                        case "hello":
                            if (-756 != -582) {
                                v43 := GenS6{F0: 42, F1: 42}
                                var v44 float32 = 1e10
                                var v45 int8 = 42
                                _ = v43
                                _ = v44
                                _ = v45
                            }
                        case "foo bar":
                            fmt.Println(fi42, v38)
                        default:
                            v2 = v2
                        }
                        v46 := GenS6{F0: -1, F1: 37}
                        _ = v46
                    }
                } else {
                    switch fi42 {
                    case 1:
                        switch v41 {
                        case "foo bar":
                            if false {
                                fmt.Println(v38)
                            }
                        }
                    case 42:
                        if true {
                            v3 = append(v3, fi40)
                        }
                    }
                }
                var v47 int8 = -128
                _ = fi42
                _ = v47
            }
        }
        _ = fi40
        _ = v41
    }
    var v48 uint = 4294967295
    var v49 bool = (0.0 >= -2.25)
    v38 = (v38 & -851)
    var v50 int32 = -2147483648
    v1 = append(v1, "a")
    fmt.Println(v49, v2, v2)
    v51 := GenS3{F0: 620}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}

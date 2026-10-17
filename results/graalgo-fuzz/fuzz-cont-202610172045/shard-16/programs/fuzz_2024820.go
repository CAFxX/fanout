package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint64; F1 bool; F2 int }
type GenS4 struct { F0 bool; F1 uint; F2 uint }
type GenS5 struct { F0 bool; F1 uint8; F2 int }
type GenS6 struct { F0 uint }
func genFn0() string {
    var v4 uint8 = 2
    v5 := GenS1{F0: 1}
    fmt.Println(v4)
    var v6 int16 = -6
    switch -670 {
    case -76:
        var v7 int16 = v6
        _ = v7
    case -1:
        switch "a" {
        case "foo bar":
            if true {
                for fi8 := 0; fi8 < 6; fi8++ {
                    fmt.Println(v6, fi8)
                    v9 := GenS1{F0: 9223372036854775807}
                    _ = fi8
                    _ = v9
                }
                if true {
                    for fi10 := 0; fi10 < 18; fi10++ {
                        fmt.Println(v4, fi10, v6)
                        v6 -= 32767
                        for fi11 := 0; fi11 < 13; fi11++ {
                            v4 = (v4 + (v4 | 220))
                            v6 += 32767
                            v6 = -32768
                            _ = fi11
                        }
                        _ = fi10
                    }
                    switch "" {
                    case "\n\t":
                        fmt.Println(v4, v6)
                    case "a":
                        var v12 uint8 = v4
                        _ = v12
                    default:
                        v4 = 155
                    }
                } else {
                    fmt.Println(v4, v6)
                    for fi13 := 0; fi13 < 9; fi13++ {
                        v6 -= -1
                        _ = fi13
                    }
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(v6, v6)
                    case "hello":
                        if false {
                            var v14 uint = 255
                            v15 := []int{2147483647, int(v4), 0, -181}
                            _ = v14
                            _ = v15
                        } else {
                            fmt.Println(v4, v6)
                        }
                    default:
                        var v16 int8 = 42
                        _ = v16
                    }
                }
                v17 := GenS2{F0: 172, F1: 127, F2: 0}
                _ = v17
            } else {
                v18 := GenS3{F0: 42, F1: true, F2: -2147483648}
                var v19 int64 = 9223372036854775807
                _ = v18
                _ = v19
            }
        case "hello":
            var v20 string = (("\n\t" + "\n\t") + ("a" + ""))
            _ = v20
        }
    case 37:
        fmt.Println(v4)
    }
    _ = v4
    _ = v5
    _ = v6
    return "foo bar"
}

func main() {
    var v1 int8 = 1
    var v2 uint32 = 2147483647
    var v3 int = 165
    v21 := []int{(v3 * v3), 32768, 0}
    fmt.Println(v3, v3)
    if (613 != -2147483648) {
        fmt.Println(v1, v2, v3)
    } else {
        fmt.Println(v3, v3)
        if (!((true && (107 == 42)) || (61 < -32768))) {
            fmt.Println(v2, v1, v1)
            fmt.Println(v2, v1, v1)
        }
        var v22 uint64 = 42
        _ = v22
    }
    var v23 int8 = (v1 % ((v1) | 1))
    v24 := GenS4{F0: ((255 > 128) || false), F1: 4294967295, F2: 65535}
    var v25 uint = 18446744073709551615
    if false {
        if true {
            for fi26 := 0; fi26 < 15; fi26++ {
                switch "foo bar" {
                case "hello":
                    fmt.Println(v3, v2, v3)
                }
                v27 := GenS5{F0: (!true), F1: 255, F2: (fi26 >> (((v25 - 255)) & 63))}
                fmt.Println(v23, v1, v25)
                _ = fi26
                _ = v27
            }
            v28 := GenS5{F0: (v2 > (v2 - 2147483647)), F1: uint8(v2), F2: (^(v3))}
            v29 := GenS6{F0: (v25 * v25)}
            _ = v28
            _ = v29
        } else {
            v23 = (v23 / (((^(v1))) | 1))
            _ = len(v21)
            switch (v3 / ((0) | 1)) {
            case -9223372036854775808:
                if false {
                    switch (("" + ("" + "hello")) + "a") {
                    case "foo bar":
                        if ("" < "a") {
                            if false {
                                var v30 int64 = -27
                                for fi31 := 0; fi31 < 1; fi31++ {
                                    v32 := GenS4{F0: true, F1: 18446744073709551615, F2: uint(v2)}
                                    _ = fi31
                                    _ = v32
                                }
                                fmt.Println(v2, v3)
                                _ = v30
                            }
                            if (-1 != 9223372036854775807) {
                                var v33 float32 = -2.25
                                _ = v33
                            }
                            v34 := GenS3{F0: 4294967295, F1: (!("foo bar" <= "foo bar")), F2: v3}
                            _ = v34
                        }
                    }
                }
            case 2147483647:
                fmt.Println(v2, v3, v3)
            default:
                var v35 string = ("foo bar" + "")
                _ = v35
            }
        }
    }
    var v36 float64 = 1e10
    var v37 string = ((("" + "\n\t") + "foo bar") + "a")
    var v38 float64 = v36
    v38 += v38
    for fi39 := 0; fi39 < 17; fi39++ {
        var v40 float32 = ((1.5 * ((float32(-2.25) * 0.1) / ((0.0 - float32(v38)) / (0.0 + 0.1)))) + 0.1)
        fmt.Println(fi39, v23)
        for fi41 := 0; fi41 < 2; fi41++ {
            switch v37 {
            case "a":
                _ = len(v21)
            default:
                v1 += 1
            }
            var v42 uint16 = 2
            _ = fi41
            _ = v42
        }
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v36
    _ = v37
    _ = v38
}

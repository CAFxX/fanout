package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 uint32 = 1
    v2 := []uint32{}
    var v3 uint8 = 247
    var v4 uint64 = 0
    var v5 int16 = -32768
    var v6 int16 = v5
    var v7 int = -9223372036854775808
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v3, v6)
    switch "\n\t" {
    case "\n\t":
        if false {
            for fi8 := 0; fi8 < 10; fi8++ {
                var v9 bool = true
                if ((0 > 16) || false) {
                    fmt.Println(v7, v6)
                    var v10 int8 = -1
                    _ = v10
                } else {
                    fmt.Println(v4, v7)
                }
                _ = fi8
                _ = v9
            }
            var v11 bool = false
            _ = v11
        } else {
            for fi12 := 0; fi12 < 1; fi12++ {
                v2 = append(v2, v1)
                if ((v4 | (v4 % ((v4) | 1))) <= 64) {
                    if true {
                        for fi13 := 0; fi13 < 16; fi13++ {
                            fmt.Println(v7)
                            switch "" {
                            case "a":
                                if (false || true) {
                                    _ = len(v2)
                                } else {
                                    v4 = v4
                                    if ((!(true || (!false))) && (-13 > 100)) {
                                        v5 -= (v6 | (-(v5)))
                                        var v14 int16 = 168
                                        v15 := GenS1{F0: false}
                                        _ = v14
                                        _ = v15
                                    } else {
                                        switch ("hello" + "hello") {
                                        case "\n\t":
                                            var v16 uint32 = 771
                                            var v17 int32 = 277
                                            _ = v16
                                            _ = v17
                                        case "hello":
                                            if len(v2) > 0 {
                                                v2[0] = v1
                                            }
                                            var v18 int8 = 42
                                            _ = v18
                                        default:
                                            fmt.Println(fi13, v3, v7)
                                        }
                                        var v19 int32 = -116
                                        fmt.Println(v3, v19)
                                        v7 = (fi12 * fi12)
                                        _ = v19
                                    }
                                }
                            default:
                                fmt.Println(v6)
                            }
                            v7 = (-(v7))
                            _ = fi13
                        }
                        if true {
                            var v20 uint = 4294967295
                            _ = v20
                        }
                        var v21 uint8 = v3
                        _ = v21
                    }
                    v4 += 0
                }
                if (0 < -1) {
                    _ = len(v2)
                }
                _ = fi12
            }
            v5 = v5
        }
    case "hello":
        switch v7 {
        case -1:
            var v22 uint32 = (^(v1))
            _ = v22
        case -291:
            switch v7 {
            case -674:
                v23 := GenS2{F0: v1}
                _ = v23
            case 127:
                for fi24 := 0; fi24 < 3; fi24++ {
                    for fi25 := 0; fi25 < 2; fi25++ {
                        var v26 string = "a"
                        if (!(4294967295 >= 18446744073709551615)) {
                            v1 = (v1 / ((v1) | 1))
                        } else {
                            var v27 int32 = -2147483648
                            _ = v27
                        }
                        v4 = v4
                        _ = fi25
                        _ = v26
                    }
                    _ = fi24
                }
            case -244:
                var v28 int8 = 1
                _ = v28
            default:
                var v29 uint16 = 65535
                _ = v29
            }
        case -750:
            for fi30 := 0; fi30 < 16; fi30++ {
                switch (v7 | (v7 >> ((42) & 63))) {
                case -9223372036854775808:
                    switch "" {
                    case "foo bar":
                        fmt.Println(v1, v3, v6)
                    }
                default:
                    for fi31 := 0; fi31 < 16; fi31++ {
                        v5 -= v6
                        _ = fi31
                    }
                }
                if (!(true && (1.5 == 0.0))) {
                    var v32 int32 = 42
                    _ = v32
                }
                _ = fi30
            }
        }
    }
    switch ((((("a" + ("foo bar" + ("" + ("a" + "foo bar")))) + "a") + "a") + "") + "") {
    case "foo bar":
        var v33 uint = 461
        _ = v33
    }
    var v34 float32 = (0.1 + float32(1e10))
    if true {
        v4 = 0
    } else {
        var v35 uint = 1
        _ = v35
    }
    fmt.Println(v34, v34)
    var v36 float32 = (((-2.25 * 1.5) * 0.0) / 1e10)
    fmt.Println(v5, v5, v7)
    v37 := GenS2{F0: uint32(v3)}
    var v38 int8 = 100
    var v39 uint = 558
    if ((!((v3 << ((uint(v3)) & 63)) > 255)) || false) {
        var v40 uint64 = (v4 / (((v4 >> ((v39) & 63))) | 1))
        _ = v40
    } else {
        v4 -= (v4 << ((42) & 63))
        fmt.Println(v4)
    }
    var v41 int = 42
    if ((!((true || true) || false)) || ("hello" == "a")) {
        switch v41 {
        case 0:
            v3 += (^(v3))
        case -1:
            _ = len(v2)
        }
        switch ("foo bar" + "") {
        case "\n\t":
            v42 := GenS2{F0: v1}
            _ = v42
        case "foo bar":
            if len(v2) > 0 {
                v2[0] = v1
            }
        }
        if true {
            var v43 float64 = 3.14159
            _ = v43
        }
    }
    var v44 uint32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v41
    _ = v44
    _ = v5
    _ = v6
    _ = v7
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8; F1 string }

func main() {
    var v1 int16 = 32767
    var v2 uint64 = 1
    v3 := []int{-2147483648, -1, 9223372036854775807}
    var v4 int = -2147483648
    var v5 float64 = (1e10 * (0.0 * 0.0))
    if (true && (42 != 0)) {
        v6 := GenS1{F0: v1, F1: 363}
        for fi7 := 0; fi7 < 13; fi7++ {
            for fi8 := 0; fi8 < 18; fi8++ {
                switch fi7 {
                case -418:
                    if false {
                        v9 := GenS1{F0: int16(v4), F1: 61}
                        fmt.Println(fi7, v2, v5)
                        if (-1 > -998) {
                            var v10 int = (v4 % (((v4 / ((fi7) | 1))) | 1))
                            v5 = v5
                            var v11 float64 = v5
                            _ = v10
                            _ = v11
                        } else {
                            var v12 uint = 336
                            _ = v12
                        }
                        _ = v9
                    } else {
                        v13 := GenS1{F0: -311, F1: 529}
                        _ = v13
                    }
                case 42:
                    if true {
                        v1 = (^(v1))
                        if false {
                            if false {
                                var v14 uint8 = 2
                                _ = v14
                            }
                            if (!false) {
                                switch "a" {
                                case "hello":
                                    switch (v4 * v4) {
                                    case 32768:
                                        v5 = 0.1
                                    case 946:
                                        switch (("foo bar" + "a") + "\n\t") {
                                        case "\n\t":
                                            var v15 int = (fi8 % ((v4) | 1))
                                            _ = v15
                                        case "a":
                                            fmt.Println(v4, v4)
                                        case "hello":
                                            v16 := GenS1{F0: -634, F1: uint32(v2)}
                                            _ = v16
                                        }
                                    }
                                case "foo bar":
                                    v4 += v4
                                case "\n\t":
                                    v17 := GenS1{F0: -82, F1: 4294967295}
                                    _ = v17
                                }
                                v18 := GenS2{F0: -54}
                                var v19 float64 = 1e10
                                _ = v18
                                _ = v19
                            }
                            switch fi8 {
                            case 2147483647:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            default:
                                switch 0 {
                                case 0:
                                    fmt.Println(v1, v2, v4)
                                case 127:
                                    v5 += -2.25
                                }
                            }
                        } else {
                            var v20 float32 = -2.25
                            v4 -= fi7
                            _ = v20
                        }
                    } else {
                        if len(v3) > 0 {
                            v3[0] = (v4 & (fi7 % ((-302) | 1)))
                        }
                        switch "" {
                        case "hello":
                            var v21 uint = 255
                            _ = v21
                        case "":
                            v22 := GenS1{F0: (v1 + v1), F1: 42}
                            _ = v22
                        case "a":
                            if len(v3) > 0 {
                                v3[0] = fi8
                            }
                        }
                    }
                    v5 -= float64(((0.1 - float32(v5)) - 1e10))
                    if ((!(false && true)) || false) {
                        v2 -= (v2 + (v2 & 0))
                        var v23 bool = false
                        if len(v3) > 0 {
                            v3[0] = 127
                        }
                        _ = v23
                    } else {
                        if true {
                            var v24 float64 = float64(float32((((float64(-2.25) - 0.1) + 1.5) * 0.0)))
                            _ = v24
                        } else {
                            var v25 int8 = 91
                            _ = v25
                        }
                        var v26 float64 = v5
                        fmt.Println(v1, fi8, v26)
                        _ = v26
                    }
                    switch -910 {
                    case 128:
                        if true {
                            v5 = v5
                        }
                    }
                    var v27 float64 = 0.0
                    v27 -= v5
                    v2 -= v2
                    _ = v27
                default:
                    v1 -= -32768
                }
                _ = fi8
            }
            _ = fi7
        }
        _ = v6
    }
    v5 += -2.25
    fmt.Println(v1, v4)
    fmt.Println(v1, v2)
    for fi28 := 0; fi28 < 12; fi28++ {
        fmt.Println(v5, v5, v5)
        var v29 uint64 = 475
        _ = fi28
        _ = v29
    }
    v30 := GenS3{F0: 42}
    v31 := GenS1{F0: (^(v1)), F1: 2147483647}
    for fi32 := 0; fi32 < 3; fi32++ {
        var v33 int64 = 1
        for fi34 := 0; fi34 < 20; fi34++ {
            v35 := GenS4{F0: -128, F1: ("\n\t" + ("hello" + "\n\t"))}
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v4)
            case "":
                v3 = append(v3, (v4 % ((v4) | 1)))
            case "hello":
                v5 = ((v5 - 0.1) / ((float64((1e10 * 0.1)) + v5) + (0.1 - 1e10)))
            }
            if false {
                v36 := GenS3{F0: 687}
                if (!false) {
                    if (true && false) {
                        v2 += (v2 ^ (^(v2)))
                        switch (v4 * 2147483647) {
                        case 32767:
                            v37 := GenS3{F0: 876}
                            _ = v37
                        }
                    } else {
                        var v38 uint64 = (v2 + v2)
                        _ = v38
                    }
                    v1 -= 42
                    switch int(v33) {
                    case -285:
                        switch v4 {
                        case 815:
                            var v39 uint64 = v2
                            _ = v39
                        default:
                            switch ("hello" + "a") {
                            case "hello":
                                fmt.Println(v33)
                            }
                        }
                    case -1:
                        v5 = 0.1
                    case 32767:
                        fmt.Println(fi32, v1)
                    }
                }
                v40 := GenS1{F0: (v1 * (v1 - 0)), F1: 0}
                _ = v36
                _ = v40
            } else {
                v33 = (^(v33))
            }
            _ = fi34
            _ = v35
        }
        if false {
            _ = len(v3)
            v4 = fi32
            v33 += v33
        } else {
            var v41 int8 = 42
            v1 -= v1
            v42 := GenS3{F0: 32767}
            _ = v41
            _ = v42
        }
        _ = fi32
        _ = v33
    }
    if len(v3) > 0 {
        v3[0] = (v4 << ((18446744073709551615) & 63))
    }
    var v43 uint32 = 2147483648
    v2 += (v2 << ((4294967295) & 63))
    var v44 float32 = 1.5
    var v45 int64 = -1
    fmt.Println(v43, v44, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
}

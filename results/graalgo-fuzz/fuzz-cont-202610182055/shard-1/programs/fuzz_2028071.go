package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 int16 = -190
    v2 := []string{}
    v3 := []uint32{}
    var v4 int = -608
    v1 = v1
    fmt.Println(v4)
    var v5 int32 = 70
    v6 := GenS1{F0: 18446744073709551615, F1: 158}
    var v7 float64 = 1e10
    var v8 int32 = -2147483648
    var v9 uint8 = 2
    fmt.Println(v4, v9, v8)
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 uint64 = 4294967295
        _ = fi10
        _ = v11
    }
    if (0.0 == 1e10) {
        var v12 uint64 = 0
        var v13 int64 = -1
        fmt.Println(v5, v1, v4)
        _ = v12
        _ = v13
    } else {
        v14 := GenS1{F0: 1, F1: 4294967295}
        switch (-(v4)) {
        case 32767:
            var v15 uint = 4294967295
            _ = v15
        case -2147483648:
            if false {
                _ = len(v2)
            } else {
                v5 = 2147483647
                if (((true || true) && (348 <= 46)) && (!false)) {
                    if false {
                        if (true && ((3.14159 + 0.0) <= float32((1e10 + v7)))) {
                            fmt.Println(v1)
                            fmt.Println(v9)
                            fmt.Println(v4)
                        } else {
                            fmt.Println(v9, v9, v8)
                            fmt.Println(v1)
                            v5 = (v8 ^ (v5 ^ (v8 | 409)))
                        }
                    }
                }
                v16 := GenS2{F0: 127}
                _ = v16
            }
        default:
            _ = len(v3)
        }
        var v17 int = 1
        _ = v14
        _ = v17
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        switch v4 {
        case -1:
            switch "hello" {
            case "\n\t":
                fmt.Println(v5)
            case "foo bar":
                v19 := GenS3{F0: 101}
                _ = v19
            default:
                var v20 float32 = 0.1
                _ = v20
            }
        case 9223372036854775807:
            if false {
                v21 := GenS1{F0: 351, F1: 804}
                fmt.Println(v4)
                switch (("\n\t" + "hello") + "foo bar") {
                case "hello":
                    switch (^(fi18)) {
                    case 2147483647:
                        v22 := GenS1{F0: 42, F1: 42}
                        _ = v22
                    case -2147483648:
                        v23 := GenS2{F0: 100}
                        _ = v23
                    default:
                        if len(v3) > 0 {
                            v3[0] = 2
                        }
                    }
                case "foo bar":
                    if (!(65535 > 788)) {
                        if (42 >= 4294967295) {
                            var v24 int64 = 630
                            var v25 uint = 2
                            if false {
                                switch (-(v4)) {
                                case 32768:
                                    v7 = v7
                                default:
                                    if len(v2) > 0 {
                                        v2[0] = "foo bar"
                                    }
                                }
                            } else {
                                v26 := GenS3{F0: 2}
                                _ = v26
                            }
                            _ = v24
                            _ = v25
                        }
                    } else {
                        v1 -= v1
                    }
                case "\n\t":
                    for fi27 := 0; fi27 < 4; fi27++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi27
                    }
                }
                _ = v21
            } else {
                switch "foo bar" {
                case "\n\t":
                    var v28 uint32 = uint32(v9)
                    _ = v28
                case "foo bar":
                    v2 = append(v2, "a")
                case "":
                    if true {
                        if len(v2) > 0 {
                            v2[0] = ""
                        }
                    }
                }
                v8 = v5
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
        v5 -= (v8 % (((v8 / ((v8) | 1))) | 1))
        _ = fi18
    }
    var v29 uint8 = (v9 * (v9 / ((v9) | 1)))
    v30 := GenS1{F0: 418, F1: 2}
    v8 += v5
    if ((float64(0.0) + v7) <= 3.14159) {
        v3 = append(v3, 89)
        v5 += 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

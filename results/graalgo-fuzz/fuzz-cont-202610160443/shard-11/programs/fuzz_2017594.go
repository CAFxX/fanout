package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 uint32 }

func main() {
    var v1 uint = 65535
    var v2 int8 = -1
    var v3 bool = (69 <= 2)
    v4 := []int{128}
    var v5 float64 = ((3.14159 / 1.0) / 0.1)
    var v6 string = ""
    if ((!true) || v3) {
        _ = len(v4)
    }
    fmt.Println(v6, v3, v6)
    for fi7 := 0; fi7 < 3; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := GenS1{F0: 128}
            if (v3 || (1.5 < (v5 * ((1.5 + 3.14159) / v5)))) {
                v4 = append(v4, 32767)
                v5 += (v5 + v5)
                v3 = (-7 < 42)
            } else {
                if false {
                    if true {
                        switch (v6 + "foo bar") {
                        case "hello":
                            if v3 {
                                v3 = true
                                v1 = 0
                            }
                        case "a":
                            v4 = append(v4, 128)
                        case "":
                            var v10 uint = 2
                            _ = v10
                        }
                    } else {
                        v1 = (v1 & v1)
                        if v3 {
                            if true {
                                _ = len(v4)
                                v11 := GenS2{F0: 0}
                                fmt.Println(v6, v5)
                                _ = v11
                            } else {
                                switch 42 {
                                case 42:
                                    var v12 int8 = 1
                                    _ = v12
                                }
                            }
                            v4 = append(v4, (fi8 << ((v1) & 63)))
                            v13 := GenS3{F0: 42, F1: 152}
                            _ = v13
                        }
                    }
                    var v14 uint8 = 128
                    _ = v14
                }
            }
            _ = fi8
            _ = v9
        }
        v3 = (830 != 832)
        _ = fi7
    }
    switch v6 {
    case "foo bar":
        switch -9223372036854775808 {
        case -1:
            var v15 uint64 = 18446744073709551615
            _ = v15
        }
    case "\n\t":
        if v3 {
            for fi16 := 0; fi16 < 10; fi16++ {
                v5 -= (v5 * ((3.14159 / v5) + 1.5))
                var v17 string = v6
                _ = fi16
                _ = v17
            }
        } else {
            for fi18 := 0; fi18 < 17; fi18++ {
                var v19 float32 = 1.5
                if (0 > 1) {
                    fmt.Println(v3, v2)
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                _ = fi18
                _ = v19
            }
            fmt.Println(v5)
            v2 += 100
        }
    case "hello":
        fmt.Println(v6, v3, v5)
    }
    if false {
        v1 -= v1
        switch v6 {
        case "\n\t":
            _ = len(v4)
        case "a":
            for fi20 := 0; fi20 < 15; fi20++ {
                var v21 bool = false
                _ = fi20
                _ = v21
            }
        default:
            var v22 float64 = (((v5 - v5) + (0.1 * v5)) + v5)
            _ = v22
        }
        for fi23 := 0; fi23 < 1; fi23++ {
            fmt.Println(fi23, v3, v5)
            switch v6 {
            case "\n\t":
                v6 = "a"
            default:
                for fi24 := 0; fi24 < 19; fi24++ {
                    if (4294967295 != 0) {
                        fmt.Println(fi24, v2, v2)
                        v3 = false
                        if (128 == 127) {
                            if v3 {
                                v5 = v5
                                switch fi24 {
                                case -657:
                                    _ = len(v4)
                                    v3 = (v2 == v2)
                                case 128:
                                    v2 -= 1
                                default:
                                    if ((9223372036854775807 > 9223372036854775807) || (true || v3)) {
                                        v5 -= v5
                                    }
                                }
                            }
                            var v25 float64 = v5
                            _ = v25
                        }
                    }
                    v4 = append(v4, (fi23 / (((^(fi24))) | 1)))
                    _ = fi24
                }
            }
            fmt.Println(v3)
            _ = fi23
        }
    }
    if (65535 < 2147483647) {
        var v26 uint16 = 358
        for fi27 := 0; fi27 < 9; fi27++ {
            for fi28 := 0; fi28 < 17; fi28++ {
                var v29 bool = v3
                _ = fi28
                _ = v29
            }
            _ = fi27
        }
        _ = v26
    }
    if false {
        var v30 int64 = 9223372036854775807
        _ = v30
    }
    if (((100 > (v2 >> (((v1 >> ((4294967295) & 63))) & 63))) || v3) || (0 >= 1)) {
        if true {
            var v31 bool = (v3 && v3)
            if (("foo bar" + "foo bar") < "a") {
                v32 := GenS3{F0: 1, F1: 42}
                switch "\n\t" {
                case "hello":
                    v33 := GenS2{F0: 32767}
                    _ = v33
                case "":
                    fmt.Println(v31)
                }
                _ = v32
            } else {
                v31 = (!(false && (127 != 32768)))
                var v34 float64 = (float64(float32((v5 / float64(1.5)))) - float64(0.1))
                for fi35 := 0; fi35 < 10; fi35++ {
                    for fi36 := 0; fi36 < 13; fi36++ {
                        if len(v4) > 0 {
                            v4[0] = fi35
                        }
                        v4 = append(v4, fi36)
                        v3 = ("a" >= "hello")
                        switch (^(fi35)) {
                        case 42:
                            var v37 bool = (2147483647 == -991)
                            _ = v37
                        }
                        v2 -= -22
                        _ = fi36
                    }
                    v5 += 3.14159
                    _ = fi35
                }
                _ = v34
            }
            _ = v31
        } else {
            v38 := GenS3{F0: 255, F1: 65535}
            _ = v38
        }
    }
    var v39 uint64 = 2
    var v40 uint8 = 68
    v40 = (v40 ^ (v40 / ((v40) | 1)))
    v4 = append(v4, -433)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}

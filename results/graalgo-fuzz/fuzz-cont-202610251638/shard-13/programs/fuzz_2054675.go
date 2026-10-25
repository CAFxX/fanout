package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int32; F1 string; F2 int16 }
type GenS3 struct { F0 uint; F1 uint; F2 string }

func main() {
    var v1 int16 = 1000
    v2 := []int{-238, 42, 42, 32768}
    var v3 uint = 65535
    var v4 int8 = -128
    fmt.Println(v4, v1)
    var v5 uint = (^(v3))
    v6 := GenS1{F0: 1, F1: 9223372036854775807}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v1, v4, v4)
    if ((!(18446744073709551615 > 318)) || true) {
        var v7 int64 = 42
        _ = v7
    } else {
        switch "" {
        case "a":
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 int8 = (v4 | int8(fi8))
                v10 := GenS2{F0: -292, F1: "foo bar", F2: v1}
                fmt.Println(v3)
                _ = fi8
                _ = v10
                _ = v9
            }
        default:
            for fi11 := 0; fi11 < 6; fi11++ {
                if (false && ((("foo bar" > "") || (!(float32((0.0 * -2.25)) != 0.1))) || (v4 >= -1))) {
                    fmt.Println(v3, v5)
                    fmt.Println(v5, v3, v4)
                } else {
                    switch (("foo bar" + "foo bar") + "hello") {
                    case "foo bar":
                        var v12 int = fi11
                        _ = v12
                    case "":
                        v4 = (v4 & v4)
                    case "\n\t":
                        if true {
                            _ = len(v2)
                            fmt.Println(v5, v3, v1)
                            var v13 int = (fi11 * fi11)
                            _ = v13
                        }
                    default:
                        for fi14 := 0; fi14 < 9; fi14++ {
                            _ = len(v2)
                            v4 -= v4
                            _ = fi14
                        }
                    }
                }
                _ = fi11
            }
        }
    }
    v2 = append(v2, 32768)
    if len(v2) > 0 {
        v2[0] = 32767
    }
    if ((0.0 - 0.1) <= ((1.5 + 1.5) / -2.25)) {
        if true {
            if false {
                v2 = append(v2, -249)
                for fi15 := 0; fi15 < 4; fi15++ {
                    for fi16 := 0; fi16 < 12; fi16++ {
                        var v17 int = (fi15 >> ((0) & 63))
                        switch "\n\t" {
                        case "a":
                            if len(v2) > 0 {
                                v2[0] = 1
                            }
                        }
                        _ = fi16
                        _ = v17
                    }
                    _ = fi15
                }
            } else {
                if true {
                    if (3.14159 != 3.14159) {
                        var v18 int = 32767
                        var v19 uint8 = 128
                        switch (v18 ^ (^(v18))) {
                        case -576:
                            var v20 string = ""
                            _ = v20
                        case 329:
                            v21 := GenS2{F0: 228, F1: "hello", F2: 42}
                            _ = v21
                        default:
                            var v22 int8 = v4
                            _ = v22
                        }
                        _ = v18
                        _ = v19
                    } else {
                        fmt.Println(v3, v4)
                    }
                }
                _ = len(v2)
            }
        } else {
            v4 = (v4 & v4)
            if (1e10 >= (3.14159 - 0.1)) {
                var v23 int8 = int8(v1)
                if true {
                    var v24 uint8 = 255
                    var v25 string = ""
                    _ = v24
                    _ = v25
                }
                _ = v23
            }
        }
        var v26 bool = (-776 < -2147483648)
        _ = v26
    }
    switch "a" {
    case "hello":
        v27 := GenS2{F0: 42, F1: "a", F2: 312}
        _ = v27
    case "":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "a":
        v2 = append(v2, -412)
    default:
        for fi28 := 0; fi28 < 12; fi28++ {
            if len(v2) > 0 {
                v2[0] = (fi28 ^ 0)
            }
            v5 -= (v3 & v5)
            fmt.Println(v3, v1, v1)
            _ = fi28
        }
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        for fi30 := 0; fi30 < 1; fi30++ {
            switch (fi30 / (((^(fi29))) | 1)) {
            case 4294967295:
                var v31 bool = ((fi30 / (((fi29 - fi29)) | 1)) > (fi30 ^ (^(fi29))))
                _ = v31
            case 127:
                var v32 uint16 = 1
                _ = v32
            }
            v2 = append(v2, (fi29 | 4294967295))
            _ = fi30
        }
        var v33 bool = false
        v3 -= (v3 - v5)
        _ = fi29
        _ = v33
    }
    if true {
        var v34 float32 = (0.0 / 1.5)
        _ = v34
    } else {
        if len(v2) > 0 {
            v2[0] = -2147483648
        }
        var v35 int = 0
        switch "" {
        case "hello":
            v4 += v4
        case "\n\t":
            v4 += (v4 & (v4 - 101))
        default:
            v35 -= (v35 / (((v35 | v35)) | 1))
        }
        _ = v35
    }
    v5 = 882
    v36 := GenS3{F0: (v5 / (((v3 ^ v5)) | 1)), F1: 0, F2: "hello"}
    var v37 int = -461
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := []int{32767}
    var v2 int64 = 9223372036854775807
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]int)
    var v5 uint16 = 50
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v5 = (v5 / (((v5 << ((301) & 63))) | 1))
    var v6 uint16 = 65535
    if (32767 > 589) {
        var v7 uint = 0
        _ = v7
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        for fi9 := 0; fi9 < 7; fi9++ {
            v10 := v3[633]
            var v11 uint32 = 42
            if false {
                v6 -= v5
            } else {
                delete(v3, 1)
                switch "foo bar" {
                case "foo bar":
                    var v12 int64 = v2
                    _ = v12
                default:
                    switch ("\n\t" + ((("\n\t" + "") + (("foo bar" + (("foo bar" + "\n\t") + ("hello" + ("foo bar" + ("foo bar" + "foo bar"))))) + (((("\n\t" + ("foo bar" + ("a" + ""))) + "a") + "a") + "a"))) + "foo bar")) {
                    case "foo bar":
                        v4[2147483648] = (v10 - 2147483647)
                    case "":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case "hello":
                        var v13 int = -474
                        _ = v13
                    default:
                        v5 += (v6 + 2)
                    }
                }
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        switch "a" {
        case "hello":
            v1 = append(v1, (fi14 / (((fi14 ^ 32768)) | 1)))
        case "\n\t":
            v2 = 1
        default:
            if (351 < 1) {
                if true {
                    switch 32767 {
                    case 9223372036854775807:
                        switch (fi14 ^ (fi14 ^ (fi14 ^ (fi14 & fi14)))) {
                        case 127:
                            if false {
                                for fi15 := 0; fi15 < 3; fi15++ {
                                    var v16 float64 = (0.0 / float64(0.1))
                                    v6 = v5
                                    fmt.Println(fi14, v6)
                                    _ = fi15
                                    _ = v16
                                }
                                _ = len(v1)
                            }
                        default:
                            if ((!(float64(float32((1e10 * ((-2.25 - -2.25) * (((((0.0 * ((-2.25 * 1.5) / 1.0)) - 3.14159) - (0.0 - float64((1.5 / (1e10 - ((float32(float64((1e10 + float32(float64(0.1))))) / 1.0) * -2.25)))))) + (1.5 / 1.0)) - 0.1))))) == 1.5)) || true) {
                                var v17 int16 = 0
                                _ = v17
                            }
                        }
                    }
                } else {
                    if (!false) {
                        for fi18 := 0; fi18 < 3; fi18++ {
                            var v19 float64 = -2.25
                            v19 += 1e10
                            v3[1] = int(v2)
                            var v20 uint8 = 255
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                    }
                }
                var v21 int = -2147483648
                _ = v21
            }
        }
        if (1e10 >= 0.0) {
            fmt.Println(v2)
            for fi22 := 0; fi22 < 6; fi22++ {
                switch ("" + "\n\t") {
                case "foo bar":
                    var v23 float32 = 3.14159
                    _ = v23
                case "a":
                    v5 -= (^(v6))
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                _ = fi22
            }
            var v24 uint = 602
            _ = v24
        }
        _ = fi14
    }
    var v25 string = "foo bar"
    v26 := v4[766]
    var v27 bool = (v26 > v26)
    fmt.Println(v5, v26)
    v6 = (^(v6))
    if v27 {
        if ((1.5 + -2.25) < (0.1 / 1.5)) {
            v28 := GenS1{F0: 2147483647}
            var v29 int = 128
            if len(v1) > 0 {
                v1[0] = v29
            }
            _ = v28
            _ = v29
        }
    }
    switch -1 {
    case 744:
        v5 = v6
    }
    var v30 uint16 = v5
    delete(v4, 577)
    var v31 string = "foo bar"
    switch "hello" {
    case "foo bar":
        fmt.Println(v27, v27, v5)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}

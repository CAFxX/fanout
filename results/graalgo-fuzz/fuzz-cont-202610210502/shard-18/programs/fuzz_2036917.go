package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint }
type GenS2 struct { F0 uint32; F1 int16; F2 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64; F1 bool }

func main() {
    v1 := []int{128, 9223372036854775807, -600}
    var v2 uint32 = 437
    var v3 int16 = 1
    v4 := make(map[string]string)
    var v5 int = -1
    var v6 uint8 = 2
    v5 = (v5 - (v5 * v5))
    if (!false) {
        var v7 int32 = -2147483648
        if len(v1) > 0 {
            v1[0] = -1
        }
        var v8 float32 = 1.5
        _ = v7
        _ = v8
    }
    if (0.0 <= ((1.5 / ((-2.25 - 3.14159) / 1.0)) / float32(3.14159))) {
        fmt.Println(v3, v2)
        if false {
            var v9 int8 = -128
            var v10 uint = 2
            _ = v10
            _ = v9
        } else {
            var v11 int = v5
            for fi12 := 0; fi12 < 6; fi12++ {
                var v13 int = v11
                for fi14 := 0; fi14 < 19; fi14++ {
                    fmt.Println(v2)
                    v5 -= (v11 - v5)
                    v13 = (-(v5))
                    _ = fi14
                }
                _ = fi12
                _ = v13
            }
            for fi15 := 0; fi15 < 20; fi15++ {
                switch -266 {
                case 960:
                    for fi16 := 0; fi16 < 10; fi16++ {
                        var v17 uint = 2
                        v4[""] = "foo bar"
                        v18 := GenS1{F0: -1, F1: fi16, F2: v17}
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                case -653:
                    for fi19 := 0; fi19 < 3; fi19++ {
                        v20 := GenS2{F0: v2, F1: (-(v3)), F2: (0 == -2147483648)}
                        _ = fi19
                        _ = v20
                    }
                default:
                    for fi21 := 0; fi21 < 20; fi21++ {
                        switch ("hello" + "\n\t") {
                        case "hello":
                            v4[""] = (("" + "\n\t") + "hello")
                        case "a":
                            fmt.Println(v3, v6)
                        default:
                            var v22 int64 = 9223372036854775807
                            _ = v22
                        }
                        _ = fi21
                    }
                }
                switch (v11 * 32767) {
                case -9223372036854775808:
                    v3 = 32767
                default:
                    var v23 int64 = 42
                    _ = v23
                }
                fmt.Println(v11)
                _ = fi15
            }
            _ = v11
        }
    }
    v24 := v4[""]
    var v25 int8 = -128
    if ((2147483647 > 4294967295) && false) {
        var v26 float64 = 3.14159
        _ = v26
    } else {
        v6 -= v6
        for fi27 := 0; fi27 < 13; fi27++ {
            fmt.Println(v25, v6)
            if false {
                for fi28 := 0; fi28 < 11; fi28++ {
                    switch -1 {
                    case 4294967295:
                        var v29 int32 = 999
                        _ = v29
                    case -293:
                        switch "foo bar" {
                        case "foo bar":
                            v2 = v2
                        default:
                            v30 := GenS3{F0: 42}
                            _ = v30
                        }
                    }
                    if true {
                        v24 = ((v24 + ((v24 + v24) + "a")) + v24)
                    } else {
                        v2 -= v2
                        v2 = (v2 ^ uint32(v6))
                        if false {
                            fmt.Println(fi27, fi28)
                            switch (^(fi28)) {
                            case 128:
                                if len(v1) > 0 {
                                    v1[0] = fi28
                                }
                                v5 = fi28
                            case 32768:
                                v3 -= 32767
                            }
                            var v31 float64 = -2.25
                            _ = v31
                        }
                    }
                    fmt.Println(v25, v25)
                    _ = fi28
                }
                var v32 int = int(v25)
                fmt.Println(v32, v6)
                _ = v32
            } else {
                if (42 != 1) {
                    if ((v3 ^ (v3 << ((1) & 63))) != v3) {
                        var v33 int = v5
                        fmt.Println(v25, v25, v24)
                        _ = v33
                    }
                } else {
                    if true {
                        switch (fi27 - (fi27 ^ (v5 | v5))) {
                        case -1:
                            v34 := GenS2{F0: v2, F1: 1000, F2: false}
                            _ = v34
                        case 32767:
                            _ = len(v4)
                        case 127:
                            switch "\n\t" {
                            case "":
                                if ((666 <= 65535) && true) {
                                    v35 := GenS1{F0: 2147483647, F1: (fi27 / (((fi27 * (v5 >> ((2) & 63)))) | 1)), F2: 1}
                                    _ = len(v1)
                                    _ = v35
                                }
                            case "foo bar":
                                v36 := v4[((v24 + "hello") + (("foo bar" + v24) + "foo bar"))]
                                _ = v36
                            default:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        default:
                            var v37 int64 = 0
                            _ = v37
                        }
                        if (((((0.0 + 3.14159) / 0.1) - (((float64(-2.25) - (((1.5 / -2.25) - (1.5 - float64((-2.25 - -2.25)))) * ((0.1 * -2.25) / ((float64((((1.5 - 0.0) / ((3.14159 / -2.25) / (0.0 - 3.14159))) + 0.1)) / (1.5 * 0.1)) - 0.1)))) / 1.0) * 1.5)) == float64(-2.25)) || true) {
                            v2 = (v2 ^ 65535)
                            delete(v4, v24)
                            for fi38 := 0; fi38 < 14; fi38++ {
                                var v39 int64 = 318
                                _ = fi38
                                _ = v39
                            }
                        }
                    } else {
                        if true {
                            var v40 float64 = float64(-2.25)
                            var v41 int64 = -117
                            v6 += (^(v6))
                            _ = v40
                            _ = v41
                        }
                    }
                    for fi42 := 0; fi42 < 19; fi42++ {
                        v25 = (v25 & (v25 % ((3) | 1)))
                        _ = fi42
                    }
                }
            }
            _ = fi27
        }
    }
    switch (v24 + ((v24 + v24) + "foo bar")) {
    case "\n\t":
        v2 += 860
    }
    for fi43 := 0; fi43 < 6; fi43++ {
        v5 = 127
        _ = fi43
    }
    var v44 uint8 = 42
    var v45 uint8 = v44
    var v46 uint64 = 790
    v25 -= (v25 << ((1) & 63))
    v47 := GenS4{F0: v46, F1: true}
    fmt.Println(v45)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}

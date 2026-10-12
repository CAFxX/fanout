package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 int64 = 115
    var v2 float32 = (((float32((0.0 / ((3.14159 / 1e10) * 1.5))) + (((0.1 + 0.1) + (1.5 / 0.1)) * 1e10)) * 0.1) - 0.1)
    v3 := []uint32{0}
    var v4 int16 = 461
    _ = len(v3)
    var v5 int8 = 86
    v2 += (v2 * 3.14159)
    v1 = (-(v1))
    v1 += v1
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    if len(v3) > 0 {
        v3[0] = 65535
    }
    switch "foo bar" {
    case "hello":
        v2 += 1.5
    case "":
        for fi6 := 0; fi6 < 10; fi6++ {
            for fi7 := 0; fi7 < 12; fi7++ {
                v4 -= v4
                var v8 uint8 = 214
                v5 = v5
                _ = fi7
                _ = v8
            }
            v9 := GenS1{F0: 0}
            switch "\n\t" {
            case "foo bar":
                switch (^(fi6)) {
                case 9223372036854775807:
                    var v10 string = "foo bar"
                    _ = v10
                case 42:
                    if (!false) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v11 uint8 = 127
                        _ = v11
                    } else {
                        switch (^(fi6)) {
                        case 32768:
                            v5 += v5
                        case -2147483648:
                            if false {
                                var v12 uint8 = 255
                                _ = v12
                            } else {
                                if false {
                                    v5 = v5
                                }
                                switch "" {
                                case "hello":
                                    v13 := GenS1{F0: 2147483647}
                                    _ = v13
                                }
                            }
                        }
                        for fi14 := 0; fi14 < 8; fi14++ {
                            v3 = append(v3, 65535)
                            _ = fi14
                        }
                    }
                }
            case "hello":
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
            default:
                v15 := GenS1{F0: 2147483647}
                _ = v15
            }
            _ = fi6
            _ = v9
        }
    case "a":
        v16 := GenS1{F0: 710}
        _ = v16
    }
    if true {
        var v17 float64 = (((1e10 * 0.1) * 3.14159) * -2.25)
        var v18 int = 32768
        _ = v17
        _ = v18
    }
    v1 = v1
    v19 := GenS2{F0: (v4 * (v4 & v4))}
    fmt.Println(v1, v1)
    v1 += -592
    var v20 uint16 = 1
    v1 += v1
    for fi21 := 0; fi21 < 13; fi21++ {
        var v22 float32 = (0.1 / (v2 - v2))
        switch fi21 {
        case 127:
            var v23 uint8 = 1
            _ = v23
        case 32767:
            for fi24 := 0; fi24 < 4; fi24++ {
                switch (("foo bar" + "a") + "\n\t") {
                case "foo bar":
                    var v25 int64 = (v1 - v1)
                    v5 = (-(v5))
                    v1 += v25
                    if (!(true || ((v5 + v5) >= int8(v4)))) {
                        v1 += v25
                    }
                    if true {
                        var v26 float64 = 1.5
                        if false {
                            var v27 uint = 0
                            fmt.Println(v25, v25)
                            _ = v27
                        }
                        v25 = v25
                        _ = v26
                    } else {
                        v22 = v2
                        var v28 uint8 = 2
                        _ = v28
                    }
                    _ = v25
                case "hello":
                    v29 := GenS1{F0: 442}
                    var v30 bool = true
                    _ = v29
                    _ = v30
                case "\n\t":
                    switch ("hello" + "a") {
                    case "\n\t":
                        fmt.Println(v22, v1, fi24)
                        var v31 uint8 = 1
                        _ = v31
                    case "":
                        var v32 float64 = (0.0 / 1.0)
                        _ = v32
                    case "a":
                        if false {
                            v33 := GenS2{F0: v4}
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v33
                        } else {
                            v1 = (v1 + v1)
                        }
                    default:
                        v2 -= 1.5
                    }
                }
                _ = fi24
            }
        case 775:
            for fi34 := 0; fi34 < 11; fi34++ {
                v2 = ((float32(((0.0 / 3.14159) / (float64(v22) - (float64((1e10 + v2)) * (0.1 * 1e10))))) / ((v2 - (float32(-2.25) / (v2 / ((v22 * (v2 - v22)) / float32(1e10))))) * 1.5)) * (v22 - v2))
                switch (((((("" + (("foo bar" + ("hello" + ("" + "foo bar"))) + "\n\t")) + ("a" + "a")) + "a") + "\n\t") + "hello") + "a") {
                case "\n\t":
                    v2 = v2
                case "":
                    _ = len(v3)
                case "hello":
                    var v35 int64 = (-(v1))
                    _ = v35
                }
                _ = fi34
            }
        default:
            if (255 == 65535) {
                fmt.Println(v5)
            }
        }
        _ = fi21
        _ = v22
    }
    v36 := GenS1{F0: 363}
    var v37 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}

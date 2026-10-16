package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int8 }

func main() {
    v1 := []string{"a", "\n\t", ("foo bar" + ("a" + "a"))}
    var v2 bool = ((1000 < -1) || (255 != 811))
    var v3 int64 = -9223372036854775808
    fmt.Println(v3, v3)
    v3 -= v3
    var v4 int = 633
    switch (v4 ^ 690) {
    case 32768:
        for fi5 := 0; fi5 < 14; fi5++ {
            _ = len(v1)
            switch fi5 {
            case 2147483647:
                v6 := GenS1{F0: -32768, F1: -1}
                _ = v6
            case 1:
                fmt.Println(v3)
            }
            for fi7 := 0; fi7 < 1; fi7++ {
                var v8 string = "\n\t"
                _ = fi7
                _ = v8
            }
            _ = fi5
        }
    case 42:
        var v9 int64 = v3
        _ = v9
    }
    fmt.Println(v4)
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 int16 = -32768
        var v12 bool = ((v2 && v2) || ((1e10 - (1.5 / 1.0)) >= 3.14159))
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println(v2)
    v13 := GenS2{F0: 71, F1: 114}
    v4 += v4
    v2 = false
    switch (v4 + v4) {
    case 9223372036854775807:
        for fi14 := 0; fi14 < 1; fi14++ {
            v4 = fi14
            switch "\n\t" {
            case "foo bar":
                if (22 < 100) {
                    var v15 uint8 = 42
                    v4 -= -806
                    v2 = (true || true)
                    _ = v15
                }
            default:
                var v16 uint64 = 218
                _ = v16
            }
            var v17 uint32 = 2
            _ = fi14
            _ = v17
        }
    case -993:
        v3 -= -1
    case 893:
        var v18 int8 = 0
        _ = v18
    }
    fmt.Println(v3, v3)
    var v19 int64 = (v3 & v3)
    fmt.Println(v3, v3, v3)
    switch ("a" + ("a" + (("a" + "\n\t") + "hello"))) {
    case "":
        if false {
            v4 += v4
            fmt.Println(v19)
            switch "\n\t" {
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "a":
                switch v4 {
                case 9223372036854775807:
                    var v20 uint32 = 2147483648
                    _ = v20
                default:
                    v3 = 0
                }
            case "hello":
                var v21 int = int(v19)
                _ = v21
            }
        } else {
            switch "a" {
            case "hello":
                var v22 string = ""
                _ = v22
            case "":
                switch "" {
                case "":
                    switch -9223372036854775808 {
                    case 14:
                        for fi23 := 0; fi23 < 12; fi23++ {
                            for fi24 := 0; fi24 < 10; fi24++ {
                                switch ("a" + "") {
                                case "":
                                    var v25 uint8 = 54
                                    _ = v25
                                case "hello":
                                    switch ("foo bar" + "a") {
                                    case "a":
                                        var v26 int = (fi24 - 9223372036854775807)
                                        _ = v26
                                    case "\n\t":
                                        switch 295 {
                                        case 2147483647:
                                            var v27 int32 = -1
                                            var v28 int16 = -407
                                            var v29 bool = ("hello" > "a")
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                        case 32768:
                                            v2 = (196 > -1)
                                        default:
                                            var v30 string = "\n\t"
                                            _ = v30
                                        }
                                    }
                                default:
                                    var v31 uint16 = 255
                                    _ = v31
                                }
                                var v32 int = fi23
                                v4 -= fi24
                                _ = fi24
                                _ = v32
                            }
                            if true {
                                if len(v1) > 0 {
                                    v1[0] = ""
                                }
                                v19 = v3
                            } else {
                                var v33 uint = 0
                                var v34 float32 = float32(1e10)
                                _ = v33
                                _ = v34
                            }
                            v35 := GenS1{F0: 1, F1: 1}
                            _ = fi23
                            _ = v35
                        }
                    default:
                        v1 = append(v1, "")
                    }
                case "a":
                    var v36 uint32 = 30
                    _ = v36
                case "\n\t":
                    v37 := GenS1{F0: 523, F1: 32767}
                    _ = v37
                default:
                    var v38 int16 = 32767
                    _ = v38
                }
            case "foo bar":
                fmt.Println(v3, v3, v2)
            default:
                switch (v4 >> ((141) & 63)) {
                case -729:
                    var v39 uint64 = 42
                    _ = v39
                }
            }
            var v40 int = v4
            _ = v40
        }
    case "a":
        v2 = v2
    case "\n\t":
        var v41 int64 = (v19 + v3)
        _ = v41
    }
    for fi42 := 0; fi42 < 2; fi42++ {
        fmt.Println(v4, fi42)
        if v2 {
            v19 = v19
        } else {
            v43 := GenS3{F0: -1}
            _ = v43
        }
        _ = fi42
    }
    v4 += (-(v4))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}

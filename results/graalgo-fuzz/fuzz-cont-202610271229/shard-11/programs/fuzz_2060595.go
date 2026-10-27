package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint16 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]uint32)
    v3 := make(map[string]uint32)
    var v4 uint8 = 255
    v5 := []string{("\n\t" + "hello"), ("" + "hello"), "", ""}
    v6 := GenS1{F0: 970, F1: (^(v4))}
    if len(v5) > 0 {
        v5[0] = ""
    }
    fmt.Println(v4)
    for fi7 := 0; fi7 < 10; fi7++ {
        if true {
            if (32767 == -640) {
                fmt.Println(v4)
                if true {
                    var v8 int32 = -1
                    switch 32768 {
                    case 194:
                        fmt.Println(v8, v4, v8)
                    default:
                        v4 = (v4 * (v4 >> ((18446744073709551615) & 63)))
                    }
                    _ = v8
                } else {
                    if false {
                        switch ("foo bar" + "") {
                        case "\n\t":
                            v4 = (v4 & 127)
                        case "hello":
                            if ("hello" >= "\n\t") {
                                _ = len(v5)
                            }
                        case "a":
                            switch 449 {
                            case 2147483647:
                                var v9 uint8 = v4
                                _ = v9
                            }
                        }
                        var v10 float64 = -2.25
                        _ = v10
                    } else {
                        v11 := GenS1{F0: -591, F1: (v4 | (v4 / ((v4) | 1)))}
                        fmt.Println(v4, v4)
                        _ = v11
                    }
                    v4 = v4
                    if (float64(-2.25) < 1.5) {
                        for fi12 := 0; fi12 < 6; fi12++ {
                            var v13 float64 = 1e10
                            v4 += 1
                            var v14 uint16 = 32767
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                        var v15 int = (fi7 | -748)
                        _ = v15
                    }
                }
                delete(v3, ("\n\t" + ("hello" + "\n\t")))
            } else {
                var v16 int64 = -9223372036854775808
                _ = v16
            }
        }
        _ = fi7
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        switch ("foo bar" + ("\n\t" + ((("\n\t" + "foo bar") + ("foo bar" + ("foo bar" + "foo bar"))) + "a"))) {
        case "\n\t":
            switch "a" {
            case "hello":
                fmt.Println(v4, v4)
            case "\n\t":
                v4 = 255
            }
        }
        _ = fi17
    }
    v18 := GenS2{F0: 2147483648, F1: 0, F2: 0}
    for fi19 := 0; fi19 < 10; fi19++ {
        v20 := GenS1{F0: -414, F1: 1}
        if false {
            delete(v1, "hello")
            v4 = (v4 >> ((255) & 63))
            v4 = v4
        }
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        v22 := GenS2{F0: 1, F1: 1, F2: uint16(v4)}
        _ = fi21
        _ = v22
    }
    var v23 int8 = 1
    v24 := v1["\n\t"]
    v23 = v23
    v4 = 18
    if (!false) {
        _ = len(v5)
    }
    v23 += 127
    v25 := GenS2{F0: 65535, F1: 32767, F2: 65535}
    fmt.Println(v23, v4)
    if false {
        var v26 int8 = -30
        fmt.Println(v26, v23, v24)
        v26 = v23
        _ = v26
    } else {
        var v27 uint8 = v4
        if (true && false) {
            v28 := GenS1{F0: 0, F1: v4}
            for fi29 := 0; fi29 < 17; fi29++ {
                v30 := GenS2{F0: 65535, F1: 0, F2: 65535}
                var v31 uint = 42
                v31 = (v31 << (((v31 << (((v31 + 0)) & 63))) & 63))
                _ = fi29
                _ = v30
                _ = v31
            }
            v24 = ("foo bar" + v24)
            _ = v28
        }
        switch 9223372036854775807 {
        case 32768:
            delete(v1, ((v24 + (v24 + "\n\t")) + v24))
        case -2147483648:
            fmt.Println(v27, v27)
        default:
            v32 := v3[((v24 + "") + "\n\t")]
            _ = v32
        }
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

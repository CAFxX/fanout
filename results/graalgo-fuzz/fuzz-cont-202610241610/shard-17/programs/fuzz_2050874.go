package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint }

func main() {
    v1 := []string{"hello", ("hello" + (("foo bar" + "\n\t") + ("a" + "")))}
    v2 := make(map[int]bool)
    var v3 int16 = 612
    v2[32768] = true
    var v4 float32 = 0.0
    switch "" {
    case "":
        if true {
            fmt.Println(v4)
            switch 550 {
            case -2147483648:
                switch -640 {
                case 32767:
                    for fi5 := 0; fi5 < 15; fi5++ {
                        if (458 < fi5) {
                            var v6 uint8 = 127
                            v6 += (v6 - (v6 & (v6 + v6)))
                            _ = len(v1)
                            _ = v6
                        } else {
                            for fi7 := 0; fi7 < 12; fi7++ {
                                if len(v1) > 0 {
                                    v1[0] = ((("\n\t" + "") + "") + "")
                                }
                                v8 := v2[(fi7 >> ((2) & 63))]
                                _ = fi7
                                _ = v8
                            }
                            if (0 < 18446744073709551615) {
                                for fi9 := 0; fi9 < 3; fi9++ {
                                    v4 = 1e10
                                    delete(v2, 42)
                                    _ = fi9
                                }
                                if ((820 != 0) || ((false || true) || (true || true))) {
                                    v10 := GenS1{F0: 1, F1: 42}
                                    _ = v10
                                }
                                for fi11 := 0; fi11 < 14; fi11++ {
                                    v3 -= (v3 + v3)
                                    switch "\n\t" {
                                    case "\n\t":
                                        var v12 int = fi11
                                        _ = v12
                                    default:
                                        v1 = append(v1, "hello")
                                    }
                                    _ = fi11
                                }
                            } else {
                                v13 := GenS2{F0: 769}
                                _ = v13
                            }
                            v4 = float32(3.14159)
                        }
                        _ = fi5
                    }
                }
            case 32767:
                var v14 uint16 = 255
                _ = v14
            case 42:
                if false {
                    if (((float32(1.5) + (v4 / 1e10)) * float32(0.0)) <= (v4 * 0.1)) {
                        var v15 uint32 = 2147483648
                        v3 += 176
                        switch "foo bar" {
                        case "":
                            v3 += v3
                        case "foo bar":
                            v16 := v2[2147483647]
                            _ = v16
                        case "hello":
                            if (-28 > 127) {
                                v15 = v15
                                switch "foo bar" {
                                case "\n\t":
                                    var v17 float64 = 0.0
                                    _ = v17
                                case "":
                                    v18 := GenS1{F0: -114, F1: 1}
                                    _ = v18
                                }
                            }
                        }
                        _ = v15
                    } else {
                        var v19 int32 = 0
                        _ = v19
                    }
                    v4 = float32((-2.25 - 0.1))
                    v3 = v3
                }
            }
        }
    case "foo bar":
        v3 -= (v3 ^ v3)
    default:
        var v20 int32 = -900
        _ = v20
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        var v22 int64 = 1
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        v4 -= v4
        _ = fi23
    }
    var v24 uint16 = 426
    delete(v2, -1000)
    for fi25 := 0; fi25 < 1; fi25++ {
        v24 = v24
        var v26 uint8 = 161
        _ = fi25
        _ = v26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}

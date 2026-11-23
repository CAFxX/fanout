package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 string }

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = (-2.25 - -2.25)
    var v3 uint64 = 733
    if true {
        if (!(127 > 6)) {
            fmt.Println(v3)
            v2 -= (float32(1e10) - v2)
            if (("\n\t" > "a") && true) {
                var v4 int32 = 1
                v4 -= v4
                v5 := []int{336, 2147483647, 127, 0}
                _ = v4
                _ = v5
            } else {
                v6 := []int{1, -9223372036854775808, 127}
                _ = v6
            }
        }
        v2 += (1.5 / (float32(0.1) * ((v2 - v2) + (v2 + v2))))
        for fi7 := 0; fi7 < 11; fi7++ {
            if (4294967295 <= 18446744073709551615) {
                var v8 float64 = 0.1
                var v9 uint16 = 1
                var v10 float32 = (v2 - v2)
                _ = v10
                _ = v8
                _ = v9
            } else {
                v11 := GenS1{F0: 32767, F1: 65535, F2: "foo bar"}
                v12 := GenS1{F0: 32767, F1: 42, F2: ("hello" + ("foo bar" + "foo bar"))}
                switch "" {
                case "foo bar":
                    v2 -= v2
                case "\n\t":
                    v2 += (v2 * v2)
                }
                _ = v11
                _ = v12
            }
            if (-9 != -9223372036854775808) {
                fmt.Println(v3, v2, fi7)
                switch fi7 {
                case 238:
                    var v13 uint64 = 2
                    _ = v13
                case -575:
                    for fi14 := 0; fi14 < 5; fi14++ {
                        if (!((false && (!true)) || (v2 <= ((((v2 - (3.14159 / 3.14159)) * 0.1) + (float32(3.14159) / 1.0)) + v2)))) {
                            v3 = (v3 << ((4294967295) & 63))
                            v15 := GenS1{F0: 486, F1: 35, F2: ""}
                            v3 = (v3 >> ((254) & 63))
                            _ = v15
                        }
                        _ = fi14
                    }
                }
            }
            _ = fi7
        }
    } else {
        for fi16 := 0; fi16 < 2; fi16++ {
            for fi17 := 0; fi17 < 8; fi17++ {
                switch fi17 {
                case 32767:
                    var v18 uint16 = 42
                    _ = v18
                case -436:
                    if true {
                        switch (("a" + ("foo bar" + "foo bar")) + ("a" + "")) {
                        case "foo bar":
                            v2 += v2
                        case "\n\t":
                            var v19 uint8 = 127
                            _ = v19
                        case "a":
                            if (false && true) {
                                var v20 uint16 = 1
                                v21 := GenS1{F0: 32767, F1: 65535, F2: ("" + "\n\t")}
                                _ = v20
                                _ = v21
                            }
                        }
                        var v22 uint64 = (^(v3))
                        _ = v22
                    }
                default:
                    switch fi16 {
                    case 32768:
                        if (!(!true)) {
                            var v23 float32 = (v2 / (float32(float64(1e10)) / v2))
                            _ = v23
                        }
                    default:
                        var v24 uint8 = 255
                        _ = v24
                    }
                }
                _ = fi17
            }
            var v25 int16 = 1000
            fmt.Println(v2, v3)
            _ = fi16
            _ = v25
        }
        v2 = v2
        var v26 bool = (((0 >= 32767) || false) || ((-46 != 0) || (!false)))
        _ = v26
    }
    v2 += (v2 + (1e10 / v2))
    v3 = v3
    var v27 uint = 42
    v3 += v3
    fmt.Println(v3, v3, v3)
    var v28 float64 = 1.5
    v2 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
}

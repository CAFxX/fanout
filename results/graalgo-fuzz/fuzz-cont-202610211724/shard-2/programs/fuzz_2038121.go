package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint16; F1 uint8 }

func main() {
    var v1 uint64 = 4294967295
    var v2 int16 = 42
    v3 := make(map[uint32]bool)
    v4 := make(map[string]int)
    v2 += 32767
    _ = len(v4)
    if false {
        if false {
            v1 -= v1
        } else {
            if (1 > 2147483647) {
                if (2147483647 != -269) {
                    var v5 int = 865
                    _ = v5
                } else {
                    if false {
                        v4["\n\t"] = -1
                        switch "" {
                        case "foo bar":
                            switch "" {
                            case "\n\t":
                                if (1.5 == ((0.1 - -2.25) * 3.14159)) {
                                    v6 := GenS1{F0: 18446744073709551615, F1: -32}
                                    v1 -= 1
                                    var v7 float64 = -2.25
                                    _ = v6
                                    _ = v7
                                } else {
                                    v1 -= (v1 | 796)
                                }
                            default:
                                v8 := []int{127}
                                _ = v8
                            }
                        }
                        for fi9 := 0; fi9 < 3; fi9++ {
                            for fi10 := 0; fi10 < 10; fi10++ {
                                switch -1 {
                                case 42:
                                    v11 := GenS1{F0: 65535, F1: -128}
                                    _ = v11
                                case 4294967295:
                                    var v12 uint64 = (v1 % (((v1 + v1)) | 1))
                                    _ = v12
                                }
                                var v13 int64 = -1
                                _ = fi10
                                _ = v13
                            }
                            _ = fi9
                        }
                    } else {
                        _ = len(v3)
                    }
                    if (false && (255 >= 0)) {
                        v1 = 42
                        switch 1 {
                        case -2147483648:
                            fmt.Println(v1)
                        }
                    }
                }
            }
            fmt.Println(v1, v2)
            for fi14 := 0; fi14 < 7; fi14++ {
                v15 := []int{}
                fmt.Println(fi14)
                _ = fi14
                _ = v15
            }
        }
        for fi16 := 0; fi16 < 3; fi16++ {
            fmt.Println(v2, v1, v1)
            v17 := GenS2{F0: 255, F1: 0}
            _ = fi16
            _ = v17
        }
    } else {
        fmt.Println(v1, v2)
    }
    switch -1 {
    case 893:
        v18 := []int{-437, -9223372036854775808, 2147483647}
        _ = v18
    case -2147483648:
        v19 := []int{32767}
        _ = v19
    default:
        var v20 uint8 = 127
        _ = v20
    }
    if (!true) {
        switch (("\n\t" + ("" + ("hello" + "foo bar"))) + ("\n\t" + "\n\t")) {
        case "":
            v21 := GenS1{F0: 2, F1: 127}
            _ = v21
        case "foo bar":
            if false {
                v22 := GenS2{F0: 65535, F1: 107}
                var v23 int8 = 45
                _ = v22
                _ = v23
            } else {
                v2 = v2
                for fi24 := 0; fi24 < 1; fi24++ {
                    v2 -= v2
                    fmt.Println(fi24, v1, v1)
                    switch ("\n\t" + "\n\t") {
                    case "foo bar":
                        v25 := GenS1{F0: 65535, F1: 0}
                        _ = v25
                    case "":
                        v1 = v1
                    }
                    _ = fi24
                }
                var v26 float64 = 0.1
                _ = v26
            }
        }
    }
    v27 := v4[("a" + "")]
    v28 := []int{v27, (v27 & (v27 - (v27 >> ((365) & 63))))}
    switch ("a" + "a") {
    case "":
        if (!(2147483647 > 2)) {
            fmt.Println(v27)
            var v29 int32 = 1
            if false {
                var v30 uint16 = 1
                var v31 int32 = v29
                _ = v30
                _ = v31
            } else {
                v32 := GenS2{F0: 2, F1: uint8(v1)}
                _ = v32
            }
            _ = v29
        } else {
            v33 := GenS2{F0: 255, F1: 42}
            _ = v33
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }

func main() {
    v1 := []uint32{}
    var v2 float32 = -2.25
    var v3 string = "foo bar"
    switch v3 {
    case "hello":
        if (3.14159 != (((1e10 / 1.5) + 0.1) - (0.0 - 0.0))) {
            switch -9223372036854775808 {
            case 430:
                if true {
                    for fi4 := 0; fi4 < 14; fi4++ {
                        for fi5 := 0; fi5 < 1; fi5++ {
                            if false {
                                if len(v1) > 0 {
                                    v1[0] = 122
                                }
                            } else {
                                var v6 int = fi4
                                v3 = "foo bar"
                                _ = v6
                            }
                            fmt.Println(fi4, fi4, v3)
                            switch (fi4 + (fi4 ^ 127)) {
                            case 0:
                                var v7 uint8 = 1
                                _ = v7
                            case 332:
                                switch fi4 {
                                case 9223372036854775807:
                                    v8 := GenS1{F0: 127}
                                    _ = v8
                                case 128:
                                    if false {
                                        v9 := GenS2{F0: -72}
                                        _ = v9
                                    }
                                case 0:
                                    v10 := GenS2{F0: 0}
                                    _ = v10
                                default:
                                    var v11 uint64 = 4294967295
                                    _ = v11
                                }
                            default:
                                if true {
                                    var v12 uint = 1
                                    v3 = ("foo bar" + (v3 + (("a" + v3) + (v3 + v3))))
                                    var v13 float64 = 0.1
                                    _ = v12
                                    _ = v13
                                } else {
                                    fmt.Println(v3)
                                }
                            }
                            _ = fi5
                        }
                        _ = fi4
                    }
                    fmt.Println(v2, v3)
                } else {
                    if true {
                        v3 = (("hello" + (v3 + v3)) + v3)
                        v1 = append(v1, 271)
                    }
                    var v14 int32 = 2147483647
                    v15 := GenS2{F0: -19}
                    _ = v14
                    _ = v15
                }
            case 127:
                switch 127 {
                case 127:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    fmt.Println(v2, v3)
                }
            case 1:
                if false {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v16 := GenS2{F0: 100}
                    var v17 uint8 = 87
                    _ = v16
                    _ = v17
                }
            }
            fmt.Println(v2)
            for fi18 := 0; fi18 < 2; fi18++ {
                _ = len(v1)
                var v19 int64 = 1
                _ = fi18
                _ = v19
            }
        }
    case "a":
        v3 = (v3 + (((v3 + v3) + "") + v3))
    case "foo bar":
        var v20 uint64 = 789
        _ = v20
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 uint8 = 127
        _ = fi21
        _ = v22
    }
    fmt.Println(v3, v2)
    v23 := GenS1{F0: 127}
    for fi24 := 0; fi24 < 15; fi24++ {
        for fi25 := 0; fi25 < 12; fi25++ {
            switch (fi25 & fi25) {
            case -114:
                v1 = append(v1, 0)
            case 128:
                if ((42 < 42) || (0 < 4294967295)) {
                    var v26 int8 = 127
                    switch (fi25 - fi25) {
                    case 191:
                        v3 = v3
                    case -724:
                        var v27 int8 = (v26 / ((v26) | 1))
                        _ = v27
                    case 42:
                        if true {
                            _ = len(v1)
                            switch (fi25 - fi24) {
                            case 0:
                                v3 = ("" + v3)
                            case -2147483648:
                                v1 = append(v1, 98)
                            }
                        } else {
                            var v28 int32 = -1
                            _ = v28
                        }
                    default:
                        fmt.Println(v2, fi24)
                    }
                    v29 := GenS2{F0: v26}
                    _ = v26
                    _ = v29
                } else {
                    if ((true || (2 <= 42)) || ((true || true) || false)) {
                        v3 = ((v3 + "\n\t") + v3)
                        v2 = v2
                    }
                    v2 = (v2 * 1e10)
                }
            default:
                v30 := GenS1{F0: 255}
                _ = v30
            }
            v31 := GenS1{F0: 127}
            _ = fi25
            _ = v31
        }
        _ = fi24
    }
    v1 = append(v1, 472)
    v1 = append(v1, 42)
    var v32 int32 = -1
    fmt.Println(v32, v2, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v32
}

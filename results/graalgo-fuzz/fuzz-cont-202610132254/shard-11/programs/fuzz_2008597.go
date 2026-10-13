package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint = 18446744073709551615
    v2 := []int{32767, 127}
    v3 := make(map[string]int)
    var v4 uint16 = 1
    v1 = (v1 & (v1 ^ (^(v1))))
    for fi5 := 0; fi5 < 1; fi5++ {
        for fi6 := 0; fi6 < 17; fi6++ {
            switch fi6 {
            case -9223372036854775808:
                if (0 == 2147483647) {
                    v4 -= (v4 + v4)
                    v7 := GenS1{F0: 86}
                    _ = v7
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        if false {
                            if (!(!true)) {
                                v4 -= v4
                                var v8 int32 = 957
                                _ = v8
                            } else {
                                v4 = (v4 * v4)
                            }
                        }
                    }
                }
            case 9223372036854775807:
                v2 = append(v2, (fi5 * fi5))
            case -2147483648:
                v9 := GenS1{F0: int8(fi5)}
                _ = v9
            default:
                var v10 uint64 = 967
                _ = v10
            }
            _ = fi6
        }
        _ = fi5
    }
    var v11 uint8 = 2
    var v12 int16 = int16(v4)
    var v13 int = int(v4)
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 float64 = 1e10
        if len(v2) > 0 {
            v2[0] = fi14
        }
        _ = fi14
        _ = v15
    }
    fmt.Println(v13, v12, v11)
    v1 -= v1
    if true {
        if (v1 <= v1) {
            switch "" {
            case "":
                switch "hello" {
                case "":
                    v1 = v1
                default:
                    v16 := v3[("foo bar" + "")]
                    _ = v16
                }
            case "foo bar":
                switch ("a" + "hello") {
                case "":
                    if false {
                        if (936 != 2147483647) {
                            switch -828 {
                            case -189:
                                var v17 bool = true
                                _ = v17
                            case 0:
                                switch "a" {
                                case "\n\t":
                                    v18 := GenS1{F0: -1}
                                    _ = v18
                                case "":
                                    v12 = v12
                                }
                            case 127:
                                v19 := GenS1{F0: 0}
                                _ = v19
                            }
                            v13 -= int(v12)
                        } else {
                            var v20 bool = (v4 >= (v4 | v4))
                            var v21 bool = true
                            _ = v20
                            _ = v21
                        }
                        var v22 float32 = 0.0
                        _ = v22
                    }
                case "foo bar":
                    v3[(("a" + ("foo bar" + "hello")) + "")] = (v13 / ((v13) | 1))
                }
            }
        }
        var v23 int32 = 2147483647
        _ = v23
    } else {
        var v24 uint8 = v11
        _ = len(v3)
        _ = v24
    }
    v25 := GenS1{F0: 127}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}

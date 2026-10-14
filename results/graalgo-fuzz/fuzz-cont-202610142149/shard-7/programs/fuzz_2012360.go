package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 bool }
type GenS2 struct { F0 string; F1 uint16; F2 int }

func main() {
    var v1 int64 = -967
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 1
    v4 := make(map[int]string)
    switch ("\n\t" + "") {
    case "a":
        switch -300 {
        case 0:
            fmt.Println(v1, v1)
        case 362:
            v1 -= 42
        }
    }
    v5 := []int{-9223372036854775808}
    var v6 float32 = (-2.25 - ((3.14159 * float32(1.5)) - (1.5 - ((0.1 * -2.25) * 1.5))))
    v7 := GenS1{F0: 65535, F1: 127, F2: (4294967295 <= 295)}
    _ = len(v5)
    fmt.Println(v1)
    var v8 uint16 = 255
    var v9 int64 = -1
    for fi10 := 0; fi10 < 18; fi10++ {
        switch "" {
        case "":
            var v11 uint32 = 928
            _ = v11
        case "foo bar":
            v12 := GenS1{F0: (v8 ^ v8), F1: -1, F2: (!(!true))}
            _ = v12
        case "a":
            switch ("foo bar" + ("a" + "a")) {
            case "foo bar":
                if (24 > 127) {
                    for fi13 := 0; fi13 < 4; fi13++ {
                        v2[42] = 166
                        v14 := GenS1{F0: 2, F1: -1, F2: (-128 != -128)}
                        v1 = -9223372036854775808
                        _ = fi13
                        _ = v14
                    }
                } else {
                    var v15 uint = 65535
                    for fi16 := 0; fi16 < 13; fi16++ {
                        v9 -= 42
                        _ = fi16
                    }
                    v17 := GenS1{F0: 721, F1: 127, F2: false}
                    _ = v15
                    _ = v17
                }
            case "\n\t":
                switch "a" {
                case "":
                    for fi18 := 0; fi18 < 2; fi18++ {
                        if true {
                            v19 := GenS1{F0: 32767, F1: 0, F2: true}
                            var v20 int16 = -32768
                            if true {
                                if (((float64((((1e10 * v6) + v6) - (v6 / v6))) - -2.25) / 1.0) == 1.5) {
                                    v9 = (v9 | (v9 + v1))
                                    delete(v4, -876)
                                    var v21 int64 = v9
                                    _ = v21
                                }
                                var v22 float64 = 0.1
                                _ = v22
                            } else {
                                _ = len(v4)
                            }
                            _ = v19
                            _ = v20
                        } else {
                            v1 += v1
                            var v23 uint64 = (^(v3))
                            _ = v23
                        }
                        _ = fi18
                    }
                case "\n\t":
                    v8 = v8
                default:
                    fmt.Println(v9, fi10, v6)
                }
            default:
                v24 := GenS2{F0: "foo bar", F1: v8, F2: 32767}
                _ = v24
            }
        }
        v3 -= v3
        _ = fi10
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

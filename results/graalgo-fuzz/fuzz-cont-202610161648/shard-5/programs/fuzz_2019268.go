package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 bool }

func main() {
    v1 := []int{460, -206}
    var v2 float64 = float64(1e10)
    v3 := []string{}
    if ((!false) && (0 < 0)) {
        var v4 string = ("a" + "hello")
        var v5 int64 = -289
        _ = v4
        _ = v5
    } else {
        v1 = append(v1, 9223372036854775807)
        _ = len(v1)
        v2 = 0.0
    }
    v2 += float64(1e10)
    _ = len(v1)
    v2 += (v2 * 3.14159)
    var v6 int8 = 100
    if (-845 < -418) {
        for fi7 := 0; fi7 < 9; fi7++ {
            v6 -= (v6 | (^(v6)))
            var v8 bool = false
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: 452, F1: v6}
        if ((!(true || (false && true))) || false) {
            if true {
                for fi10 := 0; fi10 < 11; fi10++ {
                    v6 = (v6 >> ((510) & 63))
                    _ = fi10
                }
                v6 -= (-(v6))
                v2 += 0.0
            }
        } else {
            var v11 float64 = 0.1
            for fi12 := 0; fi12 < 13; fi12++ {
                for fi13 := 0; fi13 < 15; fi13++ {
                    var v14 uint32 = 2147483648
                    fmt.Println(fi13, v2)
                    switch "foo bar" {
                    case "":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case "a":
                        fmt.Println(v2, v11)
                        v15 := GenS1{F0: 65535, F1: v6}
                        _ = v15
                    }
                    _ = fi13
                    _ = v14
                }
                if true {
                    for fi16 := 0; fi16 < 13; fi16++ {
                        var v17 float32 = (-2.25 + (((3.14159 - 0.1) / 0.1) / (((((float32(1.5) - 1.5) - 3.14159) - 3.14159) * 1.5) * -2.25)))
                        switch fi12 {
                        case -629:
                            switch "hello" {
                            case "hello":
                                var v18 float32 = v17
                                if false {
                                    switch "a" {
                                    case "\n\t":
                                        if len(v1) > 0 {
                                            v1[0] = (-(fi12))
                                        }
                                    case "hello":
                                        if true {
                                            var v19 int = fi16
                                            _ = v19
                                        } else {
                                            v17 += 3.14159
                                            fmt.Println(v6, fi16, v17)
                                        }
                                    }
                                    v2 = (1e10 + (v2 + (v11 * v2)))
                                }
                                v17 = float32(v2)
                                _ = v18
                            case "":
                                v2 += v11
                            }
                        }
                        _ = fi16
                        _ = v17
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = fi12
            }
            v20 := GenS1{F0: 65535, F1: (v6 & (v6 & v6))}
            _ = v11
            _ = v20
        }
        _ = v9
    } else {
        if (148 >= -245) {
            v21 := GenS2{F0: false}
            var v22 int16 = -32768
            _ = len(v3)
            _ = v21
            _ = v22
        }
        fmt.Println(v6, v2)
    }
    var v23 uint16 = 2
    v24 := GenS1{F0: (^(v23)), F1: (^(v6))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v6
}

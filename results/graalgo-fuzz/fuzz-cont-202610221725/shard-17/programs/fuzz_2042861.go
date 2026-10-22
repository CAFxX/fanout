package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int32 }

func main() {
    var v1 uint8 = 251
    var v2 uint8 = v1
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]int)
    var v5 int32 = 0
    _ = len(v3)
    fmt.Println(v5, v1)
    fmt.Println(v1, v5, v1)
    _ = len(v4)
    for fi6 := 0; fi6 < 12; fi6++ {
        v2 = v2
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 string = "foo bar"
            v1 -= v2
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 uint = 42
            v5 = (v5 >> ((v10) & 63))
            if (0 >= 42) {
                v11 := GenS1{F0: 2147483647}
                if (2 != 42) {
                    var v12 float64 = ((float64((-2.25 + (1.5 / 3.14159))) - -2.25) + 0.1)
                    switch "foo bar" {
                    case "hello":
                        v13 := GenS1{F0: uint32(v10)}
                        _ = v13
                    case "":
                        v14 := []int{0, (fi9 % ((fi6) | 1))}
                        var v15 int32 = (v5 | v5)
                        v2 = v1
                        fmt.Println(v5, fi6)
                        if len(v14) > 0 {
                            v14[0] = (fi9 % ((2147483647) | 1))
                        }
                        if (32767 < 829) {
                            _ = len(v3)
                            var v16 float64 = float64(1.5)
                            v17 := GenS2{F0: v15, F1: v15}
                            _ = v16
                            _ = v17
                        } else {
                            v1 = v1
                            v18 := v3[0]
                            v1 = v1
                            _ = v18
                        }
                        if len(v14) > 0 {
                            v14[0] = int(v15)
                        }
                        v1 -= 250
                        delete(v4, 42)
                        _ = v14
                        _ = v15
                    case "a":
                        v19 := []int{fi9, 759, 145}
                        _ = v19
                    default:
                        v20 := v3[596]
                        _ = v20
                    }
                    switch ("hello" + (("hello" + ("\n\t" + ("a" + "a"))) + "foo bar")) {
                    case "a":
                        if false {
                            v4[815] = (fi6 & (fi6 | 2147483647))
                        } else {
                            switch 9223372036854775807 {
                            case -545:
                                v5 = (v5 * (-(v5)))
                            }
                            var v21 float64 = float64(float32(v12))
                            v22 := GenS2{F0: v5, F1: (v5 - (v5 - v5))}
                            var v23 uint8 = v2
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                    case "foo bar":
                        if true {
                            var v24 uint64 = 2
                            _ = v24
                        } else {
                            v5 += v5
                            var v25 int32 = -2147483648
                            v25 = (v25 ^ v5)
                            if true {
                                v26 := GenS1{F0: uint32(v2)}
                                v5 -= -2147483648
                                _ = v26
                            } else {
                                var v27 uint8 = (v1 ^ (v2 - (^(v2))))
                                v2 = 255
                                v3[65535] = 4294967295
                                _ = v27
                            }
                            v10 -= 660
                            var v28 float64 = v12
                            v29 := v4[65535]
                            _ = v25
                            _ = v28
                            _ = v29
                        }
                    case "\n\t":
                        switch (fi9 / ((fi9) | 1)) {
                        case 0:
                            v12 = v12
                            v5 += v5
                        }
                    }
                    _ = v12
                }
                if true {
                    fmt.Println(v1, v2, fi9)
                    v30 := v3[567]
                    v1 += (v1 - (v1 % (((^(v2))) | 1)))
                    var v31 int16 = 136
                    if false {
                        v32 := []int{fi6, v30, fi9}
                        var v33 uint16 = 211
                        _ = v32
                        _ = v33
                    }
                    _ = v30
                    _ = v31
                }
                _ = v11
            }
            _ = fi9
            _ = v10
        }
        _ = fi6
    }
    v5 -= v5
    v2 -= 127
    switch -2147483648 {
    case 32767:
        for fi34 := 0; fi34 < 6; fi34++ {
            v5 += (v5 + 2147483647)
            switch ("hello" + ("a" + "\n\t")) {
            case "":
                fmt.Println(v1, v5, v2)
            }
            _ = fi34
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 int16; F1 bool; F2 uint16 }
type GenS4 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64 }
func genFn0() uint8 {
    v7 := GenS1{F0: 1, F1: 0}
    var v8 int = 32768
    v8 += v8
    var v9 int32 = 2147483647
    _ = v7
    _ = v8
    _ = v9
    return 92
}

func main() {
    var v1 uint8 = 159
    var v2 bool = (0 <= 0)
    var v3 float64 = ((0.1 * (float64((1.5 + 3.14159)) - -2.25)) - 0.1)
    var v4 int = 128
    v5 := make(map[int]bool)
    var v6 float32 = float32((float64((0.1 - (0.1 / 1.5))) * 0.1))
    fmt.Println(v4)
    switch (v4 + (v4 ^ 649)) {
    case -1:
        v2 = v2
    }
    fmt.Println(v2, v1)
    v10 := GenS2{F0: 1, F1: ""}
    var v11 float32 = v6
    v12 := GenS3{F0: 746, F1: v2, F2: 42}
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 string = ""
        _ = fi13
        _ = v14
    }
    if false {
        var v15 uint = 42
        v16 := GenS4{F0: 336, F1: 0, F2: 629}
        if (65 > (v1 - 242)) {
            v2 = (!false)
        } else {
            var v17 uint = (v15 * v15)
            var v18 int64 = 0
            for fi19 := 0; fi19 < 19; fi19++ {
                v20 := v5[-353]
                v4 = v4
                for fi21 := 0; fi21 < 7; fi21++ {
                    fmt.Println(v1, v20)
                    if v20 {
                        v22 := GenS5{F0: (v4 - (fi19 * fi19))}
                        v23 := []int{(v4 - fi19)}
                        var v24 string = "\n\t"
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    if (v20 && (0 != -1)) {
                        v2 = true
                    } else {
                        v1 = v1
                        v3 = ((0.1 + ((((float64((-2.25 / float32((v3 + (v3 / float64(v11)))))) - -2.25) / float64((((v6 - (v11 * 0.0)) * float32(v3)) + v6))) * v3) - v3)) / (-2.25 - float64((v6 - v6))))
                        var v25 uint = v17
                        switch (("\n\t" + "foo bar") + "\n\t") {
                        case "foo bar":
                            var v26 int32 = 1
                            _ = v26
                        case "a":
                            fmt.Println(v15)
                        }
                        v17 += v25
                        var v27 uint8 = (^(v1))
                        switch "hello" {
                        case "a":
                            v17 -= (v25 ^ v25)
                        case "\n\t":
                            var v28 uint64 = 42
                            _ = v28
                        default:
                            switch (fi21 + -2147483648) {
                            case 32768:
                                var v29 int = 32767
                                _ = v29
                            default:
                                if ((-(fi21)) <= (fi19 & -807)) {
                                    fmt.Println(fi21)
                                    v6 -= ((v6 - (float32((float64(((v11 - -2.25) + 1e10)) * v3)) * (v6 * 0.1))) + v6)
                                    var v30 uint64 = 4294967295
                                    _ = v30
                                } else {
                                    v31 := []int{fi19, (v4 - fi19)}
                                    _ = v31
                                }
                                v32 := GenS6{F0: (v18 << ((v25) & 63))}
                                _ = v32
                            }
                        }
                        _ = v25
                        _ = v27
                    }
                    _ = fi21
                }
                _ = fi19
                _ = v20
            }
            _ = v17
            _ = v18
        }
        _ = v15
        _ = v16
    }
    if false {
        v33 := []int{260, (v4 >> ((2) & 63)), 192, 0}
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

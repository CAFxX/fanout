package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int; F1 bool }

func main() {
    var v1 bool = (127 != 50)
    var v2 uint8 = 42
    v3 := make(map[string]int)
    v4 := []int{-571, 2147483647, -309}
    v5 := make(map[string]string)
    v3["\n\t"] = 128
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 int16 = 703
        _ = fi6
        _ = v7
    }
    var v8 uint32 = 42
    var v9 int = -2147483648
    var v10 int64 = 0
    v2 = (v2 / ((v2) | 1))
    if ((((0.1 / 1e10) + 1e10) - (((((float32(-2.25) - 3.14159) + (1.5 / 1e10)) + 1.5) * 1.5) - (1.5 / 1.0))) < (1.5 / 0.1)) {
        v8 = uint32(v2)
        fmt.Println(v9)
    } else {
        for fi11 := 0; fi11 < 18; fi11++ {
            for fi12 := 0; fi12 < 15; fi12++ {
                if v1 {
                    v4 = append(v4, fi12)
                    v9 = fi12
                    v8 = v8
                } else {
                    delete(v3, "hello")
                }
                var v13 bool = (-1 < -600)
                if (v10 <= (^(v10))) {
                    v14 := GenS1{F0: 1, F1: 18446744073709551615}
                    if len(v4) > 0 {
                        v4[0] = (^(v9))
                    }
                    var v15 uint8 = v2
                    _ = v14
                    _ = v15
                } else {
                    if v1 {
                        switch ("hello" + "hello") {
                        case "foo bar":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        default:
                            switch 497 {
                            case 2147483647:
                                var v16 int8 = 127
                                _ = v16
                            }
                        }
                        if (!(1 <= 169)) {
                            v3["\n\t"] = (v9 * -671)
                            switch (("" + "a") + "hello") {
                            case "\n\t":
                                fmt.Println(v1)
                            default:
                                v8 = v8
                            }
                            v17 := GenS1{F0: 2, F1: 4294967295}
                            _ = v17
                        }
                        var v18 int64 = -404
                        _ = v18
                    } else {
                        var v19 float32 = float32(3.14159)
                        var v20 int64 = int64(v8)
                        if (v1 && v13) {
                            v21 := GenS1{F0: 881, F1: 4294967295}
                            v10 -= v20
                            _ = v21
                        } else {
                            var v22 float32 = float32(0.0)
                            _ = v22
                        }
                        switch "foo bar" {
                        case "hello":
                            var v23 uint = 258
                            _ = v23
                        case "a":
                            var v24 float64 = (0.0 * (0.1 / ((((-2.25 - 0.1) + float64((v19 * -2.25))) + ((3.14159 - (((1.5 - 0.1) / 0.1) - (1e10 / float64(float32(0.1))))) * (0.0 / 3.14159))) - 1e10)))
                            _ = v24
                        case "":
                            fmt.Println(v1, v13)
                        }
                        v25 := GenS2{F0: v9, F1: false}
                        var v26 float32 = (1.5 / 1e10)
                        if len(v4) > 0 {
                            v4[0] = (v9 / ((v9) | 1))
                        }
                        v8 += 0
                        v13 = ((v2 ^ v2) <= (v2 - v2))
                        _ = v19
                        _ = v20
                        _ = v25
                        _ = v26
                    }
                    var v27 float32 = 0.1
                    _ = v27
                }
                _ = fi12
                _ = v13
            }
            for fi28 := 0; fi28 < 3; fi28++ {
                var v29 bool = (106 >= -517)
                fmt.Println(v8, fi11)
                _ = fi28
                _ = v29
            }
            _ = fi11
        }
        fmt.Println(v2, v10, v1)
    }
    _ = len(v4)
    fmt.Println(v2)
    if v1 {
        fmt.Println(v9)
        var v30 uint32 = (v8 + v8)
        _ = v30
    } else {
        fmt.Println(v2, v8)
        var v31 bool = v1
        v2 = 1
        _ = v31
    }
    v10 += 227
    for fi32 := 0; fi32 < 7; fi32++ {
        v2 = 255
        var v33 string = ("hello" + (((("hello" + ("a" + "hello")) + "a") + ("" + "foo bar")) + "foo bar"))
        v2 += v2
        _ = fi32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}

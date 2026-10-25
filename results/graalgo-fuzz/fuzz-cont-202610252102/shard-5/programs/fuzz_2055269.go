package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint }
type GenS2 struct { F0 bool; F1 bool }

func main() {
    var v1 float64 = (((float64(((1e10 + (1.5 - (0.0 * (3.14159 / (((0.1 + (1.5 + 0.0)) / (0.0 + -2.25)) + float32(1.5)))))) - float32((0.1 + 1e10)))) / ((((-2.25 - 3.14159) * (-2.25 - 3.14159)) + 3.14159) - 3.14159)) + 0.1) * 0.1)
    v2 := make(map[int]string)
    var v3 uint16 = 566
    var v4 int64 = 9223372036854775807
    var v5 float64 = ((v1 * (0.0 / v1)) + v1)
    switch ("\n\t" + "hello") {
    case "":
        fmt.Println(v1, v5)
    case "a":
        v6 := GenS1{F0: 1, F1: 32767, F2: 1}
        _ = v6
    default:
        v7 := []int{32767, -2147483648, -1}
        _ = v7
    }
    var v8 float32 = 1.5
    var v9 int = 9223372036854775807
    var v10 uint16 = v3
    v3 = (^(v10))
    v11 := v2[v9]
    for fi12 := 0; fi12 < 12; fi12++ {
        v10 += (v3 << ((219) & 63))
        v5 = 0.0
        v13 := []int{32767, v9, fi12}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := []int{32768}
        switch ("foo bar" + (("hello" + ("\n\t" + v11)) + "foo bar")) {
        case "a":
            var v16 int = v9
            _ = v16
        case "foo bar":
            for fi17 := 0; fi17 < 18; fi17++ {
                if false {
                    switch v11 {
                    case "\n\t":
                        var v18 float32 = v8
                        v19 := GenS1{F0: 1000, F1: 824, F2: 108}
                        v15 = append(v15, v9)
                        _ = v18
                        _ = v19
                    case "":
                        switch v11 {
                        case "a":
                            if true {
                                v11 = v11
                                if true {
                                    if (!false) {
                                        fmt.Println(v11, v4, v3)
                                        var v20 uint8 = 0
                                        fmt.Println(fi14, v11, v8)
                                        fmt.Println(fi14, fi14)
                                        _ = v20
                                    }
                                    if true {
                                        v3 -= v3
                                    }
                                } else {
                                    v5 = (((v5 * (0.1 + float64(v8))) / v1) * 0.1)
                                    v11 = "hello"
                                }
                            }
                        case "foo bar":
                            var v21 uint = 0
                            _ = v21
                        }
                    }
                    var v22 int64 = v4
                    _ = v22
                } else {
                    switch fi17 {
                    case -462:
                        var v23 uint8 = 115
                        _ = v23
                    case 9223372036854775807:
                        fmt.Println(fi17, v5, fi14)
                    default:
                        var v24 int64 = -1
                        _ = v24
                    }
                }
                _ = fi17
            }
        default:
            var v25 string = v11
            _ = v25
        }
        var v26 uint = 42
        _ = fi14
        _ = v15
        _ = v26
    }
    fmt.Println(v8, v1, v5)
    for fi27 := 0; fi27 < 17; fi27++ {
        fmt.Println(v4)
        fmt.Println(v5, v11)
        if (!(v4 > v4)) {
            v28 := GenS2{F0: (132 >= 935), F1: (255 >= 255)}
            v5 -= 0.0
            var v29 string = (v11 + (v11 + ("" + "a")))
            _ = v28
            _ = v29
        } else {
            fmt.Println(v1, v9)
            var v30 uint32 = 0
            var v31 string = (((v11 + ((("hello" + "\n\t") + "hello") + v11)) + (("foo bar" + "a") + "a")) + v11)
            _ = v30
            _ = v31
        }
        _ = fi27
    }
    switch v11 {
    case "a":
        fmt.Println(v8, v3, v11)
    case "":
        var v32 int32 = -242
        _ = v32
    default:
        v8 += float32(-2.25)
    }
    if (822 <= 65535) {
        v33 := GenS1{F0: 0, F1: 32767, F2: 65535}
        var v34 int = (v9 >> ((1) & 63))
        v10 = (^(v10))
        _ = v33
        _ = v34
    }
    v11 = (v11 + "hello")
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 int16; F2 string }
type GenS3 struct { F0 uint16; F1 int16; F2 uint8 }

func main() {
    v1 := []string{}
    var v2 int64 = -9223372036854775808
    var v3 int64 = 9223372036854775807
    v4 := []string{}
    if len(v4) > 0 {
        v4[0] = ""
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v5 float64 = (1e10 + 1.5)
    if (0 >= 2147483648) {
        v1 = append(v1, ("foo bar" + ("hello" + "foo bar")))
        v6 := GenS1{F0: 0}
        _ = v6
    } else {
        switch ((("foo bar" + "hello") + "a") + "a") {
        case "hello":
            v3 = (v2 % ((v2) | 1))
        default:
            var v7 float32 = 3.14159
            _ = v7
        }
        if len(v4) > 0 {
            v4[0] = ("foo bar" + ("a" + "a"))
        }
        var v8 uint8 = 127
        _ = v8
    }
    v3 -= (v3 - v2)
    var v9 float64 = ((float64(1e10) + v5) + float64(0.1))
    if false {
        switch 9223372036854775807 {
        case 9223372036854775807:
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 uint64 = 180
                _ = fi10
                _ = v11
            }
        }
        switch "foo bar" {
        case "":
            var v12 float32 = ((1e10 / ((1.5 * (-2.25 + ((float32((v5 / (v5 * (v5 - 1e10)))) - -2.25) + (3.14159 / 1.5)))) / 0.1)) + (1.5 / 0.1))
            _ = v12
        default:
            if true {
                var v13 uint = 65535
                var v14 int16 = -706
                if (true && false) {
                    for fi15 := 0; fi15 < 18; fi15++ {
                        switch ("hello" + (("" + "a") + ("a" + (("\n\t" + ("foo bar" + "a")) + "a")))) {
                        case "\n\t":
                            var v16 int32 = 787
                            _ = v16
                        case "foo bar":
                            for fi17 := 0; fi17 < 19; fi17++ {
                                v18 := GenS1{F0: -2147483648}
                                var v19 string = ""
                                if false {
                                    var v20 string = "foo bar"
                                    v20 = ""
                                    v5 -= float64(-2.25)
                                    var v21 int = (-(fi15))
                                    _ = v20
                                    _ = v21
                                }
                                v22 := GenS2{F0: 2, F1: v14, F2: "\n\t"}
                                _ = fi17
                                _ = v18
                                _ = v19
                                _ = v22
                            }
                        }
                        var v23 uint16 = 1
                        _ = fi15
                        _ = v23
                    }
                    var v24 int64 = v3
                    var v25 float64 = (v9 * v9)
                    _ = v24
                    _ = v25
                } else {
                    if len(v1) > 0 {
                        v1[0] = "\n\t"
                    }
                    _ = len(v4)
                    v14 -= (^(v14))
                }
                _ = v13
                _ = v14
            }
        }
        var v26 int32 = 1
        _ = v26
    } else {
        var v27 int8 = 126
        v28 := GenS1{F0: 1}
        switch 0 {
        case 9223372036854775807:
            v29 := GenS1{F0: -1}
            _ = v29
        case 682:
            for fi30 := 0; fi30 < 1; fi30++ {
                v3 = (v2 >> ((885) & 63))
                _ = fi30
            }
        }
        _ = v27
        _ = v28
    }
    v31 := GenS3{F0: 42, F1: -678, F2: 127}
    v5 = 3.14159
    fmt.Println(v5, v2, v3)
    var v32 uint16 = 323
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v9
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 string }
func genFn0() string {
    for fi6 := 0; fi6 < 10; fi6++ {
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 int8 = -52
            v8 = v8
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 int64 = -1
        if false {
            if false {
                for fi11 := 0; fi11 < 18; fi11++ {
                    if false {
                        switch "" {
                        case "":
                            v10 = (v10 << ((0) & 63))
                        default:
                            switch "" {
                            case "foo bar":
                                switch "hello" {
                                case "foo bar":
                                    if (!true) {
                                        v12 := GenS1{F0: 255, F1: 0, F2: ((("a" + "\n\t") + ("foo bar" + ("hello" + ""))) + "foo bar")}
                                        var v13 string = "foo bar"
                                        _ = v12
                                        _ = v13
                                    } else {
                                        var v14 uint16 = 899
                                        _ = v14
                                    }
                                case "hello":
                                    v10 -= 42
                                case "a":
                                    v15 := []int{fi11}
                                    _ = v15
                                }
                            case "":
                                var v16 uint16 = 261
                                _ = v16
                            case "a":
                                v10 += (-(v10))
                            }
                        }
                    }
                    v10 = v10
                    _ = fi11
                }
            }
            for fi17 := 0; fi17 < 3; fi17++ {
                if (65535 == 2) {
                    fmt.Println(v10, fi17)
                    v10 -= (v10 ^ v10)
                }
                var v18 int32 = 52
                _ = fi17
                _ = v18
            }
            if (true || false) {
                v10 -= 1
                var v19 float64 = (-2.25 + float64(1e10))
                v20 := []int{}
                _ = v19
                _ = v20
            } else {
                var v21 int8 = 42
                v22 := GenS1{F0: 128, F1: 42, F2: "hello"}
                _ = v21
                _ = v22
            }
        }
        _ = fi9
        _ = v10
    }
    return ""
}

func main() {
    var v1 int = 127
    var v2 bool = false
    var v3 uint16 = 202
    v4 := []int{v1, (v1 * (-(v1))), 2147483647}
    var v5 int8 = 109
    fmt.Println(v3, v5)
    switch ("" + "") {
    case "\n\t":
        if len(v4) > 0 {
            v4[0] = (v1 * 144)
        }
    }
    var v23 int = (v1 | (v1 >> ((634) & 63)))
    var v24 uint32 = 853
    var v25 int16 = -952
    fmt.Println(v25, v25)
    v4 = append(v4, v23)
    var v26 uint32 = v24
    fmt.Println(v3, v26, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int16 }
func genFn0() uint32 {
    v5 := []int{-636, -985}
    var v6 uint32 = 1
    var v7 uint8 = 2
    _ = v5
    _ = v6
    _ = v7
    return v6
}

func main() {
    v1 := make(map[string]string)
    var v2 bool = (!false)
    var v3 uint16 = 141
    var v4 uint = 1
    v8 := GenS1{F0: "", F1: 42, F2: v3}
    var v9 bool = (-1 == 42)
    var v10 uint = v4
    var v11 uint16 = v3
    fmt.Println(v9)
    for fi12 := 0; fi12 < 5; fi12++ {
        v11 = v11
        v13 := GenS1{F0: ("" + "a"), F1: 1, F2: v11}
        _ = fi12
        _ = v13
    }
    fmt.Println(v10, v10)
    v14 := GenS1{F0: (("a" + (("foo bar" + ("" + ((("" + ("\n\t" + "a")) + "") + "foo bar"))) + "a")) + "hello"), F1: 65535, F2: v11}
    var v15 float32 = (0.0 / 1e10)
    if (!true) {
        var v16 int8 = 42
        for fi17 := 0; fi17 < 5; fi17++ {
            if v9 {
                switch fi17 {
                case 32768:
                    switch fi17 {
                    case 32767:
                        v11 += (v11 ^ (v11 - 32767))
                    case 1:
                        v18 := GenS1{F0: "foo bar", F1: genFn0(), F2: (^(v11))}
                        _ = v18
                    case 127:
                        if ((^(v4)) > 611) {
                            if v9 {
                                v1[("" + "foo bar")] = (("a" + "a") + "foo bar")
                                fmt.Println(v4)
                                for fi19 := 0; fi19 < 17; fi19++ {
                                    v1["foo bar"] = "hello"
                                    if (v3 < (v11 - v11)) {
                                        if (false && ((((fi19 << (((v10 + v10)) & 63)) != (fi17 * (fi19 >> (((v4 & v10)) & 63)))) || ((v16 & (^(v16))) != v16)) && true)) {
                                            var v20 float32 = ((v15 * (-2.25 * (0.0 * v15))) / float32((float64(float32((-2.25 - 1e10))) + 1.5)))
                                            var v21 uint8 = 1
                                            v22 := GenS1{F0: "foo bar", F1: genFn0(), F2: v3}
                                            _ = v20
                                            _ = v21
                                            _ = v22
                                        }
                                        if (("" + "\n\t") == "foo bar") {
                                            var v23 uint8 = 255
                                            v24 := []int{}
                                            v16 = -128
                                            _ = v23
                                            _ = v24
                                        }
                                        var v25 uint32 = 65535
                                        _ = v25
                                    } else {
                                        v26 := GenS2{F0: v11, F1: 0}
                                        v27 := []int{fi17}
                                        _ = v26
                                        _ = v27
                                    }
                                    _ = fi19
                                }
                            } else {
                                v1["foo bar"] = ""
                                fmt.Println(fi17, v16)
                            }
                        }
                    }
                case -301:
                    fmt.Println(fi17, v3, v3)
                default:
                    v1["a"] = (("hello" + "a") + ("\n\t" + ""))
                }
                var v28 bool = (409 >= -32768)
                _ = v28
            }
            for fi29 := 0; fi29 < 17; fi29++ {
                var v30 int16 = -967
                _ = fi29
                _ = v30
            }
            _ = fi17
        }
        _ = v16
    } else {
        fmt.Println(v2)
        var v31 float32 = 1e10
        _ = v31
    }
    var v32 uint8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v8
    _ = v9
}

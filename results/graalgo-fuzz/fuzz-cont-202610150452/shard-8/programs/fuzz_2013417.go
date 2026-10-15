package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 uint; F1 int16 }

func main() {
    v1 := make(map[string]string)
    v2 := []string{}
    var v3 uint = 870
    var v4 int8 = 100
    v5 := []int{-952}
    var v6 float32 = float32(3.14159)
    _ = len(v1)
    var v7 int64 = 42
    v7 -= v7
    var v8 int = 0
    if (0 == 255) {
        v8 += v8
        var v9 int = 32767
        for fi10 := 0; fi10 < 6; fi10++ {
            v8 -= (v9 ^ 32768)
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi10
        }
        _ = v9
    } else {
        var v11 float64 = 0.1
        delete(v1, "a")
        _ = v11
    }
    if true {
        if (v8 != (v8 & (v8 >> ((255) & 63)))) {
            v1[(("hello" + "foo bar") + "\n\t")] = (("\n\t" + ((("a" + "foo bar") + "foo bar") + "a")) + ("hello" + ("a" + "\n\t")))
            switch "\n\t" {
            case "a":
                v1["hello"] = "a"
            default:
                if false {
                    var v12 string = ((("foo bar" + "hello") + ("foo bar" + "hello")) + "foo bar")
                    if true {
                        for fi13 := 0; fi13 < 10; fi13++ {
                            v1[(v12 + v12)] = v12
                            if true {
                                if (42 >= 0) {
                                    for fi14 := 0; fi14 < 14; fi14++ {
                                        fmt.Println(v6)
                                        _ = fi14
                                    }
                                    switch (((v12 + ("hello" + (((v12 + "\n\t") + "hello") + "hello"))) + (("foo bar" + ("a" + v12)) + "a")) + v12) {
                                    case "a":
                                        switch v12 {
                                        case "foo bar":
                                            v15 := GenS1{F0: 1}
                                            _ = v15
                                        case "":
                                            v6 -= ((v6 / ((-2.25 - 3.14159) * (-2.25 * 1e10))) * v6)
                                        case "\n\t":
                                            v4 += (v4 << ((v3) & 63))
                                        default:
                                            v12 = "a"
                                        }
                                    case "foo bar":
                                        if len(v2) > 0 {
                                            v2[0] = v12
                                        }
                                    case "":
                                        var v16 int = 42
                                        _ = v16
                                    }
                                } else {
                                    v5 = append(v5, (-(v8)))
                                    v17 := GenS2{F0: (v7 + (v7 / ((v7) | 1))), F1: -1}
                                    delete(v1, "\n\t")
                                    _ = v17
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                if len(v5) > 0 {
                                    v5[0] = (fi13 + (^(v8)))
                                }
                            }
                            _ = fi13
                        }
                    }
                    _ = v12
                } else {
                    for fi18 := 0; fi18 < 15; fi18++ {
                        var v19 float32 = ((((v6 - 0.1) * ((0.0 - 1.5) + v6)) - (v6 + (((1.5 - 1e10) + v6) + (v6 - 0.0)))) / (3.14159 + -2.25))
                        _ = fi18
                        _ = v19
                    }
                    var v20 bool = false
                    var v21 float64 = 0.0
                    _ = v20
                    _ = v21
                }
            }
        } else {
            fmt.Println(v8)
            for fi22 := 0; fi22 < 15; fi22++ {
                var v23 float32 = (0.1 - ((-2.25 + v6) - float32(0.1)))
                if (!(true || false)) {
                    v24 := GenS3{F0: v3, F1: 42}
                    switch "a" {
                    case "":
                        v25 := GenS2{F0: -1, F1: -388}
                        _ = v25
                    case "foo bar":
                        _ = len(v2)
                    default:
                        for fi26 := 0; fi26 < 2; fi26++ {
                            var v27 int64 = (v7 + int64(v8))
                            v28 := GenS4{F0: (v3 << (((v3 | (v3 / ((v3) | 1)))) & 63)), F1: 1000}
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                    }
                    _ = v24
                }
                for fi29 := 0; fi29 < 19; fi29++ {
                    v8 += (fi29 | (^(fi22)))
                    _ = fi29
                }
                _ = fi22
                _ = v23
            }
            var v30 bool = (((false || (((v7 % ((1) | 1)) <= v7) || (true || (!true)))) || false) && true)
            _ = v30
        }
        v31 := GenS1{F0: 127}
        _ = v31
    }
    v32 := GenS3{F0: (v3 / ((v3) | 1)), F1: 2}
    var v33 int16 = 435
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

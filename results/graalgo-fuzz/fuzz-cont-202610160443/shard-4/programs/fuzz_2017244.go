package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int16; F1 int32; F2 uint }

func main() {
    v1 := []string{"\n\t", "a", "foo bar"}
    var v2 int16 = -659
    var v3 bool = false
    var v4 int64 = 957
    v5 := []int{4294967295, 1}
    v6 := make(map[int]int)
    switch 2147483647 {
    case 0:
        for fi7 := 0; fi7 < 12; fi7++ {
            v3 = ("foo bar" > "a")
            switch "hello" {
            case "foo bar":
                v8 := GenS1{F0: 42, F1: "\n\t"}
                _ = v8
            }
            _ = fi7
        }
    case -9223372036854775808:
        fmt.Println(v3)
    }
    if (v3 && true) {
        var v9 int32 = -172
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        v3 = v3
        _ = fi10
    }
    var v11 int16 = 1
    var v12 int16 = (v2 ^ v2)
    var v13 float64 = 1.5
    fmt.Println(v2)
    v13 = v13
    v14 := GenS1{F0: 1, F1: ("hello" + "a")}
    if v3 {
        var v15 int16 = -996
        v16 := GenS2{F0: -32768, F1: -417, F2: 255}
        v12 -= 32767
        _ = v15
        _ = v16
    }
    v13 -= (float64(1e10) * (v13 / v13))
    for fi17 := 0; fi17 < 6; fi17++ {
        if v3 {
            fmt.Println(fi17, v12, v2)
            if true {
                v3 = (!true)
                var v18 int = -561
                _ = v18
            } else {
                v3 = (v2 != v11)
                if true {
                    if true {
                        v19 := v6[-355]
                        _ = v19
                    } else {
                        v4 -= 42
                    }
                    switch fi17 {
                    case -2147483648:
                        v20 := GenS1{F0: -1, F1: ("hello" + "a")}
                        _ = v20
                    case 2147483647:
                        for fi21 := 0; fi21 < 10; fi21++ {
                            _ = len(v6)
                            if ((v11 * 32767) > v12) {
                                _ = len(v5)
                                var v22 bool = (-2147483648 == 0)
                                _ = v22
                            }
                            _ = fi21
                        }
                    case 127:
                        v11 = v12
                    default:
                        if (v3 && v3) {
                            if v3 {
                                for fi23 := 0; fi23 < 9; fi23++ {
                                    var v24 uint = 153
                                    switch ((("a" + "foo bar") + "") + ("foo bar" + ("foo bar" + "foo bar"))) {
                                    case "hello":
                                        v24 -= 1
                                    case "foo bar":
                                        var v25 uint16 = 32767
                                        _ = v25
                                    default:
                                        fmt.Println(v12, v13, v3)
                                    }
                                    var v26 int32 = 2147483647
                                    _ = fi23
                                    _ = v24
                                    _ = v26
                                }
                                v2 -= (v2 - (-(v12)))
                            }
                            delete(v6, (fi17 | fi17))
                            var v27 uint32 = 586
                            _ = v27
                        }
                    }
                }
            }
        }
        var v28 float32 = (3.14159 + 1e10)
        _ = fi17
        _ = v28
    }
    v13 -= 0.1
    for fi29 := 0; fi29 < 16; fi29++ {
        for fi30 := 0; fi30 < 11; fi30++ {
            delete(v6, 990)
            v31 := GenS1{F0: 127, F1: ("hello" + "a")}
            _ = fi30
            _ = v31
        }
        _ = fi29
    }
    v1 = append(v1, ("foo bar" + "hello"))
    _ = len(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

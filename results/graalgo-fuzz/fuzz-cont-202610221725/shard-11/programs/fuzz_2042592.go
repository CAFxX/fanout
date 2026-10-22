package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]uint32)
    v3 := make(map[string]string)
    var v4 int16 = -295
    var v5 float64 = 1.5
    if true {
        v6 := []int{1}
        var v7 int = 2147483647
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 uint = 627
            _ = fi8
            _ = v9
        }
        _ = v6
        _ = v7
    } else {
        for fi10 := 0; fi10 < 10; fi10++ {
            v4 -= -1
            _ = fi10
        }
        if false {
            if (!false) {
                switch 854 {
                case -782:
                    var v11 bool = true
                    _ = v11
                }
            }
        } else {
            var v12 float32 = float32(3.14159)
            _ = v12
        }
    }
    v13 := []int{127}
    var v14 int = -679
    v4 += v4
    for fi15 := 0; fi15 < 12; fi15++ {
        for fi16 := 0; fi16 < 3; fi16++ {
            if true {
                var v17 bool = (("\n\t" + ("\n\t" + ("foo bar" + "foo bar"))) >= "a")
                if v17 {
                    if v17 {
                        v5 = float64(1e10)
                        v4 -= (v4 - 1000)
                        if false {
                            var v18 uint8 = 1
                            v19 := v1[(fi15 | -287)]
                            var v20 uint64 = 0
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                        fmt.Println(v17, v17)
                    } else {
                        v14 = (fi16 & fi15)
                    }
                    if len(v13) > 0 {
                        v13[0] = v14
                    }
                    var v21 string = "foo bar"
                    _ = v21
                } else {
                    if false {
                        v4 = v4
                        fmt.Println(v17)
                        if (fi15 > 32768) {
                            if false {
                                var v22 float32 = float32(v5)
                                v23 := GenS1{F0: 65535, F1: 346}
                                var v24 uint32 = 2
                                _ = v22
                                _ = v23
                                _ = v24
                            } else {
                                switch "a" {
                                case "hello":
                                    if v17 {
                                        var v25 uint8 = 226
                                        fmt.Println(v25, v14, v4)
                                        v14 += 0
                                        _ = v25
                                    }
                                }
                                var v26 int16 = (v4 / ((0) | 1))
                                _ = v26
                            }
                            if false {
                                fmt.Println(v17, v4)
                                var v27 uint32 = 2147483647
                                v28 := GenS1{F0: 255, F1: 811}
                                _ = v27
                                _ = v28
                            }
                        } else {
                            if true {
                                var v29 uint = 2
                                var v30 int = v14
                                _ = v29
                                _ = v30
                            }
                            if (127 >= 161) {
                                delete(v3, "")
                            }
                        }
                    }
                }
                delete(v2, v14)
                _ = v17
            }
            v13 = append(v13, fi15)
            _ = fi16
        }
        switch fi15 {
        case 0:
            fmt.Println(v4, v14, v4)
        }
        _ = fi15
    }
    for fi31 := 0; fi31 < 1; fi31++ {
        v4 += v4
        if ((0 == -368) || (-2147483648 == 2147483647)) {
            v5 -= ((3.14159 * (v5 * 0.0)) - 1.5)
        }
        _ = fi31
    }
    var v32 int32 = -198
    v13 = append(v13, v14)
    v4 = (v4 ^ 42)
    var v33 bool = ((-9223372036854775808 >= v14) && true)
    for fi34 := 0; fi34 < 4; fi34++ {
        v35 := GenS1{F0: 2, F1: 4294967295}
        _ = fi34
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}

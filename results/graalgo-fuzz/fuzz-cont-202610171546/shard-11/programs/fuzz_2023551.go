package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 bool; F1 uint8 }

func main() {
    var v1 uint16 = 42
    v2 := make(map[string]bool)
    v3 := []string{(("a" + "foo bar") + ("" + (("\n\t" + "foo bar") + "hello"))), "foo bar"}
    v4 := []int{}
    v5 := make(map[int]int)
    for fi6 := 0; fi6 < 14; fi6++ {
        if true {
            var v7 string = ("\n\t" + "foo bar")
            _ = v7
        }
        if (1 >= v1) {
            for fi8 := 0; fi8 < 9; fi8++ {
                var v9 int32 = -2147483648
                if false {
                    v9 = v9
                } else {
                    if true {
                        _ = len(v5)
                        if true {
                            var v10 uint8 = 0
                            v4 = append(v4, fi6)
                            v11 := GenS1{F0: "foo bar", F1: -3}
                            _ = v10
                            _ = v11
                        } else {
                            if len(v4) > 0 {
                                v4[0] = -9223372036854775808
                            }
                            _ = len(v3)
                        }
                        v9 = v9
                    }
                }
                fmt.Println(v1, v9, fi8)
                _ = fi8
                _ = v9
            }
            fmt.Println(fi6, v1)
            v12 := GenS1{F0: "\n\t", F1: 100}
            _ = v12
        } else {
            fmt.Println(fi6, fi6)
            fmt.Println(v1, v1)
        }
        _ = fi6
    }
    var v13 uint = 946
    if false {
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        v14 := GenS2{F0: true, F1: 127}
        _ = v14
    } else {
        v1 -= (v1 ^ v1)
        var v15 string = "a"
        if false {
            v4 = append(v4, 2147483647)
            _ = len(v2)
        }
        _ = v15
    }
    var v16 uint = (^(v13))
    v16 = (^(v16))
    var v17 int = -1
    v1 -= v1
    for fi18 := 0; fi18 < 12; fi18++ {
        var v19 int64 = 957
        var v20 uint32 = 42
        _ = fi18
        _ = v19
        _ = v20
    }
    if true {
        v1 -= (^(v1))
        if true {
            var v21 uint = 42
            _ = v21
        } else {
            var v22 bool = false
            if (168 == 21) {
                for fi23 := 0; fi23 < 14; fi23++ {
                    if ((1.5 < (0.1 / float64(0.1))) || v22) {
                        v24 := GenS2{F0: v22, F1: 1}
                        _ = v24
                    }
                    v13 = 4294967295
                    fmt.Println(v17, v22, v13)
                    _ = fi23
                }
                for fi25 := 0; fi25 < 11; fi25++ {
                    if v22 {
                        if v22 {
                            var v26 int32 = -908
                            _ = v26
                        }
                        fmt.Println(v17, v13)
                    } else {
                        var v27 float32 = -2.25
                        _ = v27
                    }
                    _ = fi25
                }
            }
            _ = v22
        }
        fmt.Println(v13)
    }
    var v28 uint = 4294967295
    v16 = (^(v28))
    fmt.Println(v17, v28, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}

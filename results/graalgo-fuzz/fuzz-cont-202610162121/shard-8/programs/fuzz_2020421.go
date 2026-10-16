package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int64 }

func main() {
    var v1 int8 = 0
    var v2 uint16 = 65535
    v3 := make(map[int]bool)
    var v4 float64 = ((-2.25 - 1.5) / (0.0 + 1e10))
    v5 := make(map[int]bool)
    var v6 uint8 = 205
    switch "a" {
    case "foo bar":
        v7 := []int{520, 128, 0}
        _ = v7
    }
    v4 -= (0.1 + 1.5)
    v8 := GenS1{F0: -391, F1: ((((false || (0 <= 113)) || true) || true) && false)}
    v6 += (v6 - 13)
    v4 = (v4 * v4)
    v4 += (v4 / 1.5)
    var v9 string = "hello"
    switch "\n\t" {
    case "hello":
        for fi10 := 0; fi10 < 2; fi10++ {
            if ("\n\t" != v9) {
                v1 -= (v1 << ((2) & 63))
                v6 -= 127
            } else {
                v4 = v4
            }
            _ = fi10
        }
    default:
        if (!false) {
            for fi11 := 0; fi11 < 15; fi11++ {
                switch fi11 {
                case -789:
                    if (false && false) {
                        fmt.Println(v2, v9, v2)
                        v9 = v9
                    } else {
                        if false {
                            v3[2147483647] = (1 == -1)
                            var v12 int32 = 1
                            _ = v12
                        }
                        v1 = v1
                        fmt.Println(v1, v9, v1)
                    }
                case 127:
                    var v13 uint8 = v6
                    _ = v13
                }
                _ = fi11
            }
            v14 := GenS1{F0: -576, F1: false}
            _ = v14
        }
    }
    v6 = (v6 | v6)
    if true {
        v15 := []int{-485, 191, 4294967295, 127}
        switch v9 {
        case "":
            var v16 int8 = v1
            _ = v16
        case "hello":
            var v17 float32 = -2.25
            _ = v17
        }
        if false {
            v2 -= 42
        } else {
            for fi18 := 0; fi18 < 5; fi18++ {
                var v19 uint16 = (^(v2))
                v20 := GenS1{F0: -117, F1: ("\n\t" >= v9)}
                if false {
                    delete(v5, -9223372036854775808)
                    v6 = v6
                    switch v9 {
                    case "":
                        for fi21 := 0; fi21 < 15; fi21++ {
                            var v22 float32 = 1e10
                            _ = fi21
                            _ = v22
                        }
                    case "foo bar":
                        var v23 int16 = 641
                        _ = v23
                    case "hello":
                        v1 -= (-(v1))
                    default:
                        for fi24 := 0; fi24 < 17; fi24++ {
                            if (v9 <= v9) {
                                switch v9 {
                                case "hello":
                                    v1 = v1
                                case "a":
                                    switch "hello" {
                                    case "":
                                        fmt.Println(v2)
                                    case "\n\t":
                                        _ = len(v3)
                                    }
                                }
                                var v25 uint = 1
                                switch (v9 + v9) {
                                case "a":
                                    fmt.Println(v4, v9, v6)
                                case "foo bar":
                                    v26 := v3[fi24]
                                    _ = v26
                                case "":
                                    var v27 uint8 = (v6 + 127)
                                    _ = v27
                                default:
                                    v28 := GenS1{F0: fi24, F1: true}
                                    _ = v28
                                }
                                _ = v25
                            }
                            var v29 int8 = (v1 + (-(v1)))
                            _ = fi24
                            _ = v29
                        }
                    }
                } else {
                    v2 = 2
                }
                _ = fi18
                _ = v19
                _ = v20
            }
            var v30 int32 = 2147483647
            v31 := GenS2{F0: 0}
            _ = v30
            _ = v31
        }
        _ = v15
    } else {
        delete(v3, -84)
        switch v9 {
        case "hello":
            if (!(!true)) {
                v32 := GenS1{F0: 127, F1: (-1 > -1)}
                _ = v32
            }
        case "a":
            v9 = "\n\t"
        }
    }
    for fi33 := 0; fi33 < 2; fi33++ {
        var v34 float32 = ((1e10 / (0.0 + (((3.14159 / 1.0) - 0.0) + 3.14159))) - 0.0)
        v9 = "a"
        _ = fi33
        _ = v34
    }
    var v35 float32 = (0.1 + 3.14159)
    if true {
        switch -716 {
        case 4294967295:
            var v36 int32 = 526
            _ = v36
        default:
            v6 -= v6
        }
        fmt.Println(v6)
        var v37 uint16 = v2
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}

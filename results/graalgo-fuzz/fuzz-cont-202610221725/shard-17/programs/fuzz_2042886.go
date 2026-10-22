package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }

func main() {
    v1 := make(map[int]bool)
    v2 := []int{4294967295, 1, -1}
    var v3 float64 = 0.1
    var v4 bool = (125 >= 0)
    var v5 int8 = -67
    v6 := []int{-9223372036854775808, 236, -759}
    v3 += (v3 / 1.5)
    v5 -= 0
    v7 := GenS1{F0: 252, F1: -568}
    if len(v2) > 0 {
        v2[0] = 32767
    }
    v8 := GenS1{F0: 42, F1: 127}
    var v9 float32 = 3.14159
    v3 += (v3 * 1e10)
    v3 -= v3
    switch ("foo bar" + ("" + "foo bar")) {
    case "":
        if ((v3 * v3) >= v3) {
            v10 := GenS1{F0: -66, F1: 2147483647}
            if (836 != 4294967295) {
                fmt.Println(v3)
                v4 = false
                var v11 string = ("a" + "hello")
                _ = v11
            }
            switch "\n\t" {
            case "a":
                for fi12 := 0; fi12 < 17; fi12++ {
                    v9 += 3.14159
                    var v13 string = ""
                    for fi14 := 0; fi14 < 19; fi14++ {
                        v9 += ((v9 / ((v9 - v9) / float32(((-2.25 / (0.1 * (v3 * v3))) - 0.1)))) + v9)
                        _ = fi14
                    }
                    _ = fi12
                    _ = v13
                }
            default:
                for fi15 := 0; fi15 < 15; fi15++ {
                    for fi16 := 0; fi16 < 10; fi16++ {
                        v9 -= (3.14159 / (v9 * v9))
                        _ = fi16
                    }
                    _ = fi15
                }
            }
            _ = v10
        } else {
            v9 = float32(v3)
            var v17 uint = 0
            switch -2147483648 {
            case -527:
                if false {
                    var v18 int64 = -816
                    var v19 uint64 = 4294967295
                    _ = v18
                    _ = v19
                } else {
                    switch "a" {
                    case "":
                        fmt.Println(v5, v3)
                    case "hello":
                        if false {
                            if (1 <= -466) {
                                for fi20 := 0; fi20 < 15; fi20++ {
                                    fmt.Println(v3, v4)
                                    _ = fi20
                                }
                            } else {
                                var v21 uint = (v17 - v17)
                                _ = v21
                            }
                            for fi22 := 0; fi22 < 12; fi22++ {
                                if (false || v4) {
                                    delete(v1, int(v17))
                                } else {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                }
                                _ = len(v2)
                                _ = fi22
                            }
                            if len(v6) > 0 {
                                v6[0] = 32768
                            }
                        }
                    case "\n\t":
                        var v23 uint16 = 32767
                        _ = v23
                    }
                }
            case 481:
                v4 = false
            }
            _ = v17
        }
    case "hello":
        v3 -= 1.5
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

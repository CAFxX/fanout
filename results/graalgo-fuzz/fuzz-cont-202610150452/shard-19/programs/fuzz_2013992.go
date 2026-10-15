package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 string }
type GenS2 struct { F0 uint32; F1 bool; F2 int64 }
func genFn0() int64 {
    var v7 string = "a"
    for fi8 := 0; fi8 < 17; fi8++ {
        v7 = "\n\t"
        switch v7 {
        case "\n\t":
            v7 = ""
        case "foo bar":
            if true {
                if true {
                    v7 = (("" + v7) + ((("hello" + (v7 + v7)) + "a") + ("a" + ("hello" + ((v7 + ((v7 + ("foo bar" + v7)) + v7)) + (v7 + ((v7 + ("foo bar" + (v7 + ""))) + "foo bar")))))))
                }
                if false {
                    v9 := GenS1{F0: 42, F1: fi8, F2: ""}
                    v10 := []int{fi8}
                    v7 = "a"
                    _ = v10
                    _ = v9
                }
            } else {
                for fi11 := 0; fi11 < 5; fi11++ {
                    v12 := []int{fi11, 2147483647}
                    _ = fi11
                    _ = v12
                }
            }
        case "":
            if false {
                fmt.Println(v7, fi8)
            } else {
                v7 = ("hello" + "a")
                fmt.Println(v7)
            }
        }
        v7 = "foo bar"
        _ = fi8
    }
    v7 = "a"
    fmt.Println(v7)
    _ = v7
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = (0.1 + float32((1.5 - (3.14159 * (1e10 / 1.5)))))
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]string)
    var v5 uint64 = 4294967295
    var v6 uint8 = 1
    var v13 int = 32768
    if false {
        for fi14 := 0; fi14 < 11; fi14++ {
            v13 += fi14
            v15 := v4[170]
            _ = fi14
            _ = v15
        }
        fmt.Println(v6, v2)
    } else {
        if ((v5 | (v5 - 1)) != v5) {
            v13 += (v13 - v13)
        } else {
            v16 := GenS1{F0: 1, F1: 200, F2: "\n\t"}
            fmt.Println(v13, v6)
            fmt.Println(v5)
            _ = v16
        }
        v5 -= 42
    }
    var v17 float64 = 1e10
    var v18 int32 = -1
    if (0.1 != (v2 - float32(v17))) {
        v19 := GenS2{F0: 42, F1: true, F2: -727}
        v2 -= ((3.14159 + (((float32((v17 / ((v17 / (float64((((-2.25 + 1e10) + v2) / v2)) + float64((1.5 / ((-2.25 + v2) / v2))))) + v17))) / v2) - v2) - v2)) + v2)
        v20 := GenS1{F0: 4294967295, F1: v13, F2: "foo bar"}
        _ = v19
        _ = v20
    } else {
        for fi21 := 0; fi21 < 5; fi21++ {
            _ = len(v4)
            for fi22 := 0; fi22 < 4; fi22++ {
                v23 := GenS1{F0: 2147483647, F1: (fi22 * 935), F2: "a"}
                v13 = v13
                if true {
                    var v24 int8 = -31
                    _ = v24
                } else {
                    switch "\n\t" {
                    case "a":
                        v13 += v13
                    }
                    v13 = v13
                }
                var v25 uint8 = uint8(v5)
                v26 := []int{(v13 | fi21), (v13 * (fi21 - 42))}
                switch (("" + "") + (("" + ((("" + "hello") + ("foo bar" + "")) + "a")) + ("\n\t" + "hello"))) {
                case "hello":
                    if len(v26) >= 2 {
                        v26 = v26[1:len(v26)]
                    }
                default:
                    fmt.Println(v5)
                }
                _ = fi22
                _ = v23
                _ = v25
                _ = v26
            }
            var v27 string = "a"
            _ = fi21
            _ = v27
        }
        fmt.Println(v17, v2, v6)
        v6 = (v6 | 203)
    }
    switch "hello" {
    case "hello":
        fmt.Println(v2, v13)
    }
    var v28 float32 = (0.0 * v2)
    switch (v13 * (v13 / ((v13) | 1))) {
    case 2147483647:
        v29 := []int{}
        _ = v29
    case 9223372036854775807:
        _ = len(v4)
    case 0:
        for fi30 := 0; fi30 < 16; fi30++ {
            if true {
                fmt.Println(fi30)
                var v31 int16 = 0
                _ = v31
            } else {
                v17 = -2.25
                fmt.Println(v2, v6, fi30)
                if (!false) {
                    v32 := GenS1{F0: 1, F1: (v13 / (((fi30 ^ (fi30 & fi30))) | 1)), F2: ""}
                    fmt.Println(fi30, v2, fi30)
                    for fi33 := 0; fi33 < 18; fi33++ {
                        if (!true) {
                            _ = len(v3)
                        } else {
                            var v34 bool = true
                            if (("" + "foo bar") <= "") {
                                v4[(fi30 << ((18446744073709551615) & 63))] = "a"
                            } else {
                                fmt.Println(fi30, v28)
                                v5 -= 4294967295
                                v34 = true
                            }
                            v13 = fi30
                            v13 -= (^(fi33))
                            _ = v34
                        }
                        _ = fi33
                    }
                    _ = v32
                }
            }
            _ = fi30
        }
    default:
        v18 += (v18 + (v18 | v18))
    }
    v28 += v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

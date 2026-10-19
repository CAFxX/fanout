package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int }
func genFn0() uint16 {
    var v5 string = ""
    v6 := []int{4294967295, -376}
    if ((((true || (false && false)) && true) && false) || (4294967295 < 0)) {
        v7 := GenS1{F0: 1}
        for fi8 := 0; fi8 < 10; fi8++ {
            fmt.Println(v5, v5)
            _ = fi8
        }
        _ = v7
    } else {
        v5 = v5
        switch -2147483648 {
        case 4294967295:
            var v9 uint32 = 2
            _ = v9
        default:
            if len(v6) > 0 {
                v6[0] = 32767
            }
        }
        if ((0.0 != float64(float32(-2.25))) || false) {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            var v10 uint32 = 2
            _ = v10
        } else {
            var v11 int64 = 0
            _ = v11
        }
    }
    _ = v5
    _ = v6
    return 748
}
func genFn1() int32 {
    var v12 float32 = -2.25
    if (-73 < -7) {
        var v13 uint8 = 255
        for fi14 := 0; fi14 < 5; fi14++ {
            fmt.Println(v12, v12)
            fmt.Println(v13, v13, v13)
            for fi15 := 0; fi15 < 15; fi15++ {
                fmt.Println(fi15)
                switch "" {
                case "a":
                    fmt.Println(fi14)
                case "foo bar":
                    var v16 uint64 = 972
                    _ = v16
                }
                var v17 uint16 = 255
                _ = fi15
                _ = v17
            }
            _ = fi14
        }
        v13 += 2
        _ = v13
    }
    var v18 float64 = 1.5
    _ = v12
    _ = v18
    return 42
}

func main() {
    var v1 uint64 = 1
    var v2 int64 = -259
    var v3 int8 = -24
    v4 := make(map[string]int)
    v19 := []int{2147483647}
    v20 := GenS1{F0: 262}
    fmt.Println(v3)
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    fmt.Println(v2)
    fmt.Println(v1, v2, v2)
    fmt.Println(v3)
    for fi21 := 0; fi21 < 12; fi21++ {
        v2 += 1
        for fi22 := 0; fi22 < 1; fi22++ {
            var v23 int8 = (v3 ^ (v3 - v3))
            if false {
                switch "\n\t" {
                case "hello":
                    if (47 < 2) {
                        var v24 bool = true
                        _ = v24
                    }
                default:
                    v25 := GenS1{F0: 2}
                    _ = v25
                }
            }
            _ = fi22
            _ = v23
        }
        v4[""] = 2147483647
        _ = fi21
    }
    if ((!false) && ((-2147483648 == -238) || (true || (2 != genFn0())))) {
        var v26 int = 9223372036854775807
        for fi27 := 0; fi27 < 11; fi27++ {
            var v28 bool = ("a" <= ("\n\t" + ("a" + "")))
            if true {
                v2 += v2
            }
            var v29 string = ""
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = v26
    }
    v30 := GenS2{F0: 0, F1: v1, F2: -1}
    fmt.Println(v3)
    v31 := GenS1{F0: 339}
    for fi32 := 0; fi32 < 4; fi32++ {
        var v33 uint = 255
        _ = fi32
        _ = v33
    }
    var v34 bool = false
    var v35 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v4
}

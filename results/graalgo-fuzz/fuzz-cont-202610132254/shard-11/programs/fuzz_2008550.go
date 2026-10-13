package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint8 }
func genFn0() int16 {
    var v7 int64 = 9223372036854775807
    var v8 int16 = 1000
    switch 128 {
    case 9223372036854775807:
        v9 := GenS1{F0: 439, F1: 127, F2: 670}
        _ = v9
    }
    v10 := []int{2147483647, 9223372036854775807}
    _ = v10
    _ = v7
    _ = v8
    return 0
}
func genFn1() int8 {
    var v11 int8 = 42
    if (!false) {
        v11 -= v11
    } else {
        if true {
            v12 := []int{-850, -1, -142}
            fmt.Println(v11)
            v11 = 0
            _ = v12
        }
        v11 = (v11 + 1)
    }
    var v13 string = "\n\t"
    _ = v11
    _ = v13
    return -128
}

func main() {
    v1 := []uint32{1, 800, 428, 2}
    var v2 uint64 = 0
    v3 := make(map[string]string)
    var v4 int = 872
    v5 := make(map[uint32]int)
    v6 := make(map[int]bool)
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 uint64 = v2
        for fi16 := 0; fi16 < 10; fi16++ {
            fmt.Println(fi14, v4)
            if (((0.0 - 0.1) > (3.14159 * ((float32(-2.25) * 3.14159) + 1e10))) && false) {
                var v17 uint8 = 202
                var v18 int = 727
                _ = v17
                _ = v18
            }
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    switch "\n\t" {
    case "hello":
        var v19 int8 = 1
        _ = v19
    default:
        fmt.Println(v2, v4)
    }
    var v20 int64 = -847
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 uint8 = 128
        v5[289] = (fi21 * (^(v4)))
        _ = fi21
        _ = v22
    }
    _ = len(v1)
    if true {
        v23 := GenS2{F0: 128}
        var v24 uint16 = 0
        if (!(true && true)) {
            v24 = v24
            for fi25 := 0; fi25 < 8; fi25++ {
                fmt.Println(fi25)
                _ = fi25
            }
            for fi26 := 0; fi26 < 17; fi26++ {
                if (("a" >= "a") || ((v4 / (((-(fi26))) | 1)) != 32768)) {
                    switch "hello" {
                    case "a":
                        if len(v1) > 0 {
                            v1[0] = 2
                        }
                    case "foo bar":
                        delete(v6, v4)
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                }
                var v27 float64 = (float64((((0.1 * (1e10 - (1e10 / 0.1))) + 3.14159) * 0.1)) - 3.14159)
                for fi28 := 0; fi28 < 7; fi28++ {
                    v4 = (v4 - (-(v4)))
                    _ = len(v5)
                    _ = fi28
                }
                _ = fi26
                _ = v27
            }
        }
        _ = v23
        _ = v24
    } else {
        switch (("hello" + ("" + ("a" + "\n\t"))) + "") {
        case "":
            fmt.Println(v4)
        case "hello":
            v29 := v3[(("" + ("a" + "hello")) + ("\n\t" + ""))]
            _ = v29
        }
        for fi30 := 0; fi30 < 4; fi30++ {
            v1 = append(v1, 2)
            var v31 uint8 = 146
            _ = fi30
            _ = v31
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (!(!true)) {
        var v32 uint = 668
        var v33 string = "foo bar"
        _ = v32
        _ = v33
    } else {
        v34 := GenS1{F0: 18446744073709551615, F1: 128, F2: -2147483648}
        _ = v34
    }
    var v35 uint = 1
    v20 = 216
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}

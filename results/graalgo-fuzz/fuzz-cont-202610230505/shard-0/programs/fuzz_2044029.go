package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }

func main() {
    var v1 float64 = float64(0.0)
    v2 := []uint32{1, 91}
    v3 := []int{-3}
    v4 := []string{"hello", ("foo bar" + ("a" + "hello"))}
    var v5 uint64 = 18446744073709551615
    var v6 int = -1
    v7 := GenS1{F0: 42, F1: 1000}
    fmt.Println(v5, v6, v5)
    v8 := GenS1{F0: 432, F1: 69}
    if (42 < 2) {
        switch "hello" {
        case "foo bar":
            v9 := GenS1{F0: 417, F1: 406}
            _ = v9
        case "":
            v6 = v6
        default:
            for fi10 := 0; fi10 < 7; fi10++ {
                var v11 int32 = 2147483647
                fmt.Println(v6, v11, v6)
                _ = len(v2)
                _ = fi10
                _ = v11
            }
        }
        v6 += -588
    }
    v2 = append(v2, 3)
    v12 := GenS1{F0: 0, F1: 32767}
    if ((!(42 != 32767)) && false) {
        var v13 uint32 = 835
        var v14 uint32 = (v13 & 952)
        _ = v13
        _ = v14
    }
    v5 -= v5
    for fi15 := 0; fi15 < 5; fi15++ {
        if ((0.1 + ((1e10 - -2.25) + 0.1)) >= (1.5 * 0.1)) {
            var v16 int8 = 1
            _ = v16
        }
        _ = fi15
    }
    var v17 uint = 65535
    var v18 string = "a"
    switch ((v18 + "\n\t") + (("foo bar" + v18) + "hello")) {
    case "hello":
        fmt.Println(v17, v6)
    case "a":
        switch ("a" + ((v18 + (v18 + "")) + "foo bar")) {
        case "\n\t":
            v3 = append(v3, v6)
        case "foo bar":
            v1 += 1e10
        case "hello":
            v1 = v1
        default:
            for fi19 := 0; fi19 < 9; fi19++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v1 = (v1 + (v1 * float64(1.5)))
                if true {
                    if false {
                        fmt.Println(v1, v17, v17)
                    }
                }
                _ = fi19
            }
        }
    case "":
        fmt.Println(v1, v1, v1)
    }
    if false {
        v18 = "\n\t"
        v20 := GenS1{F0: 1, F1: 32767}
        var v21 float32 = float32(v1)
        _ = v20
        _ = v21
    } else {
        for fi22 := 0; fi22 < 11; fi22++ {
            for fi23 := 0; fi23 < 3; fi23++ {
                if ((true || (127 >= -1)) && true) {
                    var v24 string = v18
                    _ = v24
                } else {
                    var v25 string = v18
                    v25 = ""
                    _ = v25
                }
                var v26 uint8 = 161
                v18 = (((("" + v18) + "") + "foo bar") + v18)
                _ = fi23
                _ = v26
            }
            var v27 int16 = 1
            var v28 int16 = v27
            _ = fi22
            _ = v27
            _ = v28
        }
        for fi29 := 0; fi29 < 5; fi29++ {
            switch fi29 {
            case 32767:
                for fi30 := 0; fi30 < 11; fi30++ {
                    var v31 uint16 = 42
                    var v32 uint64 = (v5 | 2)
                    var v33 float32 = 1e10
                    _ = fi30
                    _ = v31
                    _ = v32
                    _ = v33
                }
            default:
                var v34 int32 = 0
                _ = v34
            }
            var v35 int16 = -32768
            for fi36 := 0; fi36 < 7; fi36++ {
                var v37 uint32 = 747
                _ = fi36
                _ = v37
            }
            _ = fi29
            _ = v35
        }
        if (-378 <= -494) {
            fmt.Println(v18, v18, v5)
            v5 = (v5 / ((v5) | 1))
            v38 := GenS1{F0: 1, F1: -299}
            _ = v38
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

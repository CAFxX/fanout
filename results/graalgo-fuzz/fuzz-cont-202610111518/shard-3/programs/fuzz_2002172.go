package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 int32; F2 int64 }

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = 3.14159
    var v3 int8 = 42
    v4 := make(map[string]int)
    var v5 int8 = (v3 | v3)
    var v6 int8 = (v3 + v5)
    if true {
        for fi7 := 0; fi7 < 13; fi7++ {
            switch "foo bar" {
            case "":
                var v8 string = ("" + ("" + "a"))
                _ = v8
            case "hello":
                var v9 uint = 868
                _ = v9
            default:
                v10 := GenS1{F0: 612, F1: -1}
                _ = v10
            }
            fmt.Println(v3, v6, v6)
            _ = fi7
        }
        if (-1 > 152) {
            v6 = 42
            v11 := GenS1{F0: 2, F1: -2147483648}
            _ = v11
        }
        var v12 int = 4294967295
        _ = v12
    } else {
        for fi13 := 0; fi13 < 8; fi13++ {
            v14 := []int{(fi13 + (-(fi13))), 90, (fi13 % ((fi13) | 1))}
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            _ = fi13
            _ = v14
        }
        var v15 uint8 = 42
        v16 := GenS2{F0: 0}
        _ = v15
        _ = v16
    }
    v3 -= (v6 * v5)
    v5 += v5
    fmt.Println(v5, v6, v3)
    var v17 uint64 = 1
    if ((127 > 0) && true) {
        var v18 int32 = -1
        var v19 uint64 = (v17 + v17)
        _ = v18
        _ = v19
    }
    switch 2147483647 {
    case 1:
        var v20 float64 = float64(1.5)
        _ = v20
    }
    switch ((("\n\t" + ("\n\t" + "\n\t")) + "hello") + "a") {
    case "a":
        v17 = v17
    }
    v17 -= v17
    var v21 float64 = 0.1
    if (0 <= 1) {
        fmt.Println(v3, v6)
    } else {
        if (false && false) {
            var v22 float64 = (0.0 + float64((v2 - 1e10)))
            _ = v22
        } else {
            v21 = 3.14159
            v23 := []int{-9223372036854775808, 685}
            v24 := GenS3{F0: 684, F1: -2147483648, F2: -445}
            _ = v23
            _ = v24
        }
        var v25 uint16 = 297
        _ = v25
    }
    if false {
        fmt.Println(v5, v3, v5)
    } else {
        switch (("\n\t" + "foo bar") + ("a" + ("" + ""))) {
        case "foo bar":
            var v26 bool = (!false)
            _ = v26
        case "hello":
            v21 = 0.0
        default:
            v27 := []int{530, 12, -9223372036854775808, -1}
            _ = v27
        }
    }
    switch 984 {
    case 928:
        if false {
            var v28 int16 = 0
            for fi29 := 0; fi29 < 4; fi29++ {
                if ((0 < 2) && ((1 > -128) && (42 != 1))) {
                    switch (fi29 % ((fi29) | 1)) {
                    case -437:
                        delete(v4, "\n\t")
                    default:
                        v3 = v3
                    }
                    if (false && (!(v28 > 1000))) {
                        var v30 string = (("" + (("foo bar" + "") + "")) + "hello")
                        for fi31 := 0; fi31 < 19; fi31++ {
                            if false {
                                var v32 string = (("hello" + "a") + v30)
                                _ = v32
                            }
                            v30 = v30
                            v33 := GenS2{F0: 42}
                            _ = fi31
                            _ = v33
                        }
                        _ = v30
                    }
                    v3 = v5
                }
                v34 := GenS3{F0: 209, F1: 42, F2: -11}
                var v35 uint8 = 1
                _ = fi29
                _ = v34
                _ = v35
            }
            _ = v28
        }
    case -565:
        fmt.Println(v2, v5, v5)
    }
    v3 -= -128
    var v36 uint16 = 285
    var v37 int8 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}

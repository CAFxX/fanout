package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 bool; F1 int8; F2 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 int64; F2 int64 }

func main() {
    var v1 float64 = (0.0 / 3.14159)
    var v2 float64 = 0.1
    v3 := make(map[string]int)
    v4 := make(map[int]uint32)
    var v5 int = -9223372036854775808
    if (127 != v5) {
        for fi6 := 0; fi6 < 4; fi6++ {
            var v7 int32 = -496
            for fi8 := 0; fi8 < 3; fi8++ {
                var v9 int32 = 1
                v10 := GenS1{F0: 65535, F1: 23}
                fmt.Println(v5, v9, v7)
                v11 := GenS1{F0: 0, F1: 0}
                v12 := []int{(fi6 | fi8), v5, v5, v5}
                v1 = 0.0
                _ = fi8
                _ = v10
                _ = v11
                _ = v12
                _ = v9
            }
            _ = fi6
            _ = v7
        }
        switch 127 {
        case 1:
            v13 := []int{-411, -424}
            _ = v13
        case -532:
            if true {
                v14 := GenS2{F0: true, F1: -40, F2: 1}
                _ = v14
            }
        }
        _ = len(v3)
    }
    v15 := GenS3{F0: 2}
    v16 := GenS4{F0: (^(v5)), F1: 1, F2: 20}
    fmt.Println(v2, v5)
    for fi17 := 0; fi17 < 11; fi17++ {
        if (false || (2147483648 <= 4294967295)) {
            _ = len(v3)
            fmt.Println(fi17)
            var v18 uint = 1
            _ = v18
        } else {
            var v19 bool = false
            _ = v19
        }
        var v20 uint16 = 25
        _ = fi17
        _ = v20
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        switch (("a" + ("\n\t" + "foo bar")) + "hello") {
        case "foo bar":
            if ("a" > "foo bar") {
                var v22 uint = 968
                _ = v22
            } else {
                var v23 uint32 = 746
                v23 += v23
                if false {
                    switch v5 {
                    case 4294967295:
                        v5 -= (v5 + 4294967295)
                    case 32768:
                        var v24 int16 = 560
                        _ = v24
                    }
                    v23 -= v23
                }
                _ = v23
            }
        case "":
            v25 := []int{(-(fi21)), (v5 / (((fi21 % (((-(fi21))) | 1))) | 1)), -1}
            _ = v25
        case "hello":
            v26 := []int{(v5 * (-(fi21))), 4294967295, -2147483648, v5}
            _ = v26
        }
        v27 := []int{}
        for fi28 := 0; fi28 < 12; fi28++ {
            fmt.Println(v5, v5, v5)
            _ = len(v4)
            fmt.Println(v5)
            if false {
                fmt.Println(fi28, v1, v1)
            } else {
                if false {
                    var v29 uint16 = 220
                    fmt.Println(v5)
                    _ = v29
                }
            }
            _ = fi28
        }
        _ = fi21
        _ = v27
    }
    v30 := []int{v5}
    v2 -= (((0.1 + v2) + float64(-2.25)) * 1.5)
    if (1e10 <= (0.1 - float32(float64(((1.5 - (1e10 + (3.14159 * (1e10 * (-2.25 - 0.0))))) * ((((0.0 * 1e10) / 1.5) / 0.1) - -2.25)))))) {
        v31 := GenS3{F0: 2}
        var v32 float32 = 3.14159
        _ = v31
        _ = v32
    } else {
        var v33 float32 = 3.14159
        _ = v33
    }
    v34 := GenS1{F0: 1, F1: 1}
    fmt.Println(v2, v2)
    _ = len(v3)
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v5
}

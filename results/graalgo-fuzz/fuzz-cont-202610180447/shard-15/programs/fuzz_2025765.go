package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int64; F1 int64; F2 bool }
func genFn0(p0 string, p1 uint16, p2 uint) uint16 {
    switch p0 {
    case "a":
        v6 := []int{}
        _ = v6
    case "foo bar":
        var v7 bool = ((p2 | p2) == p2)
        _ = v7
    case "":
        fmt.Println(p0, p0)
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    return (^(p1))
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]bool)
    v3 := make(map[string]bool)
    var v4 float64 = 0.0
    var v5 uint = 345
    fmt.Println(v5, v4)
    switch ("\n\t" + (("" + ("" + "hello")) + "a")) {
    case "":
        v8 := GenS1{F0: (2147483647 <= 871), F1: 0}
        _ = v8
    case "foo bar":
        var v9 uint16 = 493
        _ = v9
    }
    switch 462 {
    case 42:
        switch -303 {
        case 9223372036854775807:
            fmt.Println(v4, v4)
        }
    case 946:
        v5 = (v5 + v5)
    case 244:
        if true {
            v10 := []int{0}
            v5 = 42
            _ = v10
        }
    }
    if true {
        if ((v5 < (v5 % (((v5 & 1)) | 1))) && true) {
            v4 -= (float64(3.14159) / (v4 + float64(0.1)))
            v11 := v2["\n\t"]
            _ = v11
        }
        switch "" {
        case "hello":
            v12 := GenS1{F0: false, F1: (v5 & (v5 * (v5 & v5)))}
            _ = v12
        case "foo bar":
            if true {
                v5 = v5
            }
        }
    } else {
        if (!(true && true)) {
            var v13 uint32 = 0
            v14 := GenS1{F0: false, F1: (v5 >> ((956) & 63))}
            var v15 int8 = 0
            _ = v13
            _ = v14
            _ = v15
        }
        if true {
            v4 = 1e10
            for fi16 := 0; fi16 < 17; fi16++ {
                for fi17 := 0; fi17 < 7; fi17++ {
                    v5 += v5
                    var v18 int8 = 120
                    if true {
                        fmt.Println(fi17)
                    } else {
                        v4 += -2.25
                        v18 += v18
                    }
                    _ = fi17
                    _ = v18
                }
                var v19 int16 = 620
                _ = fi16
                _ = v19
            }
        }
        var v20 float64 = (v4 - (-2.25 + (((v4 + (0.0 / ((1e10 - float64(-2.25)) - v4))) * (v4 - float64(float32((v4 - v4))))) + (((v4 / 3.14159) * v4) - 3.14159))))
        _ = v20
    }
    var v21 float32 = ((-2.25 - -2.25) - -2.25)
    var v22 uint16 = 1
    var v23 uint8 = 255
    var v24 int16 = 364
    var v25 uint32 = 0
    var v26 int8 = -1
    fmt.Println(v25)
    v27 := GenS2{F0: -987, F1: 87, F2: false}
    v24 = (v24 - -221)
    fmt.Println(v21)
    var v28 uint16 = v22
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}

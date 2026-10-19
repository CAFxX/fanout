package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int32 }
func genFn0() uint8 {
    if true {
        var v7 bool = (!(0 < 127))
        var v8 int8 = 42
        v9 := GenS1{F0: 0}
        _ = v7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := []int{}
        var v12 int16 = -562
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := []int{-986, 48, 128}
    var v14 bool = false
    var v15 uint = 42
    _ = v13
    _ = v14
    _ = v15
    return 127
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[string]bool)
    var v3 uint64 = 42
    var v4 uint = 2
    v5 := []string{((("" + "hello") + (("\n\t" + "hello") + "")) + ""), ("a" + "hello"), "\n\t", ("\n\t" + "a")}
    var v6 uint8 = 86
    for fi16 := 0; fi16 < 2; fi16++ {
        switch fi16 {
        case 2147483647:
            v17 := GenS1{F0: -2147483648}
            _ = v17
        case 1:
            var v18 bool = false
            _ = v18
        case 4294967295:
            v19 := GenS2{F0: 42, F1: 709}
            _ = v19
        }
        v4 += (v4 - v4)
        if (!(127 < -128)) {
            for fi20 := 0; fi20 < 16; fi20++ {
                if len(v5) > 0 {
                    v5[0] = "\n\t"
                }
                v6 -= v6
                _ = fi20
            }
        } else {
            switch (("foo bar" + "\n\t") + "") {
            case "foo bar":
                if (((0.0 + (float32(1.5) - -2.25)) / 1.0) <= 1e10) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v1 = 3.14159
                } else {
                    var v21 int8 = 1
                    v21 += (v21 * v21)
                    fmt.Println(fi16, v1)
                    _ = v21
                }
            case "a":
                var v22 float64 = (1e10 / 3.14159)
                _ = v22
            }
            v4 -= v4
        }
        _ = fi16
    }
    v6 = (v6 / ((v6) | 1))
    _ = len(v2)
    for fi23 := 0; fi23 < 6; fi23++ {
        var v24 int32 = -1
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi23
        _ = v24
    }
    _ = len(v5)
    v1 -= float32(0.0)
    v6 = (v6 & 127)
    fmt.Println(v6, v6)
    var v25 uint64 = 2
    if (18446744073709551615 != v4) {
        v3 += (^(v25))
    } else {
        var v26 int64 = -646
        v27 := GenS2{F0: (v4 * v4), F1: 105}
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 11; fi28++ {
        v29 := GenS1{F0: -263}
        fmt.Println(v6, fi28, v6)
        v30 := GenS1{F0: 0}
        _ = fi28
        _ = v29
        _ = v30
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint }
func genFn0(p0 uint8) uint32 {
    fmt.Println(p0)
    if false {
        p0 = (p0 % (((^(p0))) | 1))
        var v5 int = -2147483648
        _ = v5
    }
    for fi6 := 0; fi6 < 19; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            p0 = 255
            var v8 uint16 = 42
            if (!true) {
                if (17 == -128) {
                    v8 = 32767
                    v9 := GenS1{F0: 32767, F1: v8, F2: 213}
                    var v10 int64 = int64(v8)
                    _ = v10
                    _ = v9
                }
            } else {
                var v11 uint8 = p0
                if false {
                    fmt.Println(v8, fi6, fi7)
                    v12 := []int{fi6, fi7, 0, fi7}
                    switch fi7 {
                    case 734:
                        var v13 int8 = 0
                        _ = v13
                    case 9223372036854775807:
                        _ = len(v12)
                        if false {
                            p0 -= (p0 / (((^(v11))) | 1))
                        } else {
                            var v14 int8 = 100
                            _ = v14
                        }
                    case 45:
                        v8 = (^(v8))
                    }
                    _ = v12
                }
                var v15 string = "\n\t"
                _ = v11
                _ = v15
            }
            _ = fi7
            _ = v8
        }
        p0 += p0
        var v16 float64 = float64(3.14159)
        _ = fi6
        _ = v16
    }
    if (!true) {
        var v17 uint8 = p0
        for fi18 := 0; fi18 < 3; fi18++ {
            if true {
                for fi19 := 0; fi19 < 7; fi19++ {
                    var v20 uint = 65535
                    v17 += p0
                    _ = fi19
                    _ = v20
                }
                v21 := GenS1{F0: 32767, F1: 8, F2: 255}
                _ = v21
            }
            _ = fi18
        }
        _ = v17
    }
    _ = p0
    return 0
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]uint32)
    v3 := make(map[string]bool)
    var v4 uint16 = 65
    v22 := []int{}
    fmt.Println(v4)
    for fi23 := 0; fi23 < 2; fi23++ {
        fmt.Println(v4, v4)
        _ = fi23
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        var v25 string = "\n\t"
        _ = fi24
        _ = v25
    }
    var v26 int = 9223372036854775807
    if false {
        fmt.Println(v4, v4)
    }
    v27 := GenS1{F0: -974, F1: (v4 - v4), F2: 42}
    for fi28 := 0; fi28 < 17; fi28++ {
        for fi29 := 0; fi29 < 5; fi29++ {
            switch (("a" + "\n\t") + "\n\t") {
            case "foo bar":
                _ = len(v22)
            case "hello":
                fmt.Println(v26, v4, v4)
            default:
                if ((!(false || (540 == 1000))) || false) {
                    fmt.Println(v26, fi29, fi28)
                } else {
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                }
            }
            _ = fi29
        }
        if (!false) {
            v4 -= v4
            switch "\n\t" {
            case "foo bar":
                var v30 int8 = -89
                _ = v30
            case "\n\t":
                v31 := GenS1{F0: 1, F1: (v4 >> ((0) & 63)), F2: 4294967295}
                _ = v31
            }
        } else {
            v22 = append(v22, fi28)
            v4 = v4
            var v32 int = v26
            _ = v32
        }
        _ = fi28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}

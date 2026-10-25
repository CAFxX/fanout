package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint32 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]int)
    v3 := make(map[int]bool)
    switch "\n\t" {
    case "a":
        var v4 uint64 = 722
        _ = v4
    case "":
        v5 := v1[32768]
        _ = v5
    case "hello":
        if (false && true) {
            var v6 uint16 = 0
            for fi7 := 0; fi7 < 9; fi7++ {
                v6 += v6
                _ = fi7
            }
            var v8 int = 874
            _ = v6
            _ = v8
        }
    default:
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 int = (-(fi9))
            if ((float32(1.5) / 1.0) >= 0.0) {
                v11 := GenS1{F0: ("hello" + ""), F1: 4294967295, F2: 2147483648}
                _ = v11
            } else {
                v12 := GenS1{F0: "", F1: 65535, F2: 4294967295}
                _ = v12
            }
            var v13 int64 = 0
            _ = fi9
            _ = v10
            _ = v13
        }
    }
    var v14 int = 32767
    v15 := []int{(v14 >> ((709) & 63)), v14, v14}
    v15 = append(v15, 128)
    fmt.Println(v14)
    var v16 float64 = -2.25
    var v17 int = (v14 & v14)
    var v18 uint = 1
    if (v18 >= v18) {
        if (!true) {
            var v19 string = "a"
            _ = v19
        }
    }
    var v20 uint16 = 65535
    fmt.Println(v17, v17)
    if true {
        v20 = v20
    } else {
        for fi21 := 0; fi21 < 7; fi21++ {
            switch "a" {
            case "a":
                v22 := GenS1{F0: "\n\t", F1: (v18 / ((4294967295) | 1)), F2: 2147483648}
                _ = v22
            }
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v14)
            case "hello":
                v14 += v17
            case "a":
                v16 = (v16 - float64(3.14159))
            }
            _ = fi21
        }
        if (false && false) {
            for fi23 := 0; fi23 < 4; fi23++ {
                v17 += (v17 & (fi23 + fi23))
                if len(v15) > 0 {
                    v15[0] = (v14 ^ fi23)
                }
                switch (fi23 >> ((v18) & 63)) {
                case 600:
                    v14 = v17
                default:
                    if (!(2 >= 1)) {
                        var v24 uint64 = 549
                        v18 = (v18 << ((v18) & 63))
                        switch "\n\t" {
                        case "hello":
                            fmt.Println(v16, fi23, v18)
                        }
                        _ = v24
                    }
                }
                _ = fi23
            }
        }
    }
    fmt.Println(v16)
    _ = len(v15)
    fmt.Println(v14, v14)
    v16 = (v16 / (v16 * (1.5 - (((v16 + v16) * 0.0) * v16))))
    var v25 int = v17
    for fi26 := 0; fi26 < 16; fi26++ {
        var v27 uint16 = 0
        var v28 int16 = 32767
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v25
    _ = v3
}

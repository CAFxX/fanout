package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 bool; F1 uint64; F2 int8 }
func genFn0(p0 uint, p1 string, p2 bool) uint8 {
    fmt.Println(p2, p2, p1)
    var v7 int16 = -163
    var v8 int16 = v7
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 100
}

func main() {
    v1 := []string{"", "\n\t", "hello"}
    var v2 int32 = 42
    v3 := make(map[string]int)
    var v4 float64 = 0.1
    var v5 bool = true
    var v6 int8 = 42
    v9 := v3[""]
    v10 := GenS1{F0: 65535, F1: 1000}
    if v5 {
        var v11 uint16 = 2
        _ = v11
    } else {
        v2 = 1
        v2 -= v2
    }
    switch v9 {
    case -319:
        var v12 int64 = 0
        _ = v12
    case 32767:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v6 += (^(v6))
    var v13 int16 = 0
    v4 = (-2.25 + (3.14159 / v4))
    if (218 == 1) {
        v14 := GenS1{F0: 959, F1: (v13 ^ v13)}
        v15 := GenS1{F0: 1, F1: -1}
        v16 := v3[("\n\t" + ("" + "foo bar"))]
        _ = v14
        _ = v15
        _ = v16
    } else {
        v2 = (v2 + v2)
    }
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        v2 = 731
    case "a":
        for fi17 := 0; fi17 < 7; fi17++ {
            switch "a" {
            case "":
                delete(v3, "\n\t")
            case "foo bar":
                v3["a"] = fi17
            }
            if true {
                var v18 uint = 255
                var v19 int16 = 1
                _ = len(v3)
                _ = v18
                _ = v19
            }
            v20 := GenS1{F0: 42, F1: v13}
            _ = fi17
            _ = v20
        }
    }
    v6 -= (v6 & (v6 % (((v6 + 103)) | 1)))
    fmt.Println(v6, v9, v6)
    switch 128 {
    case 42:
        var v21 uint8 = genFn0(331, "a", (42 <= -380))
        _ = v21
    case 791:
        for fi22 := 0; fi22 < 13; fi22++ {
            v23 := GenS2{F0: ((float64(1e10) / v4) < 0.1), F1: 0, F2: -12}
            for fi24 := 0; fi24 < 14; fi24++ {
                var v25 float64 = ((v4 + v4) + (float64((-2.25 + (0.1 / (0.1 + 1e10)))) * 0.0))
                if v5 {
                    v2 += (v2 + -887)
                    v13 += 0
                } else {
                    fmt.Println(v13, v2, fi22)
                }
                fmt.Println(v13, v9, fi22)
                _ = fi24
                _ = v25
            }
            _ = fi22
            _ = v23
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}

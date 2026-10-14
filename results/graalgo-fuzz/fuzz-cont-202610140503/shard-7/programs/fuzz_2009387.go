package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 uint16; F2 int32 }

func main() {
    v1 := []string{"hello", "", (("hello" + "hello") + ("foo bar" + "a"))}
    v2 := make(map[string]string)
    var v3 uint16 = 330
    v4 := make(map[string]int)
    var v5 float32 = 1.5
    fmt.Println(v3, v5)
    fmt.Println(v3)
    for fi6 := 0; fi6 < 13; fi6++ {
        for fi7 := 0; fi7 < 15; fi7++ {
            fmt.Println(v3, fi6, fi6)
            if ((fi7 + (-(fi6))) == fi7) {
                v8 := GenS1{F0: 65535, F1: 2147483647}
                _ = v8
            } else {
                v5 += (v5 * v5)
                v3 = v3
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            v9 := GenS1{F0: 1, F1: -355}
            _ = fi7
            _ = v9
        }
        for fi10 := 0; fi10 < 3; fi10++ {
            switch (fi6 ^ (fi10 % (((fi6 >> ((1) & 63))) | 1))) {
            case 0:
                v5 = ((((float32(1e10) + (-2.25 - 0.1)) - (v5 * ((v5 - v5) / v5))) * (v5 - v5)) / (float32(1.5) - v5))
            }
            var v11 int8 = 42
            v3 = v3
            _ = fi10
            _ = v11
        }
        _ = fi6
    }
    var v12 int16 = -1
    var v13 int16 = v12
    v14 := v4["hello"]
    for fi15 := 0; fi15 < 12; fi15++ {
        for fi16 := 0; fi16 < 3; fi16++ {
            v12 = (v13 | (v13 - (v13 >> ((42) & 63))))
            switch fi15 {
            case 251:
                var v17 float64 = (1e10 * float64(float32((0.0 / 1.0))))
                _ = v17
            case 32768:
                var v18 int8 = 0
                _ = v18
            default:
                v19 := GenS1{F0: 2, F1: 616}
                _ = v19
            }
            var v20 uint64 = 2
            _ = fi16
            _ = v20
        }
        for fi21 := 0; fi21 < 14; fi21++ {
            if ("a" <= ("\n\t" + "hello")) {
                var v22 float64 = (0.0 / 3.14159)
                v23 := GenS1{F0: 485, F1: -1}
                v3 -= 865
                _ = v22
                _ = v23
            } else {
                var v24 uint64 = 18446744073709551615
                var v25 int32 = -2147483648
                _ = v24
                _ = v25
            }
            fmt.Println(fi15, v14)
            _ = fi21
        }
        v14 -= v14
        _ = fi15
    }
    _ = len(v1)
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    if ("\n\t" > "a") {
        v12 = v12
    } else {
        if true {
            v13 -= v13
            var v26 float64 = 0.1
            for fi27 := 0; fi27 < 16; fi27++ {
                var v28 float64 = float64(v5)
                switch (("\n\t" + "hello") + "a") {
                case "":
                    v29 := v4["a"]
                    _ = v29
                case "\n\t":
                    if false {
                        v30 := GenS2{F0: true}
                        _ = v30
                    }
                default:
                    v14 = (^(fi27))
                }
                v31 := GenS3{F0: -68, F1: (v3 - 896), F2: 553}
                _ = fi27
                _ = v28
                _ = v31
            }
            _ = v26
        } else {
            _ = len(v1)
        }
        var v32 int32 = 1
        _ = v32
    }
    if (!true) {
        var v33 bool = true
        _ = v33
    }
    if (false || true) {
        if (1 > 449) {
            fmt.Println(v12)
        }
        var v34 string = ""
        fmt.Println(v13)
        _ = v34
    } else {
        if (4294967295 != 42) {
            var v35 int32 = -1
            _ = v35
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

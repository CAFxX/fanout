package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 uint8; F2 uint }

func main() {
    var v1 string = "a"
    v2 := make(map[string]int)
    var v3 int = -2147483648
    var v4 string = ((v1 + v1) + "\n\t")
    var v5 int = 671
    fmt.Println(v5, v5, v4)
    var v6 float64 = 3.14159
    for fi7 := 0; fi7 < 3; fi7++ {
        switch 1 {
        case -2147483648:
            fmt.Println(v6)
        }
        fmt.Println(v6)
        var v8 int64 = 1
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        switch (^(fi9)) {
        case -1:
            if true {
                var v10 uint = 0
                if (true || true) {
                    v6 = (1e10 - 0.0)
                    switch v1 {
                    case "\n\t":
                        for fi11 := 0; fi11 < 10; fi11++ {
                            fmt.Println(v10, v6, v6)
                            _ = fi11
                        }
                    case "a":
                        var v12 int = v3
                        _ = v12
                    case "":
                        v13 := GenS1{F0: 42}
                        _ = v13
                    default:
                        for fi14 := 0; fi14 < 14; fi14++ {
                            var v15 int = (v3 % ((128) | 1))
                            var v16 int64 = 9223372036854775807
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    }
                    fmt.Println(v3)
                } else {
                    v17 := GenS2{F0: -2147483648, F1: 0, F2: 136}
                    if (-1 == 20) {
                        v10 = 679
                        for fi18 := 0; fi18 < 8; fi18++ {
                            fmt.Println(v3)
                            var v19 uint = 0
                            v4 = ""
                            _ = fi18
                            _ = v19
                        }
                        fmt.Println(v6, v3)
                    }
                    _ = v17
                }
                var v20 float64 = ((-2.25 * float64(-2.25)) + float64((float32((v6 * v6)) - ((0.0 * 3.14159) * 0.1))))
                _ = v10
                _ = v20
            } else {
                v21 := []int{v3, 641}
                _ = v21
            }
        default:
            fmt.Println(v3, v4)
        }
        v22 := GenS1{F0: 1}
        _ = fi9
        _ = v22
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        fmt.Println(v6, v3)
        v4 = "hello"
        _ = fi23
    }
    if true {
        switch (v4 + "") {
        case "":
            for fi24 := 0; fi24 < 2; fi24++ {
                v2[("\n\t" + v4)] = fi24
                fmt.Println(v3)
                _ = fi24
            }
        case "a":
            fmt.Println(v4, v1)
        }
        v3 = 127
    } else {
        v25 := GenS1{F0: 0}
        v1 = v4
        _ = v25
    }
    var v26 int64 = 42
    var v27 uint = 982
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

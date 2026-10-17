package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 string }
func genFn0(p0 int16) int8 {
    fmt.Println(p0)
    v5 := []int{32768, -808, 600}
    _ = p0
    _ = v5
    return 0
}
func genFn1(p0 uint64) uint16 {
    v6 := []int{4294967295, -2147483648, 32767, 127}
    switch -951 {
    case 0:
        switch "a" {
        case "":
            for fi7 := 0; fi7 < 14; fi7++ {
                p0 += p0
                _ = fi7
            }
        case "foo bar":
            switch (((("hello" + "hello") + "a") + "") + "hello") {
            case "foo bar":
                v8 := GenS1{F0: 983}
                _ = v8
            case "\n\t":
                var v9 int64 = -691
                _ = v9
            case "a":
                for fi10 := 0; fi10 < 8; fi10++ {
                    switch 77 {
                    case 704:
                        if (0 != (p0 << ((uint(p0)) & 63))) {
                            var v11 uint16 = 755
                            _ = v11
                        }
                    }
                    fmt.Println(fi10)
                    v6 = append(v6, fi10)
                    _ = fi10
                }
            }
        case "a":
            fmt.Println(p0)
        }
    case 9223372036854775807:
        p0 += 4294967295
    case 749:
        p0 += 4294967295
    }
    v12 := GenS1{F0: 2147483647}
    if true {
        v13 := GenS1{F0: 129}
        for fi14 := 0; fi14 < 1; fi14++ {
            if (p0 < 42) {
                p0 += p0
                v15 := GenS1{F0: 2}
                _ = v15
            }
            if false {
                var v16 int16 = 1
                _ = v16
            }
            _ = fi14
        }
        _ = v13
    }
    p0 += p0
    _ = p0
    _ = v12
    _ = v6
    return 493
}

func main() {
    var v1 float32 = 3.14159
    var v2 uint8 = 127
    v3 := make(map[uint32]bool)
    v4 := []string{"foo bar", ("a" + "hello"), "hello", "a"}
    if (0 >= -310) {
        var v17 int16 = 923
        var v18 int16 = (v17 & -32768)
        _ = v17
        _ = v18
    } else {
        var v19 int = 597
        _ = v19
    }
    var v20 float64 = (1e10 * (0.0 / 1.5))
    if ((0.0 * (v1 - v1)) > (v1 - 1e10)) {
        v1 -= v1
        if false {
            v21 := GenS2{F0: 32767, F1: "\n\t"}
            var v22 uint = 65535
            _ = v21
            _ = v22
        } else {
            if (false || (!true)) {
                for fi23 := 0; fi23 < 4; fi23++ {
                    v20 += (1.5 + v20)
                    for fi24 := 0; fi24 < 1; fi24++ {
                        var v25 uint64 = 964
                        fmt.Println(v1, v20, v2)
                        v25 += v25
                        fmt.Println(v2)
                        _ = fi24
                        _ = v25
                    }
                    _ = fi23
                }
                v26 := GenS2{F0: 2, F1: "\n\t"}
                _ = v26
            }
        }
        if (9223372036854775807 > -299) {
            if true {
                var v27 uint32 = 736
                _ = v27
            }
            switch ("" + "hello") {
            case "":
                var v28 int8 = 42
                _ = v28
            case "hello":
                var v29 uint16 = 42
                _ = v29
            case "\n\t":
                var v30 uint64 = 2
                _ = v30
            default:
                if len(v4) > 0 {
                    v4[0] = "\n\t"
                }
            }
            if (false || ((-1 > -1) && false)) {
                if len(v4) > 0 {
                    v4[0] = "\n\t"
                }
                v1 = 3.14159
            }
        } else {
            if len(v4) > 0 {
                v4[0] = "hello"
            }
            switch ("\n\t" + ("" + ("" + "foo bar"))) {
            case "\n\t":
                v2 -= v2
            case "a":
                var v31 int16 = -32768
                _ = v31
            case "foo bar":
                v2 += v2
            default:
                v20 += v20
            }
        }
    }
    v3[2] = true
    for fi32 := 0; fi32 < 5; fi32++ {
        for fi33 := 0; fi33 < 19; fi33++ {
            if (("\n\t" + "") <= "a") {
                if (true && true) {
                    fmt.Println(v1, fi33, v1)
                }
                v34 := GenS1{F0: 2147483648}
                v35 := GenS1{F0: 1}
                _ = v34
                _ = v35
            } else {
                v20 -= 0.1
            }
            _ = fi33
        }
        _ = fi32
    }
    v36 := GenS2{F0: 433, F1: ("\n\t" + "a")}
    fmt.Println(v2, v1, v20)
    for fi37 := 0; fi37 < 12; fi37++ {
        fmt.Println(fi37, v1, v20)
        _ = fi37
    }
    switch 127 {
    case 127:
        for fi38 := 0; fi38 < 17; fi38++ {
            v1 = v1
            var v39 int8 = 100
            _ = fi38
            _ = v39
        }
    default:
        fmt.Println(v20)
    }
    var v40 uint8 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v36
    _ = v4
    _ = v40
}

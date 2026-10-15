package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 bool }

func main() {
    var v1 string = "\n\t"
    var v2 uint32 = 377
    var v3 string = v1
    var v4 bool = (!((!false) || (0 < 2)))
    var v5 uint16 = 2
    if v4 {
        var v6 uint8 = 133
        v5 += 255
        _ = v6
    } else {
        var v7 uint32 = v2
        v7 = 446
        _ = v7
    }
    var v8 int32 = -994
    var v9 int = -179
    v10 := []int{}
    for fi11 := 0; fi11 < 8; fi11++ {
        fmt.Println(v9, fi11, v9)
        v12 := GenS1{F0: -1, F1: 1, F2: (65535 != 1)}
        _ = fi11
        _ = v12
    }
    switch v9 {
    case 1:
        fmt.Println(v1)
    }
    v8 = -1
    var v13 uint64 = 1
    if len(v10) > 0 {
        v10[0] = v9
    }
    v9 -= (v9 - (v9 / ((940) | 1)))
    _ = len(v10)
    for fi14 := 0; fi14 < 13; fi14++ {
        switch ("foo bar" + v3) {
        case "":
            var v15 float32 = 1.5
            _ = v15
        case "a":
            _ = len(v10)
        case "\n\t":
            v1 = "\n\t"
        default:
            if false {
                switch (v3 + "foo bar") {
                case "hello":
                    if v4 {
                        var v16 float64 = -2.25
                        for fi17 := 0; fi17 < 3; fi17++ {
                            var v18 float32 = 1.5
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            v5 = v5
                            fmt.Println(v1, fi17)
                            var v19 int64 = 9223372036854775807
                            switch ((v1 + v3) + v1) {
                            case "foo bar":
                                var v20 string = (v3 + v3)
                                _ = v20
                            case "a":
                                var v21 bool = (!v4)
                                _ = v21
                            case "\n\t":
                                v9 = fi17
                            default:
                                var v22 uint = 263
                                _ = v22
                            }
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        _ = len(v10)
                        _ = v16
                    }
                case "foo bar":
                    var v23 int8 = 42
                    _ = v23
                }
                if (v5 <= v5) {
                    v5 -= 1
                    if len(v10) > 0 {
                        v10[0] = (-(fi14))
                    }
                }
                for fi24 := 0; fi24 < 1; fi24++ {
                    v9 = fi14
                    _ = fi24
                }
            }
        }
        v25 := GenS1{F0: fi14, F1: 2, F2: ((0 >= 1) || (!false))}
        for fi26 := 0; fi26 < 19; fi26++ {
            fmt.Println(v9)
            fmt.Println(v4, v13)
            if len(v10) > 0 {
                v10[0] = (fi26 ^ 4294967295)
            }
            _ = fi26
        }
        _ = fi14
        _ = v25
    }
    var v27 uint16 = (v5 | v5)
    v27 += v27
    var v28 int = 127
    if v4 {
        if v4 {
            var v29 int16 = 42
            _ = v29
        } else {
            var v30 int8 = 42
            var v31 string = ((((v1 + v3) + v1) + (("a" + ((((v3 + (v3 + ("" + ((v1 + (v3 + v3)) + "\n\t")))) + "a") + "a") + "")) + ("a" + "a"))) + v3)
            _ = v30
            _ = v31
        }
    }
    fmt.Println(v4, v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}

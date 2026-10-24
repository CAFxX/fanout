package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 uint }
func genFn0(p0 string, p1 uint16) uint32 {
    p0 = p0
    v5 := []int{-2147483648, -753, -1, 98}
    if (-488 >= -439) {
        v5 = append(v5, 314)
        fmt.Println(p0, p0)
    }
    _ = p0
    _ = p1
    _ = v5
    return 706
}
func genFn1() bool {
    if (0.0 != -2.25) {
        var v6 uint64 = 798
        _ = v6
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 uint = 0
        _ = fi7
        _ = v8
    }
    var v9 float64 = ((0.1 - 1.5) - (0.0 - (-2.25 - 0.0)))
    v9 -= 0.0
    _ = v9
    return (!(v9 == v9))
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[int]bool)
    var v3 uint = 600
    var v4 uint64 = v1
    v10 := GenS1{F0: 5}
    v1 = v4
    delete(v2, -852)
    switch "a" {
    case "\n\t":
        var v11 int = 128
        _ = v11
    case "":
        if true {
            fmt.Println(v4)
        } else {
            v3 -= v3
        }
    case "foo bar":
        fmt.Println(v1, v1)
    }
    v3 = v3
    switch 127 {
    case 127:
        v12 := GenS2{F0: (v3 - (v3 - (v3 & 255))), F1: v3}
        _ = v12
    case -2147483648:
        for fi13 := 0; fi13 < 20; fi13++ {
            fmt.Println(v4)
            _ = fi13
        }
    case 0:
        for fi14 := 0; fi14 < 17; fi14++ {
            if true {
                for fi15 := 0; fi15 < 7; fi15++ {
                    switch ((((("a" + "") + ("\n\t" + "a")) + ("" + (("\n\t" + "") + "\n\t"))) + "\n\t") + "foo bar") {
                    case "hello":
                        switch (fi15 >> ((v3) & 63)) {
                        case -528:
                            v16 := GenS2{F0: (^(v3)), F1: (^(v3))}
                            _ = v16
                        case 272:
                            var v17 uint32 = 42
                            _ = v17
                        default:
                            switch fi14 {
                            case 128:
                                fmt.Println(v1, v4)
                            case -338:
                                var v18 int8 = 127
                                _ = v18
                            default:
                                var v19 uint16 = 1
                                _ = v19
                            }
                        }
                    default:
                        var v20 string = "hello"
                        _ = v20
                    }
                    var v21 int16 = 320
                    var v22 int32 = -981
                    var v23 int32 = v22
                    _ = fi15
                    _ = v21
                    _ = v22
                    _ = v23
                }
                v24 := GenS1{F0: v4}
                _ = v24
            }
            v4 += v4
            _ = fi14
        }
    default:
        switch -864 {
        case -2147483648:
            v25 := GenS1{F0: v4}
            _ = v25
        }
    }
    var v26 bool = true
    if v26 {
        var v27 int64 = -1
        switch "" {
        case "a":
            if (!(127 == 127)) {
                v1 = (^(v4))
            }
        case "\n\t":
            delete(v2, 2147483647)
        default:
            var v28 uint = v3
            _ = v28
        }
        _ = v27
    }
    var v29 uint = v3
    var v30 int = -15
    for fi31 := 0; fi31 < 11; fi31++ {
        fmt.Println(v30)
        _ = fi31
    }
    if v26 {
        if (true || true) {
            v4 = v4
        } else {
            fmt.Println(v3, v4)
        }
        var v32 uint64 = (v4 * v1)
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}

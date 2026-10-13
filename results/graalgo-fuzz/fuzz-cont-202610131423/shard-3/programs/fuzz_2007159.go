package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int64 }

func main() {
    var v1 uint64 = 657
    v2 := []uint32{4294967295, 1, 1}
    var v3 float32 = 0.0
    v4 := []string{"foo bar"}
    var v5 int16 = 2
    v2 = append(v2, 2)
    var v6 uint8 = 1
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 uint32 = 2
        if true {
            v9 := GenS1{F0: false, F1: false, F2: 675}
            switch (("" + "foo bar") + "foo bar") {
            case "hello":
                switch (fi7 | fi7) {
                case -14:
                    fmt.Println(v8)
                }
            case "foo bar":
                if false {
                    if true {
                        var v10 float32 = v3
                        _ = len(v2)
                        if ((v8 | 449) != (v8 | v8)) {
                            var v11 uint16 = 900
                            v1 += v1
                            v4 = append(v4, "\n\t")
                            _ = v11
                        }
                        _ = v10
                    }
                    v3 += 1e10
                    v12 := GenS1{F0: (!false), F1: (false && false), F2: -9223372036854775808}
                    _ = v12
                } else {
                    var v13 string = "hello"
                    _ = v13
                }
            case "a":
                for fi14 := 0; fi14 < 8; fi14++ {
                    var v15 int = fi7
                    _ = fi14
                    _ = v15
                }
            default:
                if true {
                    var v16 uint64 = v1
                    v6 += 255
                    _ = v16
                }
            }
            _ = v9
        } else {
            if (!(false || false)) {
                _ = len(v2)
            } else {
                if (-1 <= v5) {
                    v3 = 3.14159
                    fmt.Println(v8, v1, v1)
                }
            }
        }
        _ = fi7
        _ = v8
    }
    var v17 int16 = (v5 / ((v5) | 1))
    if true {
        fmt.Println(v1)
        fmt.Println(v1, v3, v17)
    }
    if ((("" + "\n\t") > "") && ("hello" != ("foo bar" + "\n\t"))) {
        v2 = append(v2, 2)
        switch 32767 {
        case -464:
            fmt.Println(v17)
        default:
            v18 := GenS1{F0: false, F1: true, F2: 0}
            _ = v18
        }
    } else {
        var v19 bool = false
        for fi20 := 0; fi20 < 2; fi20++ {
            v1 = (v1 - 18446744073709551615)
            _ = fi20
        }
        var v21 int32 = -531
        _ = v19
        _ = v21
    }
    v22 := GenS1{F0: false, F1: (-828 >= 32768), F2: 0}
    var v23 uint16 = 65535
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

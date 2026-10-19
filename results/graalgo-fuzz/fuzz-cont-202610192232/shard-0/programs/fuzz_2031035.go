package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 string; F1 int8; F2 int8 }

func main() {
    var v1 string = "\n\t"
    v2 := []string{((("" + v1) + v1) + ((v1 + (v1 + "a")) + "a")), (v1 + v1), v1, "hello"}
    var v3 int32 = -2147483648
    v4 := make(map[string]string)
    var v5 int16 = 42
    v6 := GenS1{F0: 0, F1: 100}
    v3 = (-(v3))
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            v5 = v5
            v9 := v4["foo bar"]
            v10 := v4[(v9 + v1)]
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    v11 := GenS1{F0: 239, F1: 127}
    var v12 int32 = v3
    var v13 string = (v1 + (("\n\t" + (v1 + ((v1 + v1) + v1))) + ("a" + "hello")))
    switch -9223372036854775808 {
    case -874:
        for fi14 := 0; fi14 < 8; fi14++ {
            switch v1 {
            case "a":
                var v15 int8 = 100
                _ = v15
            }
            _ = fi14
        }
    case 2147483647:
        v16 := GenS1{F0: 32767, F1: 112}
        _ = v16
    }
    var v17 uint8 = 128
    if false {
        v12 += (v3 * (v3 + 2147483647))
        var v18 string = "hello"
        _ = v18
    }
    if (((-2.25 + 1e10) - -2.25) == 0.0) {
        switch 9223372036854775807 {
        case 127:
            if false {
                for fi19 := 0; fi19 < 1; fi19++ {
                    for fi20 := 0; fi20 < 2; fi20++ {
                        delete(v4, v1)
                        v17 = v17
                        _ = fi20
                    }
                    var v21 uint16 = uint16(v17)
                    _ = fi19
                    _ = v21
                }
            }
        }
    } else {
        v22 := GenS2{F0: v13, F1: "foo bar"}
        v23 := GenS1{F0: 2, F1: 127}
        v24 := GenS3{F0: (("a" + (v13 + "foo bar")) + (v1 + "\n\t")), F1: 1, F2: -49}
        _ = v22
        _ = v23
        _ = v24
    }
    switch ((v13 + "foo bar") + (((v13 + ("a" + ("" + ("\n\t" + ((v13 + v1) + v1))))) + (((v1 + v13) + ("foo bar" + (v13 + v13))) + v1)) + ((v1 + v1) + v1))) {
    case "\n\t":
        if (4294967295 != 255) {
            if ((127 >= -1) && true) {
                fmt.Println(v12, v17)
                v12 -= (v12 | v3)
            }
            var v25 float32 = (0.1 + 1.5)
            _ = v25
        } else {
            v17 -= 42
        }
    }
    v5 -= v5
    v26 := GenS2{F0: ("\n\t" + (v13 + "a")), F1: v13}
    fmt.Println(v3, v3, v17)
    for fi27 := 0; fi27 < 11; fi27++ {
        for fi28 := 0; fi28 < 16; fi28++ {
            fmt.Println(fi28, v1)
            _ = fi28
        }
        _ = fi27
    }
    var v29 int64 = 9223372036854775807
    v13 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

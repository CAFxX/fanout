package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = false
    var v3 int8 = 100
    fmt.Println(v3)
    fmt.Println(v3)
    v2 = true
    if (!(4294967295 >= 234)) {
        v1[4294967295] = 32768
    }
    var v4 int = 32767
    if ((v3 & -128) >= 1) {
        v5 := GenS1{F0: 42, F1: 42}
        var v6 float64 = 0.0
        _ = v5
        _ = v6
    }
    delete(v1, 849)
    if (v3 >= v3) {
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := v1[2147483648]
            _ = fi7
            _ = v8
        }
        var v9 int = v4
        _ = v9
    } else {
        if v2 {
            v10 := []int{}
            switch v4 {
            case 128:
                var v11 int8 = -128
                _ = v11
            case 421:
                v2 = v2
            }
            for fi12 := 0; fi12 < 12; fi12++ {
                for fi13 := 0; fi13 < 9; fi13++ {
                    if ("a" >= "foo bar") {
                        var v14 int = fi13
                        _ = v14
                    } else {
                        v10 = append(v10, 128)
                        if v2 {
                            var v15 bool = false
                            _ = v15
                        } else {
                            var v16 string = "a"
                            v17 := GenS1{F0: 255, F1: 18446744073709551615}
                            fmt.Println(fi12, v2, fi13)
                            _ = v16
                            _ = v17
                        }
                        var v18 float64 = ((-2.25 + 3.14159) - (1e10 * (-2.25 + (0.1 * (0.1 - 1.5)))))
                        _ = v18
                    }
                    _ = fi13
                }
                _ = fi12
            }
            _ = v10
        }
        v3 = v3
    }
    v19 := GenS1{F0: 255, F1: 18446744073709551615}
    switch ("foo bar" + ("foo bar" + ("hello" + ("" + "foo bar")))) {
    case "a":
        for fi20 := 0; fi20 < 16; fi20++ {
            var v21 int64 = -9223372036854775808
            v2 = v2
            if ((^(v4)) != 32767) {
                v2 = (float32(float64(0.0)) <= (0.1 / 3.14159))
                switch 32768 {
                case -635:
                    v21 = 400
                }
            }
            _ = fi20
            _ = v21
        }
    case "\n\t":
        var v22 int = -201
        _ = v22
    }
    v1[65535] = 32767
    var v23 float32 = float32(1.5)
    v23 -= 1.5
    v23 -= (float32(0.1) + 1e10)
    var v24 string = "foo bar"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}

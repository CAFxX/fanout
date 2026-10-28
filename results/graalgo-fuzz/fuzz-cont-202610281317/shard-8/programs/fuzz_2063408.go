package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint64 = 4294967295
    v2 := []string{}
    v3 := make(map[int]bool)
    v1 -= v1
    if (true && false) {
        if (((0.0 * -2.25) + (float64((0.0 * 0.1)) - ((1e10 - 1.5) - -2.25))) <= (((0.1 / 1e10) - ((0.0 / -2.25) * 0.0)) + float64(1.5))) {
            switch ("a" + "foo bar") {
            case "a":
                var v4 int16 = -32768
                _ = v4
            case "\n\t":
                fmt.Println(v1)
            case "hello":
                v1 -= (v1 / ((v1) | 1))
            default:
                var v5 int32 = 418
                _ = v5
            }
            v1 += 2
        } else {
            var v6 uint8 = 42
            _ = v6
        }
        v7 := GenS1{F0: 127}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v7
    } else {
        var v8 int32 = 0
        v9 := GenS1{F0: 42}
        v1 -= (v1 * (v1 ^ v1))
        _ = v8
        _ = v9
    }
    if true {
        if (29 != 2147483647) {
            if (0 == 32767) {
                var v10 uint32 = 4294967295
                switch 756 {
                case 127:
                    v3[-2147483648] = true
                case -959:
                    switch ("a" + "a") {
                    case "hello":
                        if ((!("" > ("a" + "a"))) && true) {
                            var v11 float64 = 1e10
                            _ = v11
                        } else {
                            v10 += (v10 | v10)
                            v10 -= v10
                            fmt.Println(v10, v1)
                        }
                    }
                case 42:
                    var v12 float64 = 0.0
                    _ = v12
                }
                v10 += v10
                _ = v10
            }
            var v13 int16 = 1000
            var v14 float64 = float64((0.1 + (float32(3.14159) * 0.1)))
            _ = v13
            _ = v14
        }
    } else {
        v1 -= 42
    }
    v1 -= 625
    var v15 int16 = -32768
    v1 = (v1 + 1)
    fmt.Println(v15, v1)
    v16 := GenS1{F0: -80}
    v17 := GenS1{F0: 0}
    v2 = append(v2, ("foo bar" + "foo bar"))
    switch (((("" + "hello") + "\n\t") + "\n\t") + "hello") {
    case "\n\t":
        switch 127 {
        case 703:
            var v18 int32 = 0
            _ = v18
        default:
            fmt.Println(v15)
        }
    case "":
        v15 += v15
    default:
        if false {
            v15 = (v15 * 793)
            var v19 int32 = -1
            _ = v19
        }
    }
    v15 = v15
    if true {
        v20 := GenS1{F0: 42}
        fmt.Println(v1, v1)
        _ = v20
    }
    var v21 uint16 = 717
    v22 := GenS1{F0: -73}
    v1 -= v1
    fmt.Println(v1)
    var v23 bool = ((!(false && (42 == -317))) && (0 >= 756))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int16 }
func genFn0() uint16 {
    v5 := GenS1{F0: 422, F1: -32768, F2: -32768}
    if (0.1 <= (0.0 - 0.0)) {
        for fi6 := 0; fi6 < 10; fi6++ {
            var v7 bool = true
            var v8 int = fi6
            _ = fi6
            _ = v7
            _ = v8
        }
        v9 := []int{-9223372036854775808, -179, -13, -785}
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = v9
    } else {
        var v10 uint16 = 42
        _ = v10
    }
    _ = v5
    return 1
}

func main() {
    v1 := []string{"hello"}
    var v2 uint16 = 65535
    var v3 bool = (((533 == 255) && (2 < 787)) && (!((false || true) && (100 <= 42))))
    var v4 int8 = 1
    for fi11 := 0; fi11 < 8; fi11++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi11
    }
    switch -9223372036854775808 {
    case 42:
        switch ((("foo bar" + "a") + "foo bar") + ("hello" + "\n\t")) {
        case "":
            v3 = v3
        case "foo bar":
            var v12 int8 = (-(v4))
            _ = v12
        }
    case 32768:
        v4 = 88
    default:
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 float64 = (float64((3.14159 - (-2.25 - 0.0))) / 0.1)
            var v15 uint = 645
            var v16 uint8 = 42
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
    }
    fmt.Println(v2, v4, v3)
    var v17 float32 = 1e10
    v3 = v3
    var v18 float64 = -2.25
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    if v3 {
        if (v3 && (!(true || (!true)))) {
            v17 -= 1.5
            var v19 int = 97
            _ = v19
        }
    } else {
        v1 = append(v1, "foo bar")
        switch 411 {
        case 32767:
            for fi20 := 0; fi20 < 4; fi20++ {
                v21 := GenS1{F0: 1, F1: 32767, F2: 1000}
                switch (fi20 / ((fi20) | 1)) {
                case -1:
                    v4 = -14
                case 76:
                    var v22 int = (fi20 / ((9223372036854775807) | 1))
                    _ = v22
                }
                _ = fi20
                _ = v21
            }
        case -434:
            var v23 int8 = v4
            _ = v23
        case 1:
            if v3 {
                v1 = append(v1, "foo bar")
                fmt.Println(v18, v2, v4)
                if (!v3) {
                    var v24 uint8 = 136
                    _ = v24
                } else {
                    v25 := GenS1{F0: -396, F1: 1, F2: 1000}
                    var v26 uint64 = 1
                    _ = v25
                    _ = v26
                }
            }
        default:
            fmt.Println(v3, v18, v18)
        }
    }
    var v27 uint32 = 75
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}

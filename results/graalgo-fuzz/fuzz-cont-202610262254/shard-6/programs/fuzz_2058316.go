package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
func genFn0(p0 int32, p1 uint, p2 uint) uint8 {
    p2 = p2
    if (!true) {
        for fi4 := 0; fi4 < 18; fi4++ {
            var v5 uint32 = 42
            switch 4294967295 {
            case 32768:
                var v6 uint16 = 255
                _ = v6
            case 9223372036854775807:
                var v7 uint64 = 2
                _ = v7
            default:
                v8 := GenS1{F0: -1, F1: "foo bar"}
                _ = v8
            }
            var v9 uint = p1
            _ = fi4
            _ = v5
            _ = v9
        }
        v10 := GenS1{F0: 9223372036854775807, F1: (("" + ((("" + ((("a" + "a") + "\n\t") + "")) + "hello") + ("\n\t" + ""))) + "a")}
        _ = v10
    } else {
        var v11 uint16 = 255
        p0 = 1
        _ = v11
    }
    var v12 uint = (p2 >> ((p1) & 63))
    v13 := []int{-171, -73}
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 uint16 = 42
        var v16 int32 = p0
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 134
}

func main() {
    var v1 uint = 620
    v2 := []int{-9223372036854775808, 32768}
    var v3 float64 = ((((0.0 / 1.0) / 1.5) - 3.14159) / float64((float32(-2.25) * 3.14159)))
    var v17 int = 3
    if len(v2) > 0 {
        v2[0] = v17
    }
    if (128 != (v17 & v17)) {
        for fi18 := 0; fi18 < 18; fi18++ {
            if (1 > 1) {
                fmt.Println(v3, v1)
                for fi19 := 0; fi19 < 6; fi19++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v3 = v3
                    _ = fi19
                }
            }
            _ = fi18
        }
        var v20 int16 = -1
        _ = v20
    } else {
        _ = len(v2)
        fmt.Println(v3, v1, v1)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 += 0.1
    var v21 int16 = 42
    var v22 uint64 = 97
    v23 := GenS1{F0: 443, F1: "foo bar"}
    v1 += v1
    if (-676 != 0) {
        var v24 int16 = (v21 * (v21 ^ 723))
        var v25 int64 = -971
        fmt.Println(v1, v1, v21)
        _ = v24
        _ = v25
    }
    fmt.Println(v22)
    for fi26 := 0; fi26 < 11; fi26++ {
        v22 -= 2
        _ = fi26
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        switch "" {
        case "hello":
            fmt.Println(v21)
        case "":
            fmt.Println(v22, v17)
        case "a":
            for fi28 := 0; fi28 < 4; fi28++ {
                var v29 int32 = 1
                if ((("foo bar" + ("a" + ("\n\t" + "foo bar"))) >= "hello") || ((v22 + 1) > 870)) {
                    var v30 int8 = -21
                    _ = v30
                }
                v1 = (^(v1))
                _ = fi28
                _ = v29
            }
        }
        _ = fi27
    }
    fmt.Println(v17, v21)
    v1 -= (v1 * 381)
    var v31 uint8 = 255
    v21 -= v21
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v31
}

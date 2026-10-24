package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int) uint64 {
    switch ("hello" + "\n\t") {
    case "foo bar":
        var v6 int32 = -1
        _ = v6
    case "a":
        var v7 bool = true
        _ = v7
    }
    p0 -= (^(p0))
    p0 += (p0 << ((316) & 63))
    switch p0 {
    case 32767:
        switch "hello" {
        case "hello":
            v8 := GenS1{F0: ("a" + "hello")}
            _ = v8
        }
    case 9223372036854775807:
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 int = (p0 * fi9)
            _ = fi9
            _ = v10
        }
    case 120:
        p0 -= (p0 << ((18446744073709551615) & 63))
    default:
        var v11 float64 = 3.14159
        _ = v11
    }
    if false {
        if (9223372036854775807 != 0) {
            for fi12 := 0; fi12 < 5; fi12++ {
                v13 := GenS1{F0: "\n\t"}
                _ = fi12
                _ = v13
            }
        } else {
            p0 -= p0
        }
        fmt.Println(p0)
        var v14 bool = (76 >= 255)
        _ = v14
    } else {
        if ((0 >= 571) || true) {
            v15 := []int{-608, (p0 | (p0 - p0)), p0}
            _ = v15
        }
        for fi16 := 0; fi16 < 5; fi16++ {
            if (1 < 1) {
                var v17 int = (fi16 >> ((65535) & 63))
                v18 := []int{(p0 * (p0 - v17)), p0, (-(v17)), 42}
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
                _ = v17
                _ = v18
            }
            _ = fi16
        }
    }
    _ = p0
    return 126
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[string]int)
    var v3 int64 = 372
    v4 := []uint32{1, 319, 368}
    var v5 float64 = float64(0.1)
    v1 -= v1
    if ("foo bar" <= "") {
        var v19 uint16 = 255
        _ = v19
    }
    switch 4294967295 {
    case 127:
        var v20 int64 = 42
        _ = v20
    case -2147483648:
        var v21 uint = 776
        _ = v21
    case 42:
        var v22 int8 = 1
        _ = v22
    default:
        var v23 int16 = 42
        _ = v23
    }
    if false {
        var v24 uint64 = 993
        var v25 uint32 = 4294967295
        _ = v24
        _ = v25
    } else {
        v5 += (v5 + (v5 * v5))
    }
    switch -453 {
    case -845:
        _ = len(v2)
    default:
        v1 = v1
    }
    if (-128 < 1) {
        _ = len(v2)
        v5 += (v5 + 1.5)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        v1 = (v1 >> ((4294967295) & 63))
    }
    for fi26 := 0; fi26 < 11; fi26++ {
        v27 := GenS1{F0: "foo bar"}
        var v28 int = 9223372036854775807
        _ = fi26
        _ = v27
        _ = v28
    }
    v5 = (float64(0.1) * 0.1)
    for fi29 := 0; fi29 < 7; fi29++ {
        for fi30 := 0; fi30 < 7; fi30++ {
            switch (("\n\t" + ("\n\t" + "\n\t")) + ("" + "a")) {
            case "hello":
                var v31 string = ("hello" + "foo bar")
                _ = v31
            }
            fmt.Println(fi30)
            v1 += v1
            _ = fi30
        }
        v5 = v5
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

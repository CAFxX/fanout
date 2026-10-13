package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int8 }
func genFn0(p0 int8) uint {
    var v4 float64 = (float64(3.14159) + (1e10 - 1e10))
    switch "foo bar" {
    case "":
        var v5 string = ("hello" + (("\n\t" + "hello") + (("" + "") + "a")))
        _ = v5
    case "\n\t":
        var v6 float64 = float64(float32(v4))
        _ = v6
    }
    _ = p0
    _ = v4
    return 993
}

func main() {
    var v1 int = -299
    v2 := []uint32{}
    var v3 int32 = -2147483648
    v2 = append(v2, 4294967295)
    switch (v1 << ((18446744073709551615) & 63)) {
    case 1:
        if ("\n\t" <= "") {
            v2 = append(v2, 0)
            var v7 uint = 423
            v1 -= (v1 - -1)
            _ = v7
        } else {
            if false {
                var v8 uint64 = 145
                _ = v8
            } else {
                switch 610 {
                case 127:
                    v3 -= v3
                case 9223372036854775807:
                    for fi9 := 0; fi9 < 14; fi9++ {
                        var v10 int = (^(v1))
                        var v11 uint = 18446744073709551615
                        var v12 int64 = 971
                        _ = fi9
                        _ = v10
                        _ = v11
                        _ = v12
                    }
                default:
                    v1 += 42
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v13 := GenS1{F0: 778, F1: 2, F2: 100}
                _ = v13
            }
            v2 = append(v2, 2147483647)
        }
    case -9223372036854775808:
        fmt.Println(v1)
    case 128:
        fmt.Println(v3)
    default:
        switch "" {
        case "hello":
            v3 = v3
        case "":
            var v14 string = "a"
            _ = v14
        case "foo bar":
            for fi15 := 0; fi15 < 4; fi15++ {
                var v16 bool = true
                _ = fi15
                _ = v16
            }
        }
    }
    v3 = v3
    if true {
        var v17 uint32 = 2147483648
        var v18 float64 = 1.5
        switch (("a" + "") + "") {
        case "\n\t":
            var v19 uint64 = 428
            _ = v19
        case "hello":
            v18 += (-2.25 - (v18 + -2.25))
        }
        _ = v17
        _ = v18
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        v21 := GenS1{F0: 1, F1: 265, F2: 100}
        _ = fi20
        _ = v21
    }
    v22 := GenS1{F0: genFn0(75), F1: 42, F2: 74}
    var v23 int32 = (v3 >> ((18446744073709551615) & 63))
    v23 -= 0
    if (42 != int16(v23)) {
        if true {
            if (42 >= 65535) {
                fmt.Println(v3, v3, v1)
            }
            var v24 int8 = -1
            var v25 uint64 = 474
            _ = v24
            _ = v25
        }
        for fi26 := 0; fi26 < 4; fi26++ {
            fmt.Println(fi26)
            v27 := GenS1{F0: 0, F1: 0, F2: int8(fi26)}
            for fi28 := 0; fi28 < 10; fi28++ {
                var v29 int16 = 1000
                v30 := GenS1{F0: 0, F1: 4294967295, F2: 100}
                v3 = v23
                v29 += v29
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = fi26
            _ = v27
        }
        fmt.Println(v3, v23)
    } else {
        if (0 > 1) {
            switch ("a" + ("foo bar" + ("hello" + "\n\t"))) {
            case "\n\t":
                switch (v1 << ((2) & 63)) {
                case 4294967295:
                    if len(v2) > 0 {
                        v2[0] = 2147483648
                    }
                case 127:
                    fmt.Println(v1, v1, v1)
                }
            case "hello":
                v31 := GenS1{F0: 4294967295, F1: genFn0(42), F2: 0}
                _ = v31
            }
        }
    }
    if false {
        var v32 uint = genFn0(-1)
        var v33 bool = (!(!(128 >= 81)))
        var v34 float32 = 1.5
        _ = v32
        _ = v33
        _ = v34
    }
    _ = len(v2)
    var v35 int8 = -88
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v35
}

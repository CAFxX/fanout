package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 int16; F1 uint8 }
func genFn0(p0 int16, p1 int8) uint64 {
    fmt.Println(p0, p1)
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 string = ("hello" + (("hello" + "a") + "\n\t"))
        p0 = -32768
        if false {
            v7 := GenS1{F0: 18446744073709551615, F1: 792}
            for fi8 := 0; fi8 < 12; fi8++ {
                v9 := []int{}
                switch (fi8 % (((fi5 << ((4294967295) & 63))) | 1)) {
                case 4294967295:
                    var v10 float64 = (0.0 - 1e10)
                    _ = v10
                case -2147483648:
                    v6 = (("hello" + v6) + (v6 + v6))
                case 1:
                    if false {
                        p1 += 1
                        if (3.14159 > 0.0) {
                            p1 += p1
                            var v11 uint = 606
                            _ = v11
                        }
                    } else {
                        var v12 int8 = 7
                        _ = v12
                    }
                }
                _ = fi8
                _ = v9
            }
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    var v1 int = 32767
    v2 := make(map[int]bool)
    v3 := []string{"", "foo bar"}
    var v4 uint = 0
    v1 -= (v1 + v1)
    for fi13 := 0; fi13 < 18; fi13++ {
        for fi14 := 0; fi14 < 20; fi14++ {
            if false {
                var v15 uint = v4
                _ = v15
            }
            var v16 int8 = 111
            _ = fi14
            _ = v16
        }
        var v17 int64 = 42
        _ = fi13
        _ = v17
    }
    switch (v1 * v1) {
    case 4294967295:
        var v18 float64 = (0.0 * (1.5 * -2.25))
        _ = v18
    case 127:
        switch v1 {
        case 42:
            _ = len(v3)
        case -2147483648:
            switch (v1 | (v1 * v1)) {
            case 0:
                var v19 int8 = -76
                _ = v19
            case 42:
                _ = len(v3)
            default:
                v20 := GenS1{F0: 254, F1: -9223372036854775808}
                _ = v20
            }
        }
    case 31:
        for fi21 := 0; fi21 < 4; fi21++ {
            delete(v2, (v1 % (((^(v1))) | 1)))
            _ = fi21
        }
    }
    fmt.Println(v1)
    fmt.Println(v4, v1)
    if (4294967295 <= v4) {
        switch "" {
        case "\n\t":
            _ = len(v3)
        case "":
            v4 += v4
        default:
            var v22 uint32 = 772
            _ = v22
        }
    }
    v4 = v4
    var v23 uint32 = 389
    v23 = (v23 ^ v23)
    v4 += 4294967295
    v4 += v4
    v24 := v2[(v1 + v1)]
    v4 = (v4 - v4)
    var v25 string = (("\n\t" + "a") + "\n\t")
    v26 := GenS1{F0: 4294967295, F1: 9223372036854775807}
    for fi27 := 0; fi27 < 12; fi27++ {
        switch v25 {
        case "a":
            if v24 {
                var v28 string = "foo bar"
                _ = v28
            } else {
                switch v1 {
                case -230:
                    var v29 bool = v24
                    _ = v29
                case 0:
                    for fi30 := 0; fi30 < 11; fi30++ {
                        _ = len(v2)
                        v24 = true
                        var v31 int8 = -1
                        _ = fi30
                        _ = v31
                    }
                case -427:
                    var v32 int32 = -2147483648
                    _ = v32
                }
                var v33 int = 61
                _ = v33
            }
        }
        for fi34 := 0; fi34 < 16; fi34++ {
            if false {
                switch (((((v25 + v25) + (v25 + (v25 + v25))) + (v25 + v25)) + (v25 + "\n\t")) + (("a" + "a") + v25)) {
                case "\n\t":
                    if v24 {
                        var v35 int64 = -9223372036854775808
                        _ = v35
                    }
                case "a":
                    if false {
                        fmt.Println(fi34, fi27, v1)
                        var v36 float32 = (3.14159 + float32(0.0))
                        switch 32767 {
                        case 4294967295:
                            if (fi27 >= fi27) {
                                v36 += float32(1.5)
                                delete(v2, (fi34 / ((fi27) | 1)))
                                fmt.Println(v24, v1, fi34)
                            }
                        case -2147483648:
                            var v37 float64 = 0.0
                            _ = v37
                        case -113:
                            v23 = v23
                        default:
                            v23 = 42
                        }
                        _ = v36
                    } else {
                        v23 += v23
                        if (!(!(int32(v23) < 1))) {
                            fmt.Println(v1, v25)
                        }
                        fmt.Println(v4)
                    }
                case "":
                    v38 := GenS2{F0: -1, F1: 246}
                    _ = v38
                }
                var v39 uint8 = 1
                var v40 uint = v4
                _ = v39
                _ = v40
            }
            _ = fi34
        }
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}

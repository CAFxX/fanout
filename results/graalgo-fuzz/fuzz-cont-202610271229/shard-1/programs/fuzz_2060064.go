package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int64 }
type GenS2 struct { F0 uint; F1 int; F2 int8 }

func main() {
    var v1 bool = true
    var v2 bool = v1
    var v3 uint32 = 2
    v4 := []int{-14, 1, 1}
    v5 := []uint32{}
    var v6 float64 = 0.1
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := GenS1{F0: v3, F1: 127, F2: -321}
        fmt.Println(v6, v3)
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 int8 = 0
            switch (fi7 >> ((65535) & 63)) {
            case 23:
                v11 := GenS2{F0: 65535, F1: 42, F2: (v10 - 1)}
                _ = v11
            case 0:
                if (true || (702 == -1)) {
                    var v12 uint64 = 580
                    var v13 uint32 = 713
                    _ = v12
                    _ = v13
                }
            case 2147483647:
                if false {
                    if (v2 || (166 <= 1)) {
                        v14 := GenS1{F0: v3, F1: (v10 << ((80) & 63)), F2: -535}
                        _ = v14
                    }
                    v10 -= 95
                    if v1 {
                        v10 = 127
                        v3 -= (v3 * v3)
                        v2 = ((("\n\t" + ("\n\t" + "foo bar")) < ("\n\t" + ("" + ("hello" + "\n\t")))) && false)
                    }
                } else {
                    v15 := GenS1{F0: v3, F1: (v10 >> ((354) & 63)), F2: -9223372036854775808}
                    switch (fi7 % ((fi7) | 1)) {
                    case -22:
                        var v16 uint16 = 2
                        _ = v16
                    case 2147483647:
                        var v17 uint16 = 2
                        _ = v17
                    default:
                        var v18 uint16 = 575
                        _ = v18
                    }
                    fmt.Println(v1, v6)
                    _ = v15
                }
            }
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    var v19 uint16 = 255
    fmt.Println(v3, v2, v2)
    for fi20 := 0; fi20 < 12; fi20++ {
        v21 := GenS1{F0: 2147483647, F1: 1, F2: -1}
        _ = fi20
        _ = v21
    }
    switch 4294967295 {
    case 42:
        var v22 uint32 = v3
        _ = v22
    case 705:
        switch 4294967295 {
        case 1:
            var v23 uint8 = 0
            _ = v23
        case 32767:
            var v24 uint16 = (^(v19))
            _ = v24
        }
    default:
        var v25 uint16 = v19
        _ = v25
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        var v27 int = fi26
        v28 := GenS1{F0: (v3 * v3), F1: -1, F2: -974}
        v29 := GenS1{F0: v3, F1: -84, F2: -871}
        _ = fi26
        _ = v27
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 2; fi30++ {
        var v31 int16 = 32767
        _ = fi30
        _ = v31
    }
    fmt.Println(v6, v2, v6)
    v3 = 515
    if v2 {
        fmt.Println(v2, v3)
        for fi32 := 0; fi32 < 16; fi32++ {
            var v33 uint32 = v3
            _ = fi32
            _ = v33
        }
    } else {
        v3 = (v3 >> ((4294967295) & 63))
        var v34 float64 = ((3.14159 * v6) * 1e10)
        _ = v34
    }
    for fi35 := 0; fi35 < 2; fi35++ {
        v4 = append(v4, fi35)
        _ = fi35
    }
    for fi36 := 0; fi36 < 18; fi36++ {
        if v2 {
            var v37 int = (fi36 ^ 1)
            switch "a" {
            case "hello":
                var v38 uint8 = 1
                _ = v38
            case "foo bar":
                if ((!false) || true) {
                    fmt.Println(v3, v19, v37)
                    v3 = v3
                    fmt.Println(v37)
                }
            case "":
                v2 = true
            default:
                var v39 int32 = -536
                _ = v39
            }
            _ = v37
        }
        if v1 {
            for fi40 := 0; fi40 < 20; fi40++ {
                if (128 < 255) {
                    switch (fi40 * (fi36 >> ((947) & 63))) {
                    case 42:
                        v6 -= 1e10
                    case 32767:
                        var v41 int32 = -1
                        _ = v41
                    case 4294967295:
                        fmt.Println(v1, v19, v6)
                    default:
                        fmt.Println(fi36, v3, v19)
                    }
                    v42 := GenS1{F0: v3, F1: -1, F2: 1}
                    _ = v42
                } else {
                    var v43 uint = 4294967295
                    _ = v43
                }
                v44 := GenS2{F0: 255, F1: fi36, F2: 100}
                _ = fi40
                _ = v44
            }
        }
        _ = fi36
    }
    var v45 float32 = (-2.25 / (0.1 - ((((0.1 - 0.1) + 3.14159) / (((0.0 * 1e10) - (0.1 * (0.0 / 1.5))) - (3.14159 * (1e10 + 1.5)))) * 0.0)))
    var v46 float64 = float64((1.5 - (-2.25 + v45)))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}

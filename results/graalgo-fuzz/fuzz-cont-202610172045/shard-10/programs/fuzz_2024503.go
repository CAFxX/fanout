package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint }

func main() {
    v1 := []int{-9223372036854775808, -2147483648, -2147483648}
    v2 := make(map[uint32]bool)
    var v3 string = ("a" + "\n\t")
    var v4 uint32 = 65535
    if (1 > 42) {
        fmt.Println(v4)
        v4 += (v4 + (v4 - 139))
        var v5 float32 = -2.25
        _ = v5
    }
    fmt.Println(v4)
    v3 = v3
    var v6 int = 42
    switch v6 {
    case -517:
        v2[2147483647] = true
    case 9223372036854775807:
        v4 = 515
    default:
        _ = len(v1)
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        fmt.Println(v4, fi7, fi7)
        _ = fi7
    }
    v3 = (v3 + ((v3 + v3) + ("foo bar" + v3)))
    if (!(4294967295 < 18446744073709551615)) {
        var v8 uint16 = 2
        if ((v4 != (v4 >> ((0) & 63))) && (!(-128 < 42))) {
            v8 = (v8 / ((42) | 1))
            var v9 float64 = (3.14159 * 1.5)
            v4 += (v4 << ((uint(v8)) & 63))
            _ = v9
        }
        var v10 uint32 = (v4 << ((65535) & 63))
        _ = v10
        _ = v8
    }
    if true {
        fmt.Println(v3, v3)
        if (true && true) {
            _ = len(v1)
        }
        if (915 < 747) {
            var v11 uint32 = v4
            v4 = (^(v4))
            _ = len(v2)
            _ = v11
        } else {
            if false {
                var v12 int64 = 932
                var v13 int32 = -1
                _ = v12
                _ = v13
            } else {
                for fi14 := 0; fi14 < 9; fi14++ {
                    if false {
                        v3 = v3
                    }
                    if (!(false && (4294967295 == uint64(v4)))) {
                        v15 := GenS1{F0: 688, F1: ""}
                        v4 = (v4 - v4)
                        v1 = append(v1, fi14)
                        _ = v15
                    } else {
                        var v16 float32 = (float32(0.1) * (3.14159 / (0.0 - -2.25)))
                        _ = v16
                    }
                    switch fi14 {
                    case -9223372036854775808:
                        var v17 int16 = 32767
                        _ = v17
                    case 32767:
                        fmt.Println(v4, fi14)
                    }
                    _ = fi14
                }
                v18 := v2[(v4 - 2147483647)]
                _ = v18
            }
            var v19 uint32 = v4
            for fi20 := 0; fi20 < 1; fi20++ {
                var v21 int8 = 42
                _ = fi20
                _ = v21
            }
            _ = v19
        }
    } else {
        for fi22 := 0; fi22 < 17; fi22++ {
            fmt.Println(fi22)
            switch v3 {
            case "\n\t":
                var v23 string = (v3 + (v3 + (v3 + ("" + ("foo bar" + ((v3 + v3) + "foo bar"))))))
                _ = v23
            case "a":
                fmt.Println(fi22, fi22)
            case "foo bar":
                v24 := GenS2{F0: 65535}
                _ = v24
            default:
                for fi25 := 0; fi25 < 20; fi25++ {
                    v4 -= (^(v4))
                    fmt.Println(v6, fi25, fi22)
                    _ = fi25
                }
            }
            switch (v6 % ((fi22) | 1)) {
            case 32768:
                if (!true) {
                    for fi26 := 0; fi26 < 12; fi26++ {
                        var v27 int32 = 118
                        v28 := GenS1{F0: 255, F1: v3}
                        var v29 int16 = -32768
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                } else {
                    switch (v6 + (v6 | fi22)) {
                    case -9223372036854775808:
                        _ = len(v1)
                    case 760:
                        var v30 int16 = 32767
                        _ = v30
                    default:
                        v4 = (v4 - (v4 | 2147483648))
                    }
                    var v31 uint64 = 4294967295
                    _ = v31
                }
            }
            _ = fi22
        }
        for fi32 := 0; fi32 < 13; fi32++ {
            v33 := GenS1{F0: 42, F1: v3}
            fmt.Println(v3, v3)
            v34 := GenS2{F0: 0}
            _ = fi32
            _ = v33
            _ = v34
        }
    }
    if ((v6 + v6) < v6) {
        var v35 uint32 = 65535
        fmt.Println(v35)
        if ((!true) || true) {
            var v36 string = (v3 + v3)
            _ = v36
        } else {
            var v37 float64 = 0.0
            _ = v37
        }
        _ = v35
    } else {
        var v38 bool = false
        var v39 int = 27
        _ = v38
        _ = v39
    }
    if (3.14159 != 0.1) {
        var v40 float32 = (1e10 - 0.0)
        var v41 int16 = -43
        _ = v40
        _ = v41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}

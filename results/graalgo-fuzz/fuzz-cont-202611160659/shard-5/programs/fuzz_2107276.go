package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint; F1 int; F2 bool }
func genFn0(p0 int16, p1 string) uint8 {
    p1 = "a"
    if (!(161 >= 0)) {
        var v7 int = 2147483647
        switch "" {
        case "hello":
            switch "\n\t" {
            case "foo bar":
                v8 := []int{-719}
                _ = v8
            case "\n\t":
                var v9 uint = 573
                _ = v9
            case "hello":
                p1 = (p1 + p1)
            default:
                if false {
                    switch 127 {
                    case 4294967295:
                        p1 = (p1 + ("\n\t" + (p1 + "\n\t")))
                    }
                    var v10 int8 = -128
                    v10 = (-(v10))
                    _ = v10
                }
            }
        case "\n\t":
            switch ((p1 + (p1 + ("foo bar" + ("\n\t" + "\n\t")))) + p1) {
            case "\n\t":
                v11 := GenS1{F0: (true && true), F1: 4294967295}
                _ = v11
            case "":
                if true {
                    v12 := []int{(^(v7)), (-(v7)), (v7 + v7), -2147483648}
                    var v13 string = p1
                    v14 := GenS2{F0: 709, F1: (-(v7)), F2: (127 <= 127)}
                    _ = v12
                    _ = v13
                    _ = v14
                }
            case "hello":
                fmt.Println(p0, p1, p1)
            default:
                v15 := GenS1{F0: false, F1: 1}
                _ = v15
            }
        case "foo bar":
            v7 = (v7 << ((18446744073709551615) & 63))
        }
        v7 = v7
        _ = v7
    }
    _ = p0
    _ = p1
    return 255
}
func genFn1() bool {
    var v16 float64 = 0.1
    if (127 > 100) {
        var v17 string = ("\n\t" + "a")
        if false {
            switch v17 {
            case "":
                switch v17 {
                case "a":
                    var v18 int8 = -97
                    _ = v18
                case "\n\t":
                    var v19 uint = 0
                    _ = v19
                default:
                    fmt.Println(v16, v17)
                }
            case "\n\t":
                v20 := []int{203, 543}
                _ = v20
            case "foo bar":
                v16 = 1e10
            }
            v17 = ""
        }
        if (!true) {
            v21 := []int{-1, 1, 1, 2147483647}
            fmt.Println(v16, v17)
            _ = v21
        } else {
            for fi22 := 0; fi22 < 1; fi22++ {
                v16 -= (-2.25 / v16)
                _ = fi22
            }
            var v23 bool = true
            _ = v23
        }
        _ = v17
    }
    var v24 bool = true
    _ = v16
    _ = v24
    return (0 < 32767)
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 42
    v3 := make(map[uint32]string)
    var v4 bool = (-9223372036854775808 != 9223372036854775807)
    v5 := []string{(("foo bar" + ("" + ("\n\t" + "\n\t"))) + "\n\t"), "hello", "\n\t", ""}
    var v6 int8 = (v2 ^ v2)
    for fi25 := 0; fi25 < 7; fi25++ {
        switch (("a" + ("\n\t" + ("" + "a"))) + ("foo bar" + "")) {
        case "":
            for fi26 := 0; fi26 < 12; fi26++ {
                delete(v3, 48)
                fmt.Println(v4)
                _ = fi26
            }
        default:
            var v27 int32 = 1
            _ = v27
        }
        var v28 uint64 = 2
        _ = fi25
        _ = v28
    }
    _ = len(v5)
    var v29 int16 = -436
    if v4 {
        var v30 uint16 = 0
        _ = v30
    } else {
        v4 = (-372 > 42)
    }
    v5 = append(v5, "\n\t")
    if v4 {
        v6 -= 94
    }
    _ = len(v1)
    fmt.Println(v29, v2, v29)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

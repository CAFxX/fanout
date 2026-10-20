package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 string }
func genFn0(p0 int16, p1 string, p2 uint) bool {
    switch 128 {
    case 0:
        if (!true) {
            if true {
                var v4 int16 = (^(p0))
                _ = v4
            } else {
                switch 32768 {
                case -1:
                    v5 := []int{4294967295, 97, 32767, 9223372036854775807}
                    _ = v5
                default:
                    fmt.Println(p0)
                }
                p0 = p0
                v6 := GenS1{F0: (p0 % ((1) | 1))}
                _ = v6
            }
        }
    case 32767:
        v7 := GenS1{F0: p0}
        _ = v7
    case 1:
        var v8 int8 = 7
        _ = v8
    }
    v9 := GenS1{F0: (p0 ^ 1)}
    fmt.Println(p2)
    p1 = p1
    v10 := GenS1{F0: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return false
}

func main() {
    var v1 bool = (-1 >= -1)
    var v2 int = 32767
    v3 := make(map[string]bool)
    switch v2 {
    case -2147483648:
        if (("\n\t" < "\n\t") && v1) {
            v2 = v2
            var v11 int64 = 9223372036854775807
            _ = v11
        }
    case 186:
        switch (v2 * v2) {
        case -343:
            switch ("a" + (("foo bar" + "\n\t") + ("a" + (("\n\t" + "a") + "foo bar")))) {
            case "":
                v12 := v3["a"]
                _ = v12
            }
        case 2147483647:
            delete(v3, "a")
        case -278:
            switch (v2 & (v2 - (v2 >> ((2) & 63)))) {
            case 893:
                delete(v3, ("hello" + (("a" + (("foo bar" + ("hello" + "hello")) + "")) + ((("hello" + "\n\t") + "foo bar") + "\n\t"))))
            case 32768:
                fmt.Println(v2, v1)
            case 0:
                var v13 uint16 = 1
                _ = v13
            }
        default:
            delete(v3, "a")
        }
    default:
        var v14 uint64 = 1
        _ = v14
    }
    var v15 bool = true
    if v1 {
        if (65535 >= 952) {
            switch "" {
            case "hello":
                switch "hello" {
                case "\n\t":
                    for fi16 := 0; fi16 < 4; fi16++ {
                        v17 := []int{32767, (v2 & (^(fi16))), (fi16 + (fi16 | (v2 << ((394) & 63)))), 32768}
                        v2 = (fi16 | fi16)
                        _ = fi16
                        _ = v17
                    }
                case "hello":
                    v1 = (!(3.14159 <= 1e10))
                case "a":
                    if v1 {
                        if false {
                            var v18 int8 = -1
                            v1 = v15
                            _ = len(v3)
                            _ = v18
                        }
                    } else {
                        fmt.Println(v1)
                    }
                }
            case "":
                delete(v3, "")
            default:
                var v19 int16 = 42
                _ = v19
            }
            for fi20 := 0; fi20 < 5; fi20++ {
                fmt.Println(fi20, v1, v15)
                var v21 float64 = (0.0 * -2.25)
                var v22 uint64 = 42
                _ = fi20
                _ = v21
                _ = v22
            }
            var v23 float32 = -2.25
            _ = v23
        }
    }
    var v24 float64 = 0.1
    var v25 uint = 42
    var v26 float32 = float32(v24)
    fmt.Println(v2, v2)
    var v27 uint64 = 2
    for fi28 := 0; fi28 < 2; fi28++ {
        v29 := []int{int(v27)}
        if (!((fi28 + (^(fi28))) > fi28)) {
            v30 := GenS1{F0: -540}
            for fi31 := 0; fi31 < 3; fi31++ {
                if true {
                    v2 -= (-(fi28))
                    var v32 uint16 = 344
                    v1 = (1e10 >= 3.14159)
                    var v33 float32 = (((v26 * (v26 * float32((float64(v26) * float64((((float32(float64((float32(v24) / v26))) * 3.14159) + (-2.25 + v26)) / ((v26 * v26) - v26))))))) / v26) / v26)
                    _ = v32
                    _ = v33
                }
                var v34 uint64 = 42
                v35 := GenS1{F0: 56}
                fmt.Println(v15, v1, v1)
                if (62 != -128) {
                    v2 = fi28
                } else {
                    var v36 string = "\n\t"
                    v37 := v3[""]
                    _ = v36
                    _ = v37
                }
                switch ("\n\t" + (("hello" + (("a" + ("foo bar" + ("" + ""))) + (("hello" + ("a" + "hello")) + ("a" + (("hello" + "foo bar") + ""))))) + "\n\t")) {
                case "\n\t":
                    fmt.Println(v24)
                }
                _ = fi31
                _ = v34
                _ = v35
            }
            _ = v30
        }
        _ = fi28
        _ = v29
    }
    fmt.Println(v15, v27, v1)
    v38 := GenS2{F0: 5, F1: ("\n\t" + ("hello" + "hello"))}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v38
}

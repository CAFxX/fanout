package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint64; F1 uint16 }
func genFn0(p0 uint) int32 {
    var v5 int8 = 1
    fmt.Println(v5, v5)
    fmt.Println(v5)
    _ = p0
    _ = v5
    return -2147483648
}
func genFn1(p0 bool, p1 string) uint {
    v6 := GenS1{F0: p1, F1: 574, F2: -2147483648}
    p0 = (p0 && (!((-128 == 9) || (75 <= -1))))
    var v7 int = 42
    var v8 float64 = 0.0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 248
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{-1, 128, 32768, -1}
    var v3 string = "\n\t"
    var v4 bool = (false && false)
    v4 = true
    v1[-1] = 74
    v4 = true
    v4 = (0 != genFn1((v4 && v4), ((v3 + ((v3 + (((("\n\t" + v3) + v3) + (v3 + ((v3 + "") + "\n\t"))) + v3)) + v3)) + "")))
    if v4 {
        fmt.Println(v3, v3)
        v3 = v3
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        for fi10 := 0; fi10 < 16; fi10++ {
            switch (fi9 - (fi9 % (((fi9 * 551)) | 1))) {
            case 0:
                _ = len(v2)
            default:
                v3 = ""
            }
            _ = fi10
        }
        v3 = v3
        _ = fi9
    }
    if v4 {
        v11 := GenS1{F0: v3, F1: 2, F2: 958}
        switch 1 {
        case 4294967295:
            v12 := GenS1{F0: "", F1: 42, F2: genFn0(147)}
            _ = v12
        case -256:
            v4 = true
        default:
            if (!(0 <= 0)) {
                switch (("foo bar" + "") + "foo bar") {
                case "hello":
                    for fi13 := 0; fi13 < 10; fi13++ {
                        for fi14 := 0; fi14 < 8; fi14++ {
                            var v15 uint8 = 47
                            _ = fi14
                            _ = v15
                        }
                        for fi16 := 0; fi16 < 10; fi16++ {
                            switch (((v3 + v3) + (((v3 + (v3 + (v3 + (v3 + (v3 + ("foo bar" + (("" + "\n\t") + "hello"))))))) + (v3 + "a")) + v3)) + (v3 + (v3 + v3))) {
                            case "\n\t":
                                v3 = v3
                            case "":
                                fmt.Println(v3)
                            case "a":
                                v3 = v3
                            }
                            fmt.Println(v3)
                            _ = fi16
                        }
                        _ = fi13
                    }
                case "foo bar":
                    v3 = v3
                case "\n\t":
                    v1[32768] = 4294967295
                }
            }
        }
        _ = v11
    } else {
        v17 := GenS1{F0: ("foo bar" + (v3 + v3)), F1: 827, F2: 2147483647}
        if len(v2) > 0 {
            v2[0] = -504
        }
        var v18 int16 = -854
        _ = v17
        _ = v18
    }
    if (0 < 1) {
        if v4 {
            if len(v2) > 0 {
                v2[0] = -2147483648
            }
        } else {
            if true {
                v4 = (0 >= 42)
                fmt.Println(v3)
            }
            fmt.Println(v3, v4)
        }
        if v4 {
            var v19 uint = 42
            _ = v19
        }
    }
    switch 827 {
    case 32768:
        v20 := GenS1{F0: (v3 + v3), F1: 4294967295, F2: 0}
        _ = v20
    case -2147483648:
        v2 = append(v2, 128)
    case 9223372036854775807:
        for fi21 := 0; fi21 < 15; fi21++ {
            for fi22 := 0; fi22 < 20; fi22++ {
                v3 = ("foo bar" + "\n\t")
                _ = fi22
            }
            _ = fi21
        }
    }
    _ = len(v2)
    v4 = (!(18446744073709551615 >= genFn1((!false), v3)))
    v23 := GenS2{F0: 189, F1: 1}
    var v24 string = v3
    v4 = false
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint64; F1 bool }
func genFn0(p0 uint32, p1 bool) bool {
    if p1 {
        p1 = ((((("hello" + "hello") + ("hello" + (("" + (("\n\t" + ("hello" + "")) + "")) + ("a" + "")))) + "\n\t") + ("" + "")) >= "foo bar")
        fmt.Println(p0, p0)
    } else {
        p1 = (18 >= -66)
        for fi4 := 0; fi4 < 3; fi4++ {
            p1 = (!(0 < 2))
            if (!true) {
                for fi5 := 0; fi5 < 17; fi5++ {
                    v6 := []int{}
                    switch fi4 {
                    case 32767:
                        var v7 bool = (!p1)
                        _ = v7
                    }
                    _ = fi5
                    _ = v6
                }
            }
            _ = fi4
        }
        v8 := GenS1{F0: 103, F1: 18446744073709551615, F2: (p0 * 77)}
        _ = v8
    }
    var v9 bool = false
    v9 = p1
    p1 = v9
    _ = p0
    _ = p1
    _ = v9
    return (!false)
}

func main() {
    v1 := []string{("hello" + "foo bar"), "hello", "foo bar", (((("\n\t" + (("foo bar" + "foo bar") + ("\n\t" + ("hello" + "foo bar")))) + "hello") + "\n\t") + "")}
    var v2 uint16 = 492
    v3 := make(map[int]string)
    if false {
        var v10 int8 = 0
        var v11 float32 = float32((((float64((((-2.25 / 0.1) - -2.25) / (0.0 + float32((1e10 * 1e10))))) / (3.14159 / 1.5)) + -2.25) / 1.0))
        _ = v10
        _ = v11
    }
    v2 -= v2
    v2 -= v2
    v2 = 255
    v2 -= 919
    for fi12 := 0; fi12 < 11; fi12++ {
        if true {
            v2 = (v2 >> ((1) & 63))
            v13 := v3[fi12]
            switch v13 {
            case "a":
                for fi14 := 0; fi14 < 13; fi14++ {
                    v15 := GenS2{F0: 4294967295, F1: 1}
                    if len(v1) > 0 {
                        v1[0] = v13
                    }
                    v2 += v2
                    _ = fi14
                    _ = v15
                }
            default:
                v13 = ("hello" + "foo bar")
            }
            _ = v13
        }
        _ = fi12
    }
    var v16 int16 = 0
    if ((42 != 42) && (!(-1 != -15))) {
        if true {
            fmt.Println(v2, v16)
        }
        var v17 float32 = 3.14159
        _ = v17
    } else {
        fmt.Println(v16)
        var v18 float64 = (0.0 - -2.25)
        fmt.Println(v2, v16, v2)
        _ = v18
    }
    if (!(true || true)) {
        switch "hello" {
        case "":
            switch -2147483648 {
            case 520:
                v2 = 0
            case 127:
                v19 := v3[32768]
                _ = v19
            default:
                v20 := GenS2{F0: 2, F1: 2}
                _ = v20
            }
        }
    } else {
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
        if (3.14159 != 1e10) {
            var v21 uint16 = 1
            _ = v21
        }
    }
    switch ("a" + "\n\t") {
    case "":
        for fi22 := 0; fi22 < 1; fi22++ {
            v16 += int16(fi22)
            var v23 int8 = 36
            _ = len(v3)
            _ = fi22
            _ = v23
        }
    case "\n\t":
        v24 := GenS2{F0: 4294967295, F1: 287}
        _ = v24
    case "hello":
        _ = len(v3)
    default:
        if false {
            var v25 int = 42
            for fi26 := 0; fi26 < 14; fi26++ {
                var v27 bool = (36 == 100)
                _ = fi26
                _ = v27
            }
            if true {
                fmt.Println(v25, v2)
            } else {
                switch ("a" + "foo bar") {
                case "hello":
                    for fi28 := 0; fi28 < 4; fi28++ {
                        v2 = v2
                        _ = fi28
                    }
                case "\n\t":
                    fmt.Println(v25, v25, v25)
                }
                v29 := GenS1{F0: -42, F1: 190, F2: 2}
                _ = v29
            }
            _ = v25
        }
    }
    _ = len(v1)
    if (true && false) {
        if false {
            v3[803] = ""
            v2 -= (^(v2))
        }
        var v30 int = 4294967295
        _ = v30
    } else {
        var v31 uint8 = 127
        _ = v31
    }
    for fi32 := 0; fi32 < 9; fi32++ {
        v16 = v16
        fmt.Println(v16, fi32)
        v33 := GenS3{F0: 4294967295, F1: ((^(v2)) >= v2)}
        _ = fi32
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
}

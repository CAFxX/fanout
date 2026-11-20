package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 string }
func genFn0(p0 uint16, p1 int8) int16 {
    fmt.Println(p1, p0)
    p1 -= p1
    if (-9223372036854775808 > -9223372036854775808) {
        fmt.Println(p1, p0)
        switch 127 {
        case 4294967295:
            p1 -= -1
        case 128:
            var v5 uint64 = 713
            _ = v5
        default:
            for fi6 := 0; fi6 < 16; fi6++ {
                for fi7 := 0; fi7 < 8; fi7++ {
                    var v8 int16 = -32768
                    _ = fi7
                    _ = v8
                }
                var v9 float32 = -2.25
                v10 := []int{(-(fi6))}
                _ = fi6
                _ = v10
                _ = v9
            }
        }
        var v11 int16 = 1
        _ = v11
    }
    _ = p0
    _ = p1
    return 32767
}
func genFn1(p0 int32, p1 string) int {
    if false {
        for fi12 := 0; fi12 < 16; fi12++ {
            if false {
                fmt.Println(p0, p1)
            }
            var v13 bool = (!(!((!(127 >= 42)) || (!(false && (!true))))))
            _ = fi12
            _ = v13
        }
        if false {
            fmt.Println(p0)
            for fi14 := 0; fi14 < 2; fi14++ {
                switch (fi14 << ((551) & 63)) {
                case 359:
                    var v15 int = -693
                    _ = v15
                }
                _ = fi14
            }
            v16 := GenS1{F0: 6}
            _ = v16
        }
    } else {
        p1 = p1
    }
    if true {
        if (375 < 0) {
            v17 := GenS1{F0: 1}
            var v18 int = 2147483647
            var v19 int8 = 83
            _ = v17
            _ = v18
            _ = v19
        }
    }
    var v20 float32 = 0.1
    v21 := GenS1{F0: 4294967295}
    if (false && false) {
        if (p1 >= (p1 + (p1 + p1))) {
            fmt.Println(v20)
            if true {
                v22 := []int{0, 127, int(p0), -9223372036854775808}
                switch 417 {
                case -2147483648:
                    switch p1 {
                    case "hello":
                        var v23 uint64 = 2
                        _ = v23
                    case "foo bar":
                        if (true || false) {
                            p1 = p1
                        } else {
                            var v24 string = p1
                            _ = v24
                        }
                    case "":
                        var v25 float64 = -2.25
                        _ = v25
                    }
                }
                _ = v22
            }
        }
    }
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    return -2147483648
}

func main() {
    var v1 float32 = (1e10 / 1.5)
    var v2 int8 = 97
    var v3 int16 = 42
    v4 := []uint32{}
    var v26 float32 = v1
    fmt.Println(v2, v1, v3)
    v4 = append(v4, 42)
    switch -2147483648 {
    case 1:
        if (genFn1(76, ("\n\t" + ((("hello" + "foo bar") + ((("\n\t" + "\n\t") + "foo bar") + "foo bar")) + (("hello" + "") + "\n\t")))) <= 192) {
            fmt.Println(v2, v2, v2)
        } else {
            var v27 uint64 = 880
            _ = v27
        }
    }
    v3 -= v3
    v3 = v3
    var v28 float64 = 3.14159
    _ = len(v4)
    v2 += (-(v2))
    _ = len(v4)
    v2 -= -75
    v1 = ((v26 - (v26 * 3.14159)) / v26)
    if ("a" >= "foo bar") {
        if false {
            v29 := GenS2{F0: genFn1(-1, "foo bar"), F1: ("" + ((((("\n\t" + ("foo bar" + "foo bar")) + "hello") + "foo bar") + "foo bar") + "foo bar"))}
            _ = v29
        } else {
            v1 = (((0.1 / v26) / v1) + 0.0)
            var v30 int8 = v2
            v31 := GenS1{F0: 1}
            _ = v30
            _ = v31
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    if true {
        var v32 uint32 = 42
        _ = v32
    } else {
        var v33 float32 = float32(v28)
        v33 += (v1 / v33)
        _ = v33
    }
    var v34 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v34
    _ = v4
}

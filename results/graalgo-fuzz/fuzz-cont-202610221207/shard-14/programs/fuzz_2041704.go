package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint8; F1 int16 }
func genFn0() uint {
    if (32767 >= 0) {
        v7 := []int{32768, -906, -1}
        v8 := GenS1{F0: true, F1: 65535}
        _ = v7
        _ = v8
    } else {
        var v9 bool = true
        v9 = v9
        var v10 uint16 = 781
        _ = v10
        _ = v9
    }
    var v11 uint8 = 127
    _ = v11
    return 2
}
func genFn1(p0 uint8, p1 bool, p2 int16) uint16 {
    v12 := []int{4294967295, -2147483648, 128, 32768}
    if p1 {
        p2 -= 0
        if (-78 >= -58) {
            p0 = p0
        } else {
            var v13 string = "foo bar"
            _ = v13
        }
        p1 = p1
    }
    p2 += int16(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 255
}

func main() {
    var v1 float32 = float32(3.14159)
    v2 := []uint32{42, 467}
    var v3 uint8 = 42
    v4 := make(map[int]string)
    var v5 bool = (290 > 42)
    var v6 int = 1
    var v14 bool = v5
    var v15 float64 = 1.5
    switch v6 {
    case -175:
        var v16 int64 = 9223372036854775807
        _ = v16
    case 128:
        switch "a" {
        case "hello":
            v17 := GenS1{F0: false, F1: 32767}
            _ = v17
        case "foo bar":
            var v18 int32 = 1
            _ = v18
        default:
            var v19 uint16 = genFn1((^(v3)), (!false), 0)
            _ = v19
        }
    }
    if (4294967295 >= 2147483647) {
        switch "foo bar" {
        case "hello":
            v14 = (87 <= 100)
        case "foo bar":
            var v20 uint8 = 0
            _ = v20
        }
        fmt.Println(v14, v14)
        for fi21 := 0; fi21 < 18; fi21++ {
            fmt.Println(v15)
            v6 -= v6
            _ = fi21
        }
    } else {
        v5 = ((!v14) || false)
        _ = len(v4)
        var v22 string = ("" + "\n\t")
        _ = v22
    }
    v5 = false
    v1 = (((v1 - v1) + v1) - v1)
    v23 := GenS1{F0: (true || v14), F1: genFn1(v3, v14, -331)}
    var v24 uint32 = 2147483647
    if v5 {
        if v14 {
            for fi25 := 0; fi25 < 7; fi25++ {
                var v26 int64 = 9223372036854775807
                var v27 int64 = v26
                switch (-(fi25)) {
                case -483:
                    fmt.Println(v6, v26)
                case 32768:
                    for fi28 := 0; fi28 < 4; fi28++ {
                        v29 := GenS1{F0: ((v1 * ((-2.25 * v1) - (v1 + (v1 + float32((-2.25 / 3.14159)))))) >= v1), F1: uint16(v3)}
                        _ = fi28
                        _ = v29
                    }
                case -45:
                    v4[v6] = "\n\t"
                }
                _ = fi25
                _ = v26
                _ = v27
            }
            var v30 string = (("" + (("hello" + "") + ("" + "\n\t"))) + (("a" + ("foo bar" + ("\n\t" + "foo bar"))) + ""))
            var v31 uint64 = 4294967295
            _ = v30
            _ = v31
        } else {
            if (4294967295 == 4294967295) {
                var v32 string = "hello"
                _ = v32
            }
        }
    } else {
        switch "" {
        case "hello":
            v33 := GenS2{F0: 2, F1: 454}
            _ = v33
        case "foo bar":
            var v34 int8 = 1
            _ = v34
        case "\n\t":
            v35 := GenS1{F0: v5, F1: 42}
            _ = v35
        default:
            fmt.Println(v5, v24)
        }
        for fi36 := 0; fi36 < 17; fi36++ {
            v37 := GenS1{F0: v5, F1: genFn1(v3, (42 < 2), -1)}
            fmt.Println(fi36)
            _ = fi36
            _ = v37
        }
    }
    var v38 uint = 0
    v1 += (3.14159 * (v1 / 3.14159))
    fmt.Println(v3, v6, v5)
    var v39 uint = genFn0()
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}

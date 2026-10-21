package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int; F1 int16 }
func genFn0(p0 bool, p1 bool) uint64 {
    for fi4 := 0; fi4 < 4; fi4++ {
        switch "foo bar" {
        case "":
            v5 := GenS1{F0: (3.14159 == 3.14159), F1: 677}
            _ = v5
        case "a":
            var v6 uint64 = 568
            _ = v6
        case "\n\t":
            var v7 int8 = 1
            _ = v7
        default:
            p1 = (-1 >= 47)
        }
        _ = fi4
    }
    v8 := []int{1, 9223372036854775807, 985, -312}
    if len(v8) > 0 {
        v8[0] = 4294967295
    }
    _ = p0
    _ = p1
    _ = v8
    return 0
}
func genFn1(p0 string) string {
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := GenS1{F0: (42 == 100), F1: -915}
        _ = fi9
        _ = v10
    }
    p0 = ("hello" + "hello")
    var v11 uint16 = 42
    _ = p0
    _ = v11
    return (p0 + "\n\t")
}

func main() {
    var v1 bool = false
    var v2 int16 = 32767
    var v3 bool = (-128 >= -39)
    for fi12 := 0; fi12 < 16; fi12++ {
        if false {
            var v13 uint = 913
            _ = v13
        } else {
            v14 := GenS2{F0: (fi12 % ((fi12) | 1)), F1: v2}
            _ = v14
        }
        _ = fi12
    }
    v15 := []int{}
    v15 = append(v15, int(v2))
    v1 = (v1 || (((float32(-2.25) - 1.5) * (0.1 / -2.25)) == 1e10))
    var v16 float64 = (float64((0.1 - 1e10)) / 1.0)
    for fi17 := 0; fi17 < 18; fi17++ {
        for fi18 := 0; fi18 < 5; fi18++ {
            v3 = (42 <= 0)
            _ = fi18
        }
        _ = fi17
    }
    if true {
        var v19 uint16 = 255
        if v3 {
            if v3 {
                fmt.Println(v1, v16)
                var v20 string = "a"
                _ = v20
            } else {
                for fi21 := 0; fi21 < 7; fi21++ {
                    var v22 float64 = ((v16 * 0.0) + v16)
                    _ = fi21
                    _ = v22
                }
                v23 := GenS2{F0: 32768, F1: v2}
                v19 += v19
                _ = v23
            }
            fmt.Println(v16)
        }
        _ = v19
    }
    var v24 uint = 2
    if (!false) {
        fmt.Println(v2, v1)
        v2 += v2
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        for fi26 := 0; fi26 < 7; fi26++ {
            switch ("hello" + "foo bar") {
            case "":
                switch "hello" {
                case "hello":
                    switch (-(fi26)) {
                    case 42:
                        _ = len(v15)
                    }
                }
            case "\n\t":
                switch ("hello" + ("" + "hello")) {
                case "\n\t":
                    var v27 bool = false
                    _ = v27
                case "":
                    var v28 int64 = -324
                    _ = v28
                }
            case "foo bar":
                v3 = (!v1)
            }
            if v1 {
                v2 = (v2 | v2)
                var v29 int8 = 1
                var v30 uint = (v24 | (v24 * (v24 / ((1) | 1))))
                _ = v29
                _ = v30
            } else {
                v31 := GenS1{F0: true, F1: 1}
                _ = v31
            }
            v24 = (v24 << ((v24) & 63))
            _ = fi26
        }
        var v32 float32 = 0.0
        _ = fi25
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v3
}

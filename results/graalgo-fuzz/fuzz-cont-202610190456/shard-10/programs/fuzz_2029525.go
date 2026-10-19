package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 int32 }
func genFn0(p0 uint32, p1 uint16, p2 uint) string {
    var v5 int16 = 0
    if true {
        for fi6 := 0; fi6 < 11; fi6++ {
            var v7 float32 = 0.1
            _ = fi6
            _ = v7
        }
        switch ("foo bar" + ("\n\t" + ("a" + "foo bar"))) {
        case "hello":
            switch "foo bar" {
            case "":
                v8 := []int{-9223372036854775808, 4294967295, 128}
                _ = v8
            }
        }
    } else {
        v5 += (-(v5))
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        p0 += (p0 % ((p0) | 1))
        for fi10 := 0; fi10 < 10; fi10++ {
            v11 := GenS1{F0: true, F1: 65535}
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 10; fi12++ {
            p0 += p0
            v13 := GenS1{F0: (42 < 316), F1: (p2 % ((p2) | 1))}
            if (!true) {
                v14 := GenS1{F0: (false || true), F1: (p2 * 192)}
                v15 := GenS2{F0: 128}
                var v16 float32 = 1e10
                _ = v14
                _ = v15
                _ = v16
            }
            _ = fi12
            _ = v13
        }
        _ = fi9
    }
    v17 := GenS2{F0: 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v5
    return ("foo bar" + "foo bar")
}

func main() {
    var v1 uint8 = 15
    v2 := make(map[int]uint32)
    var v3 float32 = -2.25
    v4 := []string{"", ("\n\t" + ""), "hello", "foo bar"}
    var v18 int8 = 42
    v18 = v18
    if (354 < 129) {
        var v19 uint32 = 65535
        var v20 uint16 = 208
        v21 := GenS1{F0: false, F1: 1}
        _ = v19
        _ = v20
        _ = v21
    }
    switch "a" {
    case "foo bar":
        v22 := GenS3{F0: ("foo bar" + "hello"), F1: -1}
        _ = v22
    }
    delete(v2, 9223372036854775807)
    switch "" {
    case "foo bar":
        var v23 int64 = -789
        _ = v23
    case "hello":
        if true {
            v24 := GenS1{F0: false, F1: 840}
            _ = v24
        }
    case "":
        var v25 float32 = -2.25
        _ = v25
    default:
        v18 = (v18 >> ((0) & 63))
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        v2[(^(fi26))] = 2
        _ = fi26
    }
    v1 = (^(v1))
    v27 := GenS2{F0: v1}
    for fi28 := 0; fi28 < 1; fi28++ {
        v29 := GenS1{F0: ((3.14159 / (3.14159 + (3.14159 - 1e10))) > (((-2.25 / 1.0) - 1e10) + (0.1 / 1.0))), F1: 42}
        switch 800 {
        case 302:
            var v30 int64 = 0
            _ = v30
        case 127:
            if (!false) {
                switch "foo bar" {
                case "":
                    v31 := v2[fi28]
                    _ = v31
                }
                if true {
                    v3 += v3
                    fmt.Println(v3)
                    switch (fi28 & (fi28 & -2147483648)) {
                    case 2147483647:
                        var v32 uint = 255
                        _ = v32
                    default:
                        if len(v4) > 0 {
                            v4[0] = "a"
                        }
                    }
                } else {
                    if true {
                        var v33 int16 = 0
                        var v34 int32 = 24
                        _ = v33
                        _ = v34
                    }
                }
            }
        case 573:
            v18 = (v18 * 127)
        default:
            fmt.Println(fi28)
        }
        fmt.Println(v18, v18, v3)
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}

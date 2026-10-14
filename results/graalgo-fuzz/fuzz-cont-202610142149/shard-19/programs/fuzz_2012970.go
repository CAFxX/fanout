package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int64; F1 uint16 }
func genFn0(p0 int8, p1 int8, p2 int) uint8 {
    v5 := []int{}
    v5 = append(v5, p2)
    p2 = int(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 10
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[string]bool)
    var v3 uint = 1
    v4 := []int{32768, 32767, 335, 307}
    var v6 uint = v3
    if false {
        fmt.Println(v1, v1)
        var v7 string = "foo bar"
        delete(v2, ((("" + "") + v7) + v7))
        _ = v7
    }
    v8 := GenS1{F0: 127, F1: 42}
    v1 = (v1 ^ (v1 << (((v6 - v3)) & 63)))
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 uint32 = 1
        fmt.Println(fi9, fi9)
        fmt.Println(v6, v10, fi9)
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        var v12 int = (fi11 & (^(fi11)))
        if true {
            if true {
                var v13 bool = false
                switch fi11 {
                case 128:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                _ = v13
            } else {
                switch "" {
                case "":
                    switch "hello" {
                    case "":
                        if (!true) {
                            v14 := GenS1{F0: 0, F1: 758}
                            _ = v14
                        }
                    }
                case "hello":
                    var v15 int64 = int64(v1)
                    _ = v15
                }
                v1 += v1
            }
        }
        fmt.Println(fi11, v3)
        _ = fi11
        _ = v12
    }
    v16 := GenS1{F0: -1, F1: -1}
    var v17 uint8 = (v1 * genFn0(13, -64, -1))
    if (-949 == 860) {
        var v18 int = -961
        for fi19 := 0; fi19 < 2; fi19++ {
            fmt.Println(v18)
            v20 := GenS1{F0: 127, F1: 548}
            _ = fi19
            _ = v20
        }
        _ = len(v4)
        _ = v18
    }
    switch 128 {
    case 38:
        _ = len(v2)
    case -351:
        v21 := GenS1{F0: 42, F1: -1}
        _ = v21
    case -9:
        var v22 bool = false
        _ = v22
    }
    fmt.Println(v6, v6, v6)
    switch "a" {
    case "foo bar":
        var v23 int16 = 0
        _ = v23
    case "hello":
        v17 += 128
    case "\n\t":
        v24 := GenS1{F0: -128, F1: -1}
        _ = v24
    }
    var v25 int16 = 944
    if (0 == 65535) {
        fmt.Println(v3, v3, v25)
    } else {
        v26 := GenS1{F0: int8(v25), F1: (v25 << (((v6 >> ((v3) & 63))) & 63))}
        _ = v26
    }
    v17 += 0
    delete(v2, "hello")
    if (0 != 265) {
        var v27 int = -9223372036854775808
        _ = v27
    } else {
        var v28 int8 = 42
        switch "\n\t" {
        case "foo bar":
            v17 -= 0
        case "\n\t":
            switch "" {
            case "a":
                v3 += (^(v3))
            case "foo bar":
                if (!true) {
                    var v29 bool = (v28 <= v28)
                    v30 := GenS2{F0: 42, F1: 292}
                    _ = v29
                    _ = v30
                }
            case "":
                fmt.Println(v28)
            }
        case "a":
            if (false && (120 > 233)) {
                _ = len(v4)
            } else {
                if len(v4) > 0 {
                    v4[0] = int(v25)
                }
                var v31 string = "hello"
                _ = v31
            }
        }
        v6 = v6
        _ = v28
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v6
    _ = v8
}

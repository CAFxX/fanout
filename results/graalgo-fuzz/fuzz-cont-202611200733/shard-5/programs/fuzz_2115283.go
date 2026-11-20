package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 bool; F2 bool }
func genFn0(p0 int16, p1 string) uint {
    var v7 float32 = ((float32(1e10) / float32(1.5)) / 1e10)
    if true {
        v7 -= ((v7 + v7) - v7)
    } else {
        if false {
            for fi8 := 0; fi8 < 6; fi8++ {
                var v9 int = fi8
                if ((457 <= 485) && (1 != -128)) {
                    p0 += (p0 * p0)
                    if (18446744073709551615 != 299) {
                        for fi10 := 0; fi10 < 7; fi10++ {
                            v11 := []int{(fi8 | (^(fi8)))}
                            var v12 int16 = p0
                            _ = fi10
                            _ = v11
                            _ = v12
                        }
                    } else {
                        v7 = float32(1e10)
                        v13 := GenS1{F0: false}
                        _ = v13
                    }
                }
                _ = fi8
                _ = v9
            }
            v7 = ((-2.25 - v7) - v7)
            switch p1 {
            case "hello":
                v14 := GenS2{F0: 111, F1: false, F2: false}
                _ = v14
            case "\n\t":
                p0 = 0
            case "":
                p1 = (("" + p1) + p1)
            default:
                switch p1 {
                case "hello":
                    v15 := GenS1{F0: true}
                    _ = v15
                }
            }
        } else {
            var v16 string = p1
            var v17 uint = 647
            p0 -= p0
            _ = v16
            _ = v17
        }
        switch 42 {
        case 1:
            p1 = ("a" + p1)
        case 307:
            var v18 int64 = 361
            _ = v18
        case -953:
            p1 = p1
        }
    }
    fmt.Println(p1, v7)
    p1 = p1
    var v19 int32 = 1
    _ = p0
    _ = p1
    _ = v19
    _ = v7
    return 255
}

func main() {
    var v1 uint = 479
    v2 := []int{}
    v3 := []int{128, 32768}
    v4 := make(map[string]string)
    v5 := []int{997, 128}
    var v6 int64 = -1
    v6 += 42
    delete(v4, (((("" + "\n\t") + "a") + "hello") + ""))
    v20 := GenS1{F0: (32768 < 42)}
    v6 = v6
    for fi21 := 0; fi21 < 5; fi21++ {
        if true {
            fmt.Println(v6)
            delete(v4, (("hello" + "") + "foo bar"))
            v22 := GenS2{F0: 153, F1: false, F2: true}
            _ = v22
        }
        v6 = v6
        _ = fi21
    }
    if false {
        var v23 int32 = 942
        _ = len(v2)
        _ = v23
    }
    var v24 string = "hello"
    var v25 uint16 = 2
    fmt.Println(v1, v1)
    v26 := GenS2{F0: 2147483647, F1: ((v24 != v24) && true), F2: false}
    var v27 uint32 = 4294967295
    var v28 bool = ("hello" != v24)
    fmt.Println(v28, v28, v6)
    v6 -= v6
    var v29 uint8 = 48
    fmt.Println(v29, v25)
    if v28 {
        v25 -= (v25 << ((v1) & 63))
        v30 := v4[v24]
        _ = v30
    }
    switch 42 {
    case -559:
        fmt.Println(v24)
    case 0:
        for fi31 := 0; fi31 < 5; fi31++ {
            for fi32 := 0; fi32 < 9; fi32++ {
                var v33 bool = ((v29 & v29) == v29)
                fmt.Println(fi32, fi32)
                _ = fi32
                _ = v33
            }
            _ = len(v5)
            fmt.Println(fi31, v28)
            _ = fi31
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

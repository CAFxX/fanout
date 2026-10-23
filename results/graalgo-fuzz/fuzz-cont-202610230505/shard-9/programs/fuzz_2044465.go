package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint8 }
type GenS2 struct { F0 int32; F1 int16 }
func genFn0(p0 uint64, p1 uint8, p2 string) uint {
    var v6 uint8 = (p1 >> ((4294967295) & 63))
    switch 4294967295 {
    case 9223372036854775807:
        v6 += (v6 * v6)
    case -232:
        switch p2 {
        case "\n\t":
            v7 := []int{-879, 128}
            _ = v7
        default:
            switch "hello" {
            case "\n\t":
                var v8 int64 = 9223372036854775807
                _ = v8
            case "a":
                if (-783 <= -362) {
                    fmt.Println(p2, p0, p0)
                } else {
                    v9 := GenS1{F0: 6, F1: 733, F2: 150}
                    fmt.Println(p2, p0)
                    if false {
                        if ((460 > -1) && true) {
                            p1 -= v6
                        }
                        fmt.Println(p1)
                        v10 := GenS1{F0: v6, F1: 0, F2: (v6 >> ((978) & 63))}
                        _ = v10
                    } else {
                        switch (p2 + ("hello" + "foo bar")) {
                        case "a":
                            var v11 uint64 = uint64(v6)
                            _ = v11
                        }
                        p0 = 0
                    }
                    _ = v9
                }
            }
        }
    case 42:
        if true {
            switch p2 {
            case "foo bar":
                v12 := GenS1{F0: p1, F1: -514, F2: (v6 / ((255) | 1))}
                _ = v12
            case "a":
                p2 = ""
            }
        } else {
            v6 -= (p1 ^ (p1 + (v6 / ((96) | 1))))
            var v13 uint8 = 195
            _ = v13
        }
    }
    p1 = (v6 << ((304) & 63))
    var v14 int = -355
    for fi15 := 0; fi15 < 4; fi15++ {
        for fi16 := 0; fi16 < 9; fi16++ {
            v17 := GenS1{F0: 74, F1: -541, F2: 255}
            fmt.Println(fi15, p0)
            var v18 int = v14
            _ = fi16
            _ = v17
            _ = v18
        }
        var v19 string = p2
        var v20 float64 = -2.25
        _ = fi15
        _ = v19
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v6
    return 255
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{65535, 4294967295, 938}
    var v3 uint8 = 2
    v4 := make(map[int]bool)
    var v5 float32 = 1.5
    fmt.Println(v5, v3)
    for fi21 := 0; fi21 < 10; fi21++ {
        v5 -= v5
        var v22 uint64 = 4294967295
        delete(v4, fi21)
        _ = fi21
        _ = v22
    }
    var v23 int32 = -2147483648
    for fi24 := 0; fi24 < 3; fi24++ {
        var v25 int = (fi24 & fi24)
        _ = fi24
        _ = v25
    }
    switch ("foo bar" + "a") {
    case "a":
        v23 = v23
    case "":
        var v26 int16 = 32767
        _ = v26
    case "foo bar":
        _ = len(v2)
    }
    fmt.Println(v3, v5)
    v27 := GenS2{F0: v23, F1: -525}
    var v28 uint16 = 2
    for fi29 := 0; fi29 < 15; fi29++ {
        fmt.Println(fi29)
        var v30 int16 = 862
        _ = fi29
        _ = v30
    }
    switch "a" {
    case "foo bar":
        fmt.Println(v5, v3)
    default:
        v5 -= (v5 - -2.25)
    }
    v2 = append(v2, 0)
    var v31 uint = 2
    if (true && ((-2.25 == (0.0 * 1e10)) && ((!true) && false))) {
        var v32 int16 = -1
        v5 += ((v5 * 1.5) - v5)
        v3 = v3
        _ = v32
    }
    var v33 int32 = (v23 << ((genFn0(uint64(v31), (v3 ^ v3), "hello")) & 63))
    fmt.Println(v5)
    _ = len(v4)
    var v34 int = 629
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int; F1 int8 }
func genFn0(p0 bool) int {
    if (1 <= -2147483648) {
        if (true && (!true)) {
            p0 = p0
        }
        p0 = false
        if (2 >= 509) {
            fmt.Println(p0)
        }
    }
    if true {
        var v4 uint16 = 255
        p0 = p0
        _ = v4
    }
    var v5 uint8 = 128
    v5 += 0
    _ = p0
    _ = v5
    return 32767
}
func genFn1() bool {
    var v6 float32 = (0.0 - ((1.5 / float32(0.1)) + 3.14159))
    v6 -= v6
    v7 := []int{32767, -1}
    if (428 < 2147483648) {
        for fi8 := 0; fi8 < 4; fi8++ {
            for fi9 := 0; fi9 < 5; fi9++ {
                switch "\n\t" {
                case "hello":
                    var v10 uint64 = 425
                    _ = v10
                default:
                    switch "hello" {
                    case "hello":
                        v6 = v6
                    }
                }
                _ = fi9
            }
            for fi11 := 0; fi11 < 5; fi11++ {
                if true {
                    v6 = 3.14159
                    var v12 string = "\n\t"
                    _ = v12
                } else {
                    v13 := GenS1{F0: -2147483648, F1: 168}
                    if len(v7) > 0 {
                        v7[0] = (fi11 | fi11)
                    }
                    _ = v13
                }
                var v14 int = fi8
                switch ("a" + ("hello" + ("\n\t" + "foo bar"))) {
                case "\n\t":
                    var v15 float32 = 1e10
                    _ = v15
                case "a":
                    v14 = (^(v14))
                case "foo bar":
                    v7 = append(v7, (fi8 >> ((2) & 63)))
                default:
                    var v16 uint64 = 4294967295
                    _ = v16
                }
                if (170 > 42) {
                    fmt.Println(v14)
                    var v17 int8 = -1
                    var v18 uint8 = 85
                    _ = v17
                    _ = v18
                } else {
                    switch (fi11 ^ -726) {
                    case 128:
                        var v19 float64 = 0.0
                        _ = v19
                    case 702:
                        var v20 uint64 = 0
                        _ = v20
                    case -256:
                        v7 = append(v7, fi8)
                    default:
                        v14 += (v14 >> ((148) & 63))
                    }
                    var v21 int = fi11
                    _ = v21
                }
                _ = fi11
                _ = v14
            }
            _ = fi8
        }
        v6 += (v6 + v6)
        _ = len(v7)
    } else {
        v22 := GenS1{F0: -1, F1: 2}
        var v23 uint16 = 1
        _ = v22
        _ = v23
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = v6
    _ = v7
    return ((true || (-1 != 1)) || false)
}

func main() {
    var v1 int32 = -202
    v2 := make(map[string]int)
    v3 := []string{("\n\t" + (("foo bar" + "a") + "hello")), "foo bar", "\n\t"}
    _ = len(v2)
    v1 = (v1 - v1)
    v24 := GenS1{F0: v1, F1: 255}
    v25 := GenS1{F0: -115, F1: 255}
    var v26 uint64 = 18446744073709551615
    if (-1 != -117) {
        v1 = v1
        var v27 float64 = 1e10
        for fi28 := 0; fi28 < 20; fi28++ {
            fmt.Println(v27)
            v29 := GenS2{F0: (^(fi28)), F1: 93}
            switch "\n\t" {
            case "\n\t":
                if false {
                    _ = len(v3)
                    if (!true) {
                        fmt.Println(fi28)
                    } else {
                        fmt.Println(v27, v27, v27)
                    }
                }
            }
            _ = fi28
            _ = v29
        }
        _ = v27
    } else {
        v26 -= v26
        v26 = v26
    }
    var v30 int = 128
    for fi31 := 0; fi31 < 1; fi31++ {
        var v32 int64 = 1
        _ = fi31
        _ = v32
    }
    v30 = v30
    v33 := GenS2{F0: (v30 >> ((42) & 63)), F1: 55}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v33
}

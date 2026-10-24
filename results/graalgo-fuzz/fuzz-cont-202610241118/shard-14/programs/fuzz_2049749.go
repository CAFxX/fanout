package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint16) int64 {
    p0 -= 626
    switch (("foo bar" + ("foo bar" + "hello")) + "\n\t") {
    case "foo bar":
        for fi6 := 0; fi6 < 19; fi6++ {
            p0 -= 0
            for fi7 := 0; fi7 < 7; fi7++ {
                var v8 uint8 = 210
                _ = fi7
                _ = v8
            }
            _ = fi6
        }
    default:
        p0 = (p0 & 0)
    }
    var v9 uint32 = 4294967295
    _ = p0
    _ = v9
    return 745
}
func genFn1(p0 bool, p1 bool, p2 int) string {
    switch "foo bar" {
    case "\n\t":
        v10 := GenS1{F0: 609, F1: p2}
        _ = v10
    default:
        v11 := GenS2{F0: 895}
        _ = v11
    }
    var v12 int8 = 1
    v13 := []int{p2}
    for fi14 := 0; fi14 < 14; fi14++ {
        if p1 {
            v15 := GenS1{F0: 65535, F1: 1}
            _ = v15
        } else {
            p2 += (p2 & 9223372036854775807)
            p0 = p1
        }
        p0 = p0
        var v16 uint = 1
        _ = fi14
        _ = v16
    }
    fmt.Println(p1, p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return "a"
}

func main() {
    var v1 uint32 = 132
    v2 := []string{("foo bar" + "hello"), ("\n\t" + "\n\t"), ("foo bar" + (("foo bar" + (("foo bar" + "\n\t") + (("a" + (("a" + "hello") + "hello")) + "hello"))) + ("\n\t" + "\n\t")))}
    v3 := make(map[string]uint32)
    v4 := make(map[string]bool)
    var v5 float64 = 1.5
    if false {
        _ = len(v2)
        var v17 int = 32768
        _ = v17
    } else {
        fmt.Println(v5)
        var v18 uint64 = 4294967295
        _ = v18
    }
    var v19 string = ("hello" + ("a" + "\n\t"))
    _ = len(v2)
    var v20 uint32 = v1
    v5 += 3.14159
    if (v19 >= v19) {
        switch v19 {
        case "foo bar":
            var v21 string = ("hello" + "\n\t")
            _ = v21
        case "":
            if true {
                v20 += 1
            }
        case "hello":
            fmt.Println(v5, v20)
        }
        v1 -= v20
    } else {
        v5 = v5
    }
    _ = len(v4)
    delete(v4, ("\n\t" + v19))
    v22 := GenS1{F0: 32767, F1: int(v1)}
    v23 := GenS2{F0: 42}
    var v24 int16 = 32767
    v24 -= v24
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}

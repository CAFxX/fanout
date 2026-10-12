package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int64 }
func genFn0(p0 uint16, p1 int64) uint {
    p1 += (p1 | (p1 + (^(p1))))
    if (!(128 == 128)) {
        fmt.Println(p1, p1)
        v6 := GenS1{F0: 0}
        _ = v6
    }
    _ = p0
    _ = p1
    return 18446744073709551615
}

func main() {
    var v1 int8 = 92
    v2 := []int{}
    v3 := make(map[int]bool)
    var v4 int16 = -32768
    v5 := []string{(("foo bar" + "\n\t") + (("\n\t" + "a") + "a")), "a"}
    switch (("a" + "foo bar") + "foo bar") {
    case "a":
        var v7 float64 = ((0.0 / 1.0) / 0.1)
        _ = v7
    case "hello":
        for fi8 := 0; fi8 < 7; fi8++ {
            v4 = v4
            for fi9 := 0; fi9 < 19; fi9++ {
                fmt.Println(v4)
                v1 = v1
                v10 := GenS1{F0: 1}
                _ = fi9
                _ = v10
            }
            var v11 uint64 = 782
            _ = fi8
            _ = v11
        }
    default:
        v12 := GenS1{F0: 98}
        _ = v12
    }
    v13 := GenS2{F0: 51, F1: -1}
    var v14 string = "a"
    v5 = append(v5, "hello")
    v2 = append(v2, 9223372036854775807)
    if true {
        v1 = -1
        var v15 uint32 = 827
        v14 = v14
        _ = v15
    } else {
        for fi16 := 0; fi16 < 4; fi16++ {
            v1 = (v1 >> ((1) & 63))
            var v17 float32 = 0.1
            var v18 uint16 = 42
            _ = fi16
            _ = v17
            _ = v18
        }
        v4 -= -32768
        v19 := GenS1{F0: 0}
        _ = v19
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        fmt.Println(v4, v14)
        for fi21 := 0; fi21 < 13; fi21++ {
            switch (fi20 << ((4294967295) & 63)) {
            case -2147483648:
                var v22 uint8 = 255
                _ = v22
            default:
                if false {
                    v14 = "\n\t"
                    var v23 uint8 = 33
                    _ = v23
                } else {
                    fmt.Println(v1)
                }
            }
            var v24 int32 = -1
            v4 = v4
            _ = fi21
            _ = v24
        }
        v25 := v3[(fi20 | -842)]
        _ = fi20
        _ = v25
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        if true {
            _ = len(v5)
            delete(v3, (-(fi26)))
        }
        _ = fi26
    }
    for fi27 := 0; fi27 < 4; fi27++ {
        v14 = v14
        _ = fi27
    }
    v14 = ((v14 + v14) + "foo bar")
    var v28 int16 = 1000
    v29 := GenS1{F0: 2}
    if ((false || true) || false) {
        v4 += 0
        delete(v3, 1)
        v30 := GenS2{F0: 1, F1: 259}
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}

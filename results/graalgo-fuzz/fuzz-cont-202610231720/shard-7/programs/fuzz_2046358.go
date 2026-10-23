package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int }
func genFn0(p0 uint8, p1 uint16, p2 uint64) bool {
    p0 = 36
    v6 := []int{127, 128}
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 uint16 = (p1 - (p1 | p1))
        _ = fi7
        _ = v8
    }
    switch ((("a" + ("\n\t" + "")) + "\n\t") + "hello") {
    case "a":
        switch (("hello" + ("hello" + "a")) + "foo bar") {
        case "a":
            var v9 float64 = float64(float32((((float64(1.5) - 0.0) * 0.0) + float64(1e10))))
            _ = v9
        case "foo bar":
            var v10 bool = (255 == (p0 >> ((0) & 63)))
            _ = v10
        }
    case "":
        if len(v6) > 0 {
            v6[0] = 4294967295
        }
    default:
        if (!(127 <= -9223372036854775808)) {
            for fi11 := 0; fi11 < 6; fi11++ {
                p1 += (^(p1))
                switch "foo bar" {
                case "hello":
                    v12 := GenS1{F0: 4294967295, F1: 42}
                    _ = v12
                case "\n\t":
                    var v13 int64 = 9223372036854775807
                    _ = v13
                case "a":
                    var v14 uint8 = p0
                    _ = v14
                default:
                    p2 = p2
                }
                fmt.Println(fi11, fi11)
                _ = fi11
            }
            if len(v6) > 0 {
                v6[0] = int(p2)
            }
        }
    }
    var v15 bool = ((473 > 42) && true)
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v6
    return ((33 > -128) && v15)
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 0
    v3 := []uint32{}
    var v4 int = -2147483648
    var v5 bool = true
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        v2 += v2
        _ = len(v3)
        _ = fi16
    }
    v17 := GenS1{F0: v2, F1: v2}
    if ((844 == 385) || (2 >= 899)) {
        v2 = 61
    } else {
        var v18 int = 4294967295
        var v19 uint = (v2 - v2)
        _ = v18
        _ = v19
    }
    if (127 < 237) {
        delete(v1, v4)
        for fi20 := 0; fi20 < 16; fi20++ {
            for fi21 := 0; fi21 < 5; fi21++ {
                v22 := GenS2{F0: 127}
                switch ("\n\t" + ((("" + "foo bar") + "\n\t") + "a")) {
                case "":
                    v2 += (v2 & 18446744073709551615)
                default:
                    v5 = (601 < 1)
                }
                v4 += 127
                _ = fi21
                _ = v22
            }
            v23 := GenS1{F0: v2, F1: 65535}
            _ = fi20
            _ = v23
        }
        v5 = (!((^(v2)) > (^(v2))))
    } else {
        var v24 int16 = 1000
        v24 += v24
        _ = v24
    }
    v5 = v5
    var v25 uint64 = 18446744073709551615
    v26 := GenS1{F0: 69, F1: 255}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}

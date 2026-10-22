package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 uint32; F1 int }
func genFn0(p0 bool, p1 int8) int8 {
    p1 = p1
    for fi7 := 0; fi7 < 14; fi7++ {
        p0 = false
        switch ("foo bar" + "a") {
        case "hello":
            if p0 {
                if (3.14159 < (-2.25 + 1e10)) {
                    fmt.Println(fi7, fi7)
                    var v8 float64 = (1.5 / 0.1)
                    v9 := GenS1{F0: 127}
                    _ = v8
                    _ = v9
                } else {
                    var v10 uint16 = 65535
                    _ = v10
                }
            }
        case "\n\t":
            var v11 float64 = -2.25
            _ = v11
        case "a":
            fmt.Println(p1, fi7, p1)
        default:
            p1 = p1
        }
        var v12 uint64 = 999
        _ = fi7
        _ = v12
    }
    _ = p0
    _ = p1
    return (p1 + p1)
}

func main() {
    v1 := []int{414, -178}
    v2 := make(map[uint32]uint32)
    var v3 uint = 18446744073709551615
    var v4 uint64 = 558
    var v5 uint64 = v4
    v6 := []string{"foo bar", ("foo bar" + "hello")}
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 string = "foo bar"
        _ = fi13
        _ = v14
    }
    var v15 uint16 = 65535
    if true {
        v5 += 516
        for fi16 := 0; fi16 < 9; fi16++ {
            for fi17 := 0; fi17 < 9; fi17++ {
                v6 = append(v6, ("foo bar" + "a"))
                _ = len(v6)
                _ = fi17
            }
            for fi18 := 0; fi18 < 7; fi18++ {
                fmt.Println(v5)
                switch (fi18 << ((v3) & 63)) {
                case -335:
                    fmt.Println(fi18)
                case -1:
                    v19 := GenS2{F0: uint32(v5)}
                    _ = v19
                }
                if (true || false) {
                    v4 = v4
                }
                _ = fi18
            }
            fmt.Println(v15, v4, v15)
            _ = fi16
        }
        for fi20 := 0; fi20 < 8; fi20++ {
            fmt.Println(v15, v4, v5)
            if true {
                v21 := GenS3{F0: (v15 >> ((354) & 63)), F1: uint32(v3)}
                var v22 int = fi20
                v15 = (v15 - 229)
                _ = v21
                _ = v22
            }
            for fi23 := 0; fi23 < 18; fi23++ {
                var v24 int16 = -32768
                _ = fi23
                _ = v24
            }
            _ = fi20
        }
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        v15 = 32767
        if len(v6) > 0 {
            v6[0] = ("a" + "hello")
        }
        v6 = append(v6, "a")
        _ = fi25
    }
    v26 := GenS4{F0: 42, F1: 0}
    var v27 float64 = -2.25
    v4 += (v5 & (v4 | (v4 / ((v4) | 1))))
    for fi28 := 0; fi28 < 9; fi28++ {
        v4 -= v4
        _ = fi28
    }
    var v29 int8 = genFn0((1 == 42), -1)
    var v30 uint = 255
    v5 += (v5 ^ v4)
    var v31 int8 = v29
    v6 = append(v6, "")
    fmt.Println(v27, v5)
    v27 -= float64((float32(((0.1 - float64(-2.25)) + 1.5)) + 3.14159))
    v32 := v2[65535]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int) uint32 {
    var v5 string = "foo bar"
    v6 := []int{p0, p0}
    switch (p0 * -901) {
    case -796:
        for fi7 := 0; fi7 < 9; fi7++ {
            v6 = append(v6, 42)
            var v8 uint = 440
            _ = fi7
            _ = v8
        }
    case 1:
        switch "foo bar" {
        case "foo bar":
            for fi9 := 0; fi9 < 10; fi9++ {
                for fi10 := 0; fi10 < 9; fi10++ {
                    fmt.Println(p0, p0, fi9)
                    var v11 uint = 18446744073709551615
                    _ = fi10
                    _ = v11
                }
                v12 := GenS1{F0: 869}
                _ = fi9
                _ = v12
            }
        case "a":
            fmt.Println(p0, p0)
        case "\n\t":
            v13 := GenS1{F0: 255}
            _ = v13
        }
    case 552:
        if false {
            if true {
                for fi14 := 0; fi14 < 15; fi14++ {
                    if (-774 < 701) {
                        var v15 string = "foo bar"
                        var v16 uint32 = 729
                        _ = v15
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 7; fi17++ {
                        var v18 uint16 = 65535
                        v5 = v5
                        _ = fi17
                        _ = v18
                    }
                    _ = fi14
                }
                fmt.Println(v5)
            }
            var v19 int = -323
            v19 = (-(p0))
            _ = v19
        }
    default:
        fmt.Println(p0)
    }
    var v20 float64 = (1.5 * 0.1)
    _ = p0
    _ = v20
    _ = v5
    _ = v6
    return 192
}
func genFn1(p0 uint16, p1 int, p2 uint) int8 {
    p2 = p2
    var v21 bool = (60 != 4294967295)
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return 127
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{}
    v3 := []uint32{2147483647}
    var v4 float64 = (1.5 / 0.1)
    var v22 int = -460
    if true {
        var v23 float32 = (3.14159 / 0.1)
        v22 = -597
        v4 = v4
        _ = v23
    } else {
        if (3.14159 != (v4 + (0.1 - 0.1))) {
            var v24 float32 = float32(v4)
            if true {
                switch "hello" {
                case "":
                    var v25 int8 = genFn1(0, v22, 629)
                    _ = v25
                case "\n\t":
                    v1[v22] = v22
                case "hello":
                    v24 -= (v24 - 0.1)
                default:
                    fmt.Println(v24, v22, v24)
                }
            }
            var v26 uint64 = 4294967295
            _ = v24
            _ = v26
        }
    }
    v4 -= v4
    v4 -= v4
    v27 := GenS1{F0: 65535}
    fmt.Println(v4, v22)
    if ((-2.25 * float32(0.1)) == 3.14159) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v4, v22)
    }
    switch "a" {
    case "\n\t":
        var v28 int = (v22 - v22)
        _ = v28
    default:
        v4 = ((v4 / (1.5 * 0.1)) / 1e10)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v4
}

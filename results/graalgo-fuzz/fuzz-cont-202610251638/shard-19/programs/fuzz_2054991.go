package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 uint32, p1 int16) uint8 {
    var v7 uint32 = 668
    var v8 uint32 = (p0 | 410)
    v8 += 2
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 138
}
func genFn1(p0 int8) int16 {
    p0 -= p0
    var v9 string = "foo bar"
    if false {
        v10 := GenS1{F0: 18446744073709551615}
        _ = v10
    }
    var v11 uint = 4294967295
    _ = p0
    _ = v11
    _ = v9
    return 32767
}

func main() {
    var v1 bool = false
    var v2 int8 = 127
    v3 := make(map[int]bool)
    var v4 uint16 = 1
    var v5 string = (("" + "\n\t") + ((("foo bar" + ("foo bar" + ("\n\t" + (((("" + "\n\t") + "foo bar") + (("foo bar" + "hello") + "")) + ("\n\t" + "hello"))))) + "hello") + "hello"))
    v6 := make(map[int]int)
    v12 := []int{-1, -1, 470, 127}
    if v1 {
        var v13 float64 = 0.0
        var v14 int = 921
        v14 = (v14 >> ((1) & 63))
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: 624}
    v4 += (v4 & 42)
    for fi16 := 0; fi16 < 16; fi16++ {
        switch (fi16 | fi16) {
        case 1:
            switch v5 {
            case "":
                fmt.Println(v4, fi16)
            default:
                switch v5 {
                case "hello":
                    var v17 int = (fi16 ^ (fi16 / (((^(fi16))) | 1)))
                    _ = v17
                case "\n\t":
                    if v1 {
                        var v18 uint = 273
                        _ = len(v12)
                        _ = v18
                    }
                case "foo bar":
                    v19 := GenS1{F0: 1}
                    _ = v19
                default:
                    fmt.Println(v2, v2)
                }
            }
        case 0:
            var v20 float64 = float64(1e10)
            _ = v20
        }
        var v21 uint16 = v4
        v5 = "a"
        _ = fi16
        _ = v21
    }
    v2 += v2
    var v22 uint = 1
    var v23 float64 = 1.5
    var v24 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

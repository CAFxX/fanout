package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
func genFn0(p0 int) int32 {
    v6 := GenS1{F0: -32768}
    p0 -= (p0 * p0)
    var v7 int32 = -81
    v8 := []int{}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return v7
}
func genFn1(p0 bool) uint64 {
    p0 = false
    var v9 bool = (!p0)
    _ = p0
    _ = v9
    return 1
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]bool)
    var v3 int8 = 13
    var v4 float64 = 3.14159
    v5 := []int{32768}
    v3 -= v3
    v4 += 1e10
    fmt.Println(v3, v4)
    for fi10 := 0; fi10 < 15; fi10++ {
        switch ((("\n\t" + "foo bar") + "foo bar") + ("" + "\n\t")) {
        case "":
            for fi11 := 0; fi11 < 6; fi11++ {
                _ = len(v5)
                v3 = v3
                switch ("hello" + "foo bar") {
                case "a":
                    var v12 int16 = -32768
                    _ = v12
                case "\n\t":
                    fmt.Println(fi10)
                case "foo bar":
                    var v13 uint64 = 42
                    _ = v13
                default:
                    v3 -= -105
                }
                v5 = append(v5, 32768)
                var v14 int8 = (v3 | (v3 | -128))
                _ = fi11
                _ = v14
            }
        }
        _ = fi10
    }
    var v15 uint8 = 151
    var v16 bool = false
    for fi17 := 0; fi17 < 1; fi17++ {
        var v18 bool = v16
        v19 := GenS2{F0: fi17}
        _ = fi17
        _ = v18
        _ = v19
    }
    v15 = v15
    for fi20 := 0; fi20 < 10; fi20++ {
        var v21 uint64 = 1
        _ = len(v5)
        _ = fi20
        _ = v21
    }
    if v16 {
        for fi22 := 0; fi22 < 3; fi22++ {
            for fi23 := 0; fi23 < 2; fi23++ {
                v24 := GenS1{F0: 89}
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
    }
    v16 = ((!(2147483647 == 1)) || (!(("\n\t" + "foo bar") == "a")))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

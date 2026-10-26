package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint8 }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint64, p1 uint16, p2 int8) uint16 {
    var v4 uint64 = p0
    v5 := GenS1{F0: (p0 + (^(p0))), F1: true, F2: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p1
}
func genFn1(p0 int64, p1 int32, p2 string) string {
    var v6 int16 = 1
    for fi7 := 0; fi7 < 3; fi7++ {
        p0 = p0
        v8 := GenS2{F0: 42}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "foo bar"
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 0
    var v3 uint64 = 1
    var v9 string = ("a" + "")
    v2 += -1
    if (v9 < "a") {
        var v10 string = v9
        _ = v10
    }
    if false {
        var v11 int16 = -32768
        fmt.Println(v3, v2)
        _ = v11
    } else {
        var v12 int16 = 1000
        _ = v12
    }
    switch -2147483648 {
    case 2147483647:
        v13 := []int{0, 766, 128}
        _ = v13
    case 659:
        for fi14 := 0; fi14 < 15; fi14++ {
            var v15 uint16 = 0
            if false {
                var v16 int32 = 2147483647
                _ = v16
            }
            var v17 int64 = v2
            _ = fi14
            _ = v15
            _ = v17
        }
    default:
        var v18 int64 = v2
        _ = v18
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        v3 = v3
        _ = fi19
    }
    v3 = 809
    var v20 int64 = (v2 + 1)
    switch 127 {
    case 196:
        v9 = v9
    case 32767:
        var v21 int64 = 0
        _ = v21
    case 42:
        for fi22 := 0; fi22 < 18; fi22++ {
            fmt.Println(v20, fi22)
            if false {
                v9 = v9
                _ = len(v1)
                var v23 string = ((((v9 + v9) + (v9 + "a")) + v9) + ("\n\t" + (v9 + "hello")))
                _ = v23
            }
            switch (fi22 & fi22) {
            case 2147483647:
                fmt.Println(v3)
            case 42:
                for fi24 := 0; fi24 < 14; fi24++ {
                    switch v9 {
                    case "\n\t":
                        v25 := []int{}
                        _ = v25
                    default:
                        fmt.Println(v9, v2, v9)
                    }
                    if ((fi24 / (((-(fi22))) | 1)) <= fi22) {
                        var v26 uint8 = 0
                        _ = v26
                    }
                    _ = fi24
                }
            }
            _ = fi22
        }
    default:
        var v27 int64 = (^(v20))
        _ = v27
    }
    for fi28 := 0; fi28 < 10; fi28++ {
        v29 := GenS2{F0: -2147483648}
        _ = fi28
        _ = v29
    }
    fmt.Println(v20, v20, v20)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v9
}

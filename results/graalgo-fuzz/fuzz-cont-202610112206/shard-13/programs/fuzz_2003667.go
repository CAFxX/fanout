package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint }
func genFn0(p0 int, p1 int) int {
    p1 = (p1 * 2147483647)
    v5 := []int{}
    if true {
        var v6 uint32 = 4294967295
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v5
    return p0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = 100
    var v3 int = -2147483648
    v4 := make(map[string]int)
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 13; fi8++ {
            v9 := GenS1{F0: 127}
            var v10 uint16 = 0
            _ = fi8
            _ = v10
            _ = v9
        }
        v3 = fi7
        _ = fi7
    }
    var v11 uint32 = 42
    for fi12 := 0; fi12 < 10; fi12++ {
        v13 := []int{fi12, 790, genFn0((^(fi12)), 127), fi12}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        var v15 uint8 = 161
        _ = fi14
        _ = v15
    }
    var v16 int32 = 2147483647
    fmt.Println(v11)
    v2 = v2
    for fi17 := 0; fi17 < 19; fi17++ {
        v1["hello"] = v11
        if true {
            v3 = (^(v3))
            v18 := []int{genFn0(v3, genFn0(-420, 42)), v3, (v3 | genFn0((v3 % (((fi17 - -2147483648)) | 1)), v3)), (fi17 >> ((525) & 63))}
            _ = v18
        } else {
            v3 -= (fi17 << ((42) & 63))
            fmt.Println(v16)
            v2 -= -59
        }
        if (282 == -492) {
            fmt.Println(v3)
            delete(v1, "foo bar")
            for fi19 := 0; fi19 < 9; fi19++ {
                var v20 bool = ((v2 >> ((2) & 63)) == 44)
                _ = fi19
                _ = v20
            }
        } else {
            var v21 int64 = -150
            _ = v21
        }
        _ = fi17
    }
    v22 := GenS1{F0: 255}
    v23 := []int{v3}
    v1["hello"] = 4294967295
    var v24 uint8 = 127
    for fi25 := 0; fi25 < 17; fi25++ {
        var v26 string = "\n\t"
        v27 := GenS1{F0: 130}
        v28 := GenS2{F0: fi25, F1: 42}
        _ = fi25
        _ = v26
        _ = v27
        _ = v28
    }
    var v29 uint8 = 83
    var v30 string = "\n\t"
    v29 += (v29 | 135)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}

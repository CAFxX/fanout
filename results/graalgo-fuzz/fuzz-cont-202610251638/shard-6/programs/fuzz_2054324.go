package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint16, p1 bool, p2 int16) int32 {
    if p1 {
        fmt.Println(p2)
        v6 := []int{}
        _ = v6
    }
    if (true && (134 > 42)) {
        if p1 {
            p2 += 42
            v7 := []int{-9223372036854775808, -9223372036854775808, 32768}
            _ = v7
        }
        if false {
            v8 := []int{42, 2147483647}
            _ = v8
        }
    } else {
        fmt.Println(p1, p0, p0)
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 uint64 = 0
        p2 += p2
        if p1 {
            v11 := GenS1{F0: p1}
            for fi12 := 0; fi12 < 11; fi12++ {
                var v13 uint16 = p0
                _ = fi12
                _ = v13
            }
            _ = v11
        } else {
            p0 = p0
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}
func genFn1(p0 int16) string {
    for fi14 := 0; fi14 < 8; fi14++ {
        v15 := GenS1{F0: false}
        fmt.Println(fi14, p0)
        _ = fi14
        _ = v15
    }
    var v16 int64 = 1
    _ = p0
    _ = v16
    return "\n\t"
}

func main() {
    var v1 float64 = 0.0
    var v2 uint16 = 66
    var v3 uint32 = 2
    v4 := make(map[int]int)
    var v5 string = (((("hello" + "") + "") + ("a" + "foo bar")) + "foo bar")
    v1 -= 1.5
    v17 := []int{-9223372036854775808, -520, -971}
    v18 := GenS1{F0: (false && false)}
    fmt.Println(v1)
    v5 = "a"
    v5 = ("" + v5)
    v19 := v4[58]
    v1 += (v1 - v1)
    if (-9223372036854775808 != -410) {
        var v20 string = v5
        var v21 string = ("a" + ("foo bar" + "foo bar"))
        for fi22 := 0; fi22 < 5; fi22++ {
            var v23 uint = 565
            _ = fi22
            _ = v23
        }
        _ = v20
        _ = v21
    } else {
        v1 = ((v1 / 0.1) - float64((-2.25 / 0.1)))
    }
    v5 = "a"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
func genFn0(p0 int8, p1 bool, p2 bool) uint32 {
    v6 := GenS1{F0: 127}
    var v7 bool = (((!p2) || p2) || p2)
    for fi8 := 0; fi8 < 13; fi8++ {
        v7 = ((403 >= 18446744073709551615) || ((p0 << ((962) & 63)) == p0))
        for fi9 := 0; fi9 < 20; fi9++ {
            fmt.Println(p0, fi8)
            var v10 uint16 = 0
            if (fi8 <= 635) {
                fmt.Println(fi9, v7, v7)
                v11 := []int{fi9, (fi8 - -406)}
                p2 = (-1 < -260)
                _ = v11
            } else {
                fmt.Println(p0)
            }
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        for fi13 := 0; fi13 < 14; fi13++ {
            p0 -= p0
            var v14 int8 = 42
            _ = fi13
            _ = v14
        }
        p0 = p0
        var v15 int64 = -9223372036854775808
        _ = fi12
        _ = v15
    }
    v16 := []int{2147483647, 960, 32767, 860}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 4294967295
    var v3 uint = v2
    var v4 int64 = 42
    var v5 int8 = 100
    v5 = v5
    if false {
        var v17 bool = false
        v18 := []int{}
        v19 := GenS2{F0: -1}
        _ = v17
        _ = v18
        _ = v19
    }
    v20 := []int{0, 2147483647, 141}
    var v21 int32 = 42
    v20 = append(v20, 820)
    fmt.Println(v4, v3, v21)
    switch 460 {
    case -690:
        switch ("" + "foo bar") {
        case "hello":
            for fi22 := 0; fi22 < 17; fi22++ {
                var v23 int16 = 32767
                if false {
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                    switch ("" + "") {
                    case "":
                        var v24 uint8 = 149
                        _ = v24
                    case "a":
                        v4 = v4
                    }
                    var v25 int8 = 42
                    _ = v25
                } else {
                    if len(v20) > 0 {
                        v20[0] = (-(fi22))
                    }
                    var v26 float64 = (0.0 / ((1.5 - 3.14159) + 3.14159))
                    v27 := GenS2{F0: 42}
                    _ = v26
                    _ = v27
                }
                _ = fi22
                _ = v23
            }
        }
    default:
        for fi28 := 0; fi28 < 15; fi28++ {
            fmt.Println(v21, v5)
            _ = fi28
        }
    }
    var v29 float32 = 0.1
    var v30 uint = v3
    fmt.Println(v4, v30)
    var v31 int8 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}

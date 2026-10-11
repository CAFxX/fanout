package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int16, p1 int16) uint32 {
    var v7 int = 945
    v8 := []int{(-(v7)), (v7 + 4294967295)}
    v9 := GenS1{F0: -128, F1: 42}
    var v10 int16 = (p0 | 1000)
    v8 = append(v8, 2147483647)
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 uint8, p1 uint64, p2 int32) int16 {
    for fi11 := 0; fi11 < 4; fi11++ {
        p2 += p2
        _ = fi11
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        if (-84 >= -128) {
            for fi13 := 0; fi13 < 19; fi13++ {
                v14 := GenS1{F0: 100, F1: 42}
                v15 := []int{fi13, fi13}
                _ = fi13
                _ = v14
                _ = v15
            }
            var v16 uint = 4294967295
            var v17 float32 = float32((3.14159 + ((-2.25 / ((float64(3.14159) + ((0.1 * (0.1 / 1.0)) / 1.0)) + 0.1)) - (1e10 + 1.5))))
            _ = v16
            _ = v17
        }
        v18 := GenS1{F0: 42, F1: 2}
        _ = fi12
        _ = v18
    }
    switch 4294967295 {
    case 32767:
        v19 := GenS1{F0: 100, F1: 421}
        _ = v19
    default:
        var v20 uint = 1
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    return 32767
}

func main() {
    v1 := []int{-1}
    var v2 uint16 = 0
    v3 := make(map[int]string)
    var v4 uint16 = v2
    var v5 uint16 = v4
    v6 := []int{}
    for fi21 := 0; fi21 < 5; fi21++ {
        for fi22 := 0; fi22 < 4; fi22++ {
            var v23 bool = false
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if ("" >= "hello") {
                v5 += 42
                delete(v3, fi22)
                var v24 int8 = 125
                _ = v24
            } else {
                if false {
                    switch (fi22 % ((fi21) | 1)) {
                    case 4294967295:
                        _ = len(v3)
                    }
                }
                var v25 uint = uint(v4)
                _ = v25
            }
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    switch "a" {
    case "\n\t":
        v1 = append(v1, 42)
    case "":
        v2 = (v4 >> ((18446744073709551615) & 63))
    case "hello":
        for fi26 := 0; fi26 < 17; fi26++ {
            v27 := v3[fi26]
            var v28 int = -317
            _ = fi26
            _ = v27
            _ = v28
        }
    default:
        v4 -= 449
    }
    v2 = 757
    for fi29 := 0; fi29 < 3; fi29++ {
        v4 = v2
        _ = fi29
    }
    var v30 bool = (!(0.1 > 0.0))
    switch "hello" {
    case "hello":
        fmt.Println(v4, v5, v30)
    }
    var v31 int16 = 32767
    v5 = (v4 | 2)
    v32 := GenS2{F0: -1}
    v4 = (v5 >> ((42) & 63))
    var v33 int16 = v31
    for fi34 := 0; fi34 < 2; fi34++ {
        if v30 {
            if true {
                for fi35 := 0; fi35 < 1; fi35++ {
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v31, v2)
                    case "\n\t":
                        fmt.Println(v33)
                    }
                    switch ("hello" + "hello") {
                    case "\n\t":
                        var v36 string = "hello"
                        _ = v36
                    default:
                        if v30 {
                            var v37 string = ("a" + "\n\t")
                            v4 = v5
                            _ = v37
                        }
                    }
                    v38 := v3[fi35]
                    _ = fi35
                    _ = v38
                }
            }
            v33 = genFn1(255, 42, -838)
        } else {
            v5 -= 1
            v39 := GenS1{F0: -1, F1: 255}
            _ = v39
        }
        fmt.Println(v5)
        _ = len(v3)
        _ = fi34
    }
    var v40 uint8 = 127
    fmt.Println(v4, v40)
    if v30 {
        for fi41 := 0; fi41 < 6; fi41++ {
            v42 := GenS2{F0: 42}
            var v43 int32 = 120
            fmt.Println(fi41, v30, v33)
            _ = fi41
            _ = v42
            _ = v43
        }
    }
    var v44 uint16 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v40
    _ = v44
    _ = v5
    _ = v6
}

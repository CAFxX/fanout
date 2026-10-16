package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int }
func genFn0(p0 int, p1 uint32, p2 bool) int {
    fmt.Println(p0)
    if (1 != -1) {
        switch ("" + "hello") {
        case "\n\t":
            p0 = -1
        case "foo bar":
            var v5 int8 = -128
            _ = v5
        case "hello":
            var v6 int64 = int64(p0)
            _ = v6
        default:
            for fi7 := 0; fi7 < 19; fi7++ {
                p1 = (p1 - p1)
                v8 := []int{fi7, (p0 ^ 127), (-(fi7)), 128}
                _ = fi7
                _ = v8
            }
        }
    }
    v9 := GenS1{F0: -1, F1: 508, F2: p0}
    var v10 uint = 18446744073709551615
    var v11 int32 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return 0
}
func genFn1(p0 uint8, p1 uint8, p2 uint) int {
    if true {
        for fi12 := 0; fi12 < 9; fi12++ {
            if (1000 > int16(p2)) {
                p2 += p2
                fmt.Println(p0)
                var v13 int64 = -9223372036854775808
                _ = v13
            }
            _ = fi12
        }
    }
    var v14 bool = true
    fmt.Println(p2, v14, v14)
    v15 := []int{592, -182}
    for fi16 := 0; fi16 < 1; fi16++ {
        switch fi16 {
        case -9223372036854775808:
            if (!(!((0 >= -804) || ("hello" != ((("foo bar" + ("" + "\n\t")) + "foo bar") + ("" + "hello")))))) {
                var v17 int64 = 9223372036854775807
                for fi18 := 0; fi18 < 3; fi18++ {
                    if true {
                        v19 := GenS1{F0: 32767, F1: 0, F2: fi16}
                        _ = v19
                    }
                    if (v14 || ((v14 && v14) || (3.14159 != 1e10))) {
                        var v20 float64 = float64(0.1)
                        v21 := GenS1{F0: 42, F1: 1, F2: (fi18 & (fi18 ^ -734))}
                        var v22 int64 = (v17 << (((p2 / ((2) | 1))) & 63))
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    var v23 uint16 = 5
                    _ = fi18
                    _ = v23
                }
                v15 = append(v15, -593)
                _ = v17
            }
        case 269:
            var v24 int64 = 1
            _ = v24
        case 0:
            var v25 uint32 = 2147483647
            _ = v25
        }
        _ = fi16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return genFn0(128, 2147483648, false)
}

func main() {
    var v1 uint = 1
    v2 := make(map[uint32]uint32)
    var v3 int32 = -1
    v4 := make(map[string]int)
    v26 := v4["a"]
    if true {
        v27 := []int{v26}
        _ = v27
    }
    var v28 float64 = float64(3.14159)
    v29 := GenS1{F0: 0, F1: 32767, F2: (^(v26))}
    v26 = v26
    v1 += v1
    var v30 uint = (v1 + 173)
    v31 := GenS1{F0: -63, F1: 1, F2: 32767}
    v32 := []int{}
    var v33 int32 = int32(v1)
    switch (("" + "foo bar") + "foo bar") {
    case "foo bar":
        if true {
            v32 = append(v32, genFn0((-(v26)), 0, (true || (113 < 1))))
            if true {
                var v34 uint16 = 1
                v1 += 2
                _ = v34
            } else {
                fmt.Println(v33)
            }
        } else {
            var v35 uint8 = 127
            v4["hello"] = (^(v26))
            switch v26 {
            case 0:
                v36 := GenS1{F0: -1, F1: 1, F2: v26}
                _ = v36
            default:
                v37 := GenS1{F0: 1000, F1: 266, F2: v26}
                _ = v37
            }
            _ = v35
        }
    default:
        v33 = v3
    }
    fmt.Println(v28)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}

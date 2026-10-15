package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 uint64; F2 uint32 }
func genFn0(p0 uint, p1 string) string {
    switch p1 {
    case "foo bar":
        v5 := GenS1{F0: -21}
        _ = v5
    default:
        p0 -= 0
    }
    var v6 float32 = (0.0 + 0.0)
    for fi7 := 0; fi7 < 10; fi7++ {
        v6 -= 1.5
        _ = fi7
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 float32 = (v6 * 3.14159)
        var v10 int64 = 1
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    return ""
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]string)
    v3 := make(map[string]string)
    v4 := make(map[uint32]bool)
    v1[65535] = 811
    var v11 float64 = ((-2.25 + float64(((1e10 - (1e10 / 3.14159)) * 1.5))) - (float64(-2.25) + (0.1 / (0.0 - (1.5 + 1e10)))))
    for fi12 := 0; fi12 < 5; fi12++ {
        if true {
            for fi13 := 0; fi13 < 9; fi13++ {
                var v14 int64 = 42
                _ = fi13
                _ = v14
            }
        }
        _ = fi12
    }
    v15 := GenS1{F0: 1}
    v16 := GenS1{F0: 2147483647}
    var v17 int32 = 0
    var v18 int = -9223372036854775808
    if ("a" != ("\n\t" + "\n\t")) {
        switch (v18 - v18) {
        case -309:
            if false {
                v19 := GenS1{F0: v18}
                delete(v2, 500)
                if (1 < 2) {
                    for fi20 := 0; fi20 < 14; fi20++ {
                        fmt.Println(v18)
                        v17 = (-(v17))
                        _ = fi20
                    }
                }
                _ = v19
            }
        }
    }
    var v21 int32 = v17
    for fi22 := 0; fi22 < 19; fi22++ {
        fmt.Println(v18, v21)
        _ = fi22
    }
    v23 := []int{(-(v18)), v18, v18}
    v24 := GenS1{F0: 1}
    var v25 uint8 = 44
    v26 := GenS2{F0: v18, F1: 40, F2: 780}
    for fi27 := 0; fi27 < 3; fi27++ {
        if (2 != 32767) {
            for fi28 := 0; fi28 < 15; fi28++ {
                v29 := v1[460]
                var v30 int32 = 476
                _ = fi28
                _ = v29
                _ = v30
            }
        } else {
            switch int(v17) {
            case 1:
                v21 += 2147483647
            case 32768:
                switch ("" + "") {
                case "":
                    v23 = append(v23, (^(v18)))
                default:
                    v31 := GenS1{F0: (fi27 & (v18 << ((65535) & 63)))}
                    _ = v31
                }
            }
        }
        if (v25 < 42) {
            switch "hello" {
            case "hello":
                var v32 int8 = -50
                _ = v32
            case "foo bar":
                v23 = append(v23, fi27)
            case "":
                v18 += (v18 >> ((42) & 63))
            default:
                if (42 != 9223372036854775807) {
                    switch fi27 {
                    case 223:
                        v21 -= v17
                    }
                } else {
                    fmt.Println(v21, fi27)
                }
            }
            v17 = v21
        } else {
            v33 := GenS1{F0: fi27}
            v23 = append(v23, int(v25))
            _ = v33
        }
        _ = fi27
    }
    v21 = (^(v21))
    var v34 float64 = v11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v34
    _ = v4
}

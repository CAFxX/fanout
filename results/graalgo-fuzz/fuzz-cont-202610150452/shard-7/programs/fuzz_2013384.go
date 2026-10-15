package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int16 }
type GenS2 struct { F0 int32; F1 int; F2 int }
func genFn0(p0 int8, p1 int, p2 uint64) uint {
    switch (-(p1)) {
    case -1:
        switch "hello" {
        case "":
            var v7 uint64 = p2
            _ = v7
        }
    }
    var v8 float32 = (1.5 / (-2.25 / 1e10))
    var v9 int8 = 100
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 4294967295
}
func genFn1(p0 uint64, p1 bool, p2 int) bool {
    for fi10 := 0; fi10 < 9; fi10++ {
        p1 = (1 == 2)
        for fi11 := 0; fi11 < 1; fi11++ {
            var v12 float64 = -2.25
            v13 := []int{}
            if p1 {
                if p1 {
                    v14 := GenS1{F0: 1, F1: p1, F2: 1}
                    v15 := GenS1{F0: 1, F1: (42 >= 2147483647), F2: -29}
                    _ = v14
                    _ = v15
                }
                if true {
                    v12 -= (float64(3.14159) * -2.25)
                    p0 += (p0 ^ (p0 + (p0 | (p0 | p0))))
                } else {
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    v16 := GenS1{F0: -564, F1: (!p1), F2: 32767}
                    p2 -= 2147483647
                    _ = v16
                }
            }
            switch "" {
            case "\n\t":
                fmt.Println(p0, fi10)
            case "":
                switch (fi10 & fi11) {
                case -2147483648:
                    p1 = false
                case 810:
                    if len(v13) > 0 {
                        v13[0] = p2
                    }
                }
            }
            switch fi10 {
            case 701:
                v12 = ((float64(1e10) / v12) + v12)
            case 0:
                v17 := GenS2{F0: 2147483647, F1: 32768, F2: p2}
                _ = v17
            }
            switch -21 {
            case 1:
                if false {
                    var v18 int = p2
                    if p1 {
                        var v19 uint32 = 2
                        _ = v19
                    }
                    _ = v18
                }
            default:
                _ = len(v13)
            }
            _ = fi11
            _ = v12
            _ = v13
        }
        for fi20 := 0; fi20 < 18; fi20++ {
            p1 = (0.1 <= (1.5 * (0.1 / float64((0.0 + (3.14159 + ((0.1 - 3.14159) - 0.0)))))))
            switch "hello" {
            case "foo bar":
                v21 := []int{}
                _ = v21
            default:
                if p1 {
                    var v22 float32 = (-2.25 / (3.14159 - (1e10 + 1e10)))
                    _ = v22
                } else {
                    var v23 string = ("a" + "\n\t")
                    if p1 {
                        var v24 uint8 = 127
                        _ = v24
                    }
                    _ = v23
                }
            }
            _ = fi20
        }
        _ = fi10
    }
    p0 += p0
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    var v1 int16 = 1
    var v2 uint32 = 0
    var v3 string = (("a" + "hello") + ("hello" + ("" + "foo bar")))
    v4 := make(map[string]int)
    v5 := make(map[string]uint32)
    var v6 int8 = -101
    delete(v4, "foo bar")
    v25 := []int{2147483647, -726, 179}
    var v26 uint8 = 128
    v2 = (v2 & 2147483648)
    var v27 float64 = -2.25
    v6 += -1
    for fi28 := 0; fi28 < 6; fi28++ {
        v29 := v5[v3]
        _ = fi28
        _ = v29
    }
    v1 = 913
    fmt.Println(v1)
    v30 := GenS1{F0: -937, F1: (466 >= 87), F2: (v1 << ((1) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}

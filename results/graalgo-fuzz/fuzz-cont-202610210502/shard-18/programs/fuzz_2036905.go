package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint8, p1 bool, p2 int8) uint16 {
    var v5 int16 = -851
    v5 -= 1
    v5 = (v5 * v5)
    var v6 uint16 = 65535
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            if p1 {
                if p1 {
                    v9 := GenS1{F0: (v5 << ((65535) & 63)), F1: 0}
                    v10 := GenS1{F0: 32767, F1: int(v6)}
                    v11 := []int{}
                    _ = v10
                    _ = v11
                    _ = v9
                } else {
                    v5 -= v5
                    fmt.Println(p0, p1)
                }
                var v12 uint32 = 718
                fmt.Println(p0)
                _ = v12
            }
            fmt.Println(v5)
            switch 249 {
            case 127:
                if p1 {
                    if p1 {
                        var v13 uint = 255
                        _ = v13
                    }
                    v6 += 2
                    v14 := GenS1{F0: (v5 | (v5 * 32767)), F1: fi7}
                    _ = v14
                }
                if p1 {
                    p2 -= (p2 - (-(p2)))
                }
            }
            var v15 string = "\n\t"
            v16 := []int{}
            _ = fi8
            _ = v15
            _ = v16
        }
        p1 = true
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 255
}

func main() {
    v1 := []int{-2147483648}
    v2 := []uint32{65535}
    v3 := make(map[string]int)
    v4 := []int{}
    if true {
        var v17 string = "hello"
        if false {
            if (false && ((1 == 2147483648) && (255 < 42))) {
                var v18 int64 = 0
                fmt.Println(v17, v18)
                _ = v18
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = v17
    }
    var v19 int16 = 753
    if true {
        if false {
            for fi20 := 0; fi20 < 8; fi20++ {
                v21 := v3[""]
                var v22 int = fi20
                _ = fi20
                _ = v21
                _ = v22
            }
        } else {
            v19 -= (v19 ^ v19)
            for fi23 := 0; fi23 < 6; fi23++ {
                var v24 int = 32767
                v25 := v3[("foo bar" + "\n\t")]
                v19 = 32767
                _ = fi23
                _ = v24
                _ = v25
            }
        }
        delete(v3, "foo bar")
        v26 := GenS1{F0: v19, F1: 32767}
        _ = v26
    }
    var v27 int16 = 1000
    for fi28 := 0; fi28 < 11; fi28++ {
        v29 := GenS2{F0: 127}
        _ = fi28
        _ = v29
    }
    v19 = v27
    v27 -= 810
    for fi30 := 0; fi30 < 12; fi30++ {
        v27 = v19
        _ = fi30
    }
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}

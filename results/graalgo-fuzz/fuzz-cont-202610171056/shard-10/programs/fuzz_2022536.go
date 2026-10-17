package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 uint; F2 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 bool }
func genFn0(p0 int, p1 int) uint32 {
    fmt.Println(p1, p1)
    p1 = p0
    if false {
        fmt.Println(p0, p0)
    } else {
        for fi6 := 0; fi6 < 4; fi6++ {
            p0 += p0
            for fi7 := 0; fi7 < 1; fi7++ {
                fmt.Println(p1, p0)
                if false {
                    var v8 uint32 = 0
                    v9 := []int{332, p1}
                    _ = v8
                    _ = v9
                } else {
                    v10 := []int{(fi6 + p1), (fi7 >> ((0) & 63))}
                    if (true || false) {
                        var v11 int64 = 42
                        v12 := GenS1{F0: 2}
                        v13 := GenS1{F0: 65535}
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                    var v14 uint32 = 95
                    _ = v10
                    _ = v14
                }
                _ = fi7
            }
            _ = fi6
        }
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := []int{}
    var v2 int16 = 0
    var v3 int64 = -1
    v4 := []string{("" + ""), "hello", "foo bar"}
    v5 := []uint32{2147483647}
    if true {
        var v15 int8 = 100
        _ = v15
    } else {
        v16 := GenS1{F0: 18446744073709551615}
        _ = v16
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        for fi18 := 0; fi18 < 9; fi18++ {
            v3 = v3
            v19 := GenS2{F0: 95, F1: 725, F2: 0}
            if len(v5) > 0 {
                v5[0] = 391
            }
            v20 := GenS1{F0: 6}
            _ = fi18
            _ = v19
            _ = v20
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi17
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 int = 1
        _ = fi21
        _ = v22
    }
    var v23 uint = 1
    v23 = 690
    v24 := GenS1{F0: (v23 >> (((v23 & (v23 >> ((v23) & 63)))) & 63))}
    if ((0.0 + -2.25) > ((float32(1.5) * (float32((1e10 / 1.5)) * -2.25)) * -2.25)) {
        var v25 int16 = (v2 / ((v2) | 1))
        v23 -= 255
        v26 := GenS3{F0: 160}
        _ = v25
        _ = v26
    }
    var v27 uint8 = 2
    v3 += (v3 >> (((v23 & v23)) & 63))
    var v28 int16 = v2
    v29 := GenS4{F0: 42, F1: true}
    if len(v5) > 0 {
        v5[0] = 2147483648
    }
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}

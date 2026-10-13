package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint }
func genFn0(p0 int64) int {
    for fi4 := 0; fi4 < 2; fi4++ {
        for fi5 := 0; fi5 < 18; fi5++ {
            var v6 int = (fi5 * fi5)
            var v7 uint64 = 985
            _ = fi5
            _ = v6
            _ = v7
        }
        v8 := []int{fi4, (fi4 % ((fi4) | 1)), fi4, (fi4 % ((fi4) | 1))}
        if true {
            for fi9 := 0; fi9 < 20; fi9++ {
                fmt.Println(fi9, p0)
                if len(v8) > 0 {
                    v8[0] = (fi9 & (fi4 * 846))
                }
                switch fi9 {
                case 1:
                    v10 := GenS1{F0: 4294967295, F1: (fi9 - fi4)}
                    _ = v10
                }
                _ = fi9
            }
            for fi11 := 0; fi11 < 19; fi11++ {
                p0 = (p0 ^ p0)
                fmt.Println(fi4, fi4)
                fmt.Println(fi11, fi4, fi4)
                _ = fi11
            }
        } else {
            if true {
                p0 -= 0
            }
            var v12 uint32 = 2
            fmt.Println(fi4, fi4)
            _ = v12
        }
        _ = fi4
        _ = v8
    }
    v13 := []int{32768, -546, 32767}
    _ = p0
    _ = v13
    return 128
}
func genFn1(p0 bool, p1 int64) int32 {
    var v14 uint32 = 773
    p1 = (-(p1))
    if ((!(132 > 1)) && p0) {
        switch "\n\t" {
        case "a":
            fmt.Println(p1, p1)
        case "":
            for fi15 := 0; fi15 < 4; fi15++ {
                v16 := []int{genFn0(307), fi15, fi15}
                _ = fi15
                _ = v16
            }
        }
    } else {
        for fi17 := 0; fi17 < 5; fi17++ {
            if p0 {
                var v18 uint16 = 2
                fmt.Println(p0)
                fmt.Println(v18, p0, v18)
                _ = v18
            } else {
                v19 := GenS1{F0: 255, F1: fi17}
                _ = v19
            }
            _ = fi17
        }
        p0 = (!(368 > -70))
        p0 = false
    }
    _ = p0
    _ = p1
    _ = v14
    return 42
}

func main() {
    var v1 bool = true
    var v2 float32 = (((0.1 + ((0.0 + 0.0) / 1.5)) * 0.0) * 1.5)
    v3 := make(map[int]uint32)
    for fi20 := 0; fi20 < 18; fi20++ {
        var v21 int = genFn0(589)
        for fi22 := 0; fi22 < 12; fi22++ {
            switch "hello" {
            case "hello":
                if true {
                    if v1 {
                        v23 := GenS1{F0: 65535, F1: 1}
                        _ = v23
                    } else {
                        v2 += v2
                        if false {
                            var v24 uint = 2
                            var v25 string = "\n\t"
                            _ = v24
                            _ = v25
                        } else {
                            v26 := GenS2{F0: 395}
                            _ = v26
                        }
                    }
                    v2 = 3.14159
                    if (433 < 0) {
                        fmt.Println(v21)
                    }
                } else {
                    var v27 uint32 = 1
                    _ = v27
                }
            case "a":
                v1 = (2 < 2)
            default:
                var v28 uint8 = 25
                _ = v28
            }
            _ = fi22
        }
        _ = fi20
        _ = v21
    }
    v1 = (!v1)
    v2 = v2
    v29 := []int{-1, 32768}
    var v30 int = 4294967295
    if len(v29) >= 2 {
        v29 = v29[1:len(v29)]
    }
    v2 += 0.0
    v31 := GenS2{F0: 2}
    fmt.Println(v1, v30)
    var v32 string = "hello"
    v30 = genFn0(-561)
    var v33 float32 = (v2 + float32(1.5))
    v34 := v3[genFn0(-396)]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
}

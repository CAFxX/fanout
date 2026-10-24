package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint32, p1 int64, p2 int) bool {
    if true {
        if false {
            var v4 string = "\n\t"
            v5 := GenS1{F0: 42}
            var v6 uint16 = 0
            _ = v4
            _ = v5
            _ = v6
        }
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := []int{p2, fi7, fi7, fi7}
            var v9 bool = (!false)
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 int32 = -630
        _ = v10
    } else {
        v11 := GenS2{F0: (p0 >> ((4294967295) & 63)), F1: -92}
        _ = v11
    }
    if false {
        for fi12 := 0; fi12 < 5; fi12++ {
            fmt.Println(p0, p1)
            _ = fi12
        }
        v13 := []int{}
        v14 := GenS1{F0: 0}
        _ = v13
        _ = v14
    } else {
        v15 := []int{(p2 & p2), p2, 223, (p2 % ((p2) | 1))}
        if false {
            for fi16 := 0; fi16 < 17; fi16++ {
                switch ("foo bar" + "a") {
                case "hello":
                    for fi17 := 0; fi17 < 5; fi17++ {
                        var v18 float32 = (1e10 / 3.14159)
                        if true {
                            if (!(((("a" + "foo bar") + "hello") + ("foo bar" + "")) >= "")) {
                                var v19 int = (fi16 & p2)
                                v20 := GenS2{F0: p0, F1: -47}
                                v19 += -103
                                _ = v19
                                _ = v20
                            }
                        }
                        _ = fi17
                        _ = v18
                    }
                default:
                    p1 += (-(p1))
                }
                _ = fi16
            }
        }
        _ = v15
    }
    fmt.Println(p2, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    return ((p1 & 42) < 530)
}

func main() {
    var v1 int8 = -20
    var v2 int16 = int16(v1)
    var v3 float64 = 0.0
    var v21 uint64 = 1
    fmt.Println(v3)
    switch -2147483648 {
    case 4294967295:
        switch ("\n\t" + "") {
        case "\n\t":
            v22 := GenS1{F0: v21}
            _ = v22
        case "":
            var v23 float64 = v3
            _ = v23
        case "a":
            switch "\n\t" {
            case "hello":
                v24 := []int{464, 4294967295, -1}
                _ = v24
            }
        }
    case 32767:
        var v25 int64 = 626
        _ = v25
    case 0:
        fmt.Println(v1)
    default:
        v26 := []int{-977, 468, 32768, -2147483648}
        _ = v26
    }
    for fi27 := 0; fi27 < 11; fi27++ {
        v28 := []int{}
        var v29 uint8 = 215
        var v30 float64 = v3
        _ = fi27
        _ = v28
        _ = v29
        _ = v30
    }
    switch "" {
    case "hello":
        v2 = v2
    case "\n\t":
        var v31 uint8 = 2
        _ = v31
    default:
        for fi32 := 0; fi32 < 18; fi32++ {
            v21 += 1
            var v33 int8 = v1
            v21 -= v21
            _ = fi32
            _ = v33
        }
    }
    v34 := GenS3{F0: -1}
    v1 += v1
    v35 := []int{32767, 32768}
    if (447 > 0) {
        var v36 int = 220
        _ = v36
    }
    v21 += (v21 - v21)
    var v37 float32 = float32(float64(1.5))
    _ = len(v35)
    v38 := GenS2{F0: 4294967295, F1: (v1 % ((1) | 1))}
    v39 := GenS3{F0: 0}
    var v40 int8 = (v1 ^ (-(v1)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v37:", v37)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v34
    _ = v35
    _ = v37
    _ = v38
    _ = v39
    _ = v40
}

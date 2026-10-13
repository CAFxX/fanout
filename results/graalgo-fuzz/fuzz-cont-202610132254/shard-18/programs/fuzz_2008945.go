package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int64, p1 string, p2 uint) uint {
    var v6 int32 = -1
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 int8 = -1
        if true {
            switch (p1 + p1) {
            case "a":
                fmt.Println(p1)
            case "":
                v6 += (v6 | (v6 * v6))
            }
        }
        var v9 float64 = 1.5
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 uint8 = 1
        if false {
            switch fi10 {
            case 4294967295:
                p1 = "a"
            case 9223372036854775807:
                var v12 uint64 = 187
                _ = v12
            }
            var v13 int64 = p0
            _ = v13
        } else {
            if (true || false) {
                v6 += -2147483648
                p2 = 18446744073709551615
                switch fi10 {
                case -2147483648:
                    p2 = p2
                case 32767:
                    v14 := []int{}
                    _ = v14
                case 529:
                    if ((^(p0)) <= (p0 / (((p0 << (((^(p2))) & 63))) | 1))) {
                        for fi15 := 0; fi15 < 19; fi15++ {
                            v6 += -2147483648
                            v16 := []int{208, (fi15 % (((-(fi10))) | 1)), int(p2), (fi10 % ((fi15) | 1))}
                            _ = fi15
                            _ = v16
                        }
                        p2 += 4294967295
                        v17 := []int{(fi10 | (fi10 & (fi10 % ((fi10) | 1))))}
                        _ = v17
                    }
                default:
                    var v18 uint32 = 872
                    _ = v18
                }
            } else {
                var v19 bool = true
                fmt.Println(v6, v6)
                _ = v19
            }
            if true {
                v20 := []int{(fi10 << ((42) & 63)), (-(fi10))}
                for fi21 := 0; fi21 < 12; fi21++ {
                    var v22 int32 = v6
                    _ = fi21
                    _ = v22
                }
                _ = v20
            }
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p2 * 848)
}
func genFn1() bool {
    switch -476 {
    case 4294967295:
        switch "a" {
        case "hello":
            if true {
                v23 := []int{}
                var v24 float32 = (3.14159 * 1e10)
                _ = v23
                _ = v24
            } else {
                v25 := []int{2147483647, 32768, 0, 128}
                var v26 uint8 = 128
                _ = v25
                _ = v26
            }
        case "a":
            v27 := GenS1{F0: 1, F1: -52, F2: genFn0(211, (("foo bar" + "foo bar") + ""), 42)}
            _ = v27
        case "\n\t":
            var v28 uint8 = 1
            _ = v28
        default:
            var v29 string = ("a" + "")
            _ = v29
        }
    case 391:
        if false {
            v30 := []int{}
            v31 := GenS1{F0: 1, F1: 100, F2: 1}
            _ = v30
            _ = v31
        }
    case 127:
        var v32 string = "foo bar"
        _ = v32
    default:
        var v33 uint64 = 18446744073709551615
        _ = v33
    }
    v34 := GenS1{F0: 0, F1: -26, F2: 42}
    _ = v34
    return false
}

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = 0.1
    v3 := make(map[int]uint32)
    var v4 int16 = -1
    v5 := make(map[uint32]int)
    v2 += (v2 - 1e10)
    v35 := GenS2{F0: -2147483648, F1: 0}
    v36 := GenS2{F0: 1, F1: 2}
    v2 -= (((v2 / 1e10) + (v2 + v2)) - v2)
    var v37 bool = (442 < 472)
    var v38 string = ("a" + "")
    if v37 {
        v2 = ((((v2 * 0.1) * 1.5) + 1e10) + (float64(float32(v2)) * float64(0.0)))
        v39 := []int{476}
        var v40 int64 = -1
        _ = v39
        _ = v40
    }
    v1[-714] = (851 != 18446744073709551615)
    for fi41 := 0; fi41 < 5; fi41++ {
        v42 := GenS3{F0: 2}
        var v43 uint32 = 0
        _ = fi41
        _ = v42
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}

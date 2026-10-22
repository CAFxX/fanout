package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint; F2 bool }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 bool; F1 uint8; F2 uint64 }
func genFn0(p0 bool, p1 int8) int {
    for fi6 := 0; fi6 < 18; fi6++ {
        for fi7 := 0; fi7 < 3; fi7++ {
            p0 = (fi6 <= (fi7 & fi6))
            _ = fi7
        }
        v8 := GenS1{F0: (fi6 & fi6), F1: 0, F2: 127}
        if p0 {
            v9 := []int{fi6}
            p1 = 127
            _ = v9
        } else {
            p0 = (!p0)
            fmt.Println(p0, fi6)
        }
        _ = fi6
        _ = v8
    }
    if p0 {
        if p0 {
            v10 := []int{}
            _ = v10
        } else {
            v11 := GenS2{F0: 4294967295, F1: 746, F2: false}
            for fi12 := 0; fi12 < 9; fi12++ {
                var v13 int64 = 593
                _ = fi12
                _ = v13
            }
            _ = v11
        }
        var v14 uint = 317
        switch "foo bar" {
        case "":
            if false {
                var v15 int8 = (p1 ^ p1)
                if false {
                    if (255 <= 73) {
                        switch 0 {
                        case -2147483648:
                            var v16 int8 = (-(p1))
                            _ = v16
                        default:
                            var v17 uint32 = 636
                            _ = v17
                        }
                    }
                    var v18 uint = (^(v14))
                    fmt.Println(v15, v18)
                    _ = v18
                }
                _ = v15
            } else {
                v19 := []int{2147483647, 32768}
                _ = v19
            }
        case "a":
            switch ("\n\t" + "foo bar") {
            case "a":
                var v20 int8 = 42
                _ = v20
            default:
                fmt.Println(p1, p0)
            }
        case "\n\t":
            p1 = (p1 << (((v14 - 65535)) & 63))
        default:
            p0 = true
        }
        _ = v14
    } else {
        fmt.Println(p0, p0)
        p0 = true
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        for fi22 := 0; fi22 < 10; fi22++ {
            v23 := []int{(fi22 | (fi21 & fi21)), 518, (-(fi22)), (fi21 | fi21)}
            var v24 float32 = (-2.25 + (1e10 * float32((1.5 * 1e10))))
            _ = fi22
            _ = v23
            _ = v24
        }
        if p0 {
            v25 := GenS2{F0: 42, F1: 65535, F2: false}
            p1 += (^(p1))
            for fi26 := 0; fi26 < 7; fi26++ {
                var v27 float64 = 1.5
                v28 := []int{}
                v29 := GenS1{F0: 2147483647, F1: 0, F2: 2}
                _ = fi26
                _ = v27
                _ = v28
                _ = v29
            }
            _ = v25
        } else {
            v30 := GenS3{F0: 955, F1: -2147483648}
            _ = v30
        }
        p0 = p0
        _ = fi21
    }
    v31 := []int{127}
    var v32 int32 = 224
    _ = p0
    _ = p1
    _ = v31
    _ = v32
    return 288
}
func genFn1(p0 bool) int16 {
    fmt.Println(p0)
    var v33 int8 = 42
    var v34 bool = (42 < 0)
    _ = p0
    _ = v33
    _ = v34
    return 1
}

func main() {
    v1 := []string{}
    var v2 float64 = 3.14159
    var v3 int8 = 100
    v4 := make(map[string]string)
    var v5 uint16 = 1
    fmt.Println(v3)
    v3 += (-(v3))
    fmt.Println(v3)
    var v35 string = "hello"
    if false {
        var v36 float64 = 0.0
        v37 := GenS4{F0: true, F1: 127, F2: 2}
        v38 := GenS1{F0: 1, F1: -2147483648, F2: 128}
        _ = v36
        _ = v37
        _ = v38
    }
    var v39 uint64 = 942
    if (!((v3 % ((v3) | 1)) >= 1)) {
        var v40 float64 = 1.5
        var v41 int32 = 483
        fmt.Println(v2, v40)
        _ = v40
        _ = v41
    }
    for fi42 := 0; fi42 < 12; fi42++ {
        v43 := v4[v35]
        _ = fi42
        _ = v43
    }
    if true {
        var v44 int8 = 42
        _ = v44
    }
    var v45 uint64 = (v39 ^ v39)
    v5 = (v5 % ((1) | 1))
    v3 = (v3 >> ((uint(v45)) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v45
    _ = v5
}

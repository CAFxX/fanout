package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint }
type GenS2 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint64, p1 uint64, p2 int) uint16 {
    v6 := []int{p2, 842}
    if false {
        v7 := GenS1{F0: 32767, F1: 2, F2: 255}
        for fi8 := 0; fi8 < 5; fi8++ {
            switch ("hello" + "a") {
            case "hello":
                v9 := GenS1{F0: 32767, F1: 541, F2: 2}
                _ = v9
            case "foo bar":
                if (!false) {
                    if false {
                        var v10 bool = (127 == 233)
                        _ = v10
                    }
                    var v11 string = "\n\t"
                    _ = v11
                }
            case "a":
                fmt.Println(p0)
            default:
                var v12 int8 = 1
                _ = v12
            }
            _ = fi8
        }
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}
func genFn1() uint32 {
    v13 := []int{4294967295, 32767}
    _ = len(v13)
    if false {
        var v14 int = 127
        var v15 int16 = 0
        var v16 float64 = float64(1.5)
        _ = v14
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: 1000, F1: 32767, F2: 1}
    _ = v13
    _ = v17
    return 4294967295
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 146
    v3 := []uint32{65535, 4294967295, 214}
    var v4 bool = false
    v5 := make(map[string]uint32)
    switch 127 {
    case 22:
        var v18 int16 = 42
        _ = v18
    case -1:
        v19 := GenS1{F0: 1, F1: genFn0(4294967295, 229, 814), F2: 431}
        _ = v19
    default:
        var v20 uint64 = 18446744073709551615
        _ = v20
    }
    if (!false) {
        for fi21 := 0; fi21 < 17; fi21++ {
            v4 = true
            fmt.Println(v2)
            _ = len(v5)
            _ = fi21
        }
        v4 = (((548 >= 1) || v4) && (-1 < 0))
        v2 = (v2 + (v2 + (v2 % ((v2) | 1))))
    } else {
        fmt.Println(v4)
        v2 += 138
        for fi22 := 0; fi22 < 13; fi22++ {
            if len(v3) > 0 {
                v3[0] = 2
            }
            for fi23 := 0; fi23 < 7; fi23++ {
                fmt.Println(v2, fi23)
                v24 := GenS2{F0: 0, F1: 32767}
                _ = fi23
                _ = v24
            }
            if len(v3) > 0 {
                v3[0] = 478
            }
            _ = fi22
        }
    }
    v4 = false
    if true {
        if true {
            v3 = append(v3, 65535)
            fmt.Println(v2, v2)
            if v4 {
                v2 = (^(v2))
            }
        }
        if v4 {
            v2 = v2
            v2 = (v2 ^ v2)
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v2)
    delete(v1, 4)
    v1[215] = "a"
    fmt.Println(v4)
    if (!v4) {
        switch "foo bar" {
        case "foo bar":
            var v25 uint = 2
            _ = v25
        case "\n\t":
            if v4 {
                var v26 float64 = -2.25
                for fi27 := 0; fi27 < 3; fi27++ {
                    fmt.Println(v2)
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(v4)
                    }
                    _ = fi27
                }
                _ = v26
            } else {
                v4 = v4
            }
        }
    } else {
        var v28 int16 = 1000
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

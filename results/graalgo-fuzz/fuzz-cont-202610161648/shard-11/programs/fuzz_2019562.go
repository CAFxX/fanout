package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint16; F2 string }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 int8; F1 uint32 }
func genFn0() int64 {
    var v5 uint8 = 2
    var v6 uint8 = v5
    _ = v5
    _ = v6
    return 671
}
func genFn1(p0 bool, p1 uint32) uint16 {
    v7 := []int{165, 641, 127}
    p1 = 2147483647
    _ = len(v7)
    _ = p0
    _ = p1
    _ = v7
    return 32767
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[int]string)
    var v3 int32 = -914
    var v4 int16 = 32767
    switch "" {
    case "hello":
        fmt.Println(v4, v4, v3)
    case "a":
        if (3.14159 > (1e10 / 1.5)) {
            v8 := GenS1{F0: 128}
            var v9 int64 = 42
            _ = v8
            _ = v9
        }
    case "":
        if ((((0.1 / (((1.5 - 1e10) + 0.0) * -2.25)) - 1e10) * -2.25) <= -2.25) {
            if (!(0 == 32767)) {
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(v4, v3, v3)
                case "a":
                    switch 32768 {
                    case 32767:
                        fmt.Println(v3)
                    }
                default:
                    for fi10 := 0; fi10 < 6; fi10++ {
                        v4 -= v4
                        _ = fi10
                    }
                }
                if false {
                    var v11 int64 = 567
                    v11 += v11
                    _ = v11
                } else {
                    for fi12 := 0; fi12 < 16; fi12++ {
                        var v13 uint32 = 65535
                        _ = fi12
                        _ = v13
                    }
                }
            }
            var v14 uint8 = 2
            _ = v14
        }
    }
    fmt.Println(v4)
    if true {
        v15 := []int{-2147483648, 127}
        _ = v15
    }
    if ((v4 | 32767) <= v4) {
        v16 := []int{128, -2147483648}
        fmt.Println(v3, v4)
        _ = v16
    } else {
        v2[592] = "\n\t"
        v17 := []int{int(v3), 16}
        _ = v17
    }
    var v18 string = (("\n\t" + (("a" + "\n\t") + "hello")) + "foo bar")
    v19 := []int{}
    for fi20 := 0; fi20 < 9; fi20++ {
        v4 = 1
        _ = fi20
    }
    var v21 bool = false
    if true {
        for fi22 := 0; fi22 < 20; fi22++ {
            fmt.Println(v3, v21, v21)
            var v23 int64 = 823
            _ = fi22
            _ = v23
        }
        v24 := GenS1{F0: 64}
        _ = v24
    } else {
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
        var v25 int64 = 42
        _ = len(v19)
        _ = v25
    }
    for fi26 := 0; fi26 < 19; fi26++ {
        v19 = append(v19, fi26)
        v27 := GenS2{F0: -2147483648, F1: 255, F2: ((("\n\t" + ("a" + v18)) + v18) + v18)}
        for fi28 := 0; fi28 < 11; fi28++ {
            v2[(fi28 ^ fi26)] = "a"
            var v29 uint = 42
            if ((^(v1)) < (v1 * v1)) {
                v30 := v2[fi26]
                if true {
                    _ = len(v2)
                }
                _ = v30
            }
            _ = fi28
            _ = v29
        }
        _ = fi26
        _ = v27
    }
    v31 := v2[-106]
    v32 := GenS3{F0: 65535, F1: 127}
    for fi33 := 0; fi33 < 1; fi33++ {
        v2[fi33] = v31
        for fi34 := 0; fi34 < 15; fi34++ {
            v18 = (v18 + ((v31 + (v31 + v31)) + (v31 + v31)))
            switch fi34 {
            case -2147483648:
                var v35 bool = v21
                _ = v35
            case 2147483647:
                v36 := GenS1{F0: 41}
                _ = v36
            case -9223372036854775808:
                switch "foo bar" {
                case "foo bar":
                    _ = len(v19)
                case "hello":
                    v18 = "a"
                case "a":
                    var v37 int32 = v3
                    _ = v37
                }
            }
            v38 := GenS4{F0: 9, F1: 919}
            _ = fi34
            _ = v38
        }
        v39 := GenS1{F0: 2}
        _ = fi33
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v32
    _ = v4
}

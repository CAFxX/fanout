package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 int64; F2 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 int; F2 uint16 }
func genFn0(p0 int, p1 string, p2 uint32) uint64 {
    var v6 uint64 = 779
    if true {
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 uint32 = p2
            v9 := []int{fi7, fi7, fi7, fi7}
            p2 = 1
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 int16 = 156
        v11 := []int{p0, -2147483648, (p0 | p0)}
        _ = v10
        _ = v11
    } else {
        p1 = "a"
        var v12 uint64 = (v6 ^ (v6 >> ((0) & 63)))
        _ = v12
    }
    v13 := []int{p0}
    p0 -= p0
    if true {
        v14 := GenS1{F0: v6, F1: p0, F2: v6}
        var v15 uint64 = v6
        _ = v14
        _ = v15
    } else {
        v6 = 2
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v6
    return v6
}

func main() {
    var v1 int64 = 0
    var v2 int16 = 1
    var v3 int64 = -1
    var v4 int16 = -787
    var v5 uint8 = 74
    for fi16 := 0; fi16 < 6; fi16++ {
        v2 = 42
        for fi17 := 0; fi17 < 6; fi17++ {
            v18 := []int{(fi17 ^ (-(fi17)))}
            switch fi16 {
            case 42:
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
            case -363:
                switch (fi16 & 32767) {
                case 448:
                    _ = len(v18)
                }
            default:
                v19 := GenS2{F0: v3}
                _ = v19
            }
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    v4 = v4
    v1 = (v3 / ((v3) | 1))
    v20 := []int{1, -2147483648, int(v4)}
    v21 := GenS1{F0: 4294967295, F1: -367, F2: 1}
    v22 := GenS3{F0: false, F1: v1, F2: 35}
    if false {
        if false {
            v2 -= 1
        } else {
            v23 := GenS4{F0: 1}
            fmt.Println(v3, v5)
            if (2 <= (v5 | 137)) {
                v20 = append(v20, -49)
                fmt.Println(v4)
            }
            _ = v23
        }
        var v24 uint32 = 4294967295
        if true {
            v25 := GenS5{F0: (("a" + "") + "\n\t"), F1: 42, F2: 1}
            _ = v25
        }
        _ = v24
    }
    for fi26 := 0; fi26 < 11; fi26++ {
        var v27 bool = (v3 != v1)
        fmt.Println(v3)
        v5 -= v5
        _ = fi26
        _ = v27
    }
    if true {
        _ = len(v20)
        switch ((("" + (("\n\t" + "\n\t") + "")) + "") + "\n\t") {
        case "foo bar":
            v28 := GenS2{F0: v3}
            _ = v28
        case "hello":
            var v29 uint8 = 128
            _ = v29
        case "a":
            v5 = 255
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}

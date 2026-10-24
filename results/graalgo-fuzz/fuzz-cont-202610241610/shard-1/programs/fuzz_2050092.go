package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint8; F1 bool; F2 int32 }
func genFn0(p0 int16, p1 int32, p2 uint) int8 {
    v4 := []int{-756}
    switch "a" {
    case "foo bar":
        fmt.Println(p0, p0)
    case "a":
        var v5 int32 = (p1 + (p1 * (p1 & p1)))
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 127
}

func main() {
    v1 := make(map[int]string)
    var v2 float32 = 0.0
    v3 := make(map[uint32]int)
    if (941 > 1) {
        if true {
            fmt.Println(v2)
        }
        v2 = v2
        v2 += ((float32(1.5) / 3.14159) - 3.14159)
    }
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: 995, F1: 9223372036854775807}
        _ = fi6
        _ = v7
    }
    var v8 int = -2147483648
    if true {
        fmt.Println(v8)
        v1[(v8 % ((v8) | 1))] = (("" + ("hello" + "a")) + ((("\n\t" + ("hello" + "hello")) + ("a" + "")) + (("foo bar" + ("\n\t" + "foo bar")) + "")))
    }
    switch (v8 & v8) {
    case -1:
        v9 := []int{v8}
        _ = v9
    case -866:
        v10 := GenS2{F0: 42, F1: true, F2: -823}
        _ = v10
    case 652:
        for fi11 := 0; fi11 < 12; fi11++ {
            fmt.Println(fi11, v8)
            switch "hello" {
            case "foo bar":
                fmt.Println(fi11, v8, v8)
            case "":
                v12 := GenS1{F0: 686, F1: v8}
                _ = v12
            }
            if (true || true) {
                v13 := v3[465]
                switch 127 {
                case 0:
                    v14 := []int{v13}
                    _ = v14
                case 1:
                    v15 := []int{(fi11 % (((v13 | v8)) | 1))}
                    _ = v15
                }
                for fi16 := 0; fi16 < 17; fi16++ {
                    fmt.Println(v13, fi16, v8)
                    v17 := GenS2{F0: 0, F1: false, F2: -957}
                    _ = fi16
                    _ = v17
                }
                _ = v13
            }
            _ = fi11
        }
    }
    v18 := []int{32767, v8, (v8 >> ((348) & 63)), v8}
    _ = len(v18)
    if (42 < -1) {
        switch 0 {
        case 921:
            if true {
                var v19 int32 = -529
                v20 := v1[(v8 % ((v8) | 1))]
                if (!(!(65535 <= 18446744073709551615))) {
                    v21 := GenS2{F0: 128, F1: (934 >= 32767), F2: 856}
                    _ = v21
                } else {
                    var v22 bool = (0 <= 583)
                    _ = v22
                }
                _ = v19
                _ = v20
            } else {
                var v23 uint = 290
                if len(v18) > 0 {
                    v18[0] = (v8 + (-(v8)))
                }
                _ = v23
            }
        case 32767:
            var v24 float64 = 0.0
            _ = v24
        }
        fmt.Println(v2)
        var v25 uint = 65535
        _ = v25
    }
    v26 := GenS1{F0: 35, F1: v8}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v8
}

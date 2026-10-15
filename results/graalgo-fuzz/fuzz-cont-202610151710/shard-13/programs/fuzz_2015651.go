package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int; F1 int32; F2 bool }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 uint64; F1 int }
func genFn0(p0 bool) bool {
    p0 = true
    var v5 int32 = 42
    fmt.Println(v5, p0)
    var v6 uint64 = 42
    _ = p0
    _ = v5
    _ = v6
    return (v5 != 0)
}

func main() {
    v1 := []string{"\n\t", "a", "a"}
    v2 := make(map[uint32]bool)
    v3 := []int{-445}
    var v4 string = "a"
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int8 = -65
        switch (^(fi7)) {
        case 4294967295:
            switch (-(fi7)) {
            case 0:
                if false {
                    if (1000 < -74) {
                        v8 = v8
                        v8 += (v8 + (v8 >> ((466) & 63)))
                        for fi9 := 0; fi9 < 5; fi9++ {
                            var v10 int64 = 0
                            _ = fi9
                            _ = v10
                        }
                    }
                    switch fi7 {
                    case 1:
                        v4 = "a"
                    case 567:
                        var v11 int32 = 0
                        _ = v11
                    case -2147483648:
                        v4 = v4
                    }
                    if ((false && true) && false) {
                        v8 = v8
                        fmt.Println(v4)
                    }
                }
            case -231:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        case 0:
            fmt.Println(v8)
        case 2147483647:
            if (-992 < 123) {
                switch "foo bar" {
                case "":
                    v4 = v4
                default:
                    if false {
                        var v12 float32 = 1.5
                        for fi13 := 0; fi13 < 15; fi13++ {
                            var v14 float32 = v12
                            _ = fi13
                            _ = v14
                        }
                        _ = v12
                    }
                }
                var v15 float32 = 3.14159
                var v16 int8 = (v8 % ((v8) | 1))
                _ = v15
                _ = v16
            }
        default:
            var v17 int64 = -1
            _ = v17
        }
        for fi18 := 0; fi18 < 3; fi18++ {
            v8 = -29
            _ = fi18
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(v4)
    var v19 int8 = 88
    _ = len(v3)
    v4 = (v4 + "a")
    v4 = v4
    fmt.Println(v4, v19)
    v19 += (v19 * v19)
    for fi20 := 0; fi20 < 7; fi20++ {
        if ((v19 / (((^(v19))) | 1)) >= (-(v19))) {
            fmt.Println(v19)
        }
        var v21 uint16 = 42
        for fi22 := 0; fi22 < 20; fi22++ {
            v19 = v19
            v23 := GenS1{F0: 1, F1: 702}
            switch (-(fi20)) {
            case -1:
                if false {
                    if true {
                        switch "hello" {
                        case "foo bar":
                            v24 := GenS2{F0: fi20, F1: 0, F2: true}
                            _ = v24
                        case "\n\t":
                            v25 := GenS3{F0: 32767, F1: 4294967295}
                            _ = v25
                        default:
                            var v26 uint64 = 0
                            _ = v26
                        }
                        if (!false) {
                            v19 = 106
                            var v27 float32 = 3.14159
                            _ = v27
                        }
                    }
                    v28 := GenS4{F0: 42, F1: (fi22 / (((fi22 * fi22)) | 1))}
                    _ = v28
                } else {
                    v4 = ("a" + v4)
                }
            case -2147483648:
                _ = len(v2)
            default:
                v21 -= v21
            }
            _ = fi22
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    switch (v4 + v4) {
    case "a":
        var v29 bool = (42 >= 1)
        _ = v29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}

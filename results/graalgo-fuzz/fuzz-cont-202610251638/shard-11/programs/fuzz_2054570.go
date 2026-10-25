package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int64; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 uint8 }
func genFn0(p0 string) uint64 {
    p0 = p0
    if false {
        switch (p0 + p0) {
        case "a":
            fmt.Println(p0)
        }
    } else {
        p0 = (p0 + p0)
        var v7 int64 = -9223372036854775808
        if true {
            switch "foo bar" {
            case "hello":
                switch -406 {
                case -94:
                    p0 = p0
                case 1:
                    v7 = (v7 % (((^(v7))) | 1))
                case 128:
                    v8 := GenS1{F0: 0, F1: ""}
                    _ = v8
                default:
                    if (!true) {
                        v9 := GenS1{F0: 42, F1: p0}
                        for fi10 := 0; fi10 < 10; fi10++ {
                            v7 += v7
                            _ = fi10
                        }
                        p0 = p0
                        _ = v9
                    }
                }
            case "":
                if (255 != 155) {
                    v11 := GenS2{F0: v7, F1: 0, F2: -599}
                    fmt.Println(v7, v7)
                    _ = v11
                }
            case "a":
                var v12 float32 = 0.1
                _ = v12
            }
            v13 := []int{-1, 127, -1}
            fmt.Println(p0, v7)
            _ = v13
        }
        _ = v7
    }
    _ = p0
    return 2
}
func genFn1(p0 uint32, p1 uint, p2 int64) uint {
    if false {
        var v14 int64 = p2
        if (!(76 != 1)) {
            var v15 uint = p1
            p2 += (v14 << ((899) & 63))
            _ = v15
        }
        _ = v14
    } else {
        v16 := GenS1{F0: 42, F1: "hello"}
        switch "hello" {
        case "":
            p1 += p1
        case "a":
            for fi17 := 0; fi17 < 13; fi17++ {
                switch fi17 {
                case 2147483647:
                    var v18 float64 = 0.1
                    _ = v18
                default:
                    var v19 uint16 = 255
                    _ = v19
                }
                if (-128 >= 100) {
                    p2 += (p2 | (p2 | p2))
                }
                _ = fi17
            }
        case "\n\t":
            var v20 int8 = 0
            _ = v20
        }
        p2 -= -1
        _ = v16
    }
    if true {
        v21 := []int{457, 32768}
        for fi22 := 0; fi22 < 9; fi22++ {
            if false {
                v23 := GenS2{F0: 9223372036854775807, F1: 435, F2: -1}
                _ = v23
            } else {
                if len(v21) > 0 {
                    v21[0] = 384
                }
            }
            for fi24 := 0; fi24 < 4; fi24++ {
                p1 += p1
                _ = fi24
            }
            _ = fi22
        }
        for fi25 := 0; fi25 < 19; fi25++ {
            var v26 string = "\n\t"
            switch v26 {
            case "":
                switch (fi25 % (((fi25 ^ (^(fi25)))) | 1)) {
                case 32768:
                    v27 := GenS3{F0: 2}
                    _ = v27
                case -2147483648:
                    v28 := GenS3{F0: 255}
                    _ = v28
                case 128:
                    _ = len(v21)
                default:
                    p2 = -1
                }
            case "a":
                fmt.Println(p0, p1)
            }
            fmt.Println(p2, p0)
            _ = fi25
            _ = v26
        }
        _ = v21
    } else {
        switch -9223372036854775808 {
        case 800:
            var v29 int16 = -32768
            _ = v29
        }
    }
    p1 -= 42
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    v4 := make(map[uint32]bool)
    var v5 int8 = 62
    v6 := make(map[uint32]string)
    var v30 uint = 808
    var v31 uint32 = 974
    var v32 uint32 = (v31 | 2147483647)
    v33 := []int{127, 32768}
    fmt.Println(v5, v31)
    v32 -= 447
    v34 := GenS4{F0: 42, F1: 116}
    v6[v31] = ("a" + "")
    if (("\n\t" > "foo bar") || (0 < 2)) {
        v35 := GenS1{F0: -1, F1: ("\n\t" + ("foo bar" + ("" + "a")))}
        var v36 float32 = 0.0
        _ = v35
        _ = v36
    } else {
        v37 := GenS1{F0: 0, F1: "foo bar"}
        v3[2147483647] = int(v30)
        _ = len(v3)
        _ = v37
    }
    fmt.Println(v31, v32)
    for fi38 := 0; fi38 < 12; fi38++ {
        var v39 int8 = v5
        v5 = (-(v5))
        var v40 int8 = (-(v39))
        _ = fi38
        _ = v39
        _ = v40
    }
    var v41 int8 = v5
    var v42 int8 = 1
    var v43 int64 = -1
    fmt.Println(v43, v43, v41)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}

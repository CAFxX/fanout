package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 string }
type GenS2 struct { F0 uint; F1 int; F2 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 uint8; F2 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint; F1 int8 }
type GenS7 struct { F0 bool; F1 string }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]bool)
    v3 := make(map[uint32]string)
    var v4 uint16 = 0
    v5 := make(map[int]int)
    v4 -= (v4 ^ (v4 + (v4 * v4)))
    v6 := GenS1{F0: 42, F1: 368, F2: "hello"}
    if false {
        var v7 uint64 = 528
        var v8 uint32 = 432
        switch -9223372036854775808 {
        case 1:
            v9 := GenS2{F0: 0, F1: 764, F2: false}
            _ = v9
        case -9223372036854775808:
            var v10 int32 = -1
            _ = v10
        default:
            switch 127 {
            case 232:
                v11 := GenS2{F0: 42, F1: 366, F2: ((false && true) || (true || (true && true)))}
                _ = v11
            default:
                if false {
                    v12 := []int{-908, 4294967295}
                    var v13 int16 = -224
                    v4 = v4
                    _ = v12
                    _ = v13
                } else {
                    v8 += (v8 | v8)
                }
            }
        }
        _ = v7
        _ = v8
    }
    if (4294967295 < 2147483648) {
        v14 := GenS3{F0: 669}
        v15 := GenS3{F0: 1}
        fmt.Println(v4)
        _ = v14
        _ = v15
    }
    if (!true) {
        var v16 uint16 = 32767
        v4 = v4
        if true {
            v17 := GenS4{F0: -1, F1: 0, F2: (("foo bar" + "a") + "\n\t")}
            _ = v17
        }
        _ = v16
    }
    var v18 float64 = ((1e10 * 1.5) - 1e10)
    var v19 int64 = -9223372036854775808
    var v20 string = "a"
    v21 := v1[-2147483648]
    if (-103 != -1) {
        v22 := GenS5{F0: 2147483647}
        _ = v22
    } else {
        var v23 uint16 = v4
        _ = v23
    }
    v4 -= (v4 & 0)
    if true {
        for fi24 := 0; fi24 < 20; fi24++ {
            if true {
                v19 -= 56
                switch (v20 + "hello") {
                case "hello":
                    for fi25 := 0; fi25 < 13; fi25++ {
                        v26 := GenS6{F0: 2, F1: 127}
                        switch 241 {
                        case 32767:
                            v27 := GenS5{F0: 526}
                            _ = v27
                        }
                        var v28 int64 = v19
                        _ = fi25
                        _ = v26
                        _ = v28
                    }
                case "foo bar":
                    v19 = v19
                case "\n\t":
                    if false {
                        fmt.Println(fi24, v20)
                        v4 = v4
                        var v29 string = (v21 + "foo bar")
                        _ = v29
                    }
                default:
                    v30 := []int{-176, 32767, fi24, fi24}
                    _ = v30
                }
            }
            if true {
                v31 := v3[2147483647]
                _ = v31
            }
            _ = fi24
        }
        fmt.Println(v20, v19, v20)
    }
    if (!true) {
        for fi32 := 0; fi32 < 4; fi32++ {
            var v33 uint = 18446744073709551615
            if true {
                v33 -= v33
            } else {
                for fi34 := 0; fi34 < 2; fi34++ {
                    v21 = v21
                    _ = fi34
                }
                v35 := []int{}
                v36 := v1[fi32]
                _ = v35
                _ = v36
            }
            v37 := GenS5{F0: 0}
            _ = fi32
            _ = v33
            _ = v37
        }
    } else {
        for fi38 := 0; fi38 < 4; fi38++ {
            switch v20 {
            case "a":
                var v39 int = fi38
                _ = v39
            }
            _ = fi38
        }
        fmt.Println(v19)
    }
    v40 := GenS7{F0: true, F1: v20}
    for fi41 := 0; fi41 < 17; fi41++ {
        fmt.Println(v18, v18)
        fmt.Println(v4)
        var v42 uint = 391
        _ = fi41
        _ = v42
    }
    v43 := []int{}
    v19 = v19
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v40
    _ = v43
    _ = v5
    _ = v6
}

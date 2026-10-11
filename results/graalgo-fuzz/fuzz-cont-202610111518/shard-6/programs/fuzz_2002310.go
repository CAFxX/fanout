package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 uint32; F2 bool }
func genFn0(p0 string, p1 uint) int64 {
    var v6 string = ("a" + p0)
    if false {
        switch "" {
        case "a":
            var v7 int16 = 42
            _ = v7
        case "hello":
            v8 := []int{-1, 128, -2147483648}
            _ = v8
        case "foo bar":
            if false {
                for fi9 := 0; fi9 < 5; fi9++ {
                    var v10 uint16 = 32767
                    fmt.Println(p1, p1, p1)
                    _ = fi9
                    _ = v10
                }
            }
        }
    }
    var v11 string = "\n\t"
    _ = p0
    _ = p1
    _ = v11
    _ = v6
    return 1
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[int]string)
    var v3 uint8 = 27
    v4 := make(map[int]bool)
    v5 := make(map[string]int)
    var v12 float64 = v1
    v13 := GenS1{F0: 4294967295}
    fmt.Println(v12)
    v14 := v5["foo bar"]
    if (!true) {
        var v15 uint64 = 4294967295
        _ = v15
    } else {
        v16 := GenS1{F0: 717}
        var v17 uint64 = 4294967295
        _ = v16
        _ = v17
    }
    v18 := []int{(v14 % ((753) | 1)), 1, 257}
    if (true || false) {
        fmt.Println(v14, v14)
    }
    if (true || (v3 != v3)) {
        v19 := GenS1{F0: 2}
        v20 := GenS1{F0: 18446744073709551615}
        v2[(v14 ^ 32768)] = ""
        _ = v19
        _ = v20
    } else {
        delete(v4, (^(v14)))
        v21 := v5[((((("\n\t" + "") + (("\n\t" + ("\n\t" + "\n\t")) + "hello")) + ("" + "\n\t")) + "a") + ("\n\t" + "foo bar"))]
        _ = v21
    }
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    if true {
        for fi22 := 0; fi22 < 7; fi22++ {
            fmt.Println(v1)
            for fi23 := 0; fi23 < 11; fi23++ {
                switch ("foo bar" + "\n\t") {
                case "":
                    switch "hello" {
                    case "hello":
                        var v24 uint32 = 2147483647
                        _ = v24
                    case "foo bar":
                        if (uint32(v3) > 0) {
                            var v25 uint16 = 904
                            var v26 int32 = 2147483647
                            _ = v25
                            _ = v26
                        }
                    }
                default:
                    var v27 uint64 = 42
                    _ = v27
                }
                _ = fi23
            }
            v3 = 42
            _ = fi22
        }
        v3 = (^(v3))
    } else {
        v14 -= (-(v14))
    }
    fmt.Println(v3, v3, v14)
    var v28 int32 = -22
    if true {
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
        switch (("\n\t" + "foo bar") + "\n\t") {
        case "a":
            switch "hello" {
            case "foo bar":
                if (!true) {
                    v3 = (v3 & v3)
                } else {
                    if true {
                        v12 = v12
                        if (1 >= 42) {
                            var v29 float32 = 0.1
                            var v30 float32 = float32(v12)
                            var v31 uint64 = 4294967295
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                    }
                }
            case "\n\t":
                if (!false) {
                    _ = len(v18)
                    if ((v3 != v3) && false) {
                        for fi32 := 0; fi32 < 11; fi32++ {
                            if false {
                                var v33 float64 = v12
                                _ = v33
                            }
                            v3 = (v3 / ((130) | 1))
                            _ = fi32
                        }
                        v1 = v1
                    }
                } else {
                    v34 := GenS1{F0: 4294967295}
                    _ = len(v4)
                    _ = v34
                }
            }
        default:
            v35 := GenS2{F0: v14}
            _ = v35
        }
        switch -820 {
        case 32768:
            v28 -= v28
        case 2147483647:
            v14 += (-(v14))
        default:
            var v36 uint64 = 2
            _ = v36
        }
    } else {
        switch v14 {
        case -9223372036854775808:
            v37 := GenS2{F0: 128}
            _ = v37
        case 117:
            v38 := GenS3{F0: 42, F1: 0, F2: true}
            _ = v38
        }
    }
    switch ("hello" + (("hello" + "a") + "")) {
    case "hello":
        v28 = 864
    case "":
        fmt.Println(v28)
    }
    switch v14 {
    case 32768:
        var v39 uint64 = 62
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}

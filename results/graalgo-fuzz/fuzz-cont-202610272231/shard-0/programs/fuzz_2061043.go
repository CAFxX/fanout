package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 uint64; F2 uint }
func genFn0(p0 uint) int {
    v5 := GenS1{F0: 2}
    fmt.Println(p0)
    p0 = (p0 % ((p0) | 1))
    fmt.Println(p0)
    v6 := []int{128, 256, 32767}
    _ = p0
    _ = v5
    _ = v6
    return 127
}
func genFn1(p0 uint64, p1 int, p2 int32) int64 {
    if true {
        if (!(!true)) {
            fmt.Println(p1)
        }
        var v7 int32 = p2
        _ = v7
    }
    switch p1 {
    case 594:
        for fi8 := 0; fi8 < 10; fi8++ {
            switch "\n\t" {
            case "hello":
                v9 := []int{fi8, fi8, 1}
                _ = v9
            case "a":
                if false {
                    v10 := []int{p1, genFn0(2)}
                    _ = v10
                } else {
                    var v11 uint = uint(p0)
                    fmt.Println(p0, v11, v11)
                    _ = v11
                }
            }
            v12 := GenS2{F0: -1}
            _ = fi8
            _ = v12
        }
    case 1:
        v13 := []int{(p1 << ((0) & 63)), p1, -9223372036854775808}
        _ = v13
    case 0:
        v14 := []int{(p1 / (((p1 % ((p1) | 1))) | 1)), (p1 % ((-9223372036854775808) | 1))}
        _ = v14
    default:
        fmt.Println(p2)
    }
    switch p1 {
    case 32767:
        for fi15 := 0; fi15 < 7; fi15++ {
            fmt.Println(fi15, p1)
            if (!false) {
                p1 = (p1 + fi15)
                v16 := GenS3{F0: fi15}
                v17 := GenS3{F0: (^(p1))}
                _ = v16
                _ = v17
            }
            v18 := GenS3{F0: (^(fi15))}
            _ = fi15
            _ = v18
        }
    case 9223372036854775807:
        p0 += p0
    default:
        v19 := GenS1{F0: 2}
        _ = v19
    }
    var v20 uint16 = 32767
    if false {
        v21 := GenS1{F0: 4294967295}
        v22 := []int{p1, -9223372036854775808, genFn0(65535)}
        var v23 uint32 = 4294967295
        _ = v21
        _ = v22
        _ = v23
    } else {
        if false {
            var v24 float32 = 0.1
            fmt.Println(p1, p1)
            var v25 string = (("" + "foo bar") + ("\n\t" + (("a" + "") + "\n\t")))
            _ = v24
            _ = v25
        }
        fmt.Println(p0, v20, p0)
        for fi26 := 0; fi26 < 12; fi26++ {
            if (true && false) {
                switch (p1 ^ p1) {
                case 4294967295:
                    var v27 uint8 = 0
                    _ = v27
                default:
                    switch "a" {
                    case "a":
                        v28 := GenS1{F0: 2}
                        _ = v28
                    case "":
                        v29 := GenS2{F0: -32768}
                        _ = v29
                    case "\n\t":
                        switch (fi26 / ((fi26) | 1)) {
                        case -908:
                            if ((p1 & 0) >= 32768) {
                                fmt.Println(p2, p0, v20)
                                p1 = fi26
                            } else {
                                var v30 int8 = 56
                                _ = v30
                            }
                        case -9223372036854775808:
                            fmt.Println(p0)
                        }
                    }
                }
                var v31 uint64 = p0
                var v32 float32 = -2.25
                _ = v31
                _ = v32
            }
            p1 = (fi26 % (((p1 + fi26)) | 1))
            _ = fi26
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return 0
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[string]bool)
    v3 := make(map[string]string)
    v4 := make(map[string]uint32)
    var v33 int8 = 53
    var v34 int32 = -2147483648
    v35 := []int{127}
    if true {
        _ = len(v35)
    }
    var v36 bool = (!true)
    v37 := GenS1{F0: 2147483648}
    switch (v1 + "hello") {
    case "\n\t":
        switch v1 {
        case "\n\t":
            v36 = v36
        case "":
            fmt.Println(v34, v36)
        case "a":
            if true {
                v33 = v33
                v38 := GenS2{F0: 32767}
                v36 = v36
                _ = v38
            }
        default:
            v1 = (v1 + v1)
        }
    case "a":
        fmt.Println(v1, v36, v33)
    case "":
        v33 += 100
    default:
        _ = len(v3)
    }
    var v39 string = "foo bar"
    for fi40 := 0; fi40 < 14; fi40++ {
        var v41 bool = ((fi40 ^ -825) >= fi40)
        var v42 uint32 = 42
        _ = fi40
        _ = v41
        _ = v42
    }
    fmt.Println(v1)
    v1 = "hello"
    v43 := GenS4{F0: v33, F1: 1, F2: 255}
    for fi44 := 0; fi44 < 7; fi44++ {
        for fi45 := 0; fi45 < 12; fi45++ {
            switch (("hello" + v39) + "foo bar") {
            case "\n\t":
                fmt.Println(v1, v1)
            }
            _ = fi45
        }
        var v46 int8 = v33
        delete(v3, v39)
        _ = fi44
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v43
}

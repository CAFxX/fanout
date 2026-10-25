package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int; F1 int; F2 uint8 }

func main() {
    var v1 uint16 = 1
    var v2 int = -1
    v3 := make(map[int]string)
    v4 := make(map[string]string)
    v5 := make(map[int]int)
    var v6 uint32 = 0
    if (0 == -32768) {
        for fi7 := 0; fi7 < 4; fi7++ {
            switch "a" {
            case "a":
                if false {
                    fmt.Println(v6, v1)
                }
            case "foo bar":
                if false {
                    delete(v5, 9223372036854775807)
                    fmt.Println(fi7, v2)
                }
            case "":
                for fi8 := 0; fi8 < 11; fi8++ {
                    v6 = 2147483648
                    switch 4294967295 {
                    case -9223372036854775808:
                        fmt.Println(v1, fi7, v2)
                    }
                    fmt.Println(v2, v6, v1)
                    _ = fi8
                }
            }
            var v9 uint64 = 1
            _ = fi7
            _ = v9
        }
        switch 227 {
        case -1:
            var v10 uint16 = v1
            _ = v10
        case 4294967295:
            var v11 int32 = 1
            _ = v11
        default:
            var v12 uint32 = 4294967295
            _ = v12
        }
        if false {
            v13 := []int{0}
            _ = v13
        }
    }
    if true {
        v1 = (^(v1))
    } else {
        switch (v2 - v2) {
        case 9223372036854775807:
            if false {
                v14 := GenS1{F0: 127}
                switch v2 {
                case 336:
                    for fi15 := 0; fi15 < 7; fi15++ {
                        fmt.Println(v1, v2, v1)
                        var v16 uint64 = uint64(v1)
                        fmt.Println(v6)
                        _ = fi15
                        _ = v16
                    }
                case -9223372036854775808:
                    var v17 int16 = 1
                    _ = v17
                }
                _ = v14
            } else {
                var v18 int64 = 9223372036854775807
                _ = v18
            }
        default:
            var v19 uint32 = (v6 + (v6 + 767))
            _ = v19
        }
        v3[42] = "a"
        v1 = 42
    }
    switch -2147483648 {
    case 128:
        v20 := GenS2{F0: -208, F1: -128}
        _ = v20
    default:
        if (9223372036854775807 != (v2 / ((v2) | 1))) {
            fmt.Println(v1, v1)
        }
    }
    fmt.Println(v2, v2, v2)
    var v21 int8 = -74
    switch "hello" {
    case "\n\t":
        v3[(-(v2))] = ""
    case "a":
        v5[v2] = -1
    case "hello":
        var v22 bool = false
        _ = v22
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        v24 := []int{fi23, 4294967295, (v2 & 139), fi23}
        v1 = (v1 ^ 42)
        _ = fi23
        _ = v24
    }
    var v25 int16 = -32768
    v26 := GenS3{F0: -9223372036854775808, F1: v2, F2: 42}
    for fi27 := 0; fi27 < 8; fi27++ {
        if false {
            if (false || true) {
                var v28 uint = 726
                var v29 int64 = 42
                _ = v28
                _ = v29
            } else {
                v25 += (v25 + (v25 % (((v25 | v25)) | 1)))
                if false {
                    var v30 int = int(v1)
                    _ = v30
                }
            }
            for fi31 := 0; fi31 < 11; fi31++ {
                switch (fi31 << ((4294967295) & 63)) {
                case 0:
                    v25 += 0
                case 1:
                    var v32 int = v2
                    _ = v32
                case 42:
                    if (-32768 > v25) {
                        v25 -= v25
                        v33 := []int{(v2 / (((fi31 >> ((2) & 63))) | 1)), fi27, (-(v2)), (v2 | 572)}
                        _ = v33
                    } else {
                        var v34 int8 = (v21 & v21)
                        v6 = 2147483647
                        var v35 uint8 = 156
                        var v36 uint64 = 1
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                default:
                    switch ("a" + ("" + (("foo bar" + "hello") + "hello"))) {
                    case "":
                        fmt.Println(fi31, v1, v25)
                    case "hello":
                        if true {
                            v37 := GenS1{F0: 127}
                            _ = v37
                        }
                    }
                }
                _ = fi31
            }
            var v38 bool = (3.14159 != float32(0.0))
            _ = v38
        }
        for fi39 := 0; fi39 < 1; fi39++ {
            v21 = -90
            switch (^(fi39)) {
            case 332:
                v6 += v6
            case 180:
                var v40 int32 = 1
                _ = v40
            }
            var v41 int64 = 612
            _ = fi39
            _ = v41
        }
        var v42 int32 = 0
        _ = fi27
        _ = v42
    }
    var v43 float64 = -2.25
    v44 := []int{-355, v2, v2}
    v44 = append(v44, (v2 % ((v2) | 1)))
    var v45 uint16 = (v1 + (v1 | (v1 + v1)))
    fmt.Println(v25, v21, v21)
    var v46 int64 = 183
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}

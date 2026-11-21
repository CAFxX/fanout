package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 string; F2 uint16 }
func genFn0() uint8 {
    v7 := GenS1{F0: 9223372036854775807, F1: 0, F2: -1}
    if (108 > 67) {
        var v8 int = -1
        v8 += 2147483647
        if (-2.25 < 1.5) {
            v8 = v8
            v9 := []int{v8, -427, v8}
            _ = v9
        }
        _ = v8
    } else {
        switch -621 {
        case 32768:
            var v10 uint32 = 431
            _ = v10
        case 1:
            var v11 int = 42
            _ = v11
        default:
            switch -129 {
            case -2147483648:
                v12 := GenS1{F0: 9223372036854775807, F1: -466, F2: 755}
                _ = v12
            default:
                var v13 uint = 2
                _ = v13
            }
        }
        var v14 uint64 = 42
        var v15 string = ""
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: 0, F1: -53, F2: 32768}
    _ = v16
    _ = v7
    return 42
}
func genFn1() uint64 {
    v17 := GenS1{F0: -298, F1: -2147483648, F2: -9223372036854775808}
    for fi18 := 0; fi18 < 20; fi18++ {
        if (true && true) {
            v19 := []int{(fi18 | (fi18 * fi18))}
            var v20 uint32 = 258
            _ = v19
            _ = v20
        } else {
            if (!false) {
                fmt.Println(fi18)
                var v21 uint64 = 18446744073709551615
                _ = v21
            }
            v22 := []int{127, fi18}
            if true {
                fmt.Println(fi18)
            }
            _ = v22
        }
        _ = fi18
    }
    var v23 int16 = -548
    _ = v17
    _ = v23
    return 42
}

func main() {
    v1 := []int{32767, 51, 2147483647, 520}
    v2 := make(map[string]int)
    v3 := make(map[int]bool)
    var v4 string = "hello"
    v5 := make(map[int]uint32)
    v6 := []int{943, 42}
    v4 = v4
    for fi24 := 0; fi24 < 15; fi24++ {
        for fi25 := 0; fi25 < 6; fi25++ {
            fmt.Println(fi24, fi24)
            v4 = v4
            v4 = ""
            _ = fi25
        }
        _ = fi24
    }
    if false {
        var v26 uint32 = 2147483648
        _ = v26
    } else {
        var v27 int64 = 60
        switch 4294967295 {
        case 2147483647:
            v28 := GenS2{F0: ((v4 + (v4 + v4)) + v4)}
            _ = v28
        case 809:
            switch "hello" {
            case "":
                v29 := GenS2{F0: v4}
                _ = v29
            case "hello":
                delete(v5, -9223372036854775808)
            }
        }
        _ = v27
    }
    v30 := GenS1{F0: 1, F1: 29, F2: -9223372036854775808}
    v4 = v4
    var v31 uint32 = 1
    v32 := GenS3{F0: 42, F1: ((v4 + "\n\t") + "\n\t"), F2: 42}
    v31 -= v31
    v4 = (v4 + "hello")
    var v33 uint = 255
    var v34 uint8 = 255
    switch v4 {
    case "a":
        fmt.Println(v33, v34)
    case "\n\t":
        var v35 int64 = 9223372036854775807
        _ = v35
    case "foo bar":
        if false {
            if false {
                var v36 uint32 = 54
                v37 := GenS3{F0: 42, F1: ("foo bar" + v4), F2: 1}
                if false {
                    v38 := GenS1{F0: -9223372036854775808, F1: -1, F2: -555}
                    v39 := GenS2{F0: "a"}
                    fmt.Println(v4)
                    _ = v38
                    _ = v39
                }
                _ = v36
                _ = v37
            } else {
                var v40 int64 = 53
                _ = len(v2)
                _ = v40
            }
            for fi41 := 0; fi41 < 8; fi41++ {
                var v42 string = "\n\t"
                _ = fi41
                _ = v42
            }
            var v43 float64 = (0.0 / 3.14159)
            _ = v43
        } else {
            v4 = v4
            var v44 uint = v33
            var v45 int = 32768
            _ = v44
            _ = v45
        }
    }
    var v46 bool = (!(0.1 == 0.0))
    v31 = (v31 | (v31 - v31))
    var v47 uint8 = v34
    v34 += 125
    fmt.Println(v34, v47, v34)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}

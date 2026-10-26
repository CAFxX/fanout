package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int64; F1 uint8; F2 int }
type GenS3 struct { F0 int16; F1 int64; F2 int }
func genFn0(p0 uint64, p1 int16, p2 int8) string {
    switch "" {
    case "hello":
        v4 := GenS1{F0: 8, F1: 42}
        _ = v4
    case "\n\t":
        v5 := []int{}
        _ = v5
    case "":
        fmt.Println(p1, p1)
    }
    var v6 int16 = p1
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        v7 := GenS1{F0: 0, F1: 2}
        _ = v7
    case "\n\t":
        p2 = p2
    case "a":
        v8 := GenS1{F0: 478, F1: 4294967295}
        _ = v8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        fmt.Println(p0, p0, p1)
        _ = fi9
    }
    var v10 uint64 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    return "foo bar"
}
func genFn1(p0 int8, p1 string, p2 int8) uint64 {
    var v11 int64 = -535
    if false {
        p2 = p0
    }
    p2 += 0
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 2
}

func main() {
    var v1 int8 = 75
    var v2 uint64 = 0
    v3 := make(map[int]string)
    v12 := GenS1{F0: 762, F1: 65535}
    for fi13 := 0; fi13 < 1; fi13++ {
        v2 = (v2 << ((1) & 63))
        v14 := []int{42, fi13, fi13}
        _ = fi13
        _ = v14
    }
    var v15 int16 = -32768
    v16 := v3[int(v15)]
    switch (v16 + v16) {
    case "foo bar":
        var v17 int8 = 0
        _ = v17
    case "":
        if true {
            var v18 string = v16
            _ = v18
        }
    default:
        if (40 == 65535) {
            var v19 uint64 = v2
            _ = v19
        }
    }
    var v20 string = v16
    for fi21 := 0; fi21 < 16; fi21++ {
        v15 = (^(v15))
        fmt.Println(v1)
        v22 := GenS1{F0: v15, F1: 65535}
        _ = fi21
        _ = v22
    }
    switch v16 {
    case "a":
        v23 := GenS1{F0: v15, F1: 1}
        _ = v23
    case "\n\t":
        if false {
            v2 -= (v2 - (v2 / (((v2 & v2)) | 1)))
            fmt.Println(v2, v1)
            var v24 int8 = (-(v1))
            _ = v24
        }
    }
    if (280 > 255) {
        v25 := GenS1{F0: (v15 * v15), F1: 0}
        var v26 float64 = (1e10 + 1e10)
        _ = v25
        _ = v26
    } else {
        v15 = (v15 | -463)
        v27 := GenS1{F0: v15, F1: 2}
        for fi28 := 0; fi28 < 10; fi28++ {
            for fi29 := 0; fi29 < 2; fi29++ {
                v30 := GenS1{F0: v15, F1: 160}
                v20 = "\n\t"
                var v31 int8 = v1
                v1 += (v31 ^ 0)
                _ = fi29
                _ = v30
                _ = v31
            }
            switch (fi28 | fi28) {
            case 1:
                if (true || true) {
                    switch v20 {
                    case "foo bar":
                        v32 := GenS2{F0: 9223372036854775807, F1: 26, F2: (fi28 / ((int(v15)) | 1))}
                        _ = v32
                    default:
                        if (!true) {
                            var v33 int = 1
                            switch -924 {
                            case 4294967295:
                                v34 := v3[(v33 & v33)]
                                _ = v34
                            default:
                                fmt.Println(v2, v16, v1)
                            }
                            _ = v33
                        }
                    }
                    var v35 string = ""
                    var v36 string = ("a" + ((v16 + "\n\t") + v35))
                    _ = v35
                    _ = v36
                }
            default:
                var v37 uint8 = 172
                _ = v37
            }
            _ = fi28
        }
        _ = v27
    }
    v16 = (v20 + v20)
    for fi38 := 0; fi38 < 3; fi38++ {
        for fi39 := 0; fi39 < 2; fi39++ {
            var v40 bool = true
            var v41 string = v20
            v40 = true
            _ = fi39
            _ = v40
            _ = v41
        }
        v42 := GenS3{F0: v15, F1: 776, F2: (fi38 ^ (fi38 / (((fi38 - (fi38 | fi38))) | 1)))}
        var v43 uint8 = 2
        _ = fi38
        _ = v42
        _ = v43
    }
    v44 := GenS2{F0: -9223372036854775808, F1: 127, F2: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v44
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 int64; F2 int32 }
func genFn0(p0 string, p1 uint8) uint32 {
    switch "" {
    case "":
        for fi6 := 0; fi6 < 19; fi6++ {
            for fi7 := 0; fi7 < 15; fi7++ {
                v8 := GenS1{F0: -1, F1: 9223372036854775807}
                fmt.Println(p0, fi6)
                var v9 uint32 = 358
                _ = fi7
                _ = v8
                _ = v9
            }
            _ = fi6
        }
    case "hello":
        v10 := []int{2147483647}
        _ = v10
    default:
        switch p0 {
        case "hello":
            for fi11 := 0; fi11 < 1; fi11++ {
                if (false || true) {
                    for fi12 := 0; fi12 < 2; fi12++ {
                        switch "foo bar" {
                        case "foo bar":
                            var v13 int8 = 113
                            _ = v13
                        case "a":
                            p1 = p1
                        default:
                            fmt.Println(fi11, p1)
                        }
                        _ = fi12
                    }
                    switch (fi11 - -2147483648) {
                    case -9223372036854775808:
                        v14 := []int{}
                        _ = v14
                    case -898:
                        p0 = (p0 + "foo bar")
                    default:
                        var v15 uint64 = 2
                        _ = v15
                    }
                } else {
                    v16 := []int{(fi11 % ((817) | 1))}
                    for fi17 := 0; fi17 < 12; fi17++ {
                        if false {
                            switch p0 {
                            case "foo bar":
                                v18 := GenS1{F0: -260, F1: -9223372036854775808}
                                v16 = append(v16, 482)
                                _ = v18
                            case "hello":
                                var v19 int8 = -2
                                _ = v19
                            case "a":
                                if len(v16) > 0 {
                                    v16[0] = 32767
                                }
                            }
                            var v20 float32 = -2.25
                            _ = v20
                        }
                        var v21 uint64 = 42
                        p1 = p1
                        v22 := GenS1{F0: -1, F1: 1}
                        var v23 int = (fi11 >> ((18446744073709551615) & 63))
                        _ = fi17
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    if ((0 > 0) || false) {
                        var v24 int64 = -1
                        _ = v24
                    } else {
                        p1 -= 1
                        fmt.Println(fi11)
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                    }
                    _ = v16
                }
                var v25 int8 = 1
                var v26 int64 = 0
                _ = fi11
                _ = v25
                _ = v26
            }
        case "foo bar":
            v27 := GenS1{F0: -1, F1: -1}
            _ = v27
        case "a":
            v28 := []int{}
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 10; fi29++ {
        var v30 uint16 = 293
        switch ((p0 + p0) + ("foo bar" + p0)) {
        case "\n\t":
            v31 := GenS1{F0: 349, F1: -9223372036854775808}
            _ = v31
        default:
            var v32 uint16 = (^(v30))
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    var v33 float32 = float32(0.0)
    _ = p0
    _ = p1
    _ = v33
    return 42
}

func main() {
    var v1 int64 = 406
    v2 := make(map[uint32]string)
    var v3 uint = 225
    var v4 string = ""
    v5 := []int{2147483647}
    if true {
        var v34 uint8 = 42
        _ = v34
    }
    switch (v4 + v4) {
    case "hello":
        var v35 uint64 = 1
        _ = v35
    case "foo bar":
        v36 := GenS2{F0: 609}
        _ = v36
    case "a":
        if (100 >= -1) {
            for fi37 := 0; fi37 < 12; fi37++ {
                for fi38 := 0; fi38 < 18; fi38++ {
                    v1 -= v1
                    if false {
                        v4 = (v4 + "a")
                        switch v4 {
                        case "":
                            var v39 float64 = 1.5
                            _ = v39
                        case "\n\t":
                            v40 := GenS3{F0: 1000, F1: v1, F2: 241}
                            _ = v40
                        }
                        fmt.Println(v1, fi37, fi38)
                    } else {
                        v1 -= (v1 << (((v3 + v3)) & 63))
                        var v41 uint16 = 32767
                        _ = v41
                    }
                    _ = fi38
                }
                _ = fi37
            }
        }
    default:
        v1 += v1
    }
    var v42 uint32 = 65535
    delete(v2, v42)
    _ = len(v2)
    var v43 uint = (^(v3))
    for fi44 := 0; fi44 < 12; fi44++ {
        v42 += v42
        v42 = (v42 * (v42 + 2))
        v45 := GenS2{F0: (v42 | 4294967295)}
        _ = fi44
        _ = v45
    }
    switch (v4 + (v4 + "\n\t")) {
    case "\n\t":
        v46 := GenS1{F0: 1, F1: (v1 + (v1 & -82))}
        _ = v46
    }
    var v47 int = 127
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v47
    _ = v5
}

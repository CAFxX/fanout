package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 uint8; F2 bool }
type GenS3 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS4 struct { F0 bool }

func main() {
    var v1 bool = true
    var v2 uint64 = 1
    v3 := make(map[string]string)
    var v4 int32 = -1
    var v5 int8 = -1
    var v6 uint32 = 4294967295
    if true {
        if v1 {
            _ = len(v3)
            for fi7 := 0; fi7 < 17; fi7++ {
                v4 = v4
                _ = fi7
            }
            if (!false) {
                var v8 uint8 = 1
                _ = v8
            }
        } else {
            var v9 int = 32767
            fmt.Println(v9, v4, v5)
            var v10 uint64 = (v2 << ((228) & 63))
            _ = v10
            _ = v9
        }
        for fi11 := 0; fi11 < 10; fi11++ {
            v12 := []int{fi11, 4294967295}
            var v13 uint32 = (v6 % ((v6) | 1))
            var v14 int64 = 553
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
    } else {
        for fi15 := 0; fi15 < 3; fi15++ {
            if v1 {
                var v16 uint32 = v6
                if v1 {
                    v17 := GenS1{F0: 32767}
                    v5 += v5
                    _ = v17
                } else {
                    for fi18 := 0; fi18 < 18; fi18++ {
                        var v19 int32 = (v4 >> ((65535) & 63))
                        var v20 int32 = (^(v4))
                        var v21 uint = 0
                        delete(v3, ("\n\t" + ("hello" + ("\n\t" + "\n\t"))))
                        _ = fi18
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    if ((v1 || v1) || v1) {
                        v16 -= (^(v6))
                    }
                }
                _ = v16
            } else {
                var v22 uint8 = 0
                var v23 uint64 = v2
                var v24 bool = v1
                _ = v22
                _ = v23
                _ = v24
            }
            _ = fi15
        }
    }
    v25 := v3["hello"]
    var v26 uint = 65535
    delete(v3, "")
    v27 := []int{-1, -851}
    v5 = v5
    v26 += v26
    switch v25 {
    case "hello":
        for fi28 := 0; fi28 < 11; fi28++ {
            for fi29 := 0; fi29 < 20; fi29++ {
                switch (("a" + v25) + ((v25 + v25) + v25)) {
                case "foo bar":
                    if v1 {
                        v30 := GenS2{F0: 279, F1: 42, F2: true}
                        _ = v30
                    } else {
                        v31 := GenS3{F0: 2, F1: 1, F2: 1}
                        switch "" {
                        case "\n\t":
                            fmt.Println(v5)
                        default:
                            v25 = (v25 + "a")
                        }
                        switch 1 {
                        case 4294967295:
                            switch v25 {
                            case "\n\t":
                                _ = len(v27)
                            case "foo bar":
                                if false {
                                    var v32 float64 = (float64(3.14159) + -2.25)
                                    v1 = false
                                    if ((-(v4)) <= v4) {
                                        v5 = 100
                                        switch fi28 {
                                        case 42:
                                            var v33 uint64 = (v2 & (v2 | v2))
                                            _ = v33
                                        case -1:
                                            if (v6 < (v6 * 2147483647)) {
                                                v27 = append(v27, (fi28 + (-(fi29))))
                                            }
                                        case 9223372036854775807:
                                            v3[v25] = (v25 + v25)
                                        }
                                        switch ("foo bar" + v25) {
                                        case "\n\t":
                                            switch "\n\t" {
                                            case "a":
                                                _ = len(v27)
                                            case "":
                                                var v34 uint16 = 492
                                                _ = v34
                                            case "foo bar":
                                                switch ((v25 + ("" + "\n\t")) + v25) {
                                                case "a":
                                                    var v35 uint32 = 855
                                                    _ = v35
                                                }
                                            default:
                                                var v36 int = (fi29 / ((fi28) | 1))
                                                _ = v36
                                            }
                                        default:
                                            v6 += v6
                                        }
                                        v26 = (v26 & 248)
                                    } else {
                                        v37 := GenS1{F0: 32767}
                                        _ = len(v27)
                                        _ = v37
                                    }
                                    _ = v32
                                }
                            case "":
                                _ = len(v3)
                            }
                        case 536:
                            fmt.Println(fi29, v6)
                        case 127:
                            _ = len(v3)
                        }
                        var v38 uint64 = uint64(v6)
                        _ = v31
                        _ = v38
                    }
                    v26 = 65535
                    switch "foo bar" {
                    case "a":
                        v25 = v25
                    case "\n\t":
                        fmt.Println(v1, fi29, v5)
                    default:
                        v39 := GenS4{F0: v1}
                        _ = v39
                    }
                }
                _ = fi29
            }
            v40 := v3[v25]
            fmt.Println(v2, v26)
            _ = fi28
            _ = v40
        }
    case "":
        switch 0 {
        case 0:
            if ((^(v4)) != v4) {
                fmt.Println(v26, v4, v6)
            }
        case 128:
            if (v26 != 1) {
                var v41 uint = v26
                _ = v41
            }
        }
    default:
        var v42 uint = v26
        _ = v42
    }
    switch -1 {
    case 4294967295:
        v25 = ("" + v25)
    case 0:
        v43 := v3[((((v25 + v25) + (v25 + v25)) + v25) + v25)]
        _ = v43
    default:
        for fi44 := 0; fi44 < 18; fi44++ {
            fmt.Println(v2, v25, v26)
            v6 = 42
            _ = fi44
        }
    }
    switch v25 {
    case "foo bar":
        if v1 {
            switch 995 {
            case -909:
                v27 = append(v27, 142)
            default:
                v6 -= (v6 - uint32(v26))
            }
            var v45 int64 = 42
            _ = v45
        } else {
            v1 = (v4 == (v4 << ((v26) & 63)))
            var v46 uint = v26
            _ = v46
        }
    case "hello":
        fmt.Println(v1, v5)
    }
    var v47 uint32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}

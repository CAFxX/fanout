package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool }

func main() {
    var v1 uint64 = 1
    var v2 float32 = 1e10
    var v3 string = "hello"
    v4 := make(map[string]bool)
    var v5 int = 128
    v5 -= 127
    v5 = v5
    v2 += (float32(1.5) * 1.5)
    switch "a" {
    case "":
        if true {
            fmt.Println(v5)
            v6 := GenS1{F0: 2}
            _ = v6
        }
    case "hello":
        var v7 int8 = -54
        _ = v7
    case "\n\t":
        var v8 string = "hello"
        _ = v8
    default:
        v1 = 18446744073709551615
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        for fi10 := 0; fi10 < 7; fi10++ {
            switch (fi10 | (fi10 & fi9)) {
            case 544:
                switch v3 {
                case "a":
                    delete(v4, v3)
                default:
                    v2 += 3.14159
                }
            case 1:
                fmt.Println(v3, v1)
            case -765:
                switch v3 {
                case "":
                    var v11 uint32 = 669
                    _ = v11
                case "foo bar":
                    switch (fi10 << ((uint(v1)) & 63)) {
                    case -97:
                        var v12 uint = 0
                        _ = v12
                    case 704:
                        var v13 uint32 = 526
                        _ = v13
                    case 446:
                        v14 := GenS1{F0: 2}
                        v4[("" + ((v3 + "foo bar") + "hello"))] = (false || true)
                        var v15 float64 = (0.0 + ((0.1 / float64(0.1)) * 3.14159))
                        v3 = "\n\t"
                        if (879 == 42) {
                            fmt.Println(fi9, v15, fi10)
                            var v16 int32 = -841
                            var v17 float32 = v2
                            _ = v16
                            _ = v17
                        } else {
                            var v18 uint16 = 430
                            _ = v18
                        }
                        if false {
                            v19 := GenS2{F0: -101}
                            v15 -= 0.0
                            v15 += 1.5
                            _ = v19
                        }
                        _ = v14
                        _ = v15
                    default:
                        var v20 string = v3
                        _ = v20
                    }
                default:
                    fmt.Println(v5, fi10, v3)
                }
            default:
                var v21 bool = false
                _ = v21
            }
            var v22 int64 = -1
            _ = fi10
            _ = v22
        }
        _ = fi9
    }
    var v23 float32 = (v2 + 1e10)
    v24 := v4["\n\t"]
    v5 = (v5 | v5)
    v25 := []int{(v5 - 0), (v5 >> ((42) & 63)), 32768}
    if (!((42 <= 2) && true)) {
        if ((v23 + float32(float64(-2.25))) == v2) {
            var v26 uint64 = 927
            delete(v4, (v3 + "foo bar"))
            _ = v26
        }
        v4[("hello" + "foo bar")] = false
    } else {
        v2 = (((v2 * 1.5) * float32(float64(-2.25))) * v2)
        var v27 bool = (!(255 <= 2))
        switch v3 {
        case "a":
            if (!(v1 < 0)) {
                v5 -= (v5 ^ v5)
                v27 = (!(-1 == 0))
                if true {
                    v5 = 252
                    v1 += (v1 - v1)
                    for fi28 := 0; fi28 < 3; fi28++ {
                        if ((86 <= 237) || false) {
                            if true {
                                var v29 int16 = 93
                                _ = v29
                            }
                        }
                        _ = fi28
                    }
                } else {
                    var v30 int64 = -9223372036854775808
                    _ = v30
                }
            } else {
                v27 = (v1 < (v1 | v1))
                if len(v25) > 0 {
                    v25[0] = v5
                }
            }
        case "":
            v24 = false
        case "foo bar":
            v1 = (v1 & (v1 * v1))
        }
        _ = v27
    }
    fmt.Println(v5, v2, v1)
    var v31 float32 = v2
    for fi32 := 0; fi32 < 2; fi32++ {
        for fi33 := 0; fi33 < 5; fi33++ {
            if v24 {
                var v34 uint64 = 18446744073709551615
                if v24 {
                    if len(v25) >= 2 {
                        v25 = v25[1:len(v25)]
                    }
                    v31 = 0.1
                }
                _ = len(v25)
                _ = v34
            }
            switch v3 {
            case "":
                if true {
                    v23 = (0.1 / v2)
                } else {
                    v35 := GenS1{F0: 2}
                    _ = v35
                }
            default:
                v24 = v24
            }
            _ = fi33
        }
        var v36 uint64 = v1
        _ = fi32
        _ = v36
    }
    v37 := GenS3{F0: 854}
    if true {
        fmt.Println(v1, v3, v3)
        var v38 uint64 = (^(v1))
        v31 += v31
        _ = v38
    } else {
        v24 = v24
        v39 := GenS1{F0: 110}
        var v40 uint8 = 0
        _ = v39
        _ = v40
    }
    if (554 != 437) {
        v41 := GenS4{F0: v24}
        v5 = int(v1)
        _ = v41
    } else {
        fmt.Println(v31)
        v42 := GenS1{F0: 2}
        _ = v42
    }
    var v43 bool = (2147483648 != 2147483647)
    if len(v25) >= 2 {
        v25 = v25[1:len(v25)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v37
    _ = v4
    _ = v43
    _ = v5
}

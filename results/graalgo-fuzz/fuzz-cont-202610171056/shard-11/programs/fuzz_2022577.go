package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 int64; F2 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = -2147483648
    var v3 int32 = -48
    var v4 uint = 781
    var v5 string = "\n\t"
    switch -2147483648 {
    case 32768:
        for fi6 := 0; fi6 < 8; fi6++ {
            v4 = (v4 - v4)
            _ = fi6
        }
    case 4294967295:
        if (false || (42 >= 4294967295)) {
            v4 -= (v4 - (v4 + 18446744073709551615))
        }
    case -1:
        v7 := []int{63, -757}
        _ = v7
    }
    switch v5 {
    case "\n\t":
        v4 = v4
    case "foo bar":
        var v8 int64 = 42
        _ = v8
    case "":
        for fi9 := 0; fi9 < 8; fi9++ {
            switch v5 {
            case "\n\t":
                fmt.Println(fi9)
            case "":
                v4 = v4
            case "a":
                var v10 float32 = -2.25
                _ = v10
            }
            if (((0.0 + 1.5) + (float64(0.1) + (1e10 - (float64(-2.25) * float64((float32((3.14159 - 1e10)) - 3.14159)))))) == (float64(0.1) * 1.5)) {
                for fi11 := 0; fi11 < 20; fi11++ {
                    if ((1 != 0) && false) {
                        if true {
                            var v12 int32 = (v3 - v2)
                            v5 = "foo bar"
                            switch v5 {
                            case "foo bar":
                                v13 := GenS1{F0: v5}
                                _ = v13
                            case "hello":
                                _ = len(v1)
                            }
                            _ = v12
                        }
                    } else {
                        v14 := []int{fi9, (fi9 % (((fi11 - 32767)) | 1)), fi9, fi9}
                        _ = v14
                    }
                    _ = fi11
                }
                var v15 bool = true
                _ = v15
            } else {
                v16 := GenS1{F0: v5}
                for fi17 := 0; fi17 < 11; fi17++ {
                    v4 = (v4 * v4)
                    v18 := GenS2{F0: 4294967295}
                    v19 := []int{}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = v16
            }
            for fi20 := 0; fi20 < 15; fi20++ {
                v3 = 2147483647
                _ = fi20
            }
            _ = fi9
        }
    default:
        for fi21 := 0; fi21 < 7; fi21++ {
            v22 := GenS2{F0: 702}
            _ = fi21
            _ = v22
        }
    }
    v23 := []int{4294967295, -2147483648, -2147483648}
    fmt.Println(v3, v5, v5)
    v24 := GenS2{F0: 982}
    var v25 bool = true
    for fi26 := 0; fi26 < 18; fi26++ {
        var v27 float32 = 1.5
        _ = fi26
        _ = v27
    }
    fmt.Println(v2, v2)
    if len(v23) > 0 {
        v23[0] = 0
    }
    switch -2147483648 {
    case 231:
        _ = len(v1)
    }
    if v25 {
        fmt.Println(v4, v2, v2)
        switch (v5 + v5) {
        case "\n\t":
            if v25 {
                v28 := GenS1{F0: v5}
                _ = v28
            }
        case "hello":
            for fi29 := 0; fi29 < 15; fi29++ {
                switch v5 {
                case "hello":
                    var v30 uint16 = 220
                    _ = v30
                case "":
                    if v25 {
                        var v31 int32 = v2
                        for fi32 := 0; fi32 < 16; fi32++ {
                            fmt.Println(fi29)
                            if (v5 > "\n\t") {
                                v33 := v1[42]
                                var v34 int8 = -1
                                _ = v33
                                _ = v34
                            }
                            delete(v1, 2)
                            _ = fi32
                        }
                        _ = v31
                    }
                }
                v35 := GenS3{F0: 1, F1: -242, F2: "\n\t"}
                var v36 uint64 = 177
                _ = fi29
                _ = v35
                _ = v36
            }
        default:
            switch 850 {
            case 612:
                fmt.Println(v2)
            case -698:
                fmt.Println(v25, v5, v25)
            default:
                v1[65535] = 42
            }
        }
        fmt.Println(v25, v3)
    }
    delete(v1, 42)
    v3 = v2
    _ = len(v1)
    var v37 int32 = (v3 - v2)
    var v38 string = v5
    fmt.Println(v4, v38, v3)
    v3 += v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}

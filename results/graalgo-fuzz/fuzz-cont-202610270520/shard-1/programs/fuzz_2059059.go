package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 int; F2 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 int; F2 int64 }

func main() {
    var v1 float64 = (float64(-2.25) + 3.14159)
    v2 := make(map[string]string)
    var v3 bool = false
    var v4 int32 = -2147483648
    v5 := []int{2147483647, 355, 0, 4294967295}
    var v6 uint = 745
    var v7 uint16 = 32767
    v3 = (v3 && v3)
    switch -1 {
    case 1:
        if (2 > 18446744073709551615) {
            for fi8 := 0; fi8 < 10; fi8++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = len(v2)
                _ = fi8
            }
            if v3 {
                if (443 > 4294967295) {
                    for fi9 := 0; fi9 < 15; fi9++ {
                        var v10 int32 = -1
                        for fi11 := 0; fi11 < 17; fi11++ {
                            var v12 uint64 = 4294967295
                            _ = fi11
                            _ = v12
                        }
                        var v13 uint8 = 242
                        _ = fi9
                        _ = v10
                        _ = v13
                    }
                } else {
                    var v14 uint32 = 518
                    var v15 uint = v6
                    _ = v14
                    _ = v15
                }
                v6 = v6
            }
            if (v3 || (v4 <= v4)) {
                for fi16 := 0; fi16 < 17; fi16++ {
                    for fi17 := 0; fi17 < 13; fi17++ {
                        _ = len(v5)
                        if ((v4 + (v4 | (v4 | v4))) < v4) {
                            switch fi17 {
                            case 42:
                                fmt.Println(fi16, fi16)
                            }
                            v18 := GenS1{F0: v6, F1: (v6 / ((v6) | 1))}
                            _ = v18
                        } else {
                            v5 = append(v5, (fi16 / ((fi17) | 1)))
                            var v19 string = "foo bar"
                            _ = v19
                        }
                        v20 := GenS1{F0: 35, F1: v6}
                        switch -1 {
                        case 4294967295:
                            v21 := GenS2{F0: fi16}
                            _ = v21
                        default:
                            if (v6 <= v6) {
                                if (792 <= 2147483647) {
                                    var v22 int32 = v4
                                    _ = v22
                                } else {
                                    if (!true) {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    } else {
                                        if len(v5) > 0 {
                                            v5[0] = -9223372036854775808
                                        }
                                        var v23 int32 = int32(v7)
                                        _ = v23
                                    }
                                    var v24 uint64 = 0
                                    v1 += 1.5
                                    _ = v24
                                }
                            }
                        }
                        _ = fi17
                        _ = v20
                    }
                    _ = fi16
                }
                fmt.Println(v7, v4)
            }
        }
    case -9223372036854775808:
        if false {
            v25 := GenS2{F0: 42}
            for fi26 := 0; fi26 < 19; fi26++ {
                v27 := GenS3{F0: v6, F1: (fi26 | fi26), F2: -9223372036854775808}
                v5 = append(v5, (fi26 >> (((v6 - (v6 << ((v6) & 63)))) & 63)))
                for fi28 := 0; fi28 < 16; fi28++ {
                    v3 = true
                    _ = fi28
                }
                _ = fi26
                _ = v27
            }
            _ = v25
        } else {
            _ = len(v2)
        }
    default:
        v6 = v6
    }
    v29 := GenS1{F0: (v6 ^ (v6 >> (((^(v6))) & 63))), F1: v6}
    var v30 uint = v6
    var v31 uint16 = (^(v7))
    var v32 uint = v30
    for fi33 := 0; fi33 < 1; fi33++ {
        _ = len(v5)
        v1 -= (float64(1e10) + v1)
        _ = fi33
    }
    v34 := GenS4{F0: true}
    switch "\n\t" {
    case "hello":
        if true {
            v35 := GenS2{F0: 4294967295}
            _ = v35
        }
    default:
        v32 = v6
    }
    for fi36 := 0; fi36 < 7; fi36++ {
        v37 := v2["\n\t"]
        switch (fi36 << ((v30) & 63)) {
        case 9223372036854775807:
            for fi38 := 0; fi38 < 12; fi38++ {
                fmt.Println(v31, v7)
                switch v37 {
                case "hello":
                    var v39 uint8 = 211
                    _ = v39
                case "":
                    fmt.Println(v30, fi38)
                default:
                    v40 := GenS5{F0: 42, F1: fi36, F2: 1}
                    _ = v40
                }
                fmt.Println(v1)
                _ = fi38
            }
        case -9223372036854775808:
            v41 := GenS5{F0: 9223372036854775807, F1: fi36, F2: 0}
            _ = v41
        default:
            switch (fi36 * 32767) {
            case 506:
                v32 = v30
            case -9223372036854775808:
                v5 = append(v5, fi36)
            default:
                v1 = ((v1 - v1) - (v1 / v1))
            }
        }
        v42 := GenS3{F0: (^(v32)), F1: fi36, F2: 0}
        _ = fi36
        _ = v37
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

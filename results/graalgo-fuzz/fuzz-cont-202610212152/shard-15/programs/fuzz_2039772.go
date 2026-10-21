package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int, p1 uint32, p2 int16) uint16 {
    switch "foo bar" {
    case "hello":
        for fi6 := 0; fi6 < 1; fi6++ {
            fmt.Println(fi6)
            switch "a" {
            case "\n\t":
                v7 := GenS1{F0: (3.14159 < (0.0 * 3.14159))}
                _ = v7
            }
            var v8 int32 = 42
            _ = fi6
            _ = v8
        }
    case "":
        var v9 uint64 = 42
        _ = v9
    case "\n\t":
        switch (p0 / ((p0) | 1)) {
        case 0:
            var v10 uint8 = 255
            _ = v10
        case 32767:
            if (0 >= 0) {
                for fi11 := 0; fi11 < 1; fi11++ {
                    if (136 < -1) {
                        v12 := []int{-9223372036854775808, fi11}
                        var v13 uint64 = 799
                        _ = v12
                        _ = v13
                    } else {
                        var v14 string = "a"
                        _ = v14
                    }
                    if (42 <= 4294967295) {
                        v15 := []int{-22, -2147483648, (-(fi11))}
                        v15 = append(v15, (fi11 - (p0 + (p0 % ((fi11) | 1)))))
                        _ = v15
                    }
                    if false {
                        p1 = p1
                        for fi16 := 0; fi16 < 18; fi16++ {
                            if false {
                                var v17 int8 = 1
                                if (fi11 == (fi16 - p0)) {
                                    v18 := GenS1{F0: false}
                                    var v19 float64 = 3.14159
                                    _ = v18
                                    _ = v19
                                }
                                var v20 uint = 255
                                _ = v17
                                _ = v20
                            } else {
                                switch "\n\t" {
                                case "":
                                    var v21 uint32 = (p1 / ((42) | 1))
                                    _ = v21
                                case "hello":
                                    p0 += (fi11 >> ((1) & 63))
                                    v22 := []int{}
                                    switch "foo bar" {
                                    case "a":
                                        p0 = p0
                                    }
                                    _ = v22
                                }
                            }
                            switch "hello" {
                            case "a":
                                fmt.Println(p2, fi11, fi11)
                            }
                            _ = fi16
                        }
                        p1 = (p1 ^ p1)
                    }
                    _ = fi11
                }
            } else {
                var v23 uint64 = 631
                _ = v23
            }
        case -9223372036854775808:
            fmt.Println(p1)
        }
    default:
        var v24 uint16 = 42
        _ = v24
    }
    var v25 uint = 18446744073709551615
    var v26 int8 = -40
    if false {
        var v27 uint64 = 18446744073709551615
        _ = v27
    } else {
        v26 -= v26
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    _ = v26
    return 2
}
func genFn1(p0 uint32) int16 {
    var v28 uint32 = p0
    switch ("" + "a") {
    case "\n\t":
        if (7 > -1) {
            var v29 int8 = -82
            if true {
                for fi30 := 0; fi30 < 11; fi30++ {
                    var v31 uint8 = 103
                    _ = fi30
                    _ = v31
                }
                for fi32 := 0; fi32 < 15; fi32++ {
                    v33 := GenS1{F0: (("hello" + ("a" + "foo bar")) != "\n\t")}
                    _ = fi32
                    _ = v33
                }
            }
            var v34 float64 = 0.0
            _ = v29
            _ = v34
        } else {
            var v35 bool = (42 >= 806)
            fmt.Println(v35)
            v36 := GenS2{F0: -99}
            _ = v35
            _ = v36
        }
    case "":
        for fi37 := 0; fi37 < 6; fi37++ {
            v28 = (^(p0))
            _ = fi37
        }
    default:
        p0 = (^(v28))
    }
    if false {
        var v38 bool = ((p0 << ((2) & 63)) == (v28 | (p0 | v28)))
        v39 := GenS3{F0: 4294967295}
        var v40 uint = 1
        _ = v38
        _ = v39
        _ = v40
    }
    _ = p0
    _ = v28
    return int16(v28)
}

func main() {
    var v1 int64 = 611
    var v2 int64 = (v1 - (v1 << ((18446744073709551615) & 63)))
    var v3 uint16 = 42
    v4 := []string{}
    var v5 uint16 = 32767
    v41 := GenS3{F0: 18446744073709551615}
    switch "" {
    case "\n\t":
        fmt.Println(v3, v3)
    default:
        var v42 int = 509
        _ = v42
    }
    var v43 int = 127
    switch (v43 >> ((18446744073709551615) & 63)) {
    case 128:
        if len(v4) > 0 {
            v4[0] = "a"
        }
    case 9223372036854775807:
        fmt.Println(v1, v43)
    }
    fmt.Println(v43, v43, v2)
    for fi44 := 0; fi44 < 9; fi44++ {
        switch (("\n\t" + "\n\t") + "a") {
        case "a":
            v45 := GenS2{F0: -116}
            _ = v45
        default:
            v1 = v1
        }
        _ = fi44
    }
    v2 = v2
    v1 = (v2 / (((v2 - v1)) | 1))
    v43 = v43
    var v46 int8 = -125
    v46 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v43
    _ = v46
    _ = v5
}

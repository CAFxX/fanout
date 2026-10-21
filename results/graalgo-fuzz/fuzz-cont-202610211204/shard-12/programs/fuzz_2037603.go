package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint16; F1 uint }
type GenS7 struct { F0 int16; F1 bool }
type GenS8 struct { F0 int; F1 uint8; F2 uint16 }
type GenS9 struct { F0 uint16; F1 string }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint32 }
func genFn0(p0 int16, p1 int64, p2 bool) bool {
    if p2 {
        for fi7 := 0; fi7 < 19; fi7++ {
            p0 += p0
            fmt.Println(p2, fi7, p0)
            v8 := []int{(fi7 >> ((0) & 63)), (^(fi7)), 293, fi7}
            _ = fi7
            _ = v8
        }
    }
    fmt.Println(p1)
    if p2 {
        for fi9 := 0; fi9 < 8; fi9++ {
            v10 := []int{(fi9 << ((42) & 63)), 4294967295, fi9, fi9}
            if true {
                switch (("foo bar" + ("" + ("\n\t" + "hello"))) + "hello") {
                case "a":
                    v11 := GenS1{F0: 42, F1: 65535}
                    _ = v11
                case "foo bar":
                    v10 = append(v10, fi9)
                default:
                    v10 = append(v10, (fi9 << ((18446744073709551615) & 63)))
                }
            } else {
                if false {
                    v10 = append(v10, (-(fi9)))
                    v12 := GenS2{F0: 0, F1: 0}
                    _ = v12
                }
                fmt.Println(p0, fi9, p2)
                var v13 bool = true
                _ = v13
            }
            v14 := GenS3{F0: 1}
            _ = fi9
            _ = v10
            _ = v14
        }
        if (1 > 11) {
            fmt.Println(p2, p2, p0)
            p1 += (-(p1))
            var v15 float64 = 1e10
            _ = v15
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    var v1 uint16 = 0
    var v2 uint = 4294967295
    v3 := make(map[string]bool)
    var v4 string = ""
    v5 := []int{}
    v6 := make(map[int]int)
    v16 := GenS1{F0: 127, F1: 1}
    fmt.Println(v4, v2, v4)
    v17 := GenS3{F0: 678}
    if false {
        if len(v5) > 0 {
            v5[0] = 1
        }
        var v18 string = v4
        switch 4294967295 {
        case -9223372036854775808:
            v19 := GenS4{F0: (32767 != 1), F1: 0}
            _ = v19
        case 2147483647:
            v1 = (v1 * v1)
        }
        _ = v18
    } else {
        if (!true) {
            v2 = (v2 & uint(v1))
            v20 := GenS5{F0: 2}
            fmt.Println(v4, v1, v1)
            _ = v20
        }
    }
    var v21 uint = (v2 & v2)
    for fi22 := 0; fi22 < 12; fi22++ {
        _ = len(v3)
        for fi23 := 0; fi23 < 20; fi23++ {
            switch "foo bar" {
            case "\n\t":
                fmt.Println(fi22, fi22)
            case "foo bar":
                if false {
                    v24 := GenS5{F0: 65535}
                    _ = v24
                } else {
                    var v25 int = fi23
                    switch "hello" {
                    case "foo bar":
                        if false {
                            v26 := GenS6{F0: (v1 % (((v1 / ((32767) | 1))) | 1)), F1: v21}
                            v27 := GenS5{F0: (v21 & 18446744073709551615)}
                            _ = v26
                            _ = v27
                        } else {
                            if len(v5) > 0 {
                                v5[0] = (fi23 - (-(fi22)))
                            }
                            fmt.Println(fi23, v25)
                            v2 += v2
                        }
                    case "hello":
                        fmt.Println(v2, fi22)
                    }
                    _ = v25
                }
            default:
                var v28 int32 = -616
                _ = v28
            }
            var v29 uint32 = 2147483648
            _ = fi23
            _ = v29
        }
        var v30 uint64 = 6
        _ = fi22
        _ = v30
    }
    v31 := GenS7{F0: -1, F1: (false && false)}
    switch (v4 + v4) {
    case "\n\t":
        if false {
            fmt.Println(v1, v2)
        }
    case "a":
        if (2147483647 > 1) {
            if false {
                if ((-2.25 >= 0.1) && true) {
                    v1 -= v1
                }
                var v32 int16 = 0
                if (42 > 1) {
                    switch 2147483647 {
                    case 42:
                        v33 := GenS7{F0: v32, F1: false}
                        _ = v33
                    case -392:
                        v34 := GenS8{F0: 601, F1: 127, F2: uint16(v21)}
                        _ = v34
                    case 0:
                        var v35 float64 = -2.25
                        _ = v35
                    }
                    var v36 int32 = -1
                    _ = v36
                } else {
                    var v37 int = 2147483647
                    v38 := GenS6{F0: (v1 / (((^(v1))) | 1)), F1: v21}
                    v39 := GenS9{F0: (v1 >> (((v2 >> ((uint(v1)) & 63))) & 63)), F1: v4}
                    _ = v37
                    _ = v38
                    _ = v39
                }
                _ = v32
            } else {
                if (!true) {
                    if false {
                        var v40 uint16 = (^(v1))
                        fmt.Println(v4)
                        _ = v40
                    }
                    var v41 int16 = -792
                    v42 := v3["hello"]
                    _ = v41
                    _ = v42
                }
                if false {
                    v2 += 299
                    for fi43 := 0; fi43 < 12; fi43++ {
                        v5 = append(v5, -990)
                        for fi44 := 0; fi44 < 15; fi44++ {
                            v4 = (("\n\t" + (v4 + "a")) + "hello")
                            fmt.Println(fi44)
                            var v45 int64 = 0
                            var v46 bool = true
                            _ = fi44
                            _ = v45
                            _ = v46
                        }
                        v47 := GenS5{F0: (^(v21))}
                        _ = fi43
                        _ = v47
                    }
                    var v48 string = v4
                    _ = v48
                } else {
                    var v49 uint32 = 2
                    v1 += 42
                    _ = v49
                }
                v1 = v1
            }
            fmt.Println(v1, v1)
            fmt.Println(v1)
        } else {
            for fi50 := 0; fi50 < 18; fi50++ {
                var v51 int8 = -106
                _ = fi50
                _ = v51
            }
            switch ("\n\t" + ("\n\t" + v4)) {
            case "foo bar":
                fmt.Println(v4, v21)
            case "\n\t":
                var v52 bool = (1 != -915)
                _ = v52
            }
            v2 += v2
        }
    }
    v53 := GenS10{F0: -128}
    v1 = v1
    switch ((((v4 + ("a" + (v4 + "foo bar"))) + v4) + v4) + v4) {
    case "hello":
        v54 := GenS11{F0: 4294967295}
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v53
    _ = v6
}

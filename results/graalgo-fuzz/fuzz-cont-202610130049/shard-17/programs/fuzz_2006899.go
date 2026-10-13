package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 string; F2 int64 }
type GenS4 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS5 struct { F0 int32 }
func genFn0(p0 uint16) int8 {
    switch "a" {
    case "":
        var v7 uint8 = 2
        _ = v7
    case "\n\t":
        if true {
            if (127 <= 2) {
                v8 := GenS1{F0: "\n\t", F1: "a"}
                p0 = 42
                _ = v8
            }
        } else {
            var v9 int64 = -9223372036854775808
            var v10 int16 = 0
            fmt.Println(v9)
            _ = v10
            _ = v9
        }
    case "a":
        v11 := []int{-129, 0, 2147483647}
        _ = v11
    }
    fmt.Println(p0)
    if (!((42 < -2147483648) && true)) {
        if true {
            v12 := GenS2{F0: true}
            _ = v12
        } else {
            p0 += (^(p0))
        }
        var v13 int64 = -1
        _ = v13
    }
    if false {
        switch ("" + ("a" + ("foo bar" + "\n\t"))) {
        case "foo bar":
            switch 128 {
            case 4294967295:
                if (1.5 <= 1e10) {
                    var v14 uint16 = 437
                    _ = v14
                }
            case 282:
                p0 = 1
            default:
                for fi15 := 0; fi15 < 4; fi15++ {
                    if true {
                        p0 = p0
                        var v16 int32 = 298
                        _ = v16
                    }
                    _ = fi15
                }
            }
        case "\n\t":
            if (!true) {
                var v17 uint16 = 0
                v17 += (^(p0))
                var v18 int32 = 713
                _ = v17
                _ = v18
            } else {
                switch "\n\t" {
                case "a":
                    var v19 float32 = 0.0
                    _ = v19
                case "hello":
                    var v20 float32 = 1.5
                    _ = v20
                }
            }
        }
    }
    _ = p0
    return 42
}
func genFn1(p0 int16, p1 int8) uint16 {
    switch 9223372036854775807 {
    case 4294967295:
        var v21 int = -1
        _ = v21
    }
    switch 0 {
    case -1:
        p1 = p1
    case 357:
        var v22 int64 = -885
        _ = v22
    default:
        v23 := GenS1{F0: "hello", F1: (((("foo bar" + ("a" + "foo bar")) + ("" + ("\n\t" + (("foo bar" + "a") + "hello")))) + "a") + "foo bar")}
        _ = v23
    }
    if (0.1 >= 0.0) {
        p0 += (^(p0))
        p1 += p1
        fmt.Println(p0, p0)
    } else {
        v24 := GenS3{F0: 2147483647, F1: ("a" + "a"), F2: -9223372036854775808}
        var v25 uint64 = 42
        if ((false && (v25 == v25)) && (662 < 32767)) {
            p1 = -119
            v26 := GenS3{F0: 0, F1: "foo bar", F2: 9223372036854775807}
            _ = v26
        }
        _ = v24
        _ = v25
    }
    fmt.Println(p0)
    if false {
        p0 = p0
        var v27 string = "\n\t"
        _ = v27
    } else {
        var v28 float32 = 1.5
        _ = v28
    }
    _ = p0
    _ = p1
    return 65535
}

func main() {
    var v1 int8 = -128
    var v2 string = "\n\t"
    var v3 string = v2
    var v4 int = 42
    v5 := make(map[uint32]int)
    var v6 uint64 = 0
    v1 += (v1 ^ v1)
    switch "foo bar" {
    case "foo bar":
        for fi29 := 0; fi29 < 14; fi29++ {
            v30 := []int{}
            v3 = "hello"
            _ = fi29
            _ = v30
        }
    default:
        v31 := v5[42]
        _ = v31
    }
    for fi32 := 0; fi32 < 19; fi32++ {
        switch 139 {
        case -9223372036854775808:
            v3 = v3
        case 42:
            v4 = 4294967295
        default:
            switch v3 {
            case "foo bar":
                v33 := GenS1{F0: v2, F1: v2}
                _ = v33
            case "\n\t":
                v34 := GenS4{F0: 32767, F1: 127, F2: 0}
                _ = v34
            case "a":
                v35 := v5[168]
                _ = v35
            default:
                for fi36 := 0; fi36 < 7; fi36++ {
                    var v37 bool = false
                    if v37 {
                        var v38 int8 = v1
                        if (!v37) {
                            v39 := GenS5{F0: -2147483648}
                            fmt.Println(v6, fi32)
                            _ = v39
                        }
                        _ = v38
                    } else {
                        v37 = (!(42 > -1))
                        v1 -= v1
                        switch 798 {
                        case 2147483647:
                            var v40 string = v3
                            _ = v40
                        case 943:
                            v2 = ("\n\t" + (v3 + "a"))
                        }
                    }
                    var v41 int = fi32
                    v41 += fi32
                    _ = fi36
                    _ = v37
                    _ = v41
                }
            }
        }
        delete(v5, 39)
        var v42 float64 = 1.5
        _ = fi32
        _ = v42
    }
    var v43 int = -1
    v44 := GenS5{F0: int32(v6)}
    switch v4 {
    case 32768:
        v45 := []int{v4}
        _ = v45
    case 4294967295:
        v2 = "hello"
    }
    v46 := []int{-971, (v43 / ((v43) | 1)), -12}
    v3 = ("\n\t" + v2)
    var v47 uint8 = 2
    delete(v5, 4294967295)
    fmt.Println(v47)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}

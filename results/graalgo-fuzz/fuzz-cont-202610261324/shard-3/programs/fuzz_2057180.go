package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int8 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 uint32; F1 int; F2 int8 }
func genFn0(p0 int16, p1 int) string {
    for fi4 := 0; fi4 < 11; fi4++ {
        switch (("a" + "\n\t") + "hello") {
        case "foo bar":
            v5 := GenS1{F0: (p0 / ((1000) | 1)), F1: 1, F2: 5}
            _ = v5
        case "hello":
            var v6 uint = 18446744073709551615
            _ = v6
        }
        var v7 int16 = -32768
        var v8 string = ("" + "foo bar")
        _ = fi4
        _ = v7
        _ = v8
    }
    p0 = (p0 << ((255) & 63))
    _ = p0
    _ = p1
    return "\n\t"
}
func genFn1(p0 uint8, p1 uint64, p2 bool) string {
    v9 := GenS1{F0: 0, F1: -2147483648, F2: 42}
    fmt.Println(p2, p0)
    if p2 {
        var v10 uint = 2
        var v11 bool = (!(-2147483648 <= 175))
        _ = v10
        _ = v11
    } else {
        v12 := GenS1{F0: -32768, F1: 32767, F2: 1}
        _ = v12
    }
    v13 := []int{397}
    for fi14 := 0; fi14 < 18; fi14++ {
        switch ("foo bar" + ("\n\t" + "\n\t")) {
        case "":
            v15 := GenS1{F0: 1, F1: fi14, F2: -128}
            _ = v15
        case "foo bar":
            v16 := GenS2{F0: (!(p1 <= 1)), F1: 0}
            _ = v16
        default:
            p2 = (p2 || p2)
        }
        fmt.Println(fi14, fi14, fi14)
        if ((p2 || (p2 && p2)) || true) {
            switch -48 {
            case 0:
                if p2 {
                    var v17 int = (fi14 << ((18446744073709551615) & 63))
                    _ = v17
                }
            case -1:
                v13 = append(v13, (^(fi14)))
            case 32767:
                if p2 {
                    v18 := GenS2{F0: p2, F1: -94}
                    if true {
                        fmt.Println(p2, p2)
                        if true {
                            fmt.Println(p1, p2, fi14)
                        }
                        fmt.Println(p0, p2)
                    }
                    v13 = append(v13, (fi14 % ((fi14) | 1)))
                    _ = v18
                } else {
                    var v19 uint64 = 4294967295
                    for fi20 := 0; fi20 < 6; fi20++ {
                        fmt.Println(p1, p2)
                        p0 -= 13
                        if (p2 || false) {
                            v19 -= 18446744073709551615
                            v21 := GenS3{F0: 2, F1: 1, F2: -113}
                            if p2 {
                                v13 = append(v13, fi14)
                            }
                            _ = v21
                        }
                        _ = fi20
                    }
                    for fi22 := 0; fi22 < 1; fi22++ {
                        var v23 int64 = -189
                        _ = fi22
                        _ = v23
                    }
                    _ = v19
                }
            default:
                var v24 int16 = 42
                _ = v24
            }
            switch "foo bar" {
            case "foo bar":
                var v25 int32 = 702
                _ = v25
            case "\n\t":
                var v26 int = fi14
                _ = v26
            }
        }
        _ = fi14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v9
    return "\n\t"
}

func main() {
    v1 := []int{}
    var v2 int8 = -67
    var v3 string = "foo bar"
    switch "\n\t" {
    case "foo bar":
        var v27 uint = 677
        _ = v27
    default:
        var v28 int16 = 32767
        _ = v28
    }
    if false {
        v2 = (v2 << ((18446744073709551615) & 63))
        fmt.Println(v3, v3)
        v3 = v3
    } else {
        var v29 uint16 = 1
        var v30 uint32 = 65535
        _ = v29
        _ = v30
    }
    v31 := GenS1{F0: 32767, F1: -1, F2: (v2 + 42)}
    v3 = ""
    switch "foo bar" {
    case "a":
        switch ("\n\t" + "") {
        case "foo bar":
            var v32 float64 = 0.0
            _ = v32
        case "":
            var v33 int32 = -2147483648
            _ = v33
        case "\n\t":
            var v34 int16 = 32767
            _ = v34
        default:
            v1 = append(v1, 9223372036854775807)
        }
    case "\n\t":
        var v35 float64 = ((0.0 - 1.5) - float64(3.14159))
        _ = v35
    case "hello":
        v2 -= 100
    default:
        v3 = v3
    }
    if len(v1) > 0 {
        v1[0] = -977
    }
    if len(v1) > 0 {
        v1[0] = -2147483648
    }
    var v36 int8 = (v2 / ((v2) | 1))
    switch -794 {
    case 395:
        var v37 uint = 557
        _ = v37
    case -331:
        v38 := GenS1{F0: -852, F1: 62, F2: 127}
        _ = v38
    case 128:
        switch "" {
        case "foo bar":
            var v39 int16 = 498
            _ = v39
        case "a":
            var v40 string = (v3 + v3)
            _ = v40
        }
    default:
        var v41 uint32 = 2147483647
        _ = v41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v36
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 uint; F2 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16; F1 uint64; F2 int8 }
type GenS7 struct { F0 uint16; F1 int16; F2 uint16 }
func genFn0(p0 uint, p1 bool, p2 int) int8 {
    if false {
        var v6 string = ""
        _ = v6
    } else {
        v7 := GenS1{F0: "foo bar", F1: 1, F2: 42}
        _ = v7
    }
    if false {
        var v8 uint8 = 43
        _ = v8
    }
    if (float64(0.0) < ((3.14159 + 1e10) / ((3.14159 * (-2.25 / 1e10)) + 0.1))) {
        switch "hello" {
        case "a":
            if (p1 || p1) {
                v9 := GenS1{F0: (((("foo bar" + ("a" + "a")) + "foo bar") + "\n\t") + ("foo bar" + "\n\t")), F1: 115, F2: 0}
                v10 := []int{(p2 % ((p2) | 1)), (-(p2)), (p2 << (((^(p0))) & 63))}
                if p1 {
                    var v11 float32 = 0.0
                    p0 += p0
                    _ = v11
                }
                _ = v10
                _ = v9
            }
        case "\n\t":
            p0 = (p0 | p0)
        default:
            p1 = ((129 > 0) && p1)
        }
        switch "a" {
        case "hello":
            var v12 uint = p0
            _ = v12
        case "":
            if ((p0 + (p0 % ((p0) | 1))) > (p0 * p0)) {
                var v13 int8 = -40
                var v14 float32 = (3.14159 * (3.14159 / 0.1))
                switch p2 {
                case -2147483648:
                    switch (p2 & p2) {
                    case 32767:
                        v15 := GenS1{F0: "hello", F1: -20, F2: 867}
                        _ = v15
                    case 32768:
                        for fi16 := 0; fi16 < 1; fi16++ {
                            fmt.Println(p2, p1)
                            _ = fi16
                        }
                    default:
                        if (("foo bar" + "hello") > "") {
                            for fi17 := 0; fi17 < 3; fi17++ {
                                v18 := []int{349, (fi17 % ((0) | 1)), 4294967295, (^(p2))}
                                var v19 int = p2
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                        }
                    }
                }
                _ = v13
                _ = v14
            } else {
                p2 = (p2 >> ((p0) & 63))
                var v20 uint = 42
                _ = v20
            }
        case "foo bar":
            p2 -= p2
        }
        fmt.Println(p1, p2)
    } else {
        var v21 int16 = -117
        var v22 float32 = 1.5
        _ = v21
        _ = v22
    }
    if false {
        var v23 uint8 = 42
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    return -128
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{"", "\n\t"}
    v3 := make(map[int]bool)
    var v4 uint32 = 1
    var v5 int = -9223372036854775808
    if len(v2) > 0 {
        v2[0] = ""
    }
    _ = len(v3)
    v24 := GenS2{F0: 1}
    v25 := GenS3{F0: 0}
    if true {
        if true {
            if len(v2) > 0 {
                v2[0] = "a"
            }
            switch "\n\t" {
            case "a":
                for fi26 := 0; fi26 < 4; fi26++ {
                    _ = len(v1)
                    v27 := GenS4{F0: 0, F1: 1, F2: 1}
                    _ = fi26
                    _ = v27
                }
            }
        }
        for fi28 := 0; fi28 < 12; fi28++ {
            v5 += (^(v5))
            fmt.Println(v4)
            _ = fi28
        }
    }
    delete(v1, (^(v4)))
    var v29 uint8 = 0
    v30 := GenS5{F0: 0}
    var v31 float64 = (3.14159 + 3.14159)
    switch "hello" {
    case "foo bar":
        switch (("hello" + ((("a" + (("hello" + "a") + "\n\t")) + ("" + "hello")) + "a")) + "hello") {
        case "\n\t":
            delete(v3, 521)
        }
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    for fi32 := 0; fi32 < 11; fi32++ {
        v33 := GenS6{F0: 0, F1: 0, F2: -121}
        _ = fi32
        _ = v33
    }
    var v34 uint = 2
    fmt.Println(v34, v4)
    if true {
        v31 = (float64(3.14159) / 1e10)
        v5 = (^(v5))
    } else {
        if (false || (false || true)) {
            var v35 string = (("\n\t" + (("\n\t" + "") + "")) + "hello")
            v34 -= (^(v34))
            v36 := GenS6{F0: 42, F1: 18446744073709551615, F2: -128}
            _ = v35
            _ = v36
        }
        v34 = 255
        if false {
            if (!true) {
                v37 := GenS7{F0: 1, F1: 32767, F2: 908}
                var v38 uint32 = (^(v4))
                _ = v37
                _ = v38
            }
            v34 = (v34 + 255)
            var v39 int8 = -111
            _ = v39
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v5
}

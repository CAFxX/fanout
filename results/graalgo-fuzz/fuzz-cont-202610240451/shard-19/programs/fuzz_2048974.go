package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint; F1 bool; F2 uint16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64; F1 string }
func genFn0(p0 string) string {
    var v4 uint = 146
    var v5 uint16 = 1
    v6 := []int{9223372036854775807, 42, 128, -847}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return p0
}
func genFn1(p0 int8, p1 uint64) uint32 {
    var v7 float64 = 0.0
    if (!true) {
        p0 -= (p0 << ((1) & 63))
        v8 := []int{-1, -2147483648, -2147483648}
        _ = v8
    } else {
        if (1 == 1000) {
            switch 555 {
            case 4294967295:
                var v9 uint64 = 42
                _ = v9
            }
        }
        fmt.Println(p1)
        if ((p1 > 2) && (!true)) {
            p1 += p1
        }
    }
    if false {
        switch "\n\t" {
        case "hello":
            fmt.Println(p1)
        case "a":
            for fi10 := 0; fi10 < 10; fi10++ {
                v11 := GenS1{F0: -339}
                p1 = (p1 >> ((4294967295) & 63))
                v12 := []int{128, 128, (-(fi10)), 544}
                _ = fi10
                _ = v11
                _ = v12
            }
        case "":
            v13 := []int{}
            _ = v13
        }
        var v14 string = ""
        var v15 string = (("a" + "hello") + v14)
        _ = v14
        _ = v15
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = v7
    return 1
}

func main() {
    var v1 string = ""
    var v2 int32 = 42
    var v3 int16 = -785
    v1 = "a"
    for fi16 := 0; fi16 < 5; fi16++ {
        if (((v1 + ("" + ((v1 + (v1 + v1)) + v1))) + v1) < (v1 + v1)) {
            var v17 float64 = 1.5
            _ = v17
        }
        v2 += v2
        _ = fi16
    }
    v18 := GenS1{F0: -463}
    v19 := []int{-1, -945, 32768}
    fmt.Println(v2, v2)
    v19 = append(v19, 0)
    v2 += (v2 << ((18446744073709551615) & 63))
    if false {
        switch v1 {
        case "\n\t":
            for fi20 := 0; fi20 < 13; fi20++ {
                if (false && false) {
                    fmt.Println(v3, v3, v3)
                    v2 -= -2147483648
                    v21 := GenS1{F0: 1}
                    _ = v21
                }
                _ = fi20
            }
        case "hello":
            for fi22 := 0; fi22 < 15; fi22++ {
                v23 := GenS2{F0: v3, F1: 2}
                var v24 bool = true
                _ = fi22
                _ = v23
                _ = v24
            }
        case "":
            if (127 != 127) {
                var v25 int8 = -128
                _ = v25
            } else {
                if len(v19) > 0 {
                    v19[0] = 0
                }
                var v26 string = v1
                _ = v26
            }
        }
        for fi27 := 0; fi27 < 11; fi27++ {
            _ = len(v19)
            v28 := GenS2{F0: v3, F1: 613}
            _ = len(v19)
            _ = fi27
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 17; fi29++ {
        var v30 int64 = -438
        switch (fi29 >> ((4294967295) & 63)) {
        case 885:
            switch v1 {
            case "":
                for fi31 := 0; fi31 < 18; fi31++ {
                    if (!(42 < 1)) {
                        switch (^(fi31)) {
                        case -836:
                            v30 -= (v30 ^ v30)
                        case 307:
                            v32 := GenS3{F0: 1, F1: false, F2: 208}
                            _ = v32
                        default:
                            v30 -= v30
                        }
                    }
                    _ = fi31
                }
            case "hello":
                var v33 int32 = (v2 * 42)
                _ = v33
            }
        }
        v34 := GenS4{F0: 827}
        _ = fi29
        _ = v30
        _ = v34
    }
    if true {
        var v35 string = v1
        for fi36 := 0; fi36 < 14; fi36++ {
            fmt.Println(v1, v3, v35)
            switch v35 {
            case "a":
                if (1 == 42) {
                    if true {
                        v37 := GenS5{F0: 739}
                        _ = v37
                    } else {
                        for fi38 := 0; fi38 < 18; fi38++ {
                            fmt.Println(v2)
                            v39 := GenS6{F0: (v2 / ((v2) | 1))}
                            _ = fi38
                            _ = v39
                        }
                    }
                } else {
                    var v40 uint8 = 255
                    for fi41 := 0; fi41 < 2; fi41++ {
                        fmt.Println(fi36, v35)
                        var v42 string = v1
                        if len(v19) > 0 {
                            v19[0] = fi36
                        }
                        _ = fi41
                        _ = v42
                    }
                    _ = v40
                }
            case "foo bar":
                var v43 float32 = 1.5
                _ = v43
            default:
                v19 = append(v19, fi36)
            }
            v44 := GenS7{F0: -1, F1: "hello"}
            _ = fi36
            _ = v44
        }
        for fi45 := 0; fi45 < 4; fi45++ {
            fmt.Println(fi45, v1, v3)
            _ = fi45
        }
        _ = v35
    }
    v19 = append(v19, -242)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}

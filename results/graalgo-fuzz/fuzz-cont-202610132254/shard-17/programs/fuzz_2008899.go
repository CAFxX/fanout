package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int32; F1 int64 }
func genFn0(p0 int32) int64 {
    var v6 int16 = 100
    fmt.Println(v6)
    p0 = p0
    switch 0 {
    case -479:
        v6 += -1
    }
    _ = p0
    _ = v6
    return -1
}
func genFn1(p0 int32, p1 string, p2 int16) uint64 {
    if false {
        if (!false) {
            for fi7 := 0; fi7 < 13; fi7++ {
                v8 := []int{(-(fi7)), fi7, -964}
                if (2147483648 > 2) {
                    switch (fi7 / ((833) | 1)) {
                    case 42:
                        v8 = append(v8, (-(fi7)))
                    case 712:
                        switch "" {
                        case "hello":
                            if len(v8) > 0 {
                                v8[0] = fi7
                            }
                        }
                    case -2147483648:
                        p1 = (p1 + "hello")
                    default:
                        for fi9 := 0; fi9 < 8; fi9++ {
                            v10 := GenS1{F0: 4294967295}
                            _ = fi9
                            _ = v10
                        }
                    }
                    var v11 uint8 = 0
                    if (true && (false || false)) {
                        var v12 string = p1
                        _ = v12
                    }
                    _ = v11
                }
                _ = fi7
                _ = v8
            }
        } else {
            p2 = (p2 ^ p2)
        }
        var v13 uint32 = 2147483647
        p0 = p0
        _ = v13
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        switch p1 {
        case "hello":
            for fi15 := 0; fi15 < 3; fi15++ {
                p0 += (p0 & p0)
                _ = fi15
            }
        case "foo bar":
            v16 := GenS1{F0: 42}
            _ = v16
        }
        p0 += 2147483647
        var v17 uint = 392
        _ = fi14
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    return 769
}

func main() {
    var v1 string = ("hello" + "a")
    var v2 int32 = 945
    v3 := make(map[int]string)
    v4 := make(map[uint32]string)
    var v5 uint8 = 55
    v1 = ("" + v1)
    if true {
        var v18 int = 4294967295
        v19 := GenS2{F0: -1, F1: -32768}
        _ = v18
        _ = v19
    } else {
        if false {
            fmt.Println(v5, v5, v1)
            for fi20 := 0; fi20 < 1; fi20++ {
                var v21 uint32 = 335
                if (((v5 >= 0) && (!(42 >= 18446744073709551615))) || true) {
                    switch (^(fi20)) {
                    case -944:
                        v22 := v4[1]
                        _ = v22
                    case 32767:
                        var v23 int64 = 42
                        _ = v23
                    case 2147483647:
                        fmt.Println(v2, v21, v1)
                    default:
                        var v24 int16 = 0
                        _ = v24
                    }
                    var v25 int32 = (v2 | (-(v2)))
                    _ = v25
                }
                _ = fi20
                _ = v21
            }
            for fi26 := 0; fi26 < 10; fi26++ {
                v2 += (v2 ^ (v2 + (v2 & v2)))
                switch v1 {
                case "hello":
                    v27 := []int{(fi26 - 9223372036854775807), fi26, fi26, fi26}
                    _ = v27
                case "a":
                    var v28 int64 = 42
                    _ = v28
                default:
                    v29 := []int{(-(fi26))}
                    _ = v29
                }
                if false {
                    var v30 bool = ((genFn1((v2 >> ((1) & 63)), "", -1) != 47) && true)
                    var v31 int32 = v2
                    v32 := GenS2{F0: v31, F1: 42}
                    _ = v30
                    _ = v31
                    _ = v32
                }
                _ = fi26
            }
        } else {
            v33 := GenS3{F0: -406, F1: 9223372036854775807}
            var v34 float64 = -2.25
            _ = v33
            _ = v34
        }
        var v35 uint32 = 65535
        _ = v35
    }
    for fi36 := 0; fi36 < 8; fi36++ {
        var v37 uint8 = (v5 ^ (v5 & 61))
        v38 := GenS1{F0: 91}
        for fi39 := 0; fi39 < 2; fi39++ {
            switch (-(fi39)) {
            case 0:
                var v40 float64 = (-2.25 * (1e10 - 1.5))
                _ = v40
            case 9223372036854775807:
                fmt.Println(fi36)
            case -146:
                switch "foo bar" {
                case "":
                    var v41 uint32 = 903
                    _ = v41
                case "\n\t":
                    var v42 uint16 = 0
                    _ = v42
                }
            }
            var v43 float32 = 0.0
            v3[fi39] = v1
            _ = fi39
            _ = v43
        }
        _ = fi36
        _ = v37
        _ = v38
    }
    fmt.Println(v5)
    var v44 float32 = -2.25
    var v45 string = v1
    v44 = 1.5
    fmt.Println(v44)
    for fi46 := 0; fi46 < 5; fi46++ {
        v44 = v44
        if false {
            var v47 int16 = 1000
            _ = v47
        }
        var v48 int64 = 42
        _ = fi46
        _ = v48
    }
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
}

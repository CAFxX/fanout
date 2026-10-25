package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 string; F1 uint32 }
func genFn0(p0 int16, p1 int8, p2 bool) string {
    var v4 int16 = (p0 * p0)
    for fi5 := 0; fi5 < 6; fi5++ {
        p0 = (p0 & (^(p0)))
        for fi6 := 0; fi6 < 18; fi6++ {
            p1 -= 42
            if (((1e10 / 1.0) > 0.0) && true) {
                p2 = (2 == 42)
            } else {
                p2 = p2
            }
            v7 := GenS1{F0: fi5}
            _ = fi6
            _ = v7
        }
        v4 = 1000
        _ = fi5
    }
    var v8 float64 = 0.1
    var v9 int32 = -1
    var v10 uint64 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v4
    _ = v8
    _ = v9
    return "hello"
}
func genFn1() uint64 {
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 int = 0
        switch (fi11 - 249) {
        case -1:
            var v13 uint32 = 1
            _ = v13
        case 82:
            if true {
                for fi14 := 0; fi14 < 19; fi14++ {
                    var v15 string = "hello"
                    fmt.Println(fi14, fi11)
                    _ = fi14
                    _ = v15
                }
                v12 = 128
            } else {
                v16 := []int{}
                v17 := GenS2{F0: -749, F1: 32767}
                _ = v16
                _ = v17
            }
        default:
            v18 := []int{}
            _ = v18
        }
        _ = fi11
        _ = v12
    }
    var v19 bool = ("foo bar" == "hello")
    v19 = false
    switch "hello" {
    case "foo bar":
        v20 := GenS3{F0: "hello", F1: 65535}
        _ = v20
    }
    switch 4294967295 {
    case -1:
        v19 = false
    case 358:
        switch 128 {
        case 32767:
            v21 := []int{}
            _ = v21
        case -925:
            v19 = true
        case 1:
            if true {
                if true {
                    switch 936 {
                    case 1:
                        var v22 int16 = -32768
                        _ = v22
                    }
                }
                v19 = ((!(524 < 0)) || v19)
                fmt.Println(v19)
            }
        default:
            var v23 uint8 = 2
            _ = v23
        }
    case -2147483648:
        fmt.Println(v19)
    default:
        v24 := []int{127}
        _ = v24
    }
    _ = v19
    return 18446744073709551615
}

func main() {
    var v1 int8 = 42
    var v2 uint16 = 819
    var v3 int = -9223372036854775808
    v25 := GenS1{F0: v3}
    v26 := GenS2{F0: 731, F1: 1000}
    for fi27 := 0; fi27 < 16; fi27++ {
        var v28 int8 = (v1 ^ (v1 * (v1 / ((v1) | 1))))
        v29 := []int{128}
        var v30 uint8 = 128
        _ = fi27
        _ = v28
        _ = v29
        _ = v30
    }
    switch "\n\t" {
    case "\n\t":
        switch "\n\t" {
        case "\n\t":
            var v31 int = -675
            _ = v31
        }
    case "hello":
        if (0.0 != 3.14159) {
            if false {
                v1 -= (v1 % ((100) | 1))
                for fi32 := 0; fi32 < 12; fi32++ {
                    v2 -= (v2 * (v2 / ((v2) | 1)))
                    switch (("\n\t" + "\n\t") + "") {
                    case "foo bar":
                        var v33 uint32 = 2147483647
                        _ = v33
                    case "":
                        for fi34 := 0; fi34 < 8; fi34++ {
                            var v35 int64 = 55
                            v36 := GenS2{F0: (v35 * (v35 << ((4294967295) & 63))), F1: 1000}
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                    default:
                        v2 += (^(v2))
                    }
                    _ = fi32
                }
                var v37 int64 = -9223372036854775808
                _ = v37
            }
        } else {
            if (0 > 42) {
                v2 -= 464
                v3 -= 0
            }
            fmt.Println(v2, v3)
        }
    case "a":
        v38 := GenS1{F0: (^(v3))}
        _ = v38
    }
    v1 += (v1 + 0)
    v3 -= v3
    v3 = 1
    for fi39 := 0; fi39 < 3; fi39++ {
        for fi40 := 0; fi40 < 10; fi40++ {
            v41 := []int{fi40, fi39}
            var v42 uint32 = 0
            if true {
                switch "" {
                case "":
                    v43 := GenS3{F0: "", F1: v42}
                    _ = v43
                case "foo bar":
                    v1 += 127
                default:
                    var v44 float64 = 0.1
                    _ = v44
                }
                if false {
                    var v45 uint = 42
                    v46 := GenS3{F0: "", F1: v42}
                    var v47 uint = v45
                    _ = v45
                    _ = v46
                    _ = v47
                } else {
                    v2 = 719
                    var v48 int = fi40
                    v49 := GenS3{F0: "", F1: v42}
                    _ = len(v41)
                    _ = v48
                    _ = v49
                }
                var v50 uint8 = 2
                _ = v50
            } else {
                _ = len(v41)
            }
            _ = fi40
            _ = v41
            _ = v42
        }
        var v51 uint32 = 42
        _ = fi39
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
}

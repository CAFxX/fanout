package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 int32; F2 int }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int64, p1 int16) int16 {
    switch ("hello" + "\n\t") {
    case "":
        p1 = p1
    case "a":
        var v6 float32 = 3.14159
        _ = v6
    case "foo bar":
        switch 128 {
        case 21:
            for fi7 := 0; fi7 < 13; fi7++ {
                if true {
                    v8 := []int{-9223372036854775808}
                    if (844 > 4294967295) {
                        v9 := GenS1{F0: "a"}
                        _ = v9
                    } else {
                        p1 += (^(p1))
                    }
                    for fi10 := 0; fi10 < 14; fi10++ {
                        if len(v8) > 0 {
                            v8[0] = (fi10 << ((917) & 63))
                        }
                        _ = fi10
                    }
                    _ = v8
                }
                _ = fi7
            }
        case 705:
            var v11 int = int(p1)
            _ = v11
        }
    default:
        var v12 float64 = ((1.5 + 3.14159) * 1e10)
        _ = v12
    }
    p1 += p1
    _ = p0
    _ = p1
    return p1
}

func main() {
    var v1 uint8 = 2
    var v2 uint = 444
    v3 := make(map[string]int)
    var v4 int8 = 0
    v5 := []int{-711, 0}
    v1 = v1
    v1 = v1
    var v13 int = -9223372036854775808
    fmt.Println(v4)
    if (-671 == v13) {
        v14 := GenS2{F0: v2, F1: 89, F2: -1}
        var v15 uint = v2
        _ = v14
        _ = v15
    }
    switch -238 {
    case 32767:
        switch ("a" + "hello") {
        case "foo bar":
            if (127 == 32767) {
                if (!false) {
                    v1 -= v1
                    v4 -= v4
                    v13 -= (v13 / ((-2147483648) | 1))
                } else {
                    v13 = (-(v13))
                }
                var v16 int = 42
                _ = v16
            }
        case "hello":
            if (!true) {
                if (v1 == v1) {
                    for fi17 := 0; fi17 < 20; fi17++ {
                        switch ("" + ("" + ("a" + "foo bar"))) {
                        case "hello":
                            var v18 int64 = 1
                            _ = v18
                        case "":
                            fmt.Println(v13, fi17)
                        }
                        _ = fi17
                    }
                }
            }
        }
    }
    var v19 float64 = 3.14159
    for fi20 := 0; fi20 < 10; fi20++ {
        if true {
            var v21 uint = 761
            _ = v21
        }
        _ = fi20
    }
    v22 := GenS1{F0: ""}
    var v23 int = (v13 / ((v13) | 1))
    for fi24 := 0; fi24 < 20; fi24++ {
        v25 := GenS3{F0: 0}
        _ = fi24
        _ = v25
    }
    v1 = (v1 / (((v1 ^ 0)) | 1))
    v23 = 74
    for fi26 := 0; fi26 < 2; fi26++ {
        if (439 <= 2) {
            switch "hello" {
            case "hello":
                fmt.Println(fi26, v2)
            default:
                v13 = fi26
            }
            if true {
                if (42 > 2147483647) {
                    for fi27 := 0; fi27 < 11; fi27++ {
                        var v28 uint64 = 18446744073709551615
                        _ = fi27
                        _ = v28
                    }
                    var v29 uint = (v2 / (((v2 << ((639) & 63))) | 1))
                    _ = v29
                }
            }
        } else {
            v23 -= (^(v13))
            var v30 uint32 = 65535
            for fi31 := 0; fi31 < 17; fi31++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi31
            }
            _ = v30
        }
        var v32 bool = false
        switch v23 {
        case -994:
            _ = len(v3)
        case 2147483647:
            v1 = 2
        default:
            var v33 uint16 = 546
            _ = v33
        }
        _ = fi26
        _ = v32
    }
    v34 := GenS2{F0: (v2 << ((v2) & 63)), F1: -2147483648, F2: (^(v13))}
    var v35 bool = (!true)
    v36 := GenS1{F0: ""}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}

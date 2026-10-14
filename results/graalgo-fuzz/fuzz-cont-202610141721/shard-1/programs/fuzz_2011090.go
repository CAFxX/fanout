package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 bool; F2 int16 }
func genFn0(p0 int64, p1 uint32, p2 uint8) string {
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := GenS1{F0: p0, F1: int32(p2), F2: (-(fi7))}
        v9 := []int{(^(fi7))}
        _ = fi7
        _ = v8
        _ = v9
    }
    if false {
        v10 := []int{}
        p2 += p2
        if true {
            for fi11 := 0; fi11 < 5; fi11++ {
                fmt.Println(p1, p2, p1)
                _ = fi11
            }
            var v12 int16 = 42
            _ = v12
        }
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return "hello"
}
func genFn1(p0 int8) string {
    if false {
        p0 -= (p0 % (((p0 & p0)) | 1))
        var v13 uint = 920
        if false {
            var v14 uint64 = 4294967295
            switch 128 {
            case 32768:
                p0 = p0
            case 0:
                for fi15 := 0; fi15 < 5; fi15++ {
                    var v16 uint8 = 128
                    _ = fi15
                    _ = v16
                }
            case 1:
                v13 = (v13 ^ (v13 % (((v13 % ((42) | 1))) | 1)))
            default:
                var v17 uint64 = 880
                _ = v17
            }
            switch ("" + (("a" + "hello") + "foo bar")) {
            case "":
                var v18 string = ("a" + "foo bar")
                _ = v18
            case "hello":
                v19 := GenS1{F0: -576, F1: 1, F2: -9223372036854775808}
                _ = v19
            case "foo bar":
                var v20 uint8 = 124
                _ = v20
            }
            _ = v14
        } else {
            p0 -= -128
            var v21 bool = (32767 < 65535)
            p0 += (^(p0))
            _ = v21
        }
        _ = v13
    }
    var v22 int32 = -1
    v22 += (v22 * (v22 + 42))
    _ = p0
    _ = v22
    return ("foo bar" + ((("\n\t" + "a") + "foo bar") + "\n\t"))
}

func main() {
    var v1 int8 = -128
    v2 := make(map[int]int)
    var v3 string = "\n\t"
    v4 := make(map[string]int)
    var v5 uint16 = 0
    var v6 uint16 = v5
    v5 = v6
    if ((1e10 < 0.0) && false) {
        var v23 int16 = 1
        switch -9223372036854775808 {
        case 9223372036854775807:
            for fi24 := 0; fi24 < 17; fi24++ {
                var v25 int8 = v1
                _ = fi24
                _ = v25
            }
        case 2147483647:
            v26 := GenS1{F0: 921, F1: -609, F2: 9223372036854775807}
            _ = v26
        case 0:
            v27 := GenS2{F0: 1, F1: (v5 % (((v6 >> ((4294967295) & 63))) | 1))}
            _ = v27
        }
        _ = v23
    }
    if true {
        v5 = v6
        var v28 bool = (false || (0 >= (v1 * v1)))
        _ = v28
    } else {
        if false {
            v1 += -58
            switch v3 {
            case "\n\t":
                v29 := GenS3{F0: 453}
                _ = v29
            default:
                switch 2147483647 {
                case -2147483648:
                    var v30 uint64 = 18446744073709551615
                    _ = v30
                case -1:
                    v5 = v5
                default:
                    for fi31 := 0; fi31 < 8; fi31++ {
                        v32 := []int{(fi31 % ((127) | 1)), (fi31 - 0)}
                        v3 = v3
                        _ = fi31
                        _ = v32
                    }
                }
            }
        }
    }
    v1 -= v1
    v33 := GenS4{F0: v1, F1: ((65535 == 2147483648) && (false || true)), F2: -1}
    fmt.Println(v5)
    var v34 uint32 = 65535
    if true {
        v1 += (v1 ^ (v1 + v1))
        switch "" {
        case "hello":
            v35 := GenS1{F0: -1, F1: 42, F2: 128}
            _ = v35
        case "":
            for fi36 := 0; fi36 < 9; fi36++ {
                switch "hello" {
                case "foo bar":
                    fmt.Println(v6)
                default:
                    var v37 int8 = v1
                    _ = v37
                }
                fmt.Println(fi36, fi36)
                for fi38 := 0; fi38 < 9; fi38++ {
                    switch 162 {
                    case 32768:
                        fmt.Println(fi38, v6)
                    case 330:
                        v4[v3] = fi36
                    case 128:
                        v39 := v2[(fi36 % (((fi36 + (fi38 - (fi38 * fi38)))) | 1))]
                        _ = v39
                    }
                    v6 -= (v5 + v6)
                    _ = fi38
                }
                _ = fi36
            }
        }
        fmt.Println(v1, v34, v3)
    } else {
        var v40 uint64 = 42
        var v41 int16 = 0
        _ = v40
        _ = v41
    }
    v42 := []int{}
    fmt.Println(v1)
    var v43 int8 = v1
    if true {
        _ = len(v2)
        v42 = append(v42, 494)
        v3 = (v3 + ((((v3 + ("hello" + "foo bar")) + "\n\t") + (v3 + (v3 + ("" + v3)))) + ("" + (v3 + ("" + v3)))))
    }
    fmt.Println(v6, v6)
    if (v5 >= 0) {
        if true {
            v6 -= v6
            fmt.Println(v43, v1)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}

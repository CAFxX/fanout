package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 int8; F1 string; F2 int64 }
func genFn0(p0 uint8) int64 {
    var v5 int = 32768
    for fi6 := 0; fi6 < 10; fi6++ {
        for fi7 := 0; fi7 < 3; fi7++ {
            v5 += (fi7 - (fi7 - (-(fi7))))
            switch "foo bar" {
            case "\n\t":
                if ((!true) || true) {
                    p0 = p0
                    var v8 float64 = -2.25
                    switch ("a" + "") {
                    case "hello":
                        var v9 uint64 = 703
                        _ = v9
                    case "":
                        var v10 int8 = 127
                        _ = v10
                    case "foo bar":
                        var v11 int8 = 83
                        _ = v11
                    default:
                        v12 := GenS1{F0: "\n\t", F1: false}
                        _ = v12
                    }
                    _ = v8
                }
            default:
                v13 := []int{int(p0), -9223372036854775808, v5}
                _ = v13
            }
            v14 := GenS1{F0: "foo bar", F1: true}
            _ = fi7
            _ = v14
        }
        p0 = p0
        for fi15 := 0; fi15 < 18; fi15++ {
            v16 := GenS2{F0: p0, F1: 1}
            _ = fi15
            _ = v16
        }
        _ = fi6
    }
    var v17 int64 = 1
    var v18 int8 = -48
    switch v5 {
    case -102:
        v19 := []int{(v5 * (v5 - 2147483647)), (v5 << ((932) & 63)), (^(v5))}
        _ = v19
    case -501:
        v20 := GenS2{F0: p0, F1: 255}
        _ = v20
    }
    _ = p0
    _ = v17
    _ = v18
    _ = v5
    return 0
}
func genFn1(p0 bool) int {
    var v21 uint = 0
    if p0 {
        v22 := []int{}
        _ = v22
    }
    p0 = (358 < 4294967295)
    _ = p0
    _ = v21
    return -541
}

func main() {
    var v1 int = -9223372036854775808
    var v2 int = 32767
    var v3 uint16 = 2
    v4 := []uint32{734, 65535}
    v4 = append(v4, 0)
    var v23 float64 = 1.5
    var v24 string = ("hello" + "")
    v3 = (v3 ^ 81)
    v25 := GenS1{F0: ("a" + ""), F1: false}
    _ = len(v4)
    v24 = v24
    switch (v24 + v24) {
    case "a":
        if ("foo bar" < ((("hello" + "\n\t") + v24) + v24)) {
            if ((!((genFn0(255) > int64(v3)) || true)) && (65535 != 174)) {
                for fi26 := 0; fi26 < 1; fi26++ {
                    var v27 bool = false
                    var v28 uint8 = 126
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            } else {
                if (127 < 93) {
                    v29 := GenS3{F0: 121, F1: 9223372036854775807}
                    v24 = ("a" + ("" + v24))
                    _ = v29
                }
            }
            var v30 string = "\n\t"
            _ = v30
        } else {
            if (!(((v23 - (v23 * float64(0.0))) - float64(1e10)) <= v23)) {
                _ = len(v4)
                v31 := GenS4{F0: 127, F1: "a", F2: 1}
                _ = v31
            } else {
                switch "\n\t" {
                case "":
                    for fi32 := 0; fi32 < 18; fi32++ {
                        v33 := GenS3{F0: 0, F1: -126}
                        v34 := GenS3{F0: 2, F1: 1}
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                case "\n\t":
                    v35 := GenS3{F0: 999, F1: -1}
                    _ = v35
                }
                var v36 string = ((v24 + (v24 + (v24 + v24))) + "hello")
                _ = v36
            }
            if false {
                v4 = append(v4, 2)
                v37 := GenS2{F0: 212, F1: (v3 - v3)}
                _ = v37
            }
        }
    }
    var v38 uint32 = 408
    v24 = ("foo bar" + v24)
    if true {
        var v39 bool = true
        v24 = v24
        fmt.Println(v23, v24, v2)
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v38
    _ = v4
}

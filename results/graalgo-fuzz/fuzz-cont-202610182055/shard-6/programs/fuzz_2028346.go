package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int64 }
func genFn0(p0 uint8) int8 {
    var v6 uint32 = 1
    v7 := GenS1{F0: -32768}
    fmt.Println(v6)
    _ = p0
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 int32) uint64 {
    fmt.Println(p0)
    if false {
        var v8 int64 = -1
        fmt.Println(v8, v8)
        p0 = (p0 ^ (p0 | (p0 >> ((18446744073709551615) & 63))))
        _ = v8
    }
    _ = p0
    return 1
}

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = 0.1
    v3 := []string{"foo bar", (("" + "a") + "foo bar"), "\n\t", ""}
    var v4 uint = 1
    v5 := make(map[string]bool)
    fmt.Println(v2, v4)
    for fi9 := 0; fi9 < 8; fi9++ {
        v2 = v2
        v2 += 0.0
        var v10 uint32 = 42
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 float32 = 3.14159
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        v2 = 3.14159
        _ = len(v5)
        switch "hello" {
        case "\n\t":
            v14 := v5[""]
            _ = v14
        case "foo bar":
            if false {
                fmt.Println(fi13, v4, v4)
                switch fi13 {
                case 1:
                    var v15 int16 = 32767
                    _ = v15
                }
            }
        case "":
            for fi16 := 0; fi16 < 5; fi16++ {
                switch (fi13 | fi13) {
                case 128:
                    var v17 bool = true
                    _ = v17
                case -1:
                    if len(v3) > 0 {
                        v3[0] = "a"
                    }
                case -99:
                    v18 := GenS2{F0: 1, F1: 9223372036854775807}
                    _ = v18
                default:
                    switch "hello" {
                    case "\n\t":
                        v19 := GenS2{F0: -634, F1: -1}
                        _ = v19
                    default:
                        v4 = v4
                    }
                }
                if (-1 <= 1000) {
                    var v20 uint16 = 0
                    _ = v20
                }
                v2 = 3.14159
                v21 := v5["hello"]
                v4 = (v4 | v4)
                _ = fi16
                _ = v21
            }
        }
        _ = fi13
    }
    var v22 uint32 = 42
    fmt.Println(v4, v2)
    var v23 int32 = 42
    if true {
        if false {
            v23 -= -1
            var v24 int = 4294967295
            var v25 uint8 = 255
            _ = v24
            _ = v25
        }
    }
    if (!((!true) && (!(v2 < v2)))) {
        var v26 uint64 = 42
        for fi27 := 0; fi27 < 20; fi27++ {
            switch 189 {
            case -1:
                for fi28 := 0; fi28 < 20; fi28++ {
                    v4 += v4
                    v4 = (^(v4))
                    fmt.Println(fi27, v22, v23)
                    _ = fi28
                }
            case 619:
                for fi29 := 0; fi29 < 11; fi29++ {
                    switch "foo bar" {
                    case "\n\t":
                        if (!(1 >= -33)) {
                            var v30 bool = ((^(v26)) == genFn1(-1))
                            if (!((-(fi27)) <= fi27)) {
                                fmt.Println(fi29)
                                var v31 uint16 = 2
                                _ = v31
                            } else {
                                v4 = 42
                                fmt.Println(v22, v2)
                            }
                            var v32 string = ""
                            _ = v30
                            _ = v32
                        } else {
                            var v33 int32 = v23
                            var v34 uint64 = 18446744073709551615
                            _ = v33
                            _ = v34
                        }
                    case "":
                        if false {
                            fmt.Println(v2, fi27)
                            _ = len(v1)
                        } else {
                            if len(v3) > 0 {
                                v3[0] = "foo bar"
                            }
                            v4 = v4
                        }
                    case "hello":
                        var v35 bool = false
                        _ = v35
                    }
                    _ = fi29
                }
            case 0:
                switch int(v22) {
                case -9223372036854775808:
                    v23 = (-(v23))
                case -527:
                    delete(v1, (fi27 / (((-(fi27))) | 1)))
                case 493:
                    v23 -= int32(fi27)
                default:
                    var v36 int = (fi27 ^ (fi27 / ((fi27) | 1)))
                    _ = v36
                }
            default:
                var v37 uint8 = 29
                _ = v37
            }
            if len(v3) > 0 {
                v3[0] = "\n\t"
            }
            _ = fi27
        }
        _ = v26
    } else {
        if len(v3) > 0 {
            v3[0] = ("" + "\n\t")
        }
    }
    var v38 int = -371
    v23 = (-(v23))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}

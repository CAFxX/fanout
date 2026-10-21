package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 bool; F1 string }
func genFn0() uint {
    var v7 float64 = 0.0
    if false {
        fmt.Println(v7)
    } else {
        switch ("\n\t" + "\n\t") {
        case "a":
            var v8 string = "a"
            _ = v8
        }
        if true {
            v7 -= v7
            var v9 uint16 = 440
            _ = v9
        }
        v10 := []int{2147483647, 127}
        _ = v10
    }
    if true {
        v11 := []int{1, 4294967295, 4294967295}
        _ = v11
    } else {
        v7 = v7
    }
    v7 += ((v7 / v7) - (3.14159 - 1e10))
    v12 := GenS1{F0: true}
    _ = v12
    _ = v7
    return 1000
}
func genFn1(p0 int16) string {
    v13 := GenS2{F0: 0}
    switch 9223372036854775807 {
    case 4294967295:
        v14 := GenS2{F0: 1}
        _ = v14
    case 842:
        for fi15 := 0; fi15 < 13; fi15++ {
            fmt.Println(p0)
            switch fi15 {
            case 128:
                var v16 int16 = p0
                _ = v16
            case -2147483648:
                var v17 int32 = 2147483647
                _ = v17
            default:
                p0 += (-(p0))
            }
            _ = fi15
        }
    case 1:
        p0 = (p0 - p0)
    default:
        var v18 uint = 910
        _ = v18
    }
    if false {
        var v19 uint64 = 1
        _ = v19
    } else {
        if true {
            v20 := GenS1{F0: (!(!(true && true)))}
            v21 := GenS3{F0: "", F1: 9223372036854775807}
            _ = v20
            _ = v21
        } else {
            v22 := []int{4294967295, -2147483648}
            switch "" {
            case "a":
                for fi23 := 0; fi23 < 18; fi23++ {
                    for fi24 := 0; fi24 < 7; fi24++ {
                        switch "hello" {
                        case "foo bar":
                            v22 = append(v22, (fi24 & (fi24 & fi24)))
                        default:
                            p0 += p0
                        }
                        fmt.Println(fi24, fi24, fi24)
                        p0 += p0
                        v22 = append(v22, fi24)
                        var v25 uint8 = 42
                        _ = fi24
                        _ = v25
                    }
                    var v26 uint32 = 42
                    _ = fi23
                    _ = v26
                }
            case "\n\t":
                v22 = append(v22, 127)
            default:
                var v27 int = 2147483647
                _ = v27
            }
            _ = v22
        }
        for fi28 := 0; fi28 < 19; fi28++ {
            fmt.Println(p0, fi28)
            _ = fi28
        }
    }
    var v29 uint = 18446744073709551615
    _ = p0
    _ = v13
    _ = v29
    return "a"
}

func main() {
    var v1 bool = true
    v2 := []int{9223372036854775807, 372, -527, 2147483647}
    var v3 uint32 = 4294967295
    v4 := make(map[uint32]int)
    var v5 uint = 18446744073709551615
    var v6 float32 = ((3.14159 / 1.0) / 1e10)
    if v1 {
        v30 := GenS2{F0: 18446744073709551615}
        if true {
            v4[v3] = 0
        } else {
            v5 += (v5 - v5)
            for fi31 := 0; fi31 < 9; fi31++ {
                v32 := v4[0]
                var v33 bool = (v1 || v1)
                var v34 uint = 255
                _ = fi31
                _ = v32
                _ = v33
                _ = v34
            }
            var v35 int = 32767
            _ = v35
        }
        _ = v30
    }
    var v36 int64 = -648
    delete(v4, (v3 - v3))
    var v37 int = 127
    v5 -= v5
    var v38 uint16 = 441
    var v39 bool = ("" == "foo bar")
    var v40 uint32 = (v3 | (^(v3)))
    var v41 uint = uint(v38)
    for fi42 := 0; fi42 < 11; fi42++ {
        switch "foo bar" {
        case "foo bar":
            v2 = append(v2, v37)
        default:
            switch (("foo bar" + ("a" + (("foo bar" + ("\n\t" + "")) + ("\n\t" + ("hello" + ("hello" + "hello")))))) + ("hello" + "a")) {
            case "a":
                var v43 uint = 65535
                _ = v43
            case "\n\t":
                v44 := GenS4{F0: (!((v1 || (!v39)) || v1)), F1: "foo bar"}
                _ = v44
            default:
                v5 = uint(v38)
            }
        }
        var v45 uint32 = (^(v3))
        _ = fi42
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}

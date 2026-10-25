package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 float64 = 0.1
    v2 := []string{"\n\t", "foo bar", ("a" + "a"), "\n\t"}
    var v3 float64 = v1
    v4 := make(map[int]string)
    var v5 uint = 42
    v6 := []string{}
    if (!(!(779 >= 42))) {
        v7 := GenS1{F0: 171}
        for fi8 := 0; fi8 < 12; fi8++ {
            for fi9 := 0; fi9 < 4; fi9++ {
                v1 = v1
                v10 := GenS1{F0: 36}
                _ = fi9
                _ = v10
            }
            switch 271 {
            case 127:
                var v11 uint64 = 0
                _ = v11
            case 128:
                v12 := GenS1{F0: 128}
                _ = v12
            default:
                var v13 float64 = 0.0
                _ = v13
            }
            _ = fi8
        }
        _ = v7
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch -2147483648 {
    case 32768:
        var v14 int = -1
        _ = v14
    case 450:
        if true {
            var v15 uint = 18446744073709551615
            for fi16 := 0; fi16 < 15; fi16++ {
                var v17 uint8 = 255
                v18 := v4[-143]
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = v15
        }
    default:
        var v19 uint16 = 0
        _ = v19
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        if ("foo bar" == "hello") {
            for fi21 := 0; fi21 < 3; fi21++ {
                if true {
                    switch "\n\t" {
                    case "a":
                        var v22 uint16 = 65535
                        _ = v22
                    case "foo bar":
                        switch fi20 {
                        case 0:
                            if true {
                                v4[(-(fi20))] = "hello"
                                fmt.Println(fi20)
                            }
                        case 2147483647:
                            switch "\n\t" {
                            case "\n\t":
                                if true {
                                    if (65535 <= 2147483648) {
                                        _ = len(v4)
                                        v1 -= (float64(-2.25) * 3.14159)
                                        fmt.Println(fi20)
                                        var v23 uint16 = 969
                                        var v24 uint = v5
                                        _ = v23
                                        _ = v24
                                    } else {
                                        var v25 int8 = 42
                                        var v26 int64 = int64(v5)
                                        _ = v25
                                        _ = v26
                                    }
                                    v3 -= (float64(0.1) / ((v3 - 1.5) + v1))
                                }
                            case "hello":
                                v5 -= v5
                            }
                        }
                    }
                }
                v1 = v1
                _ = fi21
            }
        } else {
            fmt.Println(v1, v3)
        }
        fmt.Println(v3)
        fmt.Println(v1)
        _ = fi20
    }
    v5 += v5
    var v27 float32 = (0.0 - -2.25)
    var v28 string = ""
    v29 := GenS1{F0: 128}
    v30 := GenS2{F0: -1}
    v3 = (v1 - (v3 * 1.5))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}

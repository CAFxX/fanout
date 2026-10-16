package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }

func main() {
    var v1 int8 = 100
    v2 := []int{-9223372036854775808}
    v3 := []string{}
    v4 := make(map[string]bool)
    var v5 uint32 = 2
    var v6 string = "\n\t"
    v5 = 65535
    switch 127 {
    case -1:
        if (v6 <= "foo bar") {
            if true {
                if true {
                    if true {
                        v5 += v5
                        v7 := GenS1{F0: -2147483648}
                        _ = v7
                    } else {
                        for fi8 := 0; fi8 < 18; fi8++ {
                            if (!((^(v1)) < v1)) {
                                var v9 uint16 = 42
                                _ = v9
                            } else {
                                var v10 float64 = -2.25
                                v1 -= v1
                                _ = v10
                            }
                            var v11 uint16 = 0
                            switch "" {
                            case "\n\t":
                                fmt.Println(fi8)
                            case "hello":
                                v2 = append(v2, 128)
                            case "":
                                v5 = v5
                            }
                            _ = fi8
                            _ = v11
                        }
                        switch 127 {
                        case 9223372036854775807:
                            var v12 uint16 = 2
                            _ = v12
                        case -9223372036854775808:
                            var v13 int32 = 0
                            _ = v13
                        case 42:
                            var v14 int8 = v1
                            _ = v14
                        default:
                            _ = len(v4)
                        }
                        _ = len(v4)
                    }
                } else {
                    v6 = (("\n\t" + ("" + v6)) + v6)
                    switch v6 {
                    case "foo bar":
                        v6 = "\n\t"
                    case "\n\t":
                        fmt.Println(v6, v5)
                    }
                    if len(v3) > 0 {
                        v3[0] = v6
                    }
                }
            }
            var v15 uint8 = 127
            var v16 bool = false
            _ = v15
            _ = v16
        } else {
            v1 -= (v1 ^ -124)
            var v17 int64 = 0
            var v18 uint16 = 2
            _ = v17
            _ = v18
        }
    }
    v4["foo bar"] = (!false)
    v19 := GenS1{F0: -2147483648}
    var v20 uint32 = v5
    for fi21 := 0; fi21 < 6; fi21++ {
        v4[v6] = true
        v5 = (v20 + 454)
        _ = fi21
    }
    var v22 int64 = 0
    for fi23 := 0; fi23 < 20; fi23++ {
        switch (fi23 & fi23) {
        case 1:
            v5 -= (v5 * (v20 * (^(v5))))
        case -1:
            switch fi23 {
            case 32768:
                fmt.Println(v1)
            case -523:
                v24 := GenS2{F0: (v6 + v6)}
                _ = v24
            default:
                switch "\n\t" {
                case "":
                    var v25 uint32 = 0
                    _ = v25
                default:
                    var v26 uint64 = 42
                    _ = v26
                }
            }
        default:
            if (!true) {
                if len(v3) > 0 {
                    v3[0] = "\n\t"
                }
                if true {
                    if (36 == 127) {
                        fmt.Println(v20, v5)
                        var v27 bool = (18446744073709551615 != 192)
                        _ = v27
                    }
                } else {
                    for fi28 := 0; fi28 < 7; fi28++ {
                        var v29 uint32 = v5
                        _ = fi28
                        _ = v29
                    }
                    var v30 int64 = (^(v22))
                    v6 = ((((v6 + "") + v6) + "foo bar") + "hello")
                    _ = v30
                }
            } else {
                var v31 uint16 = 32767
                _ = v31
            }
        }
        v32 := GenS1{F0: 522}
        _ = fi23
        _ = v32
    }
    if true {
        fmt.Println(v6)
        v20 -= v5
        var v33 int64 = -9223372036854775808
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

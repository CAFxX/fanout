package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 uint16; F2 int64 }

func main() {
    v1 := []string{"\n\t", ("\n\t" + ("\n\t" + "\n\t")), "foo bar"}
    var v2 bool = (!(false && true))
    v3 := []string{"", (("foo bar" + "foo bar") + "hello")}
    var v4 int = 64
    switch "" {
    case "a":
        switch "a" {
        case "foo bar":
            var v5 int16 = 945
            _ = v5
        case "a":
            if v2 {
                if (42 <= 255) {
                    var v6 string = "hello"
                    v7 := GenS1{F0: v4, F1: 2147483648}
                    v4 = 32768
                    _ = v6
                    _ = v7
                } else {
                    var v8 int32 = -1
                    var v9 float64 = (0.1 - 1.5)
                    v4 = v4
                    _ = v8
                    _ = v9
                }
                var v10 float32 = -2.25
                for fi11 := 0; fi11 < 1; fi11++ {
                    v2 = (!(-469 >= 104))
                    v1 = append(v1, "")
                    switch (^(fi11)) {
                    case 128:
                        switch -525 {
                        case 127:
                            var v12 int16 = 0
                            _ = v12
                        case 42:
                            fmt.Println(v4, v10, fi11)
                        case 2147483647:
                            v2 = ((v10 + 0.0) < 0.0)
                        }
                    case 2147483647:
                        if (32767 > 1) {
                            fmt.Println(fi11)
                            var v13 int = 1
                            v13 = (v4 | v4)
                            _ = v13
                        } else {
                            if len(v3) > 0 {
                                v3[0] = "hello"
                            }
                            for fi14 := 0; fi14 < 20; fi14++ {
                                v10 -= (v10 * (0.0 - v10))
                                if v2 {
                                    switch fi11 {
                                    case -470:
                                        v2 = (!(1 != 1))
                                    }
                                    v10 += 0.0
                                } else {
                                    var v15 int64 = -610
                                    v4 += (v4 * fi14)
                                    v4 += (fi14 >> ((19) & 63))
                                    _ = v15
                                }
                                _ = fi14
                            }
                        }
                    case -2147483648:
                        v10 = (1e10 * v10)
                    default:
                        var v16 uint = 0
                        _ = v16
                    }
                    _ = fi11
                }
                _ = v10
            } else {
                var v17 uint32 = 427
                _ = v17
            }
        case "hello":
            v18 := GenS2{F0: 4294967295}
            _ = v18
        default:
            for fi19 := 0; fi19 < 12; fi19++ {
                v2 = true
                v20 := GenS3{F0: v2}
                v2 = (-2.25 == 1.5)
                _ = fi19
                _ = v20
            }
        }
    case "foo bar":
        fmt.Println(v4, v4)
    default:
        v4 -= v4
    }
    v21 := GenS4{F0: 1, F1: 42, F2: -342}
    v2 = (-536 != 1)
    var v22 float32 = float32(0.0)
    switch "hello" {
    case "a":
        var v23 int64 = 42
        _ = v23
    case "hello":
        switch "" {
        case "foo bar":
            var v24 uint32 = 1
            _ = v24
        }
    case "":
        var v25 float64 = (1.5 * 1e10)
        _ = v25
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch (v4 >> ((418) & 63)) {
    case -961:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case 32768:
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
    default:
        for fi26 := 0; fi26 < 5; fi26++ {
            fmt.Println(v22, v4)
            v1 = append(v1, "\n\t")
            for fi27 := 0; fi27 < 17; fi27++ {
                v28 := GenS3{F0: false}
                switch fi27 {
                case 4294967295:
                    var v29 uint16 = 1
                    _ = v29
                case -9223372036854775808:
                    var v30 int32 = 42
                    _ = v30
                }
                _ = fi27
                _ = v28
            }
            _ = fi26
        }
    }
    fmt.Println(v4, v22)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}

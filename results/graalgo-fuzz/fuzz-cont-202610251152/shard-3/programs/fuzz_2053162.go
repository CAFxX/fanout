package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 bool }

func main() {
    v1 := []string{"", "\n\t", ("" + ("a" + ("" + "a")))}
    var v2 uint8 = 255
    var v3 uint32 = 2
    var v4 uint16 = 255
    var v5 float32 = (1.5 - 1e10)
    var v6 uint16 = (v4 / (((v4 | 567)) | 1))
    v7 := GenS1{F0: 42, F1: -1}
    v6 -= v4
    var v8 uint32 = 2
    fmt.Println(v5)
    v2 = v2
    var v9 uint = 134
    switch ("" + "") {
    case "hello":
        v5 += v5
    case "":
        if (0 < 32767) {
            v1 = append(v1, "hello")
            var v10 float64 = -2.25
            if false {
                if true {
                    switch int(v9) {
                    case 32767:
                        switch 32768 {
                        case 42:
                            v11 := GenS2{F0: true}
                            _ = v11
                        case -743:
                            v12 := GenS2{F0: false}
                            _ = v12
                        case -1:
                            var v13 float64 = v10
                            _ = v13
                        }
                    default:
                        var v14 uint16 = (v6 * (v6 & (v6 * (v4 & v4))))
                        _ = v14
                    }
                }
                var v15 int8 = 42
                var v16 uint = (^(v9))
                _ = v15
                _ = v16
            }
            _ = v10
        }
    }
    var v17 int = 128
    for fi18 := 0; fi18 < 20; fi18++ {
        for fi19 := 0; fi19 < 16; fi19++ {
            v5 = (v5 - (v5 * 1e10))
            var v20 int = fi19
            _ = fi19
            _ = v20
        }
        fmt.Println(v9, v9, fi18)
        _ = fi18
    }
    fmt.Println(v4, v4, v9)
    for fi21 := 0; fi21 < 18; fi21++ {
        fmt.Println(v2, v9, fi21)
        v9 = v9
        _ = fi21
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        v5 -= v5
        _ = fi22
    }
    fmt.Println(v17, v8, v3)
    v2 = 1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        if false {
            v2 += (^(v2))
        }
        switch "" {
        case "foo bar":
            var v24 uint8 = (v2 - 127)
            _ = v24
        default:
            switch "\n\t" {
            case "hello":
                _ = len(v1)
            case "":
                if (!false) {
                    if ((v2 - v2) <= 0) {
                        switch "\n\t" {
                        case "foo bar":
                            switch (fi23 * fi23) {
                            case 42:
                                for fi25 := 0; fi25 < 8; fi25++ {
                                    var v26 bool = true
                                    v6 -= v6
                                    var v27 int32 = -807
                                    var v28 int = (^(v17))
                                    v17 -= 182
                                    _ = fi25
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                            default:
                                var v29 int = (fi23 % ((fi23) | 1))
                                _ = v29
                            }
                        }
                        fmt.Println(v17, v17, v6)
                        _ = len(v1)
                    }
                    if len(v1) > 0 {
                        v1[0] = "hello"
                    }
                }
            case "foo bar":
                v1 = append(v1, ("foo bar" + ("foo bar" + ("foo bar" + ("a" + ("\n\t" + "a"))))))
            }
        }
        _ = fi23
    }
    for fi30 := 0; fi30 < 4; fi30++ {
        v8 = v3
        if false {
            v4 += (^(v4))
            fmt.Println(fi30, v9)
        } else {
            for fi31 := 0; fi31 < 7; fi31++ {
                v1 = append(v1, ("foo bar" + ""))
                var v32 int8 = -128
                if false {
                    var v33 uint16 = uint16(v9)
                    _ = v33
                }
                switch "foo bar" {
                case "foo bar":
                    v1 = append(v1, "\n\t")
                case "\n\t":
                    if (42 < (v8 << ((42) & 63))) {
                        if (true || (!(true && true))) {
                            v34 := GenS1{F0: 0, F1: -1}
                            v3 = (v3 - 0)
                            _ = v34
                        }
                        fmt.Println(v3)
                    } else {
                        var v35 int32 = 836
                        var v36 int = v17
                        _ = v35
                        _ = v36
                    }
                case "":
                    var v37 int8 = (v32 & 99)
                    v9 -= v9
                    fmt.Println(v8, fi30)
                    _ = v37
                }
                _ = fi31
                _ = v32
            }
        }
        v4 = 65535
        _ = fi30
    }
    var v38 uint64 = 533
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

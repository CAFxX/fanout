package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 bool }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 int64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{65535, 2147483648}
    var v3 int8 = 127
    var v4 int = 324
    var v5 int = v4
    if false {
        v6 := GenS1{F0: int32(v5), F1: "\n\t", F2: true}
        _ = v6
    } else {
        v7 := GenS1{F0: -1, F1: "", F2: true}
        v4 = (-(v4))
        _ = v7
    }
    fmt.Println(v4, v4, v4)
    delete(v1, "")
    var v8 int16 = -32768
    var v9 uint32 = 2147483648
    if (true && ((!false) || (!(!(-2.25 >= float32(((3.14159 - 0.0) * (1e10 / 1.0)))))))) {
        if (!false) {
            var v10 float64 = 0.0
            var v11 int8 = v3
            _ = v10
            _ = v11
        }
        var v12 int8 = 42
        switch -184 {
        case 128:
            var v13 float32 = 1.5
            _ = v13
        case -319:
            switch (("\n\t" + "foo bar") + ("a" + ("" + "a"))) {
            case "\n\t":
                switch (("" + "") + "") {
                case "a":
                    switch -280 {
                    case 32767:
                        v5 = v5
                    case 42:
                        var v14 uint32 = 2147483647
                        _ = v14
                    default:
                        v15 := GenS1{F0: 0, F1: "a", F2: ((v4 << ((2) & 63)) != (v4 << ((0) & 63)))}
                        _ = v15
                    }
                case "\n\t":
                    if false {
                        if len(v2) > 0 {
                            v2[0] = 1
                        }
                        var v16 uint32 = (v9 % (((v9 & (v9 & v9))) | 1))
                        _ = v16
                    } else {
                        v9 += v9
                        var v17 uint32 = (v9 + (v9 & (v9 | (v9 ^ (^(v9))))))
                        var v18 float32 = 1e10
                        _ = v17
                        _ = v18
                    }
                case "hello":
                    var v19 float32 = float32(-2.25)
                    _ = v19
                default:
                    for fi20 := 0; fi20 < 12; fi20++ {
                        var v21 bool = true
                        v21 = true
                        _ = fi20
                        _ = v21
                    }
                }
            case "hello":
                var v22 float64 = -2.25
                _ = v22
            case "":
                if true {
                    delete(v1, "")
                    v4 = (v4 >> ((2) & 63))
                }
            default:
                v23 := GenS1{F0: 456, F1: "a", F2: false}
                _ = v23
            }
        case 2147483647:
            if false {
                v12 -= v12
            }
        }
        _ = v12
    }
    var v24 uint64 = 611
    switch "\n\t" {
    case "hello":
        var v25 int16 = (v8 - int16(v5))
        _ = v25
    case "\n\t":
        v8 -= v8
    case "foo bar":
        if false {
            for fi26 := 0; fi26 < 19; fi26++ {
                if false {
                    var v27 int = fi26
                    v28 := GenS2{F0: v24, F1: int16(v9)}
                    _ = v27
                    _ = v28
                } else {
                    v29 := GenS2{F0: 733, F1: (v8 | (^(v8)))}
                    var v30 uint = uint(v24)
                    _ = v29
                    _ = v30
                }
                _ = len(v2)
                if (0.1 <= 0.1) {
                    for fi31 := 0; fi31 < 3; fi31++ {
                        _ = len(v1)
                        switch "hello" {
                        case "a":
                            fmt.Println(v24)
                        case "hello":
                            fmt.Println(fi31)
                        }
                        _ = fi31
                    }
                    if len(v2) > 0 {
                        v2[0] = v9
                    }
                }
                _ = fi26
            }
            var v32 uint8 = 255
            for fi33 := 0; fi33 < 6; fi33++ {
                switch "hello" {
                case "foo bar":
                    for fi34 := 0; fi34 < 20; fi34++ {
                        delete(v1, "\n\t")
                        if false {
                            v35 := GenS2{F0: 4294967295, F1: 1000}
                            v32 = 2
                            fmt.Println(fi33, v32)
                            _ = v35
                        }
                        _ = fi34
                    }
                case "a":
                    fmt.Println(v32, v24, v5)
                case "\n\t":
                    fmt.Println(v32, v4)
                default:
                    if false {
                        var v36 int16 = -821
                        _ = v36
                    }
                }
                _ = fi33
            }
            _ = v32
        } else {
            if (v24 <= v24) {
                switch v5 {
                case 321:
                    v5 -= (v5 >> ((42) & 63))
                case 127:
                    v24 = 1
                case -733:
                    for fi37 := 0; fi37 < 20; fi37++ {
                        for fi38 := 0; fi38 < 11; fi38++ {
                            var v39 string = "\n\t"
                            v39 = ((v39 + "a") + v39)
                            if (0 == -913) {
                                v40 := GenS3{F0: 9223372036854775807}
                                _ = v40
                            } else {
                                var v41 uint16 = 121
                                fmt.Println(v5, v24, v5)
                                _ = v41
                            }
                            _ = fi38
                            _ = v39
                        }
                        var v42 uint16 = 862
                        _ = fi37
                        _ = v42
                    }
                }
                var v43 int32 = 0
                _ = v43
            } else {
                var v44 float32 = float32(1e10)
                _ = v44
            }
        }
    default:
        _ = len(v2)
    }
    v24 -= (v24 << ((10) & 63))
    v3 += (v3 ^ (v3 << ((2) & 63)))
    if true {
        fmt.Println(v3, v9, v9)
        v45 := v1["foo bar"]
        _ = v45
    }
    switch v5 {
    case -727:
        var v46 float64 = -2.25
        _ = v46
    case 341:
        switch "foo bar" {
        case "a":
            if (!false) {
                for fi47 := 0; fi47 < 5; fi47++ {
                    var v48 int8 = v3
                    v48 = (v48 * 1)
                    _ = fi47
                    _ = v48
                }
                v3 = (^(v3))
            }
        case "\n\t":
            var v49 int64 = -279
            _ = v49
        }
    default:
        v24 = (v24 % ((v24) | 1))
    }
    var v50 int16 = v8
    var v51 int64 = -157
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v8
    _ = v9
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 uint32 }
func genFn0(p0 int64, p1 uint8, p2 uint32) uint64 {
    switch 265 {
    case 32767:
        v4 := GenS1{F0: -698, F1: 4294967295, F2: (-(p0))}
        _ = v4
    case -1:
        v5 := []int{178}
        _ = v5
    default:
        p0 = 0
    }
    if ((^(p1)) < 255) {
        var v6 uint32 = (p2 | p2)
        var v7 int16 = 32767
        _ = v6
        _ = v7
    }
    fmt.Println(p2, p1, p1)
    v8 := GenS2{F0: "a"}
    v9 := []int{685}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    v1 := []string{("foo bar" + "\n\t"), "", "hello", "a"}
    var v2 uint8 = 0
    var v3 uint64 = 0
    if false {
        v10 := GenS1{F0: -2147483648, F1: v3, F2: -592}
        fmt.Println(v2, v3)
        _ = v10
    } else {
        v11 := GenS2{F0: "a"}
        v2 = (^(v2))
        _ = v11
    }
    if true {
        switch -1 {
        case 32767:
            for fi12 := 0; fi12 < 12; fi12++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi12
            }
        default:
            for fi13 := 0; fi13 < 8; fi13++ {
                v14 := GenS1{F0: fi13, F1: genFn0(9223372036854775807, (v2 & 42), 214), F2: -824}
                switch (((("" + ("a" + "\n\t")) + "foo bar") + ("" + "a")) + "\n\t") {
                case "a":
                    var v15 uint = 1
                    _ = v15
                default:
                    for fi16 := 0; fi16 < 20; fi16++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v17 := GenS3{F0: fi13}
                        _ = fi16
                        _ = v17
                    }
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi13
                _ = v14
            }
        }
        v2 += 128
    }
    switch 0 {
    case -9223372036854775808:
        fmt.Println(v3)
    case -2147483648:
        var v18 uint32 = uint32(v2)
        _ = v18
    }
    var v19 string = "foo bar"
    v3 = (v3 >> ((509) & 63))
    var v20 string = v19
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 float32 = 0.1
        v23 := GenS4{F0: v3, F1: 65535}
        v2 = (v2 & v2)
        _ = fi21
        _ = v22
        _ = v23
    }
    if false {
        switch 4294967295 {
        case 32768:
            v20 = v19
        default:
            v20 = v20
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = len(v1)
    } else {
        v24 := GenS1{F0: 2147483647, F1: v3, F2: 42}
        _ = v24
    }
    var v25 int = 127
    fmt.Println(v25, v2)
    for fi26 := 0; fi26 < 3; fi26++ {
        var v27 uint = 448
        var v28 string = "\n\t"
        var v29 float64 = 0.0
        _ = fi26
        _ = v27
        _ = v28
        _ = v29
    }
    var v30 int = (v25 << ((18446744073709551615) & 63))
    var v31 int8 = 127
    fmt.Println(v30, v19)
    switch v30 {
    case 9223372036854775807:
        for fi32 := 0; fi32 < 15; fi32++ {
            var v33 uint8 = (v2 | v2)
            _ = fi32
            _ = v33
        }
    case -2147483648:
        for fi34 := 0; fi34 < 12; fi34++ {
            v35 := GenS3{F0: 2147483647}
            v36 := GenS1{F0: v25, F1: (^(v3)), F2: 197}
            v37 := GenS2{F0: "foo bar"}
            _ = fi34
            _ = v35
            _ = v36
            _ = v37
        }
    case 103:
        for fi38 := 0; fi38 < 8; fi38++ {
            var v39 float32 = 0.0
            if (610 >= 1) {
                v39 += v39
            } else {
                if false {
                    for fi40 := 0; fi40 < 14; fi40++ {
                        switch "hello" {
                        case "hello":
                            var v41 uint32 = 36
                            _ = v41
                        }
                        _ = fi40
                    }
                    var v42 uint16 = 245
                    _ = len(v1)
                    _ = v42
                }
                for fi43 := 0; fi43 < 7; fi43++ {
                    v19 = v19
                    _ = fi43
                }
            }
            _ = fi38
            _ = v39
        }
    default:
        for fi44 := 0; fi44 < 4; fi44++ {
            v31 = (v31 / (((v31 % ((v31) | 1))) | 1))
            if len(v1) > 0 {
                v1[0] = ((v19 + ("foo bar" + v19)) + v20)
            }
            _ = fi44
        }
    }
    var v45 uint8 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v45
}

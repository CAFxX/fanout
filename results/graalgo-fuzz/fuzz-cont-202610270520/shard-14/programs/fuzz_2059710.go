package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint, p1 uint64, p2 int16) uint16 {
    v5 := []int{}
    p1 += (p1 | p1)
    fmt.Println(p0, p0)
    switch 4294967295 {
    case 32768:
        for fi6 := 0; fi6 < 3; fi6++ {
            for fi7 := 0; fi7 < 8; fi7++ {
                var v8 int8 = 0
                switch "hello" {
                case "":
                    fmt.Println(p0)
                case "foo bar":
                    v5 = append(v5, fi6)
                default:
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(fi6, p0, fi7)
                    case "":
                        v9 := GenS1{F0: p1}
                        _ = v9
                    case "hello":
                        var v10 bool = false
                        _ = v10
                    default:
                        if true {
                            p0 = p0
                        }
                    }
                    fmt.Println(p0, v8, fi6)
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi7
                _ = v8
            }
            fmt.Println(fi6)
            _ = fi6
        }
    case 338:
        switch ("foo bar" + "a") {
        case "hello":
            v11 := GenS1{F0: (p1 % ((p1) | 1))}
            _ = v11
        default:
            for fi12 := 0; fi12 < 2; fi12++ {
                if (fi12 <= 9223372036854775807) {
                    switch fi12 {
                    case -869:
                        var v13 int = fi12
                        _ = v13
                    case -1:
                        var v14 int = (fi12 & fi12)
                        _ = v14
                    }
                    if len(v5) > 0 {
                        v5[0] = 1
                    }
                    var v15 int16 = 1
                    _ = v15
                }
                _ = fi12
            }
        }
    case 98:
        var v16 int8 = -128
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 391
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = 0
    v3 := make(map[string]uint32)
    v4 := []string{}
    if len(v4) > 0 {
        v4[0] = "a"
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        switch "foo bar" {
        case "hello":
            if (-12 > 1) {
                v18 := GenS2{F0: "foo bar"}
                var v19 int64 = -87
                _ = v18
                _ = v19
            } else {
                v4 = append(v4, "hello")
                switch (fi17 | (fi17 + 4294967295)) {
                case 4294967295:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case 32767:
                    var v20 uint32 = 4294967295
                    _ = v20
                case -2147483648:
                    v2 = (v2 + (v2 - (^(v2))))
                }
                v21 := GenS1{F0: 201}
                _ = v21
            }
        case "":
            switch 2147483647 {
            case -9223372036854775808:
                v1[2] = (fi17 % (((fi17 & fi17)) | 1))
            }
        case "a":
            v4 = append(v4, "hello")
        }
        fmt.Println(v2, fi17)
        v3[""] = 4294967295
        _ = fi17
    }
    v2 = v2
    switch 221 {
    case 128:
        v22 := GenS3{F0: 887}
        _ = v22
    default:
        v23 := GenS3{F0: 649}
        _ = v23
    }
    v24 := GenS3{F0: 352}
    var v25 uint32 = 2147483647
    v25 = v25
    var v26 uint64 = 0
    fmt.Println(v26, v25, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}

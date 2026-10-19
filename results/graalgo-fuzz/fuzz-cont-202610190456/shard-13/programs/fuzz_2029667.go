package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 int }
func genFn0(p0 uint, p1 int8, p2 uint64) bool {
    var v4 string = ((("a" + "") + "foo bar") + "foo bar")
    p2 = (p2 & 2)
    if (!true) {
        p1 = p1
    } else {
        v5 := []int{32767, 1, 32767}
        _ = v5
    }
    p0 += p0
    var v6 uint64 = 325
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    return (p0 < (^(p0)))
}
func genFn1(p0 uint64, p1 int16, p2 int8) int32 {
    if true {
        var v7 uint64 = p0
        fmt.Println(p0, v7)
        _ = v7
    }
    var v8 uint = 666
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 1
}

func main() {
    var v1 uint64 = 1
    v2 := []uint32{}
    var v3 uint16 = 2
    v1 += 4294967295
    v3 += (v3 + v3)
    switch -2147483648 {
    case -2147483648:
        var v9 bool = true
        _ = v9
    case 127:
        for fi10 := 0; fi10 < 10; fi10++ {
            v3 += v3
            if len(v2) > 0 {
                v2[0] = uint32(v1)
            }
            switch 128 {
            case 0:
                for fi11 := 0; fi11 < 2; fi11++ {
                    if (!false) {
                        _ = len(v2)
                        v3 += (v3 / (((v3 + v3)) | 1))
                        var v12 string = ("" + ("foo bar" + ("\n\t" + ("\n\t" + ""))))
                        _ = v12
                    }
                    if ((v3 >> ((911) & 63)) == v3) {
                        var v13 uint16 = 65535
                        _ = v13
                    } else {
                        if (((-2.25 - (0.0 / (1.5 / 1e10))) / 1.0) <= 1e10) {
                            v3 = 65535
                            v1 += (v1 * v1)
                        } else {
                            v14 := GenS1{F0: (v1 != (^(v1)))}
                            if false {
                                var v15 int8 = 42
                                _ = v15
                            } else {
                                var v16 float32 = 0.1
                                fmt.Println(v3)
                                _ = v16
                            }
                            v17 := GenS1{F0: true}
                            _ = v14
                            _ = v17
                        }
                    }
                    _ = fi11
                }
            default:
                if true {
                    if len(v2) > 0 {
                        v2[0] = 905
                    }
                }
            }
            _ = fi10
        }
    case 4294967295:
        v1 -= (^(v1))
    }
    var v18 bool = false
    fmt.Println(v18, v1)
    switch "\n\t" {
    case "hello":
        for fi19 := 0; fi19 < 15; fi19++ {
            for fi20 := 0; fi20 < 13; fi20++ {
                fmt.Println(v3)
                if len(v2) > 0 {
                    v2[0] = 42
                }
                switch "a" {
                case "":
                    var v21 int8 = 1
                    _ = v21
                default:
                    v1 = (v1 / (((v1 + (v1 / ((175) | 1)))) | 1))
                }
                fmt.Println(fi19)
                v1 = (v1 + (^(v1)))
                _ = fi20
            }
            _ = fi19
        }
    case "a":
        _ = len(v2)
    case "foo bar":
        v3 += v3
    }
    v22 := GenS1{F0: (!v18)}
    v2 = append(v2, 802)
    if v18 {
        for fi23 := 0; fi23 < 9; fi23++ {
            switch fi23 {
            case -2147483648:
                switch "" {
                case "hello":
                    v18 = v18
                }
            case 128:
                v24 := GenS2{F0: (^(v1))}
                _ = v24
            default:
                var v25 uint16 = v3
                _ = v25
            }
            if len(v2) > 0 {
                v2[0] = 735
            }
            v26 := GenS2{F0: v1}
            _ = fi23
            _ = v26
        }
        v27 := GenS1{F0: v18}
        _ = v27
    }
    fmt.Println(v1, v1)
    if v18 {
        var v28 uint32 = 482
        v29 := GenS3{F0: 118, F1: 4294967295}
        _ = v28
        _ = v29
    } else {
        switch 42 {
        case 1:
            v3 = v3
        case 4294967295:
            v30 := GenS1{F0: true}
            _ = v30
        default:
            var v31 int64 = 9223372036854775807
            _ = v31
        }
    }
    fmt.Println(v1)
    switch ("hello" + "a") {
    case "foo bar":
        fmt.Println(v18, v3, v1)
    }
    v2 = append(v2, 4294967295)
    var v32 uint = 45
    v33 := GenS2{F0: (^(v1))}
    fmt.Println(v3, v18, v32)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v32
    _ = v33
}

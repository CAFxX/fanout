package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int16; F1 int8 }
func genFn0(p0 int32, p1 int64, p2 uint) int {
    p0 = p0
    p0 = (^(p0))
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            p2 += (^(p2))
            _ = fi8
        }
        switch (fi7 & (fi7 << ((p2) & 63))) {
        case 32768:
            if (1.5 != 1e10) {
                var v9 float64 = 0.0
                for fi10 := 0; fi10 < 17; fi10++ {
                    v11 := GenS1{F0: 42, F1: fi10}
                    _ = fi10
                    _ = v11
                }
                _ = v9
            }
        case 953:
            p0 += 2147483647
        case -9223372036854775808:
            switch "a" {
            case "\n\t":
                v12 := GenS2{F0: p2, F1: -78}
                _ = v12
            }
        default:
            for fi13 := 0; fi13 < 18; fi13++ {
                v14 := []int{(fi7 | (fi7 - 2147483647)), (fi7 | (fi7 | fi13)), (fi13 ^ 32768), (fi7 | (fi7 & fi7))}
                _ = fi13
                _ = v14
            }
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    return -708
}
func genFn1(p0 bool, p1 uint16) uint8 {
    v15 := []int{int(p1), genFn0(2147483647, 18, 625), -958, 0}
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 float32 = (((1e10 - float32(float64(3.14159))) - 1.5) * float32(1e10))
        switch genFn0(-813, 9223372036854775807, 255) {
        case -229:
            for fi18 := 0; fi18 < 12; fi18++ {
                var v19 int16 = -1
                _ = fi18
                _ = v19
            }
        case 1:
            p1 -= 42
        case 4294967295:
            var v20 uint32 = 65535
            _ = v20
        }
        v21 := GenS1{F0: 0, F1: genFn0(0, 79, 790)}
        _ = fi16
        _ = v17
        _ = v21
    }
    p1 += (p1 % ((441) | 1))
    p0 = (!((1e10 - (0.1 + (-2.25 / 1e10))) <= (1e10 - 1e10)))
    _ = p0
    _ = p1
    _ = v15
    return 1
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 42
    v3 := make(map[uint32]uint32)
    v4 := []string{("a" + "a")}
    var v5 int = 127
    var v6 string = "hello"
    v2 -= -32768
    v22 := GenS1{F0: 2, F1: (v5 >> ((370) & 63))}
    var v23 uint16 = 42
    v23 = 2
    v24 := v1[v6]
    v23 = v23
    v23 = (v23 - v23)
    var v25 int32 = 42
    if true {
        switch v6 {
        case "foo bar":
            switch "hello" {
            case "\n\t":
                switch v6 {
                case "\n\t":
                    fmt.Println(v6)
                case "a":
                    for fi26 := 0; fi26 < 5; fi26++ {
                        if len(v4) > 0 {
                            v4[0] = ""
                        }
                        _ = fi26
                    }
                default:
                    for fi27 := 0; fi27 < 7; fi27++ {
                        var v28 bool = true
                        var v29 uint32 = 906
                        var v30 uint8 = 127
                        _ = fi27
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                }
            case "":
                delete(v1, "foo bar")
            case "foo bar":
                v31 := GenS3{F0: -1, F1: 100}
                _ = v31
            default:
                if (9223372036854775807 >= 1) {
                    switch v5 {
                    case 32768:
                        v6 = (v24 + v6)
                    case 526:
                        v25 = (^(v25))
                    default:
                        var v32 uint16 = 737
                        _ = v32
                    }
                    var v33 int16 = -641
                    _ = v33
                } else {
                    for fi34 := 0; fi34 < 1; fi34++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi34
                    }
                    delete(v3, 2)
                    var v35 uint64 = 42
                    _ = v35
                }
            }
        case "hello":
            v6 = "a"
        case "\n\t":
            fmt.Println(v23, v25)
        default:
            if true {
                var v36 uint8 = 200
                v37 := GenS3{F0: -765, F1: -128}
                switch v6 {
                case "":
                    if false {
                        v38 := GenS1{F0: 786, F1: 4294967295}
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v38
                    }
                }
                _ = v36
                _ = v37
            }
        }
        for fi39 := 0; fi39 < 7; fi39++ {
            var v40 float64 = ((0.0 + 0.1) - 0.1)
            v23 -= 32767
            v24 = v24
            _ = fi39
            _ = v40
        }
        var v41 float32 = 0.1
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int16 }
func genFn0(p0 uint16, p1 int16, p2 uint8) string {
    v5 := GenS1{F0: p0, F1: -158}
    if (18446744073709551615 < 65535) {
        v6 := []int{-1}
        v6 = append(v6, -611)
        _ = v6
    } else {
        v7 := []int{32767}
        _ = len(v7)
        _ = v7
    }
    if true {
        v8 := []int{32768, 32767, 389}
        var v9 uint32 = 42
        for fi10 := 0; fi10 < 19; fi10++ {
            for fi11 := 0; fi11 < 7; fi11++ {
                _ = len(v8)
                if false {
                    switch "foo bar" {
                    case "\n\t":
                        var v12 int8 = 24
                        _ = v12
                    case "foo bar":
                        v13 := GenS1{F0: (p0 * (p0 | (p0 | (p0 | (p0 % ((650) | 1)))))), F1: 9223372036854775807}
                        var v14 uint8 = uint8(p0)
                        var v15 float64 = 3.14159
                        _ = v13
                        _ = v14
                        _ = v15
                    default:
                        var v16 uint = 149
                        _ = v16
                    }
                }
                fmt.Println(p1, fi10)
                p2 += p2
                _ = fi11
            }
            _ = fi10
        }
        _ = v8
        _ = v9
    } else {
        if false {
            switch 32767 {
            case -1:
                switch 127 {
                case -2147483648:
                    var v17 uint16 = p0
                    _ = v17
                case 2147483647:
                    switch 128 {
                    case 196:
                        for fi18 := 0; fi18 < 20; fi18++ {
                            var v19 int16 = p1
                            _ = fi18
                            _ = v19
                        }
                    case 127:
                        v20 := []int{}
                        _ = v20
                    default:
                        for fi21 := 0; fi21 < 17; fi21++ {
                            var v22 float64 = 0.1
                            var v23 int32 = 1
                            var v24 int16 = (-(p1))
                            _ = fi21
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                    }
                }
            case 32767:
                v25 := []int{}
                _ = v25
            case 4294967295:
                v26 := []int{333, 32767, 42, -698}
                _ = v26
            }
            for fi27 := 0; fi27 < 10; fi27++ {
                if (9223372036854775807 == 340) {
                    v28 := []int{fi27}
                    _ = len(v28)
                    _ = v28
                } else {
                    if true {
                        switch 795 {
                        case -294:
                            fmt.Println(p2)
                        case -310:
                            v29 := []int{fi27, fi27, -1, (-(fi27))}
                            _ = v29
                        default:
                            var v30 bool = (!(232 >= p2))
                            _ = v30
                        }
                        p0 += 65535
                        v31 := GenS1{F0: (p0 / ((952) | 1)), F1: -1}
                        _ = v31
                    }
                    p0 -= uint16(p2)
                    v32 := GenS1{F0: p0, F1: -1}
                    _ = v32
                }
                p2 -= 0
                p0 -= (^(p0))
                _ = fi27
            }
        } else {
            var v33 int8 = 100
            fmt.Println(v33)
            _ = v33
        }
    }
    var v34 int64 = -259
    v35 := []int{2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v34
    _ = v35
    _ = v5
    return ("a" + ("hello" + ("" + "")))
}

func main() {
    var v1 uint16 = 0
    var v2 uint = 18446744073709551615
    v3 := make(map[uint32]string)
    var v4 int = 1
    var v36 uint64 = 148
    v37 := GenS1{F0: v1, F1: 1}
    fmt.Println(v1)
    v38 := GenS2{F0: 68}
    fmt.Println(v1, v1)
    v39 := []int{}
    var v40 bool = false
    fmt.Println(v40)
    switch v4 {
    case 2147483647:
        if len(v39) >= 2 {
            v39 = v39[1:len(v39)]
        }
    default:
        switch 32767 {
        case 32767:
            v41 := GenS3{F0: (^(v4)), F1: 267}
            _ = v41
        case -172:
            if len(v39) > 0 {
                v39[0] = 42
            }
        default:
            v42 := GenS3{F0: 4294967295, F1: 0}
            _ = v42
        }
    }
    var v43 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v43
}

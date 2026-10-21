package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 int, p1 int16) int32 {
    v5 := []int{p0, (p0 << ((803) & 63)), (^(p0)), p0}
    var v6 uint = 1
    var v7 uint = 65535
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1() uint16 {
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 int64 = 326
        v9 = (v9 + 0)
        _ = fi8
        _ = v9
    }
    if true {
        for fi10 := 0; fi10 < 8; fi10++ {
            fmt.Println(fi10)
            _ = fi10
        }
        var v11 uint64 = 4294967295
        switch 32768 {
        case 4294967295:
            fmt.Println(v11)
        case 32768:
            v11 = 544
        case 127:
            fmt.Println(v11)
        default:
            v11 = (^(v11))
        }
        _ = v11
    }
    return 1
}

func main() {
    var v1 uint = 4294967295
    v2 := make(map[int]uint32)
    var v3 int32 = -2147483648
    v4 := []uint32{}
    var v12 bool = (false || false)
    for fi13 := 0; fi13 < 18; fi13++ {
        v14 := GenS1{F0: int8(fi13), F1: 0, F2: -61}
        fmt.Println(fi13, fi13, v3)
        _ = fi13
        _ = v14
    }
    switch 32768 {
    case 58:
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 uint = v1
            if v12 {
                switch ((("hello" + "a") + "a") + "foo bar") {
                case "":
                    var v17 float64 = -2.25
                    _ = v17
                case "hello":
                    var v18 string = "hello"
                    _ = v18
                default:
                    v19 := v2[-217]
                    _ = v19
                }
                v16 = 2
            } else {
                _ = len(v4)
            }
            _ = fi15
            _ = v16
        }
    case 541:
        for fi20 := 0; fi20 < 2; fi20++ {
            v21 := GenS2{F0: v1}
            _ = fi20
            _ = v21
        }
    }
    v3 = v3
    if (0.1 < float32(((1.5 * 0.1) * 1.5))) {
        for fi22 := 0; fi22 < 14; fi22++ {
            v23 := GenS2{F0: (v1 << ((42) & 63))}
            _ = fi22
            _ = v23
        }
        var v24 float32 = 1.5
        v12 = v12
        _ = v24
    }
    switch "a" {
    case "hello":
        v4 = append(v4, 448)
    }
    v25 := GenS3{F0: 32767}
    _ = len(v4)
    var v26 int8 = -1
    if (3.14159 != (-2.25 / ((3.14159 + 0.0) / ((0.1 - -2.25) / -2.25)))) {
        var v27 int32 = v3
        v4 = append(v4, 42)
        _ = v27
    } else {
        var v28 int64 = 1
        _ = v28
    }
    var v29 int32 = v3
    var v30 uint = (v1 << ((v1) & 63))
    v3 = (v29 * (v29 - v3))
    for fi31 := 0; fi31 < 6; fi31++ {
        if v12 {
            fmt.Println(fi31, v30)
        } else {
            if false {
                var v32 int8 = v26
                _ = v32
            } else {
                switch "" {
                case "foo bar":
                    v33 := v2[(fi31 - (fi31 + int(v30)))]
                    _ = v33
                case "a":
                    for fi34 := 0; fi34 < 2; fi34++ {
                        if true {
                            v35 := GenS2{F0: (^(v30))}
                            fmt.Println(fi31)
                            _ = v35
                        }
                        v12 = false
                        var v36 int = (fi34 - fi34)
                        _ = fi34
                        _ = v36
                    }
                case "\n\t":
                    for fi37 := 0; fi37 < 15; fi37++ {
                        v26 -= v26
                        _ = fi37
                    }
                default:
                    v3 -= (v3 % ((2147483647) | 1))
                }
                switch fi31 {
                case 421:
                    _ = len(v2)
                case 127:
                    v38 := GenS4{F0: 42}
                    _ = v38
                case 4294967295:
                    var v39 int64 = -1
                    _ = v39
                }
                v12 = true
            }
            v12 = false
        }
        var v40 uint64 = 659
        v41 := GenS5{F0: 2}
        _ = fi31
        _ = v40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}

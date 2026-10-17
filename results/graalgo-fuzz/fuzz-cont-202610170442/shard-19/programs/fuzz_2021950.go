package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint, p1 int64, p2 int64) uint {
    var v4 uint64 = 770
    switch ("foo bar" + ("" + ("hello" + "\n\t"))) {
    case "":
        for fi5 := 0; fi5 < 2; fi5++ {
            v6 := []int{fi5, (fi5 & (fi5 + fi5)), 32767, 177}
            _ = fi5
            _ = v6
        }
    case "a":
        v7 := GenS1{F0: p0}
        _ = v7
    case "foo bar":
        v8 := []int{-2147483648, 0}
        _ = v8
    }
    var v9 int64 = 0
    v10 := GenS1{F0: p0}
    fmt.Println(v9, v4)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v4
    _ = v9
    return (p0 << (((p0 / ((p0) | 1))) & 63))
}
func genFn1(p0 uint16, p1 int8) int64 {
    for fi11 := 0; fi11 < 4; fi11++ {
        v12 := []int{-613, 339, fi11}
        if false {
            var v13 uint32 = 890
            p1 = -1
            _ = v13
        } else {
            if false {
                switch fi11 {
                case 127:
                    switch fi11 {
                    case 128:
                        for fi14 := 0; fi14 < 9; fi14++ {
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                            p1 += (p1 * 24)
                            _ = fi14
                        }
                    case -9223372036854775808:
                        p0 -= p0
                    case 260:
                        if (!(!true)) {
                            v15 := GenS1{F0: 87}
                            fmt.Println(fi11, p1)
                            p0 = p0
                            _ = v15
                        } else {
                            v12 = append(v12, int(p1))
                            p1 -= (p1 % (((p1 + (p1 << ((255) & 63)))) | 1))
                        }
                    default:
                        switch "\n\t" {
                        case "\n\t":
                            var v16 bool = false
                            _ = v16
                        }
                    }
                case 768:
                    var v17 float64 = 3.14159
                    _ = v17
                case -2147483648:
                    for fi18 := 0; fi18 < 7; fi18++ {
                        p1 += 42
                        _ = fi18
                    }
                default:
                    _ = len(v12)
                }
                p0 += (p0 >> ((4294967295) & 63))
                v19 := GenS1{F0: 42}
                _ = v19
            } else {
                if true {
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                    fmt.Println(p1, fi11, p0)
                    fmt.Println(p0)
                }
                if false {
                    var v20 uint = 219
                    if false {
                        v21 := GenS1{F0: (v20 * (v20 % (((v20 ^ (^(v20)))) | 1)))}
                        p0 += p0
                        _ = v21
                    } else {
                        p1 -= (^(p1))
                        var v22 int32 = 1
                        _ = v22
                    }
                    _ = v20
                }
                if false {
                    _ = len(v12)
                    for fi23 := 0; fi23 < 1; fi23++ {
                        switch "hello" {
                        case "hello":
                            if len(v12) > 0 {
                                v12[0] = (^(fi11))
                            }
                        default:
                            v24 := GenS2{F0: 38}
                            _ = v24
                        }
                        v25 := GenS3{F0: 0}
                        _ = fi23
                        _ = v25
                    }
                    var v26 float64 = 0.1
                    _ = v26
                } else {
                    var v27 uint = 607
                    fmt.Println(p0, p1, p1)
                    _ = v27
                }
            }
            var v28 uint16 = (^(p0))
            var v29 int64 = 656
            _ = v28
            _ = v29
        }
        _ = fi11
        _ = v12
    }
    p1 = p1
    _ = p0
    _ = p1
    return -1
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[string]string)
    var v3 int32 = -2147483648
    v30 := GenS3{F0: v3}
    v3 += v3
    v31 := GenS2{F0: 498}
    fmt.Println(v1, v1)
    fmt.Println(v1, v3)
    if (3.14159 <= -2.25) {
        v1 = v1
    } else {
        var v32 uint = 1
        _ = v32
    }
    var v33 int = 9223372036854775807
    switch "foo bar" {
    case "":
        delete(v2, "\n\t")
    case "\n\t":
        var v34 uint32 = 569
        _ = v34
    case "hello":
        fmt.Println(v33, v1)
    default:
        v35 := v2["a"]
        _ = v35
    }
    v36 := []int{32768, v33, (^(v33))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v36:", len(v36), v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v36
}

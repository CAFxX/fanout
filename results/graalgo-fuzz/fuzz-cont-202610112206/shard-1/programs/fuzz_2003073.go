package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int64) uint {
    switch 32767 {
    case 127:
        for fi7 := 0; fi7 < 2; fi7++ {
            fmt.Println(fi7, p0)
            _ = fi7
        }
    }
    if true {
        if true {
            fmt.Println(p0)
        } else {
            v8 := []int{-2147483648, 127, -1, -2147483648}
            if (!(!(65 >= 1))) {
                for fi9 := 0; fi9 < 3; fi9++ {
                    for fi10 := 0; fi10 < 7; fi10++ {
                        v11 := GenS1{F0: 104}
                        var v12 int = (-(fi9))
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 7; fi13++ {
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        switch fi13 {
                        case -2147483648:
                            var v14 int = fi13
                            _ = v14
                        case 4294967295:
                            p0 += p0
                            if true {
                                var v15 uint16 = 65535
                                p0 = (^(p0))
                                v15 = v15
                                p0 -= int64(v15)
                                _ = v15
                            } else {
                                p0 += (-(p0))
                            }
                        }
                        var v16 uint16 = 0
                        _ = fi13
                        _ = v16
                    }
                    fmt.Println(p0, fi9)
                    _ = fi9
                }
                var v17 int8 = 74
                v17 -= v17
                _ = v17
            } else {
                fmt.Println(p0)
                v8 = append(v8, 128)
                switch -55 {
                case 32768:
                    v18 := GenS2{F0: 42}
                    _ = v18
                case -285:
                    if len(v8) > 0 {
                        v8[0] = 1
                    }
                case 32767:
                    _ = len(v8)
                default:
                    var v19 float64 = 0.0
                    _ = v19
                }
            }
            if false {
                var v20 int32 = 1
                p0 -= 1
                v20 = (v20 & -1)
                _ = v20
            } else {
                switch "\n\t" {
                case "a":
                    if false {
                        var v21 uint64 = 1
                        fmt.Println(v21, p0)
                        var v22 int16 = 1
                        _ = v21
                        _ = v22
                    }
                default:
                    p0 = p0
                }
            }
            _ = v8
        }
    } else {
        fmt.Println(p0)
    }
    p0 += p0
    v23 := []int{-2147483648}
    _ = p0
    _ = v23
    return 924
}
func genFn1(p0 int64, p1 bool) uint8 {
    fmt.Println(p1, p0)
    if p1 {
        fmt.Println(p1, p1)
        var v24 bool = (!((p1 && (0.0 < float32(1e10))) || p1))
        _ = v24
    }
    v25 := GenS1{F0: 18446744073709551615}
    _ = p0
    _ = p1
    _ = v25
    return 255
}

func main() {
    var v1 int8 = 13
    var v2 uint64 = 148
    v3 := make(map[uint32]int)
    v4 := make(map[string]uint32)
    v5 := []string{"hello"}
    v6 := make(map[int]string)
    delete(v4, ((("\n\t" + "a") + (((("\n\t" + "\n\t") + "hello") + ("hello" + "")) + "")) + "foo bar"))
    var v26 uint = 595
    v26 = (v26 + genFn0(-9223372036854775808))
    if false {
        v27 := GenS1{F0: 0}
        _ = v27
    }
    v28 := GenS2{F0: 0}
    var v29 int = -422
    fmt.Println(v1, v26)
    var v30 uint32 = 692
    if false {
        delete(v6, v29)
    } else {
        var v31 int8 = v1
        var v32 uint16 = 1
        _ = v31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}

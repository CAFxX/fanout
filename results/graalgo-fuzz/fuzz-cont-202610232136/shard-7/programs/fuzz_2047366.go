package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint }
func genFn0(p0 uint8, p1 uint16, p2 string) uint32 {
    if false {
        fmt.Println(p1, p1, p1)
    } else {
        var v6 int = 32767
        _ = v6
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        switch p2 {
        case "a":
            v8 := []int{fi7, (fi7 << ((134) & 63))}
            _ = v8
        case "\n\t":
            var v9 int32 = 324
            _ = v9
        }
        for fi10 := 0; fi10 < 3; fi10++ {
            fmt.Println(fi7)
            p0 += (p0 << ((2) & 63))
            _ = fi10
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 2
    var v4 string = ("foo bar" + "a")
    var v5 int16 = 42
    v11 := []int{-60, -324, 1, -926}
    var v12 string = (((v4 + ("hello" + v4)) + "\n\t") + ("a" + v4))
    for fi13 := 0; fi13 < 2; fi13++ {
        v14 := GenS1{F0: 0, F1: 1}
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        var v15 int16 = (v5 << ((42) & 63))
        _ = fi13
        _ = v14
        _ = v15
    }
    switch 4294967295 {
    case 958:
        v4 = v4
    default:
        if (!((!false) && true)) {
            if (770 >= 464) {
                if (v5 < (v5 / ((-32768) | 1))) {
                    for fi16 := 0; fi16 < 6; fi16++ {
                        if (0 > -1) {
                            if false {
                                for fi17 := 0; fi17 < 4; fi17++ {
                                    v3 += (v3 / (((v3 & (v3 >> ((2) & 63)))) | 1))
                                    _ = len(v11)
                                    _ = fi17
                                }
                                for fi18 := 0; fi18 < 7; fi18++ {
                                    v19 := GenS1{F0: 42, F1: 4294967295}
                                    _ = len(v2)
                                    _ = fi18
                                    _ = v19
                                }
                            }
                            var v20 int8 = 42
                            if false {
                                var v21 int32 = -705
                                v3 = 255
                                _ = v21
                            } else {
                                var v22 int64 = 42
                                v5 = 1
                                fmt.Println(v20, fi16)
                                _ = v22
                            }
                            _ = v20
                        } else {
                            v23 := GenS1{F0: 100, F1: 42}
                            v5 -= v5
                            v5 += (v5 * 32767)
                            _ = v23
                        }
                        _ = fi16
                    }
                    if (!true) {
                        if len(v11) > 0 {
                            v11[0] = 42
                        }
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                    } else {
                        delete(v1, genFn0(v3, 255, ((("" + "a") + ("a" + "\n\t")) + v4)))
                        v5 = v5
                        var v24 float32 = 0.1
                        _ = v24
                    }
                    var v25 bool = (!true)
                    _ = v25
                } else {
                    v26 := GenS1{F0: 100, F1: 410}
                    var v27 int = -1
                    _ = v26
                    _ = v27
                }
                v3 += v3
                var v28 bool = true
                _ = v28
            }
        }
    }
    var v29 int64 = 9223372036854775807
    v30 := GenS2{F0: 0, F1: 0, F2: 2}
    fmt.Println(v4)
    v4 = "hello"
    switch -1 {
    case 32768:
        fmt.Println(v3, v12, v29)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}

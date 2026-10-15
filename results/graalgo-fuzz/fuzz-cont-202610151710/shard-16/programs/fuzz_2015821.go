package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int8; F1 bool }
func genFn0() int32 {
    var v6 uint8 = 255
    fmt.Println(v6)
    _ = v6
    return 42
}

func main() {
    v1 := []string{}
    v2 := []int{-2147483648, 1}
    v3 := []int{-2147483648, 0, 32767}
    var v4 string = ("\n\t" + ("a" + "\n\t"))
    v5 := []uint32{358, 381, 741, 187}
    var v7 int16 = 1
    _ = len(v2)
    var v8 uint = 222
    if len(v1) > 0 {
        v1[0] = ""
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        switch "" {
        case "":
            var v10 uint16 = 1
            _ = v10
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = fi9
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        fmt.Println(v4, v8, v7)
        v8 -= 255
        _ = fi11
    }
    if false {
        fmt.Println(v4)
        if true {
            var v12 string = v4
            for fi13 := 0; fi13 < 3; fi13++ {
                for fi14 := 0; fi14 < 1; fi14++ {
                    if false {
                        if true {
                            var v15 int32 = genFn0()
                            _ = v15
                        } else {
                            var v16 float32 = 0.0
                            _ = v16
                        }
                        var v17 uint8 = 42
                        _ = v17
                    } else {
                        v12 = (v12 + (v12 + v12))
                        v2 = append(v2, fi13)
                        v1 = append(v1, v12)
                    }
                    fmt.Println(v4, v8, fi13)
                    var v18 int32 = 1
                    _ = fi14
                    _ = v18
                }
                v19 := GenS1{F0: 42, F1: 0}
                _ = fi13
                _ = v19
            }
            _ = v12
        }
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        if true {
            if true {
                for fi21 := 0; fi21 < 9; fi21++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi21
                }
            } else {
                for fi22 := 0; fi22 < 11; fi22++ {
                    if len(v5) > 0 {
                        v5[0] = 65535
                    }
                    v8 = (v8 ^ (v8 & v8))
                    fmt.Println(v8, v8)
                    _ = fi22
                }
                for fi23 := 0; fi23 < 16; fi23++ {
                    v24 := GenS2{F0: -3, F1: (!(-119 == 1))}
                    v7 += v7
                    _ = fi23
                    _ = v24
                }
            }
        } else {
            v25 := GenS2{F0: -12, F1: (!false)}
            v7 = v7
            _ = v25
        }
        _ = fi20
    }
    v8 = v8
    v4 = v4
    for fi26 := 0; fi26 < 7; fi26++ {
        var v27 float32 = float32(0.0)
        switch v4 {
        case "foo bar":
            _ = len(v5)
        }
        var v28 string = "foo bar"
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 string = (v4 + ((v4 + v4) + ""))
    v8 = v8
    for fi30 := 0; fi30 < 6; fi30++ {
        v31 := GenS1{F0: 2, F1: -109}
        fmt.Println(v29, fi30)
        for fi32 := 0; fi32 < 20; fi32++ {
            v7 += (v7 + (v7 * v7))
            fmt.Println(v4)
            v33 := GenS1{F0: 255, F1: -21}
            _ = fi32
            _ = v33
        }
        _ = fi30
        _ = v31
    }
    for fi34 := 0; fi34 < 6; fi34++ {
        var v35 int32 = -2147483648
        if len(v5) > 0 {
            v5[0] = 2
        }
        _ = fi34
        _ = v35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}

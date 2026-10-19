package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 int32; F1 uint; F2 uint32 }
type GenS5 struct { F0 uint; F1 uint; F2 bool }
func genFn0() uint8 {
    if false {
        for fi7 := 0; fi7 < 9; fi7++ {
            fmt.Println(fi7)
            for fi8 := 0; fi8 < 12; fi8++ {
                fmt.Println(fi8)
                _ = fi8
            }
            _ = fi7
        }
        var v9 int8 = -1
        _ = v9
    }
    switch "foo bar" {
    case "\n\t":
        v10 := []int{}
        _ = v10
    case "a":
        var v11 uint64 = 42
        _ = v11
    case "hello":
        var v12 int64 = 42
        _ = v12
    default:
        for fi13 := 0; fi13 < 4; fi13++ {
            if true {
                var v14 float32 = 3.14159
                if true {
                    v14 += v14
                } else {
                    if false {
                        v15 := []int{(-(fi13))}
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        fmt.Println(fi13)
                        _ = v15
                    } else {
                        fmt.Println(v14, v14)
                        v14 = v14
                    }
                }
                fmt.Println(v14)
                _ = v14
            }
            _ = fi13
        }
    }
    var v16 int64 = 864
    if (true || false) {
        v16 = v16
        v17 := GenS1{F0: (v16 * v16)}
        _ = v17
    } else {
        for fi18 := 0; fi18 < 9; fi18++ {
            v19 := []int{-841}
            _ = fi18
            _ = v19
        }
        if (true && (false || ((-2147483648 != -821) || (!((128 > 0) && false))))) {
            if ((32767 != -2147483648) || (true || false)) {
                v20 := GenS2{F0: 0, F1: 1}
                _ = v20
            }
            if false {
                if true {
                    var v21 bool = false
                    v22 := []int{-734, -2147483648, -82, 1}
                    _ = v21
                    _ = v22
                } else {
                    var v23 uint = 1
                    _ = v23
                }
            } else {
                switch -9223372036854775808 {
                case 2147483647:
                    var v24 int8 = 127
                    _ = v24
                default:
                    var v25 float32 = 3.14159
                    _ = v25
                }
            }
            var v26 float32 = ((3.14159 - -2.25) + (float32(-2.25) + 3.14159))
            _ = v26
        }
    }
    v16 -= v16
    _ = v16
    return 1
}
func genFn1(p0 string, p1 uint32) int16 {
    p0 = ((p0 + (p0 + "a")) + ((p0 + (((p0 + p0) + p0) + "a")) + p0))
    fmt.Println(p1, p1)
    switch p0 {
    case "\n\t":
        p0 = "\n\t"
    case "foo bar":
        v27 := GenS3{F0: 760, F1: 0}
        _ = v27
    case "a":
        v28 := []int{0, 128, -1}
        _ = v28
    default:
        var v29 uint = 0
        _ = v29
    }
    v30 := []int{1, 42, 42, 32767}
    if len(v30) >= 2 {
        v30 = v30[1:len(v30)]
    }
    _ = p0
    _ = p1
    _ = v30
    return 0
}

func main() {
    var v1 string = ""
    v2 := make(map[uint32]bool)
    v3 := []uint32{2147483647, 594, 11, 2147483648}
    var v4 uint32 = 2147483647
    var v5 uint16 = 295
    v6 := make(map[int]string)
    var v31 int64 = 1
    var v32 int64 = -1
    delete(v2, v4)
    v33 := GenS1{F0: (v32 ^ (v32 * (v32 >> ((748) & 63))))}
    v34 := GenS2{F0: v31, F1: 833}
    var v35 int = 127
    var v36 int64 = v31
    v37 := GenS4{F0: 0, F1: 1, F2: 2}
    switch v1 {
    case "a":
        fmt.Println(v36)
    default:
        if true {
            var v38 uint32 = v4
            _ = v38
        } else {
            fmt.Println(v31, v36)
            for fi39 := 0; fi39 < 16; fi39++ {
                v40 := GenS5{F0: 296, F1: 4294967295, F2: true}
                _ = fi39
                _ = v40
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}

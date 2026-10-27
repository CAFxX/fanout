package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int8; F1 int; F2 string }
func genFn0(p0 uint64, p1 uint64, p2 bool) bool {
    for fi5 := 0; fi5 < 12; fi5++ {
        fmt.Println(p1)
        fmt.Println(fi5, p2, p2)
        _ = fi5
    }
    switch 127 {
    case -9223372036854775808:
        if p2 {
            fmt.Println(p2, p2, p0)
            for fi6 := 0; fi6 < 9; fi6++ {
                v7 := []int{}
                var v8 int32 = -1
                _ = fi6
                _ = v7
                _ = v8
            }
            p1 = (^(p0))
        } else {
            v9 := []int{1, 1, 32768}
            _ = v9
        }
    case 127:
        v10 := []int{4294967295, 127, 128}
        _ = v10
    default:
        v11 := []int{-869}
        _ = v11
    }
    v12 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return false
}
func genFn1() int8 {
    var v13 int8 = 42
    fmt.Println(v13)
    v14 := []int{-164}
    _ = v13
    _ = v14
    return (v13 / (((v13 + v13)) | 1))
}

func main() {
    v1 := []uint32{279, 2147483647, 957, 941}
    v2 := []string{}
    v3 := []uint32{}
    var v4 float32 = (-2.25 / 0.1)
    v4 += v4
    var v15 int = -9223372036854775808
    var v16 int16 = 853
    var v17 int = v15
    switch "hello" {
    case "\n\t":
        var v18 uint8 = 128
        _ = v18
    case "":
        if true {
            for fi19 := 0; fi19 < 15; fi19++ {
                for fi20 := 0; fi20 < 4; fi20++ {
                    fmt.Println(v4, v4)
                    v15 = 1
                    var v21 int16 = 378
                    _ = fi20
                    _ = v21
                }
                v2 = append(v2, "hello")
                fmt.Println(fi19, v17)
                _ = fi19
            }
        } else {
            v22 := GenS1{F0: 42, F1: "hello"}
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        if (!true) {
            var v24 int32 = -914
            v25 := GenS1{F0: 582, F1: "a"}
            v26 := GenS2{F0: -1, F1: (v15 >> ((962) & 63)), F2: "\n\t"}
            _ = v24
            _ = v25
            _ = v26
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v27 int32 = -755
            _ = v27
        }
        var v28 int = -189
        _ = fi23
        _ = v28
    }
    if (4294967295 >= (v17 * (v15 * (v15 & (v17 + -2147483648))))) {
        fmt.Println(v17, v17, v15)
    } else {
        for fi29 := 0; fi29 < 3; fi29++ {
            v17 = (-(fi29))
            _ = fi29
        }
    }
    fmt.Println(v16, v17, v16)
    v16 = v16
    for fi30 := 0; fi30 < 16; fi30++ {
        if true {
            var v31 int8 = -101
            _ = v31
        } else {
            v17 += (v17 << ((653) & 63))
        }
        v15 = fi30
        for fi32 := 0; fi32 < 16; fi32++ {
            if false {
                if false {
                    v16 = (v16 * v16)
                    v16 += v16
                    var v33 uint16 = 32767
                    _ = v33
                }
            }
            fmt.Println(v4)
            _ = fi32
        }
        _ = fi30
    }
    if (false || false) {
        var v34 uint16 = 32767
        v15 += 147
        _ = v34
    }
    v35 := GenS2{F0: -128, F1: v17, F2: (("" + "a") + "foo bar")}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}

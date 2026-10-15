package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int16; F1 int }
func genFn0(p0 uint16, p1 int64, p2 int) int16 {
    fmt.Println(p2)
    var v6 float64 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -32768
}
func genFn1(p0 string, p1 int8) uint16 {
    var v7 int64 = -1
    v8 := []int{9223372036854775807, 32767}
    var v9 uint = 65535
    var v10 int8 = p1
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 65535
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = 0.0
    v3 := make(map[string]string)
    v4 := []int{4294967295}
    v5 := make(map[string]int)
    v2 -= v2
    var v11 string = ""
    v11 = "a"
    var v12 float64 = (((((-2.25 + float64(v2)) * float64((-2.25 * v2))) / (-2.25 * 1e10)) - 1e10) / float64(v2))
    v13 := GenS1{F0: 127, F1: 215}
    for fi14 := 0; fi14 < 14; fi14++ {
        v2 = float32(-2.25)
        _ = fi14
    }
    var v15 int32 = -781
    v11 = ((v11 + ((("\n\t" + ("hello" + ("a" + v11))) + "foo bar") + ("hello" + v11))) + v11)
    var v16 uint = 4294967295
    var v17 float64 = (0.1 - float64(v2))
    var v18 uint = 2
    var v19 int16 = 142
    switch 0 {
    case -1:
        v17 += 1.5
    case -643:
        fmt.Println(v19)
    }
    if (43 == 710) {
        if true {
            v19 += 526
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    } else {
        var v20 string = ("foo bar" + (v11 + "\n\t"))
        _ = v20
    }
    fmt.Println(v16, v16)
    v15 = v15
    if false {
        v21 := GenS2{F0: (v19 % (((v19 >> ((v18) & 63))) | 1)), F1: -278}
        v11 = v11
        fmt.Println(v15, v11, v18)
        _ = v21
    } else {
        if false {
            for fi22 := 0; fi22 < 9; fi22++ {
                v23 := GenS1{F0: uint8(v18), F1: -9223372036854775808}
                v4 = append(v4, fi22)
                for fi24 := 0; fi24 < 12; fi24++ {
                    delete(v1, 319)
                    v2 = -2.25
                    _ = fi24
                }
                _ = fi22
                _ = v23
            }
            for fi25 := 0; fi25 < 16; fi25++ {
                fmt.Println(fi25)
                var v26 int32 = (v15 + (v15 << ((v18) & 63)))
                _ = fi25
                _ = v26
            }
            if (4294967295 == 842) {
                fmt.Println(v18, v15)
                for fi27 := 0; fi27 < 7; fi27++ {
                    var v28 int16 = (v19 << (((v18 % ((v16) | 1))) & 63))
                    v18 = (v18 ^ v18)
                    _ = fi27
                    _ = v28
                }
                v18 = v16
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

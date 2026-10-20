package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }

func main() {
    var v1 uint16 = 0
    var v2 int64 = 1
    var v3 bool = (true || (v1 <= (v1 << ((uint(v1)) & 63))))
    v4 := make(map[string]bool)
    v3 = (112 != 184)
    if v3 {
        v5 := GenS1{F0: 4294967295, F1: (1 <= 0)}
        if v3 {
            if v3 {
                for fi6 := 0; fi6 < 11; fi6++ {
                    v4[""] = (("a" + "") < "hello")
                    var v7 int32 = -614
                    _ = fi6
                    _ = v7
                }
                if false {
                    fmt.Println(v2, v3)
                    fmt.Println(v1, v3, v2)
                    v2 -= v2
                } else {
                    for fi8 := 0; fi8 < 19; fi8++ {
                        var v9 float64 = 1e10
                        v10 := GenS1{F0: 0, F1: v3}
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                    var v11 int32 = 2147483647
                    _ = v11
                }
                v12 := GenS1{F0: 891, F1: v3}
                _ = v12
            }
            v4[("foo bar" + (("a" + "") + "foo bar"))] = true
            v2 = (^(v2))
        } else {
            _ = len(v4)
            v3 = (!true)
        }
        v13 := []int{32768, -236, -2147483648, 983}
        _ = v13
        _ = v5
    } else {
        v14 := GenS1{F0: 42, F1: true}
        if v3 {
            v1 -= 255
        }
        _ = v14
    }
    v15 := []int{1, 182, 32768, -9223372036854775808}
    var v16 bool = ((v2 & v2) <= (v2 << ((4294967295) & 63)))
    for fi17 := 0; fi17 < 8; fi17++ {
        var v18 int8 = -1
        _ = fi17
        _ = v18
    }
    v1 = 147
    var v19 string = (("a" + "") + "")
    fmt.Println(v1, v19)
    fmt.Println(v19, v3)
    v1 += (v1 / (((v1 ^ (v1 * (v1 << ((11) & 63))))) | 1))
    for fi20 := 0; fi20 < 1; fi20++ {
        switch fi20 {
        case 42:
            for fi21 := 0; fi21 < 17; fi21++ {
                var v22 uint32 = 4294967295
                switch (v19 + v19) {
                case "\n\t":
                    if true {
                        v1 = 0
                    }
                default:
                    var v23 uint8 = 128
                    _ = v23
                }
                _ = fi21
                _ = v22
            }
        default:
            var v24 float32 = 1e10
            _ = v24
        }
        if ((int16(v1) < 1000) && v16) {
            if false {
                v3 = (false || (v3 || v16))
            } else {
                for fi25 := 0; fi25 < 15; fi25++ {
                    var v26 float32 = (0.0 * (3.14159 * 0.0))
                    _ = fi25
                    _ = v26
                }
            }
            var v27 float32 = 0.0
            _ = v27
        } else {
            for fi28 := 0; fi28 < 9; fi28++ {
                var v29 bool = v3
                _ = fi28
                _ = v29
            }
            v16 = v3
        }
        _ = fi20
    }
    v19 = v19
    switch 94 {
    case 510:
        if v16 {
            switch -9223372036854775808 {
            case -145:
                var v30 float32 = float32((1e10 * (0.0 * ((3.14159 / 1.0) + (1.5 / 1.0)))))
                _ = v30
            }
        }
    }
    var v31 float32 = 0.0
    v19 = v19
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v4
}

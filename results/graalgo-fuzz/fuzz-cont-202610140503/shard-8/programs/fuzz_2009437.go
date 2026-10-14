package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint32, p1 uint32, p2 uint32) uint16 {
    var v7 float64 = 1.5
    v8 := GenS1{F0: 4294967295, F1: 32767}
    fmt.Println(v7, p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 32767
}

func main() {
    var v1 bool = false
    var v2 uint16 = 42
    var v3 uint64 = 2
    var v4 uint8 = 1
    var v5 bool = ((4294967295 == 2) || ((1.5 == 1e10) || false))
    var v6 uint64 = uint64(v4)
    fmt.Println(v2, v6, v1)
    if (!(-686 > 510)) {
        v9 := GenS2{F0: 553, F1: 127}
        var v10 float32 = (((-2.25 / 1.5) * ((-2.25 + float32(1e10)) - 1e10)) * 0.0)
        _ = v10
        _ = v9
    } else {
        if false {
            v11 := GenS2{F0: -730, F1: -2147483648}
            for fi12 := 0; fi12 < 2; fi12++ {
                for fi13 := 0; fi13 < 17; fi13++ {
                    var v14 bool = ((!v5) || (!v1))
                    v15 := []int{(fi12 >> ((571) & 63)), (fi12 & fi13)}
                    switch ("" + "") {
                    case "a":
                        var v16 uint32 = 722
                        _ = v16
                    case "":
                        _ = len(v15)
                        var v17 float32 = 1.5
                        v18 := GenS2{F0: -32768, F1: (fi13 * 9223372036854775807)}
                        _ = v17
                        _ = v18
                    }
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                v19 := []int{127, fi12, fi12, (-(fi12))}
                _ = fi12
                _ = v19
            }
            var v20 int32 = 1
            _ = v11
            _ = v20
        } else {
            var v21 uint64 = v6
            switch "" {
            case "hello":
                var v22 int16 = 0
                _ = v22
            }
            if ((v4 / ((1) | 1)) == (v4 | (v4 << ((132) & 63)))) {
                if v5 {
                    fmt.Println(v3)
                } else {
                    var v23 uint = 1
                    _ = v23
                }
            }
            _ = v21
        }
    }
    v3 = (v3 / ((v6) | 1))
    fmt.Println(v4, v6)
    v24 := GenS1{F0: 65535, F1: (v2 - 42)}
    for fi25 := 0; fi25 < 3; fi25++ {
        v5 = (0.0 == (float64(1.5) - -2.25))
        for fi26 := 0; fi26 < 11; fi26++ {
            var v27 float64 = float64(1.5)
            if (v1 && ((v4 - v4) > v4)) {
                switch (fi26 << ((810) & 63)) {
                case 392:
                    v2 = 65535
                default:
                    v1 = ((2147483647 < 4294967295) && false)
                }
                switch "hello" {
                case "":
                    var v28 uint = 65535
                    _ = v28
                case "a":
                    v29 := GenS3{F0: 1}
                    _ = v29
                default:
                    switch "a" {
                    case "":
                        var v30 string = "foo bar"
                        _ = v30
                    default:
                        var v31 uint8 = (v4 << ((2) & 63))
                        _ = v31
                    }
                }
            }
            _ = fi26
            _ = v27
        }
        v1 = v5
        _ = fi25
    }
    if (v4 != 54) {
        if ((uint8(v2) <= (v4 + v4)) || (!true)) {
            v5 = true
            v32 := []int{2147483647, -251, 42, -234}
            _ = v32
        }
    }
    var v33 uint16 = 65535
    v34 := []int{-477, -2147483648}
    fmt.Println(v5, v5, v4)
    if len(v34) >= 2 {
        v34 = v34[1:len(v34)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}

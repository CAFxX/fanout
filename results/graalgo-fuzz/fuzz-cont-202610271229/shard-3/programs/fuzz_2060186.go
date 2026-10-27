package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint8 }
func genFn0() uint32 {
    var v5 int32 = 0
    if true {
        var v6 int32 = v5
        for fi7 := 0; fi7 < 17; fi7++ {
            for fi8 := 0; fi8 < 19; fi8++ {
                v6 = (v5 % ((-1) | 1))
                var v9 float32 = 3.14159
                v5 += 641
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
        v10 := GenS1{F0: 0, F1: "\n\t"}
        _ = v10
        _ = v6
    } else {
        var v11 uint = 140
        _ = v11
    }
    fmt.Println(v5)
    var v12 int16 = 32767
    if (!(!(309 < 209))) {
        if true {
            v13 := []int{42, -1, 85}
            var v14 uint8 = 0
            if ((float64(1.5) > ((float64(1.5) + 1e10) - 0.0)) || (!false)) {
                v14 -= v14
                for fi15 := 0; fi15 < 15; fi15++ {
                    if true {
                        var v16 uint32 = 755
                        v14 -= 255
                        _ = v16
                    }
                    var v17 bool = (!((v5 ^ (v5 << ((4294967295) & 63))) < v5))
                    _ = fi15
                    _ = v17
                }
                var v18 string = ("hello" + ("" + "a"))
                _ = v18
            }
            _ = v13
            _ = v14
        } else {
            v19 := GenS2{F0: 42, F1: 2}
            switch 2147483647 {
            case 127:
                var v20 int32 = (v5 * v5)
                _ = v20
            case 42:
                fmt.Println(v12)
            case -286:
                fmt.Println(v12)
            default:
                var v21 float64 = ((0.1 - 0.0) + (-2.25 + -2.25))
                _ = v21
            }
            _ = v19
        }
        switch "hello" {
        case "hello":
            v12 = (v12 >> ((164) & 63))
        }
        v22 := GenS1{F0: 18446744073709551615, F1: ("" + ("" + ("a" + "hello")))}
        _ = v22
    }
    _ = v12
    _ = v5
    return 42
}
func genFn1() int {
    v23 := []int{}
    v23 = append(v23, -2147483648)
    v23 = append(v23, 9223372036854775807)
    _ = v23
    return -177
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]int)
    var v3 int8 = 27
    v4 := []uint32{2147483647, 2147483648, 2147483648}
    v3 += v3
    fmt.Println(v3)
    fmt.Println(v3)
    v3 -= v3
    fmt.Println(v3)
    v3 += (v3 ^ (v3 << ((547) & 63)))
    fmt.Println(v3)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (!false) {
        if true {
            v4 = append(v4, 1)
        }
        var v24 uint32 = 167
        _ = v24
    }
    v25 := GenS3{F0: 42}
    v3 = (v3 << ((0) & 63))
    switch "hello" {
    case "\n\t":
        fmt.Println(v3)
    case "":
        if false {
            v3 += 0
            v26 := GenS2{F0: 127, F1: 2}
            _ = v26
        }
    case "hello":
        var v27 uint16 = 2
        _ = v27
    }
    var v28 int16 = 32767
    v28 -= -1
    fmt.Println(v28, v3)
    var v29 int8 = v3
    v28 = (-(v28))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}

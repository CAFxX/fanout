package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int32; F1 int16; F2 int64 }
func genFn0() int8 {
    var v7 int = 4294967295
    switch ("\n\t" + ("foo bar" + ("hello" + ""))) {
    case "\n\t":
        v8 := []int{32767, v7, (v7 & (v7 / ((v7) | 1))), (v7 | 1)}
        _ = v8
    case "a":
        v9 := []int{0, v7, -2147483648, (v7 ^ v7)}
        _ = v9
    }
    v10 := []int{(v7 ^ 42), 685}
    v11 := GenS1{F0: 295, F1: 56, F2: 0}
    v7 = (v7 + v7)
    _ = v10
    _ = v11
    _ = v7
    return -41
}

func main() {
    var v1 int = -480
    v2 := make(map[uint32]string)
    var v3 int16 = 1000
    var v4 float32 = (float32(0.0) / 1.5)
    v5 := make(map[uint32]uint32)
    v6 := []uint32{2147483647, 42, 42, 2147483648}
    var v12 string = "\n\t"
    v13 := GenS1{F0: 1, F1: 25, F2: 1}
    v14 := GenS2{F0: v3, F1: 619}
    for fi15 := 0; fi15 < 15; fi15++ {
        v1 = (fi15 >> ((0) & 63))
        v4 -= -2.25
        _ = fi15
    }
    v16 := GenS1{F0: 0, F1: 52, F2: 65535}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        fmt.Println(fi17, v12)
        if (0 > 4294967295) {
            v1 = (v1 << ((4294967295) & 63))
        }
        _ = fi17
    }
    var v18 string = "hello"
    var v19 uint64 = 4294967295
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 uint = 255
        v4 = v4
        for fi22 := 0; fi22 < 7; fi22++ {
            v12 = ("\n\t" + (v18 + ("hello" + "")))
            if (v18 > (v18 + (v12 + (((v18 + v18) + (v18 + "foo bar")) + v12)))) {
                v23 := GenS1{F0: 255, F1: 2, F2: v21}
                _ = v23
            } else {
                v4 += 0.1
            }
            _ = fi22
        }
        _ = fi20
        _ = v21
    }
    v24 := GenS2{F0: v3, F1: 374}
    switch (v1 & (v1 ^ v1)) {
    case 32768:
        v25 := GenS3{F0: 1, F1: v3, F2: 252}
        _ = v25
    case 127:
        v3 -= v3
    case 978:
        v18 = (v18 + v18)
    default:
        var v26 uint = 42
        _ = v26
    }
    v27 := GenS3{F0: 2147483647, F1: (v3 / ((v3) | 1)), F2: 12}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

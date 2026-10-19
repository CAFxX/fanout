package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 int8; F2 uint }
func genFn0() int64 {
    var v5 uint16 = 255
    v6 := GenS1{F0: v5, F1: 42}
    _ = v5
    _ = v6
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 0
    var v3 int = -2147483648
    var v4 uint8 = 255
    if (1.5 == (3.14159 + -2.25)) {
        var v7 uint = 255
        var v8 uint16 = 2
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 string = ""
            _ = fi10
            _ = v11
        }
        v12 := GenS1{F0: 783, F1: 42}
        var v13 string = "\n\t"
        _ = fi9
        _ = v12
        _ = v13
    }
    var v14 int64 = 0
    v3 += v3
    v4 -= v4
    if false {
        _ = len(v1)
    } else {
        if ((true && false) && (0 != genFn0())) {
            var v15 bool = (true || (!(false || true)))
            v16 := GenS2{F0: v4}
            _ = v15
            _ = v16
        }
        v14 += v14
    }
    if (255 != 18446744073709551615) {
        if (false && (false || true)) {
            switch (v3 >> ((722) & 63)) {
            case 0:
                fmt.Println(v3, v14)
            case 32767:
                v17 := GenS1{F0: 65535, F1: (v14 & v14)}
                _ = v17
            case 943:
                for fi18 := 0; fi18 < 10; fi18++ {
                    var v19 int16 = int16(v2)
                    v20 := GenS2{F0: 1}
                    v21 := v1[(v2 + v2)]
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v21
                }
            default:
                v3 -= 1
            }
            var v22 int64 = genFn0()
            v2 = v2
            _ = v22
        } else {
            for fi23 := 0; fi23 < 19; fi23++ {
                for fi24 := 0; fi24 < 20; fi24++ {
                    fmt.Println(fi24, fi23, fi23)
                    _ = fi24
                }
                _ = fi23
            }
            v3 += 865
        }
        v3 = v3
    }
    var v25 int = v3
    v26 := GenS3{F0: 1, F1: 0, F2: 65535}
    if false {
        v27 := GenS2{F0: v4}
        _ = v27
    }
    for fi28 := 0; fi28 < 6; fi28++ {
        var v29 string = ""
        switch "\n\t" {
        case "hello":
            fmt.Println(v4, v29, v14)
        }
        _ = fi28
        _ = v29
    }
    v25 = v25
    v30 := GenS1{F0: 32767, F1: v14}
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v4
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int32 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint64 }
func genFn0() int16 {
    var v5 bool = (367 == 623)
    if (4294967295 > 4294967295) {
        v5 = (!v5)
        if true {
            v6 := []int{-58}
            _ = v6
        }
    } else {
        var v7 int32 = 0
        v7 = (v7 >> ((2) & 63))
        _ = v7
    }
    _ = v5
    return 0
}

func main() {
    v1 := []string{("\n\t" + ("a" + "foo bar")), "hello", ("\n\t" + "\n\t")}
    var v2 string = ("" + ((("a" + ("\n\t" + ("\n\t" + "foo bar"))) + "foo bar") + ("a" + "")))
    var v3 uint8 = 129
    var v4 uint8 = v3
    var v8 int64 = 0
    fmt.Println(v3)
    if true {
        var v9 float64 = 0.1
        var v10 int64 = (v8 * (v8 % ((9223372036854775807) | 1)))
        v10 = v8
        _ = v10
        _ = v9
    }
    var v11 int = -469
    var v12 uint64 = 845
    v4 += uint8(v12)
    v13 := GenS1{F0: v12, F1: v8, F2: 470}
    v8 = v8
    for fi14 := 0; fi14 < 6; fi14++ {
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 string = "a"
            v4 = v3
            fmt.Println(v11, v16)
            var v17 uint16 = 0
            fmt.Println(v2)
            if (true || (0 >= uint32(v4))) {
                var v18 uint16 = 255
                if true {
                    v3 += 128
                    v16 = "hello"
                    v19 := GenS2{F0: uint64(v18), F1: v3}
                    _ = v19
                }
                v8 = (^(v8))
                if true {
                    v20 := GenS3{F0: v12}
                    _ = v20
                } else {
                    v21 := GenS2{F0: (^(v12)), F1: (v4 / (((v4 * 104)) | 1))}
                    _ = v21
                }
                _ = v18
            }
            var v22 float64 = ((float64(1.5) - 1e10) - 0.1)
            v3 += 255
            var v23 bool = false
            _ = fi15
            _ = v16
            _ = v17
            _ = v22
            _ = v23
        }
        switch (fi14 >> ((1) & 63)) {
        case 32768:
            var v24 string = "\n\t"
            _ = v24
        case -105:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = fi14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int16; F1 uint16; F2 int }
func genFn0(p0 uint8, p1 int8) string {
    v5 := []int{-2147483648}
    switch (("\n\t" + "foo bar") + "hello") {
    case "hello":
        for fi6 := 0; fi6 < 11; fi6++ {
            p0 += (p0 - (p0 | (p0 & (^(p0)))))
            _ = fi6
        }
    case "foo bar":
        var v7 bool = false
        _ = v7
    }
    if false {
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 int8 = (p1 | -1)
            _ = fi8
            _ = v9
        }
    } else {
        _ = len(v5)
        var v10 uint8 = (p0 >> ((2) & 63))
        if false {
            var v11 string = ""
            _ = v11
        } else {
            p0 -= 128
        }
        _ = v10
    }
    if (265 == 446) {
        if true {
            fmt.Println(p1)
        } else {
            p0 = (p0 ^ p0)
            var v12 uint8 = 2
            v5 = append(v5, 128)
            _ = v12
        }
    } else {
        v13 := GenS1{F0: 45, F1: 65535, F2: 2}
        _ = v13
    }
    var v14 float32 = float32(1e10)
    _ = p0
    _ = p1
    _ = v14
    _ = v5
    return (("foo bar" + "a") + "foo bar")
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[int]int)
    v3 := []uint32{1, 2147483648}
    var v4 int = 4294967295
    if (false && (v4 > -497)) {
        if false {
            v1 = 136
            v1 = 2
        }
    } else {
        v1 -= v1
        v15 := GenS2{F0: -31, F1: 1, F2: (v4 + (v4 << ((1) & 63)))}
        var v16 uint = 4294967295
        _ = v15
        _ = v16
    }
    var v17 uint64 = 2
    for fi18 := 0; fi18 < 16; fi18++ {
        for fi19 := 0; fi19 < 3; fi19++ {
            switch ("" + "") {
            case "\n\t":
                v1 = (^(v1))
            case "foo bar":
                if false {
                    var v20 int16 = 1
                    var v21 uint = 687
                    if (1.5 >= (0.0 * -2.25)) {
                        v21 = v21
                    } else {
                        fmt.Println(v1, fi19)
                    }
                    _ = v20
                    _ = v21
                }
            case "hello":
                switch (fi18 ^ (-(fi18))) {
                case -9223372036854775808:
                    if false {
                        switch 560 {
                        case 149:
                            v4 -= v4
                        case -9223372036854775808:
                            var v22 int64 = 1
                            _ = v22
                        case 127:
                            fmt.Println(fi18, fi19, v17)
                        }
                        var v23 uint16 = 65535
                        _ = v23
                    }
                    if (false && true) {
                        var v24 uint32 = 42
                        _ = v24
                    }
                    if true {
                        var v25 bool = true
                        v26 := GenS1{F0: 935, F1: 2, F2: 4294967295}
                        v27 := GenS1{F0: 4294967295, F1: 2147483648, F2: uint64(v1)}
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                default:
                    fmt.Println(v1, v1, fi19)
                }
            }
            v28 := GenS1{F0: 800, F1: 2, F2: v17}
            _ = fi19
            _ = v28
        }
        for fi29 := 0; fi29 < 15; fi29++ {
            v1 = v1
            delete(v2, (fi29 | (v4 << ((42) & 63))))
            var v30 float32 = (((float32((1.5 - 1.5)) * (0.1 - -2.25)) / 1.0) + -2.25)
            _ = fi29
            _ = v30
        }
        _ = fi18
    }
    if false {
        var v31 int32 = 2147483647
        _ = v31
    } else {
        var v32 int16 = 1
        var v33 uint8 = (v1 % ((v1) | 1))
        _ = v32
        _ = v33
    }
    v17 -= v17
    v17 = v17
    v34 := v2[(v4 % (((v4 | v4)) | 1))]
    if true {
        v17 += 17
        var v35 int32 = -1
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}

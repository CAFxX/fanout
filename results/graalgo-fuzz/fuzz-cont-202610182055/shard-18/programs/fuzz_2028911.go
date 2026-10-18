package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint16 }
type GenS2 struct { F0 bool; F1 bool; F2 int64 }
type GenS3 struct { F0 int32; F1 uint32 }

func main() {
    v1 := []uint32{}
    var v2 int16 = 32767
    var v3 int = 1
    var v4 int64 = -762
    v5 := make(map[int]uint32)
    fmt.Println(v2, v4)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v5)
    if false {
        for fi6 := 0; fi6 < 18; fi6++ {
            for fi7 := 0; fi7 < 9; fi7++ {
                var v8 int = v3
                if false {
                    v9 := GenS1{F0: ("" + "a"), F1: ((("a" + "") + "foo bar") + "\n\t"), F2: 570}
                    fmt.Println(v4)
                    _ = v9
                } else {
                    switch (^(fi7)) {
                    case 4294967295:
                        v8 = (fi7 - 4294967295)
                    default:
                        fmt.Println(v3)
                    }
                }
                if (v2 >= v2) {
                    var v10 int16 = v2
                    _ = v10
                } else {
                    v11 := v5[(v3 ^ fi7)]
                    if false {
                        v12 := GenS2{F0: (!true), F1: (true && false), F2: v4}
                        _ = v12
                    }
                    if (0 >= 1) {
                        fmt.Println(fi7)
                        if (false || true) {
                            var v13 int8 = -1
                            _ = v13
                        }
                    }
                    var v14 uint16 = 230
                    _ = v11
                    _ = v14
                }
                _ = fi7
                _ = v8
            }
            if true {
                var v15 uint8 = 1
                _ = v15
            }
            for fi16 := 0; fi16 < 18; fi16++ {
                var v17 uint8 = 2
                fmt.Println(v2, fi6)
                _ = fi16
                _ = v17
            }
            _ = fi6
        }
        var v18 uint16 = 65535
        v2 -= 0
        _ = v18
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        if (!true) {
            var v20 string = ("" + "foo bar")
            _ = v20
        }
        fmt.Println(v2, v3)
        _ = len(v1)
        _ = fi19
    }
    v3 -= v3
    if len(v1) > 0 {
        v1[0] = 775
    }
    var v21 int16 = v2
    v22 := GenS3{F0: -496, F1: 65535}
    delete(v5, (v3 >> ((1) & 63)))
    if (639 >= 18446744073709551615) {
        var v23 uint16 = 1
        _ = v23
    } else {
        v24 := GenS3{F0: -999, F1: 2147483648}
        _ = len(v5)
        v25 := GenS2{F0: false, F1: true, F2: v4}
        _ = v24
        _ = v25
    }
    var v26 uint64 = 0
    v27 := GenS2{F0: false, F1: true, F2: v4}
    v26 = 284
    if (502 < 65535) {
        var v28 uint = 2
        _ = v28
    } else {
        v4 -= (v4 & (v4 * v4))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}

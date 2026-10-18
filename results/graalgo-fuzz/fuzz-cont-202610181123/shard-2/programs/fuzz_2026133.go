package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int64 }
func genFn0(p0 int, p1 uint8, p2 uint64) int8 {
    fmt.Println(p0, p1)
    fmt.Println(p1, p0)
    p2 = (p2 * (p2 % ((p2) | 1)))
    var v5 float32 = 1e10
    v6 := GenS1{F0: 42, F1: 1, F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 56
}

func main() {
    v1 := []int{}
    var v2 int64 = 130
    v3 := make(map[int]int)
    v4 := make(map[string]bool)
    var v7 bool = (-1 > -32768)
    if (("a" + ("a" + ("hello" + ("foo bar" + "hello")))) >= "\n\t") {
        if len(v1) > 0 {
            v1[0] = -1
        }
        switch int(v2) {
        case -587:
            if ((42 == 42) || (13 > 552)) {
                if (false || v7) {
                    var v8 uint = 209
                    _ = v8
                }
                v9 := GenS1{F0: (^(v2)), F1: 0, F2: v2}
                _ = v9
            }
        case 9223372036854775807:
            if v7 {
                var v10 int = -2147483648
                fmt.Println(v10, v10)
                for fi11 := 0; fi11 < 9; fi11++ {
                    v10 -= (fi11 % (((v10 + 674)) | 1))
                    var v12 int64 = (v2 << ((4294967295) & 63))
                    var v13 uint = 170
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                _ = v10
            }
        }
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        v15 := v3[(fi14 * fi14)]
        v16 := v4[("\n\t" + "a")]
        v17 := GenS1{F0: (^(v2)), F1: 633, F2: (-(v2))}
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    var v18 uint64 = 42
    if ((1e10 * 1.5) < 1e10) {
        for fi19 := 0; fi19 < 20; fi19++ {
            if false {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            var v20 int16 = -32768
            v7 = true
            _ = fi19
            _ = v20
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if (558 <= 1) {
            var v21 uint32 = 2147483648
            var v22 float32 = (((float32(1e10) * 0.1) + float32(0.0)) - 0.1)
            if v7 {
                if false {
                    for fi23 := 0; fi23 < 14; fi23++ {
                        v21 = v21
                        _ = fi23
                    }
                }
                var v24 float32 = -2.25
                v18 = 42
                _ = v24
            } else {
                for fi25 := 0; fi25 < 18; fi25++ {
                    _ = len(v1)
                    fmt.Println(v2, v18)
                    _ = fi25
                }
                var v26 uint64 = 0
                _ = v26
            }
            _ = v21
            _ = v22
        } else {
            v18 = 1
        }
        var v27 uint16 = 885
        _ = v27
    }
    var v28 int32 = 1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v7, v7, v18)
    switch ("hello" + "\n\t") {
    case "foo bar":
        if (!v7) {
            fmt.Println(v28, v28, v7)
            v2 = v2
        } else {
            var v29 uint32 = 586
            _ = v29
        }
    case "hello":
        fmt.Println(v2)
    case "a":
        v28 = (v28 / ((v28) | 1))
    }
    _ = len(v1)
    var v30 string = ""
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v7
}

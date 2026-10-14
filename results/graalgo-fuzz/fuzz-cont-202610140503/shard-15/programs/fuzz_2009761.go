package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int }
func genFn0(p0 int8) int64 {
    switch 32767 {
    case 128:
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 uint16 = 646
            _ = fi7
            _ = v8
        }
    case -232:
        switch 9223372036854775807 {
        case -575:
            var v9 int8 = 0
            _ = v9
        case 214:
            p0 = p0
        case -821:
            if (true || true) {
                p0 += p0
                p0 = (p0 / (((p0 - (p0 >> ((179) & 63)))) | 1))
            } else {
                v10 := []int{}
                _ = v10
            }
        default:
            v11 := []int{32768}
            _ = v11
        }
    }
    switch (("a" + ("\n\t" + "a")) + ("a" + "a")) {
    case "foo bar":
        v12 := GenS1{F0: -2147483648, F1: 90, F2: 0}
        _ = v12
    case "":
        p0 = (p0 % (((p0 + p0)) | 1))
    case "\n\t":
        var v13 int16 = 1000
        _ = v13
    }
    switch 1 {
    case -554:
        p0 -= (p0 | p0)
    default:
        if false {
            if (577 != 65535) {
                v14 := []int{9223372036854775807, 42, 0, -174}
                p0 = 127
                if true {
                    if len(v14) > 0 {
                        v14[0] = 145
                    }
                } else {
                    v15 := GenS2{F0: 2}
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                    _ = v15
                }
                _ = v14
            } else {
                switch "a" {
                case "foo bar":
                    v16 := GenS1{F0: -537, F1: 42, F2: 372}
                    _ = v16
                case "":
                    if (true || (-445 < -949)) {
                        var v17 uint32 = 154
                        fmt.Println(p0, v17)
                        v17 = v17
                        _ = v17
                    } else {
                        p0 += (p0 & (p0 % ((p0) | 1)))
                        var v18 int = 101
                        _ = v18
                    }
                default:
                    switch 0 {
                    case 32767:
                        if (!false) {
                            v19 := GenS2{F0: 42}
                            _ = v19
                        }
                    }
                }
            }
        } else {
            for fi20 := 0; fi20 < 18; fi20++ {
                v21 := []int{(fi20 | (fi20 - (fi20 & (fi20 / ((9223372036854775807) | 1))))), (fi20 ^ (fi20 * fi20)), fi20}
                switch (fi20 * (fi20 * 326)) {
                case 440:
                    p0 -= p0
                case 294:
                    v22 := GenS3{F0: fi20}
                    _ = v22
                }
                _ = fi20
                _ = v21
            }
        }
    }
    _ = p0
    return 359
}
func genFn1() uint16 {
    var v23 int8 = -1
    var v24 int8 = 0
    v23 = v24
    _ = v23
    _ = v24
    return 65535
}

func main() {
    var v1 uint16 = 410
    v2 := make(map[int]bool)
    v3 := []int{}
    var v4 float32 = 3.14159
    var v5 float32 = v4
    var v6 int32 = -1
    v4 = -2.25
    _ = len(v3)
    fmt.Println(v5)
    v4 = 3.14159
    var v25 uint16 = 255
    var v26 int16 = -1
    var v27 int16 = (v26 ^ 32767)
    if false {
        v25 = (v1 + (v1 << ((255) & 63)))
    } else {
        switch -891 {
        case -2147483648:
            v6 -= -2147483648
        }
        v5 = ((((v4 - 0.1) - v4) / float32(((1e10 - 1.5) / (-2.25 - ((((1e10 - 1.5) * -2.25) + (1.5 - (3.14159 + 0.1))) + (-2.25 + 1e10)))))) / v4)
    }
    delete(v2, 127)
    _ = len(v2)
    delete(v2, -2147483648)
    for fi28 := 0; fi28 < 5; fi28++ {
        v25 += (v1 / ((1) | 1))
        _ = fi28
    }
    v1 += (v25 ^ 42)
    _ = len(v3)
    switch -192 {
    case 2147483647:
        var v29 uint8 = 19
        _ = v29
    default:
        if (-9223372036854775808 > -490) {
            v26 += v27
            var v30 int8 = int8(v26)
            fmt.Println(v1, v26, v5)
            _ = v30
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint16, p1 uint16) int64 {
    for fi5 := 0; fi5 < 4; fi5++ {
        p0 = 32767
        _ = fi5
    }
    p0 -= (p0 & p1)
    _ = p0
    _ = p1
    return -1
}

func main() {
    var v1 int32 = 1
    v2 := make(map[string]bool)
    var v3 float64 = 1e10
    var v4 uint16 = 42
    v6 := []int{-1}
    v3 = v3
    var v7 int8 = 1
    if ((^(v7)) != v7) {
        v2["hello"] = true
        if true {
            v8 := GenS1{F0: (float32(float64(float32((v3 * 1.5)))) < 3.14159), F1: 0, F2: 1}
            var v9 int64 = -897
            if true {
                for fi10 := 0; fi10 < 9; fi10++ {
                    var v11 bool = (genFn0((v4 << ((65535) & 63)), 255) >= v9)
                    _ = fi10
                    _ = v11
                }
                for fi12 := 0; fi12 < 4; fi12++ {
                    var v13 int32 = (v1 & (v1 ^ (v1 >> ((908) & 63))))
                    v6 = append(v6, fi12)
                    v13 = (v1 / ((-956) | 1))
                    _ = fi12
                    _ = v13
                }
            }
            _ = v8
            _ = v9
        }
    } else {
        v14 := GenS2{F0: -32768}
        _ = v14
    }
    var v15 int32 = (v1 ^ (v1 ^ v1))
    var v16 uint32 = 4294967295
    var v17 int = 294
    v18 := GenS1{F0: true, F1: v7, F2: -416}
    if (495 == 18446744073709551615) {
        var v19 uint32 = (v16 | (v16 & (v16 & v16)))
        if false {
            switch "a" {
            case "a":
                var v20 int64 = -1
                _ = v20
            case "\n\t":
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case "foo bar":
                var v21 float32 = 3.14159
                _ = v21
            }
            var v22 uint8 = 2
            var v23 uint64 = 2
            _ = v22
            _ = v23
        } else {
            switch ("foo bar" + "hello") {
            case "a":
                v6 = append(v6, v17)
            case "\n\t":
                v24 := GenS1{F0: false, F1: v7, F2: 1000}
                _ = v24
            case "foo bar":
                fmt.Println(v3, v15, v19)
            default:
                v19 = 4294967295
            }
            if true {
                for fi25 := 0; fi25 < 15; fi25++ {
                    v3 += v3
                    v26 := GenS1{F0: false, F1: -49, F2: 32767}
                    switch (v17 >> ((42) & 63)) {
                    case 2147483647:
                        fmt.Println(v1, v1)
                    case 4294967295:
                        v15 = v15
                    case -846:
                        var v27 uint = 213
                        _ = v27
                    default:
                        var v28 int64 = genFn0((v4 ^ (v4 & v4)), v4)
                        _ = v28
                    }
                    _ = fi25
                    _ = v26
                }
            } else {
                for fi29 := 0; fi29 < 14; fi29++ {
                    var v30 int16 = -1
                    _ = fi29
                    _ = v30
                }
                for fi31 := 0; fi31 < 18; fi31++ {
                    var v32 int16 = -32768
                    for fi33 := 0; fi33 < 20; fi33++ {
                        fmt.Println(v7)
                        v15 -= (v1 >> ((948) & 63))
                        _ = fi33
                    }
                    _ = fi31
                    _ = v32
                }
            }
            fmt.Println(v7, v15, v1)
        }
        _ = v19
    }
    delete(v2, "hello")
    v34 := GenS2{F0: -1}
    var v35 uint32 = v16
    var v36 float32 = 0.1
    if (-32768 > -591) {
        v17 = (v17 * v17)
    }
    v4 -= v4
    if false {
        if len(v6) > 0 {
            v6[0] = (v17 % ((int(v7)) | 1))
        }
    } else {
        delete(v2, ("foo bar" + "\n\t"))
        switch (v17 & (v17 & v17)) {
        case -438:
            switch v17 {
            case 32768:
                v15 = (v15 / (((v15 | v1)) | 1))
            case 32767:
                v37 := GenS1{F0: (v36 > v36), F1: v7, F2: -752}
                _ = v37
            }
        case 42:
            if true {
                fmt.Println(v1, v15)
                v38 := GenS2{F0: 1000}
                fmt.Println(v15, v15)
                _ = v38
            } else {
                v16 -= (v35 * (v35 & (v35 / ((v16) | 1))))
                v39 := GenS2{F0: -32768}
                for fi40 := 0; fi40 < 10; fi40++ {
                    var v41 bool = ((v15 / ((v15) | 1)) < v15)
                    _ = fi40
                    _ = v41
                }
                _ = v39
            }
        default:
            var v42 int = v17
            _ = v42
        }
    }
    if len(v6) > 0 {
        v6[0] = (v17 & v17)
    }
    if (true && false) {
        fmt.Println(v4, v36)
        if ((32767 > 42) && false) {
            _ = len(v6)
            v36 += -2.25
        }
    } else {
        v4 -= 255
        var v43 int8 = v7
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v6
    _ = v7
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 bool }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64; F1 uint; F2 bool }
type GenS6 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS7 struct { F0 int }
func genFn0(p0 int64, p1 int) int64 {
    if false {
        v6 := GenS1{F0: 128, F1: "\n\t", F2: true}
        p1 = p1
        _ = v6
    } else {
        p0 -= (p0 % (((p0 - (p0 % ((p0) | 1)))) | 1))
        switch "hello" {
        case "\n\t":
            p0 = (p0 & p0)
        case "foo bar":
            for fi7 := 0; fi7 < 18; fi7++ {
                if (true || (true && ((^(p1)) <= fi7))) {
                    p1 -= p1
                    var v8 int = (-(p1))
                    p0 += p0
                    _ = v8
                }
                _ = fi7
            }
        }
    }
    v9 := GenS2{F0: p1, F1: 526}
    _ = p0
    _ = p1
    _ = v9
    return p0
}

func main() {
    v1 := []string{"foo bar", (("" + "a") + "\n\t")}
    var v2 int = -352
    var v3 uint8 = 255
    var v4 int8 = -1
    var v5 int32 = -957
    for fi10 := 0; fi10 < 4; fi10++ {
        v11 := GenS3{F0: v5}
        _ = fi10
        _ = v11
    }
    if (v3 > (v3 % ((2) | 1))) {
        var v12 uint32 = 4294967295
        _ = v12
    }
    var v13 int = (v2 >> ((148) & 63))
    if (3.14159 >= float32(3.14159)) {
        if true {
            fmt.Println(v4)
            if false {
                var v14 int16 = 1
                switch (v13 & int(v4)) {
                case 0:
                    var v15 bool = (463 > 2)
                    _ = v15
                case 42:
                    for fi16 := 0; fi16 < 4; fi16++ {
                        for fi17 := 0; fi17 < 2; fi17++ {
                            v2 += v13
                            _ = fi17
                        }
                        v4 -= 72
                        _ = fi16
                    }
                case -530:
                    var v18 int8 = v4
                    _ = v18
                }
                _ = v14
            } else {
                var v19 uint64 = 79
                _ = v19
            }
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
        } else {
            var v20 int8 = v4
            if false {
                for fi21 := 0; fi21 < 13; fi21++ {
                    v22 := GenS4{F0: 2147483647}
                    for fi23 := 0; fi23 < 20; fi23++ {
                        var v24 uint64 = 662
                        _ = len(v1)
                        fmt.Println(fi23)
                        _ = fi23
                        _ = v24
                    }
                    var v25 int16 = -286
                    _ = fi21
                    _ = v22
                    _ = v25
                }
                v5 = v5
            } else {
                switch "\n\t" {
                case "\n\t":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "hello":
                    v26 := GenS5{F0: -9223372036854775808, F1: 2, F2: false}
                    _ = v26
                default:
                    var v27 uint64 = 2
                    _ = v27
                }
            }
            fmt.Println(v20, v5, v5)
            _ = v20
        }
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        var v29 int16 = -1
        for fi30 := 0; fi30 < 10; fi30++ {
            if (!(!(v29 < (v29 * (v29 & v29))))) {
                fmt.Println(v5, v13)
            } else {
                fmt.Println(v5, v13)
            }
            _ = fi30
        }
        v31 := GenS6{F0: 0, F1: 255, F2: 0}
        _ = fi28
        _ = v29
        _ = v31
    }
    fmt.Println(v4, v3, v2)
    v4 = (-(v4))
    v4 += 0
    var v32 int16 = -3
    fmt.Println(v5)
    v33 := GenS7{F0: (v2 / (((v13 << ((65535) & 63))) | 1))}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}

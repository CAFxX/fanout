package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 int64; F2 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64; F1 int8; F2 int32 }
func genFn0(p0 uint, p1 int8) int16 {
    if true {
        if false {
            var v5 int8 = 0
            if false {
                p1 = p1
                p1 += (v5 & (p1 | p1))
                fmt.Println(p0, p1)
            }
            _ = v5
        }
    } else {
        v6 := GenS1{F0: int32(p1)}
        _ = v6
    }
    v7 := GenS1{F0: 20}
    p0 -= p0
    for fi8 := 0; fi8 < 20; fi8++ {
        p1 = (p1 | 100)
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := []int{}
            var v11 uint16 = 1
            var v12 int16 = 0
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v7
    return -32768
}
func genFn1() string {
    var v13 uint16 = 255
    var v14 float32 = (0.0 * 0.1)
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 uint8 = 73
        v14 += float32(0.1)
        _ = fi15
        _ = v16
    }
    v13 = (v13 & (v13 / ((648) | 1)))
    v13 += v13
    _ = v13
    _ = v14
    return ("a" + ((("hello" + "foo bar") + "hello") + "\n\t"))
}

func main() {
    v1 := []uint32{65535, 600, 65535}
    var v2 uint32 = 375
    var v3 bool = true
    var v4 int = -253
    for fi17 := 0; fi17 < 4; fi17++ {
        v18 := GenS2{F0: -1}
        v4 = v4
        _ = fi17
        _ = v18
    }
    var v19 uint8 = 0
    v3 = v3
    if true {
        v3 = false
        if false {
            for fi20 := 0; fi20 < 10; fi20++ {
                if v3 {
                    var v21 int64 = -9223372036854775808
                    v1 = append(v1, 2147483648)
                    _ = v21
                } else {
                    switch v4 {
                    case 2147483647:
                        v22 := GenS3{F0: 96, F1: -876, F2: ("\n\t" + "foo bar")}
                        _ = v22
                    case 4294967295:
                        var v23 int8 = -6
                        _ = v23
                    case 128:
                        if (0 > 672) {
                            v24 := GenS4{F0: 65535}
                            for fi25 := 0; fi25 < 2; fi25++ {
                                var v26 uint = 255
                                _ = fi25
                                _ = v26
                            }
                            v4 -= v4
                            _ = v24
                        } else {
                            v27 := GenS5{F0: (v19 & (v19 + 1))}
                            _ = len(v1)
                            for fi28 := 0; fi28 < 14; fi28++ {
                                v1 = append(v1, v2)
                                _ = fi28
                            }
                            _ = v27
                        }
                    }
                    fmt.Println(v3, fi20)
                }
                var v29 uint8 = 47
                _ = fi20
                _ = v29
            }
            var v30 int16 = -32768
            var v31 int8 = 0
            _ = v30
            _ = v31
        } else {
            for fi32 := 0; fi32 < 10; fi32++ {
                fmt.Println(v2)
                switch ((("\n\t" + ("foo bar" + "\n\t")) + ("foo bar" + "a")) + "") {
                case "foo bar":
                    if (0 <= 255) {
                        for fi33 := 0; fi33 < 18; fi33++ {
                            var v34 int64 = -9223372036854775808
                            v1 = append(v1, (v2 * (v2 + (v2 * (^(v2))))))
                            _ = fi33
                            _ = v34
                        }
                        fmt.Println(v19, v3, v3)
                        var v35 int32 = -255
                        _ = v35
                    }
                case "":
                    v3 = v3
                default:
                    if len(v1) > 0 {
                        v1[0] = v2
                    }
                }
                _ = fi32
            }
        }
    }
    _ = len(v1)
    v4 -= v4
    v4 = (v4 + v4)
    v36 := GenS6{F0: 546, F1: 100, F2: -1}
    switch (v4 | v4) {
    case 581:
        for fi37 := 0; fi37 < 5; fi37++ {
            for fi38 := 0; fi38 < 12; fi38++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi38
            }
            var v39 int8 = 0
            _ = fi37
            _ = v39
        }
    case 134:
        v3 = v3
    case -1:
        v1 = append(v1, 2147483648)
    default:
        v4 -= (v4 << ((65535) & 63))
    }
    fmt.Println(v3, v4)
    var v40 uint32 = (v2 >> ((1) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v40
}

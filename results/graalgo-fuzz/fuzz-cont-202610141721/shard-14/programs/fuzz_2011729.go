package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 uint8; F1 int64; F2 int }
type GenS4 struct { F0 uint64; F1 uint8; F2 string }
func genFn0(p0 int, p1 int8, p2 uint8) string {
    var v5 uint64 = 1
    p2 -= (^(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return "hello"
}
func genFn1(p0 string, p1 bool) int {
    p1 = p1
    p1 = p1
    _ = p0
    _ = p1
    return 32768
}

func main() {
    var v1 int8 = 42
    var v2 string = "foo bar"
    var v3 uint8 = 159
    var v4 int16 = 1000
    var v6 int64 = -9223372036854775808
    var v7 bool = true
    v8 := []int{}
    var v9 int32 = 355
    var v10 uint64 = 42
    v6 -= v6
    v11 := GenS1{F0: ""}
    for fi12 := 0; fi12 < 18; fi12++ {
        v8 = append(v8, fi12)
        _ = fi12
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        v14 := GenS2{F0: int16(v10), F1: 2}
        var v15 int32 = v9
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = len(v8)
    for fi16 := 0; fi16 < 16; fi16++ {
        v17 := GenS2{F0: int16(fi16), F1: 42}
        var v18 uint64 = 1
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println(v10)
    switch v2 {
    case "foo bar":
        v1 -= 1
    case "a":
        v19 := GenS1{F0: v2}
        _ = v19
    case "\n\t":
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    default:
        if (!v7) {
            v20 := GenS1{F0: v2}
            if (32767 != 4294967295) {
                if v7 {
                    var v21 uint64 = v10
                    if (0 < 2147483647) {
                        for fi22 := 0; fi22 < 3; fi22++ {
                            v23 := GenS2{F0: v4, F1: 32767}
                            _ = fi22
                            _ = v23
                        }
                        v24 := GenS2{F0: v4, F1: 155}
                        _ = v24
                    } else {
                        fmt.Println(v2, v21)
                        v25 := GenS1{F0: (((("\n\t" + v2) + v2) + v2) + ("\n\t" + "a"))}
                        v1 += (-(v1))
                        _ = v25
                    }
                    for fi26 := 0; fi26 < 14; fi26++ {
                        fmt.Println(v3, fi26)
                        v1 = -50
                        for fi27 := 0; fi27 < 1; fi27++ {
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            switch (^(fi27)) {
                            case -936:
                                v1 = -60
                            default:
                                var v28 int32 = v9
                                _ = v28
                            }
                            fmt.Println(v21, v10)
                            _ = fi27
                        }
                        _ = fi26
                    }
                    _ = v21
                } else {
                    switch v2 {
                    case "\n\t":
                        v3 -= (v3 >> ((65535) & 63))
                    }
                    for fi29 := 0; fi29 < 15; fi29++ {
                        if len(v8) > 0 {
                            v8[0] = genFn1(v2, true)
                        }
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        for fi30 := 0; fi30 < 11; fi30++ {
                            var v31 uint = 2
                            v4 += v4
                            _ = fi30
                            _ = v31
                        }
                        _ = fi29
                    }
                }
                switch int(v6) {
                case 9223372036854775807:
                    fmt.Println(v1, v3, v1)
                case -79:
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                default:
                    v4 -= v4
                }
                if len(v8) > 0 {
                    v8[0] = 32768
                }
            } else {
                var v32 int = -9223372036854775808
                v9 -= v9
                for fi33 := 0; fi33 < 11; fi33++ {
                    for fi34 := 0; fi34 < 5; fi34++ {
                        v9 -= (v9 % ((v9) | 1))
                        _ = fi34
                    }
                    v35 := GenS3{F0: 2, F1: -963, F2: v32}
                    _ = fi33
                    _ = v35
                }
                _ = v32
            }
            _ = v20
        }
    }
    var v36 uint8 = 2
    switch "hello" {
    case "":
        v10 = (v10 & (v10 - v10))
    case "\n\t":
        for fi37 := 0; fi37 < 13; fi37++ {
            v6 += v6
            v38 := GenS3{F0: (v36 << ((217) & 63)), F1: v6, F2: 4294967295}
            v39 := GenS4{F0: (^(v10)), F1: v36, F2: v2}
            _ = fi37
            _ = v38
            _ = v39
        }
    case "foo bar":
        var v40 int64 = (-(v6))
        _ = v40
    }
    if v7 {
        v41 := GenS4{F0: (v10 << ((450) & 63)), F1: 1, F2: "hello"}
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

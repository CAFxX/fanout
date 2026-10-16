package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int32; F1 int16; F2 uint }
type GenS7 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS8 struct { F0 int8; F1 uint32; F2 int32 }
type GenS9 struct { F0 int64; F1 uint64 }
func genFn0(p0 uint8, p1 int8, p2 bool) int64 {
    v5 := GenS1{F0: int8(p0)}
    for fi6 := 0; fi6 < 10; fi6++ {
        if (-649 != 9223372036854775807) {
            fmt.Println(p0)
            v7 := []int{}
            if false {
                if (!p2) {
                    p0 = (p0 / (((^(p0))) | 1))
                    fmt.Println(fi6, fi6, p0)
                }
                if (718 < 2147483648) {
                    p2 = (42 <= 828)
                }
            } else {
                var v8 uint64 = 734
                var v9 float64 = 1.5
                v10 := GenS2{F0: 912}
                _ = v10
                _ = v8
                _ = v9
            }
            _ = v7
        } else {
            p1 = 1
        }
        _ = fi6
    }
    p1 = 127
    fmt.Println(p2, p1)
    for fi11 := 0; fi11 < 15; fi11++ {
        if (false && true) {
            p2 = false
        } else {
            var v12 string = ("hello" + "foo bar")
            for fi13 := 0; fi13 < 4; fi13++ {
                v14 := GenS3{F0: 474, F1: 129}
                if false {
                    v15 := GenS3{F0: 1000, F1: 382}
                    var v16 string = (("\n\t" + ("" + v12)) + v12)
                    _ = v15
                    _ = v16
                } else {
                    fmt.Println(v12)
                    var v17 uint64 = 1
                    var v18 uint16 = 828
                    _ = v17
                    _ = v18
                }
                _ = fi13
                _ = v14
            }
            v19 := GenS4{F0: (fi11 - (-(fi11)))}
            _ = v12
            _ = v19
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 42
    var v3 uint16 = 240
    var v4 int = 128
    for fi20 := 0; fi20 < 16; fi20++ {
        fmt.Println(fi20, v2, v3)
        var v21 uint64 = (v2 | (v2 - (v2 << ((513) & 63))))
        _ = fi20
        _ = v21
    }
    if false {
        v22 := GenS5{F0: (v2 & v2)}
        v23 := GenS6{F0: 0, F1: 420, F2: 427}
        switch ("" + "foo bar") {
        case "":
            switch ((("hello" + ("\n\t" + "hello")) + "\n\t") + "\n\t") {
            case "foo bar":
                v2 += v2
            case "":
                v4 = v4
            case "\n\t":
                if false {
                    var v24 bool = (831 <= int16(v4))
                    for fi25 := 0; fi25 < 15; fi25++ {
                        for fi26 := 0; fi26 < 4; fi26++ {
                            fmt.Println(v2, v3)
                            var v27 int16 = -32768
                            v27 -= (v27 + v27)
                            _ = fi26
                            _ = v27
                        }
                        v3 += uint16(v2)
                        _ = fi25
                    }
                    v28 := GenS4{F0: v4}
                    _ = v24
                    _ = v28
                } else {
                    var v29 int = v4
                    var v30 uint64 = v2
                    fmt.Println(v4)
                    _ = v29
                    _ = v30
                }
            }
        case "a":
            fmt.Println(v3)
        case "foo bar":
            v31 := GenS7{F0: 992, F1: v2, F2: (v3 + v3)}
            _ = v31
        default:
            v32 := GenS4{F0: v4}
            _ = v32
        }
        _ = v22
        _ = v23
    } else {
        for fi33 := 0; fi33 < 10; fi33++ {
            v2 = 42
            _ = fi33
        }
        v34 := GenS8{F0: 65, F1: 4294967295, F2: 42}
        var v35 int = 1
        _ = v34
        _ = v35
    }
    var v36 uint64 = (v2 - 180)
    if true {
        var v37 bool = true
        v2 = v2
        if v37 {
            v38 := []int{}
            _ = v38
        }
        _ = v37
    } else {
        switch v4 {
        case 399:
            var v39 uint8 = 2
            _ = v39
        case 2147483647:
            var v40 uint = 4294967295
            _ = v40
        }
        if true {
            v3 = (v3 >> ((255) & 63))
            var v41 string = (("a" + "a") + ("" + "hello"))
            v42 := GenS9{F0: 0, F1: (^(v2))}
            _ = v41
            _ = v42
        } else {
            v43 := GenS3{F0: 992, F1: 631}
            fmt.Println(v2, v36)
            var v44 uint64 = v36
            _ = v43
            _ = v44
        }
    }
    _ = len(v1)
    v45 := v1["foo bar"]
    v36 += (^(v2))
    if ((^(v3)) <= v3) {
        fmt.Println(v36)
        fmt.Println(v36, v4, v36)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v45
}

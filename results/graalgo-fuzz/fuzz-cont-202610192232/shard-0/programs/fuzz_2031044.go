package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 string, p1 uint32, p2 int8) uint64 {
    p2 += (^(p2))
    var v7 int16 = -166
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2
}
func genFn1(p0 int16, p1 uint32, p2 uint16) uint64 {
    for fi8 := 0; fi8 < 16; fi8++ {
        switch -385 {
        case -465:
            v9 := []int{(^(fi8)), fi8}
            _ = v9
        }
        _ = fi8
    }
    v10 := GenS1{F0: p1}
    var v11 int8 = 100
    p2 += p2
    v12 := []int{0, -2147483648, 9223372036854775807, 847}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return 489
}

func main() {
    var v1 uint = 4294967295
    v2 := make(map[uint32]string)
    v3 := []int{32768, -111}
    var v4 uint8 = 128
    var v5 int8 = -1
    v6 := []int{}
    delete(v2, 169)
    switch "hello" {
    case "a":
        v4 += (v4 + v4)
    default:
        v1 = v1
    }
    var v13 int8 = (v5 ^ v5)
    if true {
        var v14 uint = (v1 << ((v1) & 63))
        _ = v14
    } else {
        v15 := GenS1{F0: 2147483648}
        var v16 int = -648
        for fi17 := 0; fi17 < 11; fi17++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi17
        }
        _ = v15
        _ = v16
    }
    v18 := GenS2{F0: (v13 << (((v1 + 377)) & 63))}
    v5 -= 127
    if false {
        if false {
            v19 := GenS1{F0: 42}
            for fi20 := 0; fi20 < 14; fi20++ {
                v13 = -1
                var v21 uint = (v1 * (v1 & 0))
                _ = fi20
                _ = v21
            }
            for fi22 := 0; fi22 < 18; fi22++ {
                fmt.Println(fi22)
                for fi23 := 0; fi23 < 15; fi23++ {
                    var v24 uint16 = 32767
                    v25 := GenS1{F0: uint32(v1)}
                    var v26 int32 = 42
                    v27 := GenS2{F0: v13}
                    v24 = 1
                    if (!(genFn1(-32768, 980, v24) <= 551)) {
                        v5 = v5
                        var v28 uint8 = 42
                        _ = v28
                    } else {
                        switch fi22 {
                        case 648:
                            fmt.Println(v13, v5, v4)
                        }
                        v29 := GenS3{F0: 2}
                        _ = v29
                    }
                    var v30 int64 = 331
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                    _ = v27
                    _ = v30
                }
                _ = fi22
            }
            _ = v19
        }
        v3 = append(v3, -99)
        v4 = v4
    }
    var v31 bool = ((^(v1)) <= 255)
    for fi32 := 0; fi32 < 1; fi32++ {
        switch (fi32 << (((v1 - (^(v1)))) & 63)) {
        case -2147483648:
            v33 := v2[uint32(v4)]
            _ = v33
        case 1:
            var v34 uint8 = 131
            _ = v34
        default:
            for fi35 := 0; fi35 < 4; fi35++ {
                var v36 int = (-(fi32))
                _ = fi35
                _ = v36
            }
        }
        var v37 uint32 = 359
        v13 -= v5
        _ = fi32
        _ = v37
    }
    var v38 string = "foo bar"
    v4 = v4
    v39 := GenS2{F0: (-(v5))}
    v40 := GenS1{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}

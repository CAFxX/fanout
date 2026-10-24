package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 string }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint64; F1 string; F2 int8 }

func main() {
    var v1 uint32 = 2147483648
    v2 := []uint32{2147483648, v1, (v1 + 2), v1}
    var v3 int32 = -1
    var v4 bool = true
    v5 := []uint32{}
    var v6 int64 = 500
    if false {
        if v4 {
            v1 = 2147483647
            fmt.Println(v1, v4)
            var v7 int32 = (v3 & (-(v3)))
            _ = v7
        }
    } else {
        var v8 int32 = 1
        _ = v8
    }
    var v9 uint16 = 42
    if (uint64(v9) <= 1) {
        for fi10 := 0; fi10 < 4; fi10++ {
            for fi11 := 0; fi11 < 10; fi11++ {
                switch (fi10 - (fi11 * (fi10 / ((fi11) | 1)))) {
                case -771:
                    fmt.Println(v1)
                case -669:
                    var v12 uint8 = 42
                    _ = v12
                }
                v9 += v9
                v3 = (v3 - -128)
                _ = fi11
            }
            var v13 int64 = (v6 ^ v6)
            _ = fi10
            _ = v13
        }
        var v14 bool = (!(-69 == int16(v6)))
        v1 = 122
        _ = v14
    } else {
        if ((v4 || ((3.14159 >= float32(((0.1 - (0.0 - -2.25)) / 0.1))) || v4)) && (v4 || (true || false))) {
            v15 := GenS1{F0: v4, F1: 4294967295, F2: "foo bar"}
            _ = v15
        } else {
            var v16 int64 = (v6 >> ((255) & 63))
            var v17 int16 = -96
            _ = v16
            _ = v17
        }
        for fi18 := 0; fi18 < 7; fi18++ {
            for fi19 := 0; fi19 < 10; fi19++ {
                if len(v5) > 0 {
                    v5[0] = 2147483648
                }
                _ = fi19
            }
            if len(v2) > 0 {
                v2[0] = (v1 >> ((4294967295) & 63))
            }
            _ = fi18
        }
        v9 = v9
    }
    v9 = v9
    if v4 {
        v20 := GenS2{F0: ("hello" + ""), F1: (32767 > 32767)}
        _ = v20
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        v22 := GenS1{F0: true, F1: (fi21 % ((fi21) | 1)), F2: ("hello" + (("" + "\n\t") + "a"))}
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 11; fi23++ {
        if false {
            fmt.Println(v9, v9)
            fmt.Println(v3, v4)
            v9 += 0
        }
        var v24 int32 = -595
        _ = fi23
        _ = v24
    }
    v2 = append(v2, 1)
    for fi25 := 0; fi25 < 18; fi25++ {
        var v26 uint64 = 0
        for fi27 := 0; fi27 < 15; fi27++ {
            v28 := GenS3{F0: (v26 % ((v26) | 1)), F1: "a", F2: 42}
            _ = len(v2)
            v26 = v26
            _ = fi27
            _ = v28
        }
        _ = fi25
        _ = v26
    }
    v29 := GenS2{F0: "a", F1: v4}
    switch 925 {
    case 682:
        if len(v2) > 0 {
            v2[0] = (v1 / ((v1) | 1))
        }
    case -77:
        v4 = v4
    default:
        for fi30 := 0; fi30 < 17; fi30++ {
            v6 = (v6 % ((-880) | 1))
            for fi31 := 0; fi31 < 9; fi31++ {
                fmt.Println(v1, v9, fi31)
                var v32 int64 = v6
                _ = fi31
                _ = v32
            }
            var v33 int32 = v3
            _ = fi30
            _ = v33
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}

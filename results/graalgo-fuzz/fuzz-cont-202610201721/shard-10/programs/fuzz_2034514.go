package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 uint8; F1 string; F2 uint }
func genFn0(p0 int64) int64 {
    var v7 bool = (32767 > 870)
    var v8 uint8 = 128
    v8 += v8
    if (65535 == 255) {
        for fi9 := 0; fi9 < 20; fi9++ {
            v10 := []int{fi9, (fi9 * (^(fi9)))}
            var v11 bool = (uint16(v8) != 65535)
            _ = fi9
            _ = v10
            _ = v11
        }
        p0 = p0
    } else {
        v12 := []int{42, 1, -730}
        _ = len(v12)
        _ = v12
    }
    _ = p0
    _ = v7
    _ = v8
    return 693
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = 1.5
    var v3 bool = ((161 >= 4294967295) && (1000 != -1))
    v4 := make(map[int]bool)
    v5 := []uint32{2, 42}
    var v6 int16 = 1
    var v13 float64 = 0.0
    v14 := GenS1{F0: (-(v6)), F1: (-2.25 != (float64(((v2 * ((v2 + -2.25) * v2)) - float32(1.5))) - float64(v2))), F2: 65535}
    v13 += float64(0.1)
    for fi15 := 0; fi15 < 8; fi15++ {
        for fi16 := 0; fi16 < 15; fi16++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = len(v5)
            v3 = (82 != 1)
            _ = fi16
        }
        if v3 {
            v13 = float64((v2 + float32((((v13 - 0.1) / float64(v2)) / float64(((1.5 + 0.1) + v2))))))
            var v17 uint8 = 128
            v6 -= v6
            _ = v17
        } else {
            var v18 float32 = v2
            if (v3 || v3) {
                if v3 {
                    for fi19 := 0; fi19 < 7; fi19++ {
                        switch "hello" {
                        case "":
                            v18 = 0.1
                        case "hello":
                            var v20 uint64 = 541
                            var v21 int64 = -119
                            v6 = (v6 | v6)
                            var v22 int16 = (v6 >> ((44) & 63))
                            v6 += v22
                            _ = v20
                            _ = v21
                            _ = v22
                        case "a":
                            v23 := GenS2{F0: (fi15 & (fi15 ^ (fi15 >> ((142) & 63))))}
                            _ = v23
                        }
                        _ = fi19
                    }
                    var v24 uint16 = 545
                    _ = v24
                } else {
                    var v25 int32 = -587
                    _ = v25
                }
                v26 := GenS2{F0: 1}
                _ = v26
            }
            fmt.Println(v13, v6, v3)
            _ = v18
        }
        v1[42] = 42
        _ = fi15
    }
    var v27 float64 = v13
    v28 := v4[-962]
    fmt.Println(v3, v28, v28)
    v6 += -640
    fmt.Println(v13, v2, v13)
    fmt.Println(v27, v3)
    v28 = ((("foo bar" + "foo bar") < (("hello" + "\n\t") + "foo bar")) && true)
    if (!v3) {
        v29 := GenS1{F0: 1000, F1: v28, F2: 0}
        v13 = (3.14159 * (1.5 - v27))
        fmt.Println(v6, v13, v2)
        _ = v29
    } else {
        _ = len(v4)
        switch ("foo bar" + "foo bar") {
        case "foo bar":
            fmt.Println(v27)
        case "\n\t":
            for fi30 := 0; fi30 < 9; fi30++ {
                for fi31 := 0; fi31 < 15; fi31++ {
                    fmt.Println(v3, fi31, fi31)
                    v32 := GenS3{F0: 0, F1: 18446744073709551615, F2: 2}
                    v33 := GenS1{F0: (v6 % ((-1) | 1)), F1: v28, F2: 65535}
                    var v34 uint32 = 955
                    _ = fi31
                    _ = v32
                    _ = v33
                    _ = v34
                }
                _ = fi30
            }
        }
        fmt.Println(v3, v6, v28)
    }
    v28 = v3
    switch 32767 {
    case 776:
        v13 -= v13
    case 489:
        for fi35 := 0; fi35 < 7; fi35++ {
            v36 := GenS4{F0: 452, F1: genFn0(-59)}
            _ = fi35
            _ = v36
        }
    }
    v28 = v28
    if v28 {
        if (true || v3) {
            for fi37 := 0; fi37 < 14; fi37++ {
                var v38 uint32 = 42
                v4[fi37] = v28
                for fi39 := 0; fi39 < 17; fi39++ {
                    if (!(100 >= 16)) {
                        v40 := GenS5{F0: 255, F1: "foo bar", F2: 255}
                        delete(v4, fi37)
                        _ = v40
                    } else {
                        var v41 int64 = 9223372036854775807
                        _ = v41
                    }
                    var v42 int8 = 1
                    v43 := GenS2{F0: (fi37 % ((0) | 1))}
                    v44 := v1[4294967295]
                    _ = fi39
                    _ = v42
                    _ = v43
                    _ = v44
                }
                _ = fi37
                _ = v38
            }
            fmt.Println(v13, v3)
        }
    }
    switch (("hello" + "hello") + ("" + ("hello" + (("\n\t" + "hello") + "foo bar")))) {
    case "a":
        v28 = v28
    case "foo bar":
        var v45 bool = v28
        _ = v45
    }
    if true {
        v2 += v2
        var v46 uint = 520
        for fi47 := 0; fi47 < 17; fi47++ {
            v48 := GenS5{F0: 128, F1: "foo bar", F2: v46}
            for fi49 := 0; fi49 < 12; fi49++ {
                var v50 uint32 = uint32(v46)
                _ = fi49
                _ = v50
            }
            _ = fi47
            _ = v48
        }
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

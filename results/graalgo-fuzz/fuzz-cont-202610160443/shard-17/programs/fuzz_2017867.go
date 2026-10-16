package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint16; F1 int8; F2 int8 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 uint; F1 int32 }
func genFn0() uint32 {
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 int16 = 469
        var v8 uint16 = 0
        v7 = (v7 % ((v7) | 1))
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 725}
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 int = (fi10 & -393)
        v11 -= fi10
        fmt.Println(v11, fi10)
        _ = fi10
        _ = v11
    }
    _ = v9
    return 2
}
func genFn1(p0 uint16, p1 uint) uint32 {
    p1 = (p1 & p1)
    v12 := GenS2{F0: ("a" + "\n\t"), F1: 495}
    v13 := GenS3{F0: p0, F1: -121, F2: 36}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 41
}

func main() {
    var v1 uint = 457
    var v2 uint8 = 2
    var v3 uint16 = 799
    var v4 uint16 = v3
    var v5 int8 = -1
    v14 := GenS1{F0: 249}
    for fi15 := 0; fi15 < 16; fi15++ {
        v16 := GenS4{F0: (v5 ^ v5), F1: 0}
        _ = fi15
        _ = v16
    }
    v17 := []int{-9223372036854775808}
    fmt.Println(v5, v3)
    if true {
        if (!(!(true && false))) {
            var v18 int16 = 32767
            v2 = (v2 << ((v1) & 63))
            v19 := GenS2{F0: "a", F1: int16(v3)}
            _ = v18
            _ = v19
        } else {
            var v20 uint8 = v2
            v21 := GenS2{F0: ("foo bar" + (((("\n\t" + "foo bar") + "hello") + "foo bar") + "")), F1: 921}
            _ = v20
            _ = v21
        }
    }
    if (v5 <= (v5 >> ((v1) & 63))) {
        if false {
            for fi22 := 0; fi22 < 14; fi22++ {
                if false {
                    v3 -= (v4 * v4)
                    var v23 bool = ((v5 / (((v5 / ((v5) | 1))) | 1)) > (v5 % ((v5) | 1)))
                    _ = v23
                }
                var v24 uint32 = 1
                v2 += (v2 % ((v2) | 1))
                _ = fi22
                _ = v24
            }
            for fi25 := 0; fi25 < 4; fi25++ {
                var v26 int8 = v5
                _ = fi25
                _ = v26
            }
            v27 := GenS5{F0: 42, F1: 338}
            _ = v27
        } else {
            v4 += v4
            if len(v17) > 0 {
                v17[0] = -2147483648
            }
            for fi28 := 0; fi28 < 6; fi28++ {
                var v29 bool = true
                _ = fi28
                _ = v29
            }
        }
        var v30 uint8 = 16
        _ = v30
    } else {
        var v31 string = "foo bar"
        v17 = append(v17, 632)
        v32 := GenS5{F0: v1, F1: -2147483648}
        _ = v31
        _ = v32
    }
    if false {
        for fi33 := 0; fi33 < 14; fi33++ {
            fmt.Println(v1, v5, v1)
            _ = fi33
        }
        if true {
            if len(v17) > 0 {
                v17[0] = 652
            }
            if (true && true) {
                v4 += v3
            }
        }
    } else {
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
        if (false || (!false)) {
            switch -9223372036854775808 {
            case 2147483647:
                v34 := GenS5{F0: (v1 * 320), F1: 131}
                _ = v34
            default:
                for fi35 := 0; fi35 < 11; fi35++ {
                    v17 = append(v17, -743)
                    _ = fi35
                }
            }
        } else {
            var v36 uint64 = 4294967295
            _ = v36
        }
        v5 = v5
    }
    v3 += (v3 + (^(v3)))
    v3 = v4
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

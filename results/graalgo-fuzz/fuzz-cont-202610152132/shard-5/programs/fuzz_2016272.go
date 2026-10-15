package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 bool }
type GenS2 struct { F0 int; F1 int8; F2 int64 }
func genFn0(p0 uint, p1 bool) int32 {
    p1 = (!p1)
    p1 = p1
    var v6 uint = p0
    v7 := GenS1{F0: 255, F1: 1, F2: p1}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 364
}
func genFn1(p0 int16, p1 bool) int64 {
    v8 := []int{-2147483648}
    var v9 int16 = p0
    p0 = (p0 ^ -496)
    v10 := GenS1{F0: 21, F1: -686, F2: (false || (!(621 != 1)))}
    var v11 string = ("foo bar" + (("a" + "a") + "\n\t"))
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 173
}

func main() {
    var v1 float32 = (0.1 / -2.25)
    var v2 int64 = -1
    var v3 float32 = 0.0
    var v4 int64 = 9223372036854775807
    var v5 uint = 0
    fmt.Println(v3, v5, v5)
    if false {
        var v12 int32 = -727
        _ = v12
    } else {
        for fi13 := 0; fi13 < 17; fi13++ {
            switch -9223372036854775808 {
            case 54:
                var v14 int = (fi13 + (fi13 % (((-(fi13))) | 1)))
                _ = v14
            }
            _ = fi13
        }
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 int32 = 1
            switch (-(fi15)) {
            case 518:
                v5 -= v5
            }
            switch "" {
            case "foo bar":
                v1 += (float32(0.1) + float32(3.14159))
            default:
                fmt.Println(v1, v2)
            }
            _ = fi15
            _ = v16
        }
        if (!(!(0 > 944))) {
            if ((!false) && false) {
                if false {
                    if (!(71 <= 108)) {
                        var v17 int16 = -32768
                        switch "foo bar" {
                        case "":
                            for fi18 := 0; fi18 < 16; fi18++ {
                                fmt.Println(v4)
                                var v19 int = (fi18 | 769)
                                v17 += (v17 & (v17 >> (((^(v5))) & 63)))
                                _ = fi18
                                _ = v19
                            }
                        case "hello":
                            v4 -= (v2 - -1)
                        }
                        _ = v17
                    }
                    if (!false) {
                        var v20 int = 1
                        _ = v20
                    }
                    fmt.Println(v2, v3, v1)
                } else {
                    v2 += v4
                    v21 := GenS1{F0: 255, F1: 127, F2: false}
                    v3 = v3
                    _ = v21
                }
                v1 = (((v1 + 0.0) * (1e10 * float32(1.5))) - (float32(0.1) / (v1 + (v3 / (((v3 / v3) - v1) / 0.1)))))
            }
            v2 = (v2 + 693)
            if true {
                v4 = (^(v4))
            }
        } else {
            for fi22 := 0; fi22 < 6; fi22++ {
                v23 := GenS1{F0: 128, F1: (-(fi22)), F2: false}
                fmt.Println(v2, v5)
                _ = fi22
                _ = v23
            }
        }
    }
    v2 = v2
    v24 := GenS2{F0: 540, F1: 0, F2: -1}
    v25 := GenS2{F0: 32768, F1: 127, F2: -349}
    v3 -= (v1 + v1)
    if true {
        var v26 uint16 = 65535
        _ = v26
    }
    for fi27 := 0; fi27 < 17; fi27++ {
        if (!true) {
            if true {
                var v28 int = 395
                v29 := GenS2{F0: 127, F1: -128, F2: (^(v4))}
                fmt.Println(v3)
                _ = v28
                _ = v29
            } else {
                v5 -= v5
            }
            var v30 bool = ((!(-9223372036854775808 >= v2)) && ((fi27 & -9223372036854775808) < (fi27 | (-(fi27)))))
            v31 := []int{fi27, (fi27 * 30)}
            _ = v30
            _ = v31
        }
        for fi32 := 0; fi32 < 6; fi32++ {
            switch "hello" {
            case "":
                v2 = int64(v5)
            case "a":
                v33 := []int{(fi27 + fi27)}
                _ = v33
            case "hello":
                var v34 uint64 = 0
                _ = v34
            }
            v1 = v3
            _ = fi32
        }
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}

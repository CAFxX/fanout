package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int; F1 uint32; F2 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 uint32; F1 int8; F2 bool }
func genFn0(p0 int16, p1 uint32) uint8 {
    v4 := []int{int(p1), 4294967295}
    v5 := GenS1{F0: 0, F1: -2147483648}
    var v6 bool = true
    for fi7 := 0; fi7 < 12; fi7++ {
        fmt.Println(p1, p1, v6)
        v6 = ((^(fi7)) < 127)
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 127
}
func genFn1() uint16 {
    v8 := GenS2{F0: -386, F1: 4294967295, F2: 4294967295}
    v9 := []int{-72, 2147483647, 9223372036854775807}
    var v10 uint8 = 215
    v10 = v10
    _ = v10
    _ = v8
    _ = v9
    return 32767
}

func main() {
    var v1 int8 = 3
    var v2 string = ("\n\t" + "foo bar")
    var v3 int32 = 0
    if false {
        var v11 uint64 = 4294967295
        fmt.Println(v2, v2, v3)
        _ = v11
    }
    if true {
        fmt.Println(v1)
        if (!true) {
            v2 = "hello"
        }
        if (128 <= -868) {
            v12 := []int{4294967295}
            v12 = append(v12, -858)
            for fi13 := 0; fi13 < 16; fi13++ {
                v14 := GenS1{F0: 73, F1: (-(v3))}
                v3 = v3
                _ = fi13
                _ = v14
            }
            _ = v12
        }
    } else {
        v15 := []int{42, -1, 32767, 128}
        v16 := GenS3{F0: 13}
        v17 := GenS4{F0: v3}
        _ = v15
        _ = v16
        _ = v17
    }
    v18 := GenS2{F0: -9223372036854775808, F1: 1, F2: 18446744073709551615}
    if (!true) {
        v19 := GenS4{F0: v3}
        var v20 float32 = 1.5
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        for fi22 := 0; fi22 < 16; fi22++ {
            v23 := []int{-454}
            _ = fi22
            _ = v23
        }
        if true {
            v24 := []int{}
            _ = v24
        }
        v25 := GenS5{F0: -1, F1: 255}
        _ = fi21
        _ = v25
    }
    var v26 uint8 = 42
    v26 = (v26 ^ (^(v26)))
    if true {
        for fi27 := 0; fi27 < 9; fi27++ {
            v3 = -981
            v28 := []int{}
            v3 = (-(v3))
            _ = fi27
            _ = v28
        }
        v2 = v2
        v3 = (v3 & v3)
    } else {
        var v29 uint16 = genFn1()
        v30 := GenS6{F0: 4294967295, F1: 42, F2: (v3 >= v3)}
        if true {
            if false {
                v29 += (v29 << ((255) & 63))
                if true {
                    if false {
                        if (1000 <= -1) {
                            if true {
                                var v31 int8 = 42
                                var v32 int16 = 0
                                _ = v31
                                _ = v32
                            }
                        }
                        var v33 int32 = v3
                        _ = v33
                    }
                    switch 4294967295 {
                    case -160:
                        v1 = -1
                    case 32768:
                        if true {
                            v1 += v1
                            if (-1 == -305) {
                                v3 -= (v3 * (v3 % ((v3) | 1)))
                            }
                        } else {
                            if ((0.0 + 3.14159) < 0.1) {
                                v26 = 42
                            }
                            switch (v2 + "hello") {
                            case "":
                                v34 := []int{}
                                _ = v34
                            case "hello":
                                v35 := []int{797, 127, -63, -9223372036854775808}
                                _ = v35
                            default:
                                v36 := []int{32768, -9223372036854775808, -9223372036854775808, -9223372036854775808}
                                _ = v36
                            }
                        }
                    default:
                        v37 := []int{571}
                        _ = v37
                    }
                }
            }
        }
        _ = v29
        _ = v30
    }
    switch v2 {
    case "foo bar":
        var v38 uint16 = 255
        _ = v38
    case "\n\t":
        fmt.Println(v1)
    case "":
        for fi39 := 0; fi39 < 20; fi39++ {
            v26 = (v26 % (((v26 * (^(v26)))) | 1))
            fmt.Println(v2, v26)
            v3 = (v3 | (v3 & v3))
            _ = fi39
        }
    default:
        var v40 int64 = -724
        _ = v40
    }
    if false {
        fmt.Println(v26, v3)
        var v41 string = "foo bar"
        _ = v41
    }
    v42 := GenS3{F0: v26}
    fmt.Println(v26, v3, v1)
    var v43 uint = 780
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v42
    _ = v43
}

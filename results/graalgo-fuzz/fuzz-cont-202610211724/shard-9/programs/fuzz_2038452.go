package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 bool; F1 int16; F2 bool }
type GenS3 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 string; F1 uint32 }
type GenS6 struct { F0 int }

func main() {
    var v1 int8 = -128
    var v2 uint8 = 210
    var v3 int = 4294967295
    v3 = (v3 & (-(v3)))
    fmt.Println(v1, v1, v2)
    v4 := GenS1{F0: -1, F1: true}
    var v5 uint64 = 730
    if false {
        for fi6 := 0; fi6 < 12; fi6++ {
            v7 := GenS2{F0: true, F1: -761, F2: (!false)}
            v8 := GenS2{F0: true, F1: 32767, F2: false}
            _ = fi6
            _ = v7
            _ = v8
        }
    } else {
        v5 -= (v5 | (^(v5)))
        v9 := GenS2{F0: false, F1: 32767, F2: ((!(!true)) || false)}
        v10 := GenS3{F0: 0, F1: 179, F2: v2}
        _ = v10
        _ = v9
    }
    if true {
        switch (v3 % (((v3 / (((v3 % ((4294967295) | 1))) | 1))) | 1)) {
        case 128:
            for fi11 := 0; fi11 < 9; fi11++ {
                switch "foo bar" {
                case "foo bar":
                    v12 := []int{(v3 + (v3 >> ((0) & 63)))}
                    _ = v12
                case "hello":
                    fmt.Println(fi11, fi11)
                default:
                    v13 := []int{(v3 * (fi11 / ((v3) | 1))), (fi11 | (^(v3)))}
                    _ = v13
                }
                var v14 bool = true
                _ = fi11
                _ = v14
            }
        case 2147483647:
            var v15 int32 = -2147483648
            _ = v15
        case -222:
            if (!((v1 * (-(v1))) >= (v1 - 127))) {
                v16 := GenS4{F0: 42, F1: "hello"}
                v17 := GenS3{F0: 1, F1: 42, F2: v2}
                _ = v16
                _ = v17
            }
        default:
            v18 := []int{(v3 & (v3 << ((1) & 63)))}
            _ = v18
        }
        v3 -= 239
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        if (3.14159 != (1e10 + (1e10 * 0.0))) {
            if true {
                v1 -= v1
                if false {
                    for fi20 := 0; fi20 < 5; fi20++ {
                        var v21 uint8 = v2
                        v22 := GenS2{F0: (93 == (v1 % (((v1 << ((1) & 63))) | 1))), F1: 426, F2: false}
                        v23 := []int{2147483647}
                        var v24 int64 = 42
                        v3 = v3
                        _ = fi20
                        _ = v21
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    switch 871 {
                    case 4294967295:
                        for fi25 := 0; fi25 < 9; fi25++ {
                            v26 := GenS4{F0: -2147483648, F1: ("" + "foo bar")}
                            v27 := []int{(^(fi19)), fi19, (v3 / ((2147483647) | 1))}
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                    default:
                        if false {
                            v28 := []int{(^(v3)), (v3 % (((v3 >> ((274) & 63))) | 1)), v3}
                            fmt.Println(fi19, fi19, v2)
                            _ = v28
                        }
                    }
                } else {
                    v1 -= (-(v1))
                    v29 := []int{}
                    _ = v29
                }
                v3 = (v3 - (-(fi19)))
            } else {
                v2 += v2
                for fi30 := 0; fi30 < 10; fi30++ {
                    v31 := []int{(^(fi30)), 127, fi19}
                    _ = fi30
                    _ = v31
                }
            }
        } else {
            if true {
                for fi32 := 0; fi32 < 8; fi32++ {
                    fmt.Println(v3, v2)
                    v33 := []int{}
                    _ = fi32
                    _ = v33
                }
                var v34 uint = 255
                _ = v34
            } else {
                var v35 int16 = 32767
                if false {
                    var v36 uint16 = 32767
                    if false {
                        v37 := []int{}
                        v38 := GenS3{F0: v36, F1: 990, F2: 2}
                        _ = v37
                        _ = v38
                    } else {
                        switch v3 {
                        case -727:
                            v2 -= uint8(v36)
                        case 128:
                            for fi39 := 0; fi39 < 4; fi39++ {
                                v40 := GenS5{F0: "", F1: uint32(v36)}
                                _ = fi39
                                _ = v40
                            }
                        case -1:
                            v41 := []int{v3, (v3 * (fi19 - -2147483648)), fi19}
                            _ = v41
                        default:
                            var v42 int64 = 42
                            _ = v42
                        }
                        v36 -= v36
                        v43 := []int{(fi19 - (fi19 / (((fi19 & v3)) | 1)))}
                        _ = v43
                    }
                    _ = v36
                }
                v35 += 32767
                _ = v35
            }
        }
        v44 := GenS6{F0: 447}
        _ = fi19
        _ = v44
    }
    v3 -= (-(v3))
    var v45 uint8 = v2
    v46 := GenS5{F0: ("a" + "hello"), F1: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
}

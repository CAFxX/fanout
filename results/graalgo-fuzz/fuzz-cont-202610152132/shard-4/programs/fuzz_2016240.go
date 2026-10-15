package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint8, p1 uint) uint8 {
    for fi4 := 0; fi4 < 18; fi4++ {
        var v5 int64 = 1
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: (p0 + p0), F1: 1, F2: 977}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    return (p0 * 169)
}

func main() {
    var v1 uint = 328
    var v2 int = 0
    var v3 bool = (128 == 1)
    v8 := GenS1{F0: 0, F1: 383, F2: 32767}
    v9 := GenS1{F0: 1, F1: 1, F2: 344}
    switch ("a" + "") {
    case "\n\t":
        var v10 int16 = 1000
        _ = v10
    default:
        v2 -= (v2 & (-(v2)))
    }
    v1 -= 1
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 uint = (v1 - v1)
        v1 = v12
        var v13 uint8 = 178
        _ = fi11
        _ = v12
        _ = v13
    }
    v1 -= 1
    v14 := GenS2{F0: v3}
    v1 -= v1
    var v15 uint16 = 1
    var v16 uint16 = 255
    v1 = 239
    v17 := GenS1{F0: uint8(v1), F1: (v16 + 336), F2: v15}
    if (genFn0(255, 4294967295) <= 2) {
        switch "foo bar" {
        case "\n\t":
            if true {
                for fi18 := 0; fi18 < 13; fi18++ {
                    for fi19 := 0; fi19 < 11; fi19++ {
                        v20 := GenS1{F0: 1, F1: (^(v15)), F2: v15}
                        if (-416 <= (fi19 ^ 4294967295)) {
                            v21 := []int{fi19, int(v16), -1, fi18}
                            _ = v21
                        } else {
                            if v3 {
                                if (!((1e10 > (((0.0 + 0.0) + ((1e10 / 0.1) * (3.14159 - 1.5))) / (-2.25 / 1e10))) || v3)) {
                                    var v22 int64 = 9223372036854775807
                                    _ = v22
                                } else {
                                    switch (v2 * -1) {
                                    case -919:
                                        fmt.Println(v2, v15, v3)
                                    case 973:
                                        v16 += 255
                                        if v3 {
                                            v3 = v3
                                            var v23 uint8 = 1
                                            _ = v23
                                        } else {
                                            switch -2147483648 {
                                            case 32768:
                                                fmt.Println(v16, v16, v2)
                                            case 127:
                                                v24 := []int{}
                                                _ = v24
                                            }
                                        }
                                    case -9223372036854775808:
                                        fmt.Println(fi19, fi18, v3)
                                    default:
                                        v25 := []int{1, -9223372036854775808}
                                        _ = v25
                                    }
                                }
                                var v26 bool = true
                                _ = v26
                            } else {
                                v27 := []int{32767, -470, 0}
                                switch fi19 {
                                case 9223372036854775807:
                                    var v28 int32 = -726
                                    _ = v28
                                default:
                                    v29 := GenS1{F0: 127, F1: (v16 / (((v16 * (v16 & 732))) | 1)), F2: (v15 & (^(v15)))}
                                    _ = v29
                                }
                                var v30 uint64 = 4294967295
                                _ = v27
                                _ = v30
                            }
                            v15 = 2
                            if v3 {
                                v1 -= (v1 * (v1 + (v1 ^ 2)))
                            }
                        }
                        _ = fi19
                        _ = v20
                    }
                    var v31 int64 = 602
                    v32 := []int{(v2 % (((-(fi18))) | 1)), fi18}
                    _ = fi18
                    _ = v31
                    _ = v32
                }
                var v33 int64 = 373
                _ = v33
            }
        }
        v3 = (-128 <= -17)
        for fi34 := 0; fi34 < 12; fi34++ {
            v35 := []int{fi34}
            v15 += (v15 * 506)
            _ = fi34
            _ = v35
        }
    } else {
        for fi36 := 0; fi36 < 12; fi36++ {
            if (100 < 1) {
                if false {
                    var v37 uint8 = 128
                    _ = v37
                }
            }
            _ = fi36
        }
    }
    var v38 uint8 = 42
    for fi39 := 0; fi39 < 20; fi39++ {
        v38 += 128
        _ = fi39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v38
    _ = v8
    _ = v9
}

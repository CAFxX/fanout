package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint8, p1 uint8) uint64 {
    if true {
        var v4 uint = 0
        _ = v4
    }
    fmt.Println(p1)
    v5 := []int{}
    switch 0 {
    case -270:
        if true {
            for fi6 := 0; fi6 < 9; fi6++ {
                switch fi6 {
                case 128:
                    for fi7 := 0; fi7 < 17; fi7++ {
                        v5 = append(v5, fi7)
                        if (!false) {
                            if (488 != 682) {
                                _ = len(v5)
                            }
                            v8 := GenS1{F0: 2147483647, F1: 18446744073709551615, F2: -118}
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = v8
                        } else {
                            if len(v5) > 0 {
                                v5[0] = (fi7 / (((fi7 - (^(fi6)))) | 1))
                            }
                            if len(v5) > 0 {
                                v5[0] = 42
                            }
                        }
                        _ = fi7
                    }
                }
                var v9 string = ""
                _ = fi6
                _ = v9
            }
        }
    case -500:
        fmt.Println(p0, p1)
    default:
        if false {
            p1 = (p0 * 55)
            for fi10 := 0; fi10 < 10; fi10++ {
                switch 445 {
                case 4294967295:
                    _ = len(v5)
                case -303:
                    if true {
                        var v11 uint = 0
                        _ = v11
                    }
                case -1:
                    _ = len(v5)
                }
                _ = fi10
            }
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return 1
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint16 = 42
    var v3 uint64 = uint64(v2)
    if true {
        var v12 float32 = 1e10
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 int32 = 2147483647
            _ = fi13
            _ = v14
        }
        _ = v12
    }
    if (false && true) {
        v3 = v1
        v15 := GenS2{F0: 0}
        v2 -= v2
        _ = v15
    } else {
        v1 = v3
        var v16 uint64 = (v1 - v3)
        _ = v16
    }
    fmt.Println(v2)
    v2 -= 78
    var v17 uint = 926
    v18 := GenS3{F0: 127, F1: (v2 | (v2 + 1))}
    for fi19 := 0; fi19 < 6; fi19++ {
        if true {
            v20 := []int{}
            for fi21 := 0; fi21 < 1; fi21++ {
                switch ("\n\t" + "foo bar") {
                case "a":
                    var v22 uint32 = 222
                    _ = v22
                case "\n\t":
                    if true {
                        switch "hello" {
                        case "\n\t":
                            v23 := GenS1{F0: 42, F1: 625, F2: 32767}
                            _ = v23
                        case "":
                            v20 = append(v20, (-(fi19)))
                        case "foo bar":
                            var v24 uint64 = v3
                            _ = v24
                        }
                        if len(v20) > 0 {
                            v20[0] = (^(fi21))
                        }
                        var v25 uint = v17
                        _ = v25
                    }
                    var v26 uint8 = 111
                    v27 := GenS3{F0: (v26 ^ uint8(v2)), F1: v2}
                    _ = v26
                    _ = v27
                default:
                    if (!(993 > v17)) {
                        var v28 int64 = -9223372036854775808
                        _ = v28
                    }
                }
                _ = fi21
            }
            v1 += v3
            _ = v20
        } else {
            if (false || (-32768 < -1)) {
                v1 = 18446744073709551615
            } else {
                v1 = genFn0(2, 255)
                if true {
                    v1 += (v3 * genFn0(87, 2))
                    var v29 uint32 = 0
                    v30 := []int{(fi19 >> ((uint(v1)) & 63)), fi19, int(v3)}
                    _ = v29
                    _ = v30
                } else {
                    for fi31 := 0; fi31 < 6; fi31++ {
                        v32 := []int{(fi31 << (((v17 | uint(v1))) & 63)), fi19, -406, (^(fi31))}
                        if false {
                            v1 = 685
                            v17 -= (^(v17))
                        } else {
                            var v33 float32 = 0.0
                            var v34 int64 = -180
                            var v35 int = fi31
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                        _ = fi31
                        _ = v32
                    }
                    var v36 float32 = (0.1 * 0.1)
                    v17 += v17
                    _ = v36
                }
                for fi37 := 0; fi37 < 17; fi37++ {
                    v38 := []int{(fi19 / (((fi37 << (((v17 | v17)) & 63))) | 1))}
                    v17 = (v17 << ((255) & 63))
                    fmt.Println(v3, fi37, fi37)
                    _ = fi37
                    _ = v38
                }
            }
            v17 += v17
            fmt.Println(v1, v3)
        }
        switch fi19 {
        case 42:
            switch "foo bar" {
            case "":
                var v39 float32 = 0.0
                _ = v39
            }
        case 2147483647:
            if true {
                v40 := GenS4{F0: true, F1: 1}
                v1 -= 4294967295
                _ = v40
            } else {
                var v41 int8 = -117
                v42 := GenS4{F0: (500 >= 902), F1: 127}
                _ = v41
                _ = v42
            }
        }
        _ = fi19
    }
    v1 = 123
    v43 := GenS1{F0: -1, F1: v17, F2: 1000}
    var v44 int32 = 0
    var v45 uint64 = (v1 - (^(v1)))
    v45 -= (v45 ^ (v1 << ((18446744073709551615) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v43
    _ = v44
    _ = v45
}

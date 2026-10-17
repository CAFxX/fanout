package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int8; F1 uint16 }

func main() {
    var v1 int16 = 1
    var v2 uint = 832
    v3 := make(map[uint32]string)
    var v4 int8 = 0
    fmt.Println(v4, v4, v1)
    switch -9223372036854775808 {
    case -843:
        for fi5 := 0; fi5 < 6; fi5++ {
            for fi6 := 0; fi6 < 4; fi6++ {
                _ = len(v3)
                v7 := GenS1{F0: ("\n\t" + ""), F1: v1}
                _ = fi6
                _ = v7
            }
            var v8 uint32 = 65535
            _ = fi5
            _ = v8
        }
    case 2147483647:
        if (4294967295 >= 661) {
            var v9 uint16 = 522
            _ = v9
        }
    case 32767:
        v10 := v3[1]
        _ = v10
    default:
        v11 := GenS1{F0: "foo bar", F1: (v1 - (v1 << (((v2 % ((v2) | 1))) & 63)))}
        _ = v11
    }
    fmt.Println(v1, v1, v2)
    if true {
        for fi12 := 0; fi12 < 12; fi12++ {
            var v13 uint16 = 50
            if true {
                for fi14 := 0; fi14 < 14; fi14++ {
                    switch "" {
                    case "foo bar":
                        fmt.Println(v4, v13)
                    default:
                        if ((1e10 / 1.0) <= 3.14159) {
                            var v15 float64 = -2.25
                            var v16 string = "a"
                            _ = v15
                            _ = v16
                        } else {
                            v1 -= (v1 - 1000)
                            var v17 uint16 = (^(v13))
                            fmt.Println(fi14, fi12, v4)
                            if false {
                                v18 := []int{}
                                _ = v18
                            } else {
                                if true {
                                    if false {
                                        v19 := []int{fi12, (-(fi12)), 32768, fi14}
                                        if (((v2 | (v2 % ((v2) | 1))) <= 1) && (-2.25 <= (1.5 * ((0.1 * -2.25) - (0.0 / 1.5))))) {
                                            v20 := GenS1{F0: "", F1: v1}
                                            var v21 uint32 = 255
                                            if len(v19) >= 2 {
                                                v19 = v19[1:len(v19)]
                                            }
                                            _ = v20
                                            _ = v21
                                        } else {
                                            var v22 string = (("" + "\n\t") + ("a" + "hello"))
                                            _ = v22
                                        }
                                        if true {
                                            v2 -= 42
                                        } else {
                                            v2 = (v2 * v2)
                                        }
                                        switch ("foo bar" + "hello") {
                                        case "a":
                                            fmt.Println(v1)
                                        case "hello":
                                            _ = len(v19)
                                        }
                                        switch "" {
                                        case "hello":
                                            fmt.Println(v17)
                                        case "a":
                                            v23 := GenS1{F0: "hello", F1: 261}
                                            _ = v23
                                        case "":
                                            v24 := GenS1{F0: "\n\t", F1: v1}
                                            _ = v24
                                        }
                                        _ = v19
                                    } else {
                                        var v25 uint8 = 1
                                        fmt.Println(fi12, fi12)
                                        if true {
                                            var v26 int8 = v4
                                            var v27 uint16 = v17
                                            _ = v26
                                            _ = v27
                                        }
                                        _ = v25
                                    }
                                    v17 -= (v13 - (v17 >> ((uint(v13)) & 63)))
                                }
                                fmt.Println(v4)
                                v28 := v3[65535]
                                _ = v28
                            }
                            fmt.Println(v4)
                            _ = v17
                        }
                    }
                    fmt.Println(v2)
                    v4 -= (-(v4))
                    v1 += (v1 - v1)
                    v1 -= v1
                    _ = fi14
                }
                if (-760 != -9223372036854775808) {
                    v29 := GenS1{F0: ("hello" + "a"), F1: (v1 & (^(v1)))}
                    _ = v29
                }
                v30 := []int{int(v1)}
                _ = v30
            }
            var v31 bool = true
            _ = fi12
            _ = v13
            _ = v31
        }
        v2 = 2
    }
    delete(v3, 2147483647)
    v32 := GenS2{F0: (v4 * (v4 << ((v2) & 63))), F1: 32767}
    v2 = 65535
    v33 := v3[2147483648]
    for fi34 := 0; fi34 < 14; fi34++ {
        v35 := GenS2{F0: (v4 / ((v4) | 1)), F1: 249}
        _ = fi34
        _ = v35
    }
    v1 = -1
    v4 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}

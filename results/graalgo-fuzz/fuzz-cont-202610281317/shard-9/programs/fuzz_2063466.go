package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint8 }
func genFn0(p0 uint) bool {
    p0 = 1
    p0 -= 42
    _ = p0
    return ((127 >= int8(p0)) || (-1 == 42))
}
func genFn1() uint32 {
    var v4 int = 127
    for fi5 := 0; fi5 < 14; fi5++ {
        for fi6 := 0; fi6 < 3; fi6++ {
            if false {
                var v7 int32 = 1
                v8 := GenS1{F0: 0, F1: (v7 % ((v7) | 1)), F2: 2}
                v9 := []int{v4, (fi6 * 127), 658}
                _ = v7
                _ = v8
                _ = v9
            } else {
                v4 = fi5
                var v10 int8 = 1
                switch "a" {
                case "a":
                    v10 += 42
                default:
                    fmt.Println(fi5)
                }
                _ = v10
            }
            v11 := GenS1{F0: 1, F1: 814, F2: 31}
            if false {
                var v12 int16 = 42
                var v13 int64 = 9223372036854775807
                v14 := []int{}
                _ = v12
                _ = v13
                _ = v14
            }
            _ = fi6
            _ = v11
        }
        _ = fi5
    }
    var v15 uint32 = 374
    v16 := []int{872, v4, v4}
    v17 := GenS1{F0: -148, F1: 108, F2: 214}
    _ = v15
    _ = v16
    _ = v17
    _ = v4
    return (^(v15))
}

func main() {
    var v1 uint8 = 1
    var v2 uint16 = 0
    var v3 int64 = -9223372036854775808
    if true {
        v2 = (v2 + (^(v2)))
        v3 -= v3
    } else {
        v18 := []int{}
        v1 += 2
        switch -2147483648 {
        case 1:
            var v19 int8 = 127
            _ = v19
        }
        _ = v18
    }
    var v20 string = "hello"
    v21 := []int{128, -1, 946, -954}
    for fi22 := 0; fi22 < 15; fi22++ {
        if len(v21) > 0 {
            v21[0] = -9223372036854775808
        }
        if true {
            fmt.Println(v3, fi22)
        }
        _ = fi22
    }
    if ((!((v1 | 255) < 128)) || (127 != -128)) {
        v2 += (v2 * v2)
    }
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    fmt.Println(v1, v1, v1)
    switch v20 {
    case "\n\t":
        if (!false) {
            for fi23 := 0; fi23 < 4; fi23++ {
                var v24 int8 = 15
                fmt.Println(v2, v2)
                if (-32768 == 1) {
                    if (250 != 2147483647) {
                        var v25 int64 = v3
                        var v26 float32 = 1e10
                        _ = v25
                        _ = v26
                    }
                }
                _ = fi23
                _ = v24
            }
        }
    default:
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v3, v2, v20)
        case "a":
            v1 = (v1 >> ((255) & 63))
        case "foo bar":
            for fi27 := 0; fi27 < 4; fi27++ {
                switch fi27 {
                case -662:
                    fmt.Println(v3, v3)
                case 128:
                    if (!(!(1e10 <= float64(0.0)))) {
                        v1 += v1
                    } else {
                        for fi28 := 0; fi28 < 18; fi28++ {
                            var v29 uint = 18446744073709551615
                            if (!(-662 >= -730)) {
                                v3 -= (v3 | -9223372036854775808)
                                fmt.Println(fi28, v2)
                                if len(v21) >= 2 {
                                    v21 = v21[1:len(v21)]
                                }
                                fmt.Println(v20, fi27)
                            } else {
                                v21 = append(v21, (fi27 % ((fi28) | 1)))
                                fmt.Println(v2, fi28)
                            }
                            _ = fi28
                            _ = v29
                        }
                        var v30 string = (v20 + ("hello" + ("\n\t" + "hello")))
                        _ = v30
                    }
                case -45:
                    if len(v21) > 0 {
                        v21[0] = -300
                    }
                }
                for fi31 := 0; fi31 < 19; fi31++ {
                    var v32 int = fi27
                    if false {
                        var v33 uint16 = 632
                        v21 = append(v21, (^(fi31)))
                        _ = v33
                    }
                    _ = fi31
                    _ = v32
                }
                v21 = append(v21, fi27)
                _ = fi27
            }
        default:
            if (!(-2147483648 <= 128)) {
                v20 = v20
            } else {
                v34 := GenS1{F0: v3, F1: -251, F2: (v1 ^ v1)}
                _ = v34
            }
        }
    }
    v35 := GenS1{F0: 9223372036854775807, F1: 1, F2: (^(v1))}
    fmt.Println(v1)
    for fi36 := 0; fi36 < 3; fi36++ {
        if len(v21) >= 2 {
            v21 = v21[1:len(v21)]
        }
        _ = fi36
    }
    v2 = (v2 + v2)
    var v37 int = -304
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v35
    _ = v37
}

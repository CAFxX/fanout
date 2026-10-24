package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
func genFn0(p0 int16, p1 uint, p2 int8) int {
    var v5 uint32 = 65535
    var v6 uint = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 uint32 = 0
    var v2 int32 = 855
    v3 := []int{32767, int(v2), 127, 32767}
    var v4 string = "hello"
    v4 = v4
    for fi7 := 0; fi7 < 20; fi7++ {
        v1 = (^(v1))
        _ = fi7
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        fmt.Println(v4)
        _ = fi8
    }
    var v9 int8 = 42
    var v10 bool = ((!false) || (771 <= 0))
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := GenS1{F0: v2}
        _ = fi11
        _ = v12
    }
    var v13 uint32 = (^(v1))
    fmt.Println(v13, v2)
    switch -605 {
    case 541:
        v14 := GenS1{F0: int32(v9)}
        _ = v14
    default:
        if (0 != 0) {
            var v15 int8 = -50
            v16 := GenS1{F0: (^(v2))}
            v9 = (^(v15))
            _ = v15
            _ = v16
        } else {
            switch 479 {
            case -997:
                v13 += v13
            case 127:
                if false {
                    var v17 uint16 = 244
                    _ = v17
                }
            case -9223372036854775808:
                for fi18 := 0; fi18 < 2; fi18++ {
                    for fi19 := 0; fi19 < 20; fi19++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi19
                    }
                    _ = fi18
                }
            default:
                v20 := GenS2{F0: v4}
                _ = v20
            }
            v10 = v10
            for fi21 := 0; fi21 < 19; fi21++ {
                v2 -= (v2 & (v2 >> ((4294967295) & 63)))
                for fi22 := 0; fi22 < 11; fi22++ {
                    if (v10 || ((v2 + -1) < -2147483648)) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v4 = v4
                        fmt.Println(v13, v13)
                    }
                    if true {
                        v4 = v4
                        switch v4 {
                        case "hello":
                            v2 += v2
                        case "":
                            var v23 int32 = (v2 & -2147483648)
                            _ = v23
                        case "a":
                            if v10 {
                                if v10 {
                                    v4 = (v4 + (v4 + "hello"))
                                }
                                var v24 float32 = (0.0 + 0.1)
                                _ = v24
                            }
                        default:
                            v1 += v1
                        }
                    }
                    var v25 int16 = -192
                    _ = fi22
                    _ = v25
                }
                fmt.Println(v1, fi21, v13)
                _ = fi21
            }
        }
    }
    if v10 {
        var v26 uint8 = 128
        _ = v26
    }
    v27 := GenS1{F0: (v2 & 42)}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v9
}

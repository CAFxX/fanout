package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 int; F2 int16 }
type GenS3 struct { F0 int32 }
func genFn0() uint {
    var v4 string = "hello"
    v4 = "\n\t"
    _ = v4
    return 0
}

func main() {
    var v1 bool = (!true)
    var v2 float64 = 1e10
    var v3 uint8 = 103
    v1 = true
    for fi5 := 0; fi5 < 10; fi5++ {
        for fi6 := 0; fi6 < 2; fi6++ {
            switch 51 {
            case 9223372036854775807:
                if (-2.25 != 1.5) {
                    var v7 int = (fi5 ^ (-(fi6)))
                    v7 = (v7 * -9223372036854775808)
                    _ = v7
                } else {
                    v8 := []int{fi5}
                    v9 := GenS1{F0: -1}
                    var v10 int = fi5
                    _ = v10
                    _ = v8
                    _ = v9
                }
            default:
                v11 := []int{fi5, -9223372036854775808, fi5, fi5}
                _ = v11
            }
            fmt.Println(v2, v3, fi5)
            var v12 int = (-(fi6))
            _ = fi6
            _ = v12
        }
        if v1 {
            v13 := GenS2{F0: 9223372036854775807, F1: (fi5 / ((fi5) | 1)), F2: -1}
            v3 += (v3 + v3)
            _ = v13
        }
        _ = fi5
    }
    if (!false) {
        for fi14 := 0; fi14 < 20; fi14++ {
            v2 -= (v2 + -2.25)
            v15 := []int{fi14, (fi14 | fi14), (fi14 >> ((4294967295) & 63)), fi14}
            _ = len(v15)
            _ = fi14
            _ = v15
        }
        if v1 {
            v2 += v2
        }
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        v2 = v2
        v3 -= (v3 & (v3 >> ((428) & 63)))
        var v17 uint8 = 43
        _ = fi16
        _ = v17
    }
    v3 -= v3
    if false {
        v18 := []int{}
        v2 = v2
        if ((-9223372036854775808 > 42) && v1) {
            if v1 {
                var v19 uint = 42
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
                if false {
                    for fi20 := 0; fi20 < 18; fi20++ {
                        v1 = true
                        var v21 bool = ((!(!false)) && (true && ((!(!(1 > 136))) && (-1 > 406))))
                        _ = fi20
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 2; fi22++ {
                        if len(v18) >= 2 {
                            v18 = v18[1:len(v18)]
                        }
                        var v23 int16 = -383
                        _ = fi22
                        _ = v23
                    }
                    v19 += v19
                }
                _ = v19
            } else {
                if len(v18) > 0 {
                    v18[0] = 127
                }
                v2 -= v2
            }
            v2 += (float64(1.5) / (v2 + (((v2 - -2.25) / v2) * 0.1)))
        } else {
            v24 := GenS3{F0: 42}
            _ = v24
        }
        _ = v18
    }
    v25 := GenS2{F0: -9223372036854775808, F1: 32767, F2: -32768}
    for fi26 := 0; fi26 < 8; fi26++ {
        var v27 string = ("" + "foo bar")
        fmt.Println(v1, fi26, v1)
        _ = fi26
        _ = v27
    }
    v3 -= 128
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
}

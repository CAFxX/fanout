package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint32, p1 int, p2 int64) int32 {
    switch (p1 & 0) {
    case -418:
        var v5 int32 = 42
        _ = v5
    }
    fmt.Println(p2, p1)
    var v6 uint32 = 4294967295
    if (65535 <= 0) {
        v7 := GenS1{F0: 1}
        _ = v7
    } else {
        fmt.Println(v6, v6, p2)
        for fi8 := 0; fi8 < 18; fi8++ {
            fmt.Println(p1, p1, p0)
            _ = fi8
        }
    }
    v9 := []int{(p1 / ((p1) | 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v9
    return -578
}

func main() {
    var v1 uint64 = 974
    var v2 int32 = 486
    var v3 uint64 = v1
    var v4 int = -2147483648
    var v10 float64 = (float64(-2.25) + 3.14159)
    v10 += v10
    v3 = v1
    var v11 int = (^(v4))
    switch (v11 >> ((483) & 63)) {
    case 32768:
        v12 := GenS1{F0: 255}
        _ = v12
    }
    v10 -= v10
    switch "" {
    case "":
        v11 += (-(v11))
    case "foo bar":
        var v13 bool = (!false)
        _ = v13
    case "a":
        for fi14 := 0; fi14 < 17; fi14++ {
            v15 := []int{fi14}
            fmt.Println(v4, v10, v10)
            var v16 int = (-(fi14))
            _ = fi14
            _ = v15
            _ = v16
        }
    default:
        var v17 string = ("\n\t" + "\n\t")
        _ = v17
    }
    v11 = v11
    switch int(v3) {
    case 2147483647:
        for fi18 := 0; fi18 < 19; fi18++ {
            v19 := []int{v11, v11}
            _ = fi18
            _ = v19
        }
    case 0:
        v4 = 0
    case 127:
        if (805 <= uint16(v1)) {
            var v20 string = ("" + "")
            var v21 uint64 = 42
            var v22 bool = (v4 <= 1)
            _ = v20
            _ = v21
            _ = v22
        } else {
            v11 = 2147483647
            fmt.Println(v4, v11)
        }
    }
    if false {
        switch "hello" {
        case "hello":
            switch "" {
            case "hello":
                for fi23 := 0; fi23 < 14; fi23++ {
                    if (752 > -659) {
                        v24 := []int{-201}
                        fmt.Println(fi23)
                        _ = v24
                    }
                    if false {
                        for fi25 := 0; fi25 < 5; fi25++ {
                            v2 = (v2 & (v2 + v2))
                            v4 = (fi25 - (fi23 | v11))
                            fmt.Println(fi23)
                            _ = fi25
                        }
                        for fi26 := 0; fi26 < 13; fi26++ {
                            if (!false) {
                                v27 := GenS1{F0: 2}
                                v28 := []int{}
                                _ = v27
                                _ = v28
                            } else {
                                fmt.Println(v2)
                                fmt.Println(v1)
                            }
                            v29 := GenS1{F0: uint16(v1)}
                            _ = fi26
                            _ = v29
                        }
                    }
                    var v30 float64 = v10
                    _ = fi23
                    _ = v30
                }
            case "a":
                v31 := []int{}
                _ = v31
            }
        case "a":
            v10 = (((1e10 * v10) + ((float64((-2.25 / 1.5)) + v10) * v10)) + -2.25)
        case "\n\t":
            v3 = (^(v1))
        }
        v32 := GenS2{F0: v3}
        _ = v32
    } else {
        v33 := GenS1{F0: 44}
        _ = v33
    }
    var v34 float64 = (((v10 - 1e10) - float64(-2.25)) * 0.1)
    var v35 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}

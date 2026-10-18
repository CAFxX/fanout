package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int32, p1 int32, p2 uint8) int {
    var v4 uint32 = 0
    v5 := []int{int(p1), 32767, -9223372036854775808}
    p2 += 230
    if (v4 == v4) {
        switch -891 {
        case 289:
            if false {
                switch "hello" {
                case "foo bar":
                    for fi6 := 0; fi6 < 17; fi6++ {
                        var v7 float64 = 0.0
                        for fi8 := 0; fi8 < 1; fi8++ {
                            if true {
                                v7 -= (float64(-2.25) / -2.25)
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                var v9 uint8 = (p2 | 167)
                                _ = v9
                            }
                            _ = fi8
                        }
                        for fi10 := 0; fi10 < 20; fi10++ {
                            var v11 int32 = (p1 % (((p0 % (((p0 & (p1 | (p0 + (p0 / ((int32(fi6)) | 1)))))) | 1))) | 1))
                            v7 -= v7
                            _ = fi10
                            _ = v11
                        }
                        _ = fi6
                        _ = v7
                    }
                case "\n\t":
                    if false {
                        var v12 string = "a"
                        switch v12 {
                        case "foo bar":
                            var v13 float64 = ((-2.25 * ((-2.25 + -2.25) / float64((-2.25 * 3.14159)))) / (1e10 * (0.1 - -2.25)))
                            _ = v13
                        case "hello":
                            var v14 uint32 = 4294967295
                            _ = v14
                        case "":
                            var v15 float32 = ((1e10 + 1.5) - 0.1)
                            _ = v15
                        default:
                            v4 = (v4 & 2147483648)
                        }
                        _ = v12
                    }
                case "a":
                    if len(v5) > 0 {
                        v5[0] = 4294967295
                    }
                }
            }
        case -9223372036854775808:
            v16 := GenS1{F0: 4294967295, F1: (v4 & uint32(p2))}
            _ = v16
        case -839:
            var v17 uint32 = v4
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        for fi19 := 0; fi19 < 9; fi19++ {
            switch (-(fi19)) {
            case 127:
                p2 = p2
            case -971:
                v5 = append(v5, (fi18 / (((-(fi18))) | 1)))
            }
            v20 := GenS1{F0: 2, F1: (^(v4))}
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -2147483648
}

func main() {
    var v1 string = "hello"
    var v2 uint64 = 0
    var v3 uint64 = 373
    fmt.Println(v2)
    fmt.Println(v3)
    v21 := []int{}
    var v22 uint64 = 907
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    for fi23 := 0; fi23 < 11; fi23++ {
        fmt.Println(v1, v1)
        if len(v21) > 0 {
            v21[0] = (fi23 - fi23)
        }
        var v24 int16 = 42
        _ = fi23
        _ = v24
    }
    v25 := GenS2{F0: 0}
    v26 := GenS1{F0: 18446744073709551615, F1: 797}
    fmt.Println(v22)
    if len(v21) > 0 {
        v21[0] = -2147483648
    }
    var v27 uint16 = 65535
    var v28 float64 = (3.14159 - 0.0)
    switch "a" {
    case "foo bar":
        if false {
            v27 -= v27
        }
    case "hello":
        v28 -= float64(0.0)
    case "":
        fmt.Println(v22, v28)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}

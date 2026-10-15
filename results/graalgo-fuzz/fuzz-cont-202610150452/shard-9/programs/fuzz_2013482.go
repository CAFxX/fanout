package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }

func main() {
    var v1 uint32 = 2147483648
    var v2 int16 = int16(v1)
    var v3 uint = 368
    v4 := make(map[uint32]bool)
    var v5 string = ""
    v6 := GenS1{F0: 767, F1: 127}
    switch 32767 {
    case 240:
        v7 := []int{387, 384, 1, 139}
        _ = v7
    case 4294967295:
        if (127 > 59) {
            if (4294967295 != 42) {
                var v8 float64 = float64(3.14159)
                var v9 uint = 581
                v10 := v4[2]
                _ = v10
                _ = v8
                _ = v9
            }
        } else {
            v5 = v5
            for fi11 := 0; fi11 < 3; fi11++ {
                delete(v4, 42)
                v12 := []int{fi11}
                fmt.Println(v1, v1)
                _ = fi11
                _ = v12
            }
        }
    case 127:
        v5 = (v5 + ((v5 + "") + v5))
    default:
        fmt.Println(v3)
    }
    var v13 int8 = -128
    for fi14 := 0; fi14 < 1; fi14++ {
        _ = len(v4)
        for fi15 := 0; fi15 < 4; fi15++ {
            if (!((v2 << ((v3) & 63)) >= v2)) {
                var v16 float32 = 3.14159
                v1 = (^(v1))
                v13 += (v13 * v13)
                _ = v16
            } else {
                if (2 == 128) {
                    v2 = 42
                    v2 += v2
                    v17 := []int{(fi14 * (fi14 & 708)), 2147483647, (-(fi15)), 0}
                    _ = v17
                }
                fmt.Println(v13)
                if (false && ((fi14 / ((fi15) | 1)) > fi15)) {
                    v3 += (v3 - uint(v1))
                    fmt.Println(v5, v13, v2)
                } else {
                    v1 += v1
                    var v18 float64 = 1e10
                    _ = v18
                }
            }
            var v19 uint = v3
            var v20 int16 = -1
            _ = fi15
            _ = v19
            _ = v20
        }
        _ = fi14
    }
    v5 = (v5 + v5)
    delete(v4, v1)
    for fi21 := 0; fi21 < 9; fi21++ {
        for fi22 := 0; fi22 < 12; fi22++ {
            v23 := []int{(-(fi22)), (fi21 >> (((v3 ^ 149)) & 63))}
            _ = fi22
            _ = v23
        }
        for fi24 := 0; fi24 < 15; fi24++ {
            fmt.Println(v5, v5, v2)
            if true {
                _ = len(v4)
            }
            if (v13 != 127) {
                v1 = 2
            } else {
                fmt.Println(fi24, v5)
                if (0.1 < 0.1) {
                    switch "foo bar" {
                    case "a":
                        var v25 int32 = -389
                        _ = v25
                    default:
                        fmt.Println(v1)
                    }
                    if false {
                        v26 := GenS1{F0: v1, F1: 127}
                        if (!(!false)) {
                            if true {
                                var v27 uint = v3
                                _ = v27
                            }
                        }
                        _ = v26
                    }
                    switch (v5 + (v5 + v5)) {
                    case "a":
                        fmt.Println(v3, v13, fi24)
                    case "foo bar":
                        fmt.Println(fi21, v1)
                    case "hello":
                        v3 = (v3 * v3)
                    }
                }
            }
            _ = fi24
        }
        var v28 float32 = (((1e10 + (0.0 / ((-2.25 / 3.14159) - 1.5))) - float32(3.14159)) - 1e10)
        _ = fi21
        _ = v28
    }
    if false {
        v13 += (v13 - v13)
        v5 = (v5 + v5)
        fmt.Println(v3, v2, v1)
    }
    delete(v4, (v1 | (v1 / ((42) | 1))))
    v29 := v4[v1]
    switch 9223372036854775807 {
    case -54:
        var v30 int32 = 1
        _ = v30
    case 127:
        var v31 bool = (42 == 1)
        _ = v31
    default:
        switch (("" + v5) + v5) {
        case "hello":
            v5 = v5
        default:
            fmt.Println(v1, v13)
        }
    }
    for fi32 := 0; fi32 < 11; fi32++ {
        for fi33 := 0; fi33 < 1; fi33++ {
            var v34 int32 = 1
            v35 := []int{(fi33 % (((fi33 % ((fi33) | 1))) | 1)), (fi32 - fi33), (fi33 | fi32), (fi32 * int(v2))}
            v2 = v2
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = fi32
    }
    for fi36 := 0; fi36 < 11; fi36++ {
        var v37 int = (fi36 / (((fi36 ^ fi36)) | 1))
        if v29 {
            var v38 int16 = v2
            var v39 string = (v5 + v5)
            _ = v38
            _ = v39
        }
        _ = fi36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

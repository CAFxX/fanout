package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0() uint32 {
    var v7 bool = (((1e10 / float32(1e10)) * (1.5 / (1.5 - (((0.0 * 3.14159) / (1e10 + 3.14159)) - 1.5)))) >= ((0.1 / (1e10 + (0.0 - (0.1 + ((1e10 * 1e10) * 0.0))))) - (0.1 * -2.25)))
    for fi8 := 0; fi8 < 20; fi8++ {
        fmt.Println(fi8)
        for fi9 := 0; fi9 < 2; fi9++ {
            v7 = false
            _ = fi9
        }
        v10 := []int{fi8, fi8, 127, (-(fi8))}
        _ = fi8
        _ = v10
    }
    v11 := GenS1{F0: 205}
    var v12 int32 = 0
    _ = v11
    _ = v12
    _ = v7
    return 42
}
func genFn1(p0 int, p1 uint, p2 int) uint32 {
    fmt.Println(p0, p1, p1)
    if (0 > 0) {
        for fi13 := 0; fi13 < 11; fi13++ {
            fmt.Println(p1, p1, p2)
            p2 += (p2 ^ 531)
            fmt.Println(p2)
            _ = fi13
        }
    } else {
        fmt.Println(p0, p0, p0)
        for fi14 := 0; fi14 < 8; fi14++ {
            for fi15 := 0; fi15 < 14; fi15++ {
                fmt.Println(fi15)
                var v16 uint16 = 255
                v16 = v16
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
        v17 := []int{(p0 - p0)}
        _ = v17
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    return 65535
}

func main() {
    var v1 uint8 = 166
    var v2 float32 = 1.5
    v3 := make(map[int]int)
    var v4 uint64 = 42
    v5 := make(map[int]bool)
    var v6 uint8 = (v1 + v1)
    for fi18 := 0; fi18 < 5; fi18++ {
        for fi19 := 0; fi19 < 20; fi19++ {
            if (9223372036854775807 < -9223372036854775808) {
                if (false || ((fi19 << ((18446744073709551615) & 63)) > (fi19 / ((fi18) | 1)))) {
                    delete(v3, -695)
                    fmt.Println(v6, fi18, fi19)
                }
                var v20 uint = 975
                _ = v20
            }
            _ = fi19
        }
        _ = fi18
    }
    fmt.Println(v1, v4)
    v2 = v2
    var v21 float64 = 1.5
    if false {
        if (v21 != v21) {
            fmt.Println(v4, v21, v2)
            v4 += 1
        } else {
            var v22 int64 = -1
            for fi23 := 0; fi23 < 11; fi23++ {
                v22 += (v22 % ((v22) | 1))
                v4 += v4
                _ = fi23
            }
            var v24 int16 = 42
            _ = v22
            _ = v24
        }
        for fi25 := 0; fi25 < 10; fi25++ {
            var v26 float32 = 3.14159
            v27 := GenS1{F0: (v1 % (((v6 ^ (v6 / ((1) | 1)))) | 1))}
            var v28 uint8 = (^(v1))
            _ = fi25
            _ = v26
            _ = v27
            _ = v28
        }
    }
    v2 -= v2
    v2 += v2
    for fi29 := 0; fi29 < 15; fi29++ {
        switch (^(fi29)) {
        case 1:
            if (true || ((-2.25 - (v2 + v2)) > (v2 + 0.1))) {
                for fi30 := 0; fi30 < 15; fi30++ {
                    if (2 <= 447) {
                        v1 += v6
                    }
                    var v31 int64 = 1
                    if (true || (v21 != v21)) {
                        delete(v5, (fi29 ^ -1))
                        v32 := GenS1{F0: v6}
                        fmt.Println(fi29, fi29)
                        _ = v32
                    } else {
                        fmt.Println(v2)
                    }
                    _ = fi30
                    _ = v31
                }
            } else {
                if (!(true && (662 >= (-(fi29))))) {
                    for fi33 := 0; fi33 < 20; fi33++ {
                        var v34 uint8 = (^(v1))
                        _ = fi33
                        _ = v34
                    }
                } else {
                    v35 := GenS1{F0: (v6 << ((4294967295) & 63))}
                    _ = v35
                }
            }
        case 32768:
            var v36 float64 = ((v21 + v21) / 3.14159)
            _ = v36
        default:
            v1 -= v1
        }
        _ = fi29
    }
    switch "foo bar" {
    case "":
        var v37 int32 = 1
        _ = v37
    }
    v38 := GenS1{F0: v6}
    var v39 int64 = 997
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}

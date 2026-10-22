package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int }
func genFn0(p0 uint8, p1 uint8) int8 {
    fmt.Println(p1, p1)
    if (-9223372036854775808 < 9223372036854775807) {
        p0 = (^(p0))
        v7 := []int{408, 128, 42, 9223372036854775807}
        p0 += (p1 % (((p0 - (p1 / ((p1) | 1)))) | 1))
        _ = v7
    } else {
        var v8 uint32 = 65535
        for fi9 := 0; fi9 < 5; fi9++ {
            if true {
                if false {
                    switch "\n\t" {
                    case "a":
                        p0 = p1
                    }
                    v10 := []int{(-(fi9)), fi9}
                    _ = v10
                }
            } else {
                if false {
                    var v11 float64 = float64(((1.5 + 3.14159) * 1e10))
                    _ = v11
                }
                p1 -= (p0 * p1)
                v12 := GenS1{F0: false}
                _ = v12
            }
            _ = fi9
        }
        fmt.Println(v8, p0, p0)
        _ = v8
    }
    fmt.Println(p1, p1)
    v13 := []int{241, 2147483647, 3}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v13
    return 1
}

func main() {
    var v1 int64 = 741
    v2 := []int{4294967295}
    var v3 float32 = 1e10
    v4 := []uint32{42}
    var v5 int8 = 42
    var v6 uint32 = 297
    var v14 int16 = -32768
    var v15 uint64 = 747
    if (v1 >= (v1 ^ v1)) {
        var v16 string = "\n\t"
        _ = v16
    }
    _ = len(v2)
    var v17 uint16 = 889
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (65535 <= 18446744073709551615) {
        v14 += v14
    }
    if false {
        for fi18 := 0; fi18 < 1; fi18++ {
            if true {
                v5 = (v5 << ((408) & 63))
            }
            v19 := GenS2{F0: (fi18 & fi18)}
            if len(v2) > 0 {
                v2[0] = (-(fi18))
            }
            _ = fi18
            _ = v19
        }
        var v20 int = -494
        var v21 uint8 = 0
        _ = v20
        _ = v21
    }
    v2 = append(v2, 2147483647)
    var v22 int = -2147483648
    v5 += v5
    v15 += v15
    var v23 float32 = -2.25
    if (42 > 517) {
        var v24 float64 = ((3.14159 * 1e10) - 1.5)
        if (65535 == uint(v15)) {
            for fi25 := 0; fi25 < 8; fi25++ {
                var v26 uint = 18446744073709551615
                _ = fi25
                _ = v26
            }
            var v27 bool = ((1 <= (v6 & (^(v6)))) && (!(0.0 >= (v24 * float64((v3 - (float32((float64(3.14159) - (v24 - 0.0))) - v23)))))))
            v27 = v27
            _ = v27
        } else {
            for fi28 := 0; fi28 < 10; fi28++ {
                if (true || (false && true)) {
                    v15 = (^(v15))
                    fmt.Println(v17)
                } else {
                    for fi29 := 0; fi29 < 10; fi29++ {
                        if len(v2) > 0 {
                            v2[0] = v22
                        }
                        _ = fi29
                    }
                    fmt.Println(v15, v6, v3)
                }
                _ = fi28
            }
        }
        _ = v24
    }
    var v30 int64 = (v1 & (-(v1)))
    v3 = (v3 - (float32(1.5) / (v23 / -2.25)))
    if (false || (412 < (-(v30)))) {
        v30 = v30
    } else {
        for fi31 := 0; fi31 < 7; fi31++ {
            var v32 float64 = ((-2.25 + ((3.14159 / (0.1 * 3.14159)) - (1.5 / 1e10))) / 1.0)
            _ = fi31
            _ = v32
        }
        for fi33 := 0; fi33 < 1; fi33++ {
            for fi34 := 0; fi34 < 3; fi34++ {
                switch 32767 {
                case 32768:
                    _ = len(v4)
                }
                fmt.Println(v30, v14, fi34)
                fmt.Println(fi34, v15)
                _ = len(v2)
                v35 := GenS1{F0: ((v14 * v14) != (v14 >> ((312) & 63)))}
                _ = fi34
                _ = v35
            }
            var v36 float32 = 0.1
            _ = fi33
            _ = v36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}

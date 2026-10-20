package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int, p1 uint16, p2 uint) uint32 {
    for fi7 := 0; fi7 < 1; fi7++ {
        if (!(18446744073709551615 >= 0)) {
            p1 += (p1 & 32767)
            p0 = fi7
        }
        v8 := GenS1{F0: 208}
        if false {
            var v9 uint16 = (^(p1))
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    if (!false) {
        var v10 float32 = (((0.1 / (-2.25 / 1.0)) - 3.14159) + 0.0)
        for fi11 := 0; fi11 < 12; fi11++ {
            if true {
                var v12 int16 = 1000
                var v13 float64 = 1e10
                _ = v12
                _ = v13
            }
            fmt.Println(fi11, p2, p1)
            v14 := GenS1{F0: 126}
            _ = fi11
            _ = v14
        }
        for fi15 := 0; fi15 < 20; fi15++ {
            p2 -= 148
            _ = fi15
        }
        _ = v10
    }
    var v16 uint = p2
    p0 = (-(p0))
    v17 := []int{(p0 * -788)}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return 31
}

func main() {
    v1 := make(map[string]int)
    var v2 float32 = (-2.25 + 1e10)
    v3 := []int{}
    var v4 uint = 255
    v5 := make(map[uint32]uint32)
    var v6 bool = (193 == 2147483647)
    _ = len(v3)
    var v18 int = 127
    v19 := GenS2{F0: 2}
    if v6 {
        if v6 {
            switch 670 {
            case -764:
                switch "\n\t" {
                case "a":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                default:
                    v18 -= v18
                }
            case -186:
                var v20 uint16 = 2
                _ = v20
            default:
                v1["hello"] = v18
            }
            if (!((v6 || v6) && (!false))) {
                fmt.Println(v18)
                var v21 int16 = 1
                _ = v21
            } else {
                v1[""] = v18
                v1["\n\t"] = (^(v18))
            }
        } else {
            if true {
                for fi22 := 0; fi22 < 3; fi22++ {
                    var v23 bool = v6
                    v24 := GenS2{F0: 4294967295}
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            } else {
                var v25 int = 4294967295
                v6 = v6
                v26 := GenS2{F0: 2}
                _ = v25
                _ = v26
            }
            switch v18 {
            case 149:
                if v6 {
                    for fi27 := 0; fi27 < 2; fi27++ {
                        if (42 >= 127) {
                            v4 -= v4
                            fmt.Println(fi27)
                            v28 := GenS2{F0: 55}
                            _ = v28
                        } else {
                            v4 = (v4 << ((v4) & 63))
                        }
                        _ = fi27
                    }
                    v18 -= v18
                }
            case 42:
                var v29 int = v18
                _ = v29
            }
            if false {
                v4 = 1
            }
        }
        v30 := GenS1{F0: 1}
        _ = v30
    }
    fmt.Println(v4, v6)
    v31 := GenS3{F0: -930}
    v6 = (2147483647 <= 1)
    delete(v5, 2147483648)
    var v32 uint = 65535
    if (42 >= 65535) {
        v33 := GenS4{F0: 255, F1: 2}
        v6 = (1000 != 0)
        var v34 int16 = 799
        _ = v33
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}

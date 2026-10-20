package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
func genFn0(p0 int16, p1 int, p2 int16) int {
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := []int{p1, p1, (p1 * (fi6 ^ (-(p1))))}
        p2 = p2
        _ = fi6
        _ = v7
    }
    fmt.Println(p0, p1, p1)
    _ = p0
    _ = p1
    _ = p2
    return 127
}
func genFn1(p0 string) uint {
    p0 = "a"
    for fi8 := 0; fi8 < 9; fi8++ {
        if ((2 != 481) && ((fi8 ^ (^(fi8))) < genFn0(460, (fi8 / ((-203) | 1)), -1))) {
            if (!true) {
                var v9 uint16 = 0
                if (440 != 4294967295) {
                    switch (("a" + "") + "") {
                    case "":
                        var v10 bool = (819 <= 1)
                        _ = v10
                    case "foo bar":
                        for fi11 := 0; fi11 < 6; fi11++ {
                            p0 = p0
                            fmt.Println(fi8, fi8, fi11)
                            v12 := GenS1{F0: genFn0(403, genFn0(326, fi8, 1000), 1), F1: 42}
                            _ = fi11
                            _ = v12
                        }
                    case "hello":
                        var v13 uint = 255
                        _ = v13
                    default:
                        for fi14 := 0; fi14 < 6; fi14++ {
                            v9 = 0
                            _ = fi14
                        }
                    }
                    var v15 int64 = 0
                    for fi16 := 0; fi16 < 17; fi16++ {
                        v15 = (v15 | 0)
                        v17 := GenS1{F0: 2147483647, F1: 2}
                        _ = fi16
                        _ = v17
                    }
                    _ = v15
                } else {
                    var v18 float64 = (1e10 + 1.5)
                    for fi19 := 0; fi19 < 15; fi19++ {
                        fmt.Println(fi8)
                        switch ((p0 + (p0 + p0)) + p0) {
                        case "a":
                            if false {
                                var v20 string = p0
                                v21 := GenS1{F0: (fi8 - fi8), F1: 2147483647}
                                v22 := GenS1{F0: (-(fi8)), F1: 2147483647}
                                v23 := GenS1{F0: genFn0(-457, (fi19 >> ((820) & 63)), 42), F1: 209}
                                _ = v20
                                _ = v21
                                _ = v22
                                _ = v23
                            } else {
                                var v24 string = p0
                                var v25 uint32 = 4294967295
                                _ = v24
                                _ = v25
                            }
                        }
                        _ = fi19
                    }
                    _ = v18
                }
                v26 := []int{fi8}
                _ = v26
                _ = v9
            }
            if false {
                p0 = (p0 + "\n\t")
            }
            p0 = ("a" + "hello")
        }
        fmt.Println(p0, p0)
        v27 := []int{fi8}
        _ = fi8
        _ = v27
    }
    var v28 uint8 = 127
    v29 := []int{}
    _ = p0
    _ = v28
    _ = v29
    return 308
}

func main() {
    var v1 int8 = 0
    v2 := []uint32{1, 427, 290, 2147483647}
    var v3 uint16 = 2
    var v4 uint8 = 48
    var v5 bool = false
    switch -699 {
    case 127:
        v5 = (v3 != 42)
    case 293:
        v30 := GenS1{F0: genFn0(544, -2147483648, 1000), F1: 0}
        _ = v30
    }
    v4 = uint8(v3)
    v31 := GenS1{F0: -2147483648, F1: 2}
    var v32 int = -1
    if (!v5) {
        v4 += (v4 >> ((383) & 63))
        v4 -= (v4 % ((v4) | 1))
        v1 = v1
    } else {
        v4 += v4
        v2 = append(v2, 0)
    }
    if len(v2) > 0 {
        v2[0] = uint32(v4)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v1)
    v1 -= v1
    fmt.Println(v32, v4, v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}

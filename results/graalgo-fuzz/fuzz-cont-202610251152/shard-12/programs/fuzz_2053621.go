package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 bool, p1 string, p2 uint64) uint8 {
    v5 := []int{-9223372036854775808, 845, 128, -1}
    v6 := GenS1{F0: 137}
    v7 := GenS1{F0: 0}
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 98
}

func main() {
    var v1 int16 = -144
    var v2 uint8 = 1
    var v3 bool = false
    v4 := make(map[int]string)
    fmt.Println(v3)
    fmt.Println(v2)
    switch "\n\t" {
    case "a":
        v3 = ((!v3) || ((float64(0.1) > -2.25) && ((42 <= 42) && (true && false))))
    }
    switch -9223372036854775808 {
    case -9223372036854775808:
        for fi8 := 0; fi8 < 19; fi8++ {
            for fi9 := 0; fi9 < 16; fi9++ {
                v10 := []int{fi9, (fi8 ^ -61)}
                fmt.Println(v1, v3, v2)
                _ = fi9
                _ = v10
            }
            if v3 {
                for fi11 := 0; fi11 < 18; fi11++ {
                    v1 = v1
                    delete(v4, fi11)
                    switch -517 {
                    case -924:
                        if ((v1 & v1) > v1) {
                            v12 := GenS2{F0: (v2 + (^(v2)))}
                            v13 := []int{fi8, fi8}
                            fmt.Println(fi11, fi8)
                            _ = v12
                            _ = v13
                        } else {
                            v14 := GenS1{F0: 357}
                            if (!(128 < (v2 * (v2 >> ((4294967295) & 63))))) {
                                delete(v4, (fi8 ^ (-(fi8))))
                                var v15 int8 = -1
                                _ = v15
                            }
                            _ = v14
                        }
                        if v3 {
                            var v16 uint8 = (v2 + v2)
                            v1 -= (v1 << ((4294967295) & 63))
                            var v17 int16 = v1
                            _ = v16
                            _ = v17
                        } else {
                            v18 := []int{}
                            var v19 int8 = -1
                            if v3 {
                                v19 += 72
                                v2 += 2
                            }
                            _ = v18
                            _ = v19
                        }
                    case -523:
                        var v20 float64 = 3.14159
                        _ = v20
                    case 127:
                        v1 -= 0
                    default:
                        var v21 int16 = -32768
                        _ = v21
                    }
                    _ = fi11
                }
            } else {
                fmt.Println(fi8, fi8, fi8)
            }
            _ = fi8
        }
    case 42:
        var v22 int64 = 0
        _ = v22
    default:
        for fi23 := 0; fi23 < 11; fi23++ {
            if (!(100 > -128)) {
                fmt.Println(v2, v1)
                if (("\n\t" + "hello") == "hello") {
                    fmt.Println(v2)
                    switch "" {
                    case "hello":
                        v24 := v4[128]
                        _ = v24
                    case "\n\t":
                        v3 = (2 == 4294967295)
                    default:
                        var v25 uint8 = v2
                        _ = v25
                    }
                } else {
                    var v26 int8 = 45
                    v4[(^(fi23))] = "hello"
                    _ = v26
                }
            } else {
                var v27 int32 = 0
                _ = v27
            }
            delete(v4, 127)
            _ = fi23
        }
    }
    var v28 int16 = v1
    var v29 uint = 399
    v30 := []int{9223372036854775807, -1, 1}
    fmt.Println(v2)
    v3 = (3.14159 > 1.5)
    var v31 uint8 = 127
    _ = len(v4)
    if v3 {
        fmt.Println(v2, v29)
    } else {
        v4[4294967295] = "hello"
        switch 32768 {
        case 4294967295:
            fmt.Println(v29)
        case -363:
            v31 -= (v31 - (v31 >> ((42) & 63)))
        case 0:
            var v32 int16 = (-(v28))
            _ = v32
        }
    }
    if v3 {
        for fi33 := 0; fi33 < 12; fi33++ {
            if v3 {
                v29 += v29
            }
            _ = fi33
        }
        v34 := GenS2{F0: (v2 | (^(v2)))}
        v28 = 32767
        _ = v34
    } else {
        v1 = (v1 / (((^(v1))) | 1))
        v2 += (^(v2))
    }
    var v35 uint8 = (v2 + (v2 % ((v2) | 1)))
    v36 := GenS1{F0: 224}
    _ = len(v30)
    var v37 int16 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}

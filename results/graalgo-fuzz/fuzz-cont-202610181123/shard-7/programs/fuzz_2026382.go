package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 uint8; F2 int16 }
func genFn0(p0 uint16) int {
    var v6 int16 = -93
    v7 := GenS1{F0: 1}
    v8 := []int{4294967295, 1, 9223372036854775807, 127}
    for fi9 := 0; fi9 < 4; fi9++ {
        fmt.Println(p0, v6, p0)
        v10 := GenS1{F0: -1}
        if (false || true) {
            for fi11 := 0; fi11 < 5; fi11++ {
                var v12 int16 = 1
                switch (fi11 >> ((4294967295) & 63)) {
                case 127:
                    switch "hello" {
                    case "hello":
                        v12 = v6
                    case "a":
                        var v13 int64 = 9223372036854775807
                        _ = v13
                    case "foo bar":
                        if true {
                            p0 = 32767
                        }
                    }
                case 418:
                    var v14 uint = 65535
                    _ = v14
                }
                var v15 uint32 = 795
                var v16 uint32 = 1
                _ = len(v8)
                _ = fi11
                _ = v12
                _ = v15
                _ = v16
            }
            v8 = append(v8, fi9)
        }
        _ = fi9
        _ = v10
    }
    v6 += v6
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 553
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 100
    var v3 float64 = float64(0.0)
    var v4 uint = 4294967295
    var v5 uint16 = 599
    fmt.Println(v3)
    var v17 string = ("a" + "a")
    if (-579 > 2147483647) {
        switch "a" {
        case "\n\t":
            var v18 uint8 = 2
            _ = v18
        default:
            var v19 uint16 = (v5 ^ v5)
            _ = v19
        }
        fmt.Println(v3, v5)
        v20 := []int{-1, 1, genFn0(0)}
        _ = v20
    }
    var v21 uint8 = 61
    fmt.Println(v17, v21, v2)
    switch "hello" {
    case "\n\t":
        if (2 > 42) {
            var v22 int16 = -32768
            _ = v22
        } else {
            var v23 uint32 = 0
            _ = v23
        }
    case "a":
        for fi24 := 0; fi24 < 18; fi24++ {
            v4 -= (v4 / ((v4) | 1))
            if (v3 >= 1.5) {
                if (3.14159 > -2.25) {
                    v17 = (v17 + "")
                    v4 = v4
                } else {
                    delete(v1, 42)
                }
            }
            _ = fi24
        }
    case "hello":
        v1[65535] = true
    default:
        delete(v1, 1)
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        var v26 int64 = 1
        for fi27 := 0; fi27 < 10; fi27++ {
            v28 := GenS1{F0: (-(v26))}
            v2 += (v2 + (v2 % ((v2) | 1)))
            v29 := GenS1{F0: v26}
            fmt.Println(v3)
            _ = fi27
            _ = v28
            _ = v29
        }
        v30 := GenS2{F0: -1, F1: v21, F2: -351}
        _ = fi25
        _ = v26
        _ = v30
    }
    if (14 <= 42) {
        for fi31 := 0; fi31 < 10; fi31++ {
            var v32 uint32 = 2147483648
            _ = fi31
            _ = v32
        }
    } else {
        var v33 uint32 = 901
        v1[1] = (((-1 <= -2147483648) && false) && (false || (!("" != ((v17 + v17) + v17)))))
        _ = v33
    }
    if (255 <= 1) {
        if ((v2 % (((v2 ^ v2)) | 1)) != v2) {
            v4 = (v4 & v4)
            if false {
                var v34 int64 = -9223372036854775808
                if false {
                    switch 128 {
                    case 127:
                        switch ("foo bar" + (("hello" + v17) + v17)) {
                        case "":
                            var v35 uint8 = v21
                            _ = v35
                        default:
                            for fi36 := 0; fi36 < 5; fi36++ {
                                v34 -= 0
                                for fi37 := 0; fi37 < 4; fi37++ {
                                    v5 = (v5 >> (((^(v4))) & 63))
                                    v4 = v4
                                    fmt.Println(v34)
                                    _ = fi37
                                }
                                for fi38 := 0; fi38 < 7; fi38++ {
                                    v39 := []int{fi38, fi36}
                                    _ = fi38
                                    _ = v39
                                }
                                _ = fi36
                            }
                        }
                    case -9223372036854775808:
                        fmt.Println(v21)
                    case -369:
                        v3 = (v3 / v3)
                    }
                    if false {
                        var v40 uint8 = (v21 * v21)
                        v2 += 42
                        _ = v40
                    } else {
                        v41 := []int{}
                        var v42 uint64 = 18446744073709551615
                        for fi43 := 0; fi43 < 7; fi43++ {
                            for fi44 := 0; fi44 < 8; fi44++ {
                                var v45 bool = (!false)
                                _ = fi44
                                _ = v45
                            }
                            v21 -= v21
                            fmt.Println(v4, v3)
                            _ = fi43
                        }
                        _ = v41
                        _ = v42
                    }
                    v4 = (v4 | v4)
                }
                _ = v34
            } else {
                v2 -= v2
                var v46 uint = (^(v4))
                _ = v46
            }
        }
        var v47 int32 = -1
        for fi48 := 0; fi48 < 8; fi48++ {
            v5 = v5
            _ = fi48
        }
        _ = v47
    } else {
        v49 := GenS2{F0: 0, F1: 1, F2: -768}
        _ = v49
    }
    v50 := []int{int(v4), -453, -315}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v50:", len(v50), v50)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v50
}

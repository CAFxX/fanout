package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int32 }
func genFn0(p0 string, p1 uint8, p2 uint16) uint32 {
    p2 = 42
    for fi4 := 0; fi4 < 7; fi4++ {
        if (1e10 > 1e10) {
            fmt.Println(p2)
            p1 = (p1 & (p1 * uint8(p2)))
        } else {
            v5 := []int{32767, fi4, (^(fi4))}
            _ = v5
        }
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := []int{fi6}
            _ = fi6
            _ = v7
        }
        p1 = 2
        _ = fi4
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        fmt.Println(p1)
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1() uint8 {
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 string = "\n\t"
        v11 := []int{(fi9 % ((fi9) | 1)), (fi9 | (fi9 & fi9)), fi9}
        v11 = append(v11, fi9)
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := []int{2147483647, 330, -659, 0}
    if (-32 != 127) {
        var v13 int32 = -2147483648
        v13 = v13
        _ = v13
    }
    var v14 int64 = -1
    _ = v12
    _ = v14
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 0
    v3 := make(map[uint32]int)
    v2 = (v2 << ((18446744073709551615) & 63))
    _ = len(v1)
    switch 4294967295 {
    case 586:
        switch 2147483647 {
        case -381:
            var v15 string = "a"
            _ = v15
        case -585:
            v16 := GenS1{F0: 223}
            _ = v16
        case 9223372036854775807:
            v17 := v1[15]
            _ = v17
        }
    case 2147483647:
        if (-196 <= 2147483647) {
            var v18 int = 761
            _ = v18
        } else {
            var v19 uint32 = uint32(v2)
            fmt.Println(v19, v2)
            _ = v19
        }
    }
    var v20 float32 = 3.14159
    v21 := GenS1{F0: 42}
    if false {
        v22 := GenS1{F0: genFn1()}
        fmt.Println(v2, v20)
        _ = v22
    } else {
        v23 := []int{}
        fmt.Println(v2)
        var v24 float32 = v20
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        v26 := GenS2{F0: 4294967295, F1: -2147483648}
        v2 -= 1
        for fi27 := 0; fi27 < 14; fi27++ {
            if false {
                switch "foo bar" {
                case "\n\t":
                    if false {
                        if (65535 >= 42) {
                            var v28 uint = 1
                            var v29 uint64 = (v2 + v2)
                            _ = v28
                            _ = v29
                        } else {
                            v30 := GenS1{F0: uint8(v2)}
                            v31 := GenS1{F0: 158}
                            _ = v30
                            _ = v31
                        }
                        v20 = (v20 * v20)
                    } else {
                        var v32 float32 = v20
                        if (int64(fi27) >= 1) {
                            var v33 uint32 = 4294967295
                            _ = v33
                        } else {
                            var v34 int16 = 1
                            switch 42 {
                            case -215:
                                switch "\n\t" {
                                case "hello":
                                    var v35 uint16 = uint16(v2)
                                    _ = v35
                                case "":
                                    var v36 float32 = 0.0
                                    _ = v36
                                case "foo bar":
                                    var v37 float64 = -2.25
                                    _ = v37
                                }
                            case 0:
                                _ = len(v1)
                            }
                            _ = v34
                        }
                        var v38 int16 = -9
                        v2 = 42
                        _ = v32
                        _ = v38
                    }
                case "a":
                    v2 = (v2 % ((0) | 1))
                case "":
                    var v39 bool = (float32((1e10 / 3.14159)) < v20)
                    _ = v39
                }
            }
            var v40 uint = 211
            _ = fi27
            _ = v40
        }
        _ = fi25
        _ = v26
    }
    v41 := v1[4294967295]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v41
}

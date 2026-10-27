package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]int)
    var v3 uint8 = 127
    v4 := make(map[int]uint32)
    var v5 uint16 = 1
    var v6 int32 = 42
    var v7 int8 = -96
    v6 = (v6 >> ((1) & 63))
    for fi8 := 0; fi8 < 4; fi8++ {
        switch 170 {
        case 32767:
            var v9 float32 = 1.5
            _ = v9
        case -2147483648:
            for fi10 := 0; fi10 < 13; fi10++ {
                if true {
                    v11 := GenS1{F0: -1}
                    v12 := GenS1{F0: 9223372036854775807}
                    v7 = v7
                    _ = v11
                    _ = v12
                }
                _ = fi10
            }
        case 42:
            v13 := []int{(fi8 & (fi8 | fi8)), (fi8 << ((2) & 63)), 0}
            _ = v13
        default:
            v7 += v7
        }
        for fi14 := 0; fi14 < 4; fi14++ {
            v15 := []int{194, (fi8 ^ (fi14 + fi8))}
            v16 := GenS2{F0: v7}
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi8
    }
    v6 -= -2147483648
    v17 := []int{2147483647, 32768, -1, 5}
    var v18 int16 = 32767
    v19 := GenS2{F0: -128}
    v7 += -128
    v6 += 1
    if (!true) {
        v20 := GenS1{F0: -1}
        if true {
            for fi21 := 0; fi21 < 20; fi21++ {
                fmt.Println(v5, v5)
                switch fi21 {
                case -9223372036854775808:
                    var v22 int32 = (v6 << ((289) & 63))
                    _ = v22
                case 128:
                    if false {
                        v23 := v2[323]
                        fmt.Println(fi21, v23, v5)
                        var v24 bool = true
                        _ = v23
                        _ = v24
                    } else {
                        v25 := GenS1{F0: -230}
                        _ = v25
                    }
                case 1:
                    fmt.Println(fi21, v6, v18)
                }
                _ = fi21
            }
        }
        delete(v1, 1)
        _ = v20
    } else {
        v17 = append(v17, 727)
    }
    switch 796 {
    case 42:
        v26 := v1[2147483648]
        _ = v26
    default:
        v6 += v6
    }
    fmt.Println(v3, v3, v18)
    for fi27 := 0; fi27 < 3; fi27++ {
        fmt.Println(fi27)
        if true {
            if false {
                v28 := GenS2{F0: (v7 % ((v7) | 1))}
                var v29 string = "foo bar"
                _ = v28
                _ = v29
            } else {
                for fi30 := 0; fi30 < 19; fi30++ {
                    switch "foo bar" {
                    case "a":
                        if (!true) {
                            v31 := GenS1{F0: -923}
                            if false {
                                if true {
                                    var v32 float32 = (1.5 - -2.25)
                                    v5 = v5
                                    _ = v32
                                } else {
                                    fmt.Println(v18, v5)
                                    v18 = 585
                                    fmt.Println(v7, fi30, v7)
                                    var v33 uint = 42
                                    v34 := GenS2{F0: v7}
                                    _ = v33
                                    _ = v34
                                }
                                fmt.Println(v18, v3, v3)
                            }
                            var v35 float32 = 1.5
                            v17 = append(v17, int(v7))
                            var v36 int64 = 1
                            _ = v31
                            _ = v35
                            _ = v36
                        }
                    case "":
                        v6 = 1
                    case "foo bar":
                        v18 = v18
                    }
                    _ = fi30
                }
                var v37 uint8 = (v3 * v3)
                _ = v37
            }
        } else {
            var v38 int = (fi27 | 1)
            fmt.Println(v3, fi27, v6)
            _ = v38
        }
        _ = fi27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

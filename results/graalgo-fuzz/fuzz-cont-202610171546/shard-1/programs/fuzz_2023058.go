package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 int16; F2 int }
type GenS3 struct { F0 uint64; F1 string; F2 uint64 }
func genFn0() uint64 {
    v7 := GenS1{F0: 42}
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 uint = 616
        if false {
            var v10 int8 = -114
            var v11 int64 = 0
            var v12 uint16 = 32767
            _ = v10
            _ = v11
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    var v13 int8 = -37
    if true {
        if (42 == 240) {
            if (65535 <= 65535) {
                switch ("" + "hello") {
                case "foo bar":
                    v13 -= 90
                }
                v14 := []int{}
                _ = v14
            }
            v15 := GenS2{F0: 453, F1: 32767, F2: 978}
            _ = v15
        }
    }
    var v16 int32 = 1
    _ = v13
    _ = v16
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 uint = 4294967295
    var v2 float32 = 0.0
    v3 := []int{-620}
    v4 := []uint32{42, 4294967295, 2}
    var v5 float32 = v2
    v6 := make(map[int]int)
    v5 += float32(0.0)
    if (v5 < 1.5) {
        switch -983 {
        case -1:
            switch 849 {
            case 127:
                if false {
                    if (!true) {
                        var v17 int32 = -1
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 5; fi18++ {
                        for fi19 := 0; fi19 < 10; fi19++ {
                            var v20 int32 = -983
                            _ = fi19
                            _ = v20
                        }
                        var v21 uint = (v1 ^ 920)
                        v21 -= 255
                        _ = fi18
                        _ = v21
                    }
                    var v22 uint8 = 128
                    _ = v22
                } else {
                    v23 := GenS1{F0: (v1 * v1)}
                    _ = v23
                }
            case -2147483648:
                var v24 float32 = 1.5
                _ = v24
            case -980:
                if false {
                    if true {
                        var v25 uint64 = 2
                        _ = v25
                    }
                } else {
                    v26 := GenS1{F0: 65535}
                    var v27 uint16 = 859
                    _ = v26
                    _ = v27
                }
            default:
                _ = len(v6)
            }
        }
        v3 = append(v3, 2147483647)
        v28 := GenS1{F0: v1}
        _ = v28
    }
    v5 -= (v2 / 1e10)
    _ = len(v3)
    for fi29 := 0; fi29 < 6; fi29++ {
        var v30 int32 = 1
        for fi31 := 0; fi31 < 15; fi31++ {
            var v32 uint64 = 18446744073709551615
            _ = fi31
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    if len(v4) > 0 {
        v4[0] = 643
    }
    v33 := GenS3{F0: genFn0(), F1: "hello", F2: 4294967295}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 uint32 }
func genFn0(p0 int64, p1 int16) uint16 {
    p1 = 992
    fmt.Println(p1, p0)
    p1 += -906
    if true {
        if (65535 < 1) {
            var v6 uint = 255
            for fi7 := 0; fi7 < 3; fi7++ {
                for fi8 := 0; fi8 < 10; fi8++ {
                    v9 := GenS1{F0: 76}
                    v10 := []int{(fi8 ^ fi8), (fi7 ^ (fi7 - fi8)), (fi7 ^ (fi8 << ((0) & 63))), (fi8 >> (((v6 << ((v6) & 63))) & 63))}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                fmt.Println(v6, v6, p0)
                _ = fi7
            }
            if (1 != 127) {
                fmt.Println(v6)
                if (int64(p1) != (p0 >> (((v6 - v6)) & 63))) {
                    v11 := []int{}
                    _ = v11
                }
            }
            _ = v6
        } else {
            switch "a" {
            case "foo bar":
                v12 := GenS1{F0: 689}
                _ = v12
            }
            for fi13 := 0; fi13 < 5; fi13++ {
                fmt.Println(fi13, p0)
                p0 -= -560
                p0 += p0
                _ = fi13
            }
        }
        v14 := []int{-292, 874}
        _ = len(v14)
        _ = v14
    }
    v15 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = v15
    return 42
}
func genFn1(p0 uint, p1 uint32) uint64 {
    p1 = (^(p1))
    p1 = p1
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := []uint32{2147483647, 2}
    v2 := []int{}
    v3 := make(map[string]int)
    var v4 float32 = (0.1 + float32(1e10))
    var v5 uint32 = 16
    if true {
        var v16 float32 = v4
        _ = v16
    }
    fmt.Println(v4, v4)
    var v17 float64 = 1.5
    fmt.Println(v17, v17, v17)
    if true {
        for fi18 := 0; fi18 < 1; fi18++ {
            v19 := v3[("hello" + (("" + "a") + "foo bar"))]
            fmt.Println(v19, v19, v5)
            _ = fi18
            _ = v19
        }
        for fi20 := 0; fi20 < 11; fi20++ {
            fmt.Println(v17, v17)
            v5 -= 84
            v21 := GenS2{F0: (fi20 >> ((753) & 63)), F1: 622}
            _ = fi20
            _ = v21
        }
        for fi22 := 0; fi22 < 18; fi22++ {
            v5 = 2
            var v23 float32 = 1e10
            if (v4 < float32((0.0 * v17))) {
                for fi24 := 0; fi24 < 7; fi24++ {
                    v5 -= v5
                    v4 -= 0.1
                    if false {
                        delete(v3, ("a" + "\n\t"))
                        fmt.Println(v5, v5)
                    }
                    _ = fi24
                }
            } else {
                for fi25 := 0; fi25 < 17; fi25++ {
                    if true {
                        var v26 int32 = 211
                        _ = v26
                    }
                    v5 = (v5 + 42)
                    _ = fi25
                }
                if false {
                    v27 := GenS1{F0: 869}
                    var v28 string = "foo bar"
                    _ = v27
                    _ = v28
                }
            }
            _ = fi22
            _ = v23
        }
    } else {
        v17 = 0.1
        v17 = v17
    }
    for fi29 := 0; fi29 < 17; fi29++ {
        v5 = (v5 << ((42) & 63))
        for fi30 := 0; fi30 < 9; fi30++ {
            v4 = (v4 - (v4 + (1.5 / (v4 + v4))))
            var v31 uint32 = v5
            var v32 uint32 = (v5 * 808)
            _ = fi30
            _ = v31
            _ = v32
        }
        var v33 bool = (-502 != 986)
        _ = fi29
        _ = v33
    }
    v5 = (v5 >> ((18446744073709551615) & 63))
    v17 = v17
    v4 = v4
    for fi34 := 0; fi34 < 4; fi34++ {
        var v35 int = int(v5)
        var v36 int8 = 72
        _ = fi34
        _ = v35
        _ = v36
    }
    v37 := GenS1{F0: 0}
    v17 += (1e10 * float64((v4 * v4)))
    for fi38 := 0; fi38 < 4; fi38++ {
        v5 = v5
        _ = fi38
    }
    var v39 int32 = 2147483647
    if false {
        fmt.Println(v39, v4)
        var v40 uint = 0
        for fi41 := 0; fi41 < 5; fi41++ {
            var v42 string = (("" + ("a" + "a")) + "foo bar")
            _ = fi41
            _ = v42
        }
        _ = v40
    } else {
        var v43 uint16 = 427
        for fi44 := 0; fi44 < 20; fi44++ {
            v43 = uint16(v5)
            v17 = 0.0
            _ = fi44
        }
        _ = v43
    }
    var v45 int8 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v37
    _ = v39
    _ = v4
    _ = v45
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint16; F2 int32 }
func genFn0(p0 uint32, p1 uint64, p2 int8) bool {
    p2 -= (p2 / ((42) | 1))
    fmt.Println(p0)
    v6 := GenS1{F0: p0, F1: 128}
    var v7 uint16 = 0
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 uint = 760
        var v10 uint64 = p1
        if ((v9 / ((4294967295) | 1)) > v9) {
            v11 := GenS2{F0: 0, F1: v7, F2: -42}
            _ = v11
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return true
}
func genFn1(p0 string, p1 int8) int64 {
    var v12 uint32 = 730
    if true {
        switch 100 {
        case 0:
            var v13 uint64 = 2
            _ = v13
        }
        v14 := []int{}
        var v15 int = 9223372036854775807
        _ = v14
        _ = v15
    } else {
        var v16 float32 = 1e10
        v17 := GenS2{F0: 909, F1: 2, F2: -2147483648}
        _ = v16
        _ = v17
    }
    fmt.Println(p1, v12)
    _ = p0
    _ = p1
    _ = v12
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    var v3 int8 = -1
    var v4 int64 = -459
    var v5 uint16 = 65535
    switch 42 {
    case 4294967295:
        v4 = -9223372036854775808
    case 1:
        for fi18 := 0; fi18 < 8; fi18++ {
            v19 := GenS2{F0: 0, F1: 1, F2: 42}
            var v20 int64 = v4
            _ = fi18
            _ = v19
            _ = v20
        }
    default:
        fmt.Println(v4)
    }
    v21 := v1[699]
    v3 -= v3
    fmt.Println(v5, v21)
    v22 := v2[""]
    v23 := GenS2{F0: -1, F1: 42, F2: 1}
    if true {
        if (0 <= -2147483648) {
            if false {
                fmt.Println(v4, v3, v4)
                var v24 int64 = v4
                _ = v24
            }
        }
        switch (^(v22)) {
        case -1:
            v3 += (v3 & 127)
        case 1:
            var v25 int8 = (v3 << ((837) & 63))
            _ = v25
        case 4294967295:
            var v26 int = (v22 | (v22 >> ((4294967295) & 63)))
            _ = v26
        }
        v4 = (v4 ^ (^(v4)))
    }
    if (false && false) {
        if (!(!false)) {
            v4 += -270
            fmt.Println(v21, v22, v3)
            if (v5 >= v5) {
                for fi27 := 0; fi27 < 3; fi27++ {
                    switch fi27 {
                    case 214:
                        if (!true) {
                            v28 := GenS2{F0: 32767, F1: v5, F2: 575}
                            _ = v28
                        } else {
                            fmt.Println(v22, v3, fi27)
                            for fi29 := 0; fi29 < 17; fi29++ {
                                v30 := GenS2{F0: -177, F1: (v5 % ((651) | 1)), F2: int32(v5)}
                                var v31 int32 = 967
                                v5 = 987
                                _ = fi29
                                _ = v30
                                _ = v31
                            }
                        }
                    case 599:
                        if false {
                            v32 := GenS1{F0: (v21 >> ((0) & 63)), F1: 15}
                            _ = v32
                        }
                    }
                    v5 += v5
                    _ = fi27
                }
                var v33 string = "a"
                _ = v33
            }
        }
        for fi34 := 0; fi34 < 10; fi34++ {
            if true {
                var v35 uint = 849
                _ = v35
            } else {
                v21 = v21
                for fi36 := 0; fi36 < 15; fi36++ {
                    v5 += (v5 << ((2) & 63))
                    var v37 string = "foo bar"
                    v38 := GenS2{F0: -250, F1: (v5 * 32767), F2: 42}
                    _ = fi36
                    _ = v37
                    _ = v38
                }
                v21 -= (^(v21))
            }
            var v39 string = "hello"
            for fi40 := 0; fi40 < 14; fi40++ {
                _ = len(v1)
                var v41 int64 = v4
                var v42 int8 = v3
                _ = fi40
                _ = v41
                _ = v42
            }
            _ = fi34
            _ = v39
        }
        for fi43 := 0; fi43 < 12; fi43++ {
            var v44 int64 = (v4 << ((723) & 63))
            var v45 int32 = -2147483648
            _ = fi43
            _ = v44
            _ = v45
        }
    } else {
        var v46 int64 = (-(v4))
        v47 := GenS2{F0: 1000, F1: (v5 << ((338) & 63)), F2: -554}
        var v48 int32 = -1
        _ = v46
        _ = v47
        _ = v48
    }
    v3 -= v3
    var v49 uint8 = 10
    var v50 uint8 = 2
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() uint16 {
    var v7 int8 = 117
    if true {
        v8 := []int{}
        v7 -= (v7 * v7)
        v7 += v7
        _ = v8
    } else {
        v9 := []int{1, 2147483647}
        _ = v9
    }
    _ = v7
    return 0
}
func genFn1() bool {
    var v10 float32 = 3.14159
    for fi11 := 0; fi11 < 11; fi11++ {
        v12 := []int{(fi11 & 1), fi11}
        var v13 uint64 = 437
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = v10
    return true
}

func main() {
    var v1 int = 2147483647
    var v2 int16 = -1
    var v3 uint8 = 1
    v4 := []int{v1}
    v5 := []string{}
    v6 := make(map[int]uint32)
    for fi14 := 0; fi14 < 14; fi14++ {
        if true {
            for fi15 := 0; fi15 < 8; fi15++ {
                v16 := GenS1{F0: v2}
                _ = fi15
                _ = v16
            }
            v17 := GenS1{F0: (v2 & 371)}
            if (1 > 65535) {
                v1 -= fi14
                for fi18 := 0; fi18 < 6; fi18++ {
                    var v19 uint64 = 2
                    var v20 float64 = float64(0.1)
                    var v21 float32 = 1.5
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v21
                }
            }
            _ = v17
        } else {
            v3 = v3
        }
        for fi22 := 0; fi22 < 4; fi22++ {
            switch 1 {
            case -882:
                fmt.Println(v1)
            }
            v1 = fi22
            v3 += 224
            _ = fi22
        }
        var v23 int64 = 878
        _ = fi14
        _ = v23
    }
    v1 = -559
    var v24 uint32 = 2147483647
    v1 = v1
    for fi25 := 0; fi25 < 15; fi25++ {
        v3 = (^(v3))
        v24 += v24
        _ = fi25
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v24, v1, v1)
    if (!false) {
        if ((3.14159 < 1.5) || true) {
            var v26 float64 = 3.14159
            _ = len(v6)
            v1 = v1
            _ = v26
        }
        var v27 int8 = 100
        v27 += (-(v27))
        _ = v27
    } else {
        v24 = v24
        var v28 uint = 0
        _ = v28
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        v6[(v1 + 558)] = v24
        v30 := GenS1{F0: v2}
        _ = fi29
        _ = v30
    }
    v31 := GenS1{F0: (-(v2))}
    var v32 string = "foo bar"
    v33 := GenS1{F0: v2}
    fmt.Println(v1)
    if false {
        if false {
            v1 += v1
            var v34 int64 = 734
            _ = v34
        } else {
            v35 := v6[(v1 ^ int(v3))]
            for fi36 := 0; fi36 < 13; fi36++ {
                if true {
                    fmt.Println(v2, v2, fi36)
                }
                _ = fi36
            }
            _ = v35
        }
        for fi37 := 0; fi37 < 3; fi37++ {
            if true {
                v1 -= fi37
                v1 -= fi37
            }
            _ = fi37
        }
        var v38 bool = true
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}

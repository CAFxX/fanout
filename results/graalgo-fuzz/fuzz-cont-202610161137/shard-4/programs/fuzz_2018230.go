package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int32; F1 bool; F2 uint64 }
type GenS3 struct { F0 uint16; F1 int8; F2 bool }
type GenS4 struct { F0 bool; F1 uint }
func genFn0() int32 {
    var v6 int16 = 1000
    v6 = (-(v6))
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            v6 -= (v6 ^ v6)
            _ = fi8
        }
        if (255 <= 2) {
            v9 := GenS1{F0: 4294967295, F1: 0}
            fmt.Println(v6, fi7)
            _ = v9
        } else {
            switch (fi7 & -701) {
            case -761:
                v10 := GenS2{F0: -11, F1: (!(0 <= 305)), F2: 18446744073709551615}
                _ = v10
            case 42:
                var v11 uint64 = 0
                _ = v11
            }
        }
        _ = fi7
    }
    v12 := GenS1{F0: 724, F1: 42}
    var v13 int16 = (v6 % ((-1) | 1))
    _ = v12
    _ = v13
    _ = v6
    return -1
}

func main() {
    v1 := []int{1, -463, 0}
    v2 := make(map[string]bool)
    v3 := []uint32{}
    var v4 uint64 = 1
    var v5 uint64 = (v4 % ((42) | 1))
    for fi14 := 0; fi14 < 7; fi14++ {
        if false {
            var v15 uint32 = 1
            var v16 uint = 689
            v17 := GenS2{F0: -2147483648, F1: true, F2: 0}
            _ = v15
            _ = v16
            _ = v17
        } else {
            for fi18 := 0; fi18 < 17; fi18++ {
                v5 += v4
                if (1 >= 1) {
                    var v19 bool = false
                    var v20 uint8 = 128
                    _ = v19
                    _ = v20
                } else {
                    if false {
                        var v21 int32 = 2147483647
                        var v22 bool = true
                        v4 = 18446744073709551615
                        v2["\n\t"] = (((1.5 - 3.14159) >= (1.5 + (((0.1 + -2.25) + 1e10) + -2.25))) && true)
                        var v23 int8 = 100
                        _ = v21
                        _ = v22
                        _ = v23
                    } else {
                        v4 += v4
                    }
                }
                var v24 int32 = 0
                _ = fi18
                _ = v24
            }
            for fi25 := 0; fi25 < 16; fi25++ {
                var v26 uint16 = uint16(v5)
                v2[""] = false
                if true {
                    v2["a"] = (!(!(0 <= 42)))
                } else {
                    delete(v2, "\n\t")
                    v4 = v4
                }
                _ = fi25
                _ = v26
            }
        }
        _ = fi14
    }
    v27 := GenS1{F0: v5, F1: -262}
    if false {
        var v28 bool = (-48 <= 1)
        _ = v28
    }
    _ = len(v2)
    var v29 uint64 = v5
    if false {
        var v30 int16 = 42
        fmt.Println(v30)
        if (v30 <= (v30 ^ v30)) {
            v5 -= (v5 ^ 42)
            var v31 uint64 = (v29 >> ((642) & 63))
            fmt.Println(v4, v4, v5)
            _ = v31
        }
        _ = v30
    } else {
        if (1 > 1) {
            var v32 bool = (0 <= 572)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v33 := GenS2{F0: genFn0(), F1: true, F2: 97}
            _ = v32
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 7; fi34++ {
        var v35 string = "foo bar"
        fmt.Println(v4)
        v36 := GenS2{F0: 2147483647, F1: true, F2: v4}
        _ = fi34
        _ = v35
        _ = v36
    }
    v29 = v5
    v5 = v5
    v5 -= 0
    for fi37 := 0; fi37 < 3; fi37++ {
        if (-379 >= 1) {
            var v38 uint8 = 0
            _ = len(v2)
            _ = v38
        } else {
            v39 := GenS3{F0: 2, F1: 63, F2: false}
            v29 -= (v4 / ((v5) | 1))
            _ = v39
        }
        var v40 string = "a"
        _ = fi37
        _ = v40
    }
    v5 = v4
    v4 += v5
    v29 += v4
    for fi41 := 0; fi41 < 6; fi41++ {
        v4 = 2
        v29 = (v5 ^ v29)
        fmt.Println(v4, v5, v4)
        _ = fi41
    }
    for fi42 := 0; fi42 < 7; fi42++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi43 := 0; fi43 < 3; fi43++ {
            var v44 uint32 = 2147483648
            fmt.Println(v4)
            v29 += v5
            v45 := GenS4{F0: true, F1: 981}
            v29 = v4
            v29 += 658
            _ = fi43
            _ = v44
            _ = v45
        }
        _ = fi42
    }
    for fi46 := 0; fi46 < 6; fi46++ {
        v47 := v2["foo bar"]
        switch ("" + "hello") {
        case "\n\t":
            v48 := GenS2{F0: -1, F1: (!false), F2: (v4 << ((2) & 63))}
            _ = v48
        }
        _ = fi46
        _ = v47
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}

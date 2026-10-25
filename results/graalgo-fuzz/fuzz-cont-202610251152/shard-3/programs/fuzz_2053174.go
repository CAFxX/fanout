package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint16 }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 int32 }
func genFn0(p0 uint32, p1 int32, p2 uint16) int {
    if (!true) {
        p0 -= 2147483648
    }
    for fi5 := 0; fi5 < 7; fi5++ {
        switch 42 {
        case 520:
            if (!(1 < 18446744073709551615)) {
                switch "\n\t" {
                case "a":
                    var v6 uint64 = 0
                    _ = v6
                default:
                    if false {
                        p0 = (p0 & p0)
                        p1 -= (p1 * p1)
                    } else {
                        p1 -= -2147483648
                    }
                }
            }
        case 530:
            if ((0.1 * (0.0 - (((0.1 + 0.1) - 1e10) + 1.5))) != 3.14159) {
                for fi7 := 0; fi7 < 12; fi7++ {
                    p0 -= (p0 << ((248) & 63))
                    _ = fi7
                }
                var v8 int16 = 1000
                if true {
                    fmt.Println(p2, p2)
                    var v9 int64 = 1
                    var v10 int64 = 46
                    _ = v10
                    _ = v9
                }
                _ = v8
            }
        case 42:
            for fi11 := 0; fi11 < 13; fi11++ {
                var v12 string = "a"
                var v13 int16 = -162
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        v14 := GenS1{F0: p2}
        _ = fi5
        _ = v14
    }
    var v15 bool = (0.1 > -2.25)
    var v16 int64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return 32768
}
func genFn1() uint {
    if (-1 == 1) {
        if (318 >= 2) {
            v17 := GenS2{F0: 2147483647, F1: 1, F2: 127}
            _ = v17
        } else {
            if (!false) {
                switch 379 {
                case 137:
                    var v18 uint8 = 91
                    _ = v18
                }
                var v19 bool = false
                _ = v19
            } else {
                v20 := GenS1{F0: 42}
                for fi21 := 0; fi21 < 20; fi21++ {
                    if true {
                        var v22 int64 = -261
                        v23 := GenS3{F0: 0, F1: 641, F2: 280}
                        var v24 int64 = (v22 | 0)
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    var v25 uint32 = 53
                    _ = fi21
                    _ = v25
                }
                var v26 int8 = -1
                _ = v20
                _ = v26
            }
            v27 := GenS4{F0: -1, F1: ((!true) && false)}
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 6; fi28++ {
        if true {
            var v29 int64 = 1
            var v30 bool = (0.1 < (((0.1 - 0.1) + (0.1 * 3.14159)) / -2.25))
            _ = v29
            _ = v30
        }
        var v31 uint = 1
        _ = fi28
        _ = v31
    }
    v32 := []int{32768}
    _ = v32
    return 2
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[uint32]string)
    var v3 uint64 = 4294967295
    v4 := make(map[string]bool)
    v3 -= (v3 >> ((2) & 63))
    var v33 string = ("hello" + ("a" + "\n\t"))
    v34 := GenS5{F0: -2147483648}
    var v35 bool = false
    v36 := GenS1{F0: (v1 % ((v1) | 1))}
    v33 = ("a" + (v33 + v33))
    v33 = v33
    var v37 float32 = float32(3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}

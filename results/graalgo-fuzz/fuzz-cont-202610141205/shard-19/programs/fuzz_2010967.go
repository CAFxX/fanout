package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint64, p1 uint64, p2 uint32) int {
    var v4 uint64 = p0
    switch 127 {
    case -9223372036854775808:
        v5 := GenS1{F0: 128, F1: -397, F2: 679}
        _ = v5
    case 4294967295:
        var v6 int = 2147483647
        _ = v6
    }
    switch 42 {
    case -801:
        var v7 uint16 = 0
        _ = v7
    case -2147483648:
        var v8 float64 = (3.14159 * (((((-2.25 - (1.5 / float64((3.14159 / -2.25)))) / 1.0) / float64(1e10)) / 1.0) / (-2.25 * float64((0.1 - 1.5)))))
        _ = v8
    default:
        if (0.0 <= (3.14159 * (0.1 / 1e10))) {
            v9 := []int{-1}
            p2 += p2
            for fi10 := 0; fi10 < 15; fi10++ {
                var v11 uint8 = 255
                v4 -= 42
                _ = fi10
                _ = v11
            }
            _ = v9
        } else {
            var v12 uint16 = uint16(v4)
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -2147483648
}

func main() {
    var v1 uint64 = 1
    var v2 int = 2147483647
    var v3 int8 = 42
    for fi13 := 0; fi13 < 6; fi13++ {
        v1 = (v1 ^ v1)
        var v14 string = ""
        if false {
            v1 -= 2
            v1 = (^(v1))
        }
        _ = fi13
        _ = v14
    }
    if false {
        v1 -= v1
        v15 := GenS1{F0: (^(v2)), F1: (-(v2)), F2: v2}
        _ = v15
    }
    v16 := GenS1{F0: v2, F1: (v2 & v2), F2: (v2 * v2)}
    for fi17 := 0; fi17 < 11; fi17++ {
        if (!(!(!(false && false)))) {
            if (v3 == v3) {
                for fi18 := 0; fi18 < 4; fi18++ {
                    if true {
                        v3 = v3
                    }
                    fmt.Println(v1, v2, fi18)
                    if (716 > 42) {
                        fmt.Println(fi18, v1, fi17)
                        var v19 int32 = 4
                        fmt.Println(v3, fi17, fi18)
                        _ = v19
                    }
                    _ = fi18
                }
                for fi20 := 0; fi20 < 1; fi20++ {
                    v2 = 4294967295
                    v21 := GenS1{F0: -910, F1: (fi20 % ((fi17) | 1)), F2: (fi20 << ((510) & 63))}
                    var v22 uint32 = 0
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            }
        } else {
            if (4294967295 > (fi17 & (fi17 >> ((2) & 63)))) {
                v3 -= v3
                var v23 uint32 = 2
                _ = v23
            }
        }
        if (!false) {
            v3 = 127
        } else {
            v2 = genFn0(v1, v1, 1)
            var v24 string = "hello"
            v25 := GenS1{F0: -238, F1: (^(fi17)), F2: (fi17 + (v2 * v2))}
            _ = v24
            _ = v25
        }
        if false {
            if false {
                for fi26 := 0; fi26 < 20; fi26++ {
                    fmt.Println(fi26, fi26)
                    v3 -= v3
                    var v27 float64 = 0.0
                    _ = fi26
                    _ = v27
                }
                v28 := GenS1{F0: v2, F1: (fi17 << ((65535) & 63)), F2: 42}
                fmt.Println(fi17, fi17)
                _ = v28
            }
            fmt.Println(v2, fi17)
            v29 := GenS1{F0: v2, F1: -441, F2: genFn0(v1, v1, 65535)}
            _ = v29
        }
        _ = fi17
    }
    v3 += v3
    var v30 int = 128
    var v31 int16 = -927
    var v32 float64 = 1.5
    fmt.Println(v31, v2)
    var v33 float64 = v32
    v3 = v3
    v34 := GenS1{F0: (^(v2)), F1: v30, F2: (v2 - int(v3))}
    if (709 <= 685) {
        for fi35 := 0; fi35 < 18; fi35++ {
            v36 := GenS1{F0: genFn0((v1 | v1), (v1 + 2), 42), F1: (^(v2)), F2: genFn0(v1, (v1 ^ 18446744073709551615), 0)}
            v3 = (^(v3))
            v37 := GenS1{F0: int(v31), F1: (fi35 << ((702) & 63)), F2: v30}
            _ = fi35
            _ = v36
            _ = v37
        }
    }
    for fi38 := 0; fi38 < 7; fi38++ {
        fmt.Println(v3)
        if false {
            if ((float64((3.14159 - 0.0)) < 3.14159) || false) {
                var v39 int = (fi38 / ((fi38) | 1))
                v33 = v33
                v32 = 1.5
                _ = v39
            } else {
                v40 := GenS2{F0: v1}
                v31 = (v31 | v31)
                v41 := []int{v30}
                _ = v40
                _ = v41
            }
            v32 = v32
            fmt.Println(v1, fi38)
        } else {
            v30 -= (-(v30))
            var v42 int32 = 245
            _ = v42
        }
        _ = fi38
    }
    v33 = v32
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
}

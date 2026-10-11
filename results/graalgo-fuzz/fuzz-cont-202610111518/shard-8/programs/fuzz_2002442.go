package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 uint16; F2 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 string; F2 bool }
type GenS5 struct { F0 string; F1 uint }
func genFn0(p0 uint16) uint8 {
    v4 := []int{127}
    if false {
        v5 := GenS1{F0: true}
        for fi6 := 0; fi6 < 8; fi6++ {
            v7 := GenS2{F0: -48, F1: 1, F2: "hello"}
            var v8 bool = true
            v4 = append(v4, (^(fi6)))
            _ = fi6
            _ = v7
            _ = v8
        }
        p0 = (^(p0))
        _ = v5
    } else {
        p0 = p0
        var v9 float32 = (1e10 * 0.1)
        _ = v9
    }
    fmt.Println(p0)
    p0 = p0
    _ = p0
    _ = v4
    return 175
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[uint32]uint32)
    var v3 int8 = 1
    var v10 int16 = -936
    var v11 int32 = 593
    var v12 int8 = (^(v3))
    for fi13 := 0; fi13 < 9; fi13++ {
        v11 -= (v11 - 1)
        for fi14 := 0; fi14 < 1; fi14++ {
            if ((((((1e10 * 3.14159) - float64(float32(((1.5 / 1.0) / 1.0)))) + float64(0.0)) <= -2.25) || (0 != 0)) && true) {
                switch "a" {
                case "hello":
                    v10 = 1000
                }
            } else {
                v15 := []int{(fi14 | (fi14 - 32767)), (fi13 << ((626) & 63)), (fi14 - (fi13 % ((32768) | 1)))}
                var v16 uint8 = 198
                _ = v15
                _ = v16
            }
            var v17 bool = false
            _ = fi14
            _ = v17
        }
        _ = fi13
    }
    fmt.Println(v11, v10, v1)
    if true {
        v12 -= 100
    } else {
        switch 9223372036854775807 {
        case 1:
            for fi18 := 0; fi18 < 16; fi18++ {
                if true {
                    fmt.Println(fi18, fi18)
                } else {
                    var v19 int16 = 32767
                    for fi20 := 0; fi20 < 5; fi20++ {
                        if ((v19 & int16(fi18)) < (-(v19))) {
                            if true {
                                var v21 uint64 = 2
                                var v22 uint8 = 128
                                fmt.Println(v12, v19)
                                if true {
                                    fmt.Println(v3, fi18)
                                    if (0 <= 9223372036854775807) {
                                        v22 += genFn0(255)
                                        delete(v2, uint32(v22))
                                    } else {
                                        v23 := GenS1{F0: (!false)}
                                        _ = v23
                                    }
                                    v24 := GenS1{F0: true}
                                    _ = v24
                                }
                                _ = v21
                                _ = v22
                            }
                            fmt.Println(fi20, fi20)
                            var v25 uint32 = 652
                            _ = v25
                        } else {
                            fmt.Println(v19, v19)
                        }
                        _ = fi20
                    }
                    _ = v19
                }
                _ = fi18
            }
        }
        v26 := GenS1{F0: false}
        v10 = 32767
        _ = v26
    }
    var v27 int = 2147483647
    for fi28 := 0; fi28 < 1; fi28++ {
        fmt.Println(v12, v1, v3)
        _ = fi28
    }
    if false {
        v29 := GenS2{F0: v12, F1: 32767, F2: "\n\t"}
        v12 = 1
        _ = v29
    } else {
        for fi30 := 0; fi30 < 16; fi30++ {
            for fi31 := 0; fi31 < 13; fi31++ {
                var v32 bool = (42 < 2)
                v33 := v2[447]
                v34 := GenS3{F0: 2}
                _ = fi31
                _ = v32
                _ = v33
                _ = v34
            }
            _ = fi30
        }
    }
    if (true || true) {
        fmt.Println(v11, v11)
    }
    switch "hello" {
    case "\n\t":
        var v35 int16 = (v10 / (((v10 - (v10 ^ v10))) | 1))
        _ = v35
    default:
        if true {
            v36 := GenS4{F0: 2147483647, F1: "foo bar", F2: true}
            if false {
                v37 := GenS5{F0: "\n\t", F1: 65535}
                v38 := GenS5{F0: ("" + (("hello" + "a") + "a")), F1: 2}
                fmt.Println(v27, v12, v1)
                _ = v37
                _ = v38
            }
            var v39 int16 = v10
            _ = v36
            _ = v39
        } else {
            for fi40 := 0; fi40 < 9; fi40++ {
                v41 := []int{}
                fmt.Println(v1)
                var v42 uint = 973
                _ = fi40
                _ = v41
                _ = v42
            }
        }
    }
    if (!true) {
        if true {
            for fi43 := 0; fi43 < 5; fi43++ {
                v10 += 1
                _ = fi43
            }
            v10 = -1
            if (true || ((false || (!true)) && true)) {
                v11 -= (v11 * (v11 | int32(v1)))
            }
        } else {
            for fi44 := 0; fi44 < 20; fi44++ {
                var v45 uint8 = 127
                var v46 string = (("a" + ("hello" + "\n\t")) + "a")
                _ = fi44
                _ = v45
                _ = v46
            }
            var v47 uint32 = 65535
            _ = v47
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v27
    _ = v3
}

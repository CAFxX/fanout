package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 bool }
type GenS2 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int16; F2 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int }
func genFn0(p0 uint32, p1 int16, p2 int64) uint8 {
    var v4 float32 = -2.25
    fmt.Println(p0, p2)
    var v5 uint8 = 0
    var v6 bool = (42 != -2147483648)
    for fi7 := 0; fi7 < 4; fi7++ {
        v4 = 1.5
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return v5
}
func genFn1(p0 int8, p1 int32, p2 uint) uint {
    v8 := []int{-1, 2147483647}
    var v9 int16 = 1000
    v8 = append(v8, 127)
    v10 := GenS1{F0: 2147483647, F1: 128, F2: false}
    for fi11 := 0; fi11 < 20; fi11++ {
        var v12 float32 = (0.1 - -2.25)
        if true {
            switch (fi11 % ((fi11) | 1)) {
            case 975:
                switch 128 {
                case 42:
                    _ = len(v8)
                default:
                    if len(v8) > 0 {
                        v8[0] = fi11
                    }
                }
            case -1:
                for fi13 := 0; fi13 < 11; fi13++ {
                    if false {
                        if len(v8) > 0 {
                            v8[0] = (fi11 * int(p2))
                        }
                    } else {
                        if (genFn0(0, 32767, 42) == 245) {
                            v9 += (v9 ^ (v9 % ((-1) | 1)))
                            v14 := GenS2{F0: 0, F1: (^(p2)), F2: 732}
                            if ((p2 / ((810) | 1)) < p2) {
                                if false {
                                    if ((1.5 > (-2.25 + float64(v12))) && true) {
                                        if true {
                                            v8 = append(v8, 9223372036854775807)
                                        }
                                        switch "a" {
                                        case "a":
                                            p2 = p2
                                        case "foo bar":
                                            fmt.Println(fi11)
                                            v8 = append(v8, (-(fi11)))
                                        case "hello":
                                            _ = len(v8)
                                        }
                                        if (321 > 426) {
                                            if (!(51 >= uint8(p2))) {
                                                fmt.Println(fi11, v12)
                                                var v15 float64 = 1.5
                                                var v16 bool = false
                                                var v17 string = "a"
                                                _ = v15
                                                _ = v16
                                                _ = v17
                                            }
                                            v18 := GenS3{F0: 42, F1: (-(v9)), F2: 2147483647}
                                            _ = v18
                                        }
                                    } else {
                                        p2 -= (^(p2))
                                    }
                                    var v19 uint8 = 0
                                    _ = v19
                                }
                                p1 = (p1 % ((p1) | 1))
                                var v20 bool = (-2.25 != 0.0)
                                v21 := GenS1{F0: p1, F1: 0, F2: (!("\n\t" == ("foo bar" + "\n\t")))}
                                v22 := GenS2{F0: 2, F1: p2, F2: 65535}
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                            _ = v14
                        }
                    }
                    p1 = (p1 * p1)
                    var v23 uint8 = 0
                    _ = fi13
                    _ = v23
                }
            default:
                switch "foo bar" {
                case "":
                    var v24 uint32 = 183
                    _ = v24
                default:
                    v25 := GenS4{F0: 2147483647}
                    _ = v25
                }
            }
            fmt.Println(v12)
            var v26 int16 = v9
            _ = v26
        }
        var v27 uint = p2
        _ = fi11
        _ = v12
        _ = v27
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return p2
}

func main() {
    var v1 string = (("" + "foo bar") + ("hello" + "hello"))
    var v2 uint16 = 1
    v3 := make(map[string]string)
    for fi28 := 0; fi28 < 12; fi28++ {
        v29 := GenS5{F0: (fi28 ^ (fi28 | (fi28 & (fi28 >> ((genFn1(63, -1, 457)) & 63)))))}
        var v30 uint = 2
        v30 = (v30 | v30)
        _ = fi28
        _ = v29
        _ = v30
    }
    switch 9223372036854775807 {
    case -650:
        var v31 uint8 = genFn0(42, 514, -697)
        _ = v31
    case 42:
        v2 = v2
    default:
        v32 := GenS5{F0: 2147483647}
        _ = v32
    }
    var v33 uint32 = 2147483648
    fmt.Println(v2)
    v34 := GenS5{F0: -1}
    fmt.Println(v2, v1)
    v2 -= (v2 | (v2 << ((4294967295) & 63)))
    fmt.Println(v1)
    if true {
        var v35 int32 = 149
        _ = v35
    } else {
        v36 := []int{42, 0, 32767, -1}
        var v37 uint64 = 231
        if true {
            var v38 int8 = -109
            _ = v38
        } else {
            var v39 float64 = 0.1
            _ = v39
        }
        _ = v36
        _ = v37
    }
    var v40 float32 = (0.0 - 3.14159)
    v41 := GenS2{F0: 178, F1: 2, F2: 42}
    v42 := GenS4{F0: v33}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v40
    _ = v41
    _ = v42
}

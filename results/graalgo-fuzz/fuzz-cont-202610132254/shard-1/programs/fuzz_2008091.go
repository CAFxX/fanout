package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint; F1 uint; F2 uint32 }
func genFn0(p0 uint) uint64 {
    p0 = 18446744073709551615
    if false {
        v4 := GenS1{F0: 18446744073709551615, F1: 97, F2: ("foo bar" + "hello")}
        v5 := GenS2{F0: 100}
        _ = v4
        _ = v5
    }
    _ = p0
    return 42
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 0
    var v3 uint = 4294967295
    fmt.Println(v3, v3)
    if (false && true) {
        if true {
            for fi6 := 0; fi6 < 2; fi6++ {
                v7 := v1[fi6]
                v8 := GenS1{F0: genFn0(v3), F1: (-(v2)), F2: (v7 + ("foo bar" + (v7 + (v7 + v7))))}
                switch ("a" + "foo bar") {
                case "hello":
                    for fi9 := 0; fi9 < 16; fi9++ {
                        var v10 float32 = (-2.25 - (3.14159 / 1.0))
                        _ = fi9
                        _ = v10
                    }
                }
                _ = fi6
                _ = v7
                _ = v8
            }
            v2 -= v2
            v11 := GenS3{F0: 4294967295, F1: 409}
            _ = v11
        } else {
            var v12 uint = (v3 % (((v3 ^ v3)) | 1))
            var v13 int64 = -1
            delete(v1, -1)
            _ = v12
            _ = v13
        }
    } else {
        v2 += (v2 >> ((v3) & 63))
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        switch "" {
        case "foo bar":
            for fi15 := 0; fi15 < 17; fi15++ {
                var v16 uint8 = 128
                v3 -= (v3 % (((v3 - (v3 << ((v3) & 63)))) | 1))
                _ = fi15
                _ = v16
            }
        default:
            v2 -= (-(v2))
        }
        v3 = (v3 % ((v3) | 1))
        var v17 uint32 = 4294967295
        _ = fi14
        _ = v17
    }
    if false {
        if true {
            v18 := GenS1{F0: 0, F1: v2, F2: ("" + "\n\t")}
            _ = v18
        }
        v3 += 4294967295
    }
    v19 := GenS2{F0: (v2 >> ((2) & 63))}
    v3 = 627
    fmt.Println(v2, v2)
    for fi20 := 0; fi20 < 13; fi20++ {
        v21 := v1[(fi20 + fi20)]
        if false {
            var v22 bool = true
            var v23 int16 = -32768
            v24 := []int{(fi20 + (fi20 ^ 32767)), -9223372036854775808, (fi20 - int(v2)), (^(fi20))}
            _ = v22
            _ = v23
            _ = v24
        } else {
            for fi25 := 0; fi25 < 13; fi25++ {
                fmt.Println(v2)
                v3 = v3
                _ = fi25
            }
        }
        if false {
            v26 := v1[(fi20 * (fi20 - fi20))]
            for fi27 := 0; fi27 < 10; fi27++ {
                var v28 int16 = 0
                if (true || false) {
                    fmt.Println(v3, v26)
                    v21 = (v26 + ("foo bar" + v21))
                }
                var v29 uint16 = 0
                if (!false) {
                    var v30 int64 = 9223372036854775807
                    v31 := GenS2{F0: (-(v2))}
                    fmt.Println(v30, v26, v30)
                    _ = v30
                    _ = v31
                } else {
                    fmt.Println(v3, v26, v26)
                    v28 = 817
                    if (v29 == 2) {
                        if (42 >= 765) {
                            if true {
                                fmt.Println(v26, v3, v28)
                                v21 = (v21 + v21)
                                fmt.Println(v2)
                            }
                            v28 = -32768
                        }
                    } else {
                        v32 := GenS4{F0: (0 > (v28 | (v28 - -1)))}
                        v33 := GenS5{F0: 18446744073709551615, F1: (v3 * 1), F2: 2147483648}
                        _ = v32
                        _ = v33
                    }
                    var v34 uint8 = 42
                    _ = v34
                }
                if false {
                    v35 := []int{}
                    _ = v35
                } else {
                    fmt.Println(fi20, v2)
                    var v36 int32 = -1
                    v3 = v3
                    _ = v36
                }
                _ = fi27
                _ = v28
                _ = v29
            }
            _ = v26
        }
        _ = fi20
        _ = v21
    }
    if (1e10 >= (0.0 - (float32((0.0 / 0.1)) + 0.1))) {
        var v37 uint16 = 255
        v37 += (v37 * 42)
        switch ("\n\t" + "") {
        case "hello":
            var v38 uint64 = 1000
            _ = v38
        }
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
}

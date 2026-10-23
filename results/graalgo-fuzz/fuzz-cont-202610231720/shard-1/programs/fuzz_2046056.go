package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 string }
type GenS2 struct { F0 uint64; F1 int16; F2 int32 }
type GenS3 struct { F0 int16; F1 int8; F2 uint16 }
type GenS4 struct { F0 int64; F1 int16; F2 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int; F1 uint16; F2 int }
func genFn0(p0 int, p1 uint32) int {
    v4 := GenS1{F0: 0, F1: 255, F2: "hello"}
    for fi5 := 0; fi5 < 20; fi5++ {
        if (2 <= 0) {
            p1 = (p1 * (p1 * (p1 * (^(p1)))))
            var v6 uint16 = 0
            var v7 int64 = 42
            _ = v6
            _ = v7
        }
        var v8 int8 = 68
        _ = fi5
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v4
    return 610
}
func genFn1(p0 int64, p1 uint, p2 int8) uint64 {
    var v9 int64 = (p0 - 42)
    var v10 uint = (p1 >> ((p1) & 63))
    p2 -= 42
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 4294967295
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = -1
    var v3 int32 = 110
    v3 -= v3
    for fi11 := 0; fi11 < 11; fi11++ {
        v12 := []int{(-(fi11)), fi11, fi11}
        _ = fi11
        _ = v12
    }
    if true {
        for fi13 := 0; fi13 < 12; fi13++ {
            for fi14 := 0; fi14 < 19; fi14++ {
                if true {
                    v3 = (v3 + v3)
                    v15 := GenS1{F0: 18446744073709551615, F1: 0, F2: "hello"}
                    v16 := GenS2{F0: 2, F1: 1, F2: (^(v3))}
                    var v17 int16 = 1
                    var v18 uint16 = 1
                    _ = v15
                    _ = v16
                    _ = v17
                    _ = v18
                } else {
                    fmt.Println(fi14)
                }
                switch "" {
                case "foo bar":
                    v19 := GenS3{F0: 42, F1: (v2 + v2), F2: 255}
                    _ = v19
                default:
                    v20 := GenS4{F0: 42, F1: -1, F2: fi13}
                    _ = v20
                }
                _ = len(v1)
                v3 = (v3 & 42)
                v3 = v3
                v3 = v3
                _ = fi14
            }
            if false {
                var v21 uint64 = genFn1(-432, 458, (v2 << ((255) & 63)))
                fmt.Println(v21, v2)
                _ = v21
            } else {
                var v22 int8 = 42
                v3 = -2147483648
                for fi23 := 0; fi23 < 3; fi23++ {
                    v2 -= v2
                    if (false && (1 != 568)) {
                        switch (fi23 / (((^(fi23))) | 1)) {
                        case 0:
                            var v24 int8 = (^(v2))
                            _ = v24
                        case 1:
                            var v25 bool = (!("hello" != "hello"))
                            _ = v25
                        }
                        if false {
                            var v26 int8 = int8(fi13)
                            v22 = (v26 + v26)
                            _ = v26
                        } else {
                            v3 += (v3 + 1)
                        }
                    }
                    var v27 float32 = 1.5
                    v2 += (v2 - v22)
                    _ = fi23
                    _ = v27
                }
                _ = v22
            }
            _ = fi13
        }
        var v28 int32 = (v3 % (((v3 / ((42) | 1))) | 1))
        var v29 int = 127
        _ = v28
        _ = v29
    }
    v2 = v2
    v3 = 2147483647
    if (102 <= 42) {
        v3 -= v3
        var v30 uint64 = 497
        _ = v30
    } else {
        for fi31 := 0; fi31 < 3; fi31++ {
            v3 -= (-(v3))
            v3 += v3
            _ = fi31
        }
        v32 := []int{-2147483648, 906, 654, 128}
        var v33 uint32 = 850
        _ = v32
        _ = v33
    }
    switch 127 {
    case -878:
        if (false && true) {
            var v34 int8 = v2
            v35 := GenS4{F0: -575, F1: -1, F2: -309}
            var v36 int32 = (-(v3))
            _ = v34
            _ = v35
            _ = v36
        }
    case 42:
        v3 = v3
    default:
        _ = len(v1)
    }
    v3 = v3
    v2 -= -11
    var v37 int64 = -9223372036854775808
    v38 := GenS5{F0: 42}
    v39 := []int{0, 4294967295}
    var v40 uint8 = 0
    v40 = v40
    v3 = (v3 + v3)
    v41 := GenS1{F0: 2, F1: 1, F2: "\n\t"}
    v42 := GenS3{F0: 42, F1: (v2 & v2), F2: 65535}
    if false {
        v43 := GenS6{F0: 42, F1: 32767, F2: 4294967295}
        _ = v43
    } else {
        v3 += 42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v40
    _ = v41
    _ = v42
}

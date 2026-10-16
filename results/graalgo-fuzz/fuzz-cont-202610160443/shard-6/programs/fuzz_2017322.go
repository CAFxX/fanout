package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint8 }
type GenS2 struct { F0 int8; F1 int; F2 uint64 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 uint32 }
func genFn0(p0 string, p1 uint32) uint32 {
    var v5 bool = (false || true)
    v6 := []int{}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (p1 & p1)
}
func genFn1() int16 {
    v7 := GenS1{F0: 0, F1: 1, F2: 0}
    var v8 uint8 = 255
    _ = v7
    _ = v8
    return 32767
}

func main() {
    var v1 float32 = 1.5
    var v2 int16 = 189
    var v3 float32 = -2.25
    v4 := make(map[uint32]int)
    var v9 int32 = 0
    switch "\n\t" {
    case "hello":
        if (2 == 0) {
            var v10 int64 = 1
            var v11 uint = 4294967295
            if true {
                v11 += v11
            } else {
                var v12 uint16 = 1
                switch 1 {
                case 140:
                    var v13 int16 = (v2 + genFn1())
                    _ = v13
                case 9223372036854775807:
                    v9 = 0
                }
                v14 := GenS1{F0: (v10 ^ 0), F1: v2, F2: 42}
                _ = v12
                _ = v14
            }
            _ = v10
            _ = v11
        }
    case "":
        var v15 int64 = 1
        _ = v15
    case "\n\t":
        v16 := GenS2{F0: 100, F1: 1, F2: 2}
        _ = v16
    default:
        v17 := []int{}
        _ = v17
    }
    switch ("hello" + "") {
    case "\n\t":
        v1 = float32(float64(v3))
    case "foo bar":
        v2 = v2
    case "a":
        _ = len(v4)
    default:
        var v18 uint8 = 42
        _ = v18
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        v20 := GenS3{F0: 543, F1: -374}
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        switch "" {
        case "hello":
            for fi22 := 0; fi22 < 16; fi22++ {
                if false {
                    _ = len(v4)
                    v1 -= -2.25
                    fmt.Println(v1, v2)
                } else {
                    v23 := GenS4{F0: 1}
                    _ = v23
                }
                delete(v4, 2)
                v24 := []int{(fi22 - -389)}
                _ = fi22
                _ = v24
            }
        case "":
            fmt.Println(v2, fi21)
        default:
            v25 := GenS2{F0: 42, F1: fi21, F2: 0}
            _ = v25
        }
        _ = fi21
    }
    v1 = v1
    switch 0 {
    case 14:
        v26 := GenS1{F0: 9223372036854775807, F1: -885, F2: 0}
        _ = v26
    case -2147483648:
        if true {
            for fi27 := 0; fi27 < 14; fi27++ {
                v28 := GenS5{F0: 0, F1: 1}
                _ = fi27
                _ = v28
            }
            var v29 uint = 1
            _ = v29
        } else {
            v9 = v9
        }
    case 819:
        var v30 int = 32767
        _ = v30
    default:
        var v31 uint16 = 65535
        _ = v31
    }
    v32 := GenS4{F0: 550}
    v2 -= v2
    var v33 int16 = 32767
    for fi34 := 0; fi34 < 2; fi34++ {
        for fi35 := 0; fi35 < 5; fi35++ {
            fmt.Println(v9, v9)
            if false {
                switch "" {
                case "":
                    _ = len(v4)
                case "a":
                    var v36 int16 = genFn1()
                    _ = v36
                case "hello":
                    v2 = 32767
                }
            }
            switch 42 {
            case 1:
                fmt.Println(v33)
            case -9223372036854775808:
                if false {
                    v37 := GenS5{F0: -100, F1: 2}
                    if (true && ("a" < "")) {
                        switch (fi34 + (fi34 / (((fi35 / (((fi34 ^ (-(fi35)))) | 1))) | 1))) {
                        case 42:
                            var v38 int32 = (^(v9))
                            _ = v38
                        default:
                            _ = len(v4)
                        }
                        var v39 uint16 = 6
                        _ = v39
                    }
                    _ = v37
                }
            default:
                if false {
                    _ = len(v4)
                }
            }
            _ = fi35
        }
        for fi40 := 0; fi40 < 6; fi40++ {
            v41 := GenS5{F0: 10, F1: 183}
            v33 = genFn1()
            v1 = v3
            _ = fi40
            _ = v41
        }
        switch "a" {
        case "a":
            v33 = (v33 / ((-1) | 1))
        case "\n\t":
            v33 -= (v2 >> ((18446744073709551615) & 63))
        }
        _ = fi34
    }
    v42 := v4[765]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v42
    _ = v9
}

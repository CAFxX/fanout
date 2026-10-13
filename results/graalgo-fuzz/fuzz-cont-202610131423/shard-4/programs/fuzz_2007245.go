package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint32; F1 uint64; F2 bool }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint32; F1 int16 }
type GenS8 struct { F0 bool; F1 int8; F2 bool }
type GenS9 struct { F0 int16 }
func genFn0(p0 int16, p1 int32, p2 uint) bool {
    p1 = p1
    if true {
        if false {
            for fi5 := 0; fi5 < 4; fi5++ {
                var v6 int8 = -7
                var v7 uint = (p2 << ((p2) & 63))
                for fi8 := 0; fi8 < 17; fi8++ {
                    v9 := []int{}
                    v10 := GenS1{F0: 159, F1: 128}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                _ = fi5
                _ = v6
                _ = v7
            }
            if (0 != 9223372036854775807) {
                var v11 int32 = -2147483648
                _ = v11
            } else {
                p0 += -1
                var v12 int64 = 9223372036854775807
                _ = v12
            }
            var v13 int64 = -1
            _ = v13
        } else {
            p0 = p0
            if (2 <= 42) {
                v14 := []int{-712, 32768, 677, -9223372036854775808}
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
                _ = v14
            } else {
                switch 29 {
                case 128:
                    v15 := GenS1{F0: 1, F1: 42}
                    _ = v15
                case 32768:
                    switch -991 {
                    case -718:
                        var v16 int = 127
                        _ = v16
                    case -742:
                        v17 := GenS2{F0: 377, F1: 1, F2: false}
                        _ = v17
                    case 0:
                        p2 -= p2
                    default:
                        for fi18 := 0; fi18 < 5; fi18++ {
                            v19 := []int{int(p0), (fi18 ^ fi18)}
                            v20 := GenS3{F0: "foo bar", F1: false}
                            p1 += (p1 % (((p1 & (p1 * (p1 % (((p1 * (p1 * -641))) | 1))))) | 1))
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                    }
                case 0:
                    v21 := GenS1{F0: 65535, F1: 0}
                    _ = v21
                default:
                    p1 = p1
                }
                var v22 uint = (p2 / (((p2 % ((p2) | 1))) | 1))
                var v23 uint16 = 899
                _ = v22
                _ = v23
            }
        }
    }
    v24 := GenS4{F0: 9223372036854775807}
    switch "a" {
    case "foo bar":
        for fi25 := 0; fi25 < 17; fi25++ {
            var v26 int8 = 42
            v27 := GenS5{F0: p0, F1: (^(v26))}
            _ = fi25
            _ = v26
            _ = v27
        }
    case "a":
        v28 := GenS6{F0: ("a" + "a")}
        _ = v28
    case "hello":
        fmt.Println(p1)
    default:
        p2 -= p2
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return ((740 > 42) && false)
}
func genFn1() int8 {
    v29 := GenS3{F0: "", F1: (42 != 0)}
    var v30 int = 787
    var v31 int8 = 0
    _ = v29
    _ = v30
    _ = v31
    return v31
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 0
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]uint32)
    v2 += (v2 + v2)
    var v32 string = "a"
    switch v32 {
    case "foo bar":
        var v33 uint64 = 4294967295
        _ = v33
    case "a":
        var v34 int = 32767
        _ = v34
    case "":
        if true {
            var v35 int8 = 127
            fmt.Println(v35)
            var v36 uint64 = 2
            _ = v35
            _ = v36
        }
    }
    v37 := []int{32768}
    v38 := GenS7{F0: 0, F1: 1000}
    v39 := GenS4{F0: 1}
    var v40 uint16 = 0
    v41 := GenS8{F0: false, F1: 42, F2: (!(v32 > (v32 + (v32 + ""))))}
    v42 := GenS9{F0: 42}
    switch 357 {
    case 915:
        var v43 float32 = (3.14159 / 0.1)
        _ = v43
    case 0:
        if len(v37) > 0 {
            v37[0] = 1
        }
    default:
        for fi44 := 0; fi44 < 17; fi44++ {
            if true {
                switch "foo bar" {
                case "a":
                    if true {
                        fmt.Println(v40)
                        v40 += (v40 ^ v40)
                    }
                default:
                    fmt.Println(fi44, fi44, v40)
                }
                v45 := GenS8{F0: false, F1: 42, F2: false}
                _ = v45
            }
            _ = fi44
        }
    }
    var v46 uint = 988
    v40 -= 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v46
}

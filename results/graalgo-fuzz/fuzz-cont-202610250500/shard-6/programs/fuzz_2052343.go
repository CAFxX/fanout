package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint64; F1 uint; F2 string }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 int64; F1 int8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64 }
func genFn0(p0 int32, p1 int, p2 int8) int8 {
    v7 := []int{p1, (^(p1)), (p1 | p1), 32767}
    var v8 uint8 = 255
    var v9 uint32 = 325
    _ = len(v7)
    switch "foo bar" {
    case "a":
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 bool = (false && (v9 < v9))
            v12 := GenS1{F0: (fi10 % (((^(fi10))) | 1)), F1: 65535, F2: -32768}
            v13 := GenS1{F0: 128, F1: uint16(v9), F2: -1}
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
    case "":
        for fi14 := 0; fi14 < 5; fi14++ {
            v15 := GenS2{F0: 1, F1: 0, F2: "hello"}
            _ = fi14
            _ = v15
        }
    case "\n\t":
        p2 += (p2 ^ 0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return p2
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{42}
    var v3 uint64 = 27
    var v4 int16 = 1
    v5 := []string{}
    v6 := make(map[uint32]string)
    v3 -= v3
    v16 := GenS3{F0: 418, F1: "a"}
    var v17 uint8 = 106
    v18 := GenS4{F0: (-(v4)), F1: v3}
    if false {
        if true {
            if (v4 != (v4 | 42)) {
                v19 := GenS4{F0: v4, F1: (v3 | v3)}
                if true {
                    var v20 uint16 = 42
                    v4 -= (v4 >> ((491) & 63))
                    _ = v20
                } else {
                    v3 = v3
                    if (v3 == v3) {
                        var v21 int8 = -1
                        fmt.Println(v17, v4)
                        _ = v21
                    }
                }
                _ = v19
            }
            if true {
                v22 := GenS5{F0: 0, F1: 34}
                _ = v22
            }
            if false {
                var v23 int8 = -128
                switch 4294967295 {
                case 127:
                    v24 := GenS5{F0: 1, F1: (v23 ^ -1)}
                    _ = v24
                case 4294967295:
                    var v25 string = ("" + "a")
                    _ = v25
                case 32768:
                    delete(v6, 26)
                }
                if (1 > v17) {
                    switch 9223372036854775807 {
                    case 4294967295:
                        var v26 float32 = (0.1 + 1.5)
                        _ = v26
                    case 32768:
                        v27 := GenS2{F0: v3, F1: 209, F2: ""}
                        _ = v27
                    case 0:
                        var v28 uint8 = (v17 & 255)
                        _ = v28
                    }
                    var v29 float64 = float64((-2.25 * float32((3.14159 + -2.25))))
                    _ = v29
                } else {
                    fmt.Println(v4)
                    if true {
                        for fi30 := 0; fi30 < 8; fi30++ {
                            v1[fi30] = ((v17 == 190) || (true && true))
                            fmt.Println(v3, v23)
                            _ = fi30
                        }
                    }
                    v31 := GenS6{F0: true}
                    _ = v31
                }
                _ = v23
            } else {
                delete(v6, 65535)
            }
        }
    }
    switch 127 {
    case 654:
        if false {
            if false {
                v4 -= v4
                v17 += 42
                var v32 int = 1
                _ = v32
            } else {
                v33 := GenS1{F0: 4294967295, F1: 0, F2: (v4 + v4)}
                _ = v33
            }
            switch "" {
            case "hello":
                var v34 int32 = 0
                _ = v34
            case "\n\t":
                for fi35 := 0; fi35 < 13; fi35++ {
                    fmt.Println(v4, fi35)
                    v36 := GenS3{F0: 1, F1: "hello"}
                    _ = fi35
                    _ = v36
                }
            }
            fmt.Println(v3, v3, v4)
        } else {
            if true {
                v37 := GenS2{F0: (v3 ^ (v3 & (^(v3)))), F1: 4294967295, F2: "\n\t"}
                _ = v37
            }
        }
    }
    for fi38 := 0; fi38 < 5; fi38++ {
        for fi39 := 0; fi39 < 10; fi39++ {
            delete(v1, 2147483647)
            v4 = (-(v4))
            _ = fi39
        }
        delete(v1, (fi38 >> ((2) & 63)))
        _ = fi38
    }
    v6[42] = ""
    var v40 int8 = -128
    if (!(true || (590 > -1))) {
        v6[1] = "foo bar"
        var v41 int32 = 2147483647
        var v42 int8 = (v40 % ((127) | 1))
        _ = v41
        _ = v42
    }
    fmt.Println(v3)
    v43 := GenS1{F0: 1, F1: 161, F2: 369}
    v4 = (v4 - (v4 + (v4 & v4)))
    for fi44 := 0; fi44 < 18; fi44++ {
        _ = len(v1)
        _ = fi44
    }
    for fi45 := 0; fi45 < 3; fi45++ {
        v46 := GenS7{F0: v3}
        v17 -= (v17 ^ v17)
        _ = fi45
        _ = v46
    }
    var v47 uint16 = 764
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v43
    _ = v47
    _ = v5
    _ = v6
}

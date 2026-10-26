package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32 }

func main() {
    var v1 string = "foo bar"
    var v2 uint = 4294967295
    v3 := make(map[uint32]uint32)
    var v4 int = 2147483647
    v5 := make(map[int]uint32)
    v6 := []string{"a"}
    if true {
        for fi7 := 0; fi7 < 9; fi7++ {
            if ((!((v2 | (v2 / ((18446744073709551615) | 1))) != 953)) && false) {
                switch 129 {
                case 381:
                    for fi8 := 0; fi8 < 3; fi8++ {
                        fmt.Println(v4, v4)
                        v9 := GenS1{F0: 42, F1: 1}
                        _ = fi8
                        _ = v9
                    }
                }
                v10 := GenS1{F0: 34, F1: 42}
                _ = v10
            }
            var v11 int32 = 2147483647
            delete(v5, fi7)
            _ = fi7
            _ = v11
        }
        v2 += (v2 * (v2 % ((v2) | 1)))
        v3[42] = 618
    } else {
        if (42 == 127) {
            for fi12 := 0; fi12 < 4; fi12++ {
                if true {
                    switch fi12 {
                    case 109:
                        if (-32768 <= -489) {
                            if (-957 != 9223372036854775807) {
                                v1 = v1
                                fmt.Println(v4, v2)
                            } else {
                                v13 := GenS2{F0: 127}
                                _ = v13
                            }
                        }
                    default:
                        var v14 bool = false
                        _ = v14
                    }
                    v2 = v2
                    var v15 int64 = -1
                    _ = v15
                } else {
                    fmt.Println(v1, fi12)
                }
                v4 = 1
                _ = fi12
            }
            for fi16 := 0; fi16 < 1; fi16++ {
                var v17 string = ("\n\t" + "hello")
                _ = fi16
                _ = v17
            }
        }
        for fi18 := 0; fi18 < 16; fi18++ {
            v1 = (v1 + ("" + "hello"))
            _ = fi18
        }
    }
    v1 = "hello"
    if (false && (!(136 <= 12))) {
        var v19 float32 = 0.0
        var v20 uint8 = 1
        _ = v19
        _ = v20
    } else {
        v4 = (v4 >> ((v2) & 63))
        for fi21 := 0; fi21 < 1; fi21++ {
            if (1 <= 18446744073709551615) {
                var v22 int8 = -110
                var v23 int64 = -1
                v24 := GenS3{F0: 1}
                _ = v22
                _ = v23
                _ = v24
            }
            _ = fi21
        }
        v25 := v3[563]
        _ = v25
    }
    v26 := GenS3{F0: 42}
    if false {
        var v27 uint16 = 42
        _ = v27
    } else {
        v28 := GenS4{F0: true}
        switch int(v2) {
        case 42:
            var v29 uint64 = 2
            _ = v29
        case 806:
            v30 := GenS5{F0: -2147483648}
            _ = v30
        }
        _ = v28
    }
    v31 := GenS2{F0: 0}
    var v32 uint32 = 2
    var v33 int8 = 100
    var v34 uint16 = 32767
    v3[v32] = (v32 - v32)
    v6 = append(v6, "\n\t")
    v2 -= 4294967295
    v2 = (v2 / ((v2) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}

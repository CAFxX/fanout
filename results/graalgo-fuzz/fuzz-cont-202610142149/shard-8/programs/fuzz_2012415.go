package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 string; F2 int16 }
type GenS3 struct { F0 int; F1 int16; F2 string }
type GenS4 struct { F0 int64; F1 uint64; F2 bool }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool; F1 uint8; F2 int16 }

func main() {
    var v1 int8 = -15
    v2 := make(map[string]uint32)
    var v3 int = 32767
    var v4 uint64 = 18446744073709551615
    v5 := make(map[string]int)
    v3 -= v3
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := []int{}
        var v8 uint = 0
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 float32 = 0.1
        _ = fi9
        _ = v10
    }
    fmt.Println(v4, v1)
    var v11 uint32 = 2147483648
    v1 -= (^(v1))
    if (v1 == (v1 | (v1 & v1))) {
        switch v3 {
        case -1:
            v12 := GenS1{F0: v3}
            _ = v12
        case 32768:
            v13 := GenS2{F0: 0, F1: "foo bar", F2: 32767}
            _ = v13
        default:
            v1 = (v1 << ((0) & 63))
        }
        var v14 uint32 = (v11 | 4294967295)
        _ = v14
    }
    if false {
        if false {
            fmt.Println(v11, v4, v1)
            v15 := GenS3{F0: (v3 << ((42) & 63)), F1: 42, F2: "hello"}
            var v16 uint32 = 0
            _ = v15
            _ = v16
        }
        v17 := GenS4{F0: -9223372036854775808, F1: (v4 - (v4 - (^(v4)))), F2: (1.5 >= (-2.25 + -2.25))}
        if false {
            var v18 float64 = 0.0
            for fi19 := 0; fi19 < 14; fi19++ {
                v1 -= 103
                _ = fi19
            }
            _ = v18
        }
        _ = v17
    } else {
        for fi20 := 0; fi20 < 8; fi20++ {
            delete(v2, ("hello" + ""))
            delete(v5, "hello")
            _ = fi20
        }
        v11 += (v11 - (v11 | (^(v11))))
        v21 := GenS2{F0: -2147483648, F1: ("a" + ("" + "foo bar")), F2: 1000}
        _ = v21
    }
    delete(v5, "hello")
    if (((!((0.1 * 0.0) < float64(0.1))) || false) && ((v1 / ((v1) | 1)) < 1)) {
        var v22 float64 = (0.1 - 3.14159)
        v4 -= (v4 + v4)
        for fi23 := 0; fi23 < 12; fi23++ {
            v24 := GenS5{F0: "\n\t"}
            _ = fi23
            _ = v24
        }
        _ = v22
    }
    if (!(1e10 < ((1.5 + ((-2.25 / 1e10) + 0.0)) / -2.25))) {
        for fi25 := 0; fi25 < 11; fi25++ {
            v26 := GenS2{F0: 2147483647, F1: "\n\t", F2: 0}
            if (!(v4 >= v4)) {
                switch fi25 {
                case 0:
                    v27 := GenS6{F0: (32767 < 42), F1: 128, F2: 1}
                    _ = v27
                default:
                    if true {
                        for fi28 := 0; fi28 < 10; fi28++ {
                            var v29 uint16 = 32767
                            v30 := GenS1{F0: fi25}
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                    }
                }
                v3 = (-(fi25))
                var v31 float32 = 0.1
                _ = v31
            } else {
                v3 = v3
                var v32 int = int(v4)
                _ = v32
            }
            v33 := GenS6{F0: true, F1: 255, F2: -569}
            _ = fi25
            _ = v26
            _ = v33
        }
        delete(v5, ("foo bar" + "hello"))
        for fi34 := 0; fi34 < 19; fi34++ {
            if false {
                var v35 uint64 = v4
                v35 = (v4 / ((946) | 1))
                _ = v35
            }
            if (!false) {
                for fi36 := 0; fi36 < 3; fi36++ {
                    fmt.Println(v11, v1)
                    _ = fi36
                }
            } else {
                v37 := []int{(v3 - 9223372036854775807)}
                _ = v37
            }
            _ = fi34
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

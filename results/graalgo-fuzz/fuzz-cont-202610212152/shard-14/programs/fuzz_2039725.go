package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int32 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8; F1 uint16 }
func genFn0(p0 uint8) uint16 {
    fmt.Println(p0)
    p0 = 1
    fmt.Println(p0)
    fmt.Println(p0)
    if ((false && (true && false)) && false) {
        p0 = p0
    } else {
        v6 := GenS1{F0: (^(p0)), F1: 946, F2: 1}
        var v7 int16 = 32767
        _ = v6
        _ = v7
    }
    _ = p0
    return uint16(p0)
}
func genFn1(p0 int, p1 uint16, p2 int) string {
    fmt.Println(p1)
    v8 := GenS1{F0: 0, F1: 9223372036854775807, F2: 0}
    v9 := GenS2{F0: 42, F1: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return "hello"
}

func main() {
    var v1 string = "\n\t"
    var v2 float64 = 1.5
    v3 := make(map[uint32]uint32)
    var v4 uint = 84
    v5 := make(map[uint32]uint32)
    v4 = 755
    var v10 uint = v4
    var v11 uint64 = 1
    switch "hello" {
    case "a":
        fmt.Println(v10, v11)
    case "hello":
        v11 += (v11 + (^(v11)))
    default:
        for fi12 := 0; fi12 < 2; fi12++ {
            v13 := GenS3{F0: (fi12 + fi12), F1: -395}
            switch v1 {
            case "foo bar":
                if true {
                    v14 := []int{}
                    fmt.Println(v1)
                    _ = v14
                } else {
                    v5[2] = 4294967295
                }
            }
            switch fi12 {
            case 32768:
                delete(v3, 5)
            case 2147483647:
                var v15 uint = (v10 / (((v10 % (((v10 % ((v4) | 1))) | 1))) | 1))
                _ = v15
            case 803:
                for fi16 := 0; fi16 < 18; fi16++ {
                    var v17 int = (fi16 << (((v10 | (v10 & (v10 >> ((v4) & 63))))) & 63))
                    fmt.Println(v17, v17)
                    v18 := GenS3{F0: 42, F1: 372}
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            default:
                v19 := GenS4{F0: 776}
                _ = v19
            }
            _ = fi12
            _ = v13
        }
    }
    if (true || (42 < 0)) {
        v20 := []int{128, 2147483647, -9223372036854775808}
        if (v1 <= "a") {
            var v21 int64 = 195
            var v22 uint64 = v11
            for fi23 := 0; fi23 < 1; fi23++ {
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
                v24 := GenS5{F0: 0}
                v25 := GenS4{F0: 925}
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = v21
            _ = v22
        } else {
            if len(v20) > 0 {
                v20[0] = 999
            }
        }
        _ = v20
    }
    fmt.Println(v10, v11)
    v2 += (v2 / v2)
    if true {
        v1 = v1
        if true {
            v26 := GenS5{F0: 32767}
            if true {
                v27 := GenS1{F0: 1, F1: int64(v11), F2: 0}
                _ = v27
            } else {
                fmt.Println(v1, v11)
                v4 = v4
            }
            _ = v26
        } else {
            var v28 int8 = 107
            _ = v28
        }
    }
    var v29 bool = false
    var v30 uint64 = 622
    fmt.Println(v29)
    fmt.Println(v4, v4)
    v1 = ("a" + v1)
    fmt.Println(v30, v10)
    if (42 <= 40) {
        var v31 float64 = v2
        _ = v31
    }
    fmt.Println(v10, v10)
    if (!v29) {
        if (-1 > 109) {
            var v32 float32 = -2.25
            _ = v32
        }
        if false {
            var v33 bool = ((-656 > 1) && (v29 || v29))
            _ = v33
        } else {
            for fi34 := 0; fi34 < 16; fi34++ {
                v35 := GenS5{F0: 42}
                v36 := GenS6{F0: -128, F1: 1}
                fmt.Println(v10)
                _ = fi34
                _ = v35
                _ = v36
            }
        }
    } else {
        var v37 uint = uint(v30)
        v5[uint32(v11)] = 2
        _ = v37
    }
    for fi38 := 0; fi38 < 7; fi38++ {
        var v39 int = int(v11)
        fmt.Println(v29)
        v4 += 42
        _ = fi38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}

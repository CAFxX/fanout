package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS5 struct { F0 string; F1 bool; F2 int8 }
func genFn0(p0 int16, p1 uint8, p2 uint32) int64 {
    fmt.Println(p1, p0)
    v6 := GenS1{F0: 32768, F1: 65535}
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := []int{fi7, 128, (fi7 << ((0) & 63))}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return int64(p0)
}
func genFn1(p0 int16, p1 int) int32 {
    v9 := GenS2{F0: -2147483648, F1: 1}
    var v10 uint32 = 2
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return -280
}

func main() {
    var v1 uint8 = 42
    v2 := []int{2147483647, 4294967295, 32767, 583}
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]string)
    var v5 uint32 = 42
    if (-93 > 127) {
        fmt.Println(v1, v1)
        var v11 uint16 = 65535
        switch 160 {
        case -766:
            for fi12 := 0; fi12 < 19; fi12++ {
                var v13 int = fi12
                _ = fi12
                _ = v13
            }
        default:
            fmt.Println(v11)
        }
        _ = v11
    } else {
        v14 := GenS3{F0: -32768}
        fmt.Println(v1, v1)
        _ = v14
    }
    v3[((("foo bar" + ("a" + ("a" + "a"))) + "a") + "foo bar")] = v5
    var v15 uint8 = (v1 & (^(v1)))
    fmt.Println(v1, v15, v1)
    v1 = (v15 ^ (v15 << ((18446744073709551615) & 63)))
    _ = len(v3)
    var v16 uint = 42
    v16 += 982
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 bool = true
        _ = fi17
        _ = v18
    }
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    fmt.Println(v16, v5)
    switch 1 {
    case -1:
        v19 := GenS3{F0: 0}
        _ = v19
    default:
        if false {
            if (genFn1(-32768, 42) < 2147483647) {
                v20 := GenS4{F0: 65535, F1: v5, F2: 472}
                var v21 uint64 = 18446744073709551615
                v22 := GenS5{F0: "hello", F1: true, F2: -52}
                _ = v20
                _ = v21
                _ = v22
            } else {
                var v23 int16 = -1
                var v24 string = "a"
                v5 -= (v5 << (((v16 - (v16 << ((v16) & 63)))) & 63))
                _ = v23
                _ = v24
            }
        }
    }
    v25 := v4[1]
    if false {
        if (!true) {
            var v26 int32 = 2147483647
            _ = v26
        } else {
            for fi27 := 0; fi27 < 14; fi27++ {
                for fi28 := 0; fi28 < 14; fi28++ {
                    v29 := GenS2{F0: 310, F1: (^(v1))}
                    v30 := GenS5{F0: v25, F1: false, F2: 115}
                    v16 = v16
                    var v31 uint = 706
                    _ = fi28
                    _ = v29
                    _ = v30
                    _ = v31
                }
                var v32 uint = 4294967295
                _ = fi27
                _ = v32
            }
            v33 := GenS4{F0: 657, F1: 747, F2: -301}
            _ = v33
        }
        v1 = 42
    } else {
        var v34 uint64 = uint64(v15)
        var v35 int64 = -420
        _ = v34
        _ = v35
    }
    var v36 bool = ((^(v5)) <= (v5 >> ((v16) & 63)))
    var v37 int = 9223372036854775807
    v38 := GenS3{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}

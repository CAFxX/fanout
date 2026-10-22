package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint32; F1 int16; F2 uint64 }
type GenS4 struct { F0 uint; F1 uint8 }
func genFn0() int32 {
    var v6 string = "a"
    v7 := []int{}
    v6 = v6
    v6 = (v6 + ("hello" + v6))
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 uint64, p1 uint32, p2 uint16) bool {
    v8 := GenS1{F0: 1}
    p0 = p0
    p1 -= (^(p1))
    p2 = (p2 >> ((4294967295) & 63))
    if true {
        p2 += p2
    } else {
        v9 := GenS2{F0: 42, F1: 58}
        switch "\n\t" {
        case "a":
            for fi10 := 0; fi10 < 11; fi10++ {
                v11 := GenS2{F0: -1, F1: 127}
                if true {
                    var v12 int32 = 1
                    v13 := []int{(fi10 - (fi10 & fi10)), fi10, int(p1)}
                    switch ("foo bar" + "foo bar") {
                    case "foo bar":
                        v12 += -2147483648
                    }
                    _ = v12
                    _ = v13
                } else {
                    v14 := GenS3{F0: (p1 >> ((uint(p0)) & 63)), F1: 32767, F2: p0}
                    _ = v14
                }
                _ = fi10
                _ = v11
            }
        case "":
            if (int32(p0) <= -814) {
                v15 := []int{98, -9223372036854775808}
                var v16 uint32 = 2
                for fi17 := 0; fi17 < 8; fi17++ {
                    if false {
                        var v18 int32 = 54
                        for fi19 := 0; fi19 < 4; fi19++ {
                            var v20 int8 = 0
                            var v21 float64 = (3.14159 - 3.14159)
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                        p2 = (p2 * (p2 << ((42) & 63)))
                        _ = v18
                    } else {
                        v22 := GenS3{F0: (p1 | (v16 / ((v16) | 1))), F1: 316, F2: p0}
                        p1 -= 4294967295
                        _ = v22
                    }
                    _ = fi17
                }
                _ = v15
                _ = v16
            }
        case "hello":
            v23 := GenS2{F0: 419, F1: -76}
            _ = v23
        default:
            p0 = 4294967295
        }
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return ((!false) && true)
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{"foo bar", "a", "a"}
    var v3 int32 = 0
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]int)
    var v24 uint = 42
    v25 := GenS2{F0: genFn0(), F1: 42}
    var v26 bool = (1 != 114)
    fmt.Println(v3)
    var v27 int64 = 0
    v2 = append(v2, ("hello" + "\n\t"))
    if (!v26) {
        v28 := v1[-9223372036854775808]
        _ = v28
    } else {
        if true {
            var v29 uint8 = 164
            var v30 int64 = (v27 << (((v24 >> (((v24 % (((^(v24))) | 1))) & 63))) & 63))
            fmt.Println(v30)
            _ = v29
            _ = v30
        }
        fmt.Println(v3)
        v31 := GenS4{F0: v24, F1: 2}
        _ = v31
    }
    if (-2.25 >= (0.0 - 0.0)) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v32 int32 = (^(v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}

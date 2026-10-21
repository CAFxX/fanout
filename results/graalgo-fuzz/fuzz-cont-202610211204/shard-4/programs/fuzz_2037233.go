package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 uint }
func genFn0(p0 int64, p1 uint32) bool {
    v5 := []int{260}
    var v6 uint8 = 42
    var v7 int64 = p0
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return false
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]uint32)
    var v3 int64 = 9223372036854775807
    var v4 string = "\n\t"
    fmt.Println(v4, v3)
    for fi8 := 0; fi8 < 12; fi8++ {
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 int32 = 0
            v1[333] = (false && true)
            v10 = v10
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v11 := []int{}
    _ = len(v11)
    for fi12 := 0; fi12 < 19; fi12++ {
        v4 = v4
        var v13 bool = (v3 <= (v3 | (v3 | v3)))
        v4 = v4
        _ = fi12
        _ = v13
    }
    v4 = v4
    var v14 string = v4
    if true {
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 int32 = -2147483648
            for fi17 := 0; fi17 < 9; fi17++ {
                switch v4 {
                case "hello":
                    if (((3.14159 < float32(1.5)) && (true || false)) || (-2.25 > 1.5)) {
                        _ = len(v11)
                        if (false || (true || false)) {
                            var v18 int32 = (v16 >> ((18446744073709551615) & 63))
                            _ = v18
                        } else {
                            delete(v1, 215)
                            var v19 uint = 255
                            _ = v19
                        }
                    } else {
                        v20 := GenS1{F0: (v16 + v16)}
                        var v21 int8 = 42
                        v3 = int64(fi17)
                        _ = v20
                        _ = v21
                    }
                case "foo bar":
                    var v22 string = "hello"
                    _ = v22
                }
                v14 = v4
                delete(v1, 775)
                _ = fi17
            }
            v4 = v14
            _ = fi15
            _ = v16
        }
    } else {
        fmt.Println(v4)
    }
    v11 = append(v11, -355)
    v23 := GenS1{F0: 2147483647}
    v24 := GenS1{F0: -2147483648}
    if true {
        var v25 uint = 42
        if true {
            v11 = append(v11, 9223372036854775807)
        } else {
            v26 := GenS2{F0: 769, F1: true}
            _ = v26
        }
        if false {
            switch v14 {
            case "foo bar":
                v27 := GenS3{F0: v25}
                _ = v27
            }
        } else {
            var v28 float64 = -2.25
            _ = v28
        }
        _ = v25
    } else {
        v3 = -70
    }
    v29 := GenS2{F0: 336, F1: false}
    v14 = v4
    v30 := GenS1{F0: 2147483647}
    v4 = v14
    if len(v11) > 0 {
        v11[0] = 32767
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}

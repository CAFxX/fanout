package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int8; F1 uint }
func genFn0(p0 uint8, p1 uint64) uint {
    switch -246 {
    case 4294967295:
        v5 := GenS1{F0: (p0 * uint8(p1)), F1: p1}
        _ = v5
    case 233:
        var v6 int16 = 860
        _ = v6
    case -6:
        v7 := GenS2{F0: 42, F1: (1 != 490)}
        _ = v7
    default:
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 uint64 = (p1 - p1)
            _ = fi8
            _ = v9
        }
    }
    v10 := GenS2{F0: 582, F1: ((^(p1)) < (p1 - (p1 ^ (p1 % ((p1) | 1)))))}
    v11 := []int{-144}
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return uint(p0)
}

func main() {
    v1 := []string{}
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    var v4 uint32 = 42
    if true {
        if (v4 <= v4) {
            v12 := GenS2{F0: 4294967295, F1: (!false)}
            switch -9223372036854775808 {
            case -1:
                var v13 bool = true
                _ = v13
            default:
                delete(v2, 65535)
            }
            fmt.Println(v4)
            _ = v12
        } else {
            var v14 uint64 = 930
            v4 -= 1
            _ = v14
        }
        for fi15 := 0; fi15 < 5; fi15++ {
            v16 := GenS2{F0: 528, F1: ((-2.25 / float32(float64(1.5))) != 0.1)}
            var v17 uint64 = 887
            v18 := GenS2{F0: 814, F1: (42 < -1)}
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
    } else {
        switch 32767 {
        case -228:
            var v19 int32 = 2147483647
            _ = v19
        }
        v20 := GenS3{F0: 118, F1: 255}
        _ = v20
    }
    var v21 uint8 = 1
    v4 += v4
    v22 := GenS3{F0: 100, F1: genFn0(1, 4294967295)}
    switch "a" {
    case "hello":
        switch (((("\n\t" + "foo bar") + "a") + (("a" + ((("" + "foo bar") + "a") + "a")) + "hello")) + "") {
        case "":
            for fi23 := 0; fi23 < 15; fi23++ {
                if (false || false) {
                    v24 := v2[(v4 | (v4 / ((v4) | 1)))]
                    _ = v24
                } else {
                    switch ("a" + ("a" + "\n\t")) {
                    case "hello":
                        var v25 uint16 = 65535
                        _ = v25
                    default:
                        for fi26 := 0; fi26 < 2; fi26++ {
                            fmt.Println(v4, v4)
                            var v27 int16 = 0
                            v28 := GenS2{F0: 0, F1: false}
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                    }
                    var v29 uint8 = (v21 % ((174) | 1))
                    _ = v29
                }
                switch 0 {
                case 4294967295:
                    v4 += v4
                case -46:
                    var v30 string = "hello"
                    _ = v30
                }
                v21 += 94
                _ = fi23
            }
        }
    case "foo bar":
        switch "" {
        case "":
            var v31 bool = ((false && false) && (39 < int8(v4)))
            _ = v31
        case "hello":
            var v32 uint64 = 4294967295
            _ = v32
        }
    case "\n\t":
        v4 -= (v4 >> ((18446744073709551615) & 63))
    default:
        var v33 int16 = -94
        _ = v33
    }
    v4 += v4
    var v34 float32 = 1e10
    v35 := GenS1{F0: 127, F1: 1}
    v36 := GenS3{F0: 2, F1: 271}
    var v37 float64 = 1.5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}

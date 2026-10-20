package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }
func genFn0(p0 int64, p1 int32, p2 uint64) string {
    if (42 >= uint(p2)) {
        v5 := []int{42, 173, -363, 32767}
        p1 = (p1 ^ (-(p1)))
        _ = v5
    } else {
        if true {
            v6 := []int{32768, 0, 127, 32767}
            _ = v6
        } else {
            p1 = 42
            if false {
                var v7 int16 = -32768
                v7 = (v7 / ((v7) | 1))
                _ = v7
            } else {
                switch ("" + "hello") {
                case "\n\t":
                    fmt.Println(p2)
                case "":
                    if ((p1 ^ p1) <= p1) {
                        fmt.Println(p1)
                        v8 := GenS1{F0: ((^(p2)) > p2)}
                        _ = v8
                    } else {
                        fmt.Println(p1)
                        fmt.Println(p2)
                        fmt.Println(p1, p2)
                    }
                default:
                    var v9 int32 = int32(p0)
                    _ = v9
                }
                for fi10 := 0; fi10 < 1; fi10++ {
                    var v11 float32 = float32((1e10 * 1.5))
                    _ = fi10
                    _ = v11
                }
            }
            switch 0 {
            case -588:
                for fi12 := 0; fi12 < 13; fi12++ {
                    fmt.Println(p1)
                    v13 := []int{}
                    _ = fi12
                    _ = v13
                }
            case -9223372036854775808:
                var v14 int64 = (p0 >> ((517) & 63))
                _ = v14
            case 275:
                v15 := GenS1{F0: false}
                _ = v15
            default:
                p1 = 0
            }
        }
        fmt.Println(p1, p1)
    }
    v16 := []int{764, 889, 2147483647}
    switch int(p1) {
    case 32767:
        _ = len(v16)
    default:
        p2 += p2
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return ("a" + "a")
}
func genFn1(p0 bool, p1 int64, p2 uint16) uint32 {
    p0 = p0
    p1 = (p1 | (p1 / ((p1) | 1)))
    _ = p0
    _ = p1
    _ = p2
    return 671
}

func main() {
    var v1 int64 = 1
    var v2 uint = 1
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]bool)
    v17 := []int{9223372036854775807, 432}
    var v18 int8 = -1
    if len(v17) > 0 {
        v17[0] = 749
    }
    if (18446744073709551615 < 42) {
        var v19 string = ""
        for fi20 := 0; fi20 < 5; fi20++ {
            v1 = 1
            _ = fi20
        }
        switch -187 {
        case 42:
            v4[65535] = ((v2 | (v2 % (((^(v2))) | 1))) != v2)
        case -2147483648:
            var v21 int16 = -1
            _ = v21
        default:
            _ = len(v4)
        }
        _ = v19
    } else {
        v2 = v2
    }
    if true {
        var v22 uint8 = 206
        _ = v22
    }
    var v23 uint = v2
    for fi24 := 0; fi24 < 14; fi24++ {
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
        for fi25 := 0; fi25 < 15; fi25++ {
            switch (-(fi25)) {
            case -1:
                v26 := GenS2{F0: (v1 / (((v1 - 0)) | 1))}
                _ = v26
            case -814:
                var v27 int16 = 42
                _ = v27
            default:
                v28 := GenS1{F0: (((v1 >> (((v2 | v2)) & 63)) == -9223372036854775808) && true)}
                var v29 int8 = (-(v18))
                var v30 uint32 = 2147483648
                _ = v28
                _ = v29
                _ = v30
            }
            v2 += 1
            v1 -= (v1 ^ (v1 | v1))
            _ = fi25
        }
        _ = fi24
    }
    var v31 uint16 = 2
    switch int(v18) {
    case -9223372036854775808:
        switch "hello" {
        case "":
            v1 += -559
        default:
            v18 -= v18
        }
    case 32768:
        var v32 int = 9223372036854775807
        _ = v32
    case -1:
        v23 = 0
    default:
        for fi33 := 0; fi33 < 11; fi33++ {
            if true {
                var v34 uint64 = 1
                v35 := GenS2{F0: v1}
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
                _ = v34
                _ = v35
            }
            _ = fi33
        }
    }
    var v36 int32 = -1
    var v37 float64 = 1.5
    fmt.Println(v18)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v31
    _ = v36
    _ = v37
    _ = v4
}

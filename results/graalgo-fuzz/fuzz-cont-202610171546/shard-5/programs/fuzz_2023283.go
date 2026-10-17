package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int32 }
func genFn0(p0 uint16, p1 int) int16 {
    if ((42 <= 645) || true) {
        for fi6 := 0; fi6 < 6; fi6++ {
            if (1 != 708) {
                p1 -= -778
                if true {
                    var v7 uint64 = 1
                    var v8 uint8 = 193
                    v7 -= (v7 | (v7 / (((v7 + v7)) | 1)))
                    _ = v7
                    _ = v8
                }
            } else {
                v9 := GenS1{F0: (^(p0)), F1: (!true), F2: ((!(!((0.1 * ((0.1 + float64(((3.14159 - (float32((1.5 + 0.1)) - (1.5 * 0.0))) / 0.1))) / 1.0)) == ((float64(0.1) / 0.1) * -2.25)))) || true)}
                _ = v9
            }
            v10 := GenS2{F0: 0}
            _ = fi6
            _ = v10
        }
        for fi11 := 0; fi11 < 3; fi11++ {
            for fi12 := 0; fi12 < 18; fi12++ {
                if true {
                    v13 := GenS3{F0: -1}
                    fmt.Println(fi12, p1, p1)
                    _ = v13
                }
                var v14 uint32 = 2147483648
                var v15 int16 = 32767
                _ = fi12
                _ = v14
                _ = v15
            }
            var v16 uint64 = 4294967295
            v17 := []int{p1}
            _ = fi11
            _ = v16
            _ = v17
        }
        fmt.Println(p1, p1)
    }
    p1 -= p1
    _ = p0
    _ = p1
    return 32767
}
func genFn1(p0 uint8, p1 uint32, p2 int16) int32 {
    v18 := GenS4{F0: 42}
    for fi19 := 0; fi19 < 5; fi19++ {
        for fi20 := 0; fi20 < 1; fi20++ {
            fmt.Println(p2, p0, p2)
            var v21 uint = 2
            var v22 int32 = -2147483648
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi19
    }
    var v23 uint32 = (p1 % ((610) | 1))
    for fi24 := 0; fi24 < 9; fi24++ {
        v25 := GenS4{F0: -1}
        _ = fi24
        _ = v25
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v23
    return 491
}

func main() {
    var v1 bool = true
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]bool)
    var v4 uint32 = 65535
    v5 := []int{}
    v4 = 2
    v4 = (^(v4))
    v26 := v3[1]
    v1 = (0.0 < 0.1)
    v4 += v4
    if (129 >= 0) {
        var v27 float32 = (1.5 - 3.14159)
        if v1 {
            var v28 bool = (-1 <= 419)
            v29 := GenS3{F0: 116}
            _ = v28
            _ = v29
        } else {
            var v30 int8 = 42
            v26 = (967 == 2)
            if (0 >= 42) {
                delete(v2, 128)
            }
            _ = v30
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v27
    }
    v4 = v4
    for fi31 := 0; fi31 < 6; fi31++ {
        v32 := GenS2{F0: 4294967295}
        var v33 uint8 = 255
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 bool = v1
    v34 = (-1 != -1)
    v35 := GenS2{F0: 18446744073709551615}
    var v36 string = "hello"
    v37 := GenS4{F0: -1}
    var v38 int64 = 1
    for fi39 := 0; fi39 < 16; fi39++ {
        v34 = (false && true)
        v40 := GenS2{F0: 619}
        var v41 uint64 = 4294967295
        _ = fi39
        _ = v40
        _ = v41
    }
    var v42 uint16 = 1
    v4 -= 42
    for fi43 := 0; fi43 < 13; fi43++ {
        var v44 float32 = (3.14159 * (3.14159 * 1e10))
        switch v36 {
        case "":
            var v45 int16 = -1
            _ = v45
        case "foo bar":
            var v46 uint16 = v42
            _ = v46
        case "a":
            if (0 == 2147483647) {
                v1 = (42 <= 42)
            } else {
                for fi47 := 0; fi47 < 4; fi47++ {
                    v48 := GenS5{F0: -761}
                    _ = fi47
                    _ = v48
                }
            }
        default:
            v49 := GenS6{F0: 2147483647}
            _ = v49
        }
        _ = fi43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v42
    _ = v5
}

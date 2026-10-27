package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int16; F1 int32; F2 bool }
type GenS3 struct { F0 int32 }
func genFn0(p0 string, p1 int) string {
    var v7 float32 = 0.0
    v8 := []int{p1}
    v9 := GenS1{F0: -1, F1: 2}
    p1 = (-(p1))
    p1 -= p1
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return p0
}
func genFn1() string {
    if true {
        var v10 uint8 = 42
        if false {
            fmt.Println(v10)
        } else {
            v11 := GenS1{F0: -9223372036854775808, F1: 2}
            v10 += 53
            if true {
                v12 := []int{408}
                for fi13 := 0; fi13 < 5; fi13++ {
                    if false {
                        if true {
                            v14 := GenS1{F0: 1, F1: 255}
                            var v15 bool = (((1e10 - -2.25) != float64(1e10)) || (!(!false)))
                            _ = len(v12)
                            _ = v14
                            _ = v15
                        }
                        v10 = 127
                        v16 := GenS1{F0: -1, F1: 104}
                        _ = v16
                    } else {
                        var v17 int64 = 42
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                        _ = v17
                    }
                    v10 -= (v10 << ((198) & 63))
                    fmt.Println(v10, fi13)
                    _ = fi13
                }
                _ = v12
            }
            _ = v11
        }
        v18 := []int{-656}
        _ = v10
        _ = v18
    }
    var v19 bool = (!(!(!(((1.5 + (3.14159 / -2.25)) != ((1e10 / ((0.0 - float32(3.14159)) / -2.25)) * 3.14159)) || true))))
    for fi20 := 0; fi20 < 9; fi20++ {
        var v21 bool = v19
        _ = fi20
        _ = v21
    }
    _ = v19
    return ("\n\t" + ("foo bar" + (("a" + "foo bar") + (("\n\t" + "") + ""))))
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []uint32{}
    var v3 uint16 = 815
    v4 := make(map[int]int)
    var v5 int64 = (v1 ^ -1)
    var v6 int8 = -40
    if false {
        var v22 int32 = -1
        v5 += 42
        v22 += v22
        _ = v22
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        switch "" {
        case "":
            var v24 uint16 = v3
            _ = v24
        case "a":
            v25 := GenS1{F0: (v1 / ((v5) | 1)), F1: 418}
            _ = v25
        }
        var v26 uint8 = 128
        var v27 bool = (true || (42 == 2147483647))
        _ = fi23
        _ = v26
        _ = v27
    }
    switch "foo bar" {
    case "foo bar":
        var v28 int = 2147483647
        _ = v28
    case "hello":
        v1 -= 1
    }
    v29 := GenS2{F0: 1, F1: 42, F2: (1 == -32768)}
    var v30 uint = 4294967295
    v1 -= v5
    for fi31 := 0; fi31 < 14; fi31++ {
        var v32 int32 = -2147483648
        _ = fi31
        _ = v32
    }
    var v33 uint = v30
    v34 := GenS1{F0: (v5 + (v5 & -1)), F1: 255}
    var v35 uint64 = 2
    var v36 bool = false
    fmt.Println(v36)
    if (uint64(v33) == (v35 / ((18446744073709551615) | 1))) {
        var v37 int8 = (v6 | v6)
        fmt.Println(v6, v5, v36)
        if v36 {
            var v38 int16 = 32767
            if false {
                v39 := GenS2{F0: (v38 - v38), F1: 1, F2: v36}
                fmt.Println(v37, v35, v36)
                _ = v39
            } else {
                v40 := GenS1{F0: (v1 | 1), F1: (v33 | v33)}
                for fi41 := 0; fi41 < 19; fi41++ {
                    _ = len(v4)
                    _ = len(v4)
                    v33 -= v33
                    _ = fi41
                }
                _ = v40
            }
            _ = v38
        } else {
            v42 := GenS3{F0: -486}
            fmt.Println(v30, v36)
            _ = v42
        }
        _ = v37
    }
    for fi43 := 0; fi43 < 9; fi43++ {
        if true {
            v5 = (-(v1))
        } else {
            var v44 float32 = 1e10
            v2 = append(v2, 903)
            var v45 string = ("\n\t" + "a")
            _ = v44
            _ = v45
        }
        for fi46 := 0; fi46 < 14; fi46++ {
            v35 -= (v35 % ((42) | 1))
            fmt.Println(v1)
            _ = fi46
        }
        _ = fi43
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}

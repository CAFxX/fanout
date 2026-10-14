package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int64, p1 int32) int {
    var v7 uint8 = 125
    var v8 string = (("a" + "a") + "a")
    v9 := GenS1{F0: -9223372036854775808}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 2147483647
}
func genFn1() uint32 {
    if false {
        switch -1 {
        case -1:
            var v10 int32 = 0
            _ = v10
        case 128:
            for fi11 := 0; fi11 < 8; fi11++ {
                if true {
                    var v12 float32 = (1.5 + (((0.0 / -2.25) - (float32((float64(1e10) * -2.25)) - ((float32(1.5) - -2.25) * 1.5))) * 0.0))
                    if false {
                        v12 -= (((((v12 * v12) * 1.5) + ((3.14159 + ((0.1 / (1.5 / v12)) * 1e10)) - v12)) - 0.0) - ((-2.25 - v12) + 1e10))
                        for fi13 := 0; fi13 < 9; fi13++ {
                            v14 := []int{}
                            _ = fi13
                            _ = v14
                        }
                        var v15 uint = 925
                        _ = v15
                    }
                    _ = v12
                }
                var v16 bool = (0 >= 0)
                _ = fi11
                _ = v16
            }
        case 862:
            var v17 uint16 = 2
            _ = v17
        }
        var v18 string = "hello"
        v19 := []int{32768, 9223372036854775807, 9223372036854775807, genFn0(42, -651)}
        _ = v18
        _ = v19
    } else {
        var v20 float64 = 0.0
        _ = v20
    }
    v21 := []int{4294967295, 591}
    if false {
        var v22 string = "a"
        _ = v22
    } else {
        var v23 float64 = 0.1
        _ = v23
    }
    switch -99 {
    case -1:
        var v24 int8 = 100
        _ = v24
    default:
        var v25 uint64 = 438
        _ = v25
    }
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    _ = v21
    return 1
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[string]int)
    var v3 uint32 = 2147483647
    var v4 int8 = 42
    var v5 float64 = 1.5
    var v6 int64 = 286
    v26 := GenS2{F0: -2147483648}
    var v27 uint = 255
    v28 := GenS3{F0: 2147483647}
    switch 32767 {
    case 127:
        for fi29 := 0; fi29 < 17; fi29++ {
            v30 := GenS2{F0: 0}
            _ = fi29
            _ = v30
        }
    case 2147483647:
        v3 = (v3 << ((v27) & 63))
    }
    var v31 int64 = v6
    var v32 uint64 = 18446744073709551615
    if ((v27 * (v27 >> ((v27) & 63))) >= 255) {
        v32 = (v32 << (((v27 << ((255) & 63))) & 63))
        v33 := GenS2{F0: 1}
        v4 += -59
        _ = v33
    } else {
        v6 -= 0
    }
    if (9223372036854775807 < 1) {
        v32 = 18446744073709551615
        switch 898 {
        case -382:
            v34 := []int{-9223372036854775808, genFn0(v31, 833), genFn0((v6 % (((v31 | (v6 - v6))) | 1)), -1), -1}
            _ = v34
        }
        for fi35 := 0; fi35 < 16; fi35++ {
            var v36 uint64 = v32
            fmt.Println(v4)
            if true {
                var v37 string = v1
                for fi38 := 0; fi38 < 17; fi38++ {
                    v6 = v31
                    _ = fi38
                }
                if (1 > 42) {
                    v4 += -52
                } else {
                    var v39 uint8 = 204
                    for fi40 := 0; fi40 < 3; fi40++ {
                        v41 := GenS1{F0: fi35}
                        _ = fi40
                        _ = v41
                    }
                    _ = v39
                }
                _ = v37
            }
            _ = fi35
            _ = v36
        }
    } else {
        v6 = v6
        for fi42 := 0; fi42 < 9; fi42++ {
            fmt.Println(v3, v4, v6)
            fmt.Println(v31, v27)
            fmt.Println(v1, v5)
            _ = fi42
        }
    }
    switch v1 {
    case "foo bar":
        var v43 int32 = 0
        _ = v43
    case "a":
        delete(v2, v1)
    default:
        var v44 uint = 121
        _ = v44
    }
    var v45 int32 = 930
    for fi46 := 0; fi46 < 8; fi46++ {
        v47 := []int{}
        fmt.Println(v5, v31, v3)
        v45 = v45
        _ = fi46
        _ = v47
    }
    v45 -= (v45 | -2147483648)
    var v48 int8 = v4
    switch v1 {
    case "hello":
        v32 -= v32
    case "a":
        var v49 int16 = -475
        _ = v49
    default:
        v48 = (v48 << (((^(v27))) & 63))
    }
    switch -788 {
    case 679:
        for fi50 := 0; fi50 < 15; fi50++ {
            v51 := GenS2{F0: v45}
            _ = fi50
            _ = v51
        }
    case 1:
        v27 = (^(v27))
    default:
        v3 = v3
    }
    v45 = 2147483647
    v2[("" + "\n\t")] = 9223372036854775807
    v6 -= int64(v45)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v45
    _ = v48
    _ = v5
    _ = v6
}

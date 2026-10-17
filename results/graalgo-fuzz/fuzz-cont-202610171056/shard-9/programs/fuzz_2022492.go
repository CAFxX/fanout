package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int32 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 int8; F2 uint16 }
func genFn0(p0 string) uint32 {
    p0 = p0
    for fi4 := 0; fi4 < 3; fi4++ {
        for fi5 := 0; fi5 < 15; fi5++ {
            var v6 uint = 436
            p0 = "foo bar"
            _ = fi5
            _ = v6
        }
        var v7 string = p0
        _ = fi4
        _ = v7
    }
    var v8 float32 = float32((float64(float32(1e10)) / 1.0))
    switch 128 {
    case -1:
        v8 = 1e10
    }
    _ = p0
    _ = v8
    return 65535
}

func main() {
    var v1 bool = (42 <= 0)
    var v2 int32 = -590
    var v3 int8 = -106
    v2 = (-(v2))
    v9 := GenS1{F0: 4294967295, F1: "a", F2: v2}
    switch "\n\t" {
    case "\n\t":
        v10 := GenS1{F0: 235, F1: "\n\t", F2: v2}
        _ = v10
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        v2 -= (v2 ^ (v2 - 0))
        _ = fi11
    }
    var v12 bool = (v1 || (v1 && v1))
    switch "foo bar" {
    case "foo bar":
        if v12 {
            var v13 int32 = (-(v2))
            v2 -= 0
            var v14 float32 = (1e10 - 1e10)
            _ = v13
            _ = v14
        } else {
            var v15 float64 = 3.14159
            _ = v15
        }
    case "\n\t":
        var v16 float32 = 1.5
        _ = v16
    case "":
        if v12 {
            fmt.Println(v3, v2, v3)
            v2 -= (^(v2))
        } else {
            switch 127 {
            case -2147483648:
                var v17 uint16 = 723
                _ = v17
            case -105:
                v18 := []int{-803}
                _ = v18
            default:
                v3 = (v3 - v3)
            }
            switch "\n\t" {
            case "\n\t":
                for fi19 := 0; fi19 < 20; fi19++ {
                    if false {
                        v20 := []int{(^(fi19)), fi19, -9223372036854775808, (fi19 / (((fi19 ^ fi19)) | 1))}
                        _ = len(v20)
                        switch 42 {
                        case 127:
                            for fi21 := 0; fi21 < 11; fi21++ {
                                v22 := GenS2{F0: (fi21 ^ fi19), F1: 32767}
                                v3 -= v3
                                v1 = v1
                                _ = fi21
                                _ = v22
                            }
                        case 1:
                            v3 += -1
                        }
                        _ = v20
                    } else {
                        for fi23 := 0; fi23 < 10; fi23++ {
                            v24 := GenS2{F0: fi23, F1: int16(v2)}
                            var v25 int16 = 0
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                    }
                    _ = fi19
                }
            }
            if (4294967295 >= -285) {
                if (355 == 9223372036854775807) {
                    v1 = false
                } else {
                    switch 42 {
                    case 32768:
                        v1 = v12
                    case -2147483648:
                        v26 := []int{9223372036854775807}
                        _ = v26
                    default:
                        fmt.Println(v1)
                    }
                    var v27 uint32 = 2147483648
                    _ = v27
                }
                for fi28 := 0; fi28 < 9; fi28++ {
                    v29 := GenS2{F0: 105, F1: -111}
                    v30 := []int{fi28, (fi28 + (-(fi28))), fi28}
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                for fi31 := 0; fi31 < 7; fi31++ {
                    v3 = (v3 / ((1) | 1))
                    fmt.Println(v3, fi31)
                    _ = fi31
                }
            }
        }
    default:
        if (1 > 65535) {
            fmt.Println(v2, v1)
            for fi32 := 0; fi32 < 18; fi32++ {
                var v33 bool = (-2.25 > 1e10)
                _ = fi32
                _ = v33
            }
            v34 := []int{877, 32768}
            _ = v34
        } else {
            var v35 float32 = 3.14159
            var v36 uint8 = 20
            _ = v35
            _ = v36
        }
    }
    if false {
        v37 := []int{32767, 1, 770}
        _ = v37
    }
    v12 = v1
    switch 42 {
    case 810:
        v38 := []int{32768}
        _ = v38
    case 1:
        v12 = false
    }
    v1 = (v1 || (0.1 < 1.5))
    if (((!v12) && (true && v1)) && v12) {
        if v12 {
            v1 = (v1 || v12)
            fmt.Println(v1, v1, v12)
            for fi39 := 0; fi39 < 7; fi39++ {
                v1 = (!true)
                v40 := GenS3{F0: 115}
                _ = fi39
                _ = v40
            }
        }
        v41 := []int{32767}
        _ = v41
    }
    v2 -= v2
    v42 := []int{975, 127, 1}
    _ = len(v42)
    v12 = v1
    var v43 uint16 = 531
    switch "foo bar" {
    case "":
        v44 := GenS4{F0: 100, F1: 127, F2: v43}
        _ = v44
    case "foo bar":
        if len(v42) >= 2 {
            v42 = v42[1:len(v42)]
        }
    }
    var v45 int8 = (v3 | (v3 / (((v3 | 83)) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v42
    _ = v43
    _ = v45
    _ = v9
}

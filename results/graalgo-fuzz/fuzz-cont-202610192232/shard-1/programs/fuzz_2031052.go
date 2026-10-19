package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 string }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 int; F1 int }
func genFn0(p0 uint16, p1 uint) uint32 {
    v6 := GenS1{F0: 0, F1: 4294967295, F2: "\n\t"}
    switch "\n\t" {
    case "a":
        v7 := []int{-311}
        _ = v7
    }
    fmt.Println(p0)
    fmt.Println(p0, p0)
    for fi8 := 0; fi8 < 12; fi8++ {
        fmt.Println(p1, p0, fi8)
        fmt.Println(p1, fi8)
        fmt.Println(fi8, p0, p1)
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v6
    return 42
}
func genFn1() bool {
    v9 := []int{128, 128, 4294967295, -2147483648}
    for fi10 := 0; fi10 < 6; fi10++ {
        if len(v9) > 0 {
            v9[0] = (fi10 << ((785) & 63))
        }
        _ = fi10
    }
    var v11 uint = 18446744073709551615
    var v12 bool = (!((32767 >= 222) && ((!(1000 != 0)) && (2147483648 < 0))))
    _ = v11
    _ = v12
    _ = v9
    return v12
}

func main() {
    var v1 uint16 = 42
    var v2 int32 = -2147483648
    v3 := []uint32{1, 99, 968}
    v4 := []int{2147483647, 32768}
    var v5 int8 = 100
    v2 += v2
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 int32 = (v2 << ((uint(v1)) & 63))
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        v5 += 54
        _ = fi15
    }
    v1 = (v1 % (((v1 - 2)) | 1))
    var v16 int8 = v5
    for fi17 := 0; fi17 < 19; fi17++ {
        if false {
            fmt.Println(v1, v5, v16)
            var v18 int64 = 1
            v19 := GenS2{F0: v16, F1: 1}
            _ = v18
            _ = v19
        } else {
            var v20 uint8 = 42
            _ = v20
        }
        if false {
            v1 = v1
        } else {
            v16 = v5
        }
        _ = fi17
    }
    if len(v4) > 0 {
        v4[0] = 604
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        for fi22 := 0; fi22 < 8; fi22++ {
            if true {
                var v23 uint16 = v1
                fmt.Println(v23)
                v2 -= (v2 << ((42) & 63))
                if false {
                    fmt.Println(fi22)
                }
                _ = v23
            } else {
                var v24 int = (-(fi21))
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if false {
                    v1 += 905
                    var v25 int8 = v5
                    if true {
                        v5 = (-(v5))
                        if len(v4) > 0 {
                            v4[0] = v24
                        }
                    } else {
                        fmt.Println(v25, v2, v1)
                    }
                    _ = v25
                } else {
                    var v26 float32 = float32(3.14159)
                    _ = v26
                }
                _ = v24
            }
            _ = fi22
        }
        _ = fi21
    }
    var v27 uint = 2
    switch -820 {
    case 671:
        v28 := GenS3{F0: -976, F1: 749}
        _ = v28
    case 42:
        v29 := GenS1{F0: 42, F1: 984, F2: "hello"}
        _ = v29
    case 127:
        if true {
            switch ("" + "foo bar") {
            case "":
                v5 = (v16 & -1)
            }
            v27 += (v27 - v27)
            if false {
                var v30 uint = v27
                _ = v30
            }
        } else {
            fmt.Println(v1, v2, v16)
            fmt.Println(v16, v27)
        }
    }
    for fi31 := 0; fi31 < 8; fi31++ {
        for fi32 := 0; fi32 < 18; fi32++ {
            if true {
                v33 := GenS3{F0: (fi32 << ((v27) & 63)), F1: int(v1)}
                _ = len(v3)
                fmt.Println(v1, v2, v2)
                _ = v33
            } else {
                switch fi32 {
                case 42:
                    var v34 int16 = 32767
                    _ = v34
                case 127:
                    switch ("hello" + "hello") {
                    case "":
                        var v35 float64 = (((1e10 * 0.0) * 1e10) - 1e10)
                        _ = v35
                    }
                }
                v36 := GenS2{F0: (v16 << ((v27) & 63)), F1: 232}
                switch "a" {
                case "a":
                    fmt.Println(v16)
                case "\n\t":
                    var v37 uint64 = 2
                    _ = v37
                case "hello":
                    v16 += (-(v5))
                default:
                    v38 := GenS2{F0: v16, F1: 18446744073709551615}
                    _ = v38
                }
                _ = v36
            }
            v2 = (v2 << ((v27) & 63))
            _ = fi32
        }
        _ = fi31
    }
    v1 = 42
    var v39 int = -1
    v2 += v2
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 int32; F2 uint8 }
type GenS4 struct { F0 int16; F1 int32; F2 uint }
type GenS5 struct { F0 int64; F1 bool }
func genFn0(p0 int8, p1 uint8) uint8 {
    var v7 float32 = 3.14159
    v7 = 1e10
    var v8 string = "a"
    v9 := []int{32767, 296}
    _ = len(v9)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 10
}
func genFn1(p0 bool) uint32 {
    if p0 {
        var v10 uint = 255
        var v11 float32 = 1.5
        _ = v10
        _ = v11
    } else {
        v12 := GenS1{F0: 32767}
        var v13 uint8 = 255
        v14 := []int{}
        _ = v12
        _ = v13
        _ = v14
    }
    var v15 uint16 = 32767
    if (!(!p0)) {
        if p0 {
            for fi16 := 0; fi16 < 9; fi16++ {
                p0 = false
                for fi17 := 0; fi17 < 19; fi17++ {
                    v15 -= v15
                    fmt.Println(fi17)
                    _ = fi17
                }
                fmt.Println(v15, fi16, p0)
                _ = fi16
            }
            for fi18 := 0; fi18 < 18; fi18++ {
                if false {
                    var v19 string = "hello"
                    _ = v19
                } else {
                    var v20 uint32 = 100
                    var v21 bool = false
                    _ = v20
                    _ = v21
                }
                for fi22 := 0; fi22 < 9; fi22++ {
                    v15 = (v15 << ((42) & 63))
                    v15 = (v15 << ((uint(v15)) & 63))
                    _ = fi22
                }
                _ = fi18
            }
            for fi23 := 0; fi23 < 1; fi23++ {
                v24 := GenS2{F0: 42}
                v25 := GenS1{F0: -938}
                _ = fi23
                _ = v24
                _ = v25
            }
        } else {
            var v26 int32 = 2147483647
            v27 := GenS1{F0: 1000}
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 10; fi28++ {
            if false {
                for fi29 := 0; fi29 < 15; fi29++ {
                    var v30 int8 = -90
                    p0 = p0
                    v31 := GenS2{F0: 2147483647}
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                var v32 uint = 176
                _ = v32
            }
            for fi33 := 0; fi33 < 5; fi33++ {
                v34 := []int{fi33, (fi33 / (((^(fi33))) | 1))}
                var v35 string = "\n\t"
                _ = fi33
                _ = v34
                _ = v35
            }
            _ = fi28
        }
    }
    v36 := []int{-2147483648, -2147483648, 1, 2147483647}
    p0 = p0
    _ = p0
    _ = v15
    _ = v36
    return 2
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 65535
    v3 := make(map[int]uint32)
    var v4 int16 = 723
    var v5 uint8 = 151
    v6 := []string{}
    v4 = v4
    switch -957 {
    case 4294967295:
        if (0.0 == (1e10 + ((3.14159 * float32(((1e10 / (1e10 - 1.5)) * -2.25))) * 1.5))) {
            if true {
                v37 := GenS2{F0: -461}
                _ = v37
            } else {
                switch "" {
                case "hello":
                    if (42 <= 155) {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                case "\n\t":
                    var v38 int32 = 0
                    _ = v38
                case "":
                    v39 := GenS3{F0: "a", F1: -2147483648, F2: (v5 ^ genFn0(-125, (v5 >> ((uint(v2)) & 63))))}
                    _ = v39
                }
                fmt.Println(v5)
                switch 127 {
                case 556:
                    v40 := GenS1{F0: 700}
                    _ = v40
                }
            }
        } else {
            v6 = append(v6, "hello")
            fmt.Println(v4)
        }
    case -680:
        fmt.Println(v5, v5, v2)
    case -736:
        delete(v1, 2147483648)
    default:
        var v41 int64 = -1
        _ = v41
    }
    fmt.Println(v5, v5)
    var v42 int32 = 1
    var v43 float32 = -2.25
    v2 += genFn1(true)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi44 := 0; fi44 < 10; fi44++ {
        if false {
            fmt.Println(v2, v43, v2)
            v45 := GenS4{F0: v4, F1: int32(v5), F2: 0}
            _ = v45
        }
        v46 := GenS5{F0: -137, F1: (v42 == (v42 - (^(v42))))}
        _ = fi44
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}

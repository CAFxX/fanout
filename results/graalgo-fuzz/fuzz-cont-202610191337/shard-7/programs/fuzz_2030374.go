package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 int16; F1 int8; F2 uint32 }
func genFn0() int64 {
    var v7 uint16 = 65535
    if true {
        v8 := []int{-668}
        var v9 int = 32767
        _ = v8
        _ = v9
    }
    _ = v7
    return 0
}
func genFn1() int16 {
    switch "hello" {
    case "":
        var v10 int16 = -706
        _ = v10
    default:
        var v11 float32 = (3.14159 + ((1.5 - 0.1) - (1e10 / -2.25)))
        _ = v11
    }
    v12 := GenS1{F0: 0}
    var v13 uint = 1
    var v14 int8 = -9
    _ = v12
    _ = v13
    _ = v14
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{}
    v3 := make(map[string]int)
    var v4 string = "a"
    v5 := []uint32{}
    v6 := make(map[string]int)
    var v15 int64 = 1
    switch 128 {
    case 128:
        fmt.Println(v4)
    case -920:
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 string = (v4 + "a")
            delete(v3, (("" + v17) + ("a" + "a")))
            _ = fi16
            _ = v17
        }
    case 157:
        v4 = v4
    default:
        if true {
            for fi18 := 0; fi18 < 8; fi18++ {
                v15 = v15
                if true {
                    for fi19 := 0; fi19 < 13; fi19++ {
                        _ = len(v3)
                        v20 := GenS1{F0: 653}
                        var v21 uint64 = 1
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 6; fi22++ {
                        switch v4 {
                        case "a":
                            v23 := GenS2{F0: 995, F1: 557}
                            _ = v23
                        default:
                            v5 = append(v5, 1)
                        }
                        _ = fi22
                    }
                    fmt.Println(v15)
                } else {
                    var v24 bool = ((!false) || false)
                    v25 := GenS3{F0: 4294967295, F1: (v4 + "hello")}
                    _ = v24
                    _ = v25
                }
                v4 = v4
                _ = fi18
            }
            fmt.Println(v15, v15)
        } else {
            fmt.Println(v15, v15)
        }
    }
    v26 := GenS2{F0: 47, F1: 0}
    if ((3.14159 + 1e10) == 0.0) {
        for fi27 := 0; fi27 < 17; fi27++ {
            for fi28 := 0; fi28 < 5; fi28++ {
                v4 = "a"
                v2 = append(v2, (v4 + v4))
                var v29 int64 = v15
                _ = fi28
                _ = v29
            }
            var v30 uint = 2
            v31 := v6[("\n\t" + "foo bar")]
            _ = fi27
            _ = v30
            _ = v31
        }
    } else {
        if (2 == 65535) {
            delete(v3, v4)
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        v4 = (v4 + "a")
        fmt.Println(v15)
    }
    switch 1 {
    case -2147483648:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case -674:
        var v32 int16 = 0
        _ = v32
    case 1:
        var v33 int = -477
        _ = v33
    }
    for fi34 := 0; fi34 < 1; fi34++ {
        for fi35 := 0; fi35 < 9; fi35++ {
            var v36 int32 = -275
            var v37 string = (v4 + ("foo bar" + (("\n\t" + (((((v4 + ((v4 + "foo bar") + v4)) + "foo bar") + "a") + (v4 + v4)) + (v4 + "hello"))) + v4)))
            _ = len(v5)
            var v38 uint8 = 128
            v39 := GenS4{F0: -128, F1: true}
            fmt.Println(v37)
            v36 += (v36 >> ((65535) & 63))
            _ = fi35
            _ = v36
            _ = v37
            _ = v38
            _ = v39
        }
        var v40 uint = 2
        v3[(v4 + (v4 + "hello"))] = fi34
        _ = fi34
        _ = v40
    }
    var v41 uint16 = 2
    if (!false) {
        var v42 uint8 = 255
        _ = v42
    }
    v43 := GenS4{F0: 127, F1: (0.1 == (1e10 / -2.25))}
    if (-987 <= 469) {
        var v44 uint64 = 2
        _ = v44
    }
    v41 = (v41 ^ (v41 * (v41 * (v41 - (^(v41))))))
    if len(v5) > 0 {
        v5[0] = 1
    }
    var v45 float64 = (-2.25 * 1e10)
    switch -1 {
    case 32768:
        for fi46 := 0; fi46 < 10; fi46++ {
            delete(v6, v4)
            switch "a" {
            case "hello":
                v47 := GenS5{F0: 32767, F1: 78, F2: 664}
                _ = v47
            }
            v48 := GenS1{F0: 2}
            _ = fi46
            _ = v48
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v41
    _ = v43
    _ = v45
    _ = v5
    _ = v6
}

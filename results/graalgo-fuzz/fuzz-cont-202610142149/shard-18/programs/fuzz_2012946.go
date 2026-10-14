package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 bool }
type GenS2 struct { F0 string; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 int64; F1 int; F2 uint64 }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64; F1 uint16; F2 int64 }
func genFn0(p0 string, p1 uint32) int32 {
    v7 := []int{-1, 227, 1}
    if true {
        switch 127 {
        case 32768:
            v8 := GenS1{F0: 2, F1: 0, F2: true}
            _ = v8
        case -578:
            for fi9 := 0; fi9 < 12; fi9++ {
                for fi10 := 0; fi10 < 1; fi10++ {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    p0 = p0
                    var v11 int32 = 1
                    _ = fi10
                    _ = v11
                }
                _ = fi9
            }
        case 604:
            if len(v7) > 0 {
                v7[0] = 128
            }
        default:
            for fi12 := 0; fi12 < 12; fi12++ {
                for fi13 := 0; fi13 < 7; fi13++ {
                    var v14 int64 = 1
                    var v15 uint64 = 42
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                _ = fi12
            }
        }
        var v16 bool = (((true && (((true || (1 > 42)) || (!false)) && true)) || (((70 == 1) || (uint8(p1) > 146)) || (true || ((p0 != p0) || (436 != (p1 * (p1 & 2147483648))))))) || (9223372036854775807 == -164))
        v7 = append(v7, 774)
        _ = v16
    } else {
        fmt.Println(p1, p1)
        fmt.Println(p1, p1)
    }
    var v17 string = "hello"
    v7 = append(v7, 128)
    v18 := GenS2{F0: ((p0 + v17) + "hello"), F1: 255, F2: 0}
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    _ = v7
    return 42
}

func main() {
    var v1 bool = true
    v2 := []uint32{748, 65535, 67, 2}
    v3 := make(map[int]uint32)
    var v4 int16 = 32767
    var v5 float32 = 3.14159
    var v6 float64 = 0.0
    fmt.Println(v4)
    for fi19 := 0; fi19 < 11; fi19++ {
        var v20 uint64 = 1
        switch "\n\t" {
        case "hello":
            fmt.Println(v1)
        case "foo bar":
            for fi21 := 0; fi21 < 2; fi21++ {
                v6 += 3.14159
                v22 := GenS3{F0: 42, F1: 0}
                _ = fi21
                _ = v22
            }
        }
        _ = fi19
        _ = v20
    }
    v1 = (445 <= 1)
    v6 = (v6 + (float64(v5) - v6))
    for fi23 := 0; fi23 < 12; fi23++ {
        switch fi23 {
        case 641:
            switch 736 {
            case 1:
                if (2 >= 2) {
                    for fi24 := 0; fi24 < 2; fi24++ {
                        var v25 uint8 = 121
                        _ = fi24
                        _ = v25
                    }
                    switch fi23 {
                    case -128:
                        if v1 {
                            v26 := GenS4{F0: -746, F1: fi23, F2: 2}
                            _ = v26
                        }
                    case 32767:
                        v4 += 42
                    }
                } else {
                    if v1 {
                        var v27 int8 = 127
                        v28 := GenS5{F0: true, F1: 1}
                        v6 -= (0.0 + float64((3.14159 / float32(((v6 + (v6 - ((v6 - (float64((v5 - v5)) / 1.0)) + 1.5))) - -2.25)))))
                        _ = v27
                        _ = v28
                    }
                    v1 = (!(fi23 <= fi23))
                }
            case -1:
                var v29 float32 = (v5 + v5)
                _ = v29
            case 9223372036854775807:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        default:
            _ = len(v2)
        }
        fmt.Println(v1)
        v5 += 1e10
        _ = fi23
    }
    var v30 uint = 42
    v4 += 42
    if v1 {
        fmt.Println(v6, v5)
        v31 := GenS6{F0: 2}
        _ = v31
    }
    v32 := GenS2{F0: ("\n\t" + ((("foo bar" + "") + "\n\t") + "")), F1: 647, F2: 0}
    v2 = append(v2, 2)
    v33 := GenS7{F0: 462, F1: 0, F2: -625}
    v30 = (v30 ^ (v30 & (v30 >> (((v30 + (^(v30)))) & 63))))
    if v1 {
        var v34 uint32 = 2
        _ = v34
    }
    _ = len(v2)
    fmt.Println(v6)
    if (float64((v5 + 3.14159)) < v6) {
        for fi35 := 0; fi35 < 8; fi35++ {
            for fi36 := 0; fi36 < 14; fi36++ {
                v6 += v6
                switch ("" + "hello") {
                case "hello":
                    v2 = append(v2, 411)
                default:
                    switch "hello" {
                    case "":
                        switch "" {
                        case "":
                            v1 = true
                        }
                    }
                }
                _ = fi36
            }
            _ = fi35
        }
    }
    var v37 float64 = v6
    var v38 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string; F1 uint16 }
type GenS6 struct { F0 int; F1 uint8; F2 int }

func main() {
    var v1 uint32 = 2147483647
    var v2 int8 = 100
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]string)
    var v5 int8 = 100
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 int8 = v5
        var v8 uint = 42
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        for fi10 := 0; fi10 < 7; fi10++ {
            v1 = v1
            delete(v3, fi10)
            _ = fi10
        }
        v5 += (-(v5))
        for fi11 := 0; fi11 < 6; fi11++ {
            if false {
                var v12 int8 = (^(v5))
                var v13 uint16 = 0
                _ = v12
                _ = v13
            }
            switch fi11 {
            case 4294967295:
                var v14 bool = (13 != 811)
                _ = v14
            case 0:
                v2 -= (v5 & (v2 | v2))
            case 323:
                var v15 uint64 = 4294967295
                v5 -= v2
                _ = v15
            }
            if false {
                switch "hello" {
                case "a":
                    if (!true) {
                        var v16 uint32 = 2147483647
                        if true {
                            if (-1 == 42) {
                                v17 := []int{}
                                switch "a" {
                                case "foo bar":
                                    v2 = (v5 & v5)
                                default:
                                    fmt.Println(v1, v1)
                                }
                                _ = v17
                            } else {
                                v1 -= v16
                            }
                            v18 := []int{}
                            v19 := GenS1{F0: "\n\t", F1: 0}
                            _ = v18
                            _ = v19
                        } else {
                            fmt.Println(fi11)
                        }
                        v2 -= 127
                        v1 = v1
                        fmt.Println(v16, v5, v16)
                        _ = v16
                    } else {
                        v20 := GenS1{F0: "\n\t", F1: 345}
                        v2 += 1
                        _ = v20
                    }
                case "hello":
                    v5 -= -93
                }
            } else {
                v1 -= v1
                v1 = v1
                _ = len(v4)
            }
            _ = fi11
        }
        _ = fi9
    }
    var v21 int32 = 42
    v22 := GenS2{F0: -1, F1: 734, F2: 2}
    var v23 uint64 = 18446744073709551615
    v24 := GenS3{F0: -9223372036854775808}
    fmt.Println(v23, v2, v1)
    switch "hello" {
    case "":
        v25 := v3[42]
        _ = v25
    case "foo bar":
        switch -1 {
        case 127:
            var v26 float32 = (1e10 + (float32((3.14159 - float64(float32(0.0)))) + 1e10))
            _ = v26
        case -429:
            v5 = (-(v5))
        default:
            v27 := GenS2{F0: 1, F1: 957, F2: v23}
            _ = v27
        }
    default:
        for fi28 := 0; fi28 < 13; fi28++ {
            for fi29 := 0; fi29 < 5; fi29++ {
                v30 := GenS4{F0: -32768}
                var v31 uint8 = 240
                var v32 uint8 = v31
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
            }
            for fi33 := 0; fi33 < 16; fi33++ {
                v34 := GenS5{F0: ("a" + ("" + "foo bar")), F1: 426}
                var v35 uint8 = 127
                _ = fi33
                _ = v34
                _ = v35
            }
            _ = fi28
        }
    }
    if (v5 <= -128) {
        if true {
            if ((((-2.25 / float32(float64((3.14159 / ((1e10 / 1.0) - -2.25))))) * 1.5) * float32(float64(0.0))) != ((-2.25 + 1.5) + 1e10)) {
                var v36 int32 = (v21 ^ v21)
                if true {
                    v37 := []int{1, 998, 9223372036854775807, 128}
                    for fi38 := 0; fi38 < 6; fi38++ {
                        var v39 uint = 818
                        var v40 uint = (v39 + v39)
                        if false {
                            v41 := GenS6{F0: fi38, F1: 2, F2: (fi38 % ((fi38) | 1))}
                            if false {
                                v21 += v36
                            } else {
                                fmt.Println(v40)
                                v36 -= v21
                                var v42 uint = (v39 / ((v40) | 1))
                                _ = v42
                            }
                            for fi43 := 0; fi43 < 7; fi43++ {
                                v4[v1] = "foo bar"
                                v21 -= (^(v36))
                                _ = fi43
                            }
                            _ = v41
                        }
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                    _ = v37
                } else {
                    switch -1 {
                    case -1:
                        var v44 float32 = 1.5
                        _ = v44
                    case -9223372036854775808:
                        fmt.Println(v1)
                    default:
                        v36 += v21
                    }
                }
                _ = v36
            } else {
                v45 := []int{32768, 4294967295}
                v5 -= (^(v2))
                _ = v45
            }
            fmt.Println(v1, v23)
        }
        if (42 != 924) {
            v46 := GenS1{F0: "a", F1: (v23 / ((v23) | 1))}
            _ = v46
        }
    }
    for fi47 := 0; fi47 < 10; fi47++ {
        v2 = (v5 ^ (v2 * (v2 << ((255) & 63))))
        switch (fi47 + (fi47 & fi47)) {
        case 4294967295:
            var v48 uint16 = 509
            _ = v48
        case -9223372036854775808:
            v21 += v21
        case -821:
            v49 := GenS1{F0: "hello", F1: v23}
            _ = v49
        default:
            var v50 float64 = 1.5
            _ = v50
        }
        _ = fi47
    }
    if false {
        if false {
            v51 := GenS6{F0: 740, F1: 127, F2: 2147483647}
            v52 := GenS1{F0: ("hello" + ((("\n\t" + ("" + "foo bar")) + "hello") + "")), F1: 0}
            _ = v51
            _ = v52
        }
        v21 += int32(v1)
    } else {
        var v53 float64 = 0.0
        _ = v53
    }
    v2 -= 0
    for fi54 := 0; fi54 < 5; fi54++ {
        v1 -= (v1 << ((315) & 63))
        v2 = 127
        switch int(v23) {
        case 32767:
            fmt.Println(v2, v23, v2)
        case 2147483647:
            v21 = (v21 % (((v21 * v21)) | 1))
        }
        _ = fi54
    }
    var v55 int16 = -96
    var v56 int64 = -842
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
}

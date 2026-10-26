package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS3 struct { F0 string; F1 int; F2 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32; F1 bool; F2 bool }
func genFn0() int8 {
    v4 := []int{-9223372036854775808, -2147483648}
    var v5 uint16 = 32767
    if false {
        v5 = v5
        v6 := GenS1{F0: 2}
        var v7 int8 = -109
        _ = v6
        _ = v7
    } else {
        switch -1 {
        case -948:
            switch 100 {
            case -9223372036854775808:
                switch "hello" {
                case "":
                    v5 = v5
                default:
                    v8 := GenS1{F0: v5}
                    _ = v8
                }
            default:
                if false {
                    for fi9 := 0; fi9 < 16; fi9++ {
                        _ = len(v4)
                        _ = fi9
                    }
                    v10 := GenS2{F0: v5, F1: (v5 % (((v5 | v5)) | 1)), F2: 0}
                    switch 128 {
                    case 9223372036854775807:
                        if false {
                            var v11 bool = (1 >= 1)
                            _ = v11
                        } else {
                            var v12 uint = 2
                            _ = v12
                        }
                    default:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    _ = v10
                }
            }
        default:
            var v13 string = "a"
            _ = v13
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 string = ""
            _ = fi14
            _ = v15
        }
    }
    var v16 uint32 = 0
    for fi17 := 0; fi17 < 2; fi17++ {
        for fi18 := 0; fi18 < 19; fi18++ {
            v19 := GenS3{F0: "", F1: (^(fi18)), F2: (fi18 ^ (fi18 + 127))}
            _ = fi18
            _ = v19
        }
        var v20 int64 = -1
        _ = fi17
        _ = v20
    }
    _ = v16
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[string]string)
    var v3 uint = 422
    var v21 uint8 = 1
    v22 := []int{375}
    v23 := GenS2{F0: 145, F1: 0, F2: 1}
    v24 := GenS2{F0: 2, F1: 734, F2: 51}
    _ = len(v2)
    if false {
        v25 := GenS4{F0: (v3 - v3)}
        v21 += (v21 % ((v21) | 1))
        v3 += (v3 << ((v3) & 63))
        _ = v25
    } else {
        if (!true) {
            v1 += -2.25
            v1 += v1
            v1 = float64(1.5)
        } else {
            v3 = v3
        }
        var v26 int = 1
        switch "a" {
        case "\n\t":
            for fi27 := 0; fi27 < 3; fi27++ {
                v28 := GenS5{F0: 2147483647, F1: ((0.1 + 0.0) <= (v1 - (v1 / v1))), F2: (true && true)}
                _ = fi27
                _ = v28
            }
        case "a":
            if (127 != v26) {
                for fi29 := 0; fi29 < 9; fi29++ {
                    fmt.Println(v1, v3)
                    _ = fi29
                }
            } else {
                var v30 int64 = 1
                _ = v30
            }
        case "hello":
            v3 = 466
        }
        _ = v26
    }
    fmt.Println(v21, v21, v1)
    delete(v2, ("hello" + "hello"))
    if false {
        v31 := GenS4{F0: (v3 >> ((255) & 63))}
        v32 := GenS1{F0: 2}
        _ = v31
        _ = v32
    } else {
        v33 := GenS1{F0: 255}
        switch 128 {
        case 2147483647:
            v3 += v3
        case 210:
            if false {
                for fi34 := 0; fi34 < 10; fi34++ {
                    v3 = v3
                    switch "" {
                    case "":
                        v3 = v3
                    }
                    _ = fi34
                }
                if true {
                    var v35 uint = 4294967295
                    fmt.Println(v35, v35, v1)
                    _ = v35
                } else {
                    for fi36 := 0; fi36 < 17; fi36++ {
                        var v37 uint64 = 397
                        delete(v2, "hello")
                        _ = fi36
                        _ = v37
                    }
                }
                v21 = (v21 / (((^(v21))) | 1))
            }
        }
        var v38 int32 = 689
        _ = v33
        _ = v38
    }
    for fi39 := 0; fi39 < 5; fi39++ {
        var v40 uint16 = 0
        _ = fi39
        _ = v40
    }
    _ = len(v2)
    fmt.Println(v3)
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}

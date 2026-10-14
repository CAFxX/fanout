package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 int = 127
    var v2 int64 = 42
    var v3 float32 = (1e10 + 0.0)
    v3 = (v3 - ((v3 * v3) * 0.1))
    fmt.Println(v1, v1, v3)
    v1 += (v1 << ((427) & 63))
    if (1 != v2) {
        var v4 uint32 = 42
        _ = v4
    }
    if true {
        var v5 uint16 = 920
        _ = v5
    }
    switch (v1 * v1) {
    case -1:
        for fi6 := 0; fi6 < 10; fi6++ {
            if false {
                fmt.Println(v3)
                v7 := []int{fi6}
                _ = v7
            } else {
                v8 := GenS1{F0: 2, F1: 65535}
                if false {
                    fmt.Println(fi6)
                    v1 += v1
                    if (false || (!(1 > -462))) {
                        var v9 uint8 = 250
                        v10 := []int{}
                        _ = v10
                        _ = v9
                    }
                } else {
                    fmt.Println(v1)
                }
                _ = v8
            }
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 int8 = 8
                var v13 int32 = 79
                _ = fi11
                _ = v12
                _ = v13
            }
            var v14 int8 = 42
            _ = fi6
            _ = v14
        }
    case 42:
        for fi15 := 0; fi15 < 6; fi15++ {
            v1 = fi15
            var v16 float32 = float32(float64(v3))
            var v17 float32 = (v3 + v3)
            _ = fi15
            _ = v16
            _ = v17
        }
    case 127:
        if false {
            if (42 == 413) {
                v18 := GenS1{F0: 2, F1: 255}
                v3 += 1.5
                for fi19 := 0; fi19 < 20; fi19++ {
                    fmt.Println(v1, v2)
                    v20 := []int{}
                    switch fi19 {
                    case 128:
                        switch fi19 {
                        case 0:
                            v2 -= v2
                        }
                    case 0:
                        if true {
                            v2 = v2
                            v1 -= (v1 - v1)
                        } else {
                            var v21 float32 = ((v3 * float32(((0.1 * 3.14159) / float64(v3)))) + -2.25)
                            switch v1 {
                            case 2147483647:
                                switch (^(v1)) {
                                case -429:
                                    if false {
                                        v1 -= (^(fi19))
                                        for fi22 := 0; fi22 < 1; fi22++ {
                                            var v23 uint64 = 0
                                            v24 := GenS2{F0: 127, F1: 127}
                                            _ = len(v20)
                                            if false {
                                                fmt.Println(v3, fi22)
                                            }
                                            v2 = (v2 * 0)
                                            _ = fi22
                                            _ = v23
                                            _ = v24
                                        }
                                    } else {
                                        switch (fi19 / (((v1 | fi19)) | 1)) {
                                        case 339:
                                            var v25 int16 = -48
                                            _ = v25
                                        }
                                    }
                                case -281:
                                    v3 += 0.1
                                case 202:
                                    v21 = -2.25
                                default:
                                    v2 = (^(v2))
                                }
                            case 438:
                                v26 := GenS1{F0: 2, F1: 255}
                                _ = v26
                            case -118:
                                fmt.Println(fi19, v1, fi19)
                            }
                            _ = v21
                        }
                    }
                    _ = fi19
                    _ = v20
                }
                _ = v18
            } else {
                for fi27 := 0; fi27 < 6; fi27++ {
                    fmt.Println(v1, v1, v1)
                    _ = fi27
                }
            }
        }
    default:
        switch (v1 << ((2) & 63)) {
        case -9223372036854775808:
            var v28 uint = 2
            _ = v28
        }
    }
    if false {
        for fi29 := 0; fi29 < 18; fi29++ {
            v30 := []int{32768, fi29, (fi29 >> ((35) & 63))}
            _ = fi29
            _ = v30
        }
        v2 += 231
        var v31 int = v1
        _ = v31
    }
    v3 += 1e10
    v32 := GenS3{F0: 22}
    fmt.Println(v1, v1)
    switch 241 {
    case -895:
        for fi33 := 0; fi33 < 17; fi33++ {
            v34 := GenS2{F0: 127, F1: (v1 << ((843) & 63))}
            v2 = (v2 * 0)
            _ = fi33
            _ = v34
        }
    case 382:
        var v35 bool = (true || false)
        _ = v35
    case 42:
        if false {
            switch (v1 | (^(v1))) {
            case 1:
                for fi36 := 0; fi36 < 11; fi36++ {
                    var v37 string = ""
                    _ = fi36
                    _ = v37
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
}

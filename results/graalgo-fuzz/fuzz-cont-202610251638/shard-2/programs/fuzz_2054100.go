package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int32, p1 int32, p2 int8) uint64 {
    var v4 int = 32768
    for fi5 := 0; fi5 < 1; fi5++ {
        if false {
            for fi6 := 0; fi6 < 10; fi6++ {
                v7 := GenS1{F0: (p2 & (p2 + (p2 - -118))), F1: (-(fi5)), F2: 0}
                var v8 uint8 = 1
                v9 := GenS2{F0: (p2 / (((p2 * p2)) | 1))}
                _ = fi6
                _ = v7
                _ = v8
                _ = v9
            }
        }
        v10 := []int{v4}
        _ = fi5
        _ = v10
    }
    var v11 float32 = 0.1
    p0 = (-(p1))
    var v12 float32 = (float32(0.0) * (((float32((3.14159 - 0.0)) / v11) + v11) / -2.25))
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v4
    return 954
}

func main() {
    var v1 float32 = (3.14159 + 0.0)
    var v2 uint = 4294967295
    var v3 int = 9223372036854775807
    v3 += 1
    v13 := []int{}
    var v14 bool = true
    v15 := GenS2{F0: -128}
    fmt.Println(v14, v2, v1)
    v16 := GenS2{F0: 1}
    var v17 uint64 = 42
    v14 = (1 > 1000)
    if v14 {
        if len(v13) > 0 {
            v13[0] = 42
        }
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
    } else {
        v14 = false
        var v18 int16 = -598
        switch "hello" {
        case "\n\t":
            v19 := GenS1{F0: -22, F1: v3, F2: int16(v17)}
            _ = v19
        case "a":
            if v14 {
                if v14 {
                    var v20 string = ("" + "a")
                    if ((42 == 499) && (!((v17 ^ (v17 / ((42) | 1))) >= (^(v17))))) {
                        for fi21 := 0; fi21 < 14; fi21++ {
                            var v22 int = fi21
                            v13 = append(v13, (v22 * (fi21 - -364)))
                            _ = fi21
                            _ = v22
                        }
                    } else {
                        if false {
                            v17 = genFn0(710, -587, 0)
                        } else {
                            v23 := GenS1{F0: 111, F1: v3, F2: (v18 ^ 42)}
                            if v14 {
                                v24 := GenS1{F0: 1, F1: -1, F2: v18}
                                _ = v24
                            }
                            _ = len(v13)
                            _ = v23
                        }
                    }
                    _ = v20
                } else {
                    v17 = v17
                    v17 = (v17 & genFn0(-1, -2147483648, 42))
                }
                for fi25 := 0; fi25 < 3; fi25++ {
                    var v26 uint64 = uint64(v2)
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    if (641 != v2) {
                        switch int(v18) {
                        case 128:
                            for fi27 := 0; fi27 < 18; fi27++ {
                                fmt.Println(v17, v1, v2)
                                _ = fi27
                            }
                        default:
                            var v28 int = -9223372036854775808
                            _ = v28
                        }
                        v29 := GenS3{F0: 690}
                        v17 = uint64(v2)
                        _ = v29
                    }
                    _ = fi25
                    _ = v26
                }
            }
        default:
            var v30 int = v3
            _ = v30
        }
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}

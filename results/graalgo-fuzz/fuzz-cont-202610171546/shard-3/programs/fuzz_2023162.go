package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 bool; F2 int64 }
type GenS3 struct { F0 int16; F1 uint32; F2 bool }

func main() {
    var v1 uint16 = 65535
    var v2 int8 = -1
    var v3 uint64 = 4294967295
    var v4 uint16 = (^(v1))
    var v5 string = (((("\n\t" + "hello") + "a") + "") + "a")
    for fi6 := 0; fi6 < 15; fi6++ {
        v1 = v1
        var v7 uint32 = 65535
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 2147483647}
    switch -2147483648 {
    case -858:
        v9 := GenS1{F0: 128}
        _ = v9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := []int{(fi10 ^ -579), int(v1), (fi10 % ((32768) | 1))}
        v11 = append(v11, fi10)
        var v12 uint64 = (v3 % ((4294967295) | 1))
        _ = fi10
        _ = v11
        _ = v12
    }
    if true {
        switch 172 {
        case -934:
            v5 = v5
        case -1:
            for fi13 := 0; fi13 < 13; fi13++ {
                for fi14 := 0; fi14 < 14; fi14++ {
                    if true {
                        v1 = (v1 | v1)
                        var v15 bool = false
                        _ = v15
                    }
                    _ = fi14
                }
                _ = fi13
            }
        default:
            var v16 float64 = 3.14159
            _ = v16
        }
        var v17 int64 = 1
        _ = v17
    }
    if ((^(v2)) > v2) {
        if true {
            v18 := []int{}
            _ = v18
        }
        if false {
            fmt.Println(v5, v1)
            if false {
                for fi19 := 0; fi19 < 16; fi19++ {
                    var v20 uint8 = 171
                    switch (fi19 << ((413) & 63)) {
                    case 390:
                        fmt.Println(v1)
                    case 128:
                        var v21 int16 = 32767
                        _ = v21
                    case 4294967295:
                        if (false && false) {
                            v22 := GenS2{F0: 2147483647, F1: (0 != (v1 / ((v1) | 1))), F2: 0}
                            _ = v22
                        } else {
                            v23 := GenS3{F0: -32768, F1: 0, F2: false}
                            for fi24 := 0; fi24 < 14; fi24++ {
                                var v25 uint16 = (^(v1))
                                _ = fi24
                                _ = v25
                            }
                            switch v5 {
                            case "":
                                v1 -= (^(v4))
                            default:
                                var v26 int32 = -2147483648
                                _ = v26
                            }
                            _ = v23
                        }
                    }
                    _ = fi19
                    _ = v20
                }
                for fi27 := 0; fi27 < 11; fi27++ {
                    for fi28 := 0; fi28 < 8; fi28++ {
                        v29 := GenS2{F0: 2147483648, F1: false, F2: 292}
                        v4 += (v4 ^ 0)
                        _ = fi28
                        _ = v29
                    }
                    fmt.Println(v1)
                    _ = fi27
                }
            } else {
                v30 := GenS2{F0: 645, F1: false, F2: -1}
                _ = v30
            }
            fmt.Println(v4, v1)
        }
    }
    v31 := []int{1, -2147483648, 0, 726}
    if len(v31) > 0 {
        v31[0] = 0
    }
    fmt.Println(v5, v1, v3)
    v32 := GenS3{F0: 32767, F1: 565, F2: true}
    switch "a" {
    case "hello":
        v4 -= 42
    default:
        if false {
            var v33 float32 = -2.25
            _ = v33
        } else {
            var v34 int32 = -1
            var v35 string = ((("\n\t" + "a") + v5) + v5)
            v34 = (^(v34))
            _ = v34
            _ = v35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v8
}

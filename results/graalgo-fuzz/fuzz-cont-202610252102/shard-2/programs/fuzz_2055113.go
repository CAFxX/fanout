package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 uint64 = 2
    var v2 float32 = (3.14159 * 0.0)
    var v3 uint8 = 0
    var v4 string = ("hello" + "hello")
    var v5 uint = 65535
    var v6 float32 = (((v2 + float32(((-2.25 - 0.0) * -2.25))) - v2) * float32(0.0))
    if true {
        fmt.Println(v1, v4)
        v7 := GenS1{F0: v3, F1: 42}
        _ = v7
    } else {
        v8 := []int{127, 128, 42, -946}
        _ = v8
    }
    v1 = v1
    var v9 float32 = v6
    v10 := GenS1{F0: v3, F1: uint(v3)}
    v11 := GenS1{F0: (^(v3)), F1: (^(v5))}
    v12 := []int{2147483647, 32767, 127, 32767}
    for fi13 := 0; fi13 < 11; fi13++ {
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 bool = false
            var v16 uint16 = 65535
            fmt.Println(fi14, v4, v15)
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    var v17 int = 32767
    if (((v4 + v4) + v4) > v4) {
        for fi18 := 0; fi18 < 6; fi18++ {
            if true {
                for fi19 := 0; fi19 < 9; fi19++ {
                    fmt.Println(v1, v6)
                    _ = fi19
                }
                var v20 uint32 = 515
                var v21 uint32 = v20
                _ = v20
                _ = v21
            } else {
                v5 = 42
                if true {
                    if false {
                        v4 = ""
                    }
                    for fi22 := 0; fi22 < 17; fi22++ {
                        if (!false) {
                            var v23 float64 = 1.5
                            if (-1 > -1) {
                                var v24 bool = (!(false && true))
                                _ = v24
                            }
                            if ("hello" > ((v4 + ("a" + "hello")) + (v4 + ((((v4 + (v4 + "")) + "hello") + "\n\t") + "foo bar")))) {
                                _ = len(v12)
                            } else {
                                v6 = v9
                            }
                            switch (fi22 / (((fi22 ^ v17)) | 1)) {
                            case -727:
                                v9 -= v9
                            case 4294967295:
                                v25 := GenS1{F0: v3, F1: v5}
                                _ = v25
                            default:
                                switch v4 {
                                case "":
                                    v26 := GenS1{F0: (^(v3)), F1: 4294967295}
                                    _ = v26
                                }
                            }
                            fmt.Println(fi18)
                            _ = v23
                        }
                        v27 := GenS2{F0: 0}
                        if len(v12) > 0 {
                            v12[0] = v17
                        }
                        _ = fi22
                        _ = v27
                    }
                }
                if (fi18 <= int(v3)) {
                    v28 := GenS3{F0: 42}
                    fmt.Println(v3, v1)
                    _ = v28
                }
            }
            var v29 int = fi18
            _ = fi18
            _ = v29
        }
    } else {
        var v30 int64 = -46
        _ = v30
    }
    var v31 string = ((v4 + v4) + (v4 + v4))
    fmt.Println(v31, v4, v5)
    v3 += v3
    var v32 int = (v17 << ((v5) & 63))
    v1 = v1
    v1 = 35
    switch (v4 + v31) {
    case "hello":
        var v33 uint = v5
        _ = v33
    default:
        if false {
            v12 = append(v12, (v17 + (v17 / ((128) | 1))))
            var v34 int = (v32 ^ v17)
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            _ = v34
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}

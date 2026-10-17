package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 int16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = -2.25
    v3 := make(map[uint32]uint32)
    var v4 uint16 = 65535
    var v5 int64 = -9223372036854775808
    if true {
        v2 = ((1e10 * v2) - 0.1)
    } else {
        if false {
            for fi6 := 0; fi6 < 5; fi6++ {
                v7 := GenS1{F0: (-(fi6))}
                for fi8 := 0; fi8 < 2; fi8++ {
                    var v9 bool = (-2.25 == -2.25)
                    v5 += v5
                    _ = fi8
                    _ = v9
                }
                var v10 uint8 = 42
                _ = fi6
                _ = v10
                _ = v7
            }
            var v11 int = -939
            v11 = (v11 + (^(v11)))
            _ = v11
        }
        v2 += v2
        v12 := []int{4294967295, 159}
        _ = v12
    }
    if (v2 >= 1e10) {
        for fi13 := 0; fi13 < 19; fi13++ {
            for fi14 := 0; fi14 < 14; fi14++ {
                switch fi13 {
                case 4294967295:
                    var v15 uint64 = 0
                    _ = v15
                default:
                    v16 := v3[2147483648]
                    var v17 int = (-(fi14))
                    _ = v16
                    _ = v17
                }
                switch "\n\t" {
                case "\n\t":
                    v5 += v5
                case "a":
                    if true {
                        if true {
                            if true {
                                v18 := GenS1{F0: (fi13 * (fi13 * (^(fi14))))}
                                if true {
                                    var v19 uint = 433
                                    v20 := []int{(-(fi14)), (fi13 | fi13), 32767}
                                    _ = v19
                                    _ = v20
                                }
                                v4 = v4
                                _ = v18
                            } else {
                                switch (fi13 + fi13) {
                                case 0:
                                    var v21 uint8 = 2
                                    _ = v21
                                default:
                                    v4 = 830
                                }
                            }
                            v22 := GenS1{F0: fi13}
                            if ((true && (fi14 < 4294967295)) || (!true)) {
                                v23 := GenS2{F0: (fi13 << ((612) & 63)), F1: 1}
                                v24 := GenS2{F0: fi13, F1: 255}
                                _ = v23
                                _ = v24
                            }
                            _ = v22
                        } else {
                            v25 := v1["foo bar"]
                            v4 = v4
                            v26 := GenS2{F0: fi13, F1: 255}
                            _ = v25
                            _ = v26
                        }
                        v1[(("\n\t" + "\n\t") + "foo bar")] = 55
                        var v27 bool = (v2 == v2)
                        _ = v27
                    }
                default:
                    if ((1 != (fi13 % ((fi14) | 1))) && (!false)) {
                        var v28 float32 = 1e10
                        _ = v28
                    }
                }
                v29 := []int{fi14}
                _ = fi14
                _ = v29
            }
            _ = fi13
        }
        v5 -= v5
    }
    v2 = ((-2.25 - v2) - v2)
    delete(v1, "hello")
    v30 := GenS3{F0: 1000}
    switch "hello" {
    case "hello":
        fmt.Println(v2, v2)
    case "a":
        for fi31 := 0; fi31 < 5; fi31++ {
            var v32 uint64 = uint64(v4)
            v33 := []int{(-(fi31)), fi31, (fi31 >> ((762) & 63))}
            _ = fi31
            _ = v32
            _ = v33
        }
    default:
        v34 := []int{127, 1, -9223372036854775808, -9223372036854775808}
        _ = v34
    }
    for fi35 := 0; fi35 < 12; fi35++ {
        v5 = v5
        _ = fi35
    }
    var v36 float32 = v2
    var v37 int8 = 42
    for fi38 := 0; fi38 < 3; fi38++ {
        v37 = (v37 >> ((1) & 63))
        v36 = v2
        _ = fi38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}

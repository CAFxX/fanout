package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[int]string)
    var v2 int32 = -2147483648
    v3 := make(map[uint32]string)
    switch -1 {
    case 0:
        if false {
            var v4 int64 = -358
            v4 = 240
            var v5 bool = true
            _ = v4
            _ = v5
        } else {
            v6 := GenS1{F0: 9223372036854775807, F1: 732, F2: 748}
            v7 := GenS1{F0: -422, F1: 496, F2: 947}
            v2 -= (v2 >> ((2) & 63))
            _ = v6
            _ = v7
        }
    case -9223372036854775808:
        var v8 int16 = 0
        _ = v8
    case 32767:
        var v9 uint16 = 0
        _ = v9
    }
    if true {
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := GenS1{F0: -1, F1: 659, F2: 65535}
            _ = fi10
            _ = v11
        }
        if false {
            for fi12 := 0; fi12 < 19; fi12++ {
                var v13 bool = true
                _ = fi12
                _ = v13
            }
            var v14 int16 = -96
            _ = v14
        }
        for fi15 := 0; fi15 < 12; fi15++ {
            var v16 int16 = 42
            switch (fi15 >> ((655) & 63)) {
            case -9223372036854775808:
                for fi17 := 0; fi17 < 1; fi17++ {
                    if true {
                        var v18 float32 = 3.14159
                        if true {
                            v19 := GenS1{F0: int64(fi15), F1: 451, F2: 255}
                            if true {
                                v20 := []int{}
                                v18 = 1.5
                                if true {
                                    v2 = v2
                                    v21 := GenS1{F0: 1, F1: 65535, F2: 255}
                                    switch fi15 {
                                    case 9223372036854775807:
                                        var v22 float32 = 3.14159
                                        _ = v22
                                    case -9223372036854775808:
                                        var v23 uint16 = 32767
                                        _ = v23
                                    }
                                    _ = v21
                                }
                                v24 := GenS1{F0: 746, F1: 65535, F2: 32767}
                                _ = v20
                                _ = v24
                            } else {
                                v18 += v18
                            }
                            _ = v19
                        } else {
                            switch "\n\t" {
                            case "hello":
                                fmt.Println(v18)
                            case "":
                                v18 = v18
                            default:
                                _ = len(v3)
                            }
                            delete(v1, fi15)
                            if ((fi15 >> ((42) & 63)) >= (fi15 << ((292) & 63))) {
                                fmt.Println(v18, fi15)
                                v16 -= v16
                            } else {
                                var v25 uint16 = 762
                                v26 := GenS1{F0: 42, F1: 180, F2: (v25 | v25)}
                                v25 += (v25 | v25)
                                _ = v25
                                _ = v26
                            }
                        }
                        if (403 != 580) {
                            var v27 uint16 = 2
                            _ = v27
                        }
                        _ = v18
                    }
                    _ = fi17
                }
            }
            _ = fi15
            _ = v16
        }
    } else {
        v2 = v2
        var v28 float32 = float32(1.5)
        _ = v28
    }
    var v29 uint32 = 65535
    v30 := GenS1{F0: -9223372036854775808, F1: 673, F2: 32767}
    fmt.Println(v2, v2)
    v29 -= (v29 << ((0) & 63))
    fmt.Println(v2, v29)
    v31 := GenS2{F0: (v29 & 0)}
    fmt.Println(v29, v2)
    fmt.Println(v29, v2)
    var v32 int = -1
    var v33 uint64 = 797
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
}

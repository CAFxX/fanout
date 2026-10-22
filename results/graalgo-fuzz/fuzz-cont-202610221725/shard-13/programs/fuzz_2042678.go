package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 string; F1 uint; F2 uint8 }
type GenS4 struct { F0 int16; F1 bool }

func main() {
    var v1 int = -1
    var v2 uint16 = 2
    v3 := make(map[uint32]string)
    var v4 float32 = 0.0
    v5 := []int{}
    var v6 int16 = -1
    v7 := GenS1{F0: (v6 - 0), F1: 2}
    if len(v5) > 0 {
        v5[0] = v1
    }
    v8 := v3[0]
    v4 += 3.14159
    v9 := GenS1{F0: (v6 + v6), F1: 610}
    var v10 uint8 = 0
    if true {
        if false {
            if ((v1 | v1) < -2147483648) {
                if true {
                    var v11 uint32 = 65535
                    v4 += v4
                    _ = v11
                }
                if (v1 != (v1 | v1)) {
                    for fi12 := 0; fi12 < 2; fi12++ {
                        switch v8 {
                        case "\n\t":
                            for fi13 := 0; fi13 < 11; fi13++ {
                                v10 = (v10 << ((65535) & 63))
                                _ = fi13
                            }
                        default:
                            var v14 float64 = 0.1
                            _ = v14
                        }
                        _ = fi12
                    }
                    var v15 string = v8
                    _ = v15
                } else {
                    if false {
                        for fi16 := 0; fi16 < 8; fi16++ {
                            fmt.Println(fi16, v8, fi16)
                            var v17 int8 = 127
                            _ = fi16
                            _ = v17
                        }
                        if true {
                            if (false || false) {
                                var v18 string = v8
                                _ = v18
                            } else {
                                var v19 uint16 = (v2 * (v2 | 342))
                                _ = v19
                            }
                            var v20 uint64 = 1
                            var v21 uint16 = 865
                            _ = v20
                            _ = v21
                        } else {
                            var v22 uint16 = 649
                            _ = v22
                        }
                    }
                    delete(v3, 42)
                }
            } else {
                if false {
                    _ = len(v5)
                    fmt.Println(v6)
                }
            }
            v3[2147483647] = "hello"
            if len(v5) > 0 {
                v5[0] = v1
            }
        } else {
            delete(v3, 2147483648)
            v23 := GenS1{F0: v6, F1: 103}
            if (-47 < 100) {
                fmt.Println(v1)
            }
            _ = v23
        }
        v10 = (v10 + (v10 + 128))
        switch v1 {
        case 0:
            if (1 >= 42) {
                v4 -= v4
                fmt.Println(v8, v2)
                if true {
                    v6 -= v6
                    v24 := GenS1{F0: (v6 << ((167) & 63)), F1: 2147483648}
                    _ = v24
                }
            } else {
                var v25 int32 = 1
                _ = v25
            }
        case 32768:
            if true {
                v26 := GenS1{F0: -1, F1: uint32(v2)}
                if (42 >= 4294967295) {
                    v8 = v8
                    for fi27 := 0; fi27 < 16; fi27++ {
                        var v28 int8 = 127
                        v29 := GenS2{F0: 32767, F1: v28}
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                } else {
                    var v30 float64 = 1.5
                    switch (^(v1)) {
                    case 128:
                        var v31 uint64 = uint64(v2)
                        _ = v31
                    case 1:
                        v32 := GenS1{F0: v6, F1: 1}
                        _ = v32
                    case 4294967295:
                        v33 := GenS3{F0: "\n\t", F1: 963, F2: 0}
                        _ = v33
                    }
                    v34 := GenS4{F0: -45, F1: true}
                    _ = v30
                    _ = v34
                }
                v2 = 215
                _ = v26
            } else {
                var v35 uint16 = (^(v2))
                _ = v35
            }
        case 1:
            fmt.Println(v4, v1)
        default:
            for fi36 := 0; fi36 < 8; fi36++ {
                _ = len(v5)
                _ = fi36
            }
        }
    } else {
        v6 = v6
        switch v1 {
        case -1:
            var v37 int8 = -41
            _ = v37
        case 32768:
            v8 = "\n\t"
        }
    }
    v8 = v8
    v38 := GenS4{F0: v6, F1: true}
    var v39 int8 = 127
    v5 = append(v5, (v1 >> ((493) & 63)))
    _ = len(v5)
    var v40 float32 = v4
    for fi41 := 0; fi41 < 7; fi41++ {
        for fi42 := 0; fi42 < 10; fi42++ {
            v43 := GenS1{F0: v6, F1: 4294967295}
            _ = len(v3)
            _ = fi42
            _ = v43
        }
        v44 := GenS4{F0: v6, F1: true}
        _ = fi41
        _ = v44
    }
    fmt.Println(v2, v6, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 uint16; F1 int16; F2 int16 }
type GenS6 struct { F0 int8; F1 uint8; F2 uint8 }
func genFn0(p0 int, p1 uint) uint64 {
    var v4 string = "foo bar"
    v5 := GenS1{F0: -1, F1: 340, F2: (p1 % (((p1 ^ p1)) | 1))}
    var v6 string = (v4 + "a")
    var v7 uint16 = 110
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    v1 := []uint32{1, 719, 2147483648}
    v2 := make(map[uint32]uint32)
    v3 := []uint32{330, 628, 1}
    if (true && true) {
        var v8 int32 = 42
        _ = v8
    } else {
        v9 := GenS2{F0: 101}
        _ = v9
    }
    var v10 int16 = -1
    if false {
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 string = "hello"
            _ = fi11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        for fi14 := 0; fi14 < 6; fi14++ {
            v2[676] = 2
            if true {
                var v15 float32 = 1.5
                v15 = v15
                _ = len(v2)
                _ = v15
            }
            var v16 int32 = 704
            v17 := v2[769]
            _ = fi14
            _ = v16
            _ = v17
        }
        v18 := GenS3{F0: 11}
        _ = fi13
        _ = v18
    }
    fmt.Println(v10)
    for fi19 := 0; fi19 < 6; fi19++ {
        var v20 uint = 193
        _ = fi19
        _ = v20
    }
    v10 = v10
    var v21 uint = 42
    var v22 uint16 = 0
    if false {
        if false {
            if true {
                v10 = v10
                var v23 int32 = 2147483647
                v22 += 1
                _ = v23
            } else {
                _ = len(v1)
                switch (("foo bar" + ("" + "a")) + "") {
                case "":
                    v2[2147483647] = 456
                case "\n\t":
                    var v24 uint32 = 2147483648
                    _ = v24
                case "hello":
                    v25 := GenS4{F0: 365, F1: 15}
                    _ = v25
                }
                var v26 uint8 = 144
                _ = v26
            }
            if true {
                var v27 int64 = -721
                switch -202 {
                case 892:
                    var v28 int64 = v27
                    _ = v28
                case 4294967295:
                    for fi29 := 0; fi29 < 4; fi29++ {
                        delete(v2, 65535)
                        for fi30 := 0; fi30 < 9; fi30++ {
                            v10 = v10
                            if (true || false) {
                                fmt.Println(fi29, fi30, fi29)
                                fmt.Println(v22, fi29, fi30)
                                _ = len(v2)
                            }
                            var v31 uint16 = (v22 * v22)
                            var v32 float64 = 0.1
                            var v33 int8 = -128
                            _ = fi30
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                        v34 := GenS4{F0: genFn0(int(v27), (^(v21))), F1: 193}
                        _ = fi29
                        _ = v34
                    }
                case 128:
                    v21 = (v21 ^ v21)
                default:
                    var v35 int32 = 1
                    _ = v35
                }
                for fi36 := 0; fi36 < 9; fi36++ {
                    switch (-(fi36)) {
                    case 1:
                        for fi37 := 0; fi37 < 3; fi37++ {
                            var v38 uint16 = v22
                            v21 += (v21 + v21)
                            _ = fi37
                            _ = v38
                        }
                    default:
                        v39 := GenS2{F0: 545}
                        _ = v39
                    }
                    var v40 uint64 = 18446744073709551615
                    _ = fi36
                    _ = v40
                }
                _ = v27
            }
            if false {
                delete(v2, 586)
                for fi41 := 0; fi41 < 13; fi41++ {
                    if (110 <= 1) {
                        if false {
                            switch "foo bar" {
                            case "a":
                                v42 := GenS5{F0: v22, F1: v10, F2: (v10 * -32768)}
                                _ = v42
                            }
                            v43 := GenS6{F0: 100, F1: 0, F2: 255}
                            _ = v43
                        }
                    }
                    var v44 float64 = (((0.1 / (((3.14159 / (3.14159 - (3.14159 - ((1e10 - (1e10 + (1.5 - 0.1))) * (1e10 + -2.25))))) + 1.5) + 3.14159)) + 0.0) - 0.0)
                    _ = fi41
                    _ = v44
                }
                v45 := GenS5{F0: v22, F1: 0, F2: (-(v10))}
                _ = v45
            } else {
                var v46 uint = v21
                _ = v46
            }
        }
        v21 = v21
    }
    if false {
        fmt.Println(v22, v22, v10)
        if false {
            v47 := v2[2147483647]
            _ = v47
        }
        _ = len(v2)
    }
    var v48 uint = (v21 + v21)
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v48
}

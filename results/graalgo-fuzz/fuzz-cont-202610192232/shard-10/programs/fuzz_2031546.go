package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 uint; F2 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 string }
type GenS6 struct { F0 int32; F1 int32; F2 uint64 }

func main() {
    var v1 uint8 = 1
    var v2 float32 = 0.0
    var v3 uint8 = v1
    fmt.Println(v3, v3)
    v4 := GenS1{F0: ""}
    var v5 int16 = 0
    fmt.Println(v3)
    for fi6 := 0; fi6 < 16; fi6++ {
        switch fi6 {
        case 127:
            fmt.Println(v1)
        default:
            var v7 int64 = -1
            _ = v7
        }
        _ = fi6
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := []int{283}
        v10 := GenS1{F0: "foo bar"}
        _ = fi8
        _ = v10
        _ = v9
    }
    if (false || false) {
        v11 := GenS2{F0: 42}
        v1 -= (v3 & v1)
        _ = v11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        for fi13 := 0; fi13 < 3; fi13++ {
            v14 := []int{127}
            v2 = v2
            _ = fi13
            _ = v14
        }
        v15 := GenS1{F0: ("\n\t" + ("foo bar" + "foo bar"))}
        v16 := GenS2{F0: 0}
        _ = fi12
        _ = v15
        _ = v16
    }
    if (251 < 824) {
        v1 = (v1 << ((1) & 63))
        v1 += v1
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        if false {
            for fi18 := 0; fi18 < 13; fi18++ {
                v3 += v3
                _ = fi18
            }
            if (true && true) {
                switch 288 {
                case 4294967295:
                    for fi19 := 0; fi19 < 4; fi19++ {
                        var v20 int32 = 0
                        if false {
                            v1 -= (v3 << ((580) & 63))
                        } else {
                            if false {
                                v21 := GenS3{F0: 1, F1: 186, F2: (v5 * v5)}
                                if (true && true) {
                                    if true {
                                        v5 = v5
                                    } else {
                                        if ((v3 << ((2) & 63)) == v3) {
                                            v22 := []int{}
                                            fmt.Println(v3, v20, v5)
                                            fmt.Println(fi19)
                                            _ = v22
                                        }
                                    }
                                }
                                _ = v21
                            } else {
                                if false {
                                    v1 = v3
                                }
                                fmt.Println(fi19, fi19)
                            }
                        }
                        switch (-(fi19)) {
                        case -1:
                            fmt.Println(fi17, fi17)
                        case 9223372036854775807:
                            var v23 float32 = (float32(1.5) + v2)
                            _ = v23
                        default:
                            v3 += (v3 / ((255) | 1))
                        }
                        _ = fi19
                        _ = v20
                    }
                case 32768:
                    if (false && (2147483647 < 42)) {
                        v24 := []int{(-(fi17)), 57}
                        _ = v24
                    }
                }
                for fi25 := 0; fi25 < 14; fi25++ {
                    var v26 string = "a"
                    switch ("" + (v26 + v26)) {
                    case "hello":
                        v2 -= v2
                    case "foo bar":
                        var v27 int64 = 0
                        _ = v27
                    default:
                        v5 = (v5 * 0)
                    }
                    switch v26 {
                    case "a":
                        var v28 uint64 = 902
                        _ = v28
                    case "foo bar":
                        v29 := GenS3{F0: -9223372036854775808, F1: 1, F2: (v5 / ((v5) | 1))}
                        _ = v29
                    case "":
                        var v30 int = fi17
                        if (0.1 >= 0.1) {
                            var v31 uint8 = 255
                            fmt.Println(v2, fi25, fi17)
                            v31 += v1
                            _ = v31
                        } else {
                            v32 := GenS4{F0: (v3 / (((v1 >> ((0) & 63))) | 1))}
                            _ = v32
                        }
                        _ = v30
                    }
                    _ = fi25
                    _ = v26
                }
            } else {
                if ((!true) || true) {
                    fmt.Println(v3)
                    v33 := []int{fi17, fi17, fi17}
                    _ = v33
                }
                var v34 uint32 = 42
                v35 := GenS2{F0: 30}
                _ = v34
                _ = v35
            }
        }
        switch -2147483648 {
        case -551:
            if false {
                var v36 int64 = 542
                _ = v36
            }
        case -408:
            v1 -= 127
        }
        v37 := GenS5{F0: 65535, F1: ("\n\t" + "")}
        _ = fi17
        _ = v37
    }
    v38 := GenS6{F0: 0, F1: -1, F2: 18446744073709551615}
    v39 := []int{32767, -296, -779, 0}
    switch -2147483648 {
    case 9223372036854775807:
        fmt.Println(v1)
    default:
        v2 = v2
    }
    v5 += v5
    var v40 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}

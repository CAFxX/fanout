package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 uint8; F1 string; F2 int8 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]uint32)
    v4 := GenS1{F0: 2, F1: 235}
    switch 1 {
    case 28:
        v5 := []int{2147483647, -1, -2147483648, -681}
        _ = v5
    case 4294967295:
        v6 := GenS1{F0: 923, F1: 42}
        _ = v6
    case -1:
        _ = len(v2)
    default:
        v7 := v1[1]
        _ = v7
    }
    var v8 uint64 = 539
    var v9 float32 = 0.0
    if false {
        v9 = float32(3.14159)
        v8 += 1
    }
    var v10 int16 = 1000
    switch 4294967295 {
    case -603:
        var v11 float64 = ((3.14159 / float64(float32(-2.25))) / float64((v9 - ((((3.14159 / ((3.14159 * v9) / v9)) / v9) + (v9 + v9)) + (v9 / (0.0 * float32(float64((0.1 / v9)))))))))
        _ = v11
    case -173:
        v12 := GenS1{F0: 391, F1: 1}
        _ = v12
    case 9223372036854775807:
        for fi13 := 0; fi13 < 7; fi13++ {
            var v14 string = ""
            for fi15 := 0; fi15 < 12; fi15++ {
                v16 := GenS2{F0: 42, F1: 990}
                v8 -= (^(v8))
                _ = fi15
                _ = v16
            }
            _ = fi13
            _ = v14
        }
    }
    fmt.Println(v8)
    var v17 int8 = 0
    v17 = (v17 >> ((0) & 63))
    switch 164 {
    case -9223372036854775808:
        fmt.Println(v10, v17, v10)
    default:
        v10 -= (v10 | v10)
    }
    if false {
        v18 := []int{-1, -9223372036854775808}
        _ = v18
    } else {
        var v19 float64 = 1e10
        var v20 uint8 = 134
        if false {
            switch 32767 {
            case -508:
                if false {
                    _ = len(v1)
                    var v21 uint64 = uint64(v20)
                    _ = v21
                } else {
                    for fi22 := 0; fi22 < 15; fi22++ {
                        var v23 uint = 1
                        v24 := []int{fi22}
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 6; fi25++ {
                        for fi26 := 0; fi26 < 1; fi26++ {
                            if (("\n\t" + ("a" + "a")) < "a") {
                                var v27 string = "hello"
                                v28 := GenS3{F0: (!((1.5 != (v9 - v9)) && (true || false))), F1: v20}
                                var v29 uint16 = 960
                                switch ((v27 + ((v27 + ("foo bar" + v27)) + "")) + v27) {
                                case "\n\t":
                                    var v30 uint32 = 4294967295
                                    var v31 uint32 = uint32(v20)
                                    delete(v2, uint32(v8))
                                    v8 += (v8 << ((2) & 63))
                                    _ = v30
                                    _ = v31
                                default:
                                    v32 := []int{(^(fi25)), fi25}
                                    if true {
                                        var v33 int16 = (v10 / (((v10 / (((v10 % (((v10 << ((26) & 63))) | 1))) | 1))) | 1))
                                        v34 := GenS1{F0: 18446744073709551615, F1: 63}
                                        switch fi25 {
                                        case -1:
                                            v27 = ("hello" + "\n\t")
                                        case 0:
                                            var v35 int = (fi26 % (((fi25 ^ fi26)) | 1))
                                            _ = v35
                                        case 32767:
                                            var v36 int32 = -767
                                            _ = v36
                                        default:
                                            switch v27 {
                                            case "foo bar":
                                                var v37 float64 = 1.5
                                                _ = v37
                                            default:
                                                v8 = 4294967295
                                            }
                                        }
                                        _ = v33
                                        _ = v34
                                    } else {
                                        delete(v2, 0)
                                    }
                                    _ = v32
                                }
                                fmt.Println(fi25, v19, v19)
                                v38 := GenS1{F0: 481, F1: (^(v20))}
                                if false {
                                    v39 := GenS2{F0: v17, F1: 0}
                                    if (65535 == 2) {
                                        v40 := []int{}
                                        v8 = v8
                                        v41 := v2[2]
                                        _ = v40
                                        _ = v41
                                    } else {
                                        v10 = (v10 * v10)
                                        v9 -= v9
                                        v42 := GenS4{F0: (v29 >> ((622) & 63)), F1: (^(fi26))}
                                        v43 := []int{fi26, int(v17), 42, (fi26 & (fi26 + fi25))}
                                        _ = v42
                                        _ = v43
                                    }
                                    _ = len(v3)
                                    _ = v39
                                }
                                _ = v27
                                _ = v28
                                _ = v29
                                _ = v38
                            }
                            v44 := GenS5{F0: v20, F1: "hello", F2: (-(v17))}
                            v45 := GenS3{F0: (true || false), F1: 2}
                            _ = fi26
                            _ = v44
                            _ = v45
                        }
                        _ = fi25
                    }
                }
            }
        }
        _ = v19
        _ = v20
    }
    v10 = v10
    v9 = float32(3.14159)
    v9 += v9
    for fi46 := 0; fi46 < 18; fi46++ {
        for fi47 := 0; fi47 < 18; fi47++ {
            v9 = 0.0
            v48 := []int{(fi46 ^ (fi46 >> ((65535) & 63))), (fi47 * (fi47 - fi47)), fi47}
            _ = fi47
            _ = v48
        }
        _ = fi46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}

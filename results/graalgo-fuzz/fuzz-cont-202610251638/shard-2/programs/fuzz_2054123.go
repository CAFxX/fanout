package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 int8 }
func genFn0(p0 uint32, p1 uint32) int16 {
    var v5 int32 = 381
    v6 := []int{-521, -521}
    for fi7 := 0; fi7 < 4; fi7++ {
        if true {
            var v8 int64 = 0
            _ = v8
        }
        _ = fi7
    }
    v6 = append(v6, 747)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -1
}
func genFn1() uint32 {
    var v9 uint16 = 255
    v10 := []int{}
    _ = v10
    _ = v9
    return 2147483648
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]string)
    var v3 int64 = 140
    var v4 uint16 = 973
    var v11 int = -2147483648
    for fi12 := 0; fi12 < 12; fi12++ {
        v13 := []int{9223372036854775807, v11}
        v14 := GenS1{F0: v3, F1: 560, F2: -1}
        if (true || true) {
            if len(v13) > 0 {
                v13[0] = 32767
            }
            _ = len(v13)
            switch "" {
            case "":
                if len(v13) > 0 {
                    v13[0] = (fi12 ^ (fi12 + v11))
                }
            default:
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
            }
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    v3 = (v3 / ((753) | 1))
    v15 := GenS1{F0: -100, F1: 2, F2: 1}
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 string = "a"
        _ = fi16
        _ = v17
    }
    var v18 uint64 = 42
    v3 = int64(v4)
    v4 = (^(v4))
    v4 = (^(v4))
    for fi19 := 0; fi19 < 7; fi19++ {
        if true {
            v11 += fi19
            var v20 int = v11
            _ = v20
        } else {
            switch 32767 {
            case 4294967295:
                fmt.Println(v18, v11, v4)
            }
            for fi21 := 0; fi21 < 20; fi21++ {
                if false {
                    if false {
                        v18 = 4294967295
                    } else {
                        v22 := GenS1{F0: (v3 & v3), F1: 1, F2: -11}
                        _ = v22
                    }
                    _ = len(v2)
                } else {
                    v23 := GenS2{F0: 4294967295, F1: v18}
                    _ = v23
                }
                _ = fi21
            }
        }
        v4 = 325
        _ = fi19
    }
    fmt.Println(v3)
    for fi24 := 0; fi24 < 13; fi24++ {
        if false {
            for fi25 := 0; fi25 < 4; fi25++ {
                delete(v1, 65535)
                v26 := GenS3{F0: (!false)}
                _ = fi25
                _ = v26
            }
            switch "" {
            case "foo bar":
                v27 := GenS4{F0: -2147483648, F1: 127}
                _ = v27
            case "":
                v28 := []int{(fi24 << ((1) & 63)), 128, (v11 ^ (v11 & fi24)), (v11 - (fi24 - fi24))}
                _ = v28
            default:
                for fi29 := 0; fi29 < 6; fi29++ {
                    var v30 uint = 2
                    var v31 uint = uint(v4)
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            }
            var v32 uint16 = uint16(v18)
            _ = v32
        } else {
            switch -9223372036854775808 {
            case 277:
                for fi33 := 0; fi33 < 10; fi33++ {
                    if false {
                        switch "hello" {
                        case "hello":
                            v3 = v3
                            v34 := GenS1{F0: (v3 + v3), F1: 2, F2: 127}
                            v11 = -9223372036854775808
                            var v35 uint8 = uint8(v4)
                            _ = v34
                            _ = v35
                        }
                        v4 += 280
                        delete(v2, fi24)
                    } else {
                        v36 := []int{(fi24 / (((v11 | (fi33 % (((fi33 | fi33)) | 1)))) | 1)), (-(fi33)), fi33, (fi24 >> ((4294967295) & 63))}
                        v4 += v4
                        _ = v36
                    }
                    _ = fi33
                }
            case 466:
                v37 := []int{}
                _ = v37
            case 128:
                fmt.Println(fi24, fi24, v4)
            default:
                v3 -= (^(v3))
            }
            v11 = 0
            var v38 float64 = 0.1
            _ = v38
        }
        var v39 float32 = 1e10
        _ = fi24
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 uint32 }
func genFn0(p0 int32, p1 int8, p2 int64) int32 {
    var v7 string = ""
    v8 := []int{795}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p0
}

func main() {
    var v1 uint64 = 2
    var v2 uint16 = 539
    var v3 float64 = 1.5
    v4 := make(map[string]bool)
    var v5 int16 = 0
    v6 := make(map[string]uint32)
    v2 = (v2 >> ((18446744073709551615) & 63))
    fmt.Println(v2, v5)
    if (false && (true && true)) {
        if true {
            if false {
                v9 := GenS1{F0: 1, F1: 42}
                _ = v9
            } else {
                var v10 uint64 = (v1 << ((36) & 63))
                _ = v10
            }
            if (86 >= -128) {
                if false {
                    v5 -= v5
                }
                if false {
                    if (uint8(v1) <= 2) {
                        v11 := GenS2{F0: 19}
                        if (v1 < (v1 / ((2) | 1))) {
                            v12 := v4["a"]
                            v1 = (^(v1))
                            var v13 float64 = v3
                            _ = v12
                            _ = v13
                        } else {
                            v14 := GenS3{F0: 42, F1: 732}
                            var v15 float32 = (1e10 * float32(v3))
                            switch 32768 {
                            case 19:
                                v1 = (v1 - (v1 << ((42) & 63)))
                            case 32768:
                                var v16 uint32 = 198
                                _ = v16
                            }
                            _ = v14
                            _ = v15
                        }
                        for fi17 := 0; fi17 < 11; fi17++ {
                            if (!(0 != 117)) {
                                v1 = v1
                            } else {
                                v18 := []int{fi17}
                                if (!(1 < -9223372036854775808)) {
                                    for fi19 := 0; fi19 < 10; fi19++ {
                                        fmt.Println(v5)
                                        _ = fi19
                                    }
                                    v1 += 0
                                }
                                v3 -= v3
                                _ = v18
                            }
                            var v20 float64 = v3
                            for fi21 := 0; fi21 < 19; fi21++ {
                                fmt.Println(fi17, fi17)
                                v22 := GenS2{F0: -128}
                                _ = fi21
                                _ = v22
                            }
                            _ = fi17
                            _ = v20
                        }
                        _ = v11
                    }
                }
                for fi23 := 0; fi23 < 3; fi23++ {
                    fmt.Println(v5, v1, v1)
                    if true {
                        var v24 int16 = (v5 / ((v5) | 1))
                        _ = v24
                    }
                    v5 += v5
                    _ = fi23
                }
            }
        } else {
            v25 := GenS1{F0: 128, F1: 42}
            _ = v25
        }
        var v26 uint8 = 114
        _ = v26
    }
    for fi27 := 0; fi27 < 5; fi27++ {
        if true {
            switch "a" {
            case "a":
                var v28 float64 = (v3 - (float64(0.1) + v3))
                _ = v28
            case "hello":
                for fi29 := 0; fi29 < 4; fi29++ {
                    if true {
                        v1 -= (v1 % (((^(v1))) | 1))
                    }
                    v2 += 2
                    _ = fi29
                }
            }
            v5 -= (v5 << ((4294967295) & 63))
        } else {
            fmt.Println(fi27, v1, v1)
            v30 := GenS2{F0: -16}
            _ = v30
        }
        _ = fi27
    }
    delete(v4, "\n\t")
    for fi31 := 0; fi31 < 5; fi31++ {
        var v32 float64 = 1e10
        _ = fi31
        _ = v32
    }
    var v33 uint16 = v2
    var v34 int8 = 0
    var v35 uint = 764
    v36 := GenS2{F0: v34}
    v2 = v33
    v37 := v4["\n\t"]
    var v38 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}

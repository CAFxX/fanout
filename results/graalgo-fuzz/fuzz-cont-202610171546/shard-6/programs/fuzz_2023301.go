package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 int64 }
func genFn0() uint32 {
    v5 := []int{}
    var v6 float64 = (1.5 * 1e10)
    if true {
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 uint64 = 18446744073709551615
            _ = fi7
            _ = v8
        }
        _ = len(v5)
        v6 += 0.1
    }
    v6 -= 1.5
    v6 -= (float64((float32(-2.25) * 1.5)) + ((1.5 / (v6 + v6)) + ((3.14159 + float64(float32((v6 / (0.1 + v6))))) / (float64(-2.25) - v6))))
    _ = v5
    _ = v6
    return 2147483648
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = 1
    v3 := make(map[int]bool)
    var v4 int16 = 42
    delete(v3, 32767)
    v2 += -1
    v9 := GenS1{F0: 100, F1: (v4 * v4)}
    var v10 int32 = v2
    v11 := GenS2{F0: 42}
    switch (("a" + ("\n\t" + "\n\t")) + "foo bar") {
    case "":
        switch 32768 {
        case 4294967295:
            v12 := GenS1{F0: -87, F1: -1}
            _ = v12
        }
    case "a":
        if (false && (42 >= 0)) {
            var v13 uint32 = 29
            var v14 bool = ((0.0 * (float64(((float32((3.14159 * (3.14159 * 1e10))) - 3.14159) + 3.14159)) - (float64(((float32(1e10) / 3.14159) - 3.14159)) + 0.1))) >= 1.5)
            v10 = (v2 + (v10 & v10))
            _ = v13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 int16 = v4
        switch (fi15 / (((fi15 + (fi15 - (fi15 + (fi15 + fi15))))) | 1)) {
        case -1:
            switch "\n\t" {
            case "foo bar":
                for fi17 := 0; fi17 < 8; fi17++ {
                    v2 = v10
                    var v18 uint8 = 255
                    _ = fi17
                    _ = v18
                }
            }
        case -2147483648:
            for fi19 := 0; fi19 < 13; fi19++ {
                if ("" >= ("a" + "\n\t")) {
                    if (221 < 2147483647) {
                        fmt.Println(fi19, v2, fi15)
                    } else {
                        if true {
                            if (2147483648 <= 1) {
                                v20 := []int{fi15}
                                fmt.Println(v16, v4)
                                _ = v20
                            } else {
                                v21 := GenS1{F0: 42, F1: (v16 * int16(fi19))}
                                _ = v21
                            }
                        } else {
                            if ((-(fi15)) == (fi15 >> ((556) & 63))) {
                                switch fi15 {
                                case 32767:
                                    var v22 float32 = 0.0
                                    _ = v22
                                case -9223372036854775808:
                                    if ((v16 - (-(v16))) >= (v4 | 1000)) {
                                        if false {
                                            v23 := GenS2{F0: 9223372036854775807}
                                            var v24 int = (fi15 / (((fi15 ^ (fi19 / ((fi19) | 1)))) | 1))
                                            v25 := GenS1{F0: 26, F1: 42}
                                            v26 := GenS1{F0: 63, F1: v4}
                                            v3[(fi15 ^ fi19)] = ("" == "")
                                            _ = v23
                                            _ = v24
                                            _ = v25
                                            _ = v26
                                        } else {
                                            v3[fi19] = false
                                            var v27 float32 = (3.14159 + (1.5 / 1.0))
                                            v2 = v2
                                            _ = v27
                                        }
                                        v1[285] = (405 == 924)
                                    } else {
                                        fmt.Println(fi19, v10)
                                        var v28 int = fi19
                                        var v29 uint = 652
                                        _ = v28
                                        _ = v29
                                    }
                                default:
                                    var v30 uint16 = 885
                                    v3[-1] = true
                                    _ = v30
                                }
                            } else {
                                v31 := []int{(-(fi15)), (fi19 - -290), fi15, fi15}
                                v4 += (v16 << ((65535) & 63))
                                v32 := GenS2{F0: 1}
                                _ = v31
                                _ = v32
                            }
                            if true {
                                v33 := []int{32768, -441}
                                v16 -= (v16 << ((255) & 63))
                                _ = v33
                            }
                        }
                    }
                    switch ("foo bar" + ((("\n\t" + "hello") + "a") + "")) {
                    case "hello":
                        v34 := GenS2{F0: 9223372036854775807}
                        _ = v34
                    default:
                        var v35 bool = true
                        _ = v35
                    }
                    v36 := GenS3{F0: 943, F1: 0}
                    _ = v36
                }
                _ = fi19
            }
        case 0:
            v10 = v10
        default:
            var v37 float64 = 0.0
            _ = v37
        }
        _ = fi15
        _ = v16
    }
    v38 := GenS4{F0: -9223372036854775808}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v9
}

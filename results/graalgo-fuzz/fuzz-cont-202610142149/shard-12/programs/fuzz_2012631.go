package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint }
type GenS2 struct { F0 int64; F1 uint64; F2 bool }
type GenS3 struct { F0 int16; F1 int; F2 uint16 }
type GenS4 struct { F0 int64; F1 uint32; F2 bool }
type GenS5 struct { F0 uint16 }
func genFn0() uint8 {
    var v7 float64 = (0.0 + ((0.1 - 1e10) * ((1e10 * (1.5 * (3.14159 + (1.5 * (3.14159 * (-2.25 * (float64(0.0) / (((-2.25 - 1e10) * -2.25) - 0.0)))))))) * (0.1 / ((1.5 + 0.1) - (((0.1 * 3.14159) / 0.1) + 1.5))))))
    var v8 uint64 = 862
    _ = v7
    _ = v8
    return 72
}

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 517
    var v3 int = 450
    var v4 uint64 = (^(v2))
    v5 := make(map[int]uint32)
    v6 := make(map[int]int)
    v9 := v5[868]
    v10 := GenS1{F0: 2, F1: -1, F2: 545}
    var v11 float32 = 1.5
    delete(v1, "foo bar")
    fmt.Println(v9, v9)
    v12 := GenS1{F0: uint8(v9), F1: 976, F2: 18446744073709551615}
    switch "foo bar" {
    case "a":
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 bool = true
            v14 = true
            var v15 uint8 = 42
            _ = fi13
            _ = v14
            _ = v15
        }
    case "":
        if (316 == 2) {
            if false {
                delete(v6, v3)
            } else {
                for fi16 := 0; fi16 < 5; fi16++ {
                    var v17 uint64 = v4
                    switch "" {
                    case "a":
                        v4 = 18446744073709551615
                    }
                    if false {
                        if false {
                            v18 := []int{}
                            if (!(((v9 % ((v9) | 1)) != v9) || false)) {
                                v2 -= (v4 - v2)
                                v19 := GenS2{F0: 0, F1: 2, F2: false}
                                v20 := GenS3{F0: -1, F1: (fi16 & (v3 & fi16)), F2: 32767}
                                _ = v19
                                _ = v20
                            } else {
                                v17 += v17
                                v9 = (v9 | v9)
                                var v21 uint16 = 932
                                _ = v21
                            }
                            _ = v18
                        }
                        if false {
                            v9 = (^(v9))
                            var v22 int = 128
                            if ((v2 / ((v17) | 1)) != (v17 | 1)) {
                                v23 := GenS3{F0: 32767, F1: fi16, F2: 2}
                                _ = v23
                            }
                            _ = v22
                        } else {
                            _ = len(v5)
                            v24 := []int{}
                            var v25 string = ("\n\t" + ("" + ""))
                            _ = v24
                            _ = v25
                        }
                        switch (("" + "hello") + "\n\t") {
                        case "a":
                            v9 += 639
                        case "":
                            if true {
                                v26 := []int{-47, v3, (fi16 & 9223372036854775807), v3}
                                if (true || (18 != int16(v4))) {
                                    fmt.Println(v2, v17, v9)
                                    v27 := GenS4{F0: 1, F1: v9, F2: (true || true)}
                                    _ = v27
                                }
                                _ = v26
                            }
                        }
                    } else {
                        fmt.Println(v4, v4)
                        switch 32767 {
                        case 32767:
                            v28 := GenS2{F0: -1, F1: v4, F2: (true && (((32767 < -32768) && (127 == (^(fi16)))) && (v17 != 42)))}
                            _ = v28
                        case -2147483648:
                            v3 = 32767
                        }
                    }
                    _ = fi16
                    _ = v17
                }
            }
            if false {
                fmt.Println(v3)
                var v29 uint8 = 1
                fmt.Println(v9)
                _ = v29
            } else {
                switch -1 {
                case 764:
                    for fi30 := 0; fi30 < 13; fi30++ {
                        v31 := GenS5{F0: 631}
                        var v32 uint64 = 1
                        v9 += v9
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                case 1:
                    var v33 uint64 = (^(v4))
                    _ = v33
                }
            }
            var v34 int8 = 100
            _ = v34
        }
    case "foo bar":
        switch v3 {
        case 730:
            v35 := []int{(v3 << ((255) & 63)), 32768, v3}
            _ = v35
        }
    }
    v36 := GenS2{F0: -1, F1: v4, F2: (((42 < -1) || false) || (-9223372036854775808 == -636))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}

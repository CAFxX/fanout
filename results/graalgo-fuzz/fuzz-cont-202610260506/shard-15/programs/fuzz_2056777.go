package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }

func main() {
    var v1 uint8 = 127
    var v2 int = 9223372036854775807
    var v3 uint32 = 349
    var v4 int32 = -2147483648
    v5 := GenS1{F0: 9223372036854775807, F1: 4294967295}
    var v6 uint = 4294967295
    fmt.Println(v6)
    var v7 uint16 = 65535
    fmt.Println(v7)
    if (!(!(0 < 1000))) {
        v8 := []int{(v2 ^ 9223372036854775807)}
        for fi9 := 0; fi9 < 2; fi9++ {
            for fi10 := 0; fi10 < 13; fi10++ {
                v3 = (^(v3))
                if (true && (false || (!((!false) || false)))) {
                    v4 -= (v4 << ((255) & 63))
                    var v11 int16 = 1000
                    var v12 int64 = -1
                    _ = v11
                    _ = v12
                }
                var v13 uint = v6
                _ = fi10
                _ = v13
            }
            _ = fi9
        }
        _ = v8
    } else {
        var v14 bool = true
        v1 -= v1
        _ = v14
    }
    var v15 int32 = v4
    if (65535 >= (v6 | v6)) {
        if true {
            if (!false) {
                v1 = (v1 / (((v1 * 42)) | 1))
                if true {
                    for fi16 := 0; fi16 < 18; fi16++ {
                        v3 -= (v3 - v3)
                        _ = fi16
                    }
                    for fi17 := 0; fi17 < 15; fi17++ {
                        if ((v3 >> ((v6) & 63)) >= (^(v3))) {
                            var v18 float32 = 0.1
                            if (false || ((float32(-2.25) - v18) == v18)) {
                                if false {
                                    for fi19 := 0; fi19 < 18; fi19++ {
                                        var v20 uint64 = 18446744073709551615
                                        switch "a" {
                                        case "a":
                                            v4 += int32(v3)
                                        case "\n\t":
                                            v21 := GenS1{F0: 345, F1: 0}
                                            _ = v21
                                        case "":
                                            var v22 bool = false
                                            _ = v22
                                        }
                                        v23 := []int{-47, fi19, v2}
                                        _ = fi19
                                        _ = v20
                                        _ = v23
                                    }
                                    v24 := []int{2147483647, (^(v2))}
                                    _ = v24
                                } else {
                                    v25 := GenS1{F0: -213, F1: v6}
                                    v26 := []int{(fi17 - v2), v2}
                                    var v27 int = (v2 % ((int(v7)) | 1))
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                }
                                var v28 uint = v6
                                _ = v28
                            }
                            _ = v18
                        } else {
                            for fi29 := 0; fi29 < 19; fi29++ {
                                v30 := GenS1{F0: -882, F1: v6}
                                v2 += v2
                                _ = fi29
                                _ = v30
                            }
                        }
                        v15 += 42
                        _ = fi17
                    }
                } else {
                    for fi31 := 0; fi31 < 3; fi31++ {
                        v32 := []int{-1, fi31}
                        _ = fi31
                        _ = v32
                    }
                    v33 := GenS1{F0: 42, F1: v6}
                    _ = v33
                }
            }
        } else {
            v34 := GenS1{F0: 124, F1: 1}
            _ = v34
        }
        var v35 uint32 = 362
        _ = v35
    }
    switch (v2 << ((2) & 63)) {
    case 32767:
        v36 := []int{(v2 / ((v2) | 1)), (-(v2)), (v2 % ((815) | 1)), 51}
        _ = v36
    }
    v37 := []int{}
    var v38 int64 = -504
    var v39 uint8 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint }
type GenS2 struct { F0 int8; F1 int16; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 int16; F2 bool }

func main() {
    v1 := make(map[int]int)
    var v2 float32 = ((float32((0.1 / ((((0.1 / 1.5) - 1.5) / (1.5 / 0.1)) - (1.5 + 3.14159)))) / 1e10) - 0.0)
    var v3 int32 = 0
    fmt.Println(v2, v3)
    if ((65535 <= 0) || (v3 < 1)) {
        var v4 uint64 = 114
        _ = v4
    }
    v5 := GenS1{F0: -58, F1: 1, F2: 1}
    var v6 string = "hello"
    var v7 uint16 = 597
    switch ((v6 + (v6 + "foo bar")) + "foo bar") {
    case "a":
        v8 := []int{-2147483648, 9223372036854775807, -9223372036854775808, 42}
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        v3 = 42
        for fi10 := 0; fi10 < 8; fi10++ {
            v6 = ("hello" + "a")
            _ = fi10
        }
        _ = fi9
    }
    if true {
        if ((42 >= 4294967295) && false) {
            if (42 <= 1) {
                v3 -= v3
                var v11 uint = 0
                v3 = v3
                _ = v11
            } else {
                v12 := []int{2147483647}
                v13 := GenS1{F0: 1, F1: 42, F2: 18446744073709551615}
                v14 := GenS1{F0: 1, F1: 87, F2: 973}
                _ = v12
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 15; fi15++ {
                if false {
                    v16 := GenS2{F0: -117, F1: -32768, F2: 1}
                    _ = v16
                }
                var v17 int16 = -1
                _ = fi15
                _ = v17
            }
            if true {
                for fi18 := 0; fi18 < 9; fi18++ {
                    if false {
                        for fi19 := 0; fi19 < 18; fi19++ {
                            var v20 uint32 = 2147483648
                            var v21 uint8 = 31
                            v3 = v3
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                    }
                    var v22 bool = (!true)
                    _ = fi18
                    _ = v22
                }
                v23 := []int{32767}
                if (42 == 4294967295) {
                    if (4294967295 >= 955) {
                        v6 = (v6 + "")
                        var v24 float64 = (0.1 - (0.1 / 1.5))
                        _ = v24
                    } else {
                        fmt.Println(v7, v6, v7)
                        var v25 string = v6
                        if true {
                            for fi26 := 0; fi26 < 2; fi26++ {
                                if len(v23) >= 2 {
                                    v23 = v23[1:len(v23)]
                                }
                                v27 := GenS1{F0: 0, F1: 168, F2: 18446744073709551615}
                                _ = fi26
                                _ = v27
                            }
                        } else {
                            if len(v23) >= 2 {
                                v23 = v23[1:len(v23)]
                            }
                            v2 -= v2
                        }
                        _ = v25
                    }
                    v3 += (v3 - (v3 & (^(v3))))
                    fmt.Println(v7, v6)
                }
                _ = v23
            } else {
                v6 = (((v6 + v6) + "\n\t") + (v6 + v6))
                var v28 float64 = 3.14159
                _ = v28
            }
        } else {
            var v29 float32 = (v2 + v2)
            _ = v29
        }
        var v30 uint8 = 255
        v31 := GenS1{F0: 688, F1: v30, F2: 65535}
        _ = v30
        _ = v31
    } else {
        var v32 bool = (!false)
        fmt.Println(v32, v6)
        v33 := []int{42, -494, 42}
        _ = v32
        _ = v33
    }
    v3 = (v3 % ((1) | 1))
    if true {
        var v34 uint32 = 65535
        fmt.Println(v6)
        _ = v34
    } else {
        v2 -= (0.1 + v2)
        v35 := GenS3{F0: 18446744073709551615}
        _ = v35
    }
    if (335 != 32767) {
        v36 := GenS4{F0: 0, F1: 487, F2: false}
        _ = v36
    } else {
        v1[4294967295] = 28
        v2 = (1e10 + v2)
    }
    v37 := []int{-664}
    v2 += (0.1 + v2)
    v3 = (v3 & v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v5
    _ = v6
    _ = v7
}

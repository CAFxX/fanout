package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string }

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 42
    v3 := make(map[string]int)
    for fi4 := 0; fi4 < 17; fi4++ {
        if ((540 > 199) && true) {
            v5 := GenS1{F0: (v2 + (v2 >> ((908) & 63)))}
            _ = v5
        }
        _ = fi4
    }
    for fi6 := 0; fi6 < 1; fi6++ {
        fmt.Println(v2, fi6)
        var v7 uint16 = 65535
        var v8 float64 = (-2.25 / 1.5)
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint8 = 0
    v10 := GenS2{F0: -1, F1: 0}
    v11 := GenS2{F0: 0, F1: 644}
    for fi12 := 0; fi12 < 19; fi12++ {
        v13 := GenS2{F0: 1, F1: -689}
        _ = fi12
        _ = v13
    }
    if true {
        fmt.Println(v2)
    } else {
        if (!true) {
            if false {
                v14 := []int{0}
                _ = v14
            }
        }
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        v16 := v3["a"]
        fmt.Println(v9, v2)
        var v17 uint32 = 222
        _ = fi15
        _ = v16
        _ = v17
    }
    switch -9223372036854775808 {
    case -2147483648:
        v18 := []int{32768, 16, 127, 4294967295}
        _ = v18
    case 32768:
        var v19 int8 = 52
        _ = v19
    case 32767:
        v20 := []int{128, 196, 127, 492}
        _ = v20
    }
    var v21 int = 541
    var v22 int32 = -314
    var v23 int16 = 854
    switch (v21 ^ v21) {
    case 9223372036854775807:
        if false {
            var v24 uint32 = 2147483648
            var v25 int32 = 2147483647
            _ = v24
            _ = v25
        } else {
            var v26 bool = ((true && ((18446744073709551615 > 0) || true)) && (489 >= v21))
            var v27 bool = (1e10 != float64(-2.25))
            if ((v9 | 99) != (v9 << ((1000) & 63))) {
                var v28 float32 = 3.14159
                _ = v28
            }
            _ = v26
            _ = v27
        }
    case 4294967295:
        switch "foo bar" {
        case "a":
            fmt.Println(v23)
        case "foo bar":
            if false {
                var v29 int64 = 9223372036854775807
                _ = v29
            }
        default:
            if (!false) {
                var v30 uint = 65535
                v3[""] = 2147483647
                _ = v30
            } else {
                switch (^(v21)) {
                case 42:
                    if (false && (((3.14159 + 3.14159) + -2.25) == (1.5 - 1.5))) {
                        v23 = (v23 / ((v23) | 1))
                    } else {
                        v3[("\n\t" + "\n\t")] = (v21 * v21)
                        if false {
                            v21 += 2147483647
                            var v31 int8 = 0
                            _ = v31
                        }
                        if (-9223372036854775808 != -1) {
                            v23 += (v23 - 1)
                            if (v9 <= (^(v9))) {
                                var v32 uint8 = 127
                                _ = v32
                            } else {
                                v9 = v9
                                var v33 int32 = 42
                                v21 += v21
                                _ = v33
                            }
                        } else {
                            if true {
                                v34 := GenS3{F0: 1}
                                for fi35 := 0; fi35 < 7; fi35++ {
                                    v36 := GenS1{F0: (v2 << ((495) & 63))}
                                    _ = fi35
                                    _ = v36
                                }
                                var v37 int = (v21 << ((65535) & 63))
                                _ = v34
                                _ = v37
                            } else {
                                v9 = 255
                                if ((("hello" + "a") + "hello") != "foo bar") {
                                    var v38 int64 = -717
                                    var v39 int8 = -105
                                    _ = v38
                                    _ = v39
                                }
                            }
                            v21 -= v21
                        }
                    }
                }
                fmt.Println(v21)
            }
        }
    }
    v40 := GenS4{F0: ""}
    v41 := GenS3{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v40
    _ = v41
    _ = v9
}

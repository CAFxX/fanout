package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32; F1 uint; F2 int }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int16; F1 int32 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint64 }

func main() {
    var v1 float64 = 1e10
    var v2 uint16 = 42
    var v3 int16 = 32767
    if ((-2.25 / 1.0) >= 0.1) {
        v2 += v2
        var v4 int32 = 42
        _ = v4
    } else {
        fmt.Println(v1, v2, v3)
        if false {
            var v5 int = 699
            fmt.Println(v1, v1, v3)
            var v6 int8 = -20
            _ = v5
            _ = v6
        } else {
            var v7 bool = false
            v2 = v2
            v8 := []int{}
            _ = v7
            _ = v8
        }
        v9 := []int{1, 32768, 1}
        _ = v9
    }
    v3 -= v3
    v10 := GenS1{F0: 127}
    v2 += (v2 / (((v2 / ((v2) | 1))) | 1))
    v11 := GenS2{F0: 95, F1: 2}
    v12 := GenS3{F0: true}
    v13 := GenS2{F0: 1, F1: v2}
    for fi14 := 0; fi14 < 8; fi14++ {
        v15 := []int{(^(fi14)), (-(fi14)), 2147483647, 0}
        _ = fi14
        _ = v15
    }
    v16 := GenS4{F0: v3}
    switch 9223372036854775807 {
    case 32768:
        v17 := GenS3{F0: false}
        _ = v17
    case 2147483647:
        v2 += v2
    case 1:
        switch "a" {
        case "hello":
            var v18 int64 = 9223372036854775807
            _ = v18
        case "":
            var v19 uint16 = 255
            _ = v19
        case "\n\t":
            var v20 uint = 65535
            _ = v20
        default:
            v21 := GenS3{F0: (v3 > v3)}
            _ = v21
        }
    default:
        v3 = (v3 + (^(v3)))
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        if (4294967295 != 65535) {
            if true {
                v2 += (v2 % ((v2) | 1))
                if (42 != 610) {
                    var v23 uint = 65535
                    if (false || ((v2 << ((255) & 63)) == (v2 % ((v2) | 1)))) {
                        v24 := []int{127, (fi22 / (((fi22 * -1)) | 1)), -341, (fi22 & fi22)}
                        var v25 uint32 = 2147483647
                        _ = v24
                        _ = v25
                    } else {
                        var v26 uint = 65535
                        if true {
                            var v27 uint32 = 2147483647
                            _ = v27
                        } else {
                            if false {
                                fmt.Println(v2)
                            } else {
                                v3 = (v3 ^ 1)
                                var v28 bool = ((72 >= v3) || true)
                                _ = v28
                            }
                            var v29 float32 = -2.25
                            var v30 bool = (!((v1 / v1) >= v1))
                            _ = v29
                            _ = v30
                        }
                        _ = v26
                    }
                    _ = v23
                }
                if (501 >= 42) {
                    v31 := GenS5{F0: 2147483647, F1: 1, F2: (fi22 + fi22)}
                    v32 := GenS6{F0: 42}
                    fmt.Println(v1, v2, v1)
                    _ = v31
                    _ = v32
                }
            } else {
                v33 := GenS7{F0: 0}
                fmt.Println(fi22, v1)
                v34 := GenS8{F0: 0}
                _ = v33
                _ = v34
            }
            v35 := GenS9{F0: (v3 / ((32767) | 1)), F1: 42}
            for fi36 := 0; fi36 < 5; fi36++ {
                if (!("" == "a")) {
                    fmt.Println(fi22, v1)
                    if true {
                        var v37 int = 387
                        _ = v37
                    }
                    switch (fi36 % ((0) | 1)) {
                    case -2147483648:
                        v38 := GenS10{F0: 255}
                        _ = v38
                    }
                } else {
                    var v39 uint32 = 2
                    if true {
                        fmt.Println(v1, v3)
                        v40 := []int{}
                        _ = v40
                    } else {
                        v41 := GenS1{F0: 127}
                        _ = v41
                    }
                    v42 := GenS11{F0: 1}
                    _ = v39
                    _ = v42
                }
                var v43 uint32 = 1
                var v44 bool = (true || true)
                _ = fi36
                _ = v43
                _ = v44
            }
            _ = v35
        } else {
            if (905 >= 1) {
                v2 = 2
            }
            var v45 float64 = (v1 / (float64(3.14159) + v1))
            v46 := []int{fi22, (fi22 + 1), fi22}
            _ = v45
            _ = v46
        }
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
}

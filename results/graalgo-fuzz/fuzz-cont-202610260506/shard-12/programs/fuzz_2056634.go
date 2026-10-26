package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 bool }
type GenS2 struct { F0 int; F1 string; F2 uint }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32; F1 string }
func genFn0(p0 uint16, p1 uint) int32 {
    v5 := []int{-868, 450, -2147483648, 128}
    switch ("" + "") {
    case "hello":
        if ("hello" > "\n\t") {
            var v6 string = "hello"
            var v7 uint8 = 255
            _ = v6
            _ = v7
        } else {
            _ = len(v5)
            if false {
                var v8 uint64 = 966
                _ = v8
            } else {
                p1 += 65535
                if (!false) {
                    p0 = 255
                    v9 := GenS1{F0: false, F1: 1, F2: (!true)}
                    _ = v9
                } else {
                    var v10 uint16 = p0
                    v10 = 255
                    switch -9223372036854775808 {
                    case -882:
                        if true {
                            var v11 int16 = 0
                            if true {
                                var v12 uint8 = 128
                                _ = v12
                            } else {
                                var v13 bool = (-2147483648 == 42)
                                _ = v13
                            }
                            _ = v11
                        } else {
                            v10 = 694
                        }
                    }
                    _ = v10
                }
            }
        }
    case "":
        var v14 uint32 = 2
        _ = v14
    case "\n\t":
        if len(v5) > 0 {
            v5[0] = 128
        }
    default:
        p0 = (^(p0))
    }
    _ = p0
    _ = p1
    _ = v5
    return 1
}
func genFn1(p0 uint8, p1 uint8) uint8 {
    fmt.Println(p0, p1)
    if (p0 == 128) {
        var v15 string = ""
        _ = v15
    }
    var v16 uint = 4294967295
    _ = p0
    _ = p1
    _ = v16
    return p1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]string)
    var v3 string = "hello"
    var v4 bool = ((1 > 0) || false)
    v17 := GenS1{F0: true, F1: -896, F2: (-9223372036854775808 >= -821)}
    delete(v2, 443)
    v18 := GenS2{F0: -778, F1: "a", F2: 18446744073709551615}
    switch 948 {
    case 9223372036854775807:
        for fi19 := 0; fi19 < 10; fi19++ {
            for fi20 := 0; fi20 < 5; fi20++ {
                if false {
                    fmt.Println(fi20)
                    fmt.Println(v3, v3, fi20)
                }
                _ = fi20
            }
            v21 := GenS1{F0: v4, F1: 128, F2: false}
            _ = fi19
            _ = v21
        }
    case 579:
        switch 9223372036854775807 {
        case 756:
            var v22 int64 = -9223372036854775808
            _ = v22
        case 1:
            var v23 uint16 = 970
            _ = v23
        case -703:
            v3 = "\n\t"
        }
    }
    fmt.Println(v3)
    v24 := GenS2{F0: -322, F1: v3, F2: 2}
    var v25 int8 = 6
    if v4 {
        if true {
            var v26 int = 127
            if ((v25 | v25) != (v25 | 1)) {
                v27 := GenS2{F0: 9223372036854775807, F1: "a", F2: 4294967295}
                _ = v27
            } else {
                v28 := []int{(v26 + (v26 % ((724) | 1))), (v26 * v26), v26}
                _ = v28
            }
            _ = v26
        }
    } else {
        var v29 uint = 1
        _ = v29
    }
    var v30 float64 = float64(float32(1.5))
    var v31 bool = (!(275 > 2))
    v1[2147483648] = ("" + ("hello" + "a"))
    switch v3 {
    case "foo bar":
        if v4 {
            v32 := GenS3{F0: v25, F1: 401}
            _ = v32
        } else {
            fmt.Println(v3)
        }
    case "hello":
        var v33 int64 = 0
        _ = v33
    }
    v34 := GenS4{F0: 952}
    fmt.Println(v30)
    var v35 int16 = 1
    var v36 uint16 = 65535
    if (v35 == (v35 - (v35 / ((v35) | 1)))) {
        switch (v3 + "a") {
        case "hello":
            v4 = (2 >= 0)
        case "":
            for fi37 := 0; fi37 < 11; fi37++ {
                v4 = v4
                if (!((v4 || v31) || v31)) {
                    switch v3 {
                    case "a":
                        v35 = (v35 & -32768)
                    case "foo bar":
                        v25 = v25
                    }
                    var v38 float64 = (((float64(1e10) / v30) + v30) / (v30 / v30))
                    v39 := GenS5{F0: -613, F1: "a"}
                    _ = v38
                    _ = v39
                } else {
                    v25 = (v25 << ((255) & 63))
                    fmt.Println(v25)
                    var v40 float32 = 1e10
                    _ = v40
                }
                var v41 uint8 = 54
                _ = fi37
                _ = v41
            }
        default:
            var v42 int = 4294967295
            _ = v42
        }
        var v43 bool = true
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}

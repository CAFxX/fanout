package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 bool }

func main() {
    var v1 int = -9223372036854775808
    v2 := make(map[int]bool)
    var v3 float32 = 3.14159
    var v4 int = (v1 & (v1 ^ (v1 & v1)))
    var v5 uint8 = 220
    var v6 uint8 = (v5 | 0)
    v7 := v2[v4]
    v8 := GenS1{F0: ("" + "hello")}
    switch ("hello" + "\n\t") {
    case "":
        v5 += v5
    case "a":
        v9 := GenS1{F0: "a"}
        _ = v9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        switch (v1 ^ (v4 * fi10)) {
        case 4294967295:
            if v7 {
                v11 := GenS1{F0: ("" + "hello")}
                _ = v11
            }
        case 42:
            v6 -= (v5 & 197)
        default:
            fmt.Println(v5, v1)
        }
        _ = fi10
    }
    switch ("" + "a") {
    case "":
        switch 0 {
        case -140:
            switch v4 {
            case 1:
                v6 += 175
            }
        case 9223372036854775807:
            switch (^(v4)) {
            case 255:
                for fi12 := 0; fi12 < 7; fi12++ {
                    for fi13 := 0; fi13 < 3; fi13++ {
                        v14 := []int{(-(fi12)), (v4 % (((-(v1))) | 1))}
                        _ = fi13
                        _ = v14
                    }
                    fmt.Println(v1, fi12, v3)
                    _ = fi12
                }
            case -2147483648:
                v5 -= v6
            case 228:
                if (!(100 < 92)) {
                    fmt.Println(v7)
                } else {
                    v5 = (v6 - v6)
                }
            default:
                v15 := GenS1{F0: "\n\t"}
                _ = v15
            }
        }
    default:
        v2[v4] = (!(false && v7))
    }
    v1 = v4
    v16 := GenS1{F0: ("hello" + "hello")}
    v4 = (v1 + v1)
    if false {
        var v17 int32 = 380
        for fi18 := 0; fi18 < 18; fi18++ {
            if ((v5 * 128) > (v6 / ((v6) | 1))) {
                v7 = (-121 == 63)
                v19 := GenS1{F0: "foo bar"}
                if v7 {
                    if v7 {
                        switch "a" {
                        case "foo bar":
                            fmt.Println(fi18, v7)
                        case "hello":
                            if ((((float64((float32(float64(((v3 * float32((-2.25 * 3.14159))) / 3.14159))) * 0.0)) / (((1e10 + 1e10) - 1.5) + 1e10)) + 1e10) + ((-2.25 + 3.14159) - 1.5)) > -2.25) {
                                v20 := GenS1{F0: "foo bar"}
                                _ = v20
                            } else {
                                var v21 int32 = (v17 << ((uint(v6)) & 63))
                                _ = v21
                            }
                        case "":
                            v22 := GenS1{F0: ("\n\t" + ("a" + "a"))}
                            _ = v22
                        default:
                            fmt.Println(v3, v6)
                        }
                        fmt.Println(v17)
                    } else {
                        v3 = 3.14159
                    }
                } else {
                    v23 := []int{v4, (fi18 >> ((18446744073709551615) & 63)), 4294967295}
                    _ = v23
                }
                _ = v19
            } else {
                fmt.Println(v6)
                if (v7 || (-32768 > 1000)) {
                    v24 := GenS1{F0: "hello"}
                    v25 := GenS2{F0: "a", F1: v7}
                    _ = v24
                    _ = v25
                } else {
                    v26 := GenS2{F0: "foo bar", F1: ((0.1 + -2.25) >= (0.0 - 3.14159))}
                    v27 := []int{(^(v4)), (fi18 | v4)}
                    _ = v26
                    _ = v27
                }
                if ((32767 < -126) && v7) {
                    var v28 bool = v7
                    for fi29 := 0; fi29 < 1; fi29++ {
                        v3 = 1e10
                        var v30 int = (-(v4))
                        v28 = v7
                        fmt.Println(fi29, fi18)
                        _ = fi29
                        _ = v30
                    }
                    if v28 {
                        for fi31 := 0; fi31 < 7; fi31++ {
                            var v32 int32 = (-(v17))
                            v33 := GenS1{F0: "foo bar"}
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                        switch ("hello" + "a") {
                        case "a":
                            fmt.Println(v6)
                        default:
                            var v34 string = "foo bar"
                            _ = v34
                        }
                    } else {
                        var v35 uint16 = 42
                        _ = v35
                    }
                    _ = v28
                }
            }
            v36 := []int{(v1 & (fi18 << ((0) & 63)))}
            if v7 {
                fmt.Println(v17, v6, v4)
            }
            _ = fi18
            _ = v36
        }
        _ = v17
    } else {
        fmt.Println(v1, v5, v1)
        v37 := []int{v4}
        _ = v37
    }
    fmt.Println(v4, v4, v3)
    v6 -= (v5 ^ (v5 ^ (v6 ^ v5)))
    v7 = v7
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

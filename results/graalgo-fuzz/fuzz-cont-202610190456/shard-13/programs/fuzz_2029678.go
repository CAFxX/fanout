package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
func genFn0(p0 uint8) int {
    v5 := GenS1{F0: 18446744073709551615, F1: -852}
    p0 += p0
    v6 := GenS1{F0: 4294967295, F1: 901}
    v7 := []int{-9223372036854775808, 1}
    var v8 uint8 = 255
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return -103
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = -40
    v3 := make(map[string]bool)
    var v4 bool = (!(0.0 >= 0.1))
    fmt.Println(v2, v4)
    v4 = false
    if v4 {
        v9 := []int{2147483647, genFn0(255), genFn0(127)}
        _ = v9
    } else {
        v2 = v2
        fmt.Println(v4, v4)
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        v2 -= 28
        var v11 uint32 = 42
        var v12 uint8 = 167
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println(v2)
    fmt.Println(v2, v4)
    for fi13 := 0; fi13 < 19; fi13++ {
        for fi14 := 0; fi14 < 20; fi14++ {
            v15 := []int{fi14, 4294967295}
            _ = fi14
            _ = v15
        }
        v16 := []int{fi13, genFn0(42), genFn0(1)}
        _ = fi13
        _ = v16
    }
    var v17 uint = 1
    switch -2147483648 {
    case 42:
        var v18 uint64 = 717
        _ = v18
    case -161:
        var v19 uint32 = 1
        _ = v19
    case -497:
        fmt.Println(v4)
    default:
        for fi20 := 0; fi20 < 9; fi20++ {
            switch (^(fi20)) {
            case 127:
                for fi21 := 0; fi21 < 9; fi21++ {
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(v2, fi20, v2)
                    case "":
                        if (!true) {
                            v22 := GenS1{F0: (v17 ^ v17), F1: fi21}
                            v23 := GenS1{F0: 4294967295, F1: (fi21 ^ (fi20 & (fi21 & fi20)))}
                            v17 -= (v17 - (v17 >> ((294) & 63)))
                            fmt.Println(fi21)
                            _ = v22
                            _ = v23
                        } else {
                            if (!(!(545 == 0))) {
                                var v24 string = "a"
                                _ = v24
                            } else {
                                v25 := GenS1{F0: (v17 ^ (v17 % ((v17) | 1))), F1: -9223372036854775808}
                                _ = v25
                            }
                            var v26 uint8 = 2
                            if true {
                                var v27 int64 = 42
                                _ = v27
                            } else {
                                switch fi20 {
                                case 32768:
                                    fmt.Println(v17)
                                default:
                                    v2 += v2
                                }
                            }
                            switch -302 {
                            case 128:
                                fmt.Println(fi21)
                            case -9223372036854775808:
                                fmt.Println(fi20, v4, fi20)
                            case -2147483648:
                                fmt.Println(v4, fi21, fi21)
                            }
                            _ = v26
                        }
                    case "hello":
                        switch "a" {
                        case "hello":
                            switch "hello" {
                            case "":
                                v2 = v2
                            case "foo bar":
                                var v28 int8 = (v2 - -47)
                                v17 = (v17 - 42)
                                _ = v28
                            case "\n\t":
                                delete(v3, "\n\t")
                            }
                        case "\n\t":
                            var v29 int32 = 2147483647
                            _ = v29
                        case "":
                            fmt.Println(v17, fi21)
                        default:
                            fmt.Println(fi21, v17)
                        }
                    }
                    switch fi21 {
                    case 4294967295:
                        v4 = (-708 <= -401)
                    default:
                        v17 += (^(v17))
                    }
                    _ = fi21
                }
            case -52:
                v4 = (v4 && v4)
            case -396:
                v30 := []int{(fi20 % (((fi20 + (fi20 / ((-453) | 1)))) | 1)), (fi20 << ((172) & 63))}
                _ = v30
            }
            delete(v3, "hello")
            for fi31 := 0; fi31 < 3; fi31++ {
                var v32 bool = true
                _ = fi31
                _ = v32
            }
            _ = fi20
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}

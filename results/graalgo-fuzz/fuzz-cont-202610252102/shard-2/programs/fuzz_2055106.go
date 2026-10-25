package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 uint; F1 string; F2 uint64 }
type GenS5 struct { F0 uint8; F1 int64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 2
    v3 := make(map[int]bool)
    v4 := GenS1{F0: 42, F1: (18446744073709551615 > 4294967295)}
    v5 := v1[556]
    v6 := GenS2{F0: (v2 | v2)}
    if v5 {
        switch -1 {
        case 128:
            v2 = v2
        case -9223372036854775808:
            switch "\n\t" {
            case "\n\t":
                v7 := GenS1{F0: 66, F1: v5}
                _ = v7
            case "a":
                v8 := GenS1{F0: 4294967295, F1: (v5 || v5)}
                _ = v8
            default:
                switch ("hello" + ((("\n\t" + "hello") + "foo bar") + "")) {
                case "":
                    v5 = v5
                case "\n\t":
                    var v9 float32 = float32(0.1)
                    _ = v9
                case "a":
                    for fi10 := 0; fi10 < 16; fi10++ {
                        fmt.Println(v5, v5)
                        v11 := []int{fi10, fi10}
                        _ = fi10
                        _ = v11
                    }
                default:
                    v3[-50] = false
                }
            }
        }
    } else {
        v12 := GenS3{F0: 218, F1: 65535}
        v2 += (v2 * v2)
        v13 := []int{127, -599, 2147483647}
        _ = v12
        _ = v13
    }
    if ((v2 ^ v2) != (^(v2))) {
        if false {
            switch 4294967295 {
            case 32767:
                for fi14 := 0; fi14 < 16; fi14++ {
                    v15 := GenS2{F0: (^(v2))}
                    switch "" {
                    case "hello":
                        fmt.Println(v5, fi14, v5)
                    default:
                        var v16 string = ""
                        _ = v16
                    }
                    switch "" {
                    case "a":
                        v17 := GenS4{F0: 4294967295, F1: "hello", F2: 1}
                        _ = v17
                    }
                    _ = fi14
                    _ = v15
                }
            case 9223372036854775807:
                if (930 <= 1000) {
                    var v18 uint32 = 624
                    for fi19 := 0; fi19 < 5; fi19++ {
                        switch -2147483648 {
                        case 278:
                            if (-1 < -1) {
                                fmt.Println(fi19, v2, v18)
                                v20 := v1[2]
                                v21 := []int{}
                                _ = v20
                                _ = v21
                            } else {
                                var v22 string = "\n\t"
                                _ = v22
                            }
                        case -532:
                            if v5 {
                                v18 -= 700
                                for fi23 := 0; fi23 < 20; fi23++ {
                                    v24 := GenS5{F0: v2, F1: 452}
                                    v5 = v5
                                    var v25 int = 69
                                    _ = fi23
                                    _ = v24
                                    _ = v25
                                }
                            }
                        case 42:
                            var v26 string = ""
                            _ = v26
                        }
                        v2 -= (v2 + 128)
                        v27 := GenS3{F0: 255, F1: 18446744073709551615}
                        _ = fi19
                        _ = v27
                    }
                    _ = v18
                } else {
                    if true {
                        v28 := []int{1, -2147483648}
                        var v29 bool = (!v5)
                        _ = len(v28)
                        _ = v28
                        _ = v29
                    } else {
                        v2 = (v2 + 11)
                    }
                }
            case 128:
                delete(v3, 42)
            }
            var v30 int32 = 0
            _ = v30
        } else {
            var v31 uint = 4294967295
            _ = v31
        }
    }
    switch "\n\t" {
    case "foo bar":
        for fi32 := 0; fi32 < 6; fi32++ {
            switch "a" {
            case "a":
                v2 -= v2
            case "hello":
                var v33 bool = v5
                _ = v33
            case "foo bar":
                v34 := GenS2{F0: 62}
                _ = v34
            }
            fmt.Println(v2)
            v35 := []int{}
            _ = fi32
            _ = v35
        }
    default:
        switch 9223372036854775807 {
        case 2147483647:
            if v5 {
                v36 := GenS5{F0: (^(v2)), F1: -280}
                for fi37 := 0; fi37 < 15; fi37++ {
                    fmt.Println(v2, v5)
                    _ = fi37
                }
                v38 := GenS3{F0: 1, F1: 2}
                _ = v36
                _ = v38
            }
        default:
            v39 := GenS5{F0: (^(v2)), F1: -9223372036854775808}
            _ = v39
        }
    }
    var v40 int8 = 16
    v41 := GenS1{F0: 42, F1: v5}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int64 }
type GenS2 struct { F0 bool; F1 string; F2 bool }
type GenS3 struct { F0 uint8 }
func genFn0() uint {
    var v6 uint32 = 123
    v6 -= (^(v6))
    v6 = (v6 % (((v6 >> ((1) & 63))) | 1))
    var v7 int8 = -1
    _ = v6
    _ = v7
    return 697
}

func main() {
    var v1 uint16 = 0
    var v2 int16 = 1
    var v3 int32 = 0
    var v4 uint16 = 0
    v5 := make(map[int]bool)
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 5; fi9++ {
            switch fi9 {
            case -2147483648:
                v4 += (^(v4))
            case 9223372036854775807:
                var v10 int32 = -2147483648
                _ = v10
            case 310:
                v4 = v1
            }
            switch ("foo bar" + "") {
            case "hello":
                fmt.Println(v4, v4)
            case "\n\t":
                if false {
                    v1 = v1
                } else {
                    v11 := []int{}
                    _ = v11
                }
            case "a":
                v2 = v2
            default:
                v1 = v4
            }
            _ = fi9
        }
        _ = fi8
    }
    v12 := GenS1{F0: (v3 % (((v3 / ((v3) | 1))) | 1)), F1: -128, F2: 9223372036854775807}
    fmt.Println(v4, v1, v4)
    v13 := GenS2{F0: false, F1: "a", F2: true}
    fmt.Println(v3)
    if (-1 <= 351) {
        v14 := GenS3{F0: 102}
        fmt.Println(v3, v4)
        fmt.Println(v3, v1)
        _ = v14
    }
    if true {
        v15 := GenS3{F0: 128}
        switch -9223372036854775808 {
        case 127:
            v4 += (v4 >> ((133) & 63))
        case 2147483647:
            var v16 int16 = 448
            _ = v16
        case -9223372036854775808:
            var v17 float32 = -2.25
            _ = v17
        }
        _ = v15
    } else {
        for fi18 := 0; fi18 < 17; fi18++ {
            switch -2147483648 {
            case -9223372036854775808:
                switch 0 {
                case 83:
                    for fi19 := 0; fi19 < 11; fi19++ {
                        switch "" {
                        case "foo bar":
                            v20 := GenS2{F0: false, F1: "a", F2: false}
                            _ = v20
                        }
                        v21 := []int{fi19, fi19, (fi18 * fi19)}
                        _ = fi19
                        _ = v21
                    }
                }
            case -88:
                if (1 >= 0) {
                    for fi22 := 0; fi22 < 13; fi22++ {
                        v23 := GenS3{F0: 0}
                        var v24 uint32 = 981
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 3; fi25++ {
                        v26 := GenS2{F0: true, F1: "hello", F2: false}
                        switch ("foo bar" + "") {
                        case "a":
                            var v27 uint16 = (v1 ^ v1)
                            _ = v27
                        }
                        fmt.Println(fi25, v2)
                        _ = fi25
                        _ = v26
                    }
                    v28 := GenS2{F0: (v3 == 2147483647), F1: ("a" + "\n\t"), F2: true}
                    _ = v28
                }
            default:
                v3 -= (v3 - 2147483647)
            }
            _ = fi18
        }
        v3 += (v3 >> ((42) & 63))
        v1 = 695
    }
    var v29 int = -9223372036854775808
    var v30 uint = genFn0()
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}

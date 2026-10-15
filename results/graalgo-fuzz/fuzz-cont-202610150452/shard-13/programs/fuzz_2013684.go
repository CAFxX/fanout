package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 uint16; F2 string }
func genFn0() int8 {
    var v4 int32 = 42
    if (!false) {
        v5 := []int{2147483647, 32767}
        _ = v5
    } else {
        v6 := []int{42, 0, 32767}
        v7 := GenS1{F0: 127, F1: -478, F2: true}
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: 127, F1: 742, F2: (((false || (!(42 <= 1))) || true) || true)}
    _ = v4
    _ = v8
    return -1
}
func genFn1() int {
    switch "a" {
    case "":
        var v9 int64 = 1
        _ = v9
    default:
        var v10 int64 = -655
        _ = v10
    }
    var v11 int8 = 0
    var v12 uint8 = 192
    v12 = v12
    for fi13 := 0; fi13 < 6; fi13++ {
        if (v11 != genFn0()) {
            v14 := []int{(fi13 << ((255) & 63))}
            fmt.Println(v12, v11, fi13)
            v12 += (v12 & 127)
            _ = v14
        } else {
            v15 := GenS1{F0: (v11 + 100), F1: 9223372036854775807, F2: false}
            fmt.Println(v12, v11, v11)
            for fi16 := 0; fi16 < 5; fi16++ {
                v11 += (v11 + v11)
                fmt.Println(v12, v11, v11)
                _ = fi16
            }
            _ = v15
        }
        for fi17 := 0; fi17 < 18; fi17++ {
            v18 := GenS2{F0: 18446744073709551615}
            v11 = (v11 % ((31) | 1))
            var v19 string = "a"
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi13
    }
    _ = v11
    _ = v12
    return 32768
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    var v3 uint = 240
    for fi20 := 0; fi20 < 15; fi20++ {
        switch "\n\t" {
        case "":
            for fi21 := 0; fi21 < 12; fi21++ {
                v22 := GenS3{F0: -122, F1: 1, F2: "a"}
                v3 += v3
                delete(v1, (fi20 / ((849) | 1)))
                _ = fi21
                _ = v22
            }
        case "foo bar":
            var v23 uint = (v3 | v3)
            _ = v23
        default:
            if (2147483647 != 42) {
                v24 := []int{fi20, -9223372036854775808, -957}
                var v25 uint16 = 65535
                _ = v24
                _ = v25
            }
        }
        _ = fi20
    }
    var v26 uint64 = 112
    v27 := v2[0]
    var v28 uint32 = 2147483647
    v29 := GenS3{F0: 1, F1: 32767, F2: ("a" + "foo bar")}
    for fi30 := 0; fi30 < 8; fi30++ {
        if true {
            if false {
                switch "hello" {
                case "hello":
                    switch -2147483648 {
                    case 733:
                        if false {
                            v3 = 4294967295
                        }
                    case 83:
                        v31 := GenS1{F0: 45, F1: 0, F2: (false || (0 > -99))}
                        _ = v31
                    }
                }
                v3 = v3
            }
        }
        if true {
            var v32 string = "\n\t"
            for fi33 := 0; fi33 < 18; fi33++ {
                v28 += v28
                _ = fi33
            }
            _ = v32
        }
        v27 = genFn1()
        _ = fi30
    }
    var v34 uint32 = v28
    if false {
        v27 -= v27
        fmt.Println(v26, v26, v27)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v34
}

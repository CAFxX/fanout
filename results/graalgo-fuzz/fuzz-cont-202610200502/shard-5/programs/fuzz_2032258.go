package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int16; F1 int16; F2 int32 }
type GenS3 struct { F0 int8; F1 int64 }
func genFn0(p0 uint64, p1 uint16) bool {
    var v7 int8 = 42
    fmt.Println(v7, p0)
    var v8 int8 = v7
    v9 := GenS1{F0: 669, F1: 1}
    if ((^(p1)) == p1) {
        v7 -= v7
        switch -145 {
        case 32768:
            v10 := GenS1{F0: 18446744073709551615, F1: 2}
            _ = v10
        default:
            fmt.Println(v7)
        }
        switch 1 {
        case 9223372036854775807:
            var v11 string = "a"
            _ = v11
        }
    } else {
        var v12 float64 = ((0.1 * -2.25) - 1e10)
        p1 = (p1 * (p1 ^ (^(p1))))
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return true
}
func genFn1() int32 {
    v13 := GenS1{F0: 878, F1: 390}
    v14 := []int{385, 128, -1}
    var v15 bool = false
    switch "hello" {
    case "":
        v15 = v15
    }
    _ = v13
    _ = v14
    _ = v15
    return 2147483647
}

func main() {
    var v1 int8 = 44
    var v2 int = 128
    v3 := make(map[string]uint32)
    var v4 uint64 = 339
    v5 := make(map[string]bool)
    v6 := make(map[int]string)
    v16 := v6[v2]
    v17 := GenS1{F0: 2, F1: 686}
    v18 := GenS1{F0: uint(v4), F1: 1}
    switch -2147483648 {
    case 2147483647:
        if (4294967295 >= 42) {
            var v19 int = v2
            _ = v19
        } else {
            switch "a" {
            case "":
                v20 := GenS2{F0: 92, F1: 283, F2: int32(v2)}
                _ = v20
            case "\n\t":
                if false {
                    if false {
                        v16 = (("foo bar" + ((("\n\t" + ("" + v16)) + (v16 + v16)) + ((((v16 + (v16 + v16)) + ("a" + v16)) + (v16 + v16)) + v16))) + (v16 + ("hello" + v16)))
                        var v21 bool = false
                        v22 := GenS1{F0: 698, F1: 1}
                        _ = v21
                        _ = v22
                    }
                } else {
                    var v23 uint32 = 2
                    _ = v23
                }
            }
            if (2147483647 != (v2 << ((18446744073709551615) & 63))) {
                v24 := v3[v16]
                v4 = 4294967295
                v4 = (^(v4))
                _ = v24
            } else {
                for fi25 := 0; fi25 < 2; fi25++ {
                    if (1e10 >= 0.1) {
                        switch ((("" + v16) + "a") + v16) {
                        case "\n\t":
                            v2 -= 0
                        case "a":
                            if (v16 != "a") {
                                _ = len(v3)
                            }
                        case "foo bar":
                            fmt.Println(v1, v1)
                        }
                        v4 = 1
                    }
                    v16 = (("a" + v16) + v16)
                    _ = fi25
                }
            }
            _ = len(v6)
        }
    case -9223372036854775808:
        v1 = (v1 & (v1 & v1))
    }
    var v26 int32 = 2147483647
    v26 -= (v26 | (v26 << ((902) & 63)))
    if (1 <= 752) {
        var v27 uint32 = 1
        _ = v27
    }
    v28 := GenS3{F0: v1, F1: 46}
    fmt.Println(v4)
    if true {
        v26 += (v26 << ((42) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

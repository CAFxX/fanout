package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint; F1 uint; F2 uint16 }
func genFn0() int {
    var v4 bool = false
    v4 = v4
    v4 = v4
    fmt.Println(v4)
    _ = v4
    return 128
}

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = -1
    var v3 int8 = 98
    var v5 bool = (!true)
    if (!(!v5)) {
        var v6 uint64 = 0
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 int64 = 0
            for fi9 := 0; fi9 < 19; fi9++ {
                v10 := GenS1{F0: 42, F1: v6}
                _ = fi9
                _ = v10
            }
            _ = fi7
            _ = v8
        }
        _ = v6
    } else {
        v11 := GenS1{F0: 42, F1: 224}
        v3 = (v3 / (((v3 * v3)) | 1))
        _ = v11
    }
    v5 = (((!v5) && v5) || (v5 && v5))
    if false {
        v2 = (v2 & v2)
        if v5 {
            if v5 {
                if ((("foo bar" + "foo bar") + "a") >= ("hello" + ("a" + "\n\t"))) {
                    v12 := GenS1{F0: 0, F1: 846}
                    var v13 string = "hello"
                    var v14 uint8 = 26
                    _ = v12
                    _ = v13
                    _ = v14
                }
            } else {
                if (709 >= 18446744073709551615) {
                    if v5 {
                        fmt.Println(v3, v5, v2)
                        var v15 uint16 = 1
                        fmt.Println(v15)
                        _ = v15
                    } else {
                        var v16 bool = v5
                        _ = v16
                    }
                }
                v17 := GenS2{F0: v2, F1: -9223372036854775808}
                for fi18 := 0; fi18 < 13; fi18++ {
                    v19 := []int{int(v2), fi18, genFn0(), (fi18 | (fi18 ^ genFn0()))}
                    _ = fi18
                    _ = v19
                }
                _ = v17
            }
            var v20 float32 = -2.25
            _ = v20
        } else {
            for fi21 := 0; fi21 < 19; fi21++ {
                fmt.Println(v3, v2, v3)
                v22 := GenS1{F0: -2147483648, F1: 42}
                _ = fi21
                _ = v22
            }
            switch "foo bar" {
            case "foo bar":
                switch -9223372036854775808 {
                case -171:
                    v2 = v3
                case -1:
                    v5 = v5
                case 128:
                    v1[42] = (!(-2.25 != (0.1 * 1e10)))
                }
            case "":
                for fi23 := 0; fi23 < 14; fi23++ {
                    fmt.Println(v2)
                    fmt.Println(v3, v5, v2)
                    switch 9223372036854775807 {
                    case 2147483647:
                        if ("hello" == "") {
                            var v24 int8 = (v3 >> ((18446744073709551615) & 63))
                            v25 := GenS2{F0: (v3 & (v2 ^ v3)), F1: 9223372036854775807}
                            _ = v24
                            _ = v25
                        }
                    }
                    _ = fi23
                }
            case "hello":
                v26 := GenS2{F0: (-(v3)), F1: 1}
                _ = v26
            default:
                v27 := GenS2{F0: v2, F1: 9223372036854775807}
                _ = v27
            }
            var v28 float32 = (3.14159 / 1e10)
            _ = v28
        }
    } else {
        v29 := GenS1{F0: 1, F1: 4294967295}
        _ = v29
    }
    v30 := GenS3{F0: 0, F1: 426, F2: 50}
    v2 -= v3
    v2 -= v3
    var v31 uint8 = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v5
}

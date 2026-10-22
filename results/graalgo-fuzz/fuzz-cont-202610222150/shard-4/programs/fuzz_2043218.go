package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 uint; F2 int }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 uint; F1 int8 }
func genFn0(p0 uint32) uint8 {
    switch 32767 {
    case 32768:
        p0 = 780
    default:
        for fi6 := 0; fi6 < 3; fi6++ {
            fmt.Println(p0, fi6)
            p0 = 2147483648
            _ = fi6
        }
    }
    v7 := GenS1{F0: 2147483647}
    var v8 uint = 0
    fmt.Println(v8)
    _ = p0
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]bool)
    var v3 int16 = 0
    var v4 int64 = 42
    v5 := make(map[string]int)
    switch "\n\t" {
    case "\n\t":
        v9 := GenS1{F0: 0}
        _ = v9
    default:
        var v10 int16 = (v3 * 32767)
        _ = v10
    }
    v4 -= (v4 >> ((1) & 63))
    var v11 float64 = ((((1e10 * 0.0) * 0.1) - 0.1) / ((-2.25 - 3.14159) / float64(3.14159)))
    v11 = v11
    v12 := GenS1{F0: 0}
    var v13 uint8 = 0
    var v14 int16 = v3
    var v15 bool = true
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 float32 = 0.1
        v18 := GenS2{F0: (fi16 | 4294967295), F1: uint(v13), F2: int(v13)}
        _ = fi16
        _ = v17
        _ = v18
    }
    if ((!v15) && v15) {
        if (("a" <= ("hello" + "\n\t")) && v15) {
            for fi19 := 0; fi19 < 12; fi19++ {
                fmt.Println(v15, v3, v13)
                var v20 float32 = -2.25
                if false {
                    for fi21 := 0; fi21 < 11; fi21++ {
                        v22 := []int{fi21, fi21, (fi21 / (((fi19 & (-(fi19)))) | 1))}
                        _ = fi21
                        _ = v22
                    }
                    var v23 int = fi19
                    _ = v23
                } else {
                    fmt.Println(v20)
                }
                _ = fi19
                _ = v20
            }
            if (1 <= 456) {
                switch 128 {
                case 32767:
                    switch 246 {
                    case 4294967295:
                        v24 := GenS2{F0: 42, F1: 2, F2: 2147483647}
                        _ = v24
                    }
                case -2147483648:
                    if v15 {
                        v3 -= v3
                        v25 := GenS3{F0: 2, F1: "a"}
                        var v26 string = ""
                        _ = v25
                        _ = v26
                    }
                }
                v11 += -2.25
            }
            switch -1 {
            case 42:
                delete(v5, "hello")
            case -9223372036854775808:
                v27 := GenS3{F0: 255, F1: "hello"}
                _ = v27
            default:
                v13 = genFn0(989)
            }
        } else {
            if true {
                v11 -= (v11 + 3.14159)
                v28 := GenS2{F0: -1, F1: 658, F2: -2147483648}
                v29 := GenS4{F0: 18446744073709551615, F1: 42}
                _ = v28
                _ = v29
            } else {
                v4 -= (v4 * v4)
                _ = len(v2)
            }
            for fi30 := 0; fi30 < 15; fi30++ {
                var v31 int16 = 1000
                _ = fi30
                _ = v31
            }
        }
        for fi32 := 0; fi32 < 9; fi32++ {
            v33 := GenS3{F0: 42, F1: "\n\t"}
            _ = fi32
            _ = v33
        }
    } else {
        if (!v15) {
            if (!((v3 << ((18446744073709551615) & 63)) != (v3 % ((v3) | 1)))) {
                v34 := GenS2{F0: int(v14), F1: 18446744073709551615, F2: 428}
                var v35 int16 = (v3 - v3)
                for fi36 := 0; fi36 < 20; fi36++ {
                    v13 = v13
                    _ = fi36
                }
                _ = v34
                _ = v35
            }
            v37 := GenS1{F0: 1}
            _ = v37
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

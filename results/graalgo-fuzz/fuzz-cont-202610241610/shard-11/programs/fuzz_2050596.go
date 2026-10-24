package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int32 }
type GenS2 struct { F0 int16; F1 int; F2 int8 }
func genFn0(p0 string, p1 int8) uint64 {
    v7 := []int{128, 42, -9223372036854775808}
    v8 := GenS1{F0: -1, F1: 0, F2: 1}
    for fi9 := 0; fi9 < 9; fi9++ {
        p1 -= 1
        p1 = (-(p1))
        _ = fi9
    }
    p0 = "hello"
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 721
}
func genFn1(p0 uint) uint8 {
    var v10 int8 = 1
    var v11 bool = (-2147483648 != 2147483647)
    switch 32768 {
    case 32768:
        v10 += v10
    }
    var v12 int32 = 0
    if (1.5 <= 1e10) {
        var v13 float32 = 0.0
        v11 = (1.5 < 0.1)
        _ = v13
    } else {
        if v11 {
            v10 = v10
            for fi14 := 0; fi14 < 5; fi14++ {
                var v15 float32 = (3.14159 * ((-2.25 / 3.14159) * 0.0))
                for fi16 := 0; fi16 < 13; fi16++ {
                    if v11 {
                        fmt.Println(v10)
                        v17 := []int{}
                        _ = v17
                    }
                    _ = fi16
                }
                var v18 uint8 = 128
                _ = fi14
                _ = v15
                _ = v18
            }
            fmt.Println(v10, v11, v11)
        }
        for fi19 := 0; fi19 < 5; fi19++ {
            v11 = true
            p0 -= p0
            _ = fi19
        }
        for fi20 := 0; fi20 < 10; fi20++ {
            v21 := []int{(fi20 >> ((0) & 63)), (fi20 * (fi20 & (fi20 >> ((p0) & 63)))), 610}
            var v22 uint64 = 84
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    return 183
}

func main() {
    var v1 int8 = 1
    var v2 int32 = 438
    var v3 int8 = (v1 - v1)
    var v4 uint = 2
    var v5 uint64 = 797
    var v6 int32 = v2
    var v23 uint64 = v5
    var v24 uint = uint(v5)
    switch ("\n\t" + ("hello" + "a")) {
    case "foo bar":
        var v25 uint64 = (v5 >> ((v24) & 63))
        _ = v25
    default:
        var v26 string = ("a" + "foo bar")
        _ = v26
    }
    for fi27 := 0; fi27 < 3; fi27++ {
        v3 = v1
        _ = fi27
    }
    var v28 int64 = 985
    v29 := GenS2{F0: int16(v2), F1: -1, F2: v1}
    fmt.Println(v2, v28, v24)
    if false {
        var v30 bool = true
        var v31 uint8 = 1
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 13; fi32++ {
        v33 := []int{(fi32 << (((v24 % (((v4 / ((v24) | 1))) | 1))) & 63)), -409}
        v2 -= v6
        _ = fi32
        _ = v33
    }
    for fi34 := 0; fi34 < 7; fi34++ {
        v35 := GenS1{F0: v28, F1: 0, F2: v6}
        if ((3.14159 - 1.5) >= 0.1) {
            var v36 int8 = v3
            if (true || (v28 > v28)) {
                v37 := []int{(fi34 / ((344) | 1)), (fi34 >> ((v24) & 63)), fi34, (fi34 << (((v4 % (((v4 % (((v24 % (((v24 ^ v4)) | 1))) | 1))) | 1))) & 63))}
                v37 = append(v37, fi34)
                for fi38 := 0; fi38 < 11; fi38++ {
                    v24 -= 65535
                    var v39 int32 = (v6 * (v6 % (((v2 >> ((v4) & 63))) | 1)))
                    _ = fi38
                    _ = v39
                }
                _ = v37
            }
            var v40 int = fi34
            _ = v36
            _ = v40
        } else {
            v41 := []int{fi34, (fi34 | fi34), fi34, (fi34 + -650)}
            fmt.Println(v3, v23, v28)
            _ = v41
        }
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

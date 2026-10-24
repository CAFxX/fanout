package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint16 }
func genFn0(p0 int8) int32 {
    var v7 uint32 = 4294967295
    v8 := []int{-9223372036854775808, -9223372036854775808, -7}
    v9 := GenS1{F0: 2, F1: -1, F2: 451}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return -2147483648
}
func genFn1(p0 string, p1 uint) uint8 {
    p1 -= (p1 * (p1 + (p1 ^ 153)))
    if (0 > 1) {
        var v10 int64 = 732
        var v11 uint32 = 1000
        _ = v10
        _ = v11
    }
    var v12 uint16 = 0
    _ = p0
    _ = p1
    _ = v12
    return 13
}

func main() {
    var v1 uint = 65535
    var v2 int8 = 100
    var v3 string = ("hello" + "")
    var v4 uint32 = 326
    var v5 string = "a"
    v6 := make(map[int]int)
    if (v1 <= v1) {
        var v13 string = v5
        v14 := []int{-9223372036854775808, 2147483647}
        _ = v13
        _ = v14
    } else {
        v1 = v1
        var v15 bool = (127 >= 78)
        _ = v15
    }
    if false {
        for fi16 := 0; fi16 < 6; fi16++ {
            v3 = ""
            _ = fi16
        }
    } else {
        v5 = "\n\t"
        var v17 uint64 = 733
        _ = v17
    }
    fmt.Println(v4, v1, v1)
    for fi18 := 0; fi18 < 2; fi18++ {
        if ((false || (false && true)) || true) {
            for fi19 := 0; fi19 < 19; fi19++ {
                fmt.Println(v2)
                var v20 float32 = 3.14159
                fmt.Println(v20, v3)
                fmt.Println(fi18, fi18)
                _ = fi19
                _ = v20
            }
            switch v3 {
            case "a":
                v21 := GenS1{F0: 127, F1: genFn0((v2 * (v2 | v2))), F2: 32767}
                _ = v21
            case "":
                var v22 int64 = int64(v4)
                _ = v22
            case "hello":
                var v23 uint64 = 0
                _ = v23
            default:
                var v24 uint16 = 81
                _ = v24
            }
        } else {
            var v25 uint16 = 2
            _ = v25
        }
        _ = fi18
    }
    v26 := GenS1{F0: genFn1(v3, (v1 | 240)), F1: genFn0(v2), F2: 255}
    for fi27 := 0; fi27 < 5; fi27++ {
        var v28 uint64 = 555
        _ = fi27
        _ = v28
    }
    var v29 int = -9223372036854775808
    var v30 bool = false
    if (42 == 18446744073709551615) {
        var v31 float32 = 1.5
        for fi32 := 0; fi32 < 13; fi32++ {
            v6[(v29 & fi32)] = (fi32 * v29)
            var v33 int32 = -47
            _ = fi32
            _ = v33
        }
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}

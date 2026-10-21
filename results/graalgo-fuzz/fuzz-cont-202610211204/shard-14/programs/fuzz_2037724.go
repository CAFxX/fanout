package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint8 }
type GenS2 struct { F0 string; F1 uint16; F2 int16 }
type GenS3 struct { F0 int; F1 int16; F2 int }
func genFn0() uint8 {
    if false {
        v7 := []int{32768, -1}
        var v8 float64 = (3.14159 * float64(3.14159))
        _ = v7
        _ = v8
    } else {
        v9 := []int{1, 1, -2147483648, 2147483647}
        var v10 uint = 42
        if true {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        } else {
            v10 = 809
            v10 = v10
            var v11 float32 = 0.1
            _ = v11
        }
        _ = v10
        _ = v9
    }
    v12 := GenS1{F0: -787, F1: 9223372036854775807, F2: 77}
    var v13 uint64 = 42
    var v14 bool = ("a" != ("a" + ""))
    _ = v12
    _ = v13
    _ = v14
    return 1
}

func main() {
    var v1 bool = true
    var v2 uint = 1
    v3 := make(map[int]bool)
    var v4 uint16 = 1
    var v5 string = ("a" + "\n\t")
    var v6 uint16 = (v4 | v4)
    switch v5 {
    case "a":
        v15 := GenS1{F0: -356, F1: 758, F2: 128}
        _ = v15
    case "":
        v5 = v5
    }
    var v16 bool = v1
    v17 := GenS1{F0: 127, F1: 0, F2: 0}
    var v18 float32 = -2.25
    v6 = v4
    if (127 < 32768) {
        v19 := GenS2{F0: "", F1: (v4 % ((v6) | 1)), F2: -32768}
        if true {
            v20 := []int{-1, 2147483647, 128}
            var v21 float64 = 1.5
            if (!v16) {
                switch "" {
                case "hello":
                    v3[1] = ((!true) || (!false))
                case "foo bar":
                    var v22 int8 = 0
                    _ = v22
                case "\n\t":
                    v21 += 3.14159
                default:
                    v23 := GenS1{F0: 32768, F1: 42, F2: 219}
                    _ = v23
                }
                var v24 uint16 = (v6 + 0)
                switch 2147483647 {
                case 1:
                    v16 = false
                case 2147483647:
                    fmt.Println(v21, v16)
                }
                _ = v24
            } else {
                v21 += (v21 - 3.14159)
            }
            _ = v20
            _ = v21
        }
        _ = v19
    } else {
        v6 = (v4 ^ (v6 + (^(v4))))
        v1 = true
        if (("" + v5) >= v5) {
            if true {
                switch -641 {
                case 1:
                    if (-128 == 127) {
                        v25 := GenS2{F0: v5, F1: (v6 >> (((v2 >> ((v2) & 63))) & 63)), F2: -32768}
                        _ = v25
                    }
                }
                var v26 float32 = v18
                _ = v26
            }
        }
    }
    switch "a" {
    case "":
        v5 = v5
    case "foo bar":
        v3[673] = v16
    case "hello":
        if v1 {
            v2 = v2
            v27 := GenS3{F0: 128, F1: 32767, F2: -663}
            _ = v27
        }
    default:
        for fi28 := 0; fi28 < 17; fi28++ {
            fmt.Println(v16, v6)
            var v29 int64 = 865
            _ = fi28
            _ = v29
        }
    }
    fmt.Println(v2, v5, v6)
    var v30 bool = v16
    if (-863 <= -259) {
        v31 := v3[42]
        v3[-1] = v1
        _ = v31
    }
    var v32 int64 = 9223372036854775807
    if true {
        if true {
            v18 -= (v18 + 3.14159)
            for fi33 := 0; fi33 < 7; fi33++ {
                v16 = (!v30)
                if (!(1 >= 42)) {
                    var v34 int = (fi33 << (((^(v2))) & 63))
                    v35 := []int{1, v34, 4294967295}
                    _ = v34
                    _ = v35
                }
                for fi36 := 0; fi36 < 15; fi36++ {
                    v18 -= 0.1
                    v32 += (v32 >> (((^(v2))) & 63))
                    _ = fi36
                }
                _ = fi33
            }
            switch v5 {
            case "":
                if (-1 != 0) {
                    v6 = v4
                    v37 := GenS2{F0: v5, F1: (v4 + (v4 ^ (v6 ^ 32767))), F2: 0}
                    if (257 < -9223372036854775808) {
                        var v38 int64 = (v32 + (v32 * v32))
                        v16 = (!v16)
                        _ = v38
                    } else {
                        v39 := GenS3{F0: 259, F1: -658, F2: -2147483648}
                        fmt.Println(v18, v30)
                        var v40 uint32 = 0
                        _ = v39
                        _ = v40
                    }
                    _ = v37
                } else {
                    fmt.Println(v1, v32, v32)
                    v6 = uint16(v2)
                }
            case "\n\t":
                var v41 uint = 575
                _ = v41
            }
        }
        var v42 uint32 = 2
        v43 := []int{38, 295, 127, int(v32)}
        _ = v42
        _ = v43
    }
    v44 := []int{}
    var v45 string = (("\n\t" + "hello") + v5)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}

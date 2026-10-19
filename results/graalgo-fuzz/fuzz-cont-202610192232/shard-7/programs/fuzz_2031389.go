package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 bool; F2 bool }
type GenS4 struct { F0 uint; F1 int64; F2 int16 }

func main() {
    var v1 uint16 = 1
    var v2 string = ""
    var v3 int16 = 0
    var v4 string = v2
    var v5 uint8 = 240
    v1 -= (v1 >> ((uint(v1)) & 63))
    v4 = v2
    v6 := GenS1{F0: 4294967295}
    var v7 int8 = 0
    var v8 int32 = -1
    var v9 int = 32767
    if (true || true) {
        switch -109 {
        case -9223372036854775808:
            if false {
                var v10 uint32 = 987
                _ = v10
            } else {
                v11 := GenS2{F0: 205}
                if false {
                    fmt.Println(v7)
                    v4 = v2
                } else {
                    v12 := GenS1{F0: 863}
                    _ = v12
                }
                _ = v11
            }
        case -676:
            v8 -= v8
        case 32768:
            v2 = ("a" + v2)
        }
    } else {
        switch "\n\t" {
        case "a":
            var v13 uint16 = v1
            _ = v13
        case "hello":
            fmt.Println(v2, v5, v7)
        case "\n\t":
            v14 := []int{128, 4294967295, -533, (v9 * -2147483648)}
            _ = v14
        default:
            var v15 uint = 42
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        v17 := GenS2{F0: 0}
        var v18 uint8 = v5
        var v19 float64 = ((3.14159 * 0.1) * (0.0 + 1.5))
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        var v21 string = (v4 + (v4 + "foo bar"))
        for fi22 := 0; fi22 < 17; fi22++ {
            fmt.Println(v7, v3)
            if false {
                if ((true && false) || false) {
                    if (2 == 378) {
                        var v23 bool = ((fi22 - (v9 % (((v9 >> ((1) & 63))) | 1))) == fi20)
                        v24 := GenS3{F0: 255, F1: v23, F2: v23}
                        switch ((("foo bar" + v4) + v4) + v21) {
                        case "\n\t":
                            v9 = int(v3)
                        case "":
                            fmt.Println(v23, v5, v4)
                        default:
                            var v25 int32 = v8
                            _ = v25
                        }
                        _ = v23
                        _ = v24
                    } else {
                        fmt.Println(v9)
                    }
                    switch -9223372036854775808 {
                    case 99:
                        var v26 int16 = (v3 - v3)
                        _ = v26
                    }
                    v27 := []int{1, (v9 & fi22), 42, fi22}
                    _ = v27
                } else {
                    var v28 uint32 = 27
                    var v29 uint16 = 255
                    _ = v28
                    _ = v29
                }
            } else {
                var v30 string = v21
                var v31 int16 = v3
                _ = v30
                _ = v31
            }
            var v32 string = v4
            _ = fi22
            _ = v32
        }
        _ = fi20
        _ = v21
    }
    v33 := GenS4{F0: 0, F1: -1, F2: (v3 + v3)}
    v34 := GenS2{F0: 256}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

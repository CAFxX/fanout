package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 int16; F2 uint8 }

func main() {
    var v1 uint16 = 190
    var v2 int32 = -635
    var v3 uint8 = 128
    var v4 string = "hello"
    v5 := GenS1{F0: 4294967295}
    v6 := GenS2{F0: (v3 - v3), F1: -1, F2: (^(v3))}
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 uint8 = v3
        switch v4 {
        case "a":
            var v9 uint8 = (^(v8))
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        fmt.Println(fi10, v3)
        v11 := GenS2{F0: 255, F1: -1, F2: v3}
        _ = fi10
        _ = v11
    }
    var v12 float64 = 0.1
    fmt.Println(v1, v2)
    for fi13 := 0; fi13 < 2; fi13++ {
        v1 = v1
        _ = fi13
    }
    var v14 uint32 = 242
    v1 = (v1 - v1)
    var v15 int64 = 1
    if true {
        if true {
            v16 := []int{-2147483648, int(v1), 79, 9223372036854775807}
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            if true {
                var v17 int8 = 42
                _ = v17
            }
            _ = v16
        } else {
            var v18 string = v4
            v19 := []int{44}
            if len(v19) > 0 {
                v19[0] = 280
            }
            _ = v18
            _ = v19
        }
        v4 = "hello"
        for fi20 := 0; fi20 < 2; fi20++ {
            switch "\n\t" {
            case "foo bar":
                if true {
                    var v21 float32 = float32(v12)
                    _ = v21
                }
            case "hello":
                fmt.Println(v4, v3)
            case "a":
                var v22 uint64 = 42
                _ = v22
            }
            if true {
                v23 := GenS2{F0: (v3 << ((864) & 63)), F1: 807, F2: (v3 / (((^(v3))) | 1))}
                if (((255 <= 2) && (v15 == v15)) && false) {
                    fmt.Println(v4)
                    var v24 int32 = -359
                    _ = v24
                }
                _ = v23
            } else {
                v25 := GenS1{F0: 4294967295}
                var v26 int8 = -1
                _ = v25
                _ = v26
            }
            _ = fi20
        }
    } else {
        fmt.Println(v14)
    }
    var v27 int16 = 203
    switch -2147483648 {
    case -958:
        switch 2147483647 {
        case -1:
            v28 := []int{-2147483648, -319, -2147483648}
            _ = v28
        case 0:
            if false {
                for fi29 := 0; fi29 < 18; fi29++ {
                    var v30 uint8 = (v3 % (((v3 ^ v3)) | 1))
                    v31 := GenS1{F0: 255}
                    v32 := GenS1{F0: 2}
                    _ = fi29
                    _ = v30
                    _ = v31
                    _ = v32
                }
                var v33 bool = true
                _ = v33
            }
        }
    case -9223372036854775808:
        if (false && (9223372036854775807 > -1)) {
            v3 -= uint8(v14)
            v34 := []int{-9223372036854775808, int(v3), 205, -1}
            _ = v34
        } else {
            fmt.Println(v12, v4)
            var v35 float32 = float32((v12 * v12))
            _ = v35
        }
    default:
        for fi36 := 0; fi36 < 7; fi36++ {
            var v37 uint32 = v14
            v37 = uint32(v1)
            _ = fi36
            _ = v37
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

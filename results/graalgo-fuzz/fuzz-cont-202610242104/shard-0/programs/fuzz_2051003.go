package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint64 }
func genFn0() int32 {
    var v4 int64 = 42
    if (1 != -32768) {
        v5 := GenS1{F0: 9223372036854775807}
        if (1 != 2147483648) {
            v6 := []int{0}
            fmt.Println(v4)
            if false {
                for fi7 := 0; fi7 < 14; fi7++ {
                    var v8 uint = 367
                    for fi9 := 0; fi9 < 14; fi9++ {
                        var v10 uint8 = 213
                        var v11 float32 = (3.14159 * float32(0.1))
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                    var v12 uint64 = 0
                    _ = fi7
                    _ = v12
                    _ = v8
                }
                if true {
                    v13 := GenS1{F0: 2147483647}
                    _ = v13
                }
                switch -481 {
                case 127:
                    var v14 uint16 = 65535
                    _ = v14
                case 73:
                    switch "foo bar" {
                    case "a":
                        v4 -= (v4 % ((-9223372036854775808) | 1))
                    }
                case -202:
                    fmt.Println(v4)
                }
            }
            _ = v6
        }
        _ = v5
    }
    v15 := GenS2{F0: 42, F1: 65535, F2: 0}
    _ = v15
    _ = v4
    return -1
}

func main() {
    var v1 bool = false
    var v2 int16 = 1000
    var v3 int32 = 1
    switch "hello" {
    case "a":
        v16 := GenS2{F0: 2147483647, F1: 42, F2: 1}
        _ = v16
    case "":
        v17 := GenS1{F0: 4294967295}
        _ = v17
    }
    switch 42 {
    case 32768:
        var v18 uint32 = 576
        _ = v18
    case 524:
        v2 += v2
    }
    if true {
        fmt.Println(v3, v2, v2)
    } else {
        if true {
            fmt.Println(v3, v2)
            v19 := GenS1{F0: 42}
            _ = v19
        }
    }
    var v20 int64 = -67
    switch 32767 {
    case 394:
        if (!(((0 >= 42) && (758 <= 9223372036854775807)) && false)) {
            var v21 uint16 = 2
            v22 := GenS2{F0: 911, F1: 65535, F2: 1}
            _ = v21
            _ = v22
        }
    }
    v1 = (!(!true))
    fmt.Println(v1, v3, v1)
    var v23 int16 = 1000
    var v24 int32 = v3
    var v25 uint64 = 435
    var v26 int8 = int8(v20)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
func genFn0() int {
    switch -1 {
    case 42:
        var v5 uint16 = 123
        _ = v5
    case -44:
        var v6 uint8 = 127
        _ = v6
    case 0:
        switch 42 {
        case 32767:
            var v7 int64 = 0
            _ = v7
        case 32768:
            v8 := []int{}
            _ = v8
        }
    }
    v9 := []int{9223372036854775807, 707, 9223372036854775807}
    var v10 uint16 = 0
    fmt.Println(v10)
    if true {
        var v11 uint = 255
        _ = v11
    }
    _ = v10
    _ = v9
    return 32768
}
func genFn1() int16 {
    var v12 uint8 = 2
    if false {
        var v13 uint64 = 0
        v13 += 567
        v12 -= (v12 | (v12 | v12))
        _ = v13
    }
    _ = v12
    return 32767
}

func main() {
    var v1 int32 = 2147483647
    var v2 bool = (!true)
    var v3 int16 = 336
    var v4 int = 9223372036854775807
    switch "hello" {
    case "\n\t":
        switch "" {
        case "a":
            v14 := []int{v4, v4}
            _ = v14
        default:
            var v15 uint16 = 0
            _ = v15
        }
    }
    var v16 string = ""
    v3 = v3
    var v17 bool = v2
    var v18 int64 = 530
    var v19 int = (v4 - 0)
    v16 = (v16 + v16)
    v20 := []int{}
    fmt.Println(v19)
    v20 = append(v20, v4)
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v1)
    case "":
        v19 = (v4 >> ((42) & 63))
    }
    v19 = v19
    if false {
        var v21 uint16 = 42
        var v22 uint8 = 0
        _ = v21
        _ = v22
    }
    fmt.Println(v1, v4)
    var v23 float64 = (3.14159 / 0.1)
    for fi24 := 0; fi24 < 10; fi24++ {
        if (0 <= 450) {
            var v25 uint64 = 2
            _ = v25
        } else {
            if (!((v23 + float64(3.14159)) <= v23)) {
                v4 += v4
            }
        }
        var v26 string = v16
        _ = fi24
        _ = v26
    }
    if (false || (!v17)) {
        v27 := GenS1{F0: 2, F1: 18446744073709551615}
        _ = v27
    } else {
        _ = len(v20)
        for fi28 := 0; fi28 < 11; fi28++ {
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            var v29 uint32 = 65535
            for fi30 := 0; fi30 < 5; fi30++ {
                v16 = "hello"
                if len(v20) > 0 {
                    v20[0] = genFn0()
                }
                switch genFn0() {
                case 0:
                    if true {
                        v1 = v1
                    } else {
                        var v31 int64 = (v18 - (v18 >> ((2) & 63)))
                        v16 = (((v16 + v16) + ("a" + (v16 + ""))) + ((v16 + v16) + v16))
                        v18 += v18
                        var v32 bool = false
                        if (v23 >= ((float64((-2.25 + 0.1)) * (3.14159 - 1.5)) + v23)) {
                            fmt.Println(v2, fi28, v31)
                        }
                        var v33 int = fi28
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                    v19 = (fi28 >> ((650) & 63))
                case -318:
                    v20 = append(v20, 0)
                case -157:
                    v34 := GenS1{F0: 1, F1: 1}
                    _ = v34
                }
                _ = fi30
            }
            _ = fi28
            _ = v29
        }
        var v35 uint = 1
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
}

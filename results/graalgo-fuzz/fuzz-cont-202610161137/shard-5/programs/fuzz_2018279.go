package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() string {
    v5 := []int{127}
    v5 = append(v5, 4294967295)
    if true {
        for fi6 := 0; fi6 < 18; fi6++ {
            for fi7 := 0; fi7 < 5; fi7++ {
                var v8 int32 = 1
                switch "\n\t" {
                case "hello":
                    v9 := GenS1{F0: 1}
                    _ = v9
                case "a":
                    v8 = 1
                case "\n\t":
                    v10 := GenS1{F0: -128}
                    if (true || (!false)) {
                        v5 = append(v5, fi7)
                    } else {
                        v8 = (v8 - (v8 - v8))
                        v11 := GenS1{F0: 42}
                        v8 = v8
                        _ = v11
                    }
                    _ = v10
                default:
                    var v12 int8 = 1
                    _ = v12
                }
                _ = fi7
                _ = v8
            }
            v13 := GenS1{F0: 1}
            _ = fi6
            _ = v13
        }
    }
    _ = v5
    return "hello"
}

func main() {
    v1 := []int{}
    var v2 uint64 = 383
    var v3 bool = true
    v4 := []int{-9223372036854775808, 9223372036854775807, 754}
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 uint64 = (v2 | (v2 | (^(v2))))
        if v3 {
            v2 = v15
            var v16 int8 = 0
            _ = v16
        } else {
            var v17 uint16 = 42
            if true {
                switch -9223372036854775808 {
                case 9223372036854775807:
                    for fi18 := 0; fi18 < 11; fi18++ {
                        v1 = append(v1, 4294967295)
                        _ = fi18
                    }
                case -871:
                    v19 := GenS1{F0: 100}
                    _ = v19
                }
                if (true && (!(!true))) {
                    for fi20 := 0; fi20 < 19; fi20++ {
                        v17 = v17
                        fmt.Println(fi14, v15, v17)
                        fmt.Println(v3, v15, v3)
                        _ = fi20
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v21 := GenS1{F0: 100}
                    _ = v21
                } else {
                    var v22 uint = 255
                    if false {
                        var v23 bool = true
                        _ = v23
                    }
                    fmt.Println(v17, v3)
                    _ = v22
                }
                fmt.Println(v2)
            }
            _ = v17
        }
        _ = fi14
        _ = v15
    }
    fmt.Println(v2, v2)
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v2)
    }
    v2 = v2
    v3 = false
    var v24 uint32 = 559
    for fi25 := 0; fi25 < 15; fi25++ {
        v26 := GenS1{F0: 0}
        v2 = (v2 & v2)
        v3 = false
        _ = fi25
        _ = v26
    }
    v3 = v3
    v4 = append(v4, 42)
    var v27 int16 = -32768
    var v28 int16 = (v27 << ((42) & 63))
    var v29 uint32 = v24
    switch 334 {
    case -2147483648:
        fmt.Println(v24, v29, v3)
    default:
        v29 -= v29
    }
    var v30 bool = (!false)
    var v31 bool = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}

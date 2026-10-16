package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 0
    v3 := make(map[int]uint32)
    var v4 string = "hello"
    for fi5 := 0; fi5 < 4; fi5++ {
        switch "hello" {
        case "a":
            v2 -= v2
        default:
            for fi6 := 0; fi6 < 13; fi6++ {
                fmt.Println(v2, fi6, v2)
                _ = fi6
            }
        }
        var v7 int16 = -559
        _ = fi5
        _ = v7
    }
    v8 := GenS1{F0: 1, F1: 0}
    var v9 int16 = 0
    var v10 int64 = -1
    v4 = "foo bar"
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := GenS1{F0: -1, F1: -43}
        fmt.Println(fi11, v9, fi11)
        v9 = (v9 + (^(v9)))
        _ = fi11
        _ = v12
    }
    switch v4 {
    case "a":
        switch 32767 {
        case 613:
            v13 := GenS1{F0: -30, F1: 1}
            _ = v13
        default:
            switch -2147483648 {
            case 9223372036854775807:
                for fi14 := 0; fi14 < 14; fi14++ {
                    var v15 string = v4
                    if (9223372036854775807 > 7) {
                        var v16 float64 = (3.14159 - -2.25)
                        _ = v16
                    }
                    v17 := v3[2147483647]
                    _ = fi14
                    _ = v15
                    _ = v17
                }
            case 32768:
                var v18 uint = 65
                _ = v18
            case -9223372036854775808:
                for fi19 := 0; fi19 < 2; fi19++ {
                    delete(v3, (fi19 << ((107) & 63)))
                    _ = fi19
                }
            }
        }
    case "foo bar":
        switch -1 {
        case 0:
            switch (v4 + "a") {
            case "\n\t":
                if false {
                    v20 := GenS1{F0: 1, F1: 42}
                    var v21 string = (("" + "") + ("" + "a"))
                    _ = v20
                    _ = v21
                }
            }
        case -595:
            v22 := []int{42, 42, 42, 631}
            _ = v22
        case 1:
            for fi23 := 0; fi23 < 13; fi23++ {
                v24 := []int{(^(fi23)), (fi23 >> ((1) & 63))}
                _ = fi23
                _ = v24
            }
        default:
            v9 = (^(v9))
        }
    case "\n\t":
        v25 := v3[32768]
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}

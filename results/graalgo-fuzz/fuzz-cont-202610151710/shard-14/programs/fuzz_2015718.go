package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 bool = true
    var v2 uint8 = 255
    v3 := make(map[int]int)
    var v4 uint8 = v2
    v5 := make(map[string]string)
    v6 := v5["hello"]
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 int16 = 1000
            v4 -= v4
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 int64 = -9223372036854775808
            var v13 int64 = v12
            switch ("\n\t" + (v6 + ((v6 + v6) + v6))) {
            case "hello":
                switch "hello" {
                case "a":
                    var v14 int8 = 127
                    _ = v14
                case "hello":
                    v1 = v1
                case "":
                    var v15 int16 = -399
                    _ = v15
                }
            case "foo bar":
                switch fi10 {
                case 809:
                    if false {
                        fmt.Println(v6, v2)
                        if v1 {
                            v16 := GenS1{F0: v1}
                            fmt.Println(v2, v2)
                            v17 := []int{(fi11 * (fi11 - fi10)), -2147483648, fi10}
                            _ = v16
                            _ = v17
                        } else {
                            fmt.Println(fi11, v13, fi11)
                            fmt.Println(v1, v2, fi10)
                        }
                        v3[fi11] = (fi11 % (((fi11 / ((9223372036854775807) | 1))) | 1))
                    } else {
                        v2 -= (v2 ^ v2)
                    }
                case 127:
                    var v18 uint8 = v4
                    _ = v18
                case 9223372036854775807:
                    v19 := v5[v6]
                    _ = v19
                default:
                    v20 := GenS1{F0: v1}
                    _ = v20
                }
            case "":
                var v21 uint = 695
                _ = v21
            }
            _ = fi11
            _ = v12
            _ = v13
        }
        var v22 uint16 = 2
        _ = fi10
        _ = v22
    }
    v5[((((v6 + v6) + v6) + ("\n\t" + v6)) + "\n\t")] = v6
    var v23 int64 = -1
    fmt.Println(v6)
    var v24 uint16 = 255
    var v25 uint8 = (v4 - (v2 * v4))
    v26 := []int{42}
    v27 := GenS1{F0: (((((((3.14159 * (1.5 - (-2.25 / (1e10 / 1e10)))) / 1.5) * 0.0) - (3.14159 / 1e10)) == float64((0.0 - 1e10))) && (!false)) && (!(false || true)))}
    for fi28 := 0; fi28 < 10; fi28++ {
        switch "a" {
        case "a":
            var v29 int8 = -1
            _ = v29
        }
        v6 = v6
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

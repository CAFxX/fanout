package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 float64 = 3.14159
    var v2 uint32 = 2147483647
    var v3 float64 = (float64(0.0) + -2.25)
    var v4 int16 = 1000
    v5 := make(map[int]uint32)
    if false {
        var v6 uint8 = 1
        v7 := []int{42, -2147483648, 4294967295, 42}
        if true {
            v1 = v1
            fmt.Println(v4, v4)
        } else {
            if ((0.1 * (1.5 + ((1e10 / 1.0) + (1.5 + 1e10)))) > -2.25) {
                switch "a" {
                case "hello":
                    v8 := GenS1{F0: v6}
                    _ = v8
                case "foo bar":
                    var v9 int8 = 1
                    _ = v9
                case "":
                    for fi10 := 0; fi10 < 13; fi10++ {
                        if false {
                            v4 += 658
                            var v11 bool = ((float64(1.5) - 1.5) != (v3 + v3))
                            _ = v11
                        }
                        v7 = append(v7, -1)
                        for fi12 := 0; fi12 < 1; fi12++ {
                            delete(v5, int(v6))
                            _ = fi12
                        }
                        _ = fi10
                    }
                default:
                    v3 = float64(1.5)
                }
                var v13 float64 = 1.5
                _ = v13
            }
        }
        _ = v6
        _ = v7
    }
    if (int16(v2) < (^(v4))) {
        v2 = v2
        var v14 int32 = 64
        _ = v14
    } else {
        v2 = 336
        switch 1 {
        case 42:
            fmt.Println(v3)
        case -53:
            for fi15 := 0; fi15 < 2; fi15++ {
                var v16 int32 = 42
                if false {
                    var v17 string = "a"
                    var v18 int64 = 912
                    v19 := v5[fi15]
                    _ = v17
                    _ = v18
                    _ = v19
                }
                var v20 int64 = 42
                _ = fi15
                _ = v16
                _ = v20
            }
        case -901:
            v21 := GenS1{F0: 1}
            _ = v21
        }
        var v22 uint8 = 128
        _ = v22
    }
    fmt.Println(v4, v3, v4)
    if (!false) {
        v23 := []int{1, -1, -9223372036854775808, 127}
        _ = v23
    }
    switch ("\n\t" + "a") {
    case "hello":
        for fi24 := 0; fi24 < 19; fi24++ {
            v4 = v4
            _ = fi24
        }
    case "foo bar":
        fmt.Println(v1, v2)
    case "":
        switch 128 {
        case -2147483648:
            var v25 float64 = ((v1 / (((1.5 - (float64(0.0) / float64(1e10))) * -2.25) * v3)) * -2.25)
            _ = v25
        case 32768:
            for fi26 := 0; fi26 < 15; fi26++ {
                var v27 uint64 = 2
                _ = fi26
                _ = v27
            }
        case 101:
            delete(v5, 128)
        }
    }
    if (100 <= 0) {
        if (true && (24 > 4294967295)) {
            fmt.Println(v1)
            v3 += (((v3 + 1.5) * 1.5) / v3)
            v4 = int16(v2)
        }
    } else {
        v28 := GenS1{F0: 127}
        for fi29 := 0; fi29 < 3; fi29++ {
            switch "foo bar" {
            case "foo bar":
                var v30 int32 = 0
                _ = v30
            default:
                v31 := GenS1{F0: 128}
                _ = v31
            }
            var v32 bool = true
            fmt.Println(v1, v32)
            _ = fi29
            _ = v32
        }
        _ = v28
    }
    if ((!true) && true) {
        var v33 int32 = -1
        _ = v33
    } else {
        v34 := []int{-2147483648}
        fmt.Println(v4)
        _ = v34
    }
    var v35 uint32 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}

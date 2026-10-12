package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() uint32 {
    var v5 uint8 = 0
    v6 := []int{-936, 996}
    switch 42 {
    case -2147483648:
        if false {
            _ = len(v6)
        }
    case 0:
        for fi7 := 0; fi7 < 15; fi7++ {
            switch "a" {
            case "\n\t":
                var v8 float64 = ((float64(float32(((1.5 * (3.14159 - ((0.0 * 1e10) * 0.0))) + -2.25))) * 1e10) / float64(0.1))
                _ = v8
            case "hello":
                v5 = v5
            case "foo bar":
                v6 = append(v6, 32767)
            }
            var v9 uint16 = 1
            v9 -= 824
            _ = fi7
            _ = v9
        }
    }
    _ = v5
    _ = v6
    return 4294967295
}

func main() {
    var v1 string = "foo bar"
    var v2 uint8 = 1
    v3 := make(map[uint32]int)
    var v4 uint8 = v2
    if (42 < 255) {
        v10 := GenS1{F0: 4294967295}
        if (1.5 < (3.14159 + 0.0)) {
            v4 = v2
            var v11 uint64 = 2
            _ = v11
        } else {
            v12 := []int{-9223372036854775808}
            _ = v12
        }
        _ = v10
    } else {
        var v13 string = v1
        _ = v13
    }
    if (((!(false || false)) && (((!((v2 >> ((2) & 63)) != 1)) && (((0 <= 565) || (-1 == 100)) && true)) && true)) && true) {
        for fi14 := 0; fi14 < 3; fi14++ {
            v15 := []int{0}
            var v16 int = (fi14 << ((42) & 63))
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    var v17 int32 = 42
    v17 = (v17 + v17)
    fmt.Println(v17, v17, v1)
    var v18 uint64 = 391
    var v19 int32 = (v17 | (v17 ^ -439))
    fmt.Println(v4, v2)
    for fi20 := 0; fi20 < 9; fi20++ {
        v17 = -1
        v21 := []int{}
        _ = fi20
        _ = v21
    }
    var v22 int64 = -9223372036854775808
    if true {
        var v23 int16 = 1
        _ = v23
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        delete(v3, 303)
        _ = fi24
    }
    if (0.1 < 1e10) {
        _ = len(v3)
        v25 := GenS1{F0: 65535}
        _ = v25
    } else {
        for fi26 := 0; fi26 < 7; fi26++ {
            if false {
                for fi27 := 0; fi27 < 11; fi27++ {
                    var v28 int8 = 42
                    _ = fi27
                    _ = v28
                }
            }
            _ = fi26
        }
        for fi29 := 0; fi29 < 11; fi29++ {
            v3[4294967295] = 42
            _ = fi29
        }
        v1 = ((v1 + v1) + (v1 + v1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}

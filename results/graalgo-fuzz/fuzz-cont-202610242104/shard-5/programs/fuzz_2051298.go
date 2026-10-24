package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
func genFn0(p0 int8) uint {
    var v7 uint32 = 2
    v7 = v7
    _ = p0
    _ = v7
    return 2
}

func main() {
    var v1 int64 = 1
    var v2 int = 4294967295
    v3 := make(map[string]int)
    v4 := make(map[uint32]string)
    var v5 uint32 = 651
    var v6 uint32 = 2
    switch "a" {
    case "a":
        if true {
            var v8 uint64 = 2
            _ = v8
        } else {
            fmt.Println(v1, v2, v2)
            v9 := []int{(^(v2))}
            v10 := GenS1{F0: 853, F1: 9}
            _ = v10
            _ = v9
        }
    }
    fmt.Println(v1, v2)
    if false {
        v11 := GenS1{F0: 382, F1: 75}
        _ = v11
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        if false {
            for fi13 := 0; fi13 < 5; fi13++ {
                v14 := []int{}
                _ = fi13
                _ = v14
            }
        }
        _ = fi12
    }
    fmt.Println(v6, v1)
    v5 = (v6 ^ 1)
    if false {
        fmt.Println(v5, v6)
        fmt.Println(v5, v1, v5)
        if (-1 <= 844) {
            v2 = v2
        }
    } else {
        switch v2 {
        case -2147483648:
            for fi15 := 0; fi15 < 6; fi15++ {
                if false {
                    if (23 <= 694) {
                        switch (fi15 % ((-27) | 1)) {
                        case 815:
                            switch v2 {
                            case 9223372036854775807:
                                var v16 int64 = v1
                                _ = v16
                            case 42:
                                v6 = (v6 / ((v5) | 1))
                            case 571:
                                v17 := GenS1{F0: 18446744073709551615, F1: 0}
                                _ = v17
                            }
                        }
                        switch (((("hello" + "a") + "a") + "hello") + "") {
                        case "":
                            v5 = 0
                        case "\n\t":
                            var v18 uint = 65535
                            _ = v18
                        }
                    } else {
                        var v19 float32 = ((((-2.25 * (((1.5 - (0.1 * 0.1)) * -2.25) + ((0.0 + (0.0 + 3.14159)) * -2.25))) + 0.0) + (3.14159 * 0.0)) + 0.1)
                        v20 := v3[""]
                        _ = v19
                        _ = v20
                    }
                    var v21 string = "\n\t"
                    v22 := GenS1{F0: 42, F1: 1}
                    _ = v21
                    _ = v22
                }
                v1 = (v1 * 998)
                var v23 float32 = (-2.25 / 1.5)
                _ = fi15
                _ = v23
            }
        case -11:
            var v24 bool = true
            _ = v24
        }
        fmt.Println(v2)
    }
    v2 += (-(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

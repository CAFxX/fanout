package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint }

func main() {
    var v1 int32 = 42
    var v2 string = "a"
    v3 := []uint32{}
    var v4 uint32 = 42
    var v5 bool = (-2147483648 != -2147483648)
    if v5 {
        var v6 bool = false
        _ = v6
    } else {
        switch -1 {
        case 246:
            switch 1 {
            case 32768:
                v2 = v2
            }
        case -860:
            var v7 int = 0
            _ = v7
        }
        fmt.Println(v4, v1, v2)
        if false {
            v5 = (((9223372036854775807 >= -9223372036854775808) || false) && (127 != 127))
            v8 := GenS1{F0: 0, F1: 549, F2: 65535}
            _ = v8
        } else {
            switch (v2 + "") {
            case "":
                if (931 > -403) {
                    var v9 bool = (-2.25 < 0.1)
                    var v10 string = v2
                    _ = v10
                    _ = v9
                } else {
                    for fi11 := 0; fi11 < 14; fi11++ {
                        if v5 {
                            v2 = v2
                            var v12 float32 = ((3.14159 / (0.1 + 1.5)) * float32(0.1))
                            var v13 uint16 = 0
                            _ = v12
                            _ = v13
                        }
                        _ = len(v3)
                        _ = fi11
                    }
                }
            case "a":
                v5 = v5
            }
            v14 := GenS1{F0: 0, F1: 1, F2: 65535}
            _ = v14
        }
    }
    if len(v3) > 0 {
        v3[0] = 2147483648
    }
    var v15 string = v2
    v15 = ("a" + (v15 + v2))
    if (!(true || v5)) {
        fmt.Println(v4, v5)
        var v16 uint16 = 255
        _ = v16
    } else {
        v17 := GenS1{F0: 255, F1: 18446744073709551615, F2: 419}
        _ = v17
    }
    fmt.Println(v1, v5, v1)
    fmt.Println(v15)
    v1 = v1
    var v18 uint16 = 255
    v19 := GenS1{F0: 128, F1: 0, F2: 780}
    v4 = (v4 + (^(v4)))
    fmt.Println(v5, v2, v2)
    v20 := GenS1{F0: 42, F1: 2, F2: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}

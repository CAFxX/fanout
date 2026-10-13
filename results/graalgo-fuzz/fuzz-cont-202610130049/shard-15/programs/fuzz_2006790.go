package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int8 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 float32 = (0.0 * 1e10)
    v3 := make(map[uint32]string)
    v4 := make(map[string]uint32)
    var v5 uint8 = 217
    var v6 string = ""
    var v7 float64 = 3.14159
    switch -2147483648 {
    case -9223372036854775808:
        v2 -= (v2 - v2)
    default:
        v6 = ("a" + v6)
    }
    var v8 uint64 = 11
    v7 -= (((0.0 + 1.5) * v7) * -2.25)
    v9 := GenS1{F0: 752, F1: (((((v6 + v6) + ("hello" + v6)) + (v6 + ((v6 + (("foo bar" + "\n\t") + "")) + (("foo bar" + (v6 + v6)) + v6)))) + (v6 + "")) + v6), F2: 1}
    var v10 float32 = (v2 * v2)
    v2 -= -2.25
    for fi11 := 0; fi11 < 2; fi11++ {
        switch (v6 + "\n\t") {
        case "a":
            var v12 uint = 65535
            _ = v12
        default:
            if false {
                var v13 uint64 = v8
                _ = v13
            } else {
                delete(v3, 41)
                var v14 int16 = -1
                _ = v14
            }
        }
        _ = len(v1)
        _ = fi11
    }
    if (-106 > 42) {
        v15 := GenS1{F0: 2, F1: v6, F2: 0}
        if (!true) {
            switch -1 {
            case -721:
                var v16 int32 = -2147483648
                _ = v16
            default:
                v5 += (v5 | (v5 - (^(v5))))
            }
        }
        v2 = v2
        _ = v15
    } else {
        v17 := GenS1{F0: 560, F1: (v6 + v6), F2: -128}
        _ = v17
    }
    v18 := GenS2{F0: 1}
    var v19 int64 = 9223372036854775807
    v3[1] = v6
    v10 = v2
    var v20 uint = 1
    var v21 int = 42
    if (true && false) {
        var v22 float64 = -2.25
        if (true && true) {
            fmt.Println(v20)
        } else {
            switch (v21 >> ((v20) & 63)) {
            case 127:
                v2 = 0.1
            case 28:
                if (false && ("a" == "")) {
                    v23 := []int{}
                    v24 := GenS2{F0: 1}
                    _ = v23
                    _ = v24
                } else {
                    var v25 uint = (v20 >> ((v20) & 63))
                    v26 := GenS2{F0: (v19 << ((605) & 63))}
                    _ = v25
                    _ = v26
                }
            default:
                v27 := []int{(v21 * v21), (v21 * v21), (v21 % ((v21) | 1)), (v21 + v21)}
                _ = v27
            }
        }
        v28 := v3[11]
        _ = v22
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

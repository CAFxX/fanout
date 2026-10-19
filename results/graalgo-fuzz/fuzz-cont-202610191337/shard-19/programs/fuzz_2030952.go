package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int }
type GenS2 struct { F0 uint8; F1 uint16 }
func genFn0() uint16 {
    if false {
        if (!(true && (-417 == 1))) {
            if (32767 < -9223372036854775808) {
                var v6 string = "foo bar"
                var v7 float64 = ((0.1 + (0.0 + 1.5)) + (((-2.25 / 1e10) - 1.5) + (-2.25 * 3.14159)))
                _ = v6
                _ = v7
            } else {
                v8 := GenS1{F0: 42, F1: 127, F2: -988}
                var v9 uint = 255
                v10 := GenS1{F0: 2, F1: 128, F2: 0}
                _ = v10
                _ = v8
                _ = v9
            }
            for fi11 := 0; fi11 < 8; fi11++ {
                fmt.Println(fi11)
                _ = fi11
            }
            var v12 int = 42
            _ = v12
        }
    }
    var v13 int64 = 42
    if true {
        v13 += (-(v13))
        v13 -= 854
    }
    _ = v13
    return 2
}

func main() {
    var v1 uint32 = 55
    v2 := make(map[uint32]string)
    v3 := make(map[string]int)
    v4 := []int{2147483647}
    var v5 uint16 = 2
    v5 = 2
    v14 := GenS1{F0: 32767, F1: 127, F2: 0}
    fmt.Println(v5, v1)
    var v15 uint16 = 1
    if false {
        v15 -= 255
        switch ("a" + ("\n\t" + "")) {
        case "foo bar":
            var v16 int = 32767
            _ = v16
        case "hello":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "\n\t":
            var v17 bool = (true || true)
            _ = v17
        default:
            if (true || true) {
                var v18 uint16 = (v5 / (((v15 + v15)) | 1))
                for fi19 := 0; fi19 < 4; fi19++ {
                    var v20 uint64 = 42
                    _ = fi19
                    _ = v20
                }
                v21 := GenS1{F0: 916, F1: 1, F2: 9223372036854775807}
                _ = v18
                _ = v21
            }
        }
    }
    v5 = genFn0()
    if true {
        switch "foo bar" {
        case "hello":
            var v22 string = ("" + "a")
            _ = v22
        case "\n\t":
            switch int(v15) {
            case 1:
                v23 := GenS2{F0: 255, F1: (^(v15))}
                _ = v23
            default:
                delete(v2, v1)
            }
        case "a":
            v24 := GenS1{F0: genFn0(), F1: 1, F2: 32767}
            _ = v24
        }
    } else {
        var v25 uint16 = (v15 - v5)
        _ = v25
    }
    v5 += (v15 + 676)
    var v26 uint8 = 140
    var v27 int = 1
    v27 = v27
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}

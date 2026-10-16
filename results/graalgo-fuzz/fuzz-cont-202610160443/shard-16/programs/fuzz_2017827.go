package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
func genFn0(p0 int64, p1 uint) uint32 {
    switch (((("a" + (("\n\t" + "hello") + "hello")) + ("" + ("" + ""))) + "\n\t") + "a") {
    case "hello":
        for fi4 := 0; fi4 < 12; fi4++ {
            v5 := GenS1{F0: 32767, F1: p1}
            p0 = -897
            _ = fi4
            _ = v5
        }
    case "foo bar":
        var v6 uint64 = 753
        _ = v6
    case "":
        if (-317 == 1) {
            v7 := []int{-9223372036854775808, -1}
            switch 4294967295 {
            case 32768:
                _ = len(v7)
            default:
                p1 = 4294967295
            }
            var v8 uint16 = 234
            _ = v7
            _ = v8
        }
    default:
        var v9 int = 1
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 uint64 = 974
        if ((p1 | p1) <= p1) {
            var v12 int = fi10
            if ("\n\t" == "a") {
                var v13 float32 = (1.5 * 1e10)
                switch "\n\t" {
                case "\n\t":
                    for fi14 := 0; fi14 < 8; fi14++ {
                        v11 += v11
                        _ = fi14
                    }
                case "":
                    var v15 float64 = -2.25
                    _ = v15
                }
                v13 = v13
                _ = v13
            }
            switch ("hello" + (("a" + "\n\t") + "foo bar")) {
            case "\n\t":
                v11 += 24
            default:
                v16 := GenS1{F0: 32767, F1: 0}
                _ = v16
            }
            _ = v12
        } else {
            p0 -= 0
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    return 65535
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[int]string)
    var v3 string = "foo bar"
    v1 = v1
    v17 := GenS1{F0: 49, F1: 1}
    v18 := v2[42]
    for fi19 := 0; fi19 < 14; fi19++ {
        v3 = (((("\n\t" + v3) + v18) + "a") + "\n\t")
        _ = fi19
    }
    v2[-754] = v18
    var v20 uint16 = 1
    v21 := GenS1{F0: (v20 & (^(v20))), F1: 65535}
    fmt.Println(v3)
    var v22 int64 = 775
    var v23 int64 = v22
    v24 := v2[42]
    var v25 int8 = 42
    delete(v2, -915)
    v20 = 628
    v26 := GenS1{F0: v20, F1: 18446744073709551615}
    fmt.Println(v3, v24, v1)
    v22 -= v23
    for fi27 := 0; fi27 < 1; fi27++ {
        var v28 uint = 0
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}

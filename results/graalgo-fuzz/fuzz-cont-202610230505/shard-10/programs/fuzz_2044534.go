package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int }
func genFn0(p0 uint8, p1 int) uint64 {
    fmt.Println(p0)
    p0 += p0
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 float32 = -2.25
    var v2 float32 = 1e10
    var v3 int = 0
    var v4 string = "a"
    v1 += 1e10
    var v5 float64 = ((0.1 + 1.5) * 1.5)
    var v6 int = v3
    switch 9223372036854775807 {
    case 681:
        if false {
            if true {
                v4 = (v4 + v4)
                v7 := []int{v3, v3, 32767}
                _ = v7
            }
            if (!(1 < 2)) {
                v8 := []int{32768}
                _ = v8
            }
            if (-94 == 6) {
                v9 := GenS1{F0: ((42 < 5) || (!(!true))), F1: 715}
                _ = v9
            }
        }
    case 9223372036854775807:
        v1 = v1
    case -792:
        var v10 bool = (((false && false) || (!(false || (-29 < 15)))) && true)
        _ = v10
    default:
        switch ((v4 + v4) + v4) {
        case "foo bar":
            v2 = 0.1
        case "\n\t":
            if false {
                if false {
                    for fi11 := 0; fi11 < 3; fi11++ {
                        fmt.Println(fi11, v4, v4)
                        _ = fi11
                    }
                }
            } else {
                fmt.Println(v6, v1)
                v12 := GenS2{F0: 4294967295}
                _ = v12
            }
        }
    }
    v3 -= v3
    v13 := GenS2{F0: -2147483648}
    var v14 uint = 420
    var v15 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 bool }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[string]string)
    var v2 float64 = ((3.14159 + 1e10) / (1.5 - 3.14159))
    var v3 int64 = -1
    v4 := make(map[uint32]bool)
    var v5 int16 = 1
    v3 = v3
    for fi6 := 0; fi6 < 18; fi6++ {
        fmt.Println(v3)
        v7 := v1[(("hello" + ("hello" + (("a" + "a") + ("a" + "\n\t")))) + "foo bar")]
        fmt.Println(v5)
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 100, F1: ("\n\t" + "a"), F2: (24 < -63)}
    for fi9 := 0; fi9 < 9; fi9++ {
        v2 = v2
        _ = fi9
    }
    v3 += (^(v3))
    fmt.Println(v2)
    fmt.Println(v5, v5)
    if false {
        v10 := []int{-1}
        _ = v10
    } else {
        if (v5 > (v5 << ((431) & 63))) {
            fmt.Println(v2)
            v3 = 103
            for fi11 := 0; fi11 < 15; fi11++ {
                v12 := GenS2{F0: "foo bar"}
                v5 = 42
                for fi13 := 0; fi13 < 19; fi13++ {
                    if false {
                        v5 = (-(v5))
                    }
                    delete(v4, 2147483648)
                    switch 9223372036854775807 {
                    case 1:
                        v3 = (v3 - -969)
                    default:
                        v14 := GenS2{F0: (("foo bar" + "foo bar") + "foo bar")}
                        _ = v14
                    }
                    _ = fi13
                }
                _ = fi11
                _ = v12
            }
        } else {
            var v15 uint = 757
            if (!false) {
                var v16 float32 = -2.25
                v15 += (v15 + v15)
                _ = v16
            }
            for fi17 := 0; fi17 < 15; fi17++ {
                _ = len(v1)
                _ = fi17
            }
            _ = v15
        }
    }
    v18 := GenS1{F0: 89, F1: (((((("" + ("a" + "")) + "\n\t") + "") + ("hello" + "a")) + ("hello" + ("\n\t" + ""))) + ("" + "foo bar")), F2: ((false || (v2 < (-2.25 + (v2 - v2)))) || (!(v5 >= -1)))}
    var v19 int = -1
    var v20 int64 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }

func main() {
    var v1 int16 = 1000
    v2 := make(map[uint32]bool)
    var v3 int64 = -166
    var v4 int8 = -4
    var v5 string = (("foo bar" + ((("foo bar" + ("" + (("\n\t" + "a") + "foo bar"))) + (("hello" + "hello") + ("\n\t" + "hello"))) + ("\n\t" + ("a" + "foo bar")))) + (("a" + "") + ("foo bar" + "")))
    v6 := make(map[int]string)
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 string = v5
        v9 := GenS1{F0: v1, F1: 2}
        _ = fi7
        _ = v8
        _ = v9
    }
    v4 = (v4 << ((2) & 63))
    v1 -= v1
    for fi10 := 0; fi10 < 5; fi10++ {
        v5 = v5
        if (0 >= 2) {
            v6[fi10] = v5
        } else {
            for fi11 := 0; fi11 < 9; fi11++ {
                v3 = (-(v3))
                _ = fi11
            }
            for fi12 := 0; fi12 < 15; fi12++ {
                v13 := []int{fi12, fi10}
                fmt.Println(v3)
                _ = fi12
                _ = v13
            }
        }
        _ = fi10
    }
    var v14 uint16 = 0
    v1 += v1
    switch -1 {
    case -1:
        var v15 bool = false
        _ = v15
    default:
        v16 := GenS1{F0: 1, F1: 2}
        _ = v16
    }
    var v17 int16 = -32768
    if (true || false) {
        v18 := GenS1{F0: (v1 + (^(v1))), F1: 2147483648}
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

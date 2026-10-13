package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 float32 = 0.0
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    v4 := make(map[uint32]bool)
    var v5 uint8 = 51
    var v6 bool = (65535 > 2147483648)
    v5 -= (v5 + (^(v5)))
    v5 += (v5 << ((121) & 63))
    for fi7 := 0; fi7 < 9; fi7++ {
        if v6 {
            for fi8 := 0; fi8 < 14; fi8++ {
                fmt.Println(v1, v6, v5)
                switch ("\n\t" + "hello") {
                case "a":
                    var v9 string = ("a" + ("hello" + "a"))
                    _ = v9
                case "\n\t":
                    delete(v2, (fi8 >> ((18446744073709551615) & 63)))
                case "":
                    fmt.Println(v6, v6, v5)
                default:
                    var v10 uint8 = v5
                    _ = v10
                }
                _ = fi8
            }
        }
        _ = fi7
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        for fi12 := 0; fi12 < 3; fi12++ {
            fmt.Println(v1)
            _ = fi12
        }
        _ = fi11
    }
    v13 := GenS1{F0: 0}
    if v6 {
        v6 = (v6 && v6)
    } else {
        var v14 float32 = (float32((0.1 + -2.25)) * ((v1 * (0.0 * ((3.14159 + ((v1 - v1) + ((0.0 * ((0.1 + v1) + v1)) + 0.0))) / v1))) / (float32((0.1 + 1e10)) + (float32(float64(((0.1 * 0.1) - -2.25))) - v1))))
        _ = v14
    }
    var v15 int64 = -1
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 uint = 1
        _ = fi16
        _ = v17
    }
    fmt.Println(v5)
    if (v5 <= (^(v5))) {
        v5 = (v5 * (v5 * (v5 + (v5 + (v5 * (^(v5)))))))
    }
    var v18 uint = 2
    var v19 int16 = -1
    v15 = (v15 + 1)
    var v20 int32 = 2147483647
    switch (("\n\t" + ((("" + "\n\t") + (("hello" + "foo bar") + (("hello" + (("foo bar" + "\n\t") + "")) + (("\n\t" + ("foo bar" + "a")) + "foo bar")))) + (((("foo bar" + ("\n\t" + (((("hello" + "hello") + "hello") + ("\n\t" + "foo bar")) + ""))) + "") + "\n\t") + (("\n\t" + "hello") + ("foo bar" + (("\n\t" + "") + "hello")))))) + "\n\t") {
    case "\n\t":
        var v21 int64 = 1
        _ = v21
    default:
        v5 -= v5
    }
    v19 -= -916
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

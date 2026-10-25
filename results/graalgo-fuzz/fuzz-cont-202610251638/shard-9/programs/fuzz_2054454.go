package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
func genFn0(p0 int64) uint {
    var v6 uint32 = 1
    v7 := []int{1, 32768, -1}
    var v8 int16 = 807
    fmt.Println(v6, v6)
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 10
}
func genFn1() string {
    v9 := GenS1{F0: 42, F1: 926}
    if (genFn0(1) > 579) {
        var v10 string = ""
        if true {
            v10 = v10
        }
        if false {
            v11 := GenS1{F0: 2147483647, F1: 1}
            if false {
                for fi12 := 0; fi12 < 11; fi12++ {
                    v10 = "foo bar"
                    var v13 uint64 = 0
                    _ = fi12
                    _ = v13
                }
            }
            v14 := []int{32767, -965, 32767}
            _ = v11
            _ = v14
        } else {
            v15 := []int{-465, 772, 32767, 2147483647}
            if true {
                fmt.Println(v10)
            }
            _ = v15
        }
        _ = v10
    }
    v16 := []int{-290}
    var v17 float64 = 0.0
    _ = v16
    _ = v17
    _ = v9
    return ((("foo bar" + (("\n\t" + "hello") + ("\n\t" + "foo bar"))) + ("foo bar" + "a")) + "\n\t")
}

func main() {
    var v1 uint8 = 155
    v2 := []int{127, -525, 32768, 128}
    var v3 string = "a"
    v4 := []int{318, 149}
    var v5 uint = 372
    var v18 int64 = -1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1 -= 2
    var v19 uint64 = 2
    _ = len(v4)
    var v20 string = ("hello" + v3)
    v18 = v18
    switch -9223372036854775808 {
    case -1:
        v18 -= (v18 + (v18 ^ v18))
    case -875:
        var v21 int8 = -128
        _ = v21
    case 42:
        var v22 uint16 = 920
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}

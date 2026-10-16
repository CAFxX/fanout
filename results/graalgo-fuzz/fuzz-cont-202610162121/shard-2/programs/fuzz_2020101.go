package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64 }
func genFn0() bool {
    var v6 int64 = 1
    fmt.Println(v6)
    v7 := GenS1{F0: 42}
    _ = v6
    _ = v7
    return (false && ((!(4294967295 >= 0)) || false))
}
func genFn1(p0 int64, p1 uint64) int64 {
    p0 = p0
    var v8 int32 = 1
    if false {
        p1 += (p1 * p1)
    } else {
        fmt.Println(v8, v8, v8)
        if true {
            p0 += -455
            v8 -= (v8 >> ((402) & 63))
            p1 = p1
        } else {
            switch ("a" + "foo bar") {
            case "hello":
                var v9 uint8 = 255
                _ = v9
            case "\n\t":
                for fi10 := 0; fi10 < 20; fi10++ {
                    v11 := GenS2{F0: (p0 * (p0 / ((p0) | 1)))}
                    v12 := []int{}
                    fmt.Println(p1, fi10)
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            case "":
                var v13 int16 = 0
                _ = v13
            }
            v14 := []int{}
            _ = v14
        }
        v15 := []int{}
        _ = v15
    }
    p0 = (p0 ^ (p0 / ((p0) | 1)))
    _ = p0
    _ = p1
    _ = v8
    return (^(p0))
}

func main() {
    var v1 uint8 = 42
    var v2 int8 = 1
    var v3 float32 = 3.14159
    v4 := []uint32{65535, 991}
    var v5 int = 32768
    v16 := GenS2{F0: -545}
    for fi17 := 0; fi17 < 15; fi17++ {
        _ = len(v4)
        _ = fi17
    }
    var v18 uint8 = (v1 * (v1 - (v1 % ((v1) | 1))))
    fmt.Println(v18, v5)
    v18 += v18
    v19 := GenS2{F0: -9223372036854775808}
    var v20 string = ("\n\t" + "\n\t")
    fmt.Println(v18, v20, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}

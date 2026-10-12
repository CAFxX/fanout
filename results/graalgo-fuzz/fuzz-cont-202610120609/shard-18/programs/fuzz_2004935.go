package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 int = 0
    v2 := make(map[uint32]uint32)
    var v3 int = v1
    var v4 float32 = 3.14159
    v5 := make(map[uint32]int)
    v6 := GenS1{F0: 4294967295}
    v7 := GenS2{F0: 2}
    v8 := []int{}
    v4 = v4
    v3 = v3
    for fi9 := 0; fi9 < 1; fi9++ {
        v10 := GenS2{F0: 453}
        _ = fi9
        _ = v10
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    if false {
        v1 -= v1
    }
    var v11 string = (((("hello" + ("hello" + "a")) + ("foo bar" + ((("" + "hello") + "\n\t") + (("foo bar" + ("a" + "foo bar")) + "")))) + ("a" + (("\n\t" + "a") + "\n\t"))) + "")
    var v12 uint64 = 1
    fmt.Println(v12, v12)
    var v13 float64 = ((-2.25 - ((0.1 + 1.5) + float64(v4))) - 1.5)
    v14 := GenS1{F0: (^(v12))}
    var v15 int16 = -391
    if (!true) {
        for fi16 := 0; fi16 < 20; fi16++ {
            for fi17 := 0; fi17 < 12; fi17++ {
                v4 = v4
                delete(v2, 1)
                _ = fi17
            }
            for fi18 := 0; fi18 < 3; fi18++ {
                v12 = v12
                _ = fi18
            }
            _ = fi16
        }
    } else {
        fmt.Println(v15, v1)
        v19 := GenS2{F0: 756}
        _ = v19
    }
    fmt.Println(v11, v3, v15)
    v4 = v4
    var v20 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

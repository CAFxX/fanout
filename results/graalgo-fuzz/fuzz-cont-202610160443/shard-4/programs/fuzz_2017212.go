package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32 }

func main() {
    var v1 uint64 = 0
    v2 := []int{1, -1}
    var v3 int = 32768
    var v4 uint32 = 4294967295
    var v5 uint = 1
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := GenS1{F0: ""}
        fmt.Println(fi6, fi6)
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: ("foo bar" + ("\n\t" + "hello"))}
    if (511 > -811) {
        if (true && true) {
            v9 := GenS1{F0: "hello"}
            v1 -= (v1 / (((v1 | (v1 + (v1 - v1)))) | 1))
            _ = v9
        }
    } else {
        v1 -= v1
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v10 uint8 = 1
    v11 := GenS2{F0: 920}
    var v12 uint32 = (^(v4))
    fmt.Println(v12)
    for fi13 := 0; fi13 < 18; fi13++ {
        v5 -= (v5 / ((v5) | 1))
        _ = fi13
    }
    var v14 int64 = 1
    var v15 int16 = 180
    v4 = v12
    v4 += (v4 + v4)
    var v16 int64 = v14
    var v17 int16 = -1
    v2 = append(v2, 32768)
    var v18 uint32 = (v12 % ((v12) | 1))
    for fi19 := 0; fi19 < 11; fi19++ {
        fmt.Println(v16, v15, v18)
        for fi20 := 0; fi20 < 20; fi20++ {
            v16 = v16
            var v21 string = ("\n\t" + (("foo bar" + (((("hello" + "hello") + "a") + ("hello" + "")) + "")) + ("hello" + "a")))
            _ = fi20
            _ = v21
        }
        fmt.Println(v5)
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}

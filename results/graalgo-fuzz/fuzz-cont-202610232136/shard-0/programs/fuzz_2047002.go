package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
func genFn0() uint64 {
    if true {
        for fi6 := 0; fi6 < 1; fi6++ {
            if (!false) {
                fmt.Println(fi6)
            }
            var v7 int16 = -32768
            _ = fi6
            _ = v7
        }
        v8 := []int{}
        _ = v8
    } else {
        v9 := []int{}
        var v10 int = 859
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 65}
    _ = v11
    return 38
}

func main() {
    var v1 uint = 2
    v2 := make(map[int]uint32)
    v3 := make(map[string]int)
    v4 := make(map[uint32]int)
    var v5 float64 = 1e10
    var v12 string = (("" + "a") + (("hello" + ("a" + ("a" + ("hello" + "a")))) + "\n\t"))
    if false {
        var v13 float64 = v5
        v14 := []int{-2147483648}
        v15 := GenS1{F0: 157}
        _ = v13
        _ = v14
        _ = v15
    } else {
        v5 = (v5 - v5)
    }
    v16 := v3[("\n\t" + v12)]
    v17 := GenS1{F0: 42}
    var v18 float64 = (v5 + (v5 + float64((float32(v5) / -2.25))))
    if true {
        v16 = v16
        for fi19 := 0; fi19 < 15; fi19++ {
            var v20 int64 = 205
            _ = fi19
            _ = v20
        }
    } else {
        v21 := GenS2{F0: 1}
        v5 = 0.1
        _ = v21
    }
    fmt.Println(v18)
    v22 := []int{}
    v12 = "foo bar"
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    v16 = v16
    var v23 int = 2147483647
    v16 -= (v16 + -1)
    v23 -= v16
    for fi24 := 0; fi24 < 8; fi24++ {
        v3["a"] = (v16 & (v23 >> ((v1) & 63)))
        for fi25 := 0; fi25 < 18; fi25++ {
            v26 := v3["foo bar"]
            _ = len(v22)
            _ = fi25
            _ = v26
        }
        fmt.Println(fi24, fi24, v12)
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}

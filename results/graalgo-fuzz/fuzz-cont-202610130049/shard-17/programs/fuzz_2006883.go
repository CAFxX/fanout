package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 int64; F2 int64 }
func genFn0(p0 uint16, p1 uint16) int8 {
    var v5 uint = 390
    fmt.Println(v5, p0, p0)
    _ = p0
    _ = p1
    _ = v5
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 0
    var v3 uint32 = 1
    var v4 uint64 = v2
    var v6 string = ("" + (("hello" + "a") + "hello"))
    var v7 uint16 = 765
    if (255 > v7) {
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := GenS1{F0: -9223372036854775808}
            fmt.Println(v4, v4, fi8)
            var v10 uint = 2
            _ = fi8
            _ = v10
            _ = v9
        }
        v6 = ""
        v4 = v2
    }
    if false {
        v11 := GenS2{F0: true, F1: 42, F2: -435}
        for fi12 := 0; fi12 < 10; fi12++ {
            v1[((v6 + ((v6 + (("a" + ("hello" + "foo bar")) + "hello")) + v6)) + v6)] = (^(fi12))
            _ = fi12
        }
        v13 := []int{-666}
        _ = v11
        _ = v13
    }
    v1["\n\t"] = -747
    var v14 int16 = -32768
    v4 -= v2
    v15 := GenS1{F0: -1}
    v6 = ("\n\t" + v6)
    v4 += 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint, p1 bool, p2 uint16) uint32 {
    switch (("foo bar" + "a") + "hello") {
    case "a":
        v7 := []int{}
        _ = v7
    case "foo bar":
        p1 = p1
    default:
        fmt.Println(p0, p0, p1)
    }
    var v8 uint64 = 0
    p1 = ((v8 | 19) == v8)
    var v9 int64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    var v1 bool = (255 == 2)
    var v2 float64 = (1e10 + 3.14159)
    v3 := make(map[uint32]int)
    var v4 int32 = 2147483647
    v5 := []string{"\n\t", "\n\t", "foo bar", "hello"}
    var v6 string = (("foo bar" + "foo bar") + "hello")
    v6 = v6
    v10 := GenS1{F0: 32767}
    v2 += v2
    v11 := GenS1{F0: -9223372036854775808}
    switch (v6 + v6) {
    case "hello":
        for fi12 := 0; fi12 < 4; fi12++ {
            if true {
                for fi13 := 0; fi13 < 9; fi13++ {
                    v14 := v3[945]
                    _ = fi13
                    _ = v14
                }
            }
            _ = fi12
        }
    }
    var v15 bool = v1
    for fi16 := 0; fi16 < 10; fi16++ {
        for fi17 := 0; fi17 < 15; fi17++ {
            v18 := GenS2{F0: 1}
            if len(v5) > 0 {
                v5[0] = (("a" + v6) + v6)
            }
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 18; fi19++ {
            v20 := GenS2{F0: 42}
            _ = fi19
            _ = v20
        }
        v4 = v4
        _ = fi16
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        for fi22 := 0; fi22 < 12; fi22++ {
            if len(v5) > 0 {
                v5[0] = "foo bar"
            }
            _ = len(v5)
            var v23 int32 = (v4 % (((v4 / ((v4) | 1))) | 1))
            _ = fi22
            _ = v23
        }
        var v24 string = v6
        _ = fi21
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

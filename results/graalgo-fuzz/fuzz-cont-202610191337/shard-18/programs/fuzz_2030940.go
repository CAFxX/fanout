package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint16, p1 bool) uint16 {
    v5 := []int{}
    var v6 uint = 267
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return p0
}

func main() {
    var v1 float64 = 1e10
    v2 := []int{}
    v3 := []int{0, -909, 518, -9223372036854775808}
    v4 := make(map[uint32]string)
    var v7 int32 = 42
    v7 = (v7 % ((v7) | 1))
    switch "hello" {
    case "\n\t":
        fmt.Println(v1, v1)
    case "foo bar":
        v1 -= v1
    }
    v8 := GenS1{F0: 276}
    switch ((("foo bar" + ("" + "\n\t")) + "") + "foo bar") {
    case "a":
        var v9 bool = (18446744073709551615 <= 922)
        _ = v9
    case "foo bar":
        var v10 float64 = 1e10
        _ = v10
    }
    fmt.Println(v1, v1)
    for fi11 := 0; fi11 < 11; fi11++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 uint32 = 1
            v1 = v1
            if false {
                v1 += v1
                _ = len(v3)
                if true {
                    switch fi12 {
                    case 2147483647:
                        fmt.Println(v1, v7)
                    case 42:
                        fmt.Println(v1, fi11, fi11)
                    }
                } else {
                    v14 := v4[v13]
                    var v15 uint64 = 1
                    v14 = ""
                    _ = v14
                    _ = v15
                }
                fmt.Println(fi11, fi12)
            }
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    var v16 string = "hello"
    var v17 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}

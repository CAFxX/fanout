package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 string) int8 {
    for fi5 := 0; fi5 < 17; fi5++ {
        v6 := GenS1{F0: 255}
        if true {
            p0 = ((p0 + "a") + ("\n\t" + "hello"))
        } else {
            v7 := []int{(fi5 & -1), (fi5 % ((128) | 1)), -9223372036854775808, 32768}
            fmt.Println(p0, fi5)
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    p0 = "foo bar"
    _ = p0
    return 97
}

func main() {
    var v1 int64 = 989
    var v2 bool = false
    v3 := []int{128, 823, -9223372036854775808}
    v4 := []string{}
    _ = len(v4)
    var v8 int32 = 1
    fmt.Println(v2, v8)
    if v2 {
        switch "" {
        case "":
            switch "\n\t" {
            case "":
                var v9 int8 = -1
                _ = v9
            case "hello":
                for fi10 := 0; fi10 < 14; fi10++ {
                    fmt.Println(v8)
                    _ = fi10
                }
            case "a":
                fmt.Println(v1, v8, v2)
            }
        }
        for fi11 := 0; fi11 < 20; fi11++ {
            v12 := GenS1{F0: 461}
            fmt.Println(fi11, v8, v2)
            v13 := GenS1{F0: 0}
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    fmt.Println(v1, v1)
    var v14 uint16 = 255
    var v15 bool = v2
    v15 = v2
    var v16 uint8 = 91
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}

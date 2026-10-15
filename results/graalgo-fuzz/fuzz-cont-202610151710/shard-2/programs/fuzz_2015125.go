package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int64, p1 int64) int16 {
    switch "a" {
    case "foo bar":
        if true {
            var v5 int16 = 32767
            switch (("hello" + "hello") + ("hello" + "foo bar")) {
            case "hello":
                p1 += int64(v5)
            }
            _ = v5
        } else {
            v6 := GenS1{F0: 4294967295}
            switch "a" {
            case "":
                switch 38 {
                case -2147483648:
                    v7 := GenS1{F0: 65535}
                    _ = v7
                default:
                    if (true && true) {
                        for fi8 := 0; fi8 < 14; fi8++ {
                            p1 = (p0 | p1)
                            var v9 uint8 = 255
                            _ = fi8
                            _ = v9
                        }
                    }
                }
            default:
                v10 := []int{-1, -44}
                _ = v10
            }
            _ = v6
        }
    }
    v11 := []int{}
    if true {
        if len(v11) > 0 {
            v11[0] = 1
        }
        var v12 float32 = float32(3.14159)
        _ = v12
    } else {
        v11 = append(v11, 32767)
        var v13 int = 4294967295
        _ = v13
    }
    p1 = p1
    _ = len(v11)
    _ = p0
    _ = p1
    _ = v11
    return 0
}

func main() {
    v1 := make(map[int]string)
    var v2 int32 = -1
    var v3 float32 = -2.25
    v4 := []string{("hello" + "")}
    var v14 bool = (!false)
    v3 += v3
    fmt.Println(v3, v3, v2)
    v1[-135] = "hello"
    var v15 string = "hello"
    var v16 int16 = 0
    var v17 uint64 = 25
    v18 := GenS1{F0: 0}
    _ = len(v1)
    if (!(42 >= 42)) {
        v19 := GenS1{F0: 24}
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
func genFn0() uint64 {
    var v5 uint64 = 696
    v5 = (^(v5))
    var v6 uint32 = 0
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 bool, p1 bool, p2 uint16) int32 {
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(p0)
        var v8 string = "a"
        v9 := []int{}
        _ = fi7
        _ = v8
        _ = v9
    }
    p0 = p0
    var v10 int16 = -1
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 uint64 = genFn0()
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 1
}

func main() {
    v1 := []string{""}
    v2 := make(map[uint32]bool)
    var v3 int64 = 809
    v4 := make(map[int]int)
    var v13 bool = false
    var v14 string = ("a" + (("\n\t" + "") + ("foo bar" + "hello")))
    v3 += (v3 ^ v3)
    var v15 uint16 = 32767
    v16 := v4[-378]
    var v17 int = v16
    v18 := GenS1{F0: "foo bar", F1: 206}
    if v13 {
        switch (v16 - v16) {
        case 0:
            if v13 {
                switch "a" {
                case "\n\t":
                    switch v14 {
                    case "foo bar":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                }
                var v19 int32 = genFn1(((v15 / ((121) | 1)) == (v15 << ((42) & 63))), (((v15 + v15) == 341) && (1e10 != -2.25)), 65535)
                v14 = "a"
                _ = v19
            } else {
                for fi20 := 0; fi20 < 11; fi20++ {
                    var v21 uint16 = (v15 % ((v15) | 1))
                    _ = fi20
                    _ = v21
                }
                v16 = (v16 ^ (v16 >> ((65535) & 63)))
            }
        case 2147483647:
            v14 = v14
        case 4294967295:
            var v22 float32 = 1e10
            _ = v22
        }
    } else {
        v1 = append(v1, "hello")
        fmt.Println(v14, v17, v3)
    }
    v23 := GenS1{F0: "", F1: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}

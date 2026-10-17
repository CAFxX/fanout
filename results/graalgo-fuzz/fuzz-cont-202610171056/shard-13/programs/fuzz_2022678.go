package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 uint8 = 0
    var v2 float64 = 3.14159
    v3 := []uint32{42, 4294967295}
    v4 := make(map[uint32]string)
    v5 := []string{}
    var v6 uint16 = 42
    v7 := GenS1{F0: (506 >= 604), F1: 1}
    v4[65535] = "foo bar"
    var v8 int32 = 0
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 uint16 = 721
        _ = fi9
        _ = v10
    }
    if len(v5) > 0 {
        v5[0] = "a"
    }
    if false {
        var v11 int = int(v1)
        _ = v11
    }
    v2 = v2
    _ = len(v4)
    v12 := v4[269]
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 uint64 = 2
        v15 := GenS1{F0: true, F1: 158}
        _ = fi13
        _ = v14
        _ = v15
    }
    v16 := GenS2{F0: 42}
    v1 = 1
    if true {
        if (false && false) {
            v4[2147483648] = v12
        }
        v12 = ((v12 + "hello") + v12)
    }
    if false {
        var v17 int64 = -1
        var v18 int16 = 0
        for fi19 := 0; fi19 < 2; fi19++ {
            var v20 string = (v12 + "hello")
            if true {
                switch int(v8) {
                case 2147483647:
                    v21 := GenS1{F0: (v1 != v1), F1: int64(fi19)}
                    _ = v21
                case 9223372036854775807:
                    for fi22 := 0; fi22 < 3; fi22++ {
                        v5 = append(v5, (v20 + ""))
                        fmt.Println(fi19, v2)
                        _ = fi22
                    }
                case -2147483648:
                    v17 += v17
                }
                v20 = v12
            } else {
                var v23 int = (fi19 % ((fi19) | 1))
                _ = v23
            }
            v24 := GenS2{F0: 1}
            _ = fi19
            _ = v20
            _ = v24
        }
        _ = v17
        _ = v18
    } else {
        if false {
            v25 := GenS1{F0: true, F1: 1}
            _ = v25
        }
        v2 += (-2.25 / v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

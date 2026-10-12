package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int64 = 1
    var v2 uint32 = 763
    v3 := make(map[int]bool)
    v2 = v2
    var v4 int = 878
    var v5 uint8 = 2
    if true {
        var v6 string = (("hello" + ("" + (("\n\t" + "hello") + "a"))) + "")
        v4 = (v4 * v4)
        v6 = v6
        _ = v6
    }
    v3[2147483647] = (((v1 * v1) != v1) && true)
    if (-103 != 1) {
        var v7 int8 = 127
        v8 := v3[(v4 | v4)]
        _ = v7
        _ = v8
    }
    _ = len(v3)
    delete(v3, v4)
    var v9 int64 = (v1 - 797)
    for fi10 := 0; fi10 < 1; fi10++ {
        switch -9223372036854775808 {
        case 2147483647:
            var v11 int8 = 16
            _ = v11
        case 9223372036854775807:
            v12 := GenS1{F0: (^(v2))}
            _ = v12
        case 1:
            for fi13 := 0; fi13 < 5; fi13++ {
                var v14 uint16 = 32767
                v4 = (v4 / ((v4) | 1))
                v5 = 223
                _ = fi13
                _ = v14
            }
        }
        _ = fi10
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        for fi16 := 0; fi16 < 6; fi16++ {
            var v17 int32 = 2147483647
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    v5 = (v5 & v5)
    if false {
        v2 = v2
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        var v19 bool = true
        var v20 uint16 = 839
        _ = fi18
        _ = v19
        _ = v20
    }
    delete(v3, 1)
    var v21 int16 = 750
    for fi22 := 0; fi22 < 3; fi22++ {
        if false {
            v23 := GenS1{F0: v2}
            _ = v23
        }
        v24 := v3[(-(v4))]
        _ = fi22
        _ = v24
    }
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int8 }

func main() {
    var v1 uint32 = 812
    var v2 int64 = 1
    var v3 float32 = -2.25
    for fi4 := 0; fi4 < 4; fi4++ {
        if (!(false && (4294967295 <= 2))) {
            for fi5 := 0; fi5 < 10; fi5++ {
                var v6 int64 = v2
                fmt.Println(fi4)
                _ = fi5
                _ = v6
            }
            var v7 uint16 = 568
            _ = v7
        } else {
            v8 := GenS1{F0: 42, F1: true, F2: -102}
            v9 := GenS1{F0: 127, F1: (false || false), F2: 48}
            _ = v8
            _ = v9
        }
        _ = fi4
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := []int{(fi10 | fi10), (fi10 | 2147483647), fi10}
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: 127, F1: false, F2: 0}
    for fi13 := 0; fi13 < 18; fi13++ {
        fmt.Println(v3)
        _ = fi13
    }
    v3 += float32(-2.25)
    v14 := GenS1{F0: 255, F1: ((!(!true)) && (!false)), F2: 116}
    var v15 string = "\n\t"
    fmt.Println(v1, v2, v15)
    var v16 uint32 = (v1 >> ((4294967295) & 63))
    v17 := []int{127, 9223372036854775807}
    v18 := GenS1{F0: 138, F1: true, F2: 42}
    var v19 int64 = (v2 % ((-9223372036854775808) | 1))
    v17 = append(v17, 2147483647)
    var v20 int64 = v19
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}

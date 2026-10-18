package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 uint32 = 105
    v2 := make(map[string]bool)
    var v3 uint = 1
    var v4 int16 = -32768
    v5 := GenS1{F0: 255, F1: -128, F2: 0}
    for fi6 := 0; fi6 < 10; fi6++ {
        v1 = v1
        _ = fi6
    }
    if false {
        if (!(127 != 42)) {
            v1 = v1
            fmt.Println(v4, v1, v1)
        }
    }
    var v7 uint = v3
    if (!(!true)) {
        fmt.Println(v3, v1, v7)
        v8 := []int{42, 1, 241, 32768}
        _ = v8
    }
    var v9 uint64 = 2
    if true {
        if ((^(v9)) <= (v9 << ((v7) & 63))) {
            var v10 bool = (!true)
            v11 := GenS2{F0: 128, F1: 0}
            _ = v10
            _ = v11
        }
        switch 128 {
        case -2147483648:
            fmt.Println(v3, v4)
        case 2147483647:
            _ = len(v2)
        }
        v3 += v3
    } else {
        var v12 int8 = 127
        for fi13 := 0; fi13 < 8; fi13++ {
            if (!(false || (105 >= (v12 & v12)))) {
                if true {
                    var v14 float32 = -2.25
                    v12 += (v12 ^ (v12 ^ 0))
                    _ = v14
                }
                v2[""] = false
                if true {
                    v4 = (v4 & 905)
                } else {
                    v15 := GenS3{F0: 885}
                    v7 -= (v3 >> ((v7) & 63))
                    _ = v15
                }
            }
            var v16 int16 = v4
            _ = fi13
            _ = v16
        }
        _ = v12
    }
    v17 := []int{1, 4294967295}
    var v18 string = "foo bar"
    if (127 < 0) {
        if true {
            v19 := GenS2{F0: 77, F1: 1}
            _ = v19
        }
        for fi20 := 0; fi20 < 17; fi20++ {
            if len(v17) > 0 {
                v17[0] = (fi20 << ((v3) & 63))
            }
            _ = fi20
        }
        v2["hello"] = false
    } else {
        v1 += v1
        v1 += (v1 & (v1 >> ((316) & 63)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}

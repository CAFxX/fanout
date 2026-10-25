package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 int16 }
func genFn0(p0 int64, p1 int64) int {
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := []int{fi6, fi6, (-(fi6))}
        _ = fi6
        _ = v7
    }
    var v8 string = ""
    if true {
        var v9 int32 = -2147483648
        _ = v9
    }
    var v10 uint8 = 255
    fmt.Println(v10, v10)
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    return 2147483647
}

func main() {
    v1 := []int{-2147483648}
    var v2 string = ("hello" + ((("hello" + "a") + "\n\t") + "foo bar"))
    var v3 uint8 = 0
    var v4 int64 = 1
    v5 := make(map[string]uint32)
    if (false || (1e10 >= 3.14159)) {
        var v11 uint32 = 931
        var v12 int64 = (-(v4))
        _ = v11
        _ = v12
    }
    var v13 int32 = -2147483648
    for fi14 := 0; fi14 < 3; fi14++ {
        v2 = v2
        _ = fi14
    }
    v15 := v5["\n\t"]
    fmt.Println(v4, v4, v4)
    fmt.Println(v15, v3, v13)
    _ = len(v5)
    for fi16 := 0; fi16 < 16; fi16++ {
        v17 := GenS1{F0: 343}
        v18 := GenS2{F0: v3, F1: 32767}
        if (!true) {
            fmt.Println(fi16, fi16, v15)
            if (true && ((888 >= 550) && true)) {
                fmt.Println(v15, v4, v13)
                if len(v1) > 0 {
                    v1[0] = 9223372036854775807
                }
                _ = len(v5)
            }
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 uint32 = v15
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

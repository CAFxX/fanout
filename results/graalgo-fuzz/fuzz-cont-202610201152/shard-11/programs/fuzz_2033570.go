package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int64 }
func genFn0(p0 int8) uint {
    p0 = p0
    if false {
        var v7 int64 = 0
        _ = v7
    } else {
        p0 = (^(p0))
    }
    switch (("a" + "hello") + "") {
    case "hello":
        if true {
            switch ("" + "\n\t") {
            case "a":
                if false {
                    fmt.Println(p0)
                } else {
                    var v8 uint8 = 77
                    v9 := GenS1{F0: "hello", F1: 42}
                    _ = v8
                    _ = v9
                }
            }
        }
    case "a":
        p0 += 0
    }
    var v10 string = "\n\t"
    v11 := []int{}
    _ = p0
    _ = v10
    _ = v11
    return 65535
}

func main() {
    var v1 int32 = 371
    var v2 float64 = ((0.1 - 1.5) * (0.0 - 3.14159))
    v3 := []int{}
    var v4 int64 = int64(v1)
    var v5 uint16 = 0
    var v6 float64 = 3.14159
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v12 uint8 = 42
    _ = len(v3)
    for fi13 := 0; fi13 < 14; fi13++ {
        fmt.Println(v5, fi13, v12)
        var v14 uint = 50
        if ((!false) || (42 > 0)) {
            v6 = (float64(0.1) * (float64((3.14159 * float32(v2))) / (v2 + (v6 * -2.25))))
            fmt.Println(fi13, v12, fi13)
        }
        _ = fi13
        _ = v14
    }
    v6 = (((v2 - float64(3.14159)) - 0.1) + float64(0.1))
    switch 886 {
    case 497:
        v3 = append(v3, 127)
    case -1:
        fmt.Println(v4, v2)
    case 127:
        v15 := GenS1{F0: "", F1: 4294967295}
        _ = v15
    }
    fmt.Println(v6)
    if (!true) {
        fmt.Println(v6, v6)
        v3 = append(v3, 127)
    }
    v16 := GenS2{F0: (v5 - (v5 << ((99) & 63))), F1: v4}
    var v17 uint32 = 2
    var v18 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

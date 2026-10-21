package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint32 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int32, p1 bool) bool {
    p0 = -2147483648
    if (100 > 42) {
        if ((0.0 + 1.5) <= 1e10) {
            p1 = (-9223372036854775808 != 4294967295)
            var v5 int32 = p0
            p1 = false
            _ = v5
        }
    }
    v6 := []int{819, 1}
    _ = p0
    _ = p1
    _ = v6
    return (255 >= 1)
}
func genFn1() uint {
    if true {
        v7 := GenS1{F0: 42, F1: 4294967295, F2: 0}
        var v8 int32 = 0
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 int64 = 42
            v10 += v10
            _ = fi9
            _ = v10
        }
        var v11 bool = (-890 < -1)
        _ = v11
    }
    var v12 bool = false
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 uint16 = 32767
        _ = fi13
        _ = v14
    }
    _ = v12
    return 4294967295
}

func main() {
    var v1 string = "\n\t"
    var v2 int64 = 9223372036854775807
    v3 := []string{"hello", ""}
    var v4 uint = 1
    v3 = append(v3, (v1 + (v1 + "\n\t")))
    _ = len(v3)
    var v15 uint8 = uint8(v4)
    v15 += (v15 + v15)
    v2 = 1
    fmt.Println(v15, v1)
    v2 = v2
    if false {
        v15 += (v15 | 127)
        if (1 == -1) {
            fmt.Println(v15, v4)
        }
    } else {
        switch "" {
        case "foo bar":
            switch (v1 + (((v1 + ("\n\t" + ((v1 + "foo bar") + "a"))) + (v1 + (((v1 + ("" + "foo bar")) + v1) + v1))) + v1)) {
            case "foo bar":
                v3 = append(v3, v1)
            default:
                if ((127 <= 42) || (true && (!(v4 < 4294967295)))) {
                    v16 := GenS2{F0: 571}
                    _ = v16
                }
            }
        }
        var v17 int64 = v2
        _ = v17
    }
    v18 := GenS2{F0: uint64(v15)}
    var v19 bool = false
    v20 := GenS1{F0: 9223372036854775807, F1: -9223372036854775808, F2: 4294967295}
    v4 = genFn1()
    var v21 int64 = v2
    v19 = v19
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}

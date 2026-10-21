package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 string }
type GenS2 struct { F0 uint8; F1 int32; F2 uint }
func genFn0() int32 {
    var v5 int = 42
    for fi6 := 0; fi6 < 8; fi6++ {
        v5 += (v5 % ((fi6) | 1))
        v5 = 32767
        var v7 uint16 = 32767
        _ = fi6
        _ = v7
    }
    fmt.Println(v5)
    _ = v5
    return 42
}

func main() {
    var v1 string = "a"
    var v2 int8 = -36
    var v3 uint32 = 65535
    var v4 int64 = 9223372036854775807
    v8 := []int{-1, 127}
    var v9 uint64 = 42
    v2 = 127
    if true {
        if (997 != uint16(v3)) {
            for fi10 := 0; fi10 < 5; fi10++ {
                var v11 int64 = -9223372036854775808
                _ = fi10
                _ = v11
            }
            if len(v8) > 0 {
                v8[0] = -524
            }
        }
    } else {
        _ = len(v8)
    }
    if true {
        var v12 int8 = (v2 / ((-90) | 1))
        var v13 string = (v1 + ("hello" + (v1 + (v1 + (("a" + v1) + "\n\t")))))
        fmt.Println(v9, v12)
        _ = v12
        _ = v13
    } else {
        fmt.Println(v2)
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        fmt.Println(fi14, v4)
        if false {
            v1 = v1
            v9 = v9
            v8 = append(v8, fi14)
        } else {
            var v15 float32 = 0.0
            if false {
                for fi16 := 0; fi16 < 3; fi16++ {
                    v15 += float32(3.14159)
                    _ = fi16
                }
                v9 -= (v9 / ((v9) | 1))
                for fi17 := 0; fi17 < 3; fi17++ {
                    switch "foo bar" {
                    case "foo bar":
                        var v18 int32 = 1
                        _ = v18
                    case "\n\t":
                        fmt.Println(v15, v9)
                    }
                    var v19 uint16 = 32767
                    v20 := GenS1{F0: false, F1: (v3 * v3), F2: v1}
                    _ = fi17
                    _ = v19
                    _ = v20
                }
            } else {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            }
            _ = v15
        }
        v21 := GenS2{F0: 185, F1: 481, F2: 18446744073709551615}
        _ = fi14
        _ = v21
    }
    var v22 int64 = (v4 << ((309) & 63))
    var v23 string = (v1 + (("\n\t" + v1) + (v1 + "")))
    if (true && true) {
        v3 -= v3
        v2 -= v2
    }
    v3 -= (v3 / ((v3) | 1))
    var v24 uint32 = v3
    fmt.Println(v4, v23)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}

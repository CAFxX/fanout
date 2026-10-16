package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint32 = 65535
    v2 := make(map[int]bool)
    v3 := make(map[string]string)
    var v4 uint64 = 42
    v5 := GenS1{F0: 45}
    var v6 int8 = 0
    v6 = v6
    var v7 uint = 2
    v8 := []int{-704, 128}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    if (42 < 0) {
        if ((-9223372036854775808 != 0) || (326 != 759)) {
            v9 := GenS1{F0: (v6 + (v6 << ((416) & 63)))}
            if true {
                var v10 int = -843
                _ = len(v3)
                v6 = 1
                _ = v10
            }
            _ = v9
        } else {
            v6 -= 100
            if ((!false) || ((3.14159 + 0.0) == (0.1 - 1.5))) {
                fmt.Println(v7)
            }
            if (0 <= v7) {
                if ((-11 <= -1) && false) {
                    var v11 uint = v7
                    switch int(v1) {
                    case -733:
                        var v12 float32 = 1e10
                        _ = v12
                    case 4294967295:
                        _ = len(v8)
                    }
                    var v13 float64 = 0.0
                    _ = v11
                    _ = v13
                } else {
                    var v14 int = -2147483648
                    _ = v14
                }
                var v15 uint16 = 0
                _ = v15
            }
        }
        for fi16 := 0; fi16 < 2; fi16++ {
            v17 := GenS1{F0: v6}
            v18 := v2[(fi16 / ((fi16) | 1))]
            v3["foo bar"] = ("\n\t" + "")
            _ = fi16
            _ = v17
            _ = v18
        }
    } else {
        v7 -= v7
    }
    v6 += v6
    _ = len(v2)
    if (false && true) {
        v8 = append(v8, 42)
        v19 := GenS1{F0: v6}
        _ = v19
    } else {
        var v20 bool = true
        fmt.Println(v1)
        _ = v20
    }
    v21 := v3[""]
    for fi22 := 0; fi22 < 2; fi22++ {
        v8 = append(v8, fi22)
        if len(v8) > 0 {
            v8[0] = (fi22 + int(v7))
        }
        var v23 string = v21
        _ = fi22
        _ = v23
    }
    v24 := GenS1{F0: v6}
    v25 := GenS1{F0: v6}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

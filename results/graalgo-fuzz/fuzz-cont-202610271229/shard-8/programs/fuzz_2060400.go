package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
func genFn0(p0 bool, p1 int16, p2 int32) int32 {
    var v7 uint16 = 1
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p2 + p2)
}
func genFn1(p0 int) uint32 {
    p0 -= p0
    v8 := []int{p0, p0}
    _ = len(v8)
    _ = p0
    _ = v8
    return 4294967295
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]int)
    var v3 int8 = 127
    v4 := []string{"a", "hello", "", "a"}
    v5 := make(map[string]uint32)
    var v6 int16 = 732
    v6 += 940
    var v9 bool = true
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v10 uint16 = 42
    var v11 string = ""
    if (!true) {
        var v12 string = (v11 + ((v11 + "") + v11))
        fmt.Println(v11)
        v2[2147483647] = 32767
        _ = v12
    }
    v13 := GenS1{F0: 0}
    for fi14 := 0; fi14 < 18; fi14++ {
        if (679 != 0) {
            if ((false || v9) && ((v10 / ((32767) | 1)) >= 844)) {
                _ = len(v4)
                v10 += (v10 + (^(v10)))
            } else {
                var v15 int32 = 2147483647
                if (!v9) {
                    fmt.Println(v6, v10, v11)
                    var v16 int = (fi14 & 403)
                    for fi17 := 0; fi17 < 18; fi17++ {
                        var v18 uint = 347
                        v19 := GenS2{F0: (v15 << ((42) & 63))}
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    _ = v16
                }
                _ = v15
            }
        }
        var v20 uint = 0
        _ = fi14
        _ = v20
    }
    v11 = "a"
    if len(v4) > 0 {
        v4[0] = v11
    }
    v21 := GenS1{F0: 65535}
    v11 = v11
    v4 = append(v4, ("" + "hello"))
    var v22 int32 = 1
    var v23 int32 = (v22 * (v22 | genFn0(true, -32768, v22)))
    for fi24 := 0; fi24 < 5; fi24++ {
        fmt.Println(v3, v11)
        _ = fi24
    }
    var v25 uint16 = (v10 - 0)
    v10 = v10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint8, p1 uint16) string {
    switch ("a" + "foo bar") {
    case "a":
        p0 = p0
    }
    p1 = p1
    p0 -= p0
    if (4294967295 >= 42) {
        p1 = 1
        var v6 bool = false
        _ = v6
    } else {
        var v7 string = "\n\t"
        var v8 string = v7
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    return "foo bar"
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[uint32]bool)
    var v3 int16 = -705
    var v4 bool = true
    var v5 int = -2147483648
    v9 := v2[396]
    var v10 int16 = v3
    var v11 int32 = 1
    v12 := []int{v5, -2147483648, v5}
    if (4294967295 >= 2) {
        v4 = (2 != 1)
        v13 := GenS1{F0: v5}
        _ = v13
    }
    var v14 uint32 = 2147483648
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    v10 += v10
    if (((422 > 190) || (v9 || (true && v9))) || v9) {
        for fi15 := 0; fi15 < 9; fi15++ {
            for fi16 := 0; fi16 < 16; fi16++ {
                v17 := GenS1{F0: fi16}
                v11 += v11
                v14 -= v14
                _ = fi16
                _ = v17
            }
            for fi18 := 0; fi18 < 1; fi18++ {
                if (!v9) {
                    v3 = v10
                    fmt.Println(v5, v9)
                    _ = len(v12)
                } else {
                    var v19 uint64 = 4294967295
                    v20 := GenS1{F0: 2147483647}
                    if (1 != 65535) {
                        v21 := GenS1{F0: 2147483647}
                        var v22 string = ((("" + ((("hello" + "") + "") + "hello")) + "") + "a")
                        fmt.Println(v11, fi15)
                        var v23 uint8 = 128
                        var v24 int64 = v1
                        _ = v21
                        _ = v22
                        _ = v23
                        _ = v24
                    } else {
                        fmt.Println(v14, fi15, v11)
                    }
                    v3 += (v3 >> ((758) & 63))
                    _ = v19
                    _ = v20
                }
                _ = fi18
            }
            v25 := GenS1{F0: v5}
            _ = fi15
            _ = v25
        }
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        fmt.Println(v9, v5, v5)
    } else {
        var v26 int = (v5 % ((-9223372036854775808) | 1))
        v2[(v14 / (((v14 & (v14 >> ((675) & 63)))) | 1))] = ("" > ("a" + ("\n\t" + "hello")))
        _ = v26
    }
    v9 = v9
    _ = len(v2)
    v4 = v4
    v10 = (^(v3))
    if true {
        for fi27 := 0; fi27 < 17; fi27++ {
            fmt.Println(fi27)
            var v28 uint32 = v14
            _ = fi27
            _ = v28
        }
        v1 += (v1 << ((0) & 63))
        fmt.Println(v10)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}

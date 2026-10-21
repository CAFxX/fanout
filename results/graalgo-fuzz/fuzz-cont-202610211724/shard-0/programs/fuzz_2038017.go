package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
func genFn0(p0 int64, p1 uint64, p2 uint) uint32 {
    p1 += (p1 >> (((^(p2))) & 63))
    var v6 int16 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 260
}
func genFn1(p0 bool) uint32 {
    var v7 int16 = -32768
    v8 := GenS1{F0: (0 <= 1)}
    _ = p0
    _ = v7
    _ = v8
    return 65535
}

func main() {
    v1 := []string{"hello", "foo bar", "a", "hello"}
    v2 := []string{}
    var v3 uint32 = 4294967295
    v4 := make(map[int]int)
    v5 := []uint32{(v3 - v3), v3, v3}
    if true {
        if (false || false) {
            fmt.Println(v3)
        }
    }
    if true {
        var v9 float32 = 1.5
        if false {
            if false {
                for fi10 := 0; fi10 < 15; fi10++ {
                    v11 := GenS1{F0: (4294967295 == 136)}
                    if len(v1) > 0 {
                        v1[0] = (("a" + ("\n\t" + "a")) + (((("\n\t" + "") + ("foo bar" + ("\n\t" + ""))) + (("" + ("a" + "\n\t")) + "hello")) + ("" + "a")))
                    }
                    _ = fi10
                    _ = v11
                }
                v12 := GenS1{F0: true}
                _ = v12
            } else {
                v9 = v9
            }
            v9 -= v9
        } else {
            v3 = v3
            if false {
                for fi13 := 0; fi13 < 9; fi13++ {
                    v9 -= float32((0.0 - -2.25))
                    if (true && false) {
                        fmt.Println(fi13, v9, v9)
                        v14 := v4[fi13]
                        _ = v14
                    }
                    _ = fi13
                }
                fmt.Println(v3)
            }
        }
        v15 := GenS1{F0: (1 >= 1000)}
        _ = v15
        _ = v9
    } else {
        if false {
            var v16 uint32 = v3
            _ = v16
        } else {
            var v17 int8 = 42
            v17 += (v17 % ((100) | 1))
            _ = v17
        }
    }
    var v18 int32 = 2147483647
    v18 = 320
    v18 += v18
    fmt.Println(v18, v3)
    switch "foo bar" {
    case "foo bar":
        v3 = 270
    case "hello":
        fmt.Println(v18, v18)
    }
    switch ("" + "\n\t") {
    case "\n\t":
        if true {
            for fi19 := 0; fi19 < 10; fi19++ {
                v3 = (v3 & (v3 * v3))
                switch (fi19 - 32768) {
                case 32767:
                    v20 := GenS1{F0: true}
                    _ = v20
                default:
                    v18 += v18
                }
                var v21 uint8 = 127
                _ = fi19
                _ = v21
            }
            for fi22 := 0; fi22 < 3; fi22++ {
                v4[(fi22 + fi22)] = fi22
                for fi23 := 0; fi23 < 4; fi23++ {
                    v24 := v4[-331]
                    fmt.Println(fi23, fi22, fi23)
                    _ = fi23
                    _ = v24
                }
                _ = fi22
            }
            v3 += 1
        } else {
            v25 := GenS2{F0: -1}
            _ = v25
        }
    case "":
        if len(v5) > 0 {
            v5[0] = v3
        }
    case "a":
        fmt.Println(v3, v18)
    }
    v18 = -1
    _ = len(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

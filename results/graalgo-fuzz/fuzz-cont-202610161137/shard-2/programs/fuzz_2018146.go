package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
func genFn0(p0 string, p1 int, p2 uint16) uint64 {
    var v4 uint8 = 0
    if false {
        v5 := []int{0, (p1 / (((p1 % ((p1) | 1))) | 1)), (p1 % ((p1) | 1))}
        _ = v5
    }
    fmt.Println(v4, p0)
    for fi6 := 0; fi6 < 20; fi6++ {
        fmt.Println(fi6, p2, p0)
        switch p0 {
        case "hello":
            var v7 string = p0
            _ = v7
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return uint64(p2)
}
func genFn1(p0 int) uint {
    v8 := GenS1{F0: (-68 <= 7), F1: 0}
    v9 := []int{}
    _ = p0
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    v1 := []string{"hello", "foo bar"}
    v2 := []uint32{65535, 65535, 4294967295, 380}
    var v3 uint = 0
    v3 += v3
    switch (("a" + ("foo bar" + "\n\t")) + ("hello" + "hello")) {
    case "":
        v3 += (v3 / (((v3 << (((v3 + 0)) & 63))) | 1))
    case "foo bar":
        fmt.Println(v3)
    }
    var v10 int = 2147483647
    var v11 int16 = -1
    var v12 int = v10
    v2 = append(v2, 2147483647)
    for fi13 := 0; fi13 < 19; fi13++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if true {
            v2 = append(v2, 65535)
        } else {
            if len(v2) > 0 {
                v2[0] = 0
            }
            if false {
                var v14 float32 = (float32(0.0) * 1e10)
                _ = v14
            } else {
                switch v10 {
                case -1:
                    fmt.Println(fi13, v10, v10)
                default:
                    var v15 uint16 = 2
                    _ = v15
                }
                for fi16 := 0; fi16 < 18; fi16++ {
                    var v17 int32 = -2147483648
                    _ = fi16
                    _ = v17
                }
            }
            var v18 float64 = (1e10 - -2.25)
            _ = v18
        }
        if (-882 >= 0) {
            if ((v3 << (((v3 >> ((v3) & 63))) & 63)) > 1) {
                if true {
                    for fi19 := 0; fi19 < 9; fi19++ {
                        v12 += v10
                        var v20 int8 = -62
                        _ = fi19
                        _ = v20
                    }
                    if len(v1) > 0 {
                        v1[0] = ("a" + "foo bar")
                    }
                    var v21 uint16 = 65535
                    _ = v21
                }
            }
        } else {
            var v22 string = ("" + "")
            var v23 string = "a"
            _ = v22
            _ = v23
        }
        _ = fi13
    }
    var v24 float32 = 0.0
    v11 = (v11 % ((v11) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v3
}

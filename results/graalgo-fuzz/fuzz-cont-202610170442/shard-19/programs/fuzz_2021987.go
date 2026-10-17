package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
func genFn0(p0 int8, p1 uint16) int64 {
    p0 -= p0
    var v6 int = 127
    p1 = p1
    _ = p0
    _ = p1
    _ = v6
    return -1
}
func genFn1(p0 uint32, p1 string, p2 uint8) bool {
    fmt.Println(p1, p0)
    fmt.Println(p2)
    switch p1 {
    case "foo bar":
        var v7 int16 = 0
        _ = v7
    default:
        var v8 uint8 = (p2 * (^(p2)))
        _ = v8
    }
    var v9 int32 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (!(p1 <= "hello"))
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[int]string)
    var v3 int32 = -2147483648
    var v4 uint = 0
    v5 := make(map[string]uint32)
    if (false && true) {
        v10 := GenS1{F0: 32767, F1: ((("a" + "\n\t") + "") != ("a" + "foo bar"))}
        _ = v10
    } else {
        v11 := GenS1{F0: -9223372036854775808, F1: true}
        _ = v11
    }
    var v12 bool = true
    for fi13 := 0; fi13 < 17; fi13++ {
        v14 := []int{fi13, fi13}
        _ = fi13
        _ = v14
    }
    if (42 != 595) {
        v12 = v12
        for fi15 := 0; fi15 < 2; fi15++ {
            v16 := GenS1{F0: (^(fi15)), F1: ((0 == 1) && v12)}
            var v17 int32 = (-(v3))
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    v18 := GenS1{F0: 128, F1: (float32(1e10) <= (1.5 / -2.25))}
    for fi19 := 0; fi19 < 4; fi19++ {
        switch "hello" {
        case "hello":
            if true {
                switch (fi19 | (fi19 * (fi19 * (fi19 * (fi19 / ((-2147483648) | 1)))))) {
                case 42:
                    var v20 int32 = v3
                    _ = v20
                case 128:
                    var v21 uint8 = 58
                    _ = v21
                case 4294967295:
                    var v22 bool = v12
                    _ = v22
                }
                var v23 int = (fi19 >> ((v4) & 63))
                delete(v5, "")
                _ = v23
            }
        case "foo bar":
            v12 = (!((true || false) && false))
        case "a":
            if v12 {
                var v24 int16 = 1
                if (3.14159 < 0.1) {
                    var v25 float64 = (float64((3.14159 + float32(-2.25))) * (-2.25 * 3.14159))
                    if v12 {
                        switch fi19 {
                        case 783:
                            v25 = 1.5
                        default:
                            v24 -= (v24 | (v24 & v24))
                        }
                        var v26 float32 = 3.14159
                        _ = v26
                    } else {
                        v27 := []int{fi19, int(v24)}
                        _ = v27
                    }
                    _ = v25
                }
                _ = v24
            } else {
                fmt.Println(v12, v4)
            }
        }
        _ = fi19
    }
    var v28 uint32 = 65535
    var v29 bool = true
    var v30 uint8 = 127
    var v31 int16 = 32767
    for fi32 := 0; fi32 < 6; fi32++ {
        var v33 int16 = (v31 % ((v31) | 1))
        _ = fi32
        _ = v33
    }
    v34 := []int{2147483647, int(v1), 4294967295}
    v29 = v12
    var v35 float32 = 1e10
    _ = len(v34)
    switch "a" {
    case "foo bar":
        var v36 int16 = v31
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}

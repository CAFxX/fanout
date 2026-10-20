package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint16; F2 int64 }
type GenS3 struct { F0 int8; F1 uint; F2 uint8 }
func genFn0(p0 uint, p1 uint8, p2 bool) uint {
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := []int{fi6, 128, (-(fi6)), (^(fi6))}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        fmt.Println(p0, p1, p2)
        _ = fi8
    }
    var v9 uint8 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 915
}
func genFn1() int8 {
    var v10 float64 = (0.0 * -2.25)
    var v11 int8 = 100
    _ = v10
    _ = v11
    return v11
}

func main() {
    var v1 int16 = 667
    var v2 int64 = 623
    v3 := make(map[string]string)
    v4 := []uint32{537, 4294967295, 276}
    var v5 uint32 = 2
    if true {
        var v12 uint8 = 255
        fmt.Println(v2, v1, v1)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v12
    } else {
        for fi13 := 0; fi13 < 4; fi13++ {
            v2 += (v2 >> ((2) & 63))
            switch "" {
            case "foo bar":
                if true {
                    fmt.Println(v5, v1)
                }
            case "hello":
                v14 := GenS1{F0: 0, F1: 17}
                _ = v14
            }
            v3["\n\t"] = ("\n\t" + "")
            _ = fi13
        }
        v3["a"] = "foo bar"
        if true {
            v15 := GenS1{F0: 2147483647, F1: 127}
            for fi16 := 0; fi16 < 1; fi16++ {
                fmt.Println(v5, v1)
                _ = fi16
            }
            var v17 int16 = (v1 & (v1 - 32767))
            _ = v15
            _ = v17
        } else {
            v1 += 0
            switch -390 {
            case 1:
                var v18 int64 = (v2 % ((v2) | 1))
                _ = v18
            }
            v5 = 4294967295
        }
    }
    v5 = v5
    v3[("" + "foo bar")] = ""
    fmt.Println(v5, v2)
    for fi19 := 0; fi19 < 13; fi19++ {
        switch "hello" {
        case "":
            v5 = 65535
        }
        _ = fi19
    }
    var v20 string = "hello"
    v1 = v1
    var v21 uint8 = 2
    for fi22 := 0; fi22 < 16; fi22++ {
        if true {
            v4 = append(v4, (v5 | 65535))
            _ = len(v4)
            if false {
                v23 := GenS2{F0: 127, F1: 2, F2: v2}
                _ = v23
            }
        }
        v24 := GenS1{F0: -1, F1: (^(v21))}
        _ = len(v4)
        _ = fi22
        _ = v24
    }
    if false {
        var v25 int64 = (v2 - v2)
        v26 := v3[v20]
        _ = v25
        _ = v26
    }
    var v27 int32 = 1
    v28 := GenS3{F0: 1, F1: genFn0(42, (v21 * v21), true), F2: v21}
    fmt.Println(v1, v20)
    v29 := v3[("hello" + "foo bar")]
    var v30 uint64 = 4294967295
    v29 = "\n\t"
    var v31 int8 = genFn1()
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}

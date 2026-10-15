package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 bool, p1 uint16) int {
    var v6 uint64 = 2
    v7 := []int{}
    v8 := GenS1{F0: p1}
    if ((^(v6)) > v6) {
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 string = "foo bar"
            var v11 uint = 1
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 int64 = 0
        var v13 uint64 = 0
        _ = v12
        _ = v13
    } else {
        var v14 int32 = -778
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return -853
}

func main() {
    var v1 string = "a"
    var v2 float32 = -2.25
    var v3 int8 = 1
    v4 := []uint32{2}
    var v5 string = "\n\t"
    if false {
        if len(v4) > 0 {
            v4[0] = 0
        }
        var v15 uint8 = 0
        _ = len(v4)
        _ = v15
    } else {
        for fi16 := 0; fi16 < 6; fi16++ {
            switch "foo bar" {
            case "foo bar":
                fmt.Println(fi16, v1)
            case "hello":
                if ((3.14159 - (-2.25 - 1.5)) < 3.14159) {
                    fmt.Println(v1)
                    var v17 uint = 94
                    var v18 uint64 = 411
                    _ = v17
                    _ = v18
                }
            case "a":
                fmt.Println(v2)
            default:
                var v19 uint = 255
                _ = v19
            }
            v5 = (v5 + v5)
            _ = fi16
        }
        var v20 uint = 4
        var v21 uint8 = 2
        _ = v20
        _ = v21
    }
    v2 = v2
    if (v2 == ((v2 + (v2 * v2)) / v2)) {
        for fi22 := 0; fi22 < 2; fi22++ {
            for fi23 := 0; fi23 < 5; fi23++ {
                fmt.Println(v1)
                _ = fi23
            }
            v2 -= float32(1.5)
            v1 = (((v5 + (v1 + v5)) + v1) + (v1 + v5))
            _ = fi22
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v24 := GenS1{F0: 1}
    fmt.Println(v2, v3, v5)
    var v25 int16 = 1
    var v26 uint = 42
    fmt.Println(v5, v2)
    v3 -= (-(v3))
    var v27 string = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}

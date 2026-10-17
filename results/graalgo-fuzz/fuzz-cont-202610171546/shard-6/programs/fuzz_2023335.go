package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 int8; F2 int64 }
func genFn0(p0 string, p1 int8) int16 {
    var v7 int8 = 100
    fmt.Println(p1, p1)
    if false {
        if false {
            var v8 uint16 = 377
            _ = v8
        }
        v9 := GenS1{F0: -2147483648}
        _ = v9
    }
    if false {
        v10 := GenS1{F0: -1}
        p0 = (p0 + (p0 + ("\n\t" + ((p0 + "a") + ("a" + "foo bar")))))
        v11 := GenS1{F0: 2147483647}
        _ = v10
        _ = v11
    } else {
        var v12 int16 = 42
        for fi13 := 0; fi13 < 10; fi13++ {
            fmt.Println(p1)
            _ = fi13
        }
        switch 127 {
        case 617:
            if (4294967295 <= 2) {
                var v14 float64 = -2.25
                var v15 int64 = 42
                _ = v14
                _ = v15
            }
        case 42:
            v16 := []int{128, 127}
            _ = v16
        case 4294967295:
            v17 := GenS2{F0: 32767}
            _ = v17
        }
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v7
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = (false || (42 < 492))
    v3 := make(map[int]string)
    var v4 uint32 = 769
    var v5 uint = 0
    v6 := make(map[uint32]bool)
    if true {
        if v2 {
            for fi18 := 0; fi18 < 10; fi18++ {
                var v19 uint8 = 1
                v5 -= v5
                _ = fi18
                _ = v19
            }
            if v2 {
                if ((!v2) || v2) {
                    fmt.Println(v4, v5, v2)
                } else {
                    v20 := []int{}
                    v4 = 65535
                    v4 -= (v4 | 42)
                    _ = v20
                }
            } else {
                var v21 int8 = -1
                fmt.Println(v4)
                var v22 uint16 = 42
                _ = v21
                _ = v22
            }
            for fi23 := 0; fi23 < 19; fi23++ {
                v2 = v2
                _ = fi23
            }
        }
    }
    v24 := GenS1{F0: 882}
    v2 = (128 > 42)
    switch 2147483647 {
    case 0:
        v2 = false
    case 271:
        for fi25 := 0; fi25 < 14; fi25++ {
            v26 := GenS1{F0: 1}
            var v27 string = "a"
            _ = fi25
            _ = v26
            _ = v27
        }
    case 4294967295:
        var v28 int = 577
        _ = v28
    }
    v29 := []int{706, 0, 2147483647}
    if v2 {
        if (uint16(v5) < 910) {
            v4 += 0
            var v30 int32 = 1
            _ = v30
        } else {
            for fi31 := 0; fi31 < 1; fi31++ {
                _ = len(v29)
                v4 += 1
                _ = fi31
            }
            v32 := v6[v4]
            _ = v32
        }
        v33 := GenS1{F0: 42}
        _ = v33
    } else {
        v2 = (!false)
        var v34 bool = v2
        v4 = v4
        _ = v34
    }
    var v35 int32 = -2147483648
    v36 := GenS1{F0: (v35 | (v35 - 0))}
    v37 := GenS3{F0: 0, F1: 46, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}

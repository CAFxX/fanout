package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint8, p1 bool, p2 int8) int16 {
    p2 -= 42
    fmt.Println(p2)
    var v7 float64 = float64(0.0)
    switch "" {
    case "a":
        if p1 {
            var v8 int = 128
            _ = v8
        }
    case "foo bar":
        var v9 int64 = -907
        _ = v9
    default:
        var v10 uint32 = 797
        _ = v10
    }
    p1 = true
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}
func genFn1(p0 int64) string {
    p0 = (p0 - 815)
    fmt.Println(p0)
    switch 0 {
    case -521:
        v11 := GenS1{F0: 11}
        _ = v11
    default:
        fmt.Println(p0)
    }
    _ = p0
    return "a"
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]int)
    var v3 uint64 = 378
    var v4 uint32 = 0
    v5 := []string{("hello" + "\n\t"), "\n\t", "foo bar", "foo bar"}
    var v6 int16 = -821
    if false {
        var v12 int8 = -56
        fmt.Println(v4, v4)
        _ = v12
    }
    if v1 {
        for fi13 := 0; fi13 < 9; fi13++ {
            fmt.Println(v3, fi13, v3)
            _ = len(v2)
            _ = fi13
        }
        var v14 uint = 847
        v3 -= v3
        _ = v14
    } else {
        if v1 {
            if len(v5) > 0 {
                v5[0] = "\n\t"
            }
        }
    }
    var v15 float32 = 0.1
    v16 := GenS2{F0: 100}
    switch int(v3) {
    case -2147483648:
        switch -338 {
        case 4294967295:
            v17 := GenS1{F0: 2}
            _ = v17
        default:
            for fi18 := 0; fi18 < 17; fi18++ {
                switch "a" {
                case "a":
                    v19 := GenS1{F0: 235}
                    _ = v19
                case "foo bar":
                    v1 = v1
                }
                _ = fi18
            }
        }
    }
    v4 += (^(v4))
    var v20 float32 = (v15 / v15)
    fmt.Println(v15, v20)
    switch "" {
    case "\n\t":
        switch 127 {
        case 42:
            var v21 uint64 = v3
            _ = v21
        case 128:
            fmt.Println(v4, v3)
        default:
            fmt.Println(v6)
        }
    }
    var v22 int8 = 63
    if ((-1 == -868) || ((3.14159 + 3.14159) > 0.1)) {
        var v23 float32 = v15
        for fi24 := 0; fi24 < 12; fi24++ {
            var v25 int8 = v22
            _ = fi24
            _ = v25
        }
        v20 = v15
        _ = v23
    } else {
        v26 := GenS1{F0: 127}
        for fi27 := 0; fi27 < 9; fi27++ {
            v28 := v2[(v4 + (v4 ^ v4))]
            _ = fi27
            _ = v28
        }
        switch -806 {
        case 0:
            v20 = v15
        }
        _ = v26
    }
    for fi29 := 0; fi29 < 9; fi29++ {
        if (508 >= 42) {
            for fi30 := 0; fi30 < 13; fi30++ {
                var v31 int8 = (^(v22))
                var v32 int64 = 1
                _ = fi30
                _ = v31
                _ = v32
            }
            var v33 uint32 = 2147483648
            var v34 float64 = 1.5
            _ = v33
            _ = v34
        }
        var v35 uint16 = 0
        var v36 uint = 0
        _ = fi29
        _ = v35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

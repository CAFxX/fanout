package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint; F1 string }
func genFn0() uint32 {
    v4 := []int{395, 4294967295, 2147483647, -841}
    var v5 string = ""
    v5 = v5
    _ = v4
    _ = v5
    return 2147483648
}

func main() {
    v1 := []uint32{4294967295, 42, 4294967295, 0}
    var v2 uint16 = 0
    var v3 int8 = 1
    var v6 int = 4294967295
    if len(v1) > 0 {
        v1[0] = 2
    }
    var v7 int8 = (^(v3))
    v8 := GenS1{F0: 4294967295, F1: v6}
    var v9 string = ("\n\t" + "")
    v1 = append(v1, 1)
    var v10 uint16 = v2
    v11 := GenS2{F0: 4294967295}
    var v12 uint = 0
    fmt.Println(v6, v6, v12)
    v12 = (^(v12))
    if false {
        for fi13 := 0; fi13 < 19; fi13++ {
            v14 := GenS1{F0: 42, F1: (fi13 >> (((^(v12))) & 63))}
            v15 := GenS3{F0: 2147483647}
            _ = fi13
            _ = v14
            _ = v15
        }
    } else {
        var v16 uint8 = 255
        _ = v16
    }
    var v17 uint8 = 1
    if ((v12 + 676) == (^(v12))) {
        for fi18 := 0; fi18 < 16; fi18++ {
            v6 += (fi18 - v6)
            _ = fi18
        }
        var v19 float64 = 0.1
        switch (v6 & (^(v6))) {
        case 2147483647:
            v20 := GenS2{F0: (v12 % ((v12) | 1))}
            _ = v20
        default:
            v10 += (v2 & v2)
        }
        _ = v19
    } else {
        for fi21 := 0; fi21 < 12; fi21++ {
            if (false || (0 == -917)) {
                var v22 float32 = 0.1
                _ = v22
            }
            _ = fi21
        }
        fmt.Println(v9, v3)
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        for fi24 := 0; fi24 < 12; fi24++ {
            fmt.Println(v17, fi23, v6)
            var v25 int32 = 42
            _ = fi24
            _ = v25
        }
        var v26 int8 = (v3 + 100)
        for fi27 := 0; fi27 < 18; fi27++ {
            switch (v6 / ((-721) | 1)) {
            case 4294967295:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case -1:
                v28 := GenS4{F0: v9}
                _ = v28
            case -708:
                v29 := GenS1{F0: 255, F1: fi27}
                _ = v29
            }
            switch ((((v9 + (("" + "a") + "\n\t")) + (v9 + v9)) + "") + v9) {
            case "\n\t":
                v30 := GenS5{F0: v12, F1: v9}
                _ = v30
            case "foo bar":
                var v31 string = ""
                _ = v31
            default:
                v10 = (v2 * v2)
            }
            _ = fi27
        }
        _ = fi23
        _ = v26
    }
    var v32 float32 = 1e10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

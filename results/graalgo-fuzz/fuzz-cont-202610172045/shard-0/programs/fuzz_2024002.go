package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint8) bool {
    for fi4 := 0; fi4 < 5; fi4++ {
        v5 := []int{fi4, (fi4 - fi4), fi4}
        p0 -= (p0 / ((255) | 1))
        var v6 uint16 = 2
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 bool = false
        _ = fi7
        _ = v8
    }
    v9 := []int{-879, 4294967295}
    _ = p0
    _ = v9
    return (!false)
}
func genFn1(p0 int32) string {
    fmt.Println(p0)
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := []int{42, (fi10 / ((127) | 1)), (fi10 % (((fi10 * fi10)) | 1)), -769}
        var v12 float64 = 1.5
        for fi13 := 0; fi13 < 9; fi13++ {
            _ = len(v11)
            v14 := GenS1{F0: 4294967295}
            _ = fi13
            _ = v14
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    v15 := GenS1{F0: 0}
    var v16 uint8 = 2
    var v17 int16 = 1
    _ = p0
    _ = v15
    _ = v16
    _ = v17
    return "a"
}

func main() {
    var v1 int8 = 42
    var v2 uint16 = 1
    v3 := []int{127}
    if false {
        v1 = 1
        fmt.Println(v1, v1)
    } else {
        v1 -= (v1 - v1)
        var v18 uint16 = (v2 / (((v2 >> ((698) & 63))) | 1))
        _ = v18
    }
    var v19 int8 = (v1 / ((48) | 1))
    v20 := GenS1{F0: 611}
    var v21 float32 = ((3.14159 - 1e10) / 1e10)
    var v22 int8 = (^(v19))
    var v23 float64 = 1e10
    var v24 int8 = 104
    for fi25 := 0; fi25 < 8; fi25++ {
        switch "hello" {
        case "a":
            var v26 int16 = 42
            _ = v26
        case "\n\t":
            v27 := GenS1{F0: 255}
            _ = v27
        }
        _ = fi25
    }
    v28 := GenS1{F0: 1}
    var v29 uint16 = (v2 | v2)
    v29 -= (^(v29))
    v19 = v19
    if (v21 < 0.1) {
        v22 += -128
    } else {
        var v30 string = "a"
        if (2147483647 >= 42) {
            if true {
                if (1000 <= -958) {
                    var v31 uint64 = 830
                    var v32 int16 = -48
                    var v33 uint16 = (v29 + (v2 << ((65535) & 63)))
                    _ = v31
                    _ = v32
                    _ = v33
                }
            }
        }
        _ = v30
    }
    for fi34 := 0; fi34 < 5; fi34++ {
        v23 -= -2.25
        _ = fi34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
}

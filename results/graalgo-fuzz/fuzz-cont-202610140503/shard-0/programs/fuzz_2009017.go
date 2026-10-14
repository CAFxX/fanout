package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 int8 }

func main() {
    var v1 int = 4294967295
    var v2 int = v1
    var v3 float64 = 1.5
    var v4 uint = 255
    var v5 uint = (v4 * (v4 << (((v4 + (^(v4)))) & 63)))
    var v6 uint64 = 4294967295
    if (2 >= 127) {
        if (788 != 239) {
            var v7 bool = true
            v2 = (^(v2))
            for fi8 := 0; fi8 < 10; fi8++ {
                for fi9 := 0; fi9 < 4; fi9++ {
                    var v10 int32 = 1
                    var v11 int = fi8
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                _ = fi8
            }
            _ = v7
        } else {
            if (!false) {
                v12 := []int{v2, (v2 % ((v1) | 1)), (^(v2))}
                var v13 uint64 = v6
                var v14 uint8 = 15
                _ = v12
                _ = v13
                _ = v14
            }
        }
        var v15 float32 = 3.14159
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 float64 = (3.14159 + 3.14159)
            _ = fi16
            _ = v17
        }
        _ = v15
    } else {
        fmt.Println(v5)
        v18 := GenS1{F0: 1}
        fmt.Println(v4)
        _ = v18
    }
    var v19 uint32 = 4294967295
    for fi20 := 0; fi20 < 5; fi20++ {
        if true {
            var v21 uint32 = (v19 * v19)
            v22 := GenS2{F0: 2147483647, F1: -68}
            var v23 string = ""
            _ = v21
            _ = v22
            _ = v23
        }
        fmt.Println(v19)
        _ = fi20
    }
    v24 := GenS1{F0: 183}
    v1 += 685
    var v25 float32 = float32(v3)
    v26 := []int{(v2 - v1)}
    v5 = v4
    if (false || false) {
        v26 = append(v26, v2)
        v6 = (v6 * (v6 - (v6 >> ((v4) & 63))))
    }
    if false {
        v1 += (v2 % ((-9223372036854775808) | 1))
        if len(v26) >= 2 {
            v26 = v26[1:len(v26)]
        }
        v19 = v19
    } else {
        v4 = (v4 + 42)
        _ = len(v26)
        v5 = v4
    }
    for fi27 := 0; fi27 < 11; fi27++ {
        if (369 <= 42) {
            var v28 uint8 = 128
            if false {
                v3 = 3.14159
            }
            _ = v28
        } else {
            fmt.Println(v4, v1)
        }
        var v29 string = "foo bar"
        switch (v29 + "foo bar") {
        case "hello":
            var v30 string = v29
            _ = v30
        }
        _ = fi27
        _ = v29
    }
    var v31 uint = 0
    v32 := GenS1{F0: 127}
    var v33 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}

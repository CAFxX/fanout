package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 uint; F1 int16; F2 int }
func genFn0() uint16 {
    v7 := GenS1{F0: 173, F1: 926, F2: 18446744073709551615}
    v8 := GenS1{F0: 1, F1: 1, F2: 18446744073709551615}
    var v9 uint = 18446744073709551615
    fmt.Println(v9)
    _ = v7
    _ = v8
    _ = v9
    return 831
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{-22, 127}
    var v3 int64 = 1
    var v4 int = 4294967295
    var v5 bool = (27 <= 145)
    var v6 int8 = 5
    switch 911 {
    case 647:
        for fi10 := 0; fi10 < 2; fi10++ {
            v6 = v6
            _ = fi10
        }
    case 0:
        var v11 uint8 = 42
        _ = v11
    }
    v12 := GenS1{F0: 7, F1: 931, F2: 585}
    switch -694 {
    case 4294967295:
        for fi13 := 0; fi13 < 5; fi13++ {
            if (false && v5) {
                var v14 int = 127
                v15 := GenS1{F0: 128, F1: 2, F2: 2}
                for fi16 := 0; fi16 < 3; fi16++ {
                    var v17 string = ("\n\t" + ((("\n\t" + "hello") + "") + "a"))
                    _ = fi16
                    _ = v17
                }
                _ = v14
                _ = v15
            }
            var v18 int16 = 1
            _ = fi13
            _ = v18
        }
    }
    var v19 float64 = 0.1
    if (false && false) {
        fmt.Println(v3, v6)
        v20 := GenS1{F0: 128, F1: 89, F2: 42}
        _ = len(v2)
        _ = v20
    } else {
        var v21 float32 = 3.14159
        v22 := GenS1{F0: 1, F1: 0, F2: 42}
        if false {
            for fi23 := 0; fi23 < 11; fi23++ {
                if v5 {
                    var v24 uint = 18446744073709551615
                    _ = v24
                } else {
                    v2 = append(v2, fi23)
                    v19 = 3.14159
                    var v25 float32 = (v21 * (v21 - 0.0))
                    _ = v25
                }
                _ = fi23
            }
            var v26 int16 = -490
            v27 := GenS2{F0: 42, F1: (v6 >= 100)}
            _ = v26
            _ = v27
        }
        _ = v21
        _ = v22
    }
    v28 := GenS3{F0: 2, F1: -32768, F2: 4294967295}
    var v29 string = "\n\t"
    var v30 int64 = 42
    switch (v4 / ((v4) | 1)) {
    case 32767:
        v31 := v1[0]
        _ = v31
    case 402:
        for fi32 := 0; fi32 < 3; fi32++ {
            var v33 float64 = 0.1
            v34 := GenS1{F0: 127, F1: 255, F2: 2}
            _ = fi32
            _ = v33
            _ = v34
        }
    case 128:
        v6 = -1
    }
    v35 := GenS1{F0: 1, F1: 158, F2: 42}
    delete(v1, 2147483648)
    v36 := v1[42]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}

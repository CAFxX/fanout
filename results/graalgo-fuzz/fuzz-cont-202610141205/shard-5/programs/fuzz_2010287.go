package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 float64 = (3.14159 - -2.25)
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    var v4 uint16 = 65535
    var v5 uint16 = v4
    var v6 int8 = -1
    v1 = ((v1 + (v1 * 1.5)) * v1)
    if true {
        v7 := v3[2147483647]
        _ = v7
    } else {
        v8 := GenS1{F0: 834}
        _ = v8
    }
    v9 := GenS1{F0: 1000}
    fmt.Println(v4, v5, v5)
    var v10 int8 = (v6 & v6)
    v5 += (v5 - v5)
    fmt.Println(v5)
    v11 := GenS1{F0: 1000}
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 float64 = v1
        v5 = v4
        if false {
            v4 -= v4
            for fi14 := 0; fi14 < 20; fi14++ {
                v13 -= v13
                switch (fi14 % (((fi14 ^ fi14)) | 1)) {
                case 32767:
                    v15 := GenS1{F0: 705}
                    _ = v15
                case 1:
                    v16 := v2["\n\t"]
                    _ = v16
                case 128:
                    var v17 uint8 = 93
                    _ = v17
                }
                var v18 int16 = int16(fi12)
                _ = fi14
                _ = v18
            }
        }
        _ = fi12
        _ = v13
    }
    var v19 float64 = 3.14159
    for fi20 := 0; fi20 < 17; fi20++ {
        var v21 uint32 = 192
        _ = fi20
        _ = v21
    }
    v22 := []int{1}
    _ = len(v2)
    for fi23 := 0; fi23 < 11; fi23++ {
        var v24 bool = (!(false && false))
        v22 = append(v22, fi23)
        if (("hello" + "") < ((("hello" + "hello") + ("" + "\n\t")) + ("a" + "a"))) {
            var v25 int = -458
            _ = v25
        }
        _ = fi23
        _ = v24
    }
    var v26 uint8 = 127
    v1 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}

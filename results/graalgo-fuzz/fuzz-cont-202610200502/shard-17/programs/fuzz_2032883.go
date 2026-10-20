package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 int64 = -112
    var v2 int16 = 407
    v3 := make(map[int]uint32)
    var v4 uint16 = 32767
    delete(v3, -9223372036854775808)
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := []int{32767, 1, 384}
        v7 := GenS1{F0: 115}
        fmt.Println(v1, v2, v2)
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 string = ("\n\t" + "\n\t")
    if true {
        for fi9 := 0; fi9 < 20; fi9++ {
            if true {
                v10 := GenS2{F0: v1}
                _ = v10
            }
            fmt.Println(v1)
            _ = fi9
        }
    }
    if (0 < 670) {
        _ = len(v3)
    } else {
        v11 := []int{1}
        var v12 bool = ((984 == -1) && (!true))
        if (!v12) {
            if len(v11) > 0 {
                v11[0] = -474
            }
            switch 0 {
            case 32767:
                switch -2147483648 {
                case 42:
                    var v13 int8 = 100
                    _ = v13
                case -106:
                    v12 = (42 <= 887)
                }
            case 127:
                fmt.Println(v4, v12, v1)
            }
            var v14 float64 = 1e10
            _ = v14
        }
        _ = v11
        _ = v12
    }
    v8 = v8
    v15 := GenS1{F0: 255}
    v2 = v2
    for fi16 := 0; fi16 < 16; fi16++ {
        switch (fi16 + fi16) {
        case 128:
            var v17 uint = 303
            _ = v17
        case -2147483648:
            if (1.5 != ((((1.5 * (1.5 / (-2.25 / 1.0))) - 1.5) - 3.14159) * 3.14159)) {
                v1 -= -847
                if false {
                    var v18 string = v8
                    _ = v18
                } else {
                    if true {
                        _ = len(v3)
                        v8 = ("hello" + v8)
                    }
                }
                v4 += v4
            }
        case 32767:
            v1 = 0
        default:
            v19 := []int{41}
            _ = v19
        }
        v2 = (v2 + v2)
        var v20 uint = 255
        _ = fi16
        _ = v20
    }
    v2 = v2
    fmt.Println(v8, v8, v4)
    v21 := []int{127}
    v4 += 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v8
}

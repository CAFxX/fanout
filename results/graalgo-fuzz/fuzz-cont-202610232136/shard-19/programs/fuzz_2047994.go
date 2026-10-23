package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 255
    var v3 int32 = 2147483647
    v4 := GenS1{F0: 343, F1: 2147483647, F2: 14}
    v2 += 722
    v3 -= 2147483647
    if true {
        v5 := v1[-494]
        var v6 uint32 = v5
        delete(v1, 1)
        _ = v5
        _ = v6
    }
    var v7 int8 = 0
    var v8 uint32 = 4294967295
    v9 := []int{-581, 42}
    switch 478 {
    case 42:
        switch 1 {
        case 1:
            v7 = 100
        }
    }
    var v10 int = 32768
    switch (v10 & v10) {
    case 128:
        for fi11 := 0; fi11 < 17; fi11++ {
            if false {
                v8 = (v8 << ((65535) & 63))
            } else {
                if false {
                    var v12 float32 = ((0.1 - -2.25) + (0.1 - (3.14159 - -2.25)))
                    switch "foo bar" {
                    case "":
                        for fi13 := 0; fi13 < 10; fi13++ {
                            v14 := GenS1{F0: fi11, F1: (v3 / (((v3 / ((v3) | 1))) | 1)), F2: (fi11 - 937)}
                            var v15 string = "a"
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                    case "hello":
                        _ = len(v1)
                    case "a":
                        switch fi11 {
                        case 128:
                            switch v10 {
                            case 889:
                                var v16 float64 = (1.5 / 0.1)
                                _ = v16
                            }
                        default:
                            var v17 bool = true
                            _ = v17
                        }
                    }
                    var v18 bool = (158 >= 42)
                    _ = v12
                    _ = v18
                }
                for fi19 := 0; fi19 < 4; fi19++ {
                    var v20 int16 = 558
                    _ = fi19
                    _ = v20
                }
            }
            _ = fi11
        }
    case 4294967295:
        _ = len(v1)
    }
    if ("foo bar" >= ("\n\t" + ("foo bar" + "foo bar"))) {
        switch "a" {
        case "a":
            var v21 uint32 = (v8 | 0)
            _ = v21
        case "hello":
            _ = len(v1)
        }
        var v22 uint64 = 0
        switch (v10 * v10) {
        case 273:
            if (!false) {
                var v23 bool = false
                _ = v23
            }
        }
        _ = v22
    } else {
        for fi24 := 0; fi24 < 2; fi24++ {
            if (322 <= 2) {
                var v25 bool = false
                v26 := GenS1{F0: fi24, F1: v3, F2: 42}
                _ = v25
                _ = v26
            } else {
                v2 = (v2 << ((206) & 63))
            }
            _ = fi24
        }
    }
    _ = len(v1)
    v9 = append(v9, v10)
    var v27 float32 = (((1e10 * 3.14159) - ((1e10 - 3.14159) - 1e10)) * 3.14159)
    v9 = append(v9, (v10 - (v10 + 951)))
    v10 += (v10 / (((v10 - v10)) | 1))
    var v28 int = (v10 & (-(v10)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int; F1 int; F2 int8 }

func main() {
    var v1 uint32 = 4294967295
    var v2 uint16 = 32767
    v3 := make(map[string]string)
    var v4 int16 = -32768
    var v5 int64 = 0
    var v6 int64 = (v5 | v5)
    if (!true) {
        v7 := []int{696, -363, 714}
        var v8 uint8 = 255
        for fi9 := 0; fi9 < 9; fi9++ {
            v10 := GenS1{F0: (v4 % (((v4 / ((763) | 1))) | 1)), F1: (fi9 - fi9)}
            _ = fi9
            _ = v10
        }
        _ = v7
        _ = v8
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 uint64 = 4294967295
        var v13 int = 601
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 uint32 = 0
    var v15 string = "\n\t"
    var v16 float64 = 0.0
    var v17 uint32 = (v14 | (v1 ^ 2147483647))
    v18 := GenS2{F0: 2, F1: 0}
    v19 := GenS2{F0: 4294967295, F1: 0}
    for fi20 := 0; fi20 < 5; fi20++ {
        v4 -= v4
        _ = fi20
    }
    v21 := GenS2{F0: 1, F1: 0}
    var v22 float32 = 1e10
    var v23 bool = (((v1 >> ((44) & 63)) == v17) && false)
    v23 = false
    for fi24 := 0; fi24 < 6; fi24++ {
        switch (v15 + "hello") {
        case "hello":
            v25 := []int{584}
            _ = v25
        case "\n\t":
            if v23 {
                for fi26 := 0; fi26 < 11; fi26++ {
                    v27 := GenS2{F0: 42, F1: 100}
                    var v28 uint32 = (v17 * 347)
                    switch (v15 + v15) {
                    case "\n\t":
                        _ = len(v3)
                    case "":
                        if (!false) {
                            fmt.Println(v15)
                        }
                    case "a":
                        v6 = (v5 / (((v5 / (((^(v5))) | 1))) | 1))
                    default:
                        v23 = true
                    }
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                var v29 uint32 = 42
                for fi30 := 0; fi30 < 5; fi30++ {
                    var v31 uint16 = v2
                    v1 = 624
                    v32 := GenS3{F0: (^(fi24)), F1: 127, F2: 0}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                _ = v29
            } else {
                var v33 float32 = 1e10
                var v34 uint = uint(v17)
                switch ((v15 + v15) + ((v15 + v15) + v15)) {
                case "a":
                    v3[v15] = v15
                case "hello":
                    v35 := GenS2{F0: v34, F1: 80}
                    _ = v35
                default:
                    var v36 int32 = 519
                    _ = v36
                }
                _ = v33
                _ = v34
            }
        case "foo bar":
            var v37 float64 = (v16 + 3.14159)
            _ = v37
        default:
            _ = len(v3)
        }
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

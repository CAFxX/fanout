package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint16; F1 int8 }
func genFn0() int16 {
    var v7 int16 = -32768
    if (!(752 >= 884)) {
        v7 += (v7 + v7)
        if (2 >= 1) {
            v8 := GenS1{F0: 190, F1: 32767}
            _ = v8
        } else {
            var v9 int8 = 0
            _ = v9
        }
        for fi10 := 0; fi10 < 20; fi10++ {
            switch (fi10 + (fi10 | (fi10 - (fi10 - (fi10 | fi10))))) {
            case 2147483647:
                var v11 uint64 = 698
                _ = v11
            case 127:
                v12 := []int{(fi10 / ((fi10) | 1)), fi10}
                _ = v12
            }
            v13 := []int{(fi10 | (fi10 << ((4294967295) & 63))), (fi10 - fi10), 128}
            v7 = v7
            _ = fi10
            _ = v13
        }
    } else {
        var v14 uint64 = 18446744073709551615
        for fi15 := 0; fi15 < 5; fi15++ {
            fmt.Println(fi15, v7)
            var v16 string = "hello"
            var v17 float64 = 0.1
            _ = fi15
            _ = v16
            _ = v17
        }
        var v18 float64 = 1e10
        _ = v14
        _ = v18
    }
    if false {
        for fi19 := 0; fi19 < 11; fi19++ {
            v7 = (v7 & (-(v7)))
            if (!true) {
                switch int(v7) {
                case -179:
                    v7 = v7
                case 32768:
                    v7 -= (v7 >> ((2) & 63))
                case 759:
                    var v20 int = fi19
                    _ = v20
                }
            }
            var v21 uint = 255
            _ = fi19
            _ = v21
        }
    } else {
        if false {
            switch "\n\t" {
            case "foo bar":
                v22 := GenS2{F0: 2, F1: 1, F2: 42}
                _ = v22
            }
        } else {
            v7 = v7
            for fi23 := 0; fi23 < 20; fi23++ {
                for fi24 := 0; fi24 < 19; fi24++ {
                    fmt.Println(fi24, v7, fi24)
                    var v25 float32 = 3.14159
                    v26 := GenS2{F0: 128, F1: 873, F2: 255}
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                v7 -= (v7 - v7)
                _ = fi23
            }
            v27 := []int{0, 127}
            _ = v27
        }
    }
    fmt.Println(v7)
    _ = v7
    return (^(v7))
}

func main() {
    var v1 int = -909
    var v2 string = "\n\t"
    v3 := make(map[uint32]uint32)
    var v4 int32 = 2147483647
    var v5 int64 = 917
    var v6 float64 = 0.0
    var v28 uint64 = 4294967295
    v29 := []int{}
    for fi30 := 0; fi30 < 6; fi30++ {
        switch v2 {
        case "foo bar":
            if false {
                var v31 int = 66
                _ = v31
            } else {
                _ = len(v3)
            }
        default:
            fmt.Println(v5)
        }
        if len(v29) > 0 {
            v29[0] = 9223372036854775807
        }
        v4 -= v4
        _ = fi30
    }
    switch v1 {
    case 1:
        var v32 float32 = 0.0
        _ = v32
    case 0:
        v4 = (v4 | (^(v4)))
    case 4294967295:
        _ = len(v29)
    default:
        v33 := GenS3{F0: 0, F1: -128}
        _ = v33
    }
    var v34 int16 = 1
    var v35 int = 998
    var v36 float32 = (0.0 + 0.1)
    if len(v29) > 0 {
        v29[0] = 0
    }
    var v37 int16 = -1
    v3[2] = 0
    fmt.Println(v6)
    if true {
        if (0 <= 65535) {
            for fi38 := 0; fi38 < 5; fi38++ {
                v2 = v2
                _ = fi38
            }
            var v39 uint = 42
            _ = v39
        }
    }
    var v40 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}

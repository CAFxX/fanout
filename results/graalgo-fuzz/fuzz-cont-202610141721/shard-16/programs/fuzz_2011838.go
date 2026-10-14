package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = (float64(1.5) * float64(0.1))
    var v3 uint = 42
    var v4 int = 1
    var v5 string = ""
    var v6 uint64 = 1
    v7 := GenS1{F0: -268}
    switch ("" + ((v5 + v5) + "a")) {
    case "foo bar":
        if (!true) {
            var v8 uint8 = 255
            fmt.Println(v5, v2)
            _ = v8
        }
    case "a":
        var v9 string = (v5 + ((v5 + v5) + ((v5 + v5) + "\n\t")))
        _ = v9
    default:
        v6 += v6
    }
    fmt.Println(v2)
    switch 0 {
    case 0:
        var v10 int8 = -1
        _ = v10
    case 4294967295:
        var v11 string = "\n\t"
        _ = v11
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        v3 = v3
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 float64 = ((v2 * v2) - v2)
            v14 = (v2 * float64(1.5))
            var v15 int16 = 287
            _ = fi13
            _ = v14
            _ = v15
        }
        var v16 uint16 = 2
        _ = fi12
        _ = v16
    }
    if false {
        v17 := []int{v4, (v4 | (v4 + (v4 >> ((v3) & 63))))}
        var v18 int8 = 18
        var v19 int = v4
        _ = v17
        _ = v18
        _ = v19
    }
    var v20 uint32 = 65535
    if ((0.0 + 0.0) <= 1e10) {
        var v21 int16 = -775
        v22 := []int{(v4 & 0), v4, v4, v4}
        _ = v21
        _ = v22
    } else {
        v3 = v3
        v4 = 32767
        v5 = ("hello" + ((v5 + ("" + ((v5 + (v5 + ("" + v5))) + ""))) + v5))
    }
    if ((42 <= 2) || false) {
        v23 := GenS2{F0: (v6 * (v6 << ((1) & 63)))}
        _ = v23
    } else {
        if false {
            var v24 int16 = 635
            for fi25 := 0; fi25 < 17; fi25++ {
                var v26 float32 = (0.1 + 1e10)
                if true {
                    v24 += v24
                }
                _ = fi25
                _ = v26
            }
            _ = v24
        } else {
            v27 := GenS3{F0: (v6 & v6)}
            v28 := GenS1{F0: 1}
            v29 := []int{2147483647, 2147483647, (v4 / (((^(v4))) | 1))}
            _ = v27
            _ = v28
            _ = v29
        }
        var v30 float32 = (float32(v2) * 1.5)
        var v31 int16 = 1
        _ = v30
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]uint32)
    v3 := make(map[int]int)
    var v4 int16 = -958
    v5 := []int{42, 32768, 652, 2147483647}
    if true {
        var v6 float64 = 1.5
        v7 := GenS1{F0: 830}
        var v8 int = 32768
        _ = v6
        _ = v7
        _ = v8
    } else {
        v2[32768] = 908
    }
    v4 = (v4 % ((v4) | 1))
    var v9 uint64 = 18446744073709551615
    var v10 int16 = (v4 | 32767)
    var v11 uint16 = 65535
    v9 += 4294967295
    v4 -= v10
    if (true && (-2.25 <= float32(1.5))) {
        v9 -= (v9 - v9)
        var v12 int8 = int8(v11)
        var v13 bool = (!false)
        _ = v12
        _ = v13
    }
    v11 = v11
    v14 := v2[127]
    if false {
        for fi15 := 0; fi15 < 19; fi15++ {
            v11 = (v11 >> ((42) & 63))
            if false {
                switch "\n\t" {
                case "foo bar":
                    var v16 uint16 = v11
                    _ = v16
                }
            }
            v1["foo bar"] = ("a" + ("" + ""))
            _ = fi15
        }
        fmt.Println(v10)
    } else {
        fmt.Println(v14)
    }
    v11 = 0
    switch "\n\t" {
    case "":
        v9 = v9
    case "a":
        v17 := GenS1{F0: 1}
        _ = v17
    case "\n\t":
        for fi18 := 0; fi18 < 2; fi18++ {
            for fi19 := 0; fi19 < 7; fi19++ {
                var v20 string = "foo bar"
                _ = fi19
                _ = v20
            }
            if len(v5) > 0 {
                v5[0] = fi18
            }
            var v21 int32 = 42
            _ = fi18
            _ = v21
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint64 = 659
    var v2 int = -1
    var v3 uint32 = 4294967295
    var v4 float64 = (((1e10 + float64(float32(1.5))) * 0.1) - -2.25)
    var v5 uint8 = 2
    var v6 string = "\n\t"
    for fi7 := 0; fi7 < 4; fi7++ {
        v1 = (v1 / ((828) | 1))
        v8 := GenS1{F0: 974}
        v9 := []int{}
        _ = fi7
        _ = v8
        _ = v9
    }
    switch (v2 / ((42) | 1)) {
    case 9223372036854775807:
        v10 := GenS1{F0: 65535}
        _ = v10
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 int64 = -1
        v4 -= 3.14159
        v3 += 2147483647
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: 931}
    v14 := []int{(v2 - (v2 / ((128) | 1)))}
    v1 += 2
    var v15 int8 = -36
    switch v6 {
    case "foo bar":
        if (-824 >= (v2 >> ((0) & 63))) {
            v1 += v1
        } else {
            v15 += v15
            v14 = append(v14, (v2 * (v2 * v2)))
            fmt.Println(v1, v1)
        }
    case "":
        var v16 int32 = -389
        _ = v16
    case "a":
        v15 = (v15 - 42)
    default:
        var v17 uint8 = (^(v5))
        _ = v17
    }
    var v18 string = "hello"
    for fi19 := 0; fi19 < 10; fi19++ {
        v18 = v18
        _ = fi19
    }
    var v20 int16 = -1
    v3 += (^(v3))
    fmt.Println(v4, v5, v4)
    if false {
        fmt.Println(v15)
    } else {
        _ = len(v14)
        v21 := GenS1{F0: uint(v3)}
        _ = v21
    }
    switch ("hello" + v18) {
    case "":
        var v22 int = 1
        _ = v22
    case "\n\t":
        for fi23 := 0; fi23 < 2; fi23++ {
            switch v2 {
            case -328:
                var v24 float64 = v4
                _ = v24
            }
            v6 = v18
            _ = fi23
        }
    }
    var v25 uint64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

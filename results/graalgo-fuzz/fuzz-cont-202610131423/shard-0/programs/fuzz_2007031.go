package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    var v3 int16 = -1
    var v4 float64 = 3.14159
    var v5 int64 = 76
    v6 := GenS1{F0: 32768}
    if (((v4 * v4) + v4) != (v4 + float64(-2.25))) {
        var v7 uint = 2
        switch ("" + "\n\t") {
        case "a":
            v4 = v4
        case "\n\t":
            for fi8 := 0; fi8 < 7; fi8++ {
                switch ("\n\t" + "\n\t") {
                case "\n\t":
                    var v9 int64 = (v5 / (((v5 * 224)) | 1))
                    _ = v9
                case "":
                    v10 := GenS1{F0: (fi8 | -2147483648)}
                    _ = v10
                case "a":
                    v11 := GenS1{F0: fi8}
                    _ = v11
                }
                _ = fi8
            }
        }
        var v12 string = ("hello" + "a")
        _ = v12
        _ = v7
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        v3 = (v3 >> ((18446744073709551615) & 63))
        _ = fi13
    }
    fmt.Println(v5)
    fmt.Println(v4, v4)
    var v14 uint16 = 1
    v4 -= v4
    v15 := []int{673, 32768, -2147483648}
    var v16 uint64 = 774
    for fi17 := 0; fi17 < 13; fi17++ {
        var v18 string = "a"
        for fi19 := 0; fi19 < 15; fi19++ {
            fmt.Println(v16, fi17)
            _ = fi19
        }
        fmt.Println(v3, v18, v18)
        _ = fi17
        _ = v18
    }
    v14 -= (v14 ^ v14)
    var v20 uint = 587
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

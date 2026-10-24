package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 42
    var v3 uint16 = 585
    v4 := []uint32{2, 2, 2147483648}
    var v5 int8 = (v2 / ((v2) | 1))
    for fi6 := 0; fi6 < 20; fi6++ {
        if true {
            for fi7 := 0; fi7 < 9; fi7++ {
                var v8 bool = (false && (!(2147483647 == 4294967295)))
                if v8 {
                    switch "a" {
                    case "\n\t":
                        fmt.Println(fi7, v3)
                    }
                    var v9 uint16 = 827
                    v5 = (v5 + v2)
                    _ = v9
                }
                _ = len(v4)
                _ = fi7
                _ = v8
            }
            v10 := GenS1{F0: 0}
            _ = v10
        } else {
            v3 = 380
            v11 := GenS1{F0: v3}
            _ = v11
        }
        _ = fi6
    }
    var v12 int32 = 42
    v13 := GenS1{F0: (v3 & v3)}
    var v14 float32 = (0.1 * 1.5)
    var v15 int8 = (^(v5))
    _ = len(v4)
    var v16 float32 = 1.5
    if true {
        v14 = (v16 / v14)
        var v17 bool = (v2 >= v15)
        _ = v17
    } else {
        if (false && true) {
            v1[-829] = ""
            v4 = append(v4, 65535)
            var v18 uint8 = 255
            _ = v18
        }
        if (128 > 1) {
            for fi19 := 0; fi19 < 20; fi19++ {
                var v20 uint = 993
                v16 -= (3.14159 - (v16 / float32((1.5 - 0.1))))
                _ = len(v1)
                _ = fi19
                _ = v20
            }
            if false {
                switch 1 {
                case 4294967295:
                    var v21 string = "hello"
                    _ = v21
                }
                switch ("a" + "a") {
                case "\n\t":
                    v22 := GenS1{F0: (v3 ^ v3)}
                    _ = v22
                case "hello":
                    _ = len(v4)
                case "foo bar":
                    var v23 float64 = (3.14159 / (1e10 + float64(1.5)))
                    _ = v23
                }
            }
        } else {
            _ = len(v4)
        }
    }
    if (((2 > 42) && (1 < 42)) || false) {
        v24 := v1[128]
        var v25 uint64 = 1
        var v26 int32 = v12
        _ = v24
        _ = v25
        _ = v26
    }
    var v27 string = ("hello" + (("\n\t" + "\n\t") + ("hello" + "\n\t")))
    v2 -= -1
    var v28 int32 = v12
    v5 -= v15
    v1[395] = (((v27 + "hello") + (v27 + v27)) + v27)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}

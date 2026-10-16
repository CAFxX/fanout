package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int }

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[int]bool)
    v3 := []string{"", ("foo bar" + "hello"), "hello", ("" + "hello")}
    var v4 int = 0
    switch 2147483647 {
    case 742:
        v5 := GenS1{F0: 935, F1: (v1 >= (v1 + (v1 + 739))), F2: (v4 % ((v4) | 1))}
        _ = v5
    case -341:
        if (!false) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    _ = len(v2)
    for fi6 := 0; fi6 < 14; fi6++ {
        delete(v2, (fi6 + 218))
        _ = fi6
    }
    var v7 bool = true
    switch v4 {
    case 1:
        if (2 == 42) {
            var v8 int64 = -1
            _ = v8
        } else {
            fmt.Println(v7)
            v7 = v7
            v4 += (v4 & v4)
        }
    case 0:
        fmt.Println(v7)
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    switch "" {
    case "":
        switch "" {
        case "":
            fmt.Println(v4, v1, v7)
        default:
            v7 = v7
        }
    case "\n\t":
        var v9 int8 = 0
        _ = v9
    default:
        var v10 uint16 = 42
        _ = v10
    }
    v11 := GenS1{F0: 313, F1: (0 <= 128), F2: v4}
    var v12 uint8 = 201
    var v13 int32 = -1
    switch (^(v4)) {
    case 32767:
        v7 = (!(((v1 + (v1 + v1)) > v1) && v7))
    case 211:
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 float64 = (3.14159 + 3.14159)
            var v16 int64 = -1
            v17 := GenS1{F0: 32767, F1: (127 > v12), F2: v4}
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
    case -148:
        for fi18 := 0; fi18 < 1; fi18++ {
            if false {
                for fi19 := 0; fi19 < 10; fi19++ {
                    switch v4 {
                    case 1:
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                    }
                    var v20 float32 = 1e10
                    _ = fi19
                    _ = v20
                }
                v12 = v12
            }
            _ = fi18
        }
    }
    if true {
        var v21 bool = true
        _ = v21
    } else {
        if (!v7) {
            if v7 {
                var v22 int8 = 127
                _ = v22
            }
            if v7 {
                fmt.Println(v4, v12)
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}

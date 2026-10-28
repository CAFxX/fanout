package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
func genFn0() int16 {
    var v5 int8 = 42
    switch 32767 {
    case 32768:
        fmt.Println(v5)
    }
    var v6 bool = true
    v6 = ((654 != 65535) && v6)
    _ = v5
    _ = v6
    return -1
}

func main() {
    v1 := make(map[int]bool)
    var v2 bool = ("hello" == ("" + ("foo bar" + "foo bar")))
    v3 := make(map[int]bool)
    var v4 int16 = 671
    var v7 float64 = (0.1 + 1e10)
    var v8 uint16 = 42
    switch "" {
    case "a":
        var v9 float32 = (1e10 + 0.1)
        _ = v9
    default:
        fmt.Println(v2, v8)
    }
    var v10 string = "foo bar"
    var v11 float64 = v7
    var v12 int = 260
    v11 = v11
    v7 = (1.5 * v11)
    v13 := []int{}
    for fi14 := 0; fi14 < 15; fi14++ {
        v3[(v12 ^ v12)] = v2
        _ = fi14
    }
    var v15 int16 = 0
    if (4294967295 <= 910) {
        switch v10 {
        case "a":
            var v16 int8 = -98
            _ = v16
        }
        if (0.1 <= (1e10 - 1.5)) {
            delete(v1, v12)
            var v17 uint = 1
            v11 = 0.1
            _ = v17
        } else {
            switch 352 {
            case -2147483648:
                for fi18 := 0; fi18 < 8; fi18++ {
                    v3[(fi18 + 2147483647)] = v2
                    switch 42 {
                    case -493:
                        for fi19 := 0; fi19 < 6; fi19++ {
                            v11 = (v7 + 1.5)
                            _ = fi19
                        }
                    case -296:
                        var v20 bool = (false && (!false))
                        _ = v20
                    case 1:
                        var v21 int32 = 42
                        _ = v21
                    default:
                        switch v12 {
                        case -1:
                            delete(v1, v12)
                        case 4294967295:
                            fmt.Println(v7)
                        case -2147483648:
                            v22 := GenS1{F0: 0, F1: (v15 + 1)}
                            _ = v22
                        }
                    }
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    _ = fi18
                }
            }
        }
    } else {
        var v23 int64 = 1
        switch (v12 * v12) {
        case -9223372036854775808:
            fmt.Println(v23)
        }
        _ = v23
    }
    var v24 int64 = int64(v15)
    for fi25 := 0; fi25 < 13; fi25++ {
        delete(v1, v12)
        for fi26 := 0; fi26 < 10; fi26++ {
            v24 -= -326
            var v27 uint8 = 1
            _ = fi26
            _ = v27
        }
        _ = fi25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}

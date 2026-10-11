package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int32 }

func main() {
    var v1 bool = false
    v2 := make(map[string]int)
    var v3 float64 = 0.1
    v4 := []uint32{65535, 1000, 2147483647, 2}
    v1 = (v1 || (false && v1))
    if (0 >= -1) {
        _ = len(v4)
        v3 = 1.5
        v3 += 3.14159
    } else {
        switch "" {
        case "foo bar":
            for fi5 := 0; fi5 < 12; fi5++ {
                fmt.Println(v3, v1)
                if true {
                    var v6 int64 = 808
                    _ = v6
                } else {
                    delete(v2, "foo bar")
                    v7 := GenS1{F0: 42, F1: 1, F2: -2147483648}
                    _ = v7
                }
                _ = fi5
            }
        }
    }
    v1 = (-694 > 1000)
    fmt.Println(v1, v3)
    switch 211 {
    case 1:
        v3 = (-2.25 / ((v3 / v3) + (v3 * v3)))
    case 128:
        if (-9223372036854775808 >= 42) {
            for fi8 := 0; fi8 < 2; fi8++ {
                var v9 int64 = 9223372036854775807
                var v10 uint = 4294967295
                switch (fi8 % ((fi8) | 1)) {
                case 0:
                    var v11 bool = v1
                    _ = v11
                case 2147483647:
                    _ = len(v4)
                }
                _ = fi8
                _ = v10
                _ = v9
            }
        }
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        for fi13 := 0; fi13 < 11; fi13++ {
            var v14 int8 = 100
            v1 = (v1 && v1)
            _ = fi13
            _ = v14
        }
        var v15 uint8 = 255
        _ = fi12
        _ = v15
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        v1 = true
        _ = fi16
    }
    v17 := GenS1{F0: 1, F1: 255, F2: 1}
    for fi18 := 0; fi18 < 14; fi18++ {
        for fi19 := 0; fi19 < 11; fi19++ {
            fmt.Println(v1)
            _ = fi19
        }
        var v20 bool = (!false)
        _ = fi18
        _ = v20
    }
    v1 = ((-337 != 1000) && false)
    v1 = (21 != 127)
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 int32 = 2147483647
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}

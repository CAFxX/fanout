package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }

func main() {
    v1 := []uint32{4294967295, 648}
    var v2 string = ("hello" + "hello")
    var v3 float32 = 1e10
    fmt.Println(v3, v2)
    v3 = 1e10
    var v4 float64 = float64((v3 + v3))
    if false {
        switch 2147483647 {
        case 32767:
            switch -373 {
            case 9223372036854775807:
                for fi5 := 0; fi5 < 19; fi5++ {
                    v3 += (((3.14159 / float32(float64(1.5))) / v3) - ((1.5 + (((1.5 * (1e10 * (1e10 + 1e10))) + 1e10) - 0.1)) + ((v3 / (float32(v4) + 0.0)) - float32((v4 / v4)))))
                    _ = fi5
                }
            case 922:
                v6 := GenS1{F0: 0}
                _ = v6
            }
        case 128:
            var v7 int64 = 0
            _ = v7
        case 0:
            _ = len(v1)
        }
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 3; fi9++ {
            fmt.Println(fi8, fi8)
            switch (-(fi8)) {
            case -618:
                switch 2147483647 {
                case 57:
                    v3 = (v3 + (0.0 - v3))
                case -756:
                    v10 := GenS2{F0: 255}
                    _ = v10
                }
            case 0:
                fmt.Println(v4)
            }
            _ = fi9
        }
        _ = fi8
    }
    v4 += ((0.1 - v4) + (float64(v3) - v4))
    var v11 uint = 65535
    if false {
        v12 := GenS2{F0: 0}
        v11 += (^(v11))
        _ = v12
    } else {
        if (2 > 42) {
            v13 := GenS1{F0: 0}
            if len(v1) > 0 {
                v1[0] = 293
            }
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        for fi15 := 0; fi15 < 12; fi15++ {
            var v16 int16 = 42
            v17 := GenS1{F0: 218}
            _ = fi15
            _ = v16
            _ = v17
        }
        fmt.Println(v3, v11)
        var v18 int64 = -1
        _ = fi14
        _ = v18
    }
    var v19 int32 = -983
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}

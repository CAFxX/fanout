package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 uint32 = 65535
    var v2 bool = (!((((!true) && true) && (0 <= 1)) && (-32768 <= 1)))
    var v3 int32 = -1
    v4 := []int{-784, 665, 1, 230}
    fmt.Println(v2)
    v3 = (-(v3))
    fmt.Println(v2, v2, v2)
    v1 += v1
    _ = len(v4)
    v1 = v1
    v3 -= (v3 ^ v3)
    if true {
        for fi5 := 0; fi5 < 12; fi5++ {
            fmt.Println(v2, v3)
            var v6 string = ""
            if len(v4) > 0 {
                v4[0] = fi5
            }
            _ = fi5
            _ = v6
        }
        if len(v4) > 0 {
            v4[0] = 42
        }
    } else {
        v4 = append(v4, 71)
    }
    var v7 int16 = 42
    for fi8 := 0; fi8 < 14; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 int = fi8
            v11 := GenS1{F0: 0}
            var v12 int64 = 0
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        _ = fi8
    }
    fmt.Println(v3, v2, v2)
    if (!(442 == 2)) {
        v13 := GenS1{F0: 42}
        for fi14 := 0; fi14 < 18; fi14++ {
            v2 = ((1e10 >= 1.5) && true)
            switch (fi14 ^ -590) {
            case -1:
                switch fi14 {
                case 127:
                    fmt.Println(v1, v3, v1)
                }
            case 2147483647:
                for fi15 := 0; fi15 < 7; fi15++ {
                    v16 := GenS1{F0: 128}
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi15
                    _ = v16
                }
            }
            _ = fi14
        }
        fmt.Println(v7, v1, v3)
        _ = v13
    }
    _ = len(v4)
    if (!v2) {
        var v17 bool = v2
        for fi18 := 0; fi18 < 1; fi18++ {
            var v19 int8 = 93
            _ = fi18
            _ = v19
        }
        v17 = true
        _ = v17
    }
    fmt.Println(v1, v2)
    v1 -= v1
    v20 := GenS1{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v7
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 uint64 = 200
    var v2 float64 = 1.5
    v3 := make(map[uint32]uint32)
    var v4 float32 = 1.5
    var v5 uint8 = 42
    v6 := []int{-2147483648, -1, 4294967295, 4294967295}
    switch -1 {
    case -814:
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := GenS1{F0: (fi7 % (((fi7 + int(v1))) | 1))}
            _ = fi7
            _ = v8
        }
    }
    if len(v6) > 0 {
        v6[0] = 9223372036854775807
    }
    if true {
        for fi9 := 0; fi9 < 20; fi9++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi9
        }
        fmt.Println(v2, v2, v2)
    }
    var v10 int = -2147483648
    fmt.Println(v1, v2, v5)
    v5 -= v5
    for fi11 := 0; fi11 < 5; fi11++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v12 int32 = 1
        fmt.Println(v5, v10, fi11)
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        for fi14 := 0; fi14 < 20; fi14++ {
            v2 = -2.25
            var v15 float64 = v2
            if len(v6) > 0 {
                v6[0] = (v10 | (fi14 & 128))
            }
            v2 += 1.5
            _ = fi14
            _ = v15
        }
        v1 = (v1 ^ v1)
        _ = fi13
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        switch 1 {
        case -868:
            var v17 uint16 = 187
            _ = v17
        case -789:
            v10 = fi16
        case -944:
            v10 = (fi16 % (((v10 >> ((182) & 63))) | 1))
        }
        switch (v10 + (v10 % (((v10 * fi16)) | 1))) {
        case 907:
            fmt.Println(v1, v2, v5)
        case 9223372036854775807:
            for fi18 := 0; fi18 < 8; fi18++ {
                var v19 uint32 = 280
                fmt.Println(fi18, fi16, v10)
                _ = len(v3)
                _ = fi18
                _ = v19
            }
        }
        var v20 int8 = 127
        _ = fi16
        _ = v20
    }
    if true {
        v10 = (v10 + v10)
        v4 -= -2.25
        for fi21 := 0; fi21 < 16; fi21++ {
            v5 = 127
            delete(v3, 4294967295)
            _ = fi21
        }
    } else {
        v1 = v1
    }
    v10 -= (v10 * 82)
    v22 := GenS2{F0: 1}
    v6 = append(v6, v10)
    var v23 uint = 65535
    v23 = (v23 / ((298) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

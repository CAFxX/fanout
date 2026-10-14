package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]int)
    var v2 bool = false
    v3 := make(map[int]uint32)
    var v4 int64 = 0
    v5 := []uint32{563}
    _ = len(v3)
    if false {
        if len(v5) > 0 {
            v5[0] = 683
        }
        delete(v1, 1)
        for fi6 := 0; fi6 < 2; fi6++ {
            v4 += v4
            fmt.Println(v4)
            _ = fi6
        }
    }
    var v7 float32 = (float32((0.1 * (float64(0.0) - ((1.5 + (3.14159 - (1e10 / 0.1))) / 3.14159)))) * (0.0 / 1.0))
    delete(v1, 0)
    v7 = v7
    var v8 float64 = 1e10
    fmt.Println(v4)
    var v9 int64 = v4
    for fi10 := 0; fi10 < 8; fi10++ {
        for fi11 := 0; fi11 < 2; fi11++ {
            v4 += (v9 + (v9 ^ -1))
            if false {
                fmt.Println(fi11)
            }
            _ = fi11
        }
        var v12 bool = ((1 <= 183) && true)
        _ = len(v5)
        _ = fi10
        _ = v12
    }
    _ = len(v3)
    var v13 int32 = 1
    if ((1 >= 36) || true) {
        v13 = (v13 >> ((2) & 63))
        switch 32767 {
        case -287:
            var v14 bool = (true && (true || v2))
            _ = v14
        case 4294967295:
            if v2 {
                var v15 bool = false
                v15 = (0 > 904)
                _ = v15
            }
        default:
            v8 += (((float64((v7 / -2.25)) * ((v8 + 1e10) * float64(float32(v8)))) / -2.25) * 0.1)
        }
    } else {
        for fi16 := 0; fi16 < 2; fi16++ {
            if v2 {
                switch (fi16 << ((389) & 63)) {
                case 32767:
                    switch (fi16 | (fi16 * fi16)) {
                    case 127:
                        var v17 uint = 0
                        _ = v17
                    case 42:
                        v5 = append(v5, 943)
                    case 0:
                        v8 = (3.14159 * (float64(0.1) + float64(v7)))
                    default:
                        var v18 uint32 = 42
                        _ = v18
                    }
                }
                v13 = v13
            }
            _ = fi16
        }
        fmt.Println(v4, v2, v8)
    }
    switch "a" {
    case "\n\t":
        v1[0] = 750
    case "a":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "hello":
        v4 = (v9 ^ v4)
    }
    switch 0 {
    case 127:
        var v19 uint16 = 0
        _ = v19
    case 474:
        switch "foo bar" {
        case "":
            var v20 uint64 = 4294967295
            _ = v20
        }
    case 1:
        if (!(!v2)) {
            v4 = v4
        } else {
            for fi21 := 0; fi21 < 10; fi21++ {
                var v22 int16 = 0
                _ = fi21
                _ = v22
            }
            var v23 float32 = (v7 / float32(float64((((((float32(v8) * v7) / 1.0) * (v7 - float32(1.5))) * v7) * (v7 * v7)))))
            _ = v23
        }
    default:
        fmt.Println(v13, v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}

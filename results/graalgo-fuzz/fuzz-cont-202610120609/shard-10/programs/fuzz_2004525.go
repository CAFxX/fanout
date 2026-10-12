package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 479
    var v3 string = "foo bar"
    var v4 int32 = int32(v2)
    v5 := make(map[uint32]uint32)
    var v6 uint32 = uint32(v2)
    v1[(^(v6))] = -415
    v6 = (v6 >> ((182) & 63))
    switch v3 {
    case "foo bar":
        for fi7 := 0; fi7 < 11; fi7++ {
            fmt.Println(v2, v4, v6)
            v2 -= v2
            _ = fi7
        }
    case "hello":
        for fi8 := 0; fi8 < 1; fi8++ {
            v1[v6] = 83
            _ = fi8
        }
    case "":
        var v9 int = -1
        _ = v9
    default:
        if (!(18446744073709551615 < 0)) {
            var v10 int8 = -70
            _ = v10
        } else {
            switch 832 {
            case -737:
                v2 += 4294967295
            default:
                fmt.Println(v6, v2)
            }
        }
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 int8 = -1
        _ = fi11
        _ = v12
    }
    v4 += (v4 % ((v4) | 1))
    switch 128 {
    case -1:
        var v13 uint64 = v2
        _ = v13
    }
    fmt.Println(v2, v6)
    var v14 bool = (!((^(v2)) != v2))
    var v15 uint8 = 128
    v14 = (v14 || v14)
    v14 = (int8(v2) > 1)
    var v16 uint64 = 441
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int64, p1 uint8, p2 int16) uint8 {
    p2 = (p2 / ((p2) | 1))
    p1 = p1
    fmt.Println(p0)
    v6 := []int{4294967295, 644, -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}
func genFn1(p0 int64) uint16 {
    var v7 uint16 = 32767
    p0 -= (p0 << ((1) & 63))
    _ = p0
    _ = v7
    return v7
}

func main() {
    var v1 float32 = (((0.0 / (1e10 + (0.1 / float32(0.1)))) + 1.5) - (3.14159 - (0.0 + (0.0 + ((((1.5 / 0.1) / -2.25) * 1e10) + (0.1 + (1e10 - ((3.14159 - 3.14159) - float32((-2.25 - 3.14159))))))))))
    var v2 bool = (1000 > -1)
    var v3 uint8 = 243
    var v4 string = ("" + ("" + ("" + ("hello" + ("\n\t" + "a")))))
    v5 := make(map[uint32]uint32)
    v8 := GenS1{F0: 526}
    switch "\n\t" {
    case "":
        for fi9 := 0; fi9 < 16; fi9++ {
            v4 = ("hello" + "hello")
            _ = fi9
        }
    default:
        for fi10 := 0; fi10 < 6; fi10++ {
            v3 = (^(v3))
            _ = fi10
        }
    }
    _ = len(v5)
    fmt.Println(v3, v2, v2)
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := GenS1{F0: 986}
        v1 -= 0.0
        var v13 int16 = -298
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println(v2)
    v14 := v5[386]
    v3 -= (v3 & v3)
    for fi15 := 0; fi15 < 16; fi15++ {
        var v16 int32 = 488
        if (2 == 2) {
            fmt.Println(v2)
            fmt.Println(v3, v14)
        }
        switch fi15 {
        case -9223372036854775808:
            var v17 int8 = 61
            _ = v17
        case 2147483647:
            v2 = ((v2 || false) || v2)
        }
        _ = fi15
        _ = v16
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}

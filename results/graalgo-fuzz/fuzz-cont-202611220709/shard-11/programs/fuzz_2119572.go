package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int32 }

func main() {
    var v1 int = -2147483648
    var v2 int = v1
    v3 := []int{(v2 & (v1 << ((42) & 63)))}
    var v4 uint8 = 0
    var v5 uint = 18446744073709551615
    v6 := make(map[uint32]string)
    var v7 string = "a"
    v5 += (^(v5))
    switch ((v7 + (("" + (v7 + "a")) + (v7 + "\n\t"))) + v7) {
    case "":
        switch v1 {
        case 2147483647:
            v6[2] = ""
        case 9223372036854775807:
            v3 = append(v3, v1)
        }
    case "a":
        var v8 bool = ((-(v1)) >= v1)
        _ = v8
    }
    if (true && (3.14159 > 0.0)) {
        for fi9 := 0; fi9 < 14; fi9++ {
            if true {
                if len(v3) > 0 {
                    v3[0] = (^(v1))
                }
            }
            _ = fi9
        }
        fmt.Println(v2, v4, v2)
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 uint16 = 0
            v3 = append(v3, fi10)
            _ = fi10
            _ = v11
        }
    }
    v2 += 4294967295
    v1 -= v2
    var v12 int8 = -113
    v5 += 0
    v5 -= (v5 % (((^(v5))) | 1))
    fmt.Println(v1, v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v13 := GenS1{F0: v5, F1: 0, F2: 2147483647}
    var v14 uint8 = uint8(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 float32 = 0.0
    v2 := make(map[uint32]int)
    var v3 uint32 = 320
    v4 := []int{127, 818, -904, -9223372036854775808}
    v5 := make(map[int]bool)
    var v6 int = 9223372036854775807
    _ = len(v4)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v3, v6, v1)
    v7 := GenS1{F0: 1}
    switch ("foo bar" + ((("\n\t" + "foo bar") + "\n\t") + "a")) {
    case "\n\t":
        v8 := GenS1{F0: 42}
        _ = v8
    case "hello":
        if len(v4) > 0 {
            v4[0] = (v6 - v6)
        }
    }
    var v9 uint32 = 2147483647
    v10 := GenS1{F0: -128}
    for fi11 := 0; fi11 < 14; fi11++ {
        if false {
            var v12 float32 = 1.5
            switch fi11 {
            case 42:
                if ((-40 != -1) && (false || (v3 != (^(v3))))) {
                    for fi13 := 0; fi13 < 14; fi13++ {
                        switch v6 {
                        case -376:
                            var v14 bool = false
                            _ = v14
                        case 128:
                            var v15 uint64 = 18446744073709551615
                            _ = v15
                        }
                        _ = fi13
                    }
                } else {
                    var v16 uint8 = 2
                    _ = v16
                }
            case 4294967295:
                v9 = (v9 << ((42) & 63))
            case 457:
                v12 = v1
            }
            delete(v2, v3)
            _ = v12
        }
        v9 = v9
        _ = fi11
    }
    var v17 int8 = 13
    fmt.Println(v1, v3)
    switch (v6 | 15) {
    case 0:
        v18 := v2[v3]
        _ = v18
    case -89:
        for fi19 := 0; fi19 < 1; fi19++ {
            var v20 int8 = (v17 - (v17 ^ v17))
            v4 = append(v4, (v6 | v6))
            var v21 int32 = -1
            _ = fi19
            _ = v20
            _ = v21
        }
    case 2147483647:
        for fi22 := 0; fi22 < 9; fi22++ {
            var v23 int = (v6 + (v6 - 128))
            _ = fi22
            _ = v23
        }
    default:
        v24 := v5[(v6 << ((2) & 63))]
        _ = v24
    }
    var v25 float64 = (0.1 - 1e10)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v26 float64 = v25
    delete(v5, v6)
    v27 := GenS1{F0: (v17 * v17)}
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 uint64 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    v3 := make(map[string]string)
    var v4 int = -1
    if ("\n\t" <= "hello") {
        v4 = v4
    } else {
        v5 := GenS1{F0: 1, F1: -128}
        _ = v5
    }
    fmt.Println(v4)
    v6 := GenS1{F0: 215, F1: 127}
    var v7 int32 = -2147483648
    var v8 uint64 = 68
    var v9 uint32 = 1
    fmt.Println(v9, v7)
    v10 := GenS2{F0: 128}
    for fi11 := 0; fi11 < 2; fi11++ {
        switch v4 {
        case -415:
            var v12 uint64 = (^(v8))
            _ = v12
        case 32768:
            if false {
                v7 += v7
                for fi13 := 0; fi13 < 10; fi13++ {
                    fmt.Println(v4, v4)
                    if false {
                        var v14 uint32 = v9
                        _ = v14
                    } else {
                        var v15 uint32 = 2147483647
                        var v16 uint8 = 160
                        _ = v15
                        _ = v16
                    }
                    _ = fi13
                }
                delete(v1, fi11)
            }
        default:
            fmt.Println(v9)
        }
        for fi17 := 0; fi17 < 15; fi17++ {
            v18 := v1[fi17]
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 19; fi19++ {
            fmt.Println(v7, v8)
            v20 := GenS2{F0: 195}
            _ = fi19
            _ = v20
        }
        _ = fi11
    }
    if ((32767 == 931) || false) {
        if false {
            switch "hello" {
            case "":
                if true {
                    if (206 >= 255) {
                        var v21 bool = (1e10 != ((-2.25 - (-2.25 + (-2.25 - -2.25))) * -2.25))
                        for fi22 := 0; fi22 < 1; fi22++ {
                            fmt.Println(v9, v7)
                            _ = fi22
                        }
                        if (!v21) {
                            delete(v2, (v9 * (^(v9))))
                            v23 := GenS3{F0: 241, F1: (v8 + 4294967295)}
                            _ = v23
                        }
                        _ = v21
                    }
                    if (!false) {
                        if false {
                            fmt.Println(v7)
                        }
                    }
                }
            case "a":
                v8 += v8
            case "\n\t":
                v24 := []int{856, (-(v4)), 9223372036854775807, 4294967295}
                _ = v24
            default:
                var v25 float64 = -2.25
                _ = v25
            }
        } else {
            v9 -= (v9 ^ (v9 << ((1) & 63)))
        }
        switch 1 {
        case 127:
            v9 -= v9
        case 2147483647:
            v7 = -1
        case 4294967295:
            var v26 uint8 = uint8(v9)
            _ = v26
        }
        for fi27 := 0; fi27 < 7; fi27++ {
            fmt.Println(fi27, v9)
            _ = fi27
        }
    } else {
        var v28 uint16 = 1
        for fi29 := 0; fi29 < 7; fi29++ {
            var v30 int8 = int8(v28)
            _ = fi29
            _ = v30
        }
        _ = v28
    }
    fmt.Println(v4)
    var v31 int32 = (v7 >> ((255) & 63))
    var v32 uint32 = v9
    fmt.Println(v4, v9)
    v9 -= 4294967295
    v32 = v9
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

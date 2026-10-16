package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }

func main() {
    var v1 string = ("" + ("" + "foo bar"))
    var v2 uint8 = 255
    var v3 int64 = int64(v2)
    v4 := make(map[int]string)
    var v5 int64 = v3
    var v6 uint32 = 626
    switch 96 {
    case -9223372036854775808:
        for fi7 := 0; fi7 < 10; fi7++ {
            delete(v4, (fi7 - (fi7 & (fi7 / (((-(fi7))) | 1)))))
            var v8 uint = 648
            _ = fi7
            _ = v8
        }
    case -2147483648:
        v5 -= v5
    case 177:
        fmt.Println(v3)
    default:
        if true {
            for fi9 := 0; fi9 < 17; fi9++ {
                for fi10 := 0; fi10 < 18; fi10++ {
                    v2 = v2
                    var v11 int16 = -904
                    v11 = v11
                    _ = fi10
                    _ = v11
                }
                v12 := []int{(fi9 % (((fi9 % (((fi9 << ((2) & 63))) | 1))) | 1)), (fi9 & fi9), (^(fi9))}
                _ = fi9
                _ = v12
            }
        } else {
            for fi13 := 0; fi13 < 1; fi13++ {
                fmt.Println(v3)
                for fi14 := 0; fi14 < 3; fi14++ {
                    v15 := GenS1{F0: 2147483647}
                    _ = fi14
                    _ = v15
                }
                _ = fi13
            }
            switch 884 {
            case -65:
                if false {
                    v1 = v1
                }
            }
        }
    }
    v16 := GenS1{F0: -1}
    if false {
        v1 = (v1 + "hello")
    }
    v2 = 127
    v5 -= -1
    v17 := GenS2{F0: v1}
    for fi18 := 0; fi18 < 6; fi18++ {
        for fi19 := 0; fi19 < 20; fi19++ {
            v2 = v2
            v4[fi19] = "\n\t"
            _ = fi19
        }
        v20 := GenS1{F0: 42}
        var v21 uint32 = (v6 >> ((42) & 63))
        _ = fi18
        _ = v20
        _ = v21
    }
    var v22 float32 = 1e10
    switch v1 {
    case "\n\t":
        var v23 float32 = ((v22 / 1.0) / (0.1 - (((v22 - 3.14159) / 1.0) / v22)))
        _ = v23
    default:
        var v24 uint16 = 2
        _ = v24
    }
    v5 = (v3 >> ((42) & 63))
    v25 := GenS2{F0: v1}
    switch 2147483647 {
    case 128:
        if (9223372036854775807 == v3) {
            var v26 uint32 = 2
            for fi27 := 0; fi27 < 1; fi27++ {
                v3 = 0
                _ = fi27
            }
            _ = v26
        }
    }
    if (2147483647 == v6) {
        v22 += (v22 * v22)
        var v28 float64 = float64(-2.25)
        _ = v28
    } else {
        var v29 uint64 = 701
        var v30 uint8 = (v2 + v2)
        v31 := []int{-1}
        _ = v29
        _ = v30
        _ = v31
    }
    var v32 string = v1
    v33 := GenS2{F0: "a"}
    v34 := GenS2{F0: v32}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}

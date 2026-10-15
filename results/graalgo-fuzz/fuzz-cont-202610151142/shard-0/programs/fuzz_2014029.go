package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16 }
func genFn0() uint16 {
    var v7 bool = true
    var v8 uint32 = 2147483648
    _ = v7
    _ = v8
    return 255
}

func main() {
    var v1 int8 = 6
    var v2 uint32 = 2147483648
    v3 := []uint32{2147483647, v2, v2}
    var v4 uint16 = 65535
    var v5 uint16 = v4
    v6 := []uint32{(^(v2)), (v2 | 704), (v2 / (((^(v2))) | 1)), (v2 / ((936) | 1))}
    v4 = (v5 - 165)
    v9 := GenS1{F0: 42, F1: 0, F2: -32768}
    var v10 int = -665
    v6 = append(v6, 441)
    switch v10 {
    case 4294967295:
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 uint = 328
            var v13 bool = ("" > ("a" + "\n\t"))
            _ = fi11
            _ = v12
            _ = v13
        }
    default:
        fmt.Println(v10, v4)
    }
    v14 := GenS2{F0: v5}
    for fi15 := 0; fi15 < 3; fi15++ {
        v5 = (v5 ^ 0)
        _ = fi15
    }
    if (-1 == 214) {
        if false {
            fmt.Println(v4, v5)
            switch "" {
            case "":
                _ = len(v3)
            }
        } else {
            v5 -= (v5 / ((255) | 1))
            v16 := GenS3{F0: (v5 << ((1) & 63))}
            v1 -= (v1 ^ v1)
            _ = v16
        }
    } else {
        v17 := GenS2{F0: v4}
        for fi18 := 0; fi18 < 12; fi18++ {
            v19 := GenS1{F0: 42, F1: 42, F2: 42}
            _ = fi18
            _ = v19
        }
        _ = v17
    }
    if ((!true) && ((genFn0() < v4) || (false && (((-2.25 >= 0.1) && ("" < "foo bar")) || (-563 < -1))))) {
        var v20 uint16 = v4
        v10 = -286
        var v21 uint32 = v2
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        for fi23 := 0; fi23 < 18; fi23++ {
            var v24 int16 = 888
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    fmt.Println(v5)
    switch (("" + "hello") + "a") {
    case "\n\t":
        fmt.Println(v4)
    case "a":
        fmt.Println(v1, v10, v1)
    }
    var v25 uint64 = 4294967295
    if false {
        if (!false) {
            v5 = genFn0()
            if len(v6) > 0 {
                v6[0] = uint32(v5)
            }
            var v26 bool = (1e10 >= 1e10)
            _ = v26
        }
        for fi27 := 0; fi27 < 11; fi27++ {
            _ = len(v6)
            v28 := GenS1{F0: -882, F1: -1, F2: 385}
            v4 -= (v5 << ((0) & 63))
            _ = fi27
            _ = v28
        }
    }
    if (((1e10 - 0.1) - 0.1) == ((1e10 + float64((float32(0.1) - 1.5))) / 1.0)) {
        var v29 int = v10
        if (!(-2147483648 >= -1)) {
            if len(v6) > 0 {
                v6[0] = 1
            }
            switch (v29 * v10) {
            case 42:
                v30 := GenS3{F0: v5}
                _ = v30
            default:
                fmt.Println(v5)
            }
        }
        _ = v29
    } else {
        var v31 int32 = -216
        var v32 int32 = v31
        _ = v31
        _ = v32
    }
    var v33 bool = (!(!(false || false)))
    v25 -= v25
    var v34 float32 = (0.1 + (1e10 / 1.5))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v25
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}

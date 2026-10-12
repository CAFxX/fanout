package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int; F1 int8; F2 uint8 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 int32; F1 uint16; F2 string }
type GenS6 struct { F0 uint32 }
func genFn0() uint {
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 int = fi7
        _ = fi7
        _ = v8
    }
    var v9 int = 9223372036854775807
    var v10 uint32 = 2147483648
    v11 := GenS1{F0: "\n\t", F1: 374}
    v12 := []int{(v9 >> ((552) & 63)), (v9 + v9), (v9 | (v9 / ((32768) | 1))), (v9 * (v9 + (v9 << ((537) & 63))))}
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 65535
}
func genFn1(p0 int32) uint16 {
    switch 2147483647 {
    case 9223372036854775807:
        v13 := GenS2{F0: ("\n\t" + "foo bar"), F1: false}
        _ = v13
    case 24:
        p0 += (-(p0))
    default:
        switch 9223372036854775807 {
        case 127:
            v14 := GenS3{F0: 954, F1: -59, F2: 0}
            _ = v14
        case -212:
            p0 -= 0
        }
    }
    var v15 float64 = 0.0
    _ = p0
    _ = v15
    return 255
}

func main() {
    var v1 uint32 = 2
    var v2 uint32 = v1
    v3 := make(map[int]uint32)
    var v4 uint64 = 0
    var v5 string = ""
    v6 := make(map[uint32]string)
    if false {
        if false {
            v4 = v4
            for fi16 := 0; fi16 < 1; fi16++ {
                _ = len(v6)
                _ = fi16
            }
        }
    }
    fmt.Println(v5)
    var v17 int = 1
    if true {
        for fi18 := 0; fi18 < 16; fi18++ {
            v4 = (v4 & 4294967295)
            _ = fi18
        }
        var v19 bool = (false && true)
        var v20 string = ""
        _ = v19
        _ = v20
    }
    v21 := GenS1{F0: ("a" + ((v5 + (v5 + (v5 + v5))) + ((v5 + "foo bar") + v5))), F1: (v4 | (v4 << ((65535) & 63)))}
    fmt.Println(v17, v5, v1)
    v22 := GenS4{F0: (v4 / ((v4) | 1)), F1: 0}
    switch 32768 {
    case -663:
        v23 := GenS5{F0: 42, F1: 65535, F2: v5}
        _ = v23
    }
    v24 := GenS5{F0: 235, F1: 2, F2: (v5 + "\n\t")}
    v25 := GenS6{F0: v2}
    if false {
        switch (v5 + v5) {
        case "foo bar":
            var v26 string = "foo bar"
            _ = v26
        case "a":
            v5 = ("\n\t" + (v5 + ((v5 + v5) + (v5 + v5))))
        case "":
            v4 = 42
        }
        fmt.Println(v2, v2)
        var v27 int8 = 100
        _ = v27
    } else {
        delete(v6, (^(v1)))
        fmt.Println(v17)
    }
    for fi28 := 0; fi28 < 11; fi28++ {
        var v29 int = (v17 | v17)
        var v30 int64 = 9223372036854775807
        _ = fi28
        _ = v29
        _ = v30
    }
    var v31 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}

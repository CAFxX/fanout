package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int16 }
type GenS2 struct { F0 string }
func genFn0(p0 int32, p1 int) uint16 {
    fmt.Println(p1, p0)
    p1 -= -9223372036854775808
    v6 := GenS1{F0: 417, F1: 1000, F2: -32768}
    v7 := []int{p1, (p1 + 4294967295), p1}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 533
}

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{2, 692, 781}
    var v3 int32 = -1
    var v4 int64 = 1
    v5 := []uint32{346}
    switch 882 {
    case 856:
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 int32 = 2147483647
            _ = fi8
            _ = v9
        }
    }
    var v10 uint32 = 1
    var v11 uint8 = 255
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 uint64 = 4294967295
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: 42, F1: 32767, F2: -391}
    v11 -= 2
    fmt.Println(v10, v10)
    fmt.Println(v11, v3)
    var v15 string = ("\n\t" + ("\n\t" + "a"))
    for fi16 := 0; fi16 < 5; fi16++ {
        if true {
            var v17 int8 = -1
            v18 := GenS1{F0: 882, F1: 1000, F2: 410}
            for fi19 := 0; fi19 < 8; fi19++ {
                v3 += v3
                fmt.Println(v3)
                _ = fi19
            }
            _ = v17
            _ = v18
        }
        v20 := GenS2{F0: ((((v15 + v15) + v15) + v15) + "hello")}
        _ = fi16
        _ = v20
    }
    if (42 != 18446744073709551615) {
        if false {
            v15 = v15
            v11 = (v11 ^ 178)
        } else {
            switch 779 {
            case 127:
                v10 = v10
            }
        }
        var v21 float32 = 1.5
        for fi22 := 0; fi22 < 15; fi22++ {
            fmt.Println(v21, v10)
            _ = fi22
        }
        _ = v21
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v23 bool = true
    var v24 bool = true
    fmt.Println(v10, v24, v15)
    switch -9223372036854775808 {
    case 128:
        switch -9223372036854775808 {
        case 902:
            var v25 int64 = v4
            _ = v25
        default:
            v15 = ("\n\t" + "hello")
        }
    case 4294967295:
        v26 := GenS1{F0: 630, F1: 1000, F2: 0}
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}

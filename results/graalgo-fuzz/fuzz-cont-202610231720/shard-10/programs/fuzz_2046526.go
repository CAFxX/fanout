package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 uint8; F1 int32; F2 int8 }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 uint64; F1 int8; F2 bool }

func main() {
    var v1 int32 = 1
    var v2 int32 = (v1 ^ 1)
    var v3 uint64 = 208
    var v4 string = "hello"
    v5 := make(map[uint32]bool)
    if (false && false) {
        v1 += v1
    }
    switch 127 {
    case 484:
        if true {
            v6 := []int{32768, 4294967295, -632, 0}
            v2 = v1
            _ = v6
        } else {
            if (32767 <= 65535) {
                v7 := GenS1{F0: 0, F1: 129, F2: (v3 - (^(v3)))}
                for fi8 := 0; fi8 < 3; fi8++ {
                    var v9 int16 = 0
                    v10 := GenS1{F0: 100, F1: 0, F2: 4294967295}
                    var v11 uint32 = 2147483647
                    _ = fi8
                    _ = v10
                    _ = v11
                    _ = v9
                }
                v12 := GenS1{F0: 47, F1: 201, F2: (v3 - (v3 | (^(v3))))}
                _ = v12
                _ = v7
            } else {
                v3 += (v3 * 0)
                v4 = (v4 + ((("hello" + v4) + ((v4 + "hello") + "foo bar")) + v4))
                for fi13 := 0; fi13 < 19; fi13++ {
                    v14 := []int{fi13, fi13, fi13}
                    v4 = (v4 + "\n\t")
                    v2 += v1
                    _ = fi13
                    _ = v14
                }
            }
        }
    default:
        v15 := []int{470, int(v1), 32768}
        _ = v15
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        v17 := GenS2{F0: 1, F1: fi16}
        var v18 uint32 = 920
        v18 += (^(v18))
        _ = fi16
        _ = v17
        _ = v18
    }
    if false {
        for fi19 := 0; fi19 < 15; fi19++ {
            v1 = (v1 << ((255) & 63))
            v20 := GenS3{F0: 30, F1: v1, F2: 127}
            v4 = v4
            _ = fi19
            _ = v20
        }
    }
    v4 = v4
    v21 := GenS4{F0: (v3 ^ 18446744073709551615), F1: 0}
    v22 := v5[46]
    fmt.Println(v2, v4)
    v23 := GenS4{F0: (^(v3)), F1: 0}
    v1 += v1
    v4 = v4
    if v22 {
        for fi24 := 0; fi24 < 8; fi24++ {
            var v25 bool = (v3 > 4294967295)
            v26 := GenS5{F0: v3, F1: -56, F2: ((!("hello" <= (((v4 + v4) + (v4 + v4)) + v4))) && false)}
            v25 = ((fi24 >= fi24) && v25)
            _ = fi24
            _ = v25
            _ = v26
        }
    }
    delete(v5, 42)
    fmt.Println(v1, v22, v1)
    var v27 uint8 = uint8(v3)
    if v22 {
        fmt.Println(v27, v4, v2)
        fmt.Println(v2, v2)
    } else {
        var v28 uint64 = (v3 - v3)
        _ = v28
    }
    if v22 {
        var v29 int8 = -77
        delete(v5, 42)
        _ = v29
    }
    if (true && (!v22)) {
        var v30 uint16 = 65535
        v27 += (v27 ^ (v27 * v27))
        for fi31 := 0; fi31 < 5; fi31++ {
            v1 = -27
            _ = fi31
        }
        _ = v30
    } else {
        var v32 uint32 = 4294967295
        switch ("foo bar" + v4) {
        case "foo bar":
            v33 := []int{0, 128}
            _ = v33
        case "\n\t":
            var v34 int16 = 531
            _ = v34
        case "hello":
            var v35 int8 = 21
            _ = v35
        }
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}

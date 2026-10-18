package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint8 }

func main() {
    var v1 uint8 = 2
    var v2 int32 = 1
    var v3 float32 = 3.14159
    v3 += v3
    for fi4 := 0; fi4 < 12; fi4++ {
        fmt.Println(v1)
        for fi5 := 0; fi5 < 9; fi5++ {
            if (true && (-1 == 465)) {
                v6 := []int{fi5, 263}
                _ = v6
            }
            var v7 bool = true
            if false {
                switch (fi5 >> ((2) & 63)) {
                case -2147483648:
                    var v8 string = ""
                    _ = v8
                }
            }
            _ = fi5
            _ = v7
        }
        _ = fi4
    }
    var v9 int64 = 932
    fmt.Println(v1, v1, v3)
    var v10 uint32 = 2147483647
    fmt.Println(v9, v10)
    v11 := []int{42, 42}
    for fi12 := 0; fi12 < 4; fi12++ {
        v13 := GenS1{F0: 4294967295, F1: 581, F2: (v1 / ((v1) | 1))}
        v11 = append(v11, fi12)
        v1 -= v1
        _ = fi12
        _ = v13
    }
    _ = len(v11)
    switch -1 {
    case 824:
        switch -387 {
        case -9223372036854775808:
            if (v2 >= -2147483648) {
                if false {
                    if true {
                        v1 += 117
                        var v14 uint8 = 128
                        fmt.Println(v2)
                        _ = v14
                    } else {
                        if len(v11) > 0 {
                            v11[0] = -9223372036854775808
                        }
                        fmt.Println(v2, v1, v10)
                        v15 := GenS1{F0: -1, F1: 18446744073709551615, F2: v1}
                        _ = v15
                    }
                    switch 42 {
                    case -1:
                        v2 -= -516
                    }
                    for fi16 := 0; fi16 < 17; fi16++ {
                        var v17 uint64 = 0
                        _ = fi16
                        _ = v17
                    }
                } else {
                    switch -2147483648 {
                    case 9223372036854775807:
                        var v18 uint64 = 845
                        _ = v18
                    case 2147483647:
                        for fi19 := 0; fi19 < 5; fi19++ {
                            fmt.Println(v3, v3, v3)
                            _ = fi19
                        }
                    case 925:
                        var v20 int8 = -1
                        _ = v20
                    default:
                        if false {
                            fmt.Println(v10, v1, v1)
                        }
                    }
                    if len(v11) > 0 {
                        v11[0] = 127
                    }
                    v21 := GenS1{F0: 0, F1: 0, F2: (v1 * v1)}
                    _ = v21
                }
                var v22 uint = 18446744073709551615
                v9 = (v9 ^ (v9 | v9))
                _ = v22
            } else {
                var v23 float32 = v3
                var v24 bool = (42 > 4294967295)
                _ = v23
                _ = v24
            }
        case 806:
            v10 += 915
        }
    }
    fmt.Println(v9, v1)
    v11 = append(v11, 32767)
    var v25 uint = 899
    v25 -= (v25 >> (((v25 / ((65535) | 1))) & 63))
    fmt.Println(v9, v1)
    var v26 int = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v9
}

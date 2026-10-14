package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 uint64, p1 uint64) uint64 {
    if true {
        p0 -= (p0 % (((p0 ^ (p0 * p1))) | 1))
        p1 += p0
        var v7 int32 = -2147483648
        _ = v7
    } else {
        if false {
            if true {
                fmt.Println(p0, p1)
                var v8 uint8 = 0
                if false {
                    p1 = (p0 << ((154) & 63))
                    var v9 uint64 = (^(p1))
                    for fi10 := 0; fi10 < 3; fi10++ {
                        for fi11 := 0; fi11 < 10; fi11++ {
                            v8 += (v8 ^ 169)
                            _ = fi11
                        }
                        v8 = (^(v8))
                        _ = fi10
                    }
                    _ = v9
                } else {
                    p0 = p1
                    var v12 string = "a"
                    _ = v12
                }
                _ = v8
            }
            v13 := GenS1{F0: p1}
            _ = v13
        } else {
            p1 = 1
        }
    }
    fmt.Println(p1)
    p1 -= p1
    p1 = p0
    var v14 int8 = 127
    _ = p0
    _ = p1
    _ = v14
    return (p0 * p0)
}

func main() {
    var v1 int64 = 173
    var v2 bool = (32768 != -87)
    var v3 uint64 = 1
    var v4 float32 = (0.0 / (0.0 - 1e10))
    v5 := []uint32{685, 42, 0}
    var v6 uint32 = 414
    fmt.Println(v2, v6)
    if (v4 > (3.14159 + 3.14159)) {
        var v15 uint = 65535
        for fi16 := 0; fi16 < 1; fi16++ {
            if true {
                var v17 int32 = -477
                _ = v17
            } else {
                for fi18 := 0; fi18 < 10; fi18++ {
                    v5 = append(v5, (v6 & 1))
                    v19 := GenS1{F0: (v3 & (v3 + v3))}
                    _ = fi18
                    _ = v19
                }
                switch -152 {
                case -806:
                    for fi20 := 0; fi20 < 15; fi20++ {
                        var v21 string = "a"
                        v15 = v15
                        _ = fi20
                        _ = v21
                    }
                case 996:
                    var v22 float32 = v4
                    _ = v22
                case 0:
                    fmt.Println(fi16, fi16, v1)
                }
            }
            fmt.Println(v2, v6, fi16)
            v5 = append(v5, v6)
            _ = fi16
        }
        _ = v15
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        var v24 int32 = -447
        fmt.Println(fi23)
        _ = fi23
        _ = v24
    }
    v25 := GenS1{F0: (v3 * (v3 & (v3 / (((v3 % ((genFn0(v3, (v3 & (^(v3))))) | 1))) | 1))))}
    v5 = append(v5, 2147483647)
    v5 = append(v5, 4294967295)
    var v26 int64 = v1
    v2 = true
    v27 := GenS1{F0: 1}
    v5 = append(v5, v6)
    var v28 uint16 = 376
    var v29 uint = 0
    if len(v5) > 0 {
        v5[0] = v6
    }
    if len(v5) > 0 {
        v5[0] = (v6 | v6)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

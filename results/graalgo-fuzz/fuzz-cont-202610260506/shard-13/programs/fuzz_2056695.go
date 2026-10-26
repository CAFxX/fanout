package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 uint; F2 bool }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 18446744073709551615
    v3 := []uint32{134, 2147483648, 65535, 2}
    var v4 int32 = -534
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v5 uint16 = 42
    v3 = append(v3, 1)
    v4 -= (v4 + (-(v4)))
    v2 = 752
    fmt.Println(v5)
    var v6 int8 = 42
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 uint = 18446744073709551615
        if true {
            if false {
                v9 := GenS1{F0: -476}
                _ = v9
            } else {
                for fi10 := 0; fi10 < 4; fi10++ {
                    v11 := GenS1{F0: 429}
                    _ = fi10
                    _ = v11
                }
                var v12 int8 = v6
                _ = v12
            }
        }
        _ = fi7
        _ = v8
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        for fi13 := 0; fi13 < 10; fi13++ {
            for fi14 := 0; fi14 < 4; fi14++ {
                _ = len(v3)
                if true {
                    v5 -= 65535
                    fmt.Println(v5)
                    if true {
                        var v15 int = (fi13 ^ (fi13 - fi14))
                        v1[286] = 399
                        fmt.Println(fi14, fi13)
                        _ = v15
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                } else {
                    v16 := GenS1{F0: 842}
                    var v17 int64 = 9223372036854775807
                    v18 := GenS2{F0: 4294967295}
                    _ = v16
                    _ = v17
                    _ = v18
                }
                _ = fi14
            }
            v6 += v6
            _ = fi13
        }
    } else {
        for fi19 := 0; fi19 < 13; fi19++ {
            for fi20 := 0; fi20 < 9; fi20++ {
                v6 = (-(v6))
                _ = fi20
            }
            var v21 float32 = (3.14159 - 3.14159)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi19
            _ = v21
        }
        v6 += v6
    }
    fmt.Println(v2, v5)
    var v22 string = ("\n\t" + "hello")
    if (!(0 > 42)) {
        for fi23 := 0; fi23 < 6; fi23++ {
            var v24 int8 = 0
            fmt.Println(v24)
            if len(v3) > 0 {
                v3[0] = 514
            }
            _ = fi23
            _ = v24
        }
        v6 = v6
    } else {
        fmt.Println(v4, v6)
        v25 := GenS3{F0: -560, F1: 4294967295, F2: true}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v25
    }
    fmt.Println(v5, v4)
    delete(v1, 42)
    if (v5 >= 827) {
        v26 := GenS1{F0: 674}
        _ = v26
    }
    fmt.Println(v22, v22, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

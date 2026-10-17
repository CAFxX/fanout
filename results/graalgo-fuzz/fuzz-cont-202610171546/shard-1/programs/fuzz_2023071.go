package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }

func main() {
    var v1 uint64 = 0
    v2 := []int{127}
    var v3 uint64 = (v1 ^ 42)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v4 uint16 = 0
    for fi5 := 0; fi5 < 2; fi5++ {
        v1 += (v3 & (v1 | (v1 * (v1 - 875))))
        _ = fi5
    }
    if false {
        if (0 != 147) {
            fmt.Println(v1)
            for fi6 := 0; fi6 < 15; fi6++ {
                var v7 bool = (-2.25 < 1.5)
                _ = fi6
                _ = v7
            }
        } else {
            var v8 int = 718
            for fi9 := 0; fi9 < 19; fi9++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch (-(fi9)) {
                case -907:
                    for fi10 := 0; fi10 < 15; fi10++ {
                        v1 = 42
                        _ = fi10
                    }
                }
                var v11 bool = (100 <= -1)
                _ = fi9
                _ = v11
            }
            _ = v8
        }
    }
    var v12 string = ((("" + "\n\t") + "a") + ("" + ""))
    fmt.Println(v12)
    v1 = v3
    switch 1 {
    case 192:
        var v13 int8 = -1
        _ = v13
    default:
        var v14 float64 = 0.1
        _ = v14
    }
    v3 = (v1 >> ((4294967295) & 63))
    var v15 string = v12
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 uint32 = 2147483648
        var v18 int16 = 488
        _ = fi16
        _ = v17
        _ = v18
    }
    _ = len(v2)
    if (42 > -1) {
        v3 -= v1
    } else {
        var v19 uint = 255
        for fi20 := 0; fi20 < 8; fi20++ {
            var v21 uint = 2
            fmt.Println(v12, fi20)
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        for fi23 := 0; fi23 < 16; fi23++ {
            v3 -= (v3 << ((9) & 63))
            if ((1e10 / float32(3.14159)) < 1e10) {
                v2 = append(v2, fi23)
                if true {
                    _ = len(v2)
                    var v24 uint16 = 255
                    v12 = (v15 + "\n\t")
                    v12 = (v15 + ("a" + v12))
                    v4 = (^(v24))
                    if true {
                        if false {
                            v25 := GenS1{F0: 42, F1: v4}
                            v26 := GenS1{F0: 0, F1: (v24 & v4)}
                            v24 = v24
                            _ = v25
                            _ = v26
                        }
                    } else {
                        v24 = (v4 >> ((194) & 63))
                    }
                    _ = v24
                }
            }
            _ = fi23
        }
        _ = fi22
    }
    for fi27 := 0; fi27 < 3; fi27++ {
        v28 := GenS1{F0: v1, F1: (v4 + (^(v4)))}
        _ = fi27
        _ = v28
    }
    var v29 int8 = 100
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}

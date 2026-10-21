package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[int]int)
    var v2 int64 = 0
    v3 := make(map[uint32]int)
    var v4 uint8 = 107
    var v5 int = 4294967295
    v6 := []int{v5}
    v2 = 9223372036854775807
    v7 := v3[2]
    var v8 uint = 709
    var v9 uint32 = 309
    if true {
        v5 = v7
        v6 = append(v6, -1)
    } else {
        for fi10 := 0; fi10 < 3; fi10++ {
            fmt.Println(v2, v9)
            _ = fi10
        }
        v6 = append(v6, (v5 << ((145) & 63)))
        switch "hello" {
        case "\n\t":
            for fi11 := 0; fi11 < 12; fi11++ {
                var v12 uint64 = 0
                v8 -= (v8 | (v8 * (v8 % ((v8) | 1))))
                v8 += 65535
                _ = fi11
                _ = v12
            }
        case "hello":
            if len(v6) > 0 {
                v6[0] = (-(v5))
            }
        case "":
            v13 := v3[(v9 & v9)]
            _ = v13
        default:
            var v14 bool = true
            _ = v14
        }
    }
    switch (v7 % (((-(v5))) | 1)) {
    case 32768:
        for fi15 := 0; fi15 < 6; fi15++ {
            for fi16 := 0; fi16 < 12; fi16++ {
                v6 = append(v6, fi15)
                var v17 int8 = 42
                v8 = (v8 ^ 2)
                if (!true) {
                    v9 -= (v9 ^ uint32(v8))
                } else {
                    v5 -= v7
                    fmt.Println(fi16, v7)
                    var v18 uint32 = v9
                    _ = v18
                }
                if false {
                    v9 -= v9
                } else {
                    v8 = v8
                    v4 = v4
                    var v19 int32 = -1
                    _ = v19
                }
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
    default:
        v5 -= -1
    }
    fmt.Println(v4, v7)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if len(v6) > 0 {
        v6[0] = (v5 ^ v5)
    }
    if len(v6) > 0 {
        v6[0] = 32767
    }
    switch ("\n\t" + "hello") {
    case "foo bar":
        if false {
            v9 = v9
            var v20 float64 = 1.5
            _ = v20
        }
    case "":
        var v21 int64 = (v2 & 9223372036854775807)
        _ = v21
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        var v23 uint32 = (v9 ^ uint32(v4))
        if true {
            v24 := GenS1{F0: 2147483647}
            if ((("\n\t" + "\n\t") + ((("\n\t" + "\n\t") + ("hello" + ("a" + "\n\t"))) + "")) < "") {
                v9 = (^(v23))
            } else {
                v4 += v4
            }
            v25 := GenS1{F0: 42}
            _ = v24
            _ = v25
        }
        var v26 uint = (^(v8))
        _ = fi22
        _ = v23
        _ = v26
    }
    if len(v6) > 0 {
        v6[0] = (v5 - v7)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

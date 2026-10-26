package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
func genFn0(p0 int32, p1 uint16) int64 {
    v5 := GenS1{F0: -469, F1: -9223372036854775808}
    var v6 int32 = int32(p1)
    p0 = 42
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 bool = (206 == 9223372036854775807)
        fmt.Println(fi7)
        var v9 int16 = -948
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        if true {
            v6 += p0
        } else {
            v11 := GenS2{F0: ""}
            var v12 int64 = -1
            _ = v11
            _ = v12
        }
        fmt.Println(fi10, p1)
        var v13 uint16 = (p1 % ((579) | 1))
        _ = fi10
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 bool = true
    var v2 int64 = 9223372036854775807
    var v3 string = "hello"
    v4 := make(map[uint32]uint32)
    v14 := GenS3{F0: v1}
    var v15 int8 = -128
    v16 := GenS1{F0: -582, F1: (v2 - v2)}
    var v17 float64 = (1e10 / float64((1e10 * ((1e10 + (1.5 - -2.25)) / -2.25))))
    for fi18 := 0; fi18 < 4; fi18++ {
        v17 = float64(0.0)
        _ = fi18
    }
    var v19 uint = 4294967295
    v1 = (125 >= 1)
    var v20 string = (v3 + ((((v3 + v3) + "hello") + v3) + v3))
    var v21 uint8 = 2
    v22 := v4[2147483647]
    for fi23 := 0; fi23 < 19; fi23++ {
        var v24 float64 = ((float64(3.14159) / v17) + float64(0.0))
        _ = fi23
        _ = v24
    }
    var v25 bool = false
    switch -9223372036854775808 {
    case 127:
        var v26 float64 = v17
        _ = v26
    case 2147483647:
        v25 = v1
    case -739:
        if ((1 != v22) || (0 == 42)) {
            if v1 {
                v21 = (^(v21))
                v19 -= 164
                switch ("hello" + ((v3 + "") + v20)) {
                case "\n\t":
                    v27 := GenS1{F0: 193, F1: (v2 << ((42) & 63))}
                    _ = v27
                case "":
                    for fi28 := 0; fi28 < 10; fi28++ {
                        var v29 int16 = 42
                        _ = fi28
                        _ = v29
                    }
                case "hello":
                    v25 = true
                }
            }
        }
    }
    v20 = v3
    delete(v4, v22)
    v30 := GenS3{F0: v25}
    var v31 bool = ((!((!v25) && (50 > 9223372036854775807))) && (v1 || v25))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}

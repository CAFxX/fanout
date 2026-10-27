package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int16 }

func main() {
    var v1 int16 = -303
    v2 := []int{-2147483648, -9223372036854775808, -9223372036854775808}
    var v3 uint = 42
    v4 := []string{"", "\n\t", (("\n\t" + "a") + "hello"), "\n\t"}
    var v5 uint8 = 31
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 uint = v3
            _ = fi7
            _ = v8
        }
        var v9 int32 = -1
        _ = fi6
        _ = v9
    }
    var v10 int64 = 1
    v2 = append(v2, 106)
    _ = len(v4)
    var v11 uint8 = v5
    v3 += (v3 ^ v3)
    var v12 float64 = 1.5
    switch "foo bar" {
    case "":
        v5 = v5
    case "a":
        for fi13 := 0; fi13 < 14; fi13++ {
            v14 := GenS1{F0: 532, F1: 0, F2: 1}
            _ = fi13
            _ = v14
        }
    case "\n\t":
        for fi15 := 0; fi15 < 14; fi15++ {
            if false {
                var v16 int64 = (v10 % ((v10) | 1))
                v10 += (v10 - v10)
                _ = v16
            }
            v5 = v5
            fmt.Println(v3, v1)
            _ = fi15
        }
    }
    var v17 int16 = (v1 & v1)
    for fi18 := 0; fi18 < 15; fi18++ {
        v4 = append(v4, ("foo bar" + "\n\t"))
        _ = fi18
    }
    var v19 uint8 = v5
    switch 2147483647 {
    case 127:
        v12 = float64((0.0 / 0.1))
    case 367:
        v20 := GenS1{F0: 65535, F1: 152, F2: 1}
        _ = v20
    case 2147483647:
        v21 := GenS1{F0: 305, F1: 42, F2: -1}
        _ = v21
    }
    var v22 string = "a"
    v23 := GenS2{F0: 18446744073709551615}
    var v24 float32 = 0.1
    switch 4 {
    case 0:
        if true {
            var v25 float64 = float64(v24)
            _ = v25
        } else {
            v26 := GenS3{F0: -117, F1: ((v22 + "hello") + ("" + "foo bar"))}
            v17 -= int16(v19)
            var v27 uint = (v3 - v3)
            _ = v26
            _ = v27
        }
    case 2147483647:
        for fi28 := 0; fi28 < 11; fi28++ {
            for fi29 := 0; fi29 < 8; fi29++ {
                fmt.Println(fi28)
                _ = fi29
            }
            var v30 string = "foo bar"
            switch -422 {
            case -634:
                if true {
                    switch (v30 + v22) {
                    case "foo bar":
                        if false {
                            v31 := GenS3{F0: 42, F1: v30}
                            _ = v31
                        }
                    case "":
                        v17 += -254
                    default:
                        v32 := GenS1{F0: 1, F1: 2147483648, F2: -2147483648}
                        _ = v32
                    }
                } else {
                    v1 = (^(v1))
                }
            }
            _ = fi28
            _ = v30
        }
    case 32768:
        if true {
            var v33 uint16 = 889
            switch 0 {
            case 128:
                var v34 int32 = 42
                _ = v34
            }
            _ = v33
        }
    }
    if (v24 < v24) {
        v35 := GenS2{F0: 301}
        for fi36 := 0; fi36 < 14; fi36++ {
            v37 := GenS1{F0: 255, F1: uint32(v3), F2: 42}
            v38 := GenS4{F0: v1}
            var v39 uint64 = 1
            _ = fi36
            _ = v37
            _ = v38
            _ = v39
        }
        _ = v35
    }
    fmt.Println(v10, v19, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}

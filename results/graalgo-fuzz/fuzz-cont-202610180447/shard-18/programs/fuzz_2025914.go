package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint16; F1 uint64 }

func main() {
    v1 := []string{"\n\t", "foo bar", "hello"}
    var v2 uint32 = 4294967295
    var v3 uint64 = 4294967295
    v4 := []uint32{65535}
    var v5 int32 = -2147483648
    var v6 int8 = -80
    v5 += 42
    if (128 == 255) {
        fmt.Println(v6, v2, v3)
    } else {
        var v7 uint16 = 107
        if (296 > int64(v7)) {
            v7 -= v7
        }
        switch 42 {
        case -2147483648:
            for fi8 := 0; fi8 < 3; fi8++ {
                var v9 uint8 = 215
                if true {
                    var v10 uint8 = 2
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    for fi11 := 0; fi11 < 19; fi11++ {
                        _ = len(v1)
                        var v12 uint64 = (v3 & uint64(v10))
                        var v13 bool = false
                        v14 := GenS1{F0: (v2 % ((v2) | 1)), F1: (-1 > int64(v2))}
                        if v13 {
                            v5 -= v5
                        }
                        switch ("" + "a") {
                        case "\n\t":
                            var v15 bool = (true || false)
                            _ = v15
                        }
                        var v16 string = ("a" + "foo bar")
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v14
                        _ = v16
                    }
                    _ = v10
                }
                var v17 bool = true
                _ = fi8
                _ = v17
                _ = v9
            }
        case 9223372036854775807:
            v18 := GenS1{F0: v2, F1: (!false)}
            _ = v18
        case 32767:
            switch int(v6) {
            case 127:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        default:
            if true {
                v19 := GenS1{F0: v2, F1: true}
                v20 := GenS2{F0: 164, F1: (v3 - v3)}
                _ = v19
                _ = v20
            } else {
                v3 -= 624
                var v21 bool = false
                var v22 float32 = 1e10
                _ = v21
                _ = v22
            }
        }
        _ = v7
    }
    for fi23 := 0; fi23 < 3; fi23++ {
        v1 = append(v1, ("hello" + "\n\t"))
        _ = fi23
    }
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    _ = len(v1)
    var v24 uint16 = 975
    var v25 uint = 2
    v25 += (v25 - (v25 - v25))
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

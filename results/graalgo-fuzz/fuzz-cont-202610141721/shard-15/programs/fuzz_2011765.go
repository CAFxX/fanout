package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 uint8 = 42
    var v2 bool = true
    var v3 int64 = 9223372036854775807
    v4 := []uint32{0}
    var v5 int32 = int32(v1)
    v6 := GenS1{F0: 143}
    for fi7 := 0; fi7 < 8; fi7++ {
        fmt.Println(fi7, v5, v1)
        _ = fi7
    }
    var v8 float32 = 1e10
    v5 -= (v5 + (v5 + -103))
    for fi9 := 0; fi9 < 7; fi9++ {
        for fi10 := 0; fi10 < 13; fi10++ {
            fmt.Println(fi10, v1)
            _ = fi10
        }
        v2 = ((v3 | v3) == (v3 >> ((4294967295) & 63)))
        _ = fi9
    }
    v11 := GenS1{F0: -32768}
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 uint32 = 2147483647
        var v14 float64 = 3.14159
        if false {
            v15 := GenS1{F0: 32767}
            _ = v15
        } else {
            if v2 {
                v8 = v8
                var v16 float64 = v14
                if v2 {
                    switch "" {
                    case "a":
                        var v17 uint8 = 127
                        _ = v17
                    case "":
                        v1 = (v1 >> ((937) & 63))
                    case "hello":
                        switch (fi12 | (^(fi12))) {
                        case 106:
                            v18 := GenS1{F0: 0}
                            _ = v18
                        case 128:
                            v19 := GenS1{F0: -32768}
                            _ = v19
                        default:
                            v20 := GenS1{F0: 959}
                            _ = v20
                        }
                    }
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                _ = v16
            } else {
                for fi21 := 0; fi21 < 7; fi21++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    switch fi21 {
                    case 508:
                        if ((v3 & v3) >= (v3 * v3)) {
                            v8 = (v8 / 0.1)
                        } else {
                            fmt.Println(v3, fi12)
                        }
                    case 139:
                        var v22 int32 = -691
                        _ = v22
                    }
                    _ = fi21
                }
                v3 -= v3
            }
            fmt.Println(v2, v14)
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    var v23 uint64 = 2
    var v24 float32 = v8
    v1 += v1
    for fi25 := 0; fi25 < 3; fi25++ {
        var v26 int = fi25
        _ = fi25
        _ = v26
    }
    if (!true) {
        for fi27 := 0; fi27 < 5; fi27++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v4 = append(v4, 2147483648)
            _ = len(v4)
            _ = fi27
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}

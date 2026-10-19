package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 string, p1 uint64) uint8 {
    for fi4 := 0; fi4 < 11; fi4++ {
        for fi5 := 0; fi5 < 10; fi5++ {
            p0 = "a"
            _ = fi5
        }
        _ = fi4
    }
    p1 -= (p1 / (((p1 ^ p1)) | 1))
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 127
}

func main() {
    v1 := []uint32{686}
    var v2 int8 = 127
    var v3 int32 = 42
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 int = (fi6 - (^(fi6)))
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: -32768}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v1)
    v3 += (v3 * (v3 - (^(v3))))
    fmt.Println(v3, v3)
    v3 = v3
    fmt.Println(v2, v2)
    var v9 int64 = int64(v3)
    switch ("\n\t" + "\n\t") {
    case "a":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println(v2)
    fmt.Println(v9, v2, v9)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v10 uint32 = 2
    var v11 float64 = 3.14159
    for fi12 := 0; fi12 < 13; fi12++ {
        switch "hello" {
        case "foo bar":
            switch fi12 {
            case -1:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                var v13 uint32 = 65535
                _ = v13
            }
        case "hello":
            if len(v1) > 0 {
                v1[0] = v10
            }
        default:
            v1 = append(v1, v10)
        }
        _ = fi12
    }
    if (("\n\t" + "hello") < "") {
        if false {
            var v14 int16 = -921
            v14 = int16(v3)
            if (!(1 != 42)) {
                var v15 uint64 = 1
                var v16 int64 = (v9 / ((-9223372036854775808) | 1))
                var v17 uint32 = v10
                _ = v15
                _ = v16
                _ = v17
            } else {
                var v18 int8 = v2
                var v19 int16 = (v14 << ((4294967295) & 63))
                _ = v18
                _ = v19
            }
            _ = v14
        } else {
            if true {
                var v20 int8 = (v2 ^ (v2 - 1))
                v21 := GenS2{F0: genFn0("\n\t", 4294967295)}
                _ = v20
                _ = v21
            } else {
                if false {
                    _ = len(v1)
                    switch "foo bar" {
                    case "hello":
                        v9 = (v9 * v9)
                    default:
                        var v22 uint32 = (v10 + v10)
                        _ = v22
                    }
                } else {
                    v1 = append(v1, (v10 * 234))
                }
                switch "\n\t" {
                case "hello":
                    var v23 float32 = 0.1
                    _ = v23
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
            fmt.Println(v10, v11)
        }
    } else {
        v3 = v3
    }
    var v24 float32 = 1e10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v3
    _ = v8
    _ = v9
}

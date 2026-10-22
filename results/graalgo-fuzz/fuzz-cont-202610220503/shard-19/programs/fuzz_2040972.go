package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int64, p1 int16, p2 uint64) uint64 {
    switch "\n\t" {
    case "":
        p0 += p0
    case "\n\t":
        var v7 float64 = -2.25
        _ = v7
    case "hello":
        v8 := GenS1{F0: 65535}
        _ = v8
    default:
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 int32 = 2147483647
            var v11 uint32 = 2147483648
            var v12 uint16 = 0
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    }
    v13 := GenS2{F0: "\n\t"}
    var v14 int32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return (p2 / (((p2 + p2)) | 1))
}
func genFn1(p0 uint16) uint {
    switch 2147483647 {
    case 39:
        for fi15 := 0; fi15 < 2; fi15++ {
            if (-279 > 1) {
                v16 := GenS1{F0: 2147483648}
                v17 := []int{}
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
                _ = v16
                _ = v17
            } else {
                var v18 int32 = 142
                _ = v18
            }
            _ = fi15
        }
    case 127:
        p0 = (p0 - p0)
    case 4294967295:
        var v19 int32 = 2147483647
        _ = v19
    }
    v20 := GenS3{F0: 1}
    var v21 float64 = (-2.25 + -2.25)
    p0 += (p0 / ((p0) | 1))
    _ = p0
    _ = v20
    _ = v21
    return 207
}

func main() {
    var v1 float64 = 1.5
    var v2 uint64 = 18446744073709551615
    var v3 string = "a"
    v4 := make(map[string]uint32)
    v5 := []int{4294967295}
    var v6 uint16 = 0
    for fi22 := 0; fi22 < 2; fi22++ {
        fmt.Println(v1, v6, v2)
        v3 = "hello"
        _ = fi22
    }
    if true {
        switch v3 {
        case "":
            switch -9223372036854775808 {
            case -9223372036854775808:
                delete(v4, ("foo bar" + "hello"))
            case 1:
                var v23 uint64 = (v2 - 4294967295)
                _ = v23
            default:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        }
    }
    var v24 int16 = 42
    v25 := v4["a"]
    v4[(("foo bar" + "a") + ("foo bar" + "a"))] = v25
    v1 = v1
    if true {
        v25 = 985
        v3 = v3
    } else {
        fmt.Println(v1)
        v4[v3] = (v25 | (v25 + 2147483647))
    }
    _ = len(v4)
    v26 := GenS3{F0: 369}
    if false {
        for fi27 := 0; fi27 < 11; fi27++ {
            v24 = 32767
            var v28 uint16 = 0
            _ = fi27
            _ = v28
        }
        var v29 int16 = (v24 >> ((912) & 63))
        _ = v29
    } else {
        var v30 string = "foo bar"
        if true {
            v4[v30] = 4294967295
            switch v30 {
            case "hello":
                for fi31 := 0; fi31 < 6; fi31++ {
                    v2 = (v2 >> ((1) & 63))
                    _ = fi31
                }
            default:
                switch 1 {
                case -322:
                    switch "" {
                    case "hello":
                        var v32 uint32 = 363
                        _ = v32
                    }
                case 4294967295:
                    _ = len(v4)
                }
            }
        } else {
            delete(v4, v30)
        }
        _ = v30
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

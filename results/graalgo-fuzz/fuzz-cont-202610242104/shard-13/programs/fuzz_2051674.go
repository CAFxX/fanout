package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int64 }
type GenS2 struct { F0 uint32; F1 string; F2 uint32 }
type GenS3 struct { F0 int; F1 uint; F2 int16 }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 bool = true
    v2 := make(map[string]string)
    var v3 uint64 = 4294967295
    var v4 bool = true
    v5 := []int{128}
    if (1.5 >= 1e10) {
        if v1 {
            for fi6 := 0; fi6 < 4; fi6++ {
                var v7 int32 = 0
                _ = fi6
                _ = v7
            }
            _ = len(v5)
            for fi8 := 0; fi8 < 6; fi8++ {
                for fi9 := 0; fi9 < 13; fi9++ {
                    v10 := v2["\n\t"]
                    if v4 {
                        delete(v2, "foo bar")
                    } else {
                        v10 = v10
                        v4 = v4
                    }
                    v11 := GenS1{F0: ((v10 + (("a" + ((v10 + (v10 + "\n\t")) + (v10 + v10))) + v10)) + v10), F1: 1000, F2: 42}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                var v12 uint8 = 119
                _ = fi8
                _ = v12
            }
        } else {
            v3 = 1
        }
    }
    v13 := GenS1{F0: "a", F1: 42, F2: -369}
    var v14 int32 = 0
    if (227 != 732) {
        v3 = v3
        var v15 int32 = v14
        _ = v15
    }
    var v16 uint64 = (v3 << ((18446744073709551615) & 63))
    for fi17 := 0; fi17 < 14; fi17++ {
        v16 = (^(v3))
        _ = fi17
    }
    fmt.Println(v14, v3)
    var v18 uint = 91
    for fi19 := 0; fi19 < 17; fi19++ {
        if true {
            _ = len(v5)
            switch fi19 {
            case 492:
                v20 := GenS2{F0: 183, F1: ((("foo bar" + ("foo bar" + "a")) + ("" + "foo bar")) + ""), F2: 2147483647}
                _ = v20
            case 32767:
                var v21 int16 = 395
                _ = v21
            case -526:
                v22 := GenS3{F0: fi19, F1: 545, F2: 466}
                _ = v22
            default:
                if (!(2147483648 >= 240)) {
                    var v23 bool = v1
                    _ = v23
                } else {
                    var v24 int16 = 1
                    if false {
                        fmt.Println(v24, v16)
                        if false {
                            v25 := GenS2{F0: 4294967295, F1: "hello", F2: 254}
                            _ = v25
                        }
                        _ = len(v5)
                    }
                    fmt.Println(v24, v24, v16)
                    _ = v24
                }
            }
        } else {
            v26 := GenS4{F0: 42}
            v27 := GenS3{F0: (fi19 * (fi19 % ((-9223372036854775808) | 1))), F1: v18, F2: -1}
            v3 = 18446744073709551615
            _ = v26
            _ = v27
        }
        v4 = (false || (v1 || (1 >= 2147483647)))
        fmt.Println(v3)
        _ = fi19
    }
    v18 += (^(v18))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

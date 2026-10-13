package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint64; F1 bool; F2 int }
type GenS3 struct { F0 uint; F1 uint16; F2 string }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 18446744073709551615
    v3 := []string{""}
    for fi4 := 0; fi4 < 13; fi4++ {
        var v5 uint64 = 0
        v6 := GenS1{F0: 255, F1: -25, F2: 18446744073709551615}
        var v7 uint16 = 563
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    var v8 bool = true
    v8 = v8
    var v9 int16 = 42
    for fi10 := 0; fi10 < 19; fi10++ {
        v3 = append(v3, ("a" + "foo bar"))
        v1[30] = 65535
        if (!((^(fi10)) <= 4294967295)) {
            v8 = false
            if false {
                for fi11 := 0; fi11 < 2; fi11++ {
                    if (128 >= fi10) {
                        v12 := GenS1{F0: 26, F1: 0, F2: uint64(v2)}
                        var v13 bool = v8
                        _ = len(v3)
                        _ = v12
                        _ = v13
                    }
                    _ = fi11
                }
                v2 += 18446744073709551615
                v3 = append(v3, ((("" + ("a" + "")) + ("" + "foo bar")) + "\n\t"))
            }
            var v14 float64 = 3.14159
            _ = v14
        }
        _ = fi10
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        v9 -= 42
        _ = fi15
    }
    v2 = v2
    switch 42 {
    case -2147483648:
        _ = len(v1)
    case 7:
        fmt.Println(v2, v8)
    default:
        v16 := GenS2{F0: 18446744073709551615, F1: ((v8 && false) && v8), F2: 4294967295}
        _ = v16
    }
    v17 := GenS3{F0: 4294967295, F1: 948, F2: "foo bar"}
    if len(v3) > 0 {
        v3[0] = (("a" + ("hello" + "")) + "hello")
    }
    var v18 int32 = -847
    if ((!true) && v8) {
        v19 := GenS3{F0: v2, F1: 65535, F2: "foo bar"}
        if v8 {
            delete(v1, -1)
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v20 int64 = 9223372036854775807
            _ = v20
        }
        var v21 uint = 0
        _ = v19
        _ = v21
    } else {
        v8 = (("hello" + "hello") < "\n\t")
        for fi22 := 0; fi22 < 7; fi22++ {
            if v8 {
                var v23 int32 = -2147483648
                _ = len(v1)
                var v24 string = ""
                _ = v23
                _ = v24
            }
            _ = fi22
        }
        v18 = (v18 / ((v18) | 1))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}

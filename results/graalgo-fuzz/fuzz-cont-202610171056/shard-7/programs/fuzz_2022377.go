package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 int16; F2 uint32 }
func genFn0(p0 int64, p1 uint16) string {
    v6 := []int{0}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v7 int32 = 1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return "hello"
}
func genFn1(p0 bool, p1 int8) bool {
    var v8 uint16 = 413
    v8 += v8
    if true {
        if (p0 && p0) {
            var v9 uint8 = 0
            _ = v9
        }
    } else {
        v8 -= (v8 & 794)
        v8 += 617
        var v10 uint16 = v8
        _ = v10
    }
    v11 := []int{4294967295, 9223372036854775807}
    _ = p0
    _ = p1
    _ = v11
    _ = v8
    return true
}

func main() {
    var v1 uint = 0
    v2 := make(map[string]uint32)
    v3 := make(map[int]string)
    var v4 uint = (v1 - v1)
    v5 := make(map[string]string)
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 int64 = 9223372036854775807
        v14 := v2["hello"]
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        v4 = v1
        _ = fi15
    }
    if false {
        for fi16 := 0; fi16 < 5; fi16++ {
            for fi17 := 0; fi17 < 17; fi17++ {
                switch (-(fi17)) {
                case 739:
                    var v18 int64 = 9223372036854775807
                    switch "" {
                    case "":
                        var v19 int = (fi16 - fi16)
                        _ = v19
                    default:
                        fmt.Println(fi17, fi17)
                    }
                    var v20 float64 = 3.14159
                    _ = v18
                    _ = v20
                }
                fmt.Println(fi16, fi17, fi17)
                _ = fi17
            }
            switch (fi16 | (fi16 ^ 128)) {
            case -2147483648:
                var v21 string = "\n\t"
                _ = v21
            default:
                v4 = (v1 % (((v4 | (v1 + 18446744073709551615))) | 1))
            }
            v22 := GenS1{F0: false, F1: 128}
            _ = fi16
            _ = v22
        }
        v1 = (^(v1))
        v23 := GenS2{F0: 4294967295}
        _ = v23
    }
    fmt.Println(v4, v4)
    v3[711] = ""
    var v24 bool = false
    for fi25 := 0; fi25 < 4; fi25++ {
        var v26 uint8 = 42
        _ = fi25
        _ = v26
    }
    v1 = v4
    v27 := GenS3{F0: 100, F1: -271, F2: 0}
    for fi28 := 0; fi28 < 17; fi28++ {
        var v29 uint32 = 42
        v4 += 299
        if (42 >= 42) {
            v30 := []int{529, fi28, (fi28 + (fi28 | fi28)), (fi28 << ((v4) & 63))}
            v31 := GenS3{F0: 44, F1: -32768, F2: 42}
            var v32 uint16 = uint16(v29)
            _ = v30
            _ = v31
            _ = v32
        }
        _ = fi28
        _ = v29
    }
    var v33 uint16 = uint16(v4)
    delete(v5, "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}

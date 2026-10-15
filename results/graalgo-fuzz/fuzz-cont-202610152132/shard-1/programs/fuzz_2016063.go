package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint32; F1 uint }
func genFn0() bool {
    var v6 bool = (-166 <= -1)
    var v7 bool = (42 <= 2)
    fmt.Println(v7)
    _ = v6
    _ = v7
    return (!v7)
}

func main() {
    v1 := make(map[int]int)
    var v2 int64 = -1
    v3 := []string{}
    var v4 string = ""
    var v5 uint8 = 1
    switch 2147483647 {
    case 169:
        if false {
            var v8 uint = 255
            v8 = (v8 * v8)
            if (v8 <= 18446744073709551615) {
                fmt.Println(v5)
            } else {
                var v9 uint64 = 18446744073709551615
                for fi10 := 0; fi10 < 10; fi10++ {
                    _ = len(v1)
                    v3 = append(v3, v4)
                    _ = fi10
                }
                _ = v9
            }
            _ = v8
        }
    case 243:
        switch 127 {
        case 32768:
            switch 2147483647 {
            case 127:
                var v11 float32 = (float32(3.14159) / (3.14159 + 1e10))
                _ = v11
            case -6:
                var v12 uint64 = 0
                _ = v12
            }
        default:
            v5 -= 42
        }
    }
    v1[9223372036854775807] = 32767
    switch -2147483648 {
    case 128:
        switch -2147483648 {
        case 2147483647:
            v2 = -9223372036854775808
        case 9223372036854775807:
            v13 := GenS1{F0: 0, F1: 141}
            _ = v13
        case -2147483648:
            var v14 int32 = 1
            _ = v14
        default:
            var v15 string = v4
            _ = v15
        }
    case 195:
        fmt.Println(v4, v2)
    }
    v16 := GenS1{F0: -605, F1: 849}
    switch v4 {
    case "hello":
        for fi17 := 0; fi17 < 16; fi17++ {
            v18 := v1[-2147483648]
            fmt.Println(v4, v2, v18)
            _ = fi17
            _ = v18
        }
    case "a":
        switch -2147483648 {
        case -9223372036854775808:
            fmt.Println(v4, v4, v4)
        }
    case "":
        if (127 >= 100) {
            switch 128 {
            case 1:
                v2 = (v2 % (((v2 * v2)) | 1))
            case 4294967295:
                if len(v3) > 0 {
                    v3[0] = "\n\t"
                }
            case -907:
                if true {
                    for fi19 := 0; fi19 < 14; fi19++ {
                        var v20 uint64 = 18446744073709551615
                        var v21 float64 = 1.5
                        var v22 uint64 = v20
                        _ = fi19
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    v23 := GenS1{F0: -2147483648, F1: 65535}
                    _ = v23
                } else {
                    fmt.Println(v4, v4, v4)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            default:
                fmt.Println(v4, v2, v5)
            }
        }
    }
    v2 = v2
    for fi24 := 0; fi24 < 2; fi24++ {
        delete(v1, -840)
        v25 := GenS2{F0: 0, F1: 65535}
        for fi26 := 0; fi26 < 6; fi26++ {
            var v27 int8 = -128
            _ = fi26
            _ = v27
        }
        _ = fi24
        _ = v25
    }
    if true {
        fmt.Println(v2, v5, v4)
        var v28 int32 = 1
        _ = v28
    } else {
        for fi29 := 0; fi29 < 9; fi29++ {
            var v30 int = fi29
            _ = fi29
            _ = v30
        }
        for fi31 := 0; fi31 < 13; fi31++ {
            var v32 uint16 = uint16(v5)
            _ = fi31
            _ = v32
        }
        var v33 int16 = 1
        _ = v33
    }
    v5 = v5
    fmt.Println(v5)
    switch (v4 + "foo bar") {
    case "":
        v34 := GenS2{F0: 198, F1: 1}
        _ = v34
    case "hello":
        var v35 int = 42
        _ = v35
    case "a":
        v5 += (^(v5))
    }
    _ = len(v3)
    var v36 uint16 = 1
    var v37 int = 1
    v38 := GenS1{F0: 1, F1: 0}
    v39 := GenS1{F0: 0, F1: 969}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}

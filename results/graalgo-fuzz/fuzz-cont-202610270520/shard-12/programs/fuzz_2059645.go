package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := []uint32{}
    var v2 bool = (42 == 16)
    var v3 float32 = (1.5 * 1e10)
    var v4 uint = 106
    var v5 bool = false
    switch -9223372036854775808 {
    case 32768:
        v3 += v3
    case -506:
        switch ("foo bar" + "\n\t") {
        case "\n\t":
            v5 = true
        case "a":
            if (812 < 220) {
                var v6 int = 0
                _ = v6
            }
        }
    default:
        switch 128 {
        case 2147483647:
            v2 = (!(42 <= -1))
        default:
            v5 = (42 != 1)
        }
    }
    v7 := GenS1{F0: -32768}
    var v8 bool = false
    _ = len(v1)
    if true {
        if v2 {
            if v2 {
                var v9 int64 = -9223372036854775808
                var v10 uint64 = 1
                _ = v10
                _ = v9
            } else {
                var v11 float32 = v3
                switch 127 {
                case -266:
                    v11 = v11
                case -255:
                    for fi12 := 0; fi12 < 13; fi12++ {
                        fmt.Println(fi12, v11)
                        _ = fi12
                    }
                case 85:
                    fmt.Println(v5)
                default:
                    for fi13 := 0; fi13 < 1; fi13++ {
                        fmt.Println(v11, v5, v2)
                        _ = fi13
                    }
                }
                _ = v11
            }
            v14 := GenS1{F0: -1}
            _ = v14
        }
        v8 = false
    } else {
        v2 = ((978 >= 2) && v5)
        for fi15 := 0; fi15 < 8; fi15++ {
            if v8 {
                switch 1 {
                case -638:
                    var v16 int32 = 0
                    _ = v16
                case 32767:
                    fmt.Println(v2, v8)
                default:
                    if (42 != 42) {
                        fmt.Println(v2, v2)
                    }
                }
            }
            _ = fi15
        }
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        v18 := GenS1{F0: 42}
        var v19 float64 = -2.25
        v3 -= ((3.14159 * (v3 + v3)) - (((v3 * float32(v19)) * v3) * (v3 * v3)))
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 uint16 = 2
    switch ("a" + "\n\t") {
    case "hello":
        switch 128 {
        case 9223372036854775807:
            var v21 int = -2147483648
            _ = v21
        }
    case "a":
        for fi22 := 0; fi22 < 16; fi22++ {
            var v23 float32 = 0.1
            v24 := GenS1{F0: 42}
            switch fi22 {
            case 1:
                v1 = append(v1, 2147483647)
            case 42:
                v8 = v5
            }
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    var v25 string = "a"
    _ = len(v1)
    var v26 float32 = v3
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v1 = append(v1, 33)
        var v27 int8 = 26
        _ = v27
    }
    switch 32 {
    case 32767:
        v26 += (((((v3 * 0.1) - v26) / v3) - v26) / float32((0.0 + ((3.14159 + (3.14159 + ((0.1 - -2.25) + (-2.25 / -2.25)))) + ((1e10 - 0.0) - float64(v3))))))
    default:
        v3 += float32(3.14159)
    }
    var v28 uint16 = v20
    v8 = ((!v2) || (!(1 <= 753)))
    fmt.Println(v2)
    v5 = (128 >= 1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}

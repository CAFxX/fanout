package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint; F2 bool }

func main() {
    var v1 uint = 1
    var v2 bool = (true || false)
    var v3 bool = v2
    v1 += v1
    v4 := []int{127}
    switch ("hello" + "hello") {
    case "a":
        var v5 string = ""
        _ = v5
    case "":
        fmt.Println(v3, v1, v3)
    default:
        for fi6 := 0; fi6 < 13; fi6++ {
            v3 = (926 < v1)
            v3 = (!v2)
            v7 := GenS1{F0: -2147483648, F1: 0, F2: 365}
            _ = fi6
            _ = v7
        }
    }
    if (v3 && true) {
        v8 := GenS1{F0: 0, F1: 777, F2: 42}
        switch "hello" {
        case "a":
            var v9 uint32 = 2
            _ = v9
        case "\n\t":
            v3 = v3
        case "hello":
            fmt.Println(v1)
        default:
            v2 = v3
        }
        _ = v8
    } else {
        if len(v4) > 0 {
            v4[0] = -981
        }
        if (v1 == 18446744073709551615) {
            for fi10 := 0; fi10 < 16; fi10++ {
                v4 = append(v4, 42)
                for fi11 := 0; fi11 < 9; fi11++ {
                    if v3 {
                        var v12 int16 = -1
                        _ = v12
                    }
                    v3 = v2
                    _ = fi11
                }
                _ = fi10
            }
            var v13 uint64 = 2
            v14 := GenS1{F0: 42, F1: (v13 * uint64(v1)), F2: 1}
            _ = v13
            _ = v14
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    if v3 {
        if v2 {
            var v15 int16 = 1000
            _ = v15
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi16 := 0; fi16 < 16; fi16++ {
            for fi17 := 0; fi17 < 13; fi17++ {
                v4 = append(v4, (-(fi17)))
                v3 = (false || true)
                var v18 bool = (-2147483648 <= 108)
                _ = fi17
                _ = v18
            }
            switch ("hello" + "\n\t") {
            case "hello":
                v4 = append(v4, fi16)
            }
            _ = fi16
        }
        v3 = (0 != 1)
    }
    var v19 uint8 = 127
    for fi20 := 0; fi20 < 10; fi20++ {
        v21 := GenS2{F0: -978}
        if v2 {
            if v3 {
                v19 = (v19 % ((v19) | 1))
            } else {
                fmt.Println(fi20, v19)
                var v22 int = fi20
                _ = v22
            }
            v2 = (-128 != 100)
        }
        _ = fi20
        _ = v21
    }
    fmt.Println(v3, v19, v1)
    var v23 int64 = -1
    if len(v4) > 0 {
        v4[0] = 0
    }
    var v24 int32 = 42
    fmt.Println(v1)
    v25 := GenS3{F0: 127, F1: 4294967295, F2: (float64(-2.25) > (3.14159 + 0.1))}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}

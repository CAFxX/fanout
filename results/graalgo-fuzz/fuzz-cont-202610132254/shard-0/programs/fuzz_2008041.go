package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int16 }
type GenS2 struct { F0 int8; F1 int32; F2 string }

func main() {
    v1 := []string{"a", ("foo bar" + ("foo bar" + "hello")), "hello", ("hello" + "a")}
    var v2 string = "a"
    var v3 string = v2
    v2 = v2
    fmt.Println(v3, v2)
    var v4 uint8 = 127
    if false {
        switch -9223372036854775808 {
        case -823:
            for fi5 := 0; fi5 < 2; fi5++ {
                if false {
                    var v6 uint32 = 1
                    _ = v6
                }
                var v7 string = (((v2 + v2) + (("foo bar" + v2) + "a")) + v3)
                var v8 int64 = 0
                _ = fi5
                _ = v7
                _ = v8
            }
        case 32768:
            for fi9 := 0; fi9 < 15; fi9++ {
                v4 = v4
                v1 = append(v1, v2)
                _ = fi9
            }
        case 32767:
            switch 161 {
            case 606:
                var v10 float64 = 1e10
                _ = v10
            case -146:
                if (!false) {
                    v4 = v4
                    v11 := GenS1{F0: 32767, F1: 4294967295, F2: -1}
                    _ = v11
                }
            }
        }
        if false {
            switch "" {
            case "foo bar":
                var v12 uint8 = v4
                _ = v12
            }
            var v13 uint32 = 65535
            v3 = "a"
            _ = v13
        }
        v14 := GenS2{F0: 46, F1: 752, F2: v3}
        _ = v14
    } else {
        _ = len(v1)
    }
    fmt.Println(v2, v2, v3)
    v15 := GenS2{F0: 0, F1: -2147483648, F2: v3}
    var v16 uint16 = 65535
    if len(v1) > 0 {
        v1[0] = v3
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        switch "" {
        case "":
            var v18 string = (v3 + v2)
            _ = v18
        }
        _ = fi17
    }
    v19 := GenS1{F0: v16, F1: 0, F2: 42}
    v16 = (^(v16))
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 int8 = 1
        _ = fi20
        _ = v21
    }
    var v22 int32 = 2147483647
    v23 := GenS2{F0: -1, F1: v22, F2: v3}
    v24 := GenS2{F0: -128, F1: v22, F2: (v2 + v3)}
    v25 := GenS2{F0: -128, F1: v22, F2: "a"}
    fmt.Println(v22, v4, v4)
    var v26 string = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int8; F1 int; F2 uint64 }
type GenS4 struct { F0 int }

func main() {
    var v1 int16 = 32767
    var v2 uint = 65535
    v3 := []int{565, -1, 4294967295, 1}
    v4 := GenS1{F0: v1, F1: 203}
    v3 = append(v3, 581)
    switch 970 {
    case 128:
        for fi5 := 0; fi5 < 5; fi5++ {
            v6 := GenS2{F0: v1, F1: v2}
            _ = fi5
            _ = v6
        }
    case -2147483648:
        if true {
            if false {
                var v7 uint8 = 2
                v3 = append(v3, 622)
                var v8 bool = false
                _ = v7
                _ = v8
            }
            var v9 uint = v2
            var v10 uint16 = 1
            _ = v10
            _ = v9
        } else {
            v11 := GenS3{F0: 1, F1: -408, F2: 18446744073709551615}
            v12 := GenS4{F0: 1}
            fmt.Println(v2)
            _ = v11
            _ = v12
        }
    }
    v3 = append(v3, 128)
    switch ("\n\t" + ((("a" + "foo bar") + (("hello" + "hello") + "")) + "")) {
    case "\n\t":
        switch 0 {
        case -2147483648:
            v1 = (v1 & v1)
        default:
            v2 = v2
        }
    case "foo bar":
        var v13 int32 = 1
        _ = v13
    case "":
        v1 += 0
    }
    v1 = v1
    var v14 uint16 = 1
    v15 := GenS1{F0: (v1 ^ 1000), F1: 127}
    v3 = append(v3, 45)
    var v16 uint32 = 348
    for fi17 := 0; fi17 < 2; fi17++ {
        v1 += 32767
        v16 += (^(v16))
        _ = fi17
    }
    v3 = append(v3, 42)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}

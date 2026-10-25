package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 bool; F1 uint16; F2 int16 }
type GenS4 struct { F0 uint32; F1 int32; F2 uint64 }

func main() {
    var v1 string = "foo bar"
    var v2 string = v1
    var v3 int64 = 9223372036854775807
    var v4 int64 = (v3 / (((v3 | v3)) | 1))
    var v5 string = (v1 + v1)
    v6 := GenS1{F0: 2}
    v1 = "\n\t"
    switch 0 {
    case -943:
        var v7 uint = 42
        _ = v7
    }
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 uint64 = 697
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 32767}
    v11 := []int{32768}
    v2 = "hello"
    var v12 uint64 = 0
    var v13 uint8 = 1
    v14 := GenS1{F0: 255}
    v12 = (v12 >> ((896) & 63))
    var v15 int32 = int32(v13)
    for fi16 := 0; fi16 < 8; fi16++ {
        if len(v11) > 0 {
            v11[0] = -2147483648
        }
        if false {
            v15 += -79
            for fi17 := 0; fi17 < 11; fi17++ {
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                _ = fi17
            }
            var v18 int8 = 1
            _ = v18
        }
        _ = fi16
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        fmt.Println(v13, v1)
        v15 = v15
        switch 32768 {
        case 42:
            for fi20 := 0; fi20 < 9; fi20++ {
                v21 := GenS1{F0: 671}
                _ = fi20
                _ = v21
            }
        case 9223372036854775807:
            if (!(v12 < (v12 ^ (^(v12))))) {
                v22 := GenS2{F0: "a", F1: 18446744073709551615}
                var v23 uint8 = v13
                v24 := GenS3{F0: (false || false), F1: 255, F2: -1}
                _ = v22
                _ = v23
                _ = v24
            } else {
                if len(v11) > 0 {
                    v11[0] = 42
                }
                v25 := GenS2{F0: ("foo bar" + v2), F1: v12}
                _ = v25
            }
        case 4294967295:
            for fi26 := 0; fi26 < 5; fi26++ {
                var v27 uint32 = 1
                _ = fi26
                _ = v27
            }
        }
        _ = fi19
    }
    for fi28 := 0; fi28 < 3; fi28++ {
        if len(v11) > 0 {
            v11[0] = fi28
        }
        _ = fi28
    }
    var v29 string = ("" + "hello")
    switch -1 {
    case 127:
        for fi30 := 0; fi30 < 10; fi30++ {
            if ((v15 << ((1) & 63)) == int32(fi30)) {
                v1 = (v29 + "")
            }
            var v31 int32 = v15
            _ = fi30
            _ = v31
        }
    case -9223372036854775808:
        switch ("\n\t" + ((v1 + v29) + v5)) {
        case "a":
            v32 := GenS4{F0: 2, F1: v15, F2: v12}
            _ = v32
        }
    case 341:
        v3 -= v4
    default:
        fmt.Println(v15, v12, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

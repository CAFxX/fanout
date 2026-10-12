package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int64 }

func main() {
    var v1 uint8 = 113
    var v2 bool = (("hello" > "foo bar") || (42 < (v1 & (v1 << ((18446744073709551615) & 63)))))
    var v3 uint64 = 42
    v4 := make(map[string]int)
    for fi5 := 0; fi5 < 15; fi5++ {
        for fi6 := 0; fi6 < 7; fi6++ {
            delete(v4, ((("hello" + "a") + "foo bar") + ("hello" + ("foo bar" + ""))))
            _ = fi6
        }
        fmt.Println(v3, v3)
        _ = fi5
    }
    if (v3 < v3) {
        var v7 int8 = 127
        fmt.Println(v3, v7)
        _ = v7
    } else {
        v3 = 109
    }
    if (0.0 != 0.1) {
        v2 = (2147483647 > -1)
    } else {
        var v8 uint8 = v1
        v9 := GenS1{F0: 0, F1: 2147483648}
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := []int{fi10, fi10, fi10}
        var v12 int16 = -32768
        v4["a"] = (^(fi10))
        _ = fi10
        _ = v11
        _ = v12
    }
    v2 = false
    if true {
        var v13 int16 = 1000
        v14 := []int{}
        v15 := GenS2{F0: 641}
        _ = v13
        _ = v14
        _ = v15
    } else {
        v16 := GenS2{F0: 191}
        switch -182 {
        case 32767:
            for fi17 := 0; fi17 < 16; fi17++ {
                v18 := []int{(-(fi17)), fi17}
                _ = fi17
                _ = v18
            }
        }
        for fi19 := 0; fi19 < 4; fi19++ {
            _ = len(v4)
            _ = fi19
        }
        _ = v16
    }
    v2 = false
    if true {
        switch "a" {
        case "\n\t":
            var v20 string = "a"
            _ = v20
        }
        switch 909 {
        case -25:
            v2 = (-9223372036854775808 > 42)
        case -1:
            v3 = (v3 | 941)
        }
    } else {
        v3 += (v3 & 42)
        for fi21 := 0; fi21 < 20; fi21++ {
            switch ("" + "hello") {
            case "foo bar":
                v4["foo bar"] = fi21
            case "hello":
                var v22 uint32 = 42
                _ = v22
            case "":
                v23 := GenS1{F0: 502, F1: 2147483648}
                _ = v23
            }
            fmt.Println(v1, v3, v3)
            v24 := GenS1{F0: -9223372036854775808, F1: 0}
            _ = fi21
            _ = v24
        }
        switch "\n\t" {
        case "hello":
            fmt.Println(v2, v1)
        default:
            v25 := GenS3{F0: 255, F1: -1}
            _ = v25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}

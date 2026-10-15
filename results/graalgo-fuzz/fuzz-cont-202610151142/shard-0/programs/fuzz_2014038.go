package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 int32 = -1
    var v2 bool = (32767 > 1)
    v3 := make(map[uint32]bool)
    if v2 {
        switch 868 {
        case 1:
            var v4 uint32 = 2147483647
            _ = v4
        case -2147483648:
            v1 = v1
        default:
            for fi5 := 0; fi5 < 13; fi5++ {
                fmt.Println(v2, v1)
                _ = fi5
            }
        }
        switch 778 {
        case 128:
            var v6 uint16 = 65535
            _ = v6
        case -1:
            fmt.Println(v1, v2)
        }
    } else {
        var v7 string = (("hello" + "hello") + ("foo bar" + "foo bar"))
        if false {
            fmt.Println(v7)
        }
        v8 := GenS1{F0: 12}
        _ = v7
        _ = v8
    }
    switch 42 {
    case -891:
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 int16 = -742
            v1 += (v1 * v1)
            _ = fi9
            _ = v10
        }
    }
    var v11 int8 = 70
    switch 732 {
    case 32767:
        fmt.Println(v2)
    case 324:
        for fi12 := 0; fi12 < 19; fi12++ {
            v13 := GenS2{F0: 2147483648}
            var v14 int = -881
            v14 += int(v11)
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    v1 += v1
    var v15 int32 = 2147483647
    for fi16 := 0; fi16 < 19; fi16++ {
        for fi17 := 0; fi17 < 3; fi17++ {
            v1 = -1
            switch 32767 {
            case 796:
                var v18 bool = ((v1 | v15) < (v15 & (v1 | v1)))
                _ = v18
            }
            v2 = false
            _ = fi17
        }
        var v19 int64 = -9223372036854775808
        _ = fi16
        _ = v19
    }
    if v2 {
        v20 := []int{-9223372036854775808, 42, int(v11), -358}
        _ = v20
    }
    if (v2 || v2) {
        v21 := []int{}
        _ = v21
    } else {
        v1 = (v1 + v1)
        var v22 uint32 = 2147483648
        _ = v22
    }
    if v2 {
        fmt.Println(v2, v15)
    }
    var v23 string = (("a" + ("foo bar" + "hello")) + (("hello" + ("foo bar" + "foo bar")) + ""))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v23
    _ = v3
}

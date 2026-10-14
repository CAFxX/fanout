package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int32 }
func genFn0(p0 int64, p1 bool, p2 bool) int8 {
    v7 := GenS1{F0: 4294967295}
    if (127 != 162) {
        fmt.Println(p2, p1, p0)
        if (p1 || (-32768 == 32767)) {
            var v8 uint8 = 42
            v9 := GenS2{F0: 65535, F1: -2147483648}
            v10 := GenS2{F0: 255, F1: -307}
            _ = v10
            _ = v8
            _ = v9
        } else {
            var v11 int64 = (p0 | -1)
            p1 = p2
            _ = v11
        }
    } else {
        p1 = p1
        for fi12 := 0; fi12 < 13; fi12++ {
            fmt.Println(p1, p2)
            _ = fi12
        }
        fmt.Println(p1, p2)
    }
    p2 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 49
}
func genFn1(p0 uint, p1 uint64, p2 uint64) uint8 {
    if (!true) {
        p1 = 42
        var v13 bool = (!false)
        _ = v13
    }
    var v14 uint16 = 32767
    v14 -= v14
    v15 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return 2
}

func main() {
    v1 := []string{"hello", ""}
    var v2 int32 = 2147483647
    v3 := make(map[int]string)
    var v4 bool = true
    v5 := make(map[string]string)
    var v6 int8 = 1
    var v16 int16 = -187
    for fi17 := 0; fi17 < 8; fi17++ {
        for fi18 := 0; fi18 < 9; fi18++ {
            v2 += v2
            fmt.Println(v16, v2)
            _ = fi18
        }
        v6 += genFn0(107, v4, false)
        _ = fi17
    }
    v19 := GenS1{F0: 4294967295}
    var v20 uint = 839
    if v4 {
        var v21 uint8 = 0
        if v4 {
            var v22 uint8 = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 6; fi23++ {
            v6 -= v6
            _ = fi23
        }
        _ = v21
    }
    switch 157 {
    case 1:
        v5[""] = "foo bar"
    default:
        var v24 uint32 = 42
        _ = v24
    }
    for fi25 := 0; fi25 < 10; fi25++ {
        v1 = append(v1, (("hello" + "") + "foo bar"))
        for fi26 := 0; fi26 < 19; fi26++ {
            switch ("hello" + ("" + "")) {
            case "":
                _ = len(v1)
            default:
                var v27 int8 = v6
                _ = v27
            }
            v20 += v20
            _ = fi26
        }
        _ = fi25
    }
    var v28 int64 = 716
    v4 = (465 < -9223372036854775808)
    v4 = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 uint32 }
func genFn0(p0 uint64) int {
    v7 := GenS1{F0: 255}
    if (!(!(42 > 65535))) {
        if (p0 == p0) {
            p0 -= p0
            var v8 float64 = 3.14159
            var v9 int16 = 514
            _ = v8
            _ = v9
        } else {
            p0 += 0
            fmt.Println(p0)
            p0 = p0
        }
        fmt.Println(p0)
        v10 := []int{}
        _ = v10
    }
    switch "\n\t" {
    case "\n\t":
        fmt.Println(p0)
    case "a":
        var v11 int8 = 1
        _ = v11
    default:
        v12 := []int{127}
        _ = v12
    }
    _ = p0
    _ = v7
    return 4294967295
}
func genFn1(p0 uint16, p1 string) int16 {
    p0 -= 42
    p1 = "\n\t"
    fmt.Println(p1, p0)
    var v13 uint32 = 2147483648
    _ = p0
    _ = p1
    _ = v13
    return 1
}

func main() {
    var v1 uint32 = 1
    var v2 string = (((("" + "foo bar") + "") + ("a" + "a")) + "")
    var v3 uint8 = 40
    v4 := make(map[uint32]string)
    var v5 uint64 = 4294967295
    var v6 uint = 2
    var v14 int = 32767
    fmt.Println(v3)
    v2 = ("\n\t" + (v2 + (v2 + v2)))
    v4[65535] = (v2 + ((v2 + ((v2 + (v2 + v2)) + "foo bar")) + (v2 + "\n\t")))
    for fi15 := 0; fi15 < 16; fi15++ {
        for fi16 := 0; fi16 < 19; fi16++ {
            v17 := GenS1{F0: v3}
            var v18 uint = v6
            if true {
                v1 = 4294967295
            }
            _ = fi16
            _ = v17
            _ = v18
        }
        v2 = (((("foo bar" + "hello") + "hello") + (((v2 + ("hello" + v2)) + v2) + "foo bar")) + "a")
        delete(v4, 2147483647)
        _ = fi15
    }
    v14 = genFn0(v5)
    for fi19 := 0; fi19 < 7; fi19++ {
        v2 = (("hello" + ("" + v2)) + "a")
        for fi20 := 0; fi20 < 17; fi20++ {
            v21 := GenS2{F0: "", F1: v1}
            _ = fi20
            _ = v21
        }
        v22 := GenS1{F0: 42}
        _ = fi19
        _ = v22
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        switch 638 {
        case 0:
            v24 := []int{}
            _ = v24
        case 127:
            fmt.Println(v1, v1)
        case 1:
            for fi25 := 0; fi25 < 8; fi25++ {
                var v26 int16 = 42
                var v27 int32 = 0
                v27 = 2147483647
                fmt.Println(v1, v5, v5)
                _ = fi25
                _ = v26
                _ = v27
            }
        }
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

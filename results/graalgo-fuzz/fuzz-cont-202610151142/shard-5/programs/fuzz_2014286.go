package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64 }
func genFn0(p0 int32, p1 int) uint {
    fmt.Println(p0, p1)
    for fi4 := 0; fi4 < 3; fi4++ {
        for fi5 := 0; fi5 < 7; fi5++ {
            switch ("hello" + ("foo bar" + "a")) {
            case "\n\t":
                v6 := GenS1{F0: 255}
                _ = v6
            default:
                v7 := GenS1{F0: 0}
                _ = v7
            }
            var v8 float32 = 0.0
            _ = fi5
            _ = v8
        }
        v9 := GenS2{F0: 576}
        _ = fi4
        _ = v9
    }
    switch "\n\t" {
    case "\n\t":
        v10 := GenS1{F0: 42}
        _ = v10
    case "a":
        fmt.Println(p0)
    case "hello":
        var v11 uint32 = 2
        _ = v11
    }
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{(("a" + (("foo bar" + "hello") + "foo bar")) + ""), ("hello" + "a")}
    v3 := []string{("\n\t" + (((("hello" + "foo bar") + "hello") + "hello") + ""))}
    var v12 int64 = 730
    var v13 bool = (0 <= 4294967295)
    for fi14 := 0; fi14 < 11; fi14++ {
        v13 = v13
        _ = fi14
    }
    v15 := GenS2{F0: 65535}
    fmt.Println(v13, v12)
    switch ("foo bar" + "hello") {
    case "hello":
        v12 += (-(v12))
    case "":
        v13 = (false || v13)
    case "a":
        if v13 {
            v12 = v12
            if true {
                v2 = append(v2, ("" + ("" + "\n\t")))
                v16 := GenS3{F0: (v12 ^ -1)}
                _ = v16
            }
            if false {
                var v17 string = ("hello" + "a")
                fmt.Println(v12, v13, v17)
                var v18 int = -9223372036854775808
                _ = v17
                _ = v18
            } else {
                v19 := GenS1{F0: 264}
                _ = v19
            }
        }
    default:
        v1[128] = 223
    }
    var v20 float32 = 0.1
    for fi21 := 0; fi21 < 10; fi21++ {
        v13 = false
        v12 = 632
        _ = fi21
    }
    if (1 > 100) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v12 -= (v12 >> ((85) & 63))
    }
    v1[-2147483648] = 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v20
    _ = v3
}

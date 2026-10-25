package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint64 }
func genFn0(p0 int32, p1 int64) uint8 {
    v4 := []int{-993, 4294967295, 127, -679}
    fmt.Println(p0, p1)
    var v5 int64 = (^(p1))
    switch 626 {
    case 857:
        v6 := GenS1{F0: "a", F1: 32767, F2: 0}
        _ = v6
    case -402:
        for fi7 := 0; fi7 < 11; fi7++ {
            fmt.Println(p1)
            switch fi7 {
            case 1:
                fmt.Println(v5, v5)
            case 0:
                switch 732 {
                case 2147483647:
                    p0 = p0
                }
            case 128:
                p0 += 473
            default:
                v8 := GenS1{F0: "hello", F1: 469, F2: 302}
                _ = v8
            }
            _ = fi7
        }
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    v3 := []uint32{2, 0}
    var v9 int32 = -1
    var v10 int = 32767
    var v11 int64 = -1
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 int = -9223372036854775808
        for fi14 := 0; fi14 < 19; fi14++ {
            v15 := GenS1{F0: (("hello" + "foo bar") + ((("\n\t" + ("foo bar" + "foo bar")) + (("foo bar" + "\n\t") + (("hello" + ("" + "")) + ("a" + ("a" + (((("\n\t" + (("hello" + "\n\t") + "hello")) + (("" + ("\n\t" + "a")) + ("a" + "hello"))) + "a") + "\n\t")))))) + "hello")), F1: 0, F2: 0}
            if false {
                var v16 int8 = -1
                v9 = v9
                _ = v16
            }
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    v17 := GenS1{F0: "foo bar", F1: 65535, F2: 42}
    fmt.Println(v10)
    delete(v1, 1)
    fmt.Println(v10, v11)
    v9 = 787
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v9
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 string }
func genFn0(p0 uint64) uint {
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 int32 = -633
        var v7 int = (^(fi5))
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint8 = 42
    var v9 string = (("foo bar" + (("a" + "a") + "hello")) + "")
    v10 := []int{-1}
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := []string{("" + ("\n\t" + ("\n\t" + "foo bar"))), "a", ("a" + ((("hello" + ((("a" + ("\n\t" + "")) + ("a" + ("hello" + ("\n\t" + (("hello" + "a") + (("" + ("" + "foo bar")) + "a")))))) + "")) + ("hello" + ((((("foo bar" + ("\n\t" + ("\n\t" + "hello"))) + "") + "\n\t") + "hello") + "hello"))) + "\n\t"))}
    var v2 bool = (42 > 42)
    v3 := []uint32{}
    v4 := make(map[int]uint32)
    v2 = v2
    if v2 {
        v2 = v2
        for fi11 := 0; fi11 < 20; fi11++ {
            fmt.Println(v2, fi11)
            for fi12 := 0; fi12 < 6; fi12++ {
                if v2 {
                    var v13 uint = 2
                    _ = v13
                }
                if (v2 || v2) {
                    v14 := GenS1{F0: "foo bar", F1: 2, F2: "foo bar"}
                    v2 = v2
                    _ = v14
                }
                _ = len(v3)
                _ = fi12
            }
            _ = fi11
        }
    }
    fmt.Println(v2)
    v2 = v2
    for fi15 := 0; fi15 < 2; fi15++ {
        if v2 {
            v16 := GenS1{F0: "\n\t", F1: 2147483647, F2: ("foo bar" + "\n\t")}
            var v17 string = ("foo bar" + "\n\t")
            _ = v16
            _ = v17
        } else {
            var v18 uint64 = 4294967295
            if v2 {
                v4[441] = 114
            }
            _ = v18
        }
        v2 = (18446744073709551615 != 1)
        _ = fi15
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 float32 = 3.14159
        v2 = (!(645 < 42))
        _ = fi19
        _ = v20
    }
    if len(v3) > 0 {
        v3[0] = 0
    }
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}

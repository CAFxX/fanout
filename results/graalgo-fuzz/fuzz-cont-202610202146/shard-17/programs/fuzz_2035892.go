package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 string }

func main() {
    v1 := []string{}
    v2 := make(map[int]string)
    var v3 uint16 = 32767
    v4 := make(map[uint32]uint32)
    for fi5 := 0; fi5 < 4; fi5++ {
        delete(v2, (fi5 - 1))
        var v6 int16 = 335
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: (("hello" + ((("hello" + "foo bar") + "hello") + (("hello" + "") + (("\n\t" + "a") + ("" + ("hello" + "\n\t")))))) + ("foo bar" + "")), F1: 980}
    if ((-332 == 0) || true) {
        fmt.Println(v3)
        v1 = append(v1, "hello")
        switch 4294967295 {
        case 32767:
            var v8 int32 = 2147483647
            _ = v8
        case 42:
            v9 := GenS2{F0: ""}
            _ = v9
        default:
            v3 += 332
        }
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        v11 := GenS2{F0: "hello"}
        var v12 int16 = 0
        for fi13 := 0; fi13 < 13; fi13++ {
            v3 += v3
            _ = fi13
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println(v3)
    v3 -= v3
    var v14 int32 = 0
    if true {
        var v15 uint64 = uint64(v3)
        _ = v15
    } else {
        v4[229] = 729
    }
    v3 -= (v3 ^ 725)
    var v16 string = "a"
    v17 := GenS2{F0: v16}
    v16 = (v16 + v16)
    for fi18 := 0; fi18 < 19; fi18++ {
        if false {
            v3 += v3
        }
        _ = fi18
    }
    _ = len(v1)
    v3 -= 2
    switch (("" + "a") + v16) {
    case "a":
        for fi19 := 0; fi19 < 19; fi19++ {
            var v20 uint32 = 798
            _ = fi19
            _ = v20
        }
    case "\n\t":
        for fi21 := 0; fi21 < 3; fi21++ {
            for fi22 := 0; fi22 < 12; fi22++ {
                delete(v2, fi21)
                _ = fi22
            }
            _ = fi21
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}

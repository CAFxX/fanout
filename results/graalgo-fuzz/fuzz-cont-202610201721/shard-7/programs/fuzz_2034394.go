package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint32; F1 int32 }
func genFn0() int32 {
    var v4 string = ("hello" + ((("a" + ("foo bar" + ("hello" + ("a" + "")))) + "") + ("hello" + "\n\t")))
    v4 = (v4 + (v4 + ((("a" + "hello") + v4) + v4)))
    var v5 uint32 = 470
    if true {
        v6 := []int{-1, -1, 1}
        _ = v6
    }
    var v7 uint = 2
    _ = v4
    _ = v5
    _ = v7
    return -1
}
func genFn1(p0 uint64) bool {
    fmt.Println(p0)
    fmt.Println(p0)
    v8 := GenS1{F0: "foo bar", F1: 0}
    if ("" < ("" + ("" + ((("foo bar" + "hello") + ("\n\t" + "foo bar")) + "\n\t")))) {
        p0 = 18446744073709551615
        v9 := GenS1{F0: "\n\t", F1: 462}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 7; fi10++ {
            switch "hello" {
            case "hello":
                if false {
                    p0 = p0
                    p0 += p0
                    var v11 uint64 = 42
                    _ = v11
                }
            case "a":
                for fi12 := 0; fi12 < 20; fi12++ {
                    var v13 uint16 = 255
                    _ = fi12
                    _ = v13
                }
            }
            v14 := GenS1{F0: (("hello" + "\n\t") + "foo bar"), F1: 763}
            _ = fi10
            _ = v14
        }
        var v15 float32 = -2.25
        var v16 int16 = -32768
        _ = v15
        _ = v16
    }
    _ = p0
    _ = v8
    return (-1 < 1)
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{2147483648}
    var v3 uint16 = 422
    var v17 int8 = 100
    v2 = append(v2, 2147483647)
    for fi18 := 0; fi18 < 8; fi18++ {
        switch fi18 {
        case 2147483647:
            v19 := GenS1{F0: "foo bar", F1: 2}
            _ = v19
        case 487:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 9223372036854775807:
            if true {
                var v20 int = -1
                _ = v20
            } else {
                var v21 uint64 = 2
                if false {
                    if len(v2) > 0 {
                        v2[0] = 232
                    }
                }
                var v22 uint8 = 139
                _ = v21
                _ = v22
            }
        default:
            v17 -= (v17 % ((-128) | 1))
        }
        var v23 int = fi18
        _ = fi18
        _ = v23
    }
    if ("\n\t" < "\n\t") {
        var v24 string = "hello"
        _ = v24
    } else {
        for fi25 := 0; fi25 < 15; fi25++ {
            var v26 uint = 737
            for fi27 := 0; fi27 < 6; fi27++ {
                v3 = 32767
                var v28 int16 = -32768
                _ = fi27
                _ = v28
            }
            var v29 uint64 = 266
            _ = fi25
            _ = v26
            _ = v29
        }
    }
    v3 -= (v3 ^ v3)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    _ = len(v2)
    v30 := GenS2{F0: 4294967295, F1: -811}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v30
}

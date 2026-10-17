package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int16 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint; F1 int32 }
func genFn0() int {
    if ((!true) && (!((!(42 >= 96)) && (-105 == 0)))) {
        var v4 uint32 = 366
        _ = v4
    }
    v5 := GenS1{F0: 461, F1: -541, F2: 0}
    _ = v5
    return 32768
}

func main() {
    v1 := make(map[int]int)
    var v2 string = "a"
    v3 := make(map[int]uint32)
    if true {
        switch ((v2 + ("\n\t" + v2)) + (v2 + "\n\t")) {
        case "":
            for fi6 := 0; fi6 < 4; fi6++ {
                v2 = "a"
                delete(v3, fi6)
                if false {
                    v7 := GenS2{F0: 2, F1: "a"}
                    _ = v7
                } else {
                    var v8 uint64 = 1
                    _ = v8
                }
                _ = fi6
            }
        case "\n\t":
            var v9 string = v2
            _ = v9
        case "foo bar":
            var v10 int64 = -1
            _ = v10
        default:
            for fi11 := 0; fi11 < 20; fi11++ {
                for fi12 := 0; fi12 < 12; fi12++ {
                    var v13 float32 = -2.25
                    _ = fi12
                    _ = v13
                }
                _ = fi11
            }
        }
        var v14 int = 1
        v2 = v2
        _ = v14
    }
    v2 = "foo bar"
    v2 = "hello"
    v15 := GenS1{F0: 0, F1: -9223372036854775808, F2: 797}
    v16 := []int{}
    _ = len(v3)
    v16 = append(v16, -300)
    _ = len(v16)
    for fi17 := 0; fi17 < 7; fi17++ {
        v2 = v2
        _ = fi17
    }
    var v18 uint64 = 18446744073709551615
    for fi19 := 0; fi19 < 3; fi19++ {
        fmt.Println(v2, v2, v18)
        v20 := GenS1{F0: 32767, F1: genFn0(), F2: 0}
        switch v2 {
        case "foo bar":
            var v21 int = fi19
            _ = v21
        case "\n\t":
            fmt.Println(fi19, v18)
        default:
            v22 := GenS3{F0: 65535, F1: 0}
            _ = v22
        }
        _ = fi19
        _ = v20
    }
    v2 = "foo bar"
    v2 = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
}

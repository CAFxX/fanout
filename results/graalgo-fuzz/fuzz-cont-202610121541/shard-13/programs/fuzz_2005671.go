package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int }
func genFn0() uint32 {
    var v5 uint = 795
    v5 += v5
    switch 110 {
    case 184:
        fmt.Println(v5)
    default:
        if false {
            fmt.Println(v5)
        } else {
            fmt.Println(v5)
            for fi6 := 0; fi6 < 20; fi6++ {
                switch (fi6 % ((128) | 1)) {
                case 127:
                    switch fi6 {
                    case -815:
                        fmt.Println(v5, v5)
                    case 9223372036854775807:
                        var v7 uint16 = 2
                        _ = v7
                    }
                default:
                    var v8 uint32 = 2
                    _ = v8
                }
                if true {
                    fmt.Println(v5)
                    v9 := []int{fi6, fi6, (-(fi6)), fi6}
                    _ = v9
                }
                var v10 int8 = 0
                _ = fi6
                _ = v10
            }
            var v11 uint32 = 0
            _ = v11
        }
    }
    _ = v5
    return 396
}

func main() {
    v1 := []string{"hello"}
    var v2 uint8 = 255
    v3 := make(map[uint32]int)
    var v4 int = 0
    v12 := GenS1{F0: ""}
    v13 := GenS1{F0: "a"}
    var v14 float32 = (3.14159 + -2.25)
    v14 += v14
    v2 += (v2 - 1)
    v4 += (^(v4))
    switch v4 {
    case 2147483647:
        v2 = v2
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        _ = len(v1)
        var v16 bool = (3.14159 <= 1.5)
        if true {
            for fi17 := 0; fi17 < 2; fi17++ {
                v18 := GenS2{F0: (fi17 & fi17)}
                var v19 uint64 = 4294967295
                _ = fi17
                _ = v18
                _ = v19
            }
            var v20 int = (fi15 * (^(v4)))
            var v21 string = "\n\t"
            _ = v20
            _ = v21
        }
        _ = fi15
        _ = v16
    }
    var v22 int = 2147483647
    if (true || false) {
        v3[genFn0()] = -1
        delete(v3, 65535)
    }
    fmt.Println(v2)
    for fi23 := 0; fi23 < 19; fi23++ {
        switch ("a" + "hello") {
        case "hello":
            for fi24 := 0; fi24 < 11; fi24++ {
                var v25 float32 = v14
                var v26 uint16 = 42
                switch (("foo bar" + "hello") + "a") {
                case "foo bar":
                    v22 = 128
                case "":
                    var v27 int32 = -932
                    _ = v27
                default:
                    v25 = (v14 / float32(1e10))
                }
                _ = fi24
                _ = v25
                _ = v26
            }
        default:
            v2 = v2
        }
        v4 = v4
        _ = fi23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}

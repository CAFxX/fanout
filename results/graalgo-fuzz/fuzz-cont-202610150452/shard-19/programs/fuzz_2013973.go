package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int64 }
func genFn0(p0 string, p1 string, p2 string) int32 {
    p1 = p1
    var v7 float64 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    var v1 uint16 = 0
    var v2 int8 = -1
    var v3 float64 = 1.5
    v4 := []uint32{42, 42}
    v5 := make(map[string]string)
    var v6 int8 = (^(v2))
    var v8 int = 169
    _ = len(v5)
    v9 := GenS1{F0: "foo bar"}
    for fi10 := 0; fi10 < 19; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v12 string = (("hello" + "foo bar") + "hello")
            _ = fi11
            _ = v12
        }
        if false {
            for fi13 := 0; fi13 < 11; fi13++ {
                if (251 == 128) {
                    v14 := GenS2{F0: 192}
                    _ = v14
                } else {
                    v3 = (v3 / 1.5)
                    fmt.Println(v6, fi13, v6)
                }
                _ = fi13
            }
            if (((((v3 + 3.14159) + 3.14159) < 0.1) || (!("hello" == ("foo bar" + "")))) || (0 < 32767)) {
                v2 = 0
                var v15 float32 = 1e10
                _ = v15
            } else {
                var v16 uint8 = 42
                fmt.Println(v1, v2, v1)
                _ = v16
            }
        }
        if (1 < -44) {
            var v17 bool = (0 >= 4294967295)
            var v18 uint16 = (^(v1))
            fmt.Println(fi10, v8, v2)
            _ = v17
            _ = v18
        } else {
            if len(v4) > 0 {
                v4[0] = 2
            }
            v3 = -2.25
        }
        _ = fi10
    }
    fmt.Println(v8)
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        var v19 string = (("\n\t" + "hello") + "\n\t")
        _ = v19
    }
    v20 := GenS3{F0: 904, F1: -1}
    var v21 string = ("\n\t" + "\n\t")
    for fi22 := 0; fi22 < 20; fi22++ {
        for fi23 := 0; fi23 < 7; fi23++ {
            v21 = v21
            _ = fi23
        }
        v24 := GenS1{F0: ""}
        _ = fi22
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}

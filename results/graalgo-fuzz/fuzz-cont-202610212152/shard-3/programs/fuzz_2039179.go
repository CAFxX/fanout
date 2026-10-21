package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint8; F1 uint; F2 string }
type GenS3 struct { F0 int32 }

func main() {
    var v1 string = "foo bar"
    var v2 uint32 = 0
    var v3 int = 1
    v4 := make(map[string]uint32)
    v5 := make(map[string]bool)
    var v6 uint16 = 0
    v6 = (^(v6))
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 int32 = 2147483647
        _ = fi7
        _ = v8
    }
    v1 = ((((((((("hello" + "hello") + v1) + v1) + (("hello" + (v1 + (v1 + "\n\t"))) + v1)) + "\n\t") + "a") + v1) + (((v1 + "foo bar") + v1) + (v1 + (v1 + "a")))) + ("hello" + "foo bar"))
    var v9 int32 = 0
    switch v1 {
    case "hello":
        v10 := v5["hello"]
        _ = v10
    }
    var v11 uint = 2
    fmt.Println(v2)
    var v12 uint64 = 4294967295
    for fi13 := 0; fi13 < 19; fi13++ {
        v14 := GenS1{F0: v11, F1: 832}
        if true {
            for fi15 := 0; fi15 < 15; fi15++ {
                v2 = (v2 ^ 2147483648)
                switch (-(fi13)) {
                case 4294967295:
                    fmt.Println(v11, v1)
                case 0:
                    if false {
                        var v16 int = 0
                        if (1000 <= -904) {
                            delete(v4, (("\n\t" + v1) + "foo bar"))
                            v6 = (v6 / (((^(v6))) | 1))
                            var v17 uint16 = 0
                            _ = v17
                        }
                        fmt.Println(v16, v16, fi15)
                        _ = v16
                    }
                case -463:
                    v6 -= (v6 & v6)
                }
                v18 := GenS1{F0: v11, F1: 1}
                _ = fi15
                _ = v18
            }
            var v19 bool = (!true)
            _ = v19
        }
        _ = fi13
        _ = v14
    }
    fmt.Println(v12)
    var v20 uint16 = (v6 / ((v6) | 1))
    v5[v1] = (true && (2147483647 <= v9))
    for fi21 := 0; fi21 < 19; fi21++ {
        v22 := GenS2{F0: 0, F1: v11, F2: v1}
        if true {
            var v23 float32 = 0.1
            for fi24 := 0; fi24 < 3; fi24++ {
                v25 := GenS3{F0: 441}
                var v26 uint64 = 1
                _ = fi24
                _ = v25
                _ = v26
            }
            v11 -= uint(v2)
            _ = v23
        }
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 int32 }

func main() {
    var v1 int = -788
    var v2 string = "foo bar"
    v3 := make(map[string]int)
    var v4 int64 = 125
    var v5 uint32 = 2
    var v6 uint32 = v5
    fmt.Println(v2)
    var v7 uint8 = 49
    v8 := GenS1{F0: 1000}
    for fi9 := 0; fi9 < 18; fi9++ {
        v2 = "\n\t"
        if false {
            for fi10 := 0; fi10 < 3; fi10++ {
                _ = len(v3)
                _ = fi10
            }
            fmt.Println(fi9)
            var v11 int64 = v4
            _ = v11
        } else {
            v6 += (v5 ^ v6)
            v7 += v7
            for fi12 := 0; fi12 < 8; fi12++ {
                v13 := GenS2{F0: 42}
                switch (-(fi9)) {
                case -9223372036854775808:
                    if ((v4 - (v4 << ((1) & 63))) <= (v4 + v4)) {
                        var v14 float32 = float32(0.1)
                        v15 := GenS3{F0: v2}
                        var v16 uint16 = 32767
                        var v17 bool = false
                        _ = v14
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                default:
                    if true {
                        if (42 <= 32767) {
                            v18 := []int{(v1 | fi12), fi9}
                            _ = v18
                        } else {
                            fmt.Println(v5, v2, fi9)
                            var v19 uint = 18446744073709551615
                            _ = v19
                        }
                        fmt.Println(v7, v1)
                    } else {
                        switch ("a" + "") {
                        case "a":
                            v7 -= v7
                        case "hello":
                            var v20 uint16 = 197
                            _ = v20
                        }
                        var v21 uint32 = v6
                        var v22 int16 = 939
                        _ = v21
                        _ = v22
                    }
                }
                if (!(true && (("" + "hello") < v2))) {
                    fmt.Println(fi9)
                    fmt.Println(v6, v2, v6)
                }
                _ = fi12
                _ = v13
            }
        }
        _ = fi9
    }
    delete(v3, v2)
    var v23 int64 = (v4 + (v4 >> ((2) & 63)))
    for fi24 := 0; fi24 < 1; fi24++ {
        v1 += (^(v1))
        v25 := v3[((v2 + (("foo bar" + v2) + "foo bar")) + ("foo bar" + "\n\t"))]
        _ = fi24
        _ = v25
    }
    v26 := GenS3{F0: v2}
    v27 := GenS4{F0: (1 == 0), F1: -362}
    fmt.Println(v7, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

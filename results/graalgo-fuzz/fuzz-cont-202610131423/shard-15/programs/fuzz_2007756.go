package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool }

func main() {
    v1 := []int{-9223372036854775808, 42, 368}
    var v2 uint64 = 4294967295
    var v3 uint8 = 0
    v4 := []string{"\n\t", "a", "hello", "\n\t"}
    v5 := make(map[int]int)
    var v6 uint64 = 4294967295
    v7 := GenS1{F0: -509}
    var v8 bool = false
    fmt.Println(v6)
    var v9 string = "hello"
    for fi10 := 0; fi10 < 6; fi10++ {
        v8 = (false || (-1 > 0))
        v11 := GenS2{F0: v8}
        _ = fi10
        _ = v11
    }
    fmt.Println(v2)
    for fi12 := 0; fi12 < 5; fi12++ {
        for fi13 := 0; fi13 < 13; fi13++ {
            if (!(42 < 0)) {
                switch (v9 + v9) {
                case "":
                    v6 = 18446744073709551615
                case "\n\t":
                    v5[fi12] = (-(fi13))
                }
                if v8 {
                    switch "\n\t" {
                    case "":
                        var v14 uint32 = 65535
                        _ = v14
                    case "hello":
                        if ((42 >= 0) && v8) {
                            v8 = v8
                            v15 := GenS2{F0: (821 == 445)}
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v16 uint32 = 172
                            _ = v15
                            _ = v16
                        }
                    case "a":
                        v17 := GenS1{F0: 881}
                        _ = v17
                    }
                    fmt.Println(v3, v6, v6)
                    v5[(fi13 | 127)] = (-(fi12))
                } else {
                    fmt.Println(fi12)
                }
                var v18 string = v9
                _ = v18
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v19 := GenS1{F0: -57}
                _ = v19
            }
            _ = fi13
        }
        fmt.Println(v2, fi12)
        _ = fi12
    }
    var v20 int16 = -32768
    v1 = append(v1, -9223372036854775808)
    v21 := GenS2{F0: (false || false)}
    v8 = ((!true) || true)
    v4 = append(v4, "hello")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

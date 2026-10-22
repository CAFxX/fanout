package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }

func main() {
    var v1 uint8 = 144
    var v2 float64 = 1.5
    v3 := []int{-9223372036854775808}
    var v4 float32 = -2.25
    fmt.Println(v2)
    fmt.Println(v4, v4, v2)
    for fi5 := 0; fi5 < 8; fi5++ {
        v6 := GenS1{F0: true}
        if (true || (42 < 1)) {
            _ = len(v3)
        } else {
            switch (-(fi5)) {
            case 4294967295:
                var v7 int32 = -2147483648
                _ = v7
            }
        }
        v3 = append(v3, (fi5 / (((-(fi5))) | 1)))
        _ = fi5
        _ = v6
    }
    v8 := GenS2{F0: -285}
    _ = len(v3)
    v2 = 3.14159
    if (!(42 < 42)) {
        var v9 int8 = 67
        var v10 uint64 = 18446744073709551615
        _ = v10
        _ = v9
    } else {
        v4 = 0.0
    }
    var v11 uint8 = (v1 / (((v1 & (v1 - 255))) | 1))
    var v12 uint64 = 1
    if false {
        v4 = v4
        v12 = 632
    } else {
        if (true && (322 > 196)) {
            switch ("\n\t" + "hello") {
            case "":
                if (519 <= 2) {
                    var v13 uint = 1
                    if len(v3) > 0 {
                        v3[0] = 0
                    }
                    v14 := GenS1{F0: false}
                    _ = v13
                    _ = v14
                } else {
                    var v15 bool = (2147483647 >= 42)
                    v3 = append(v3, -9223372036854775808)
                    if (!(357 <= 42)) {
                        switch -639 {
                        case 128:
                            var v16 int32 = 42
                            _ = v16
                        }
                    } else {
                        var v17 uint8 = v11
                        if (-2.25 < v2) {
                            v15 = v15
                        } else {
                            var v18 uint32 = 1
                            _ = v18
                        }
                        _ = v17
                    }
                    _ = v15
                }
            }
            v1 -= (^(v11))
        }
    }
    v12 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}

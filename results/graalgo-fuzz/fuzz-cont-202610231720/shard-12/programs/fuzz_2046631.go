package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 2147483647
    var v3 bool = false
    var v4 uint32 = 0
    var v5 uint32 = 2147483647
    v3 = (!false)
    var v6 uint64 = 4294967295
    var v7 uint64 = (v6 % ((v6) | 1))
    if false {
        v8 := GenS1{F0: 127, F1: 153}
        if v3 {
            v2 += v2
            v2 = (v5 << ((42) & 63))
            fmt.Println(v6, v3)
        }
        _ = v8
    }
    v9 := GenS2{F0: -234}
    for fi10 := 0; fi10 < 16; fi10++ {
        switch (("foo bar" + "\n\t") + (("a" + ("a" + ("hello" + "a"))) + ("hello" + ""))) {
        case "hello":
            if ((70 < -73) && v3) {
                switch fi10 {
                case -1:
                    v11 := GenS2{F0: -1}
                    _ = v11
                }
                if (42 < 2147483647) {
                    v4 = v4
                    v4 -= 2147483648
                } else {
                    v6 = 1
                    for fi12 := 0; fi12 < 7; fi12++ {
                        if (!v3) {
                            if ((127 == uint8(v5)) && v3) {
                                var v13 float32 = 0.1
                                _ = v13
                            }
                        }
                        v14 := GenS2{F0: 0}
                        _ = fi12
                        _ = v14
                    }
                    v1[""] = "foo bar"
                }
                v2 += v2
            } else {
                if ((float32((((3.14159 + 3.14159) * 0.0) / (((0.0 - 1e10) - (0.1 - 0.0)) / -2.25))) + 0.0) == (0.1 * (1.5 * (-2.25 - (1.5 + float32(1e10)))))) {
                    v7 -= v7
                }
                v15 := GenS1{F0: 255, F1: -32768}
                _ = v15
            }
        case "\n\t":
            var v16 int16 = 32767
            _ = v16
        case "":
            _ = len(v1)
        }
        _ = fi10
    }
    var v17 int16 = -882
    fmt.Println(v6, v7, v2)
    var v18 int16 = v17
    var v19 int64 = 0
    v20 := []int{-9223372036854775808, 32767, 4294967295, 340}
    var v21 uint = 255
    v19 = 1
    var v22 float32 = 1.5
    var v23 uint = v21
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}

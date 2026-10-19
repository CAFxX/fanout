package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int64 }
type GenS2 struct { F0 int8; F1 uint32; F2 string }

func main() {
    v1 := []int{}
    v2 := make(map[int]string)
    var v3 uint64 = 2
    v4 := make(map[int]string)
    v3 = 18446744073709551615
    if ((2 >= 239) || (127 >= -3)) {
        for fi5 := 0; fi5 < 2; fi5++ {
            var v6 string = "hello"
            _ = fi5
            _ = v6
        }
        switch "hello" {
        case "hello":
            if (128 < 255) {
                for fi7 := 0; fi7 < 17; fi7++ {
                    fmt.Println(fi7)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    fmt.Println(fi7, fi7)
                    _ = fi7
                }
                if (!(2147483648 == 600)) {
                    if (false || (0.1 == 0.1)) {
                        delete(v2, 32767)
                        v8 := GenS1{F0: 0, F1: 18446744073709551615, F2: -9223372036854775808}
                        _ = v8
                    }
                    var v9 int8 = 1
                    v1 = append(v1, 347)
                    _ = v9
                } else {
                    v10 := GenS1{F0: 0, F1: 0, F2: 573}
                    var v11 string = (("foo bar" + "foo bar") + "foo bar")
                    var v12 string = v11
                    _ = v10
                    _ = v11
                    _ = v12
                }
            }
        case "a":
            v3 += 2
        case "\n\t":
            for fi13 := 0; fi13 < 14; fi13++ {
                var v14 int8 = -1
                _ = fi13
                _ = v14
            }
        }
    }
    var v15 uint8 = 1
    var v16 int64 = 1
    var v17 uint16 = 0
    if len(v1) > 0 {
        v1[0] = 1
    }
    v3 -= (^(v3))
    var v18 int = -9223372036854775808
    v19 := GenS2{F0: 76, F1: 2147483648, F2: "\n\t"}
    v18 += int(v15)
    v3 = uint64(v17)
    v4[32767] = "a"
    var v20 string = (("a" + ("\n\t" + ("" + ("hello" + "")))) + "a")
    var v21 float32 = (0.1 - (-2.25 / 1.0))
    v22 := v2[(^(v18))]
    v17 += 449
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}

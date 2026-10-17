package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }
func genFn0() string {
    v6 := []int{800, -9223372036854775808, -470, -2147483648}
    switch 33 {
    case 4294967295:
        var v7 string = ("\n\t" + "foo bar")
        _ = v7
    case -1:
        if (588 <= 0) {
            var v8 float64 = 1.5
            if len(v6) > 0 {
                v6[0] = 128
            }
            v8 += 0.1
            _ = v8
        }
    case 1:
        var v9 uint32 = 2
        _ = v9
    }
    var v10 int32 = 1
    v11 := GenS1{F0: 453}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = v10
    _ = v11
    _ = v6
    return "\n\t"
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]bool)
    var v3 int16 = -1
    v4 := []string{""}
    v5 := make(map[string]bool)
    v3 = -444
    v3 = (^(v3))
    if true {
        for fi12 := 0; fi12 < 11; fi12++ {
            switch (fi12 * -107) {
            case -2147483648:
                var v13 int32 = 617
                _ = v13
            default:
                var v14 bool = (!(!false))
                _ = v14
            }
            var v15 float64 = 3.14159
            _ = fi12
            _ = v15
        }
        switch 484 {
        case 9223372036854775807:
            v16 := GenS2{F0: true}
            _ = v16
        }
        if (!false) {
            var v17 uint = 18446744073709551615
            _ = v17
        } else {
            var v18 int16 = 1
            v18 -= 1000
            _ = v18
        }
    }
    v3 = (v3 - (v3 >> ((122) & 63)))
    v3 = v3
    var v19 uint16 = 65535
    delete(v1, "\n\t")
    v19 = v19
    v3 += -631
    switch 128 {
    case -540:
        var v20 uint8 = 129
        _ = v20
    case 0:
        v3 = v3
    default:
        fmt.Println(v3, v19)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v3, v19)
    v4 = append(v4, "a")
    v3 += v3
    v21 := GenS3{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
